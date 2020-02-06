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
    self._chocolateyInstalled = {}
    self._pythonInstalled = {}

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

  def RunCommand(self, instance_name, cmd, timeout=None):
    """Run a command on the specified instance."""
    return self.clients[instance_name].RunCommand(cmd, timeout=timeout)

  def _runCommand(self, cmd):
    """Run a command."""
    try:
      logging.info("Running: %s", cmd)
      output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
      logging.info("Output: %s", output)
    except subprocess.CalledProcessError as e:
      logging.info("Command run failed with error code %s: %s" % (e.returncode,
                                                                  e.output))
      raise

  def UploadFile(self, instance_name, src_file, dest_directory):
    """Upload local file to the specified instance.

    Returns:
    the full path of the destination file.
    """
    return self.clients[instance_name].UploadFile(src_file, dest_directory)

  def EnsureChocolateyInstalled(self, instance_name):
    if instance_name in self._chocolateyInstalled:
      return

    cmd = "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"
    self.clients[instance_name].RunPowershell(cmd)

    self._chocolateyInstalled[instance_name] = True

  def EnsurePythonInstalled(self, instance_name):
    if instance_name in self._pythonInstalled:
      return

    self.InstallChocolateyPackage(instance_name, 'python2', '2.7.15')

    self._pythonInstalled[instance_name] = True

  def InstallChocolateyPackage(self, instance_name, name, version):
    self.EnsureChocolateyInstalled(instance_name)
    cmd = r'c:\ProgramData\chocolatey\bin\choco install %s -y --version %s' % (
        name, version)
    self.RunCommand(instance_name, cmd)

  def InstallChocolateyPackageLatest(self, instance_name, name):
    self.EnsureChocolateyInstalled(instance_name)
    cmd = r'c:\ProgramData\chocolatey\bin\choco install %s -y' % name
    self.RunCommand(instance_name, cmd)

  def InstallPipPackagesLatest(self, instance_name, packages):
    self.EnsurePythonInstalled(instance_name)
    self.RunCommand(
        instance_name,
        r'c:\Python27\python.exe -m pip install %s' % ' '.join(packages))
