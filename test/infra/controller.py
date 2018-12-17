# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import os
import pydoc
import re
import subprocess
import test.infra.gcp as gcp
from test.infra.core import EnterpriseTestCase, TestEnvironment
import traceback


class SingleTestController:

  def __init__(self, testCaseClassName, hostFile, cel_ctl):
    if not os.path.exists(hostFile):
      raise ValueError('Host file not found: %s' % hostFile)

    testClass = pydoc.locate(testCaseClassName)
    if testClass == None:
      message = 'Class not found: %s.' % testCaseClassName
      raise ValueError(message)

    if not hasattr(testClass, "ASSET_FILE"):
      message = 'Class found with no @environment: %s.' % testCaseClassName
      raise ValueError(message)

    if not os.path.exists(testClass.ASSET_FILE):
      raise ValueError('Asset file not found: %s' % testClass.ASSET_FILE)

    self._testClass = testClass
    self._hostFile = hostFile
    self._assetFile = testClass.ASSET_FILE
    self._deployTimeout = testClass.DEPLOY_TIMEOUT

    name, zone = self._ParseProjectInfo(hostFile)
    self._project = gcp.ComputeProject(name, zone)

    self._celCtlRunner = CelCtlRunner(cel_ctl, self._hostFile, self._assetFile)

  def DeployNewEnvironment(self, showProgress=False):
    """Deploys the test environment. Returns only when it is ready."""
    self._celCtlRunner.Deploy()

    # Wait for the on-host deployment scripts to finish.
    print("Waiting for all assets to be ready...")
    config = gcp.CloudRuntimeConfig('cel-config', self._project)

    kwargs = {'showProgress': showProgress}
    if self._deployTimeout:
      kwargs['timeout'] = self._deployTimeout

    config.WaitForAllAssetsReady(**kwargs)

  def ExecuteTestCase(self):
    """Runs all the @test methods for this TestCase.

    Prints test results & exceptions to stdout.

    Returns:
      True if all tests passed.
    """
    environment = TestEnvironment(self._project, self._celCtlRunner)

    testCaseInstance = self._testClass(environment)

    print("Running tests...\n")

    passes = 0
    tests = EnterpriseTestCase.GetTestMethods(self._testClass)
    for test in tests:
      try:
        logging.info("Running test %s" % test)
        test(self=testCaseInstance)
        print("PASSED   %s" % test.func_name)
        passes += 1
      except:
        print("FAILED   %s" % test.func_name)
        print(traceback.format_exc())

    success = (passes == len(tests))

    # Print summary (in red if there are failures)
    summary = "\n%s/%s tests passed.\n" % (passes, len(tests))
    if not success:
      summary = '\033[91m%s\033[0m' % summary
    print(summary)

    return success

  def WriteComputeLogsTo(self, destination):
    """Writes all useful logs to investigate a test failure."""
    if not os.path.exists(destination):
      os.makedirs(destination)

    # Fetch & write Compute instance logs
    logPrefix = 'WriteLogsTo(%s)' % destination
    for instance in self._project.GetComputeInstances():
      logging.info("%s - ComputeInstance %s" % (logPrefix, instance.name))
      pathToLog = os.path.join(destination, "instance-%s.log" % instance.name)
      logs = instance.GetLatestConsoleOutput()
      with open(pathToLog, 'w') as f:
        f.write(logs)

  def _ParseProjectInfo(self, hostFile):
    # TODO: Generate *_pb.py and use google.protobuf.text_format.Parse
    content = ''
    with open(hostFile, 'r') as f:
      content = f.read()

    pattern = re.compile(r'project\s*\{([^\}]*)\}', re.DOTALL)
    match = pattern.match(content)
    if not match:
      raise Exception('Failed to parse host file.')

    projectContent = match.groups()[0]
    logging.debug("ParseProjectInfo matched: %s" % repr(projectContent))

    parts = {}
    for line in projectContent.splitlines():
      line = line.strip()
      if line.startswith('#'):
        continue
      pattern = re.compile(r'(?P<key>(name)|(zone))\s*:\s*["\'](?P<v>.*)["\']')
      match = pattern.match(line)
      if match:
        p = match.groupdict()
        parts[p['key']] = p['v']

    return parts['name'], parts['zone']


class CelCtlRunner:

  def __init__(self, cel_ctl, hostFile, assetFile):
    self._cel_ctl = cel_ctl
    self._hostFile = hostFile
    self._assetFile = assetFile

  def Deploy(self):
    cmd = [
        self._cel_ctl, 'deploy', '--builtins', self._hostFile, self._assetFile
    ]

    logging.info("Running %s" % cmd)
    code = subprocess.call(cmd)
    logging.info("cel_ctl returned code=%s" % code)

    if code != 0:
      raise CelCtlError("Deployment failed.")

  def RunCommand(self, instance, command):
    cmd = [
        self._cel_ctl, 'run', '--instance', instance, '--command', command,
        '--builtins', self._hostFile, self._assetFile
    ]

    try:
      logging.debug("Running on %s: %s" % (instance, command))
      output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
      logging.debug("cel_ctl run output: %s" % output)
      return 0, output
    except subprocess.CalledProcessError, e:
      logging.debug("cel_ctl run returned %s: %s" % (e.returncode, e.output))
      return e.returncode, e.output


class CelCtlError(Exception):
  pass
