# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import inspect
import logging
import os
import pydoc
import subprocess
import test.infra.gcp as gcp
import traceback
from test.infra.core import EnterpriseTestCase, TestEnvironment
from test.infra.proto.schema.host import host_environment_pb2

from google.protobuf import text_format


class SingleTestController:

  def __init__(self, testCaseClassName, hostFile, cel_ctl):
    hostFile = os.path.expanduser(hostFile)
    if not os.path.exists(hostFile):
      raise ValueError('Host file not found: %s' % hostFile)

    testClass = pydoc.locate(testCaseClassName)
    if testClass == None:
      message = 'Class not found: %s.' % testCaseClassName
      raise ValueError(message)

    if not hasattr(testClass, "ASSET_FILE"):
      message = 'Class found with no @environment: %s.' % testCaseClassName
      raise ValueError(message)

    asset_file = os.path.join(
        os.path.dirname(inspect.getfile(testClass)), testClass.ASSET_FILE)
    if not os.path.exists(asset_file):
      raise ValueError('Asset file not found: %s' % asset_file)

    self._testClass = testClass
    self._hostFile = hostFile
    self._assetFile = asset_file
    self._deployTimeout = testClass.DEPLOY_TIMEOUT

    host = self._ParseHostFile(hostFile)
    self._project = gcp.ComputeProject(host.project.name, host.project.zone)
    self._host = host

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
    environment = TestEnvironment(self._project, self._host.storage.bucket,
                                  self._celCtlRunner)

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

    # Print summary
    summary = "\n%s/%s tests passed.\n" % (passes, len(tests))
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
        f.write(logs.encode("utf-8"))

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

    logging.info("Running on %s: %s" % (instance, command))
    try:
      output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
      logging.info("cel_ctl run output: %s" % output)
      return output
    except subprocess.CalledProcessError, e:
      logging.debug("cel_ctl run returned %s: %s" % (e.returncode, e.output))
      raise

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
