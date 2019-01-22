# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import collections
from datetime import datetime
import logging
import os
from test.infra.gcp import StorageBucket
import threading
import time
from googleapiclient.errors import HttpError


class SharedHostProvider:
  """Manages a list of gcp projects used by multiple runners at once.

  This HostProvider is tasked with handling GCP projects that can be used by
  multiple test runners at the same time. To do that, it uses GCP storage objs
  to mark projects as "busy" with a .locks/ object. Other runners use those
  objects to know when a project is busy.

  This HostProvider can also notify another object when busy hosts are later
  marked as available (via SetNewAvailableHostCallback).

  Hosts returned by this HostProvider can be safely used as exclusive hosts.
  """

  def __init__(self, hosts, bucketName):
    self._storage = StorageBucket(bucketName)
    self._candidateHosts = collections.deque(hosts)

    # This is all the hosts that we have locked.
    # Mapping of (host => lockGenerationId)
    self._lockedHosts = {}

    # This is all the hosts that are locked by somebody else.
    # Mapping of (host => True)
    self._unavailableHosts = {}

    # Start background tasks now even if there's nothing to do yet.
    self._bgThread = SharedHostProviderThread(self)
    self._bgThread.start()

  def TryAcquire(self):
    """Tries to acquire/lock a host from the list of candidate hosts.

    Returns:
      A host that has been "locked", or None if all hosts are being used.
      We can safely assume hosts returned here are exclusive to this test run.
    """
    while True:
      try:
        host = self._candidateHosts.popleft()
        self._AcquireHostOrThrow(host)
        logging.debug("Acquired lock for %s" % host)
        return host
      except HostAlreadyLockedError as e:
        logging.debug("Failed to acquire the lock for %s: %s" % (host, e))
        self._unavailableHosts[host] = True
      except IndexError:
        # We have no candidate hosts left
        return None

  def Release(self, host):
    """Releases the host by deleting the lock object for it.

    This marks the host as available for other test runners to use and
    should only be used when we no longer have more tests to run.
    """
    try:
      self._storage.DeleteObject(
          self._GetLockName(host), self._lockedHosts[host])
    except Exception as e:
      logging.warning("Failed to release the lock for %s: %s" % (host, e))

    del self._lockedHosts[host]
    self._candidateHosts.append(host)

  def SetNewAvailableHostCallback(self, callback):
    """Registers a callback to notify when a previously busy host is unlocked.

    This callback will be called from a different thread.
    """
    self._bgThread._callback = callback

  def Size(self):
    """Returns the number of hosts left in this provider."""
    return len(self._candidateHosts)

  def _AcquireHostOrThrow(self, host):
    """Tries to acquire a lock for given host.

    Returns:
      Host if succesfully locked.

    Raises:
      HostAlreadyLockedError: If someone else already locked this project.
      (...): Any other exceptions raised by StorageBucket.WriteObject().
    """
    lockName = self._GetLockName(host)
    try:
      self._lockedHosts[host] = self._storage.WriteObject(
          lockName, self._GetLockContent(host), 0)
    except HttpError as e:
      # HttpError: <HttpError 412 when requesting
      #             http...&ifGenerationMatch=0 returned "Precondition Failed">
      if e.resp.status == 412:
        raise HostAlreadyLockedError("Host %s is already locked." % host)
      else:
        raise

  def _GetLockName(self, host):
    """Returns the fullname of the lock to use to lock a host.

    Conceptually, the lock name is the name of the compute project but we rely
    on our internal convention that the hsot file name is the project name to
    avoid having to read/parse that file.
    """
    filename = os.path.basename(host)
    project = filename.split('.')[0]
    return ".locks/%s.lock" % project

  def _GetLockContent(self, host):
    return 'Project locked.'

  def _FindAvailableHosts(self, currentLocks, hosts):
    """Finds potentially available hosts from the lists passed as argument.

    Does NOT lock hosts. Caller must still lock them before using them.

    Args:
      currentLocks: The list of lock objects currently in the storage bucket.
        We expect items to be Storage Objects (e.g. from bucket.objects.list):
        https://cloud.google.com/storage/docs/json_api/v1/objects#resource
      hosts: The list of `unavailable` hosts to check against currentLocks.

    Returns:
      A list of available hosts.
    """
    lockNames = [lock['name'] for lock in currentLocks]

    availableHosts = []
    for host in hosts:
      if self._GetLockName(host) not in lockNames:
        availableHosts.append(host)

    return availableHosts

  def _CleanStaleLocks(self, currentLocks, staleTimeout=60 * 60):
    """Finds stale locks and deletes them.

    We assume a lock is stale/unused if it's been untouched for over an hour.

    Args:
      currentLocks: The list of lock objects currently in the storage bucket.
      staleTimeout: The number of seconds after which to consider a lock stale.

    Returns:
      The updated list of active locks.
    """
    activeLocks = []
    for lock in currentLocks:
      # Skip the folder
      if lock['name'] == '.locks/':
        continue

      lastUpdate = datetime.strptime(lock['updated'], "%Y-%m-%dT%H:%M:%S.%fZ")
      delta = datetime.utcnow() - lastUpdate
      if delta.total_seconds() > staleTimeout:
        logging.debug(
            "Found stale lock (%s) (%s). Deleting." % (lock['name'], delta))
        try:
          self._storage.DeleteObject(lock['name'], lock['generation'])
        except Exception as e:
          logging.warning(
              "Failed to delete stale lock for %s: %s" % (lock['name'], e))
          activeLocks.append(lock)
      else:
        activeLocks.append(lock)

    return activeLocks

  def _RefreshAllLocks(self):
    """Refreshes all active lock objects.

    Test runners will ignore lock objects that are older than an hour,
    so it's our responsibility to call this every so often (30 mins) to ensure
    we don't lose exclusivity on a GCP project.
    """
    for host in self._lockedHosts:
      try:
        genId = self._storage.WriteObject(
            self._GetLockName(host), self._GetLockContent(host),
            self._lockedHosts[host])
        self._lockedHosts[host] = genId
        logging.debug("Refreshed lock for %s" % host)
      except Exception as e:
        logging.warning("Failed to refresh lock for %s: %s" % (host, e))


