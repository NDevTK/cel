# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import os
import subprocess


class EnterpriseTestCase:
  """Base class for tests that provides test hooks and resources."""

  def __init__(self, environment):
    logging.info('Initialize Test=%s with %s' % (self.__class__, environment))
    self.clients = environment.clients
    self.gsbucket = environment.gsbucket

  @staticmethod
  def GetTestMethods(_class):
    testMethods = []
    for _, elem in _class.__dict__.items():
      if hasattr(elem, 'IS_TEST_METHOD'):
        testMethods.append(elem)
    return testMethods

  def assertTrue(self, assertion, message='Assertion failed'):
    if not assertion:
      raise Exception(message)

  def assertEqual(self, first, second, message='Assertion failed'):
    if first != second:
      raise Exception("%s [first=%s, second=%s]" % (message, first, second))

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
    file_name = os.path.basename(src_file)

    # On Windows, the gsutil program is named gsutil.cmd
    gsutil_executable = 'gsutil'
    if os.name == 'nt':
      gsutil_executable = 'gsutil.cmd'

    self._runCommand(
        [gsutil_executable, 'cp', src_file, 'gs://' + self.gsbucket])

    dest_file = os.path.join(dest_directory, file_name).replace('/', '\\')
    cmd = 'gsutil cp gs://%s/%s %s' % (self.gsbucket, file_name, dest_file)
    self.RunCommand(instance_name, cmd)
    return dest_file
