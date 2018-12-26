# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import glob
import os
from test.infra.core.test_registry import TestRegistry


class ArgsParser:

  @staticmethod
  def ParseTestsArg(tests):
    """Parses a user supplied --tests arguments.

    Supports both individual classes and prefixes.

    e.g. --tests test.tests.active_directory.*
         --tests test.tests.TestA;test.tests.TestB;...

    Returns:
      The list of EnterpriseTestCase class names (strings) found.

    Raises:
      ValueError: One of the CSV tokens doesn't match any test case class.
    """
    testsToRun = []

    for pattern in tests.split(';'):
      if pattern.endswith('*'):
        results = TestRegistry.FindAll(pattern[:-1])
        if len(results) == 0:
          message = 'No test found that matches "%s".' % pattern
          raise ValueError(message)
        testsToRun += results
      else:
        result = TestRegistry.Find(pattern)
        if result == None:
          message = 'No test found with the name "%s".' % pattern
          raise ValueError(message)
        testsToRun += [result]

    return testsToRun

  @staticmethod
  def ParseHostsArg(hosts):
    """Parses a user supplied --hosts arguments.

    Supports both individual files and directories.

    e.g. --hosts ./my/dir/with/host/files
         --hosts test_a.host.textpb;test_b.host.textpb;...

    Returns:
      The list of host files (strings) found.

    Raises:
      ValueError: One of the CSV tokens doesn't match any existing host file.
    """
    hostFilesFound = []

    hostsParts = hosts.split(';')

    for location in hosts.split(';'):
      if not os.path.exists(location):
        message = 'Path does not exist: %s.' % location
        raise ValueError(message)

      if os.path.isfile(location):
        hostFilesFound.append(location)
      elif os.path.isdir(location):
        path = os.path.join(location, '*.host.textpb')
        hostFilesInDirectory = glob.glob(path)

        if len(hostFilesInDirectory) == 0:
          message = 'Could not find *.host.textpb files in dir: %s' % location
          raise ValueError(message)

        hostFilesFound += hostFilesInDirectory
      else:
        message = 'Path is not a file or a directory: %s' % location
        raise ValueError(message)

    return hostFilesFound
