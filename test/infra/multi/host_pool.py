# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import collections
import logging
import threading


class HostPool:
  """Manages the available hosts for this test run.

  All hosts here are assumed to be safe to use (exclusive).

  Attributes:
    _provider: Provides new hosts and ensures that they are exclusive.
    _maxSize: The maximum number of hosts we can have in the pool.
    _available: A list of host files that are available for a test.
    _busy: A list of host files that are currently under test.
  """

  def __init__(self, provider, maxSize=10):
    self._provider = provider
    self._maxSize = maxSize

    # (_available, _busy) can be mutated from multiple threads (see below).
    self._available = collections.deque()
    self._busy = {}

    # Thread must use _hostsLock to read or mutate _available or _busy.
    self._hostsLock = threading.RLock()

    # _hostAvailableEvent should be used to wait on _available when it's empty.
    self._hostAvailableEvent = threading.Event()

    # Set when we have no more tests to run. Available hosts must be released.
    self._releaseAvailableHosts = False

  def Size(self):
    """Returns the total number of hosts (available and busy) in the pool."""
    with self._hostsLock:
      return len(self._available) + len(self._busy)

  def TakeOrWaitForHost(self):
    """Returns an available host to run a test on.

    It can find one in different ways:
      1) By looking in _available (if a previous test just completed)
      2) By asking HostProvider (if we are under the pool's max size)
      3) By waiting for a new host to be marked as available (1)

    It's assumed this is only called from the main thread because it does not
    have any concept of priority when waiting for (3).
    """
    if not isinstance(threading.current_thread(), threading._MainThread):
      logging.warning("TakeOrWaitForHost called from an unexpected thread.")

    with self._hostsLock:
      if len(self._available) > 0:
        return self._PopAvailableHost()

    if self.Size() < self._maxSize and self._provider.Size() > 0:
      host = self._provider.TryAcquire()
      if host != None:
        self._busy[host] = True
        return host

    return self._WaitForHost()

  def _WaitForHost(self):
    """Waits until it finds an available host.

    This function blocks until another thread adds an available host and
    flips _hostAvailableEvent and returns that host.
    """
    logging.debug("Waiting for a new available host.")
    while True:
      # wait() with no timeout blocks KeyboardInterrupt (CTRL+C) in Python 2.7,
      # but doesn't if one is passed. Let's wait 100h. (Python3 never blocks)
      if not self._hostAvailableEvent.wait(60 * 60 * 100):
        logging.warning("WaitForHost has been waiting for 100 hours! (???)")
        continue

      with self._hostsLock:
        if len(self._available) > 0:
          return self._PopAvailableHost()

      # Should not happen. Main thread is the only TakeOrWaitForHost caller.
      logging.warning("Someone else took the last available host.")

  def ReleaseAllAvailableHosts(self):
    """Releases all current and future available hosts."""
    self._releaseAvailableHosts = True

    with self._hostsLock:
      while len(self._available) > 0:
        host = self._available.popleft()
        self._provider.Release(host)

  def AddAvailableHost(self, host):
    """This marks a host as available for TakeOrWaitForHost.

    This is called by test threads when they complete and no longer need it.
    """
    with self._hostsLock:
      if host in self._busy:
        del self._busy[host]

      # If we don't need the host anymore (no tests in the queue), release it.
      if self._releaseAvailableHosts:
        logging.debug("We don't need host '%s' anymore. Releasing it." % host)
        self._provider.Release(host)
        return

      # Mark the host as available and unblock TakeOrWaitForHost
      logging.debug("Host '%s' is marked as available." % host)
      self._available.append(host)
      self._hostAvailableEvent.set()

  def _PopAvailableHost(self):
    with self._hostsLock:
      host = self._available.popleft()
      self._busy[host] = True

      if len(self._available) == 0:
        self._hostAvailableEvent.clear()

      return host
