# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import os
import subprocess
import unittest


# We inherit from unittest.TestCase so that we get all those assertXXX()
# methods. Note that the way our tests work is different from unittest,
# so please do not use anything in unittest.TestCase except those assert
# methods.
class EnterpriseTestCase(unittest.TestCase):
  """Base class for tests that provides test hooks and resources."""

  def __init__(self, environment):
    logging.info('Initialize Test=%s with %s' % (self.__class__, environment))
    super(EnterpriseTestCase, self).__init__()
    self.clients = environment.clients
    self.gsbucket = environment.gsbucket

  # this method is here to please unittest.
  def runTest(self):
    raise "Please use ./test.py to run this test."

  @staticmethod
  def GetTestMethods(_class):
    return EnterpriseTestCase._getMethods(_class, 'IS_TEST_METHOD')

  @staticmethod
  def GetBeforeAllMethods(_class):
    return EnterpriseTestCase._getMethods(_class, 'IS_BEFORE_ALL')

  @staticmethod
  def _getMethods(_class, attr):
    """Returns a list of methods that has the specified attribute."""
    methods = []
    for _, elem in _class.__dict__.items():
      if hasattr(elem, attr):
        methods.append(elem)
    return methods

  def RunCommand(self, instance_name, cmd):
    """Run a command on the specified instance."""
    return self.clients[instance_name].RunCommand(cmd)

  def _runCommand(self, cmd):
    """Run a command."""
    try:
      logging.info("Running: %s", cmd)
      output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
      logging.info("Output: %s", output)
    except subprocess.CalledProcessError, e:
      logging.info("Command run failed with error code %s: %s" % (e.returncode,
                                                                  e.output))
      raise

  def UploadFile(self, instance_name, src_file, dest_directory):
    """Upload local file to the specified instance.

    Returns:
    the full path of the destination file.
    """
    return self.clients[instance_name].UploadFile(src_file, dest_directory)