class SharedHostProviderThread(threading.Thread):
  """Runs background tasks for SharedHostProvider.

  It's responsible for refreshing active locks and deleting stale ones.
  Optional: Can also watch for new available hosts.

  This thread is started when a SharedHostProvider is initialized and ends with
  the process.
  """

  def __init__(self, provider, pollInterval=30, lockRefreshInterval=20 * 60):
    """Initializes the background thread for SharedHostProvider.

    The background thread does the following tasks:
      - Notify callback when a busy host becomes available
        (see SetNewAvailableHostCallback)
      - Refresh the host locks (see _RefreshAllLocks)

    Args:
      pollInterval: The interval (in seconds) to wait between polls.
      lockRefreshInterval: The interval (in seconds) between lock refreshes.
    """
    threading.Thread.__init__(self)
    self.daemon = True

    self._provider = provider
    self._callback = None

    self._pollInterval = pollInterval
    self._lockRefreshInterval = lockRefreshInterval

  def run(self):
    p = self._provider
    lastRefresh = -1
    while True:
      # SharedHostProvider assumes that locks older than an hour are stale. We
      # will refresh active locks every 30 mins to ensure they are not stolen.
      if lastRefresh + self._lockRefreshInterval < time.time():
        self._provider._RefreshAllLocks()
        lastRefresh = time.time()

      currentLocks = self._provider._storage.ListObjects('.locks/')

      # Clean stale locks (if any).
      currentLocks = self._provider._CleanStaleLocks(currentLocks)

      # Look for available hosts and try to lock them.
      if self._callback:
        self._LookForAvailableHosts(currentLocks)

      time.sleep(self._pollInterval)

  def _LookForAvailableHosts(self, currentLocks):
    unavailableHosts = self._provider._unavailableHosts
    candidates = self._provider._FindAvailableHosts(currentLocks,
                                                    unavailableHosts)

    callback = self._callback
    while callback and len(candidates) > 0:
      host = candidates.pop()
      try:
        if callback:
          self._provider._AcquireHostOrThrow(host)
          del unavailableHosts[host]
          callback(host)
      except Exception as e:
        logging.debug(
            "_LookForAvailableHosts: Skipping %s because of exception: %s" %
            (host, e))

      callback = self._callback


class HostAlreadyLockedError(Exception):
  pass
