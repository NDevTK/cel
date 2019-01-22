# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import collections
import logging


# TODO: Implement NewHostProvider which creates new GCP projects.
class SimpleHostProvider:

  def __init__(self, hosts):
    self._hosts = collections.deque(hosts)

  def TryAcquire(self):
    """Tries to acquire a host that can be used for tests.

    SimpleHostProvider always succeeds as long as there are hosts left to use.

    Variants can take an arbitrary long time to return a host file, but must
    ensure the caller has exclusivity over the resources in that file.

    Returns:
      Path to the host file if it succeeds. None if it fails.
    """
    try:
      host = self._hosts.popleft()
      logging.info("Acquiring host '%s'" % host)
      return host
    except IndexError:
      return None

  def Release(self, host):
    """Releases resources associated with the host file."""
    logging.info("Releasing host '%s'" % host)
    self._hosts.append(host)

  def SetNewAvailableHostCallback(self, callback):
    """Not implemented for SimpleHostProvider."""
    pass

  def Size(self):
    """Returns the number of hosts left in this provider."""
    return len(self._hosts)
