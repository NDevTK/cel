# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from google.protobuf import text_format
import logging
import os
import pydoc
import subprocess
import test.infra.gcp as gcp
from test.infra.core import EnterpriseTestCase, TestEnvironment
from test.infra.proto.schema.host import host_environment_pb2
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

    host = self._ParseHostFile(hostFile)
    self._project = gcp.ComputeProject(host.project.name, host.project.zone)

    self._celCtlRunner = CelCtlRunner(cel_ctl, self._hostFile, self._assetFile)

  def DeployNewEnvironment(self):
    """Deploys the test environment. Returns only when it is ready."""
    self._celCtlRunner.Deploy(self._deployTimeout)

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
      except KeyboardInterrupt:
        raise
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

  def TryWriteComputeLogsTo(self, destination):
    try:
      self._WriteComputeLogsTo(destination)
    except:
      print(traceback.format_exc())
      logging.error('TryWriteComputeLogsTo(%s) failed.' % destination)

  def TryCleanHostEnvironment(self):
    try:
      self._celCtlRunner.Clean()
    except:
      print(traceback.format_exc())
      logging.error('TryCleanHostEnvironment failed.')

  def _WriteComputeLogsTo(self, destination):
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

  def _ParseHostFile(self, hostFile):
    content = ''
    with open(hostFile, 'r') as f:
      content = f.read()

    host = host_environment_pb2.HostEnvironment()
    text_format.Parse(content, host)

    return host


class CelCtlRunner:

  def __init__(self, cel_ctl, hostFile, assetFile):
    self._cel_ctl = cel_ctl
    self._hostFile = hostFile
    self._assetFile = assetFile

  def Deploy(self, deployTimeout):
    cmd = [
        self._cel_ctl, 'deploy', '--builtins', self._hostFile, self._assetFile
    ]

    if deployTimeout != None:
      cmd += ['--timeout', str(deployTimeout)]

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

  def Clean(self):
    cmd = [
        self._cel_ctl, 'purge', '--builtins', self._hostFile, self._assetFile
    ]

    logging.info("Running %s" % cmd)
    code = subprocess.call(cmd)
    logging.info("cel_ctl returned code=%s" % code)

    if code != 0:
      raise CelCtlError("Clean failed.")


class CelCtlError(Exception):
  pass
