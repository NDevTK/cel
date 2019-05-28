# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import base64
import logging
import os
import subprocess
from chrome_ent_test.infra.core import EnterpriseTestCase

from absl import flags

FLAGS = flags.FLAGS
flags.DEFINE_string('chrome_installer', None,
                    'The path to the chrome installer')
flags.mark_flag_as_required('chrome_installer')

flags.DEFINE_string(
    'chromedriver', None,
    'The path to the chromedriver executable. If not specified, '
    'a chocholatey chromedriver packae will be installed and used.')


class ChromeEnterpriseTestCase(EnterpriseTestCase):
  """Base class for Chrome enterprise test cases."""

  def InstallChrome(self, instance_name):
    """Installs chrome.

    Currently supports two types of installer:
    - mini_installer.exe, and
    - *.msi
    """
    file_name = self.UploadFile(instance_name, FLAGS.chrome_installer,
                                r'c:\temp')

    if os.path.basename(file_name).lower() == 'mini_installer.exe':
      dir = os.path.dirname(os.path.abspath(__file__))
      self.UploadFile(instance_name, os.path.join(dir, 'installer_data'),
                      r'c:\temp')
      cmd = file_name + r' --installerdata=c:\temp\installer_data'
    else:
      cmd = 'msiexec /i %s' % file_name

    self.RunCommand(instance_name, cmd)

  def SetPolicy(self, instance_name, policy_name, policy_value, policy_type):
    r"""Sets a Google Chrome policy in registry.

    Args:
      policy_name: the policy name.
        The name can contain \. In this case, the last segment will be the
        real policy name, while anything before will be part of the key.
    """
    segments = policy_name.split('\\')
    policy_name = segments[-1]

    # The policy will be set for both Chrome and Chromium, since only
    # googlers can build Chrome-branded executable.
    keys = [
        r'HKLM\Software\Policies\Google\Chrome',
        r'HKLM\Software\Policies\Chromium'
    ]
    for key in keys:
      if len(segments) >= 2:
        key += '\\' + '\\'.join(segments[:-1])

      cmd = (r"Set-GPRegistryValue -Name 'Default Domain Policy' "
             "-Key %s -ValueName %s -Value %s -Type %s") % (
                 key, policy_name, policy_value, policy_type)
      self.clients[instance_name].RunPowershell(cmd)

  def _installChocolatey(self, instance_name):
    cmd = "Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"
    self.clients[instance_name].RunPowershell(cmd)

  def InstallPackage(self, instance_name, package_name, package_version):
    cmd = r'c:\ProgramData\chocolatey\bin\choco install %s -y --version %s' % (
        package_name, package_version)
    self.RunCommand(instance_name, cmd)

  def InstallWebDriver(self, instance_name):
    self._installChocolatey(instance_name)
    self.InstallPackage(instance_name, 'python2', '2.7.15')
    self.RunCommand(
        instance_name,
        r'c:\Python27\python.exe -m pip install -U selenium absl-py pywin32')
    if FLAGS.chromedriver is None:
      # chromedriver flag is not specified. In this case, install the chocolatey
      # package
      self.InstallPackage(instance_name, 'chromedriver', '74.0.3729.60')
    else:
      self.UploadFile(instance_name, FLAGS.chromedriver,
                      'C:/ProgramData/chocolatey/lib/chromedriver/tools/')

    dir = os.path.dirname(os.path.abspath(__file__))
    self.UploadFile(instance_name, os.path.join(dir, 'test_util.py'),
                    r'c:\temp')

  def RunWebDriverTest(self, instance_name, test_file, args=[]):
    """Runs a python webdriver test on an instance.

    Args:
      instance_name: name of the instance.
      test_file: the path of the webdriver test file.
      args: the list of arguments passed to the test.

    Returns:
      the output."""
    # upload the test
    file_name = self.UploadFile(instance_name, test_file, r'c:\temp')

    # run the test
    cmd = r'c:\Python27\python.exe %s %s' % (file_name, ' '.join(args))
    return self.RunCommand(instance_name, cmd)