# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import json
import logging
import os
import re
import subprocess
import unittest

CIPD_CHOCOLATEY_ROOT_PATH = 'infra/celab/third_party/chocolatey_packages/'
CIPD_PIP_ROOT_PATH = 'infra/celab/third_party/pip-packages/'
WORK_DIRECTORY = r'c:\temp'
CIPD_PATH = r'c:\depot_tools\cipd.bat'
CHOCOLATEY_EXECUTABLE_PATH = r'c:\programdata\chocolatey\bin\choco'


# We inherit from unittest.TestCase so that we get all those assertXXX()
# methods. Note that the way our tests work is different from unittest,
# so please do not use anything in unittest.TestCase except those assert
# methods.
class EnterpriseTestCase(unittest.TestCase):

  def __init__(self, environment):
    logging.info('Initialize Test=%s with %s' % (self.__class__, environment))
    super(EnterpriseTestCase, self).__init__()
    self.clients = environment.clients
    self.gsbucket = environment.gsbucket
    self._chocolateyInstalled = {}
    self._pythonExecutablePath = {}

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
      logging.info("Command run failed with error code %s: %s" %
                   (e.returncode, e.output))
      raise

  def UploadFile(self, instance_name, src_file, dest_directory):
    """Upload local file to the specified instance.

    Returns:
    the full path of the destination file.
    """
    return self.clients[instance_name].UploadFile(src_file, dest_directory)

  def _downloadCipdPackageToClient(self, instance_name, path, ref, root):
    """Downloads a package from Chrome Infra Package Deployer (CIPD)

    Executes a CIPD ensure command to download a package from CIPD to the test
    client machine. Please refer to CIPD documentation for background on how
    to use CIPD:
    https://g3doc.corp.google.com/company/teams/chrome/ops/luci/cipd.md?\
    cl=client-command-line-interface

    Args:
      instance_name: name of the target test client instance.
      path: the full package path in CIPD.
      ref: a git-like ref pointing to a CIPD package instance.
           The ref should be either an version string, like '2.7.15',
           or a reference like 'latest'.
      root: the destination directory on the client machine.
    """
    cmd = (
        # Construct a CIPD ensure file, and set the file encoding to ASCII.
        # (CIPD cannot read from unicode files, which is Windows
        # PowerShell's default text encoding scheme)
        'echo \'{0} {1}\' | Out-File -Encoding Ascii {2}; '
        # cipd has two executable scripts, cipd.bat and cipd.ps1. Only cipd.bat
        # can accept the ensure command. Explicitly invoke cipd.bat to avoid
        # confusion.
        '{3} ensure -ensure-file {2} -root {4}').format(
            path, ref, r'{0}\ensure_file.txt'.format(root), CIPD_PATH, root)
    self.clients[instance_name].RunCommand(cmd)

  def _checkIfCipdPackageExists(self, instance_name, path, ref):
    """Checks if a CIPD package with a particular version exists.

    Executes a CIPD describe command to check if a CIPD package with a
    particular version (e.g. '3.14.0') or reference (e.g. 'latest')
    exists.

    Args:
      instance_name: name of the target test client instance.
      path: the full package path in CIPD.
      ref: a git-like ref pointing to a CIPD package instance.
           The ref should be either an version string, like '3.14.0',
           or a reference like 'latest'.

    Returns:
      true if a package with the ref exists in CIPD, false, otherwise.
    """
    cmd = '{0} resolve {1} -version {2}'.format(CIPD_PATH, path, ref)
    try:
      output = self.clients[instance_name].RunCommand(cmd)
      return "Packages:" in output
    except subprocess.CalledProcessError as e:
      # If the path does not exist in CIPD at all, cipd will return an error.
      # Detect this case and return false.
      if "no such package." in e.output:
        return False
      raise

  def _addPathToClientEnvironment(self, instance_name, path):
    """Add a path to the test client machine's %PATH% environment variable.

    Note that this method adds the path to the test client's machine env. So
    that the path is present for all subsequent cmds. Very useful for ensuring
    that a newly installed piece of software is executable without exposing
    where the software binaries reside.

    Args:
      instance_name: name of the target test client instance.
      path: the path to add to the %PATH% env variable.
    """
    cmd = ('$env:Path += ";{0}"; '
           '[System.Environment]::SetEnvironmentVariable('
           '    \'Path\', $env:Path,'
           '    [System.EnvironmentVariableTarget]::Machine)').format(path)
    self.clients[instance_name].RunPowershell(cmd)

  def EnsureChocolateyInstalled(self, instance_name):
    if instance_name in self._chocolateyInstalled:
      return

    cipdPath = '{0}chocolatey'.format(CIPD_CHOCOLATEY_ROOT_PATH)
    if self._checkIfCipdPackageExists(instance_name, cipdPath, 'latest'):
      # Install Chocolatey using the package cached in CIPD.
      self._downloadCipdPackageToClient(instance_name, cipdPath, 'latest',
                                        WORK_DIRECTORY)
      self.clients[instance_name].RunPowershell(r'{0}\{1}'.format(
          WORK_DIRECTORY, 'install.chocolatey.ps1'))
    else:
      # Install chocolatey from the internet.
      cmd = (
          'Set-ExecutionPolicy Bypass -Scope Process -Force; '
          # By default Powershell uses TLS (Transport Layer Security) 1.0,
          # but the Chocolatey site requires TLS 1.2.
          '[Net.ServicePointManager]::SecurityProtocol ='
          ' [Net.SecurityProtocolType]::Tls12;'
          'iex ((New-Object System.Net.WebClient)'
          '.DownloadString(\'https://chocolatey.org/install.ps1\'))')
      self.clients[instance_name].RunPowershell(cmd)

    # The Chocolatey installation script does not properly add the
    # chocolatey binary directory to the test client machine's path.
    # So this method add the path here.
    self._addPathToClientEnvironment(instance_name,
                                     r'C:\ProgramData\chocolatey\bin')

    self._chocolateyInstalled[instance_name] = True

  def EnsurePythonInstalled(self, instance_name):
    if instance_name in self._pythonExecutablePath:
      return

    installedFromCipd = self.InstallChocolateyPackage(instance_name, 'python2',
                                                      '2.7.15')

    # Depending on whether Chocolatey installed Python2 from the Internet or
    # from a local nupkg, the Python binary path will be different.
    self._pythonExecutablePath[instance_name] =\
       r'C:\ProgramData\chocolatey\lib\python2\tools\python.exe'\
       if installedFromCipd else r'c:\Python27\python.exe'

    # Install wheel 0.34.2. The wheel that comes with the python2 2.7.15 inbox
    # installation cannot handle the pip package comtypes. (A dependency of
    # the pip package pywinauto, which ui test needs).
    self.InstallPipPackage(instance_name, 'wheel', '0.34.2')

  def InstallChocolateyPackage(self, instance_name, name, version):
    self.EnsureChocolateyInstalled(instance_name)
    installedFromCipd = False

    cipdPath = '{0}{1}'.format(CIPD_CHOCOLATEY_ROOT_PATH, name)
    if self._checkIfCipdPackageExists(instance_name, cipdPath, version):
      # Download cipd package into a special package directory.
      # CIPD does not list the content of its packages. So this function must
      # run the 'Get-ChildItem' cmdlet in a clean directory to figure out the
      # name of the nupkg file it downloaded from cipd.
      # Construct the directory name based on the package name and package
      # version. Powershell's Remove-Item cmdlet is unstable and prone to
      # failure, so we cannot rely on cleaning and reusing an existing dest
      # directory.
      dirName = re.sub('[^\w\-_\. ]', '_', '{0}_{1}'.format(name, version))
      nupkgDir = r'{0}\{1}'.format(WORK_DIRECTORY, dirName)
      # Create the package directory if it does not exist.
      cmd = 'New-Item -ItemType Directory -Force -Path {0}'.format(nupkgDir)
      self.clients[instance_name].RunPowershell(cmd)

      # Install package using the package cached in CIPD.
      self._downloadCipdPackageToClient(instance_name, cipdPath, version,
                                        nupkgDir)
      # Find the name of the nupkg file.
      cmd = (
          r'$package = Get-ChildItem -Path {0} | '
          # Find '.nupkg' files or 'nuspec' files.
          'Where-Object -FilterScript{{$_.Name -match \'.nupkg$|.nuspec$\'}};'
          # Use chocolatey to install the nupkg file.
          # By default Get-ChildItem returns the file name. '.FullName' returns
          # the file's full path.
          r'{1} install -y $package.FullName').format(
              nupkgDir, CHOCOLATEY_EXECUTABLE_PATH)
      self.clients[instance_name].RunPowershell(cmd)
      installedFromCipd = True
    else:
      # Install package from the internet.
      cmd += '{0} install {1} -y'.format(CHOCOLATEY_EXECUTABLE_PATH, name)
      if version is not None and version != 'latest':
        cmd += ' --version {0}'.format(version)
      self.RunCommand(instance_name, cmd)

    return installedFromCipd

  def InstallChocolateyPackageLatest(self, instance_name, name):
    return self.InstallChocolateyPackage(instance_name, name, 'latest')

  def InstallPipPackagesLatest(self, instance_name, packages):
    self.EnsurePythonInstalled(instance_name)
    for package in packages:
      self.InstallPipPackage(instance_name, package, 'latest')

  def InstallPipPackage(self, instance_name, package, version):
    cipdPath = '{0}{1}'.format(CIPD_PIP_ROOT_PATH, package)
    if self._checkIfCipdPackageExists(instance_name, cipdPath, version):
      # Install package using the package cached in CIPD.
      self._downloadCipdPackageToClient(instance_name, cipdPath, version,
                                        WORK_DIRECTORY)
      self.RunCommand(instance_name,
                      (r'{0} -m pip install '
                       r'-r {1}\{2}.txt --no-index --find-links {1}').format(
                           self._pythonExecutablePath[instance_name],
                           WORK_DIRECTORY, package))
    else:
      # Install package from the internet.
      # TODO(uwyiming) force the version here if the version string is valid.
      self.RunCommand(
          instance_name, r'{0} -m pip install {1}'.format(
              self._pythonExecutablePath[instance_name], package))
