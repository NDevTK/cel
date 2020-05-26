# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import os
import random
import string
import subprocess
import threading
import time
from datetime import datetime

from chrome_ent_test.infra.core import *


@category("core")
@environment(file="./assets/ui_test.asset.textpb")
class UITest(EnterpriseTestCase):
  """Tests that running UI tests works on all versions of Windows.
  """

  def _enableUITest(self, instance_name):
    """Configures the instance so that UI tests can be run on it."""
    self.RunCommand(instance_name, r'md -Force c:\temp')

    self.InstallChocolateyPackage(instance_name, 'chocolatey_core_extension',
                                  '1.3.3')
    self.InstallChocolateyPackage(instance_name, 'sysinternals',
                                  '2019.6.12.20190614')
    self.InstallPipPackagesLatest(instance_name,
                                  ['absl-py', 'requests', 'pywinauto'])

    password = self._generatePassword()
    user_name = 'ui_user'
    cmd = (r'powershell -File c:\cel\supporting_files\enable_auto_logon.ps1 '
           r'-userName %s -password %s') % (user_name, password)
    self.RunCommand(instance_name, cmd)
    self._rebootInstance(instance_name)

    cmd = (r'powershell -File c:\cel\supporting_files\set_ui_agent.ps1 '
           '-username %s') % user_name
    self.RunCommand(instance_name, cmd)
    self._rebootInstance(instance_name)

  def _rebootInstance(self, instance_name):
    self.RunCommand(instance_name, 'shutdown /r /t 0')

    # wait a little so that we can be sure the instance is
    # rebooting
    time.sleep(30)

    # wait for the instance to boot up and ready
    start_time = datetime.now()
    while True:
      try:
        self.RunCommand(instance_name, 'whoami')
        break
      except:
        logging.info('Instance is not ready. Retry')
        now = datetime.now()
        time_used = (now - start_time).total_seconds()
        if time_used > 5 * 60:
          self.fail("Time out when waiting for instance to boot up")
        else:
          time.sleep(60)

  def _generatePassword(self):
    """Generates a random password."""
    s = [random.choice(string.ascii_lowercase) for _ in range(4)]
    s += [random.choice(string.ascii_uppercase) for _ in range(4)]
    s += [random.choice(string.digits) for _ in range(4)]
    random.shuffle(s)
    return ''.join(s)

  def _runUITest(self, instance_name, test_file, timeout=300, args=[]):
    """Runs a UI test on an instance.

    Args:
      instance_name: name of the instance.
      test_file: the path of the UI test file.
      timeout: the timeout in seconds. Default is 300,
               i.e. 5 minutes.
      args: the list of arguments passed to the test.

    Returns:
      the output."""
    self.EnsurePythonInstalled(instance_name)

    # upload the test
    file_name = self.UploadFile(instance_name, test_file, r'c:\temp')

    # run the test.
    # note that '-u' flag is passed to enable unbuffered stdout and stderr.
    # Without this flag, if the test is killed because of timeout, we will not
    # get any output from stdout because the output is buffered. When this
    # happens it makes debugging really hard.
    args = subprocess.list2cmdline(args)
    ui_test_cmd = r'{0} -u {1} {2}'.format(
        self._pythonExecutablePath[instance_name], file_name, args)
    cmd = (r'{0} '
           r'c:\cel\supporting_files\run_ui_test.py '
           r'--timeout {1} -- {2}').format(
               self._pythonExecutablePath[instance_name], timeout, ui_test_cmd)
    return self.RunCommand(instance_name, cmd)

  def _runUITestOnInstance(self, instance_name, error):
    """Runs a UI test on the specified instance.

    Args:
      instance_name: the instance where the UI test is run
      error: a list. If the UI test cannot run successfully, an
         error message will be appended to the list
    """
    try:
      self._enableUITest(instance_name)
      dir = os.path.dirname(os.path.abspath(__file__))
      test_file = os.path.join(dir, 'ui_test_on_instance.py')
      output = self._runUITest(instance_name, test_file)
      self.assertIn('SUCCESS', output)
    except:
      error += ["UI test failed on %s" % instance_name]

  @test
  def runUITests(self):
    # run tests on instances in parallel. It takes about 10 minutes on
    # one instance, so, it'll take 50 minutes to run them in parallel.
    threads = []
    error = []
    for client in [
        'win2008', 'win2012', 'win2016', 'win2019', 'win-7', 'win-10'
    ]:
      threads += [
          threading.Thread(
              target=self._runUITestOnInstance, args=(client, error))
      ]

    for thread in threads:
      thread.start()

    for thread in threads:
      thread.join()

    # assert that there are no errors
    self.assertFalse(error)
