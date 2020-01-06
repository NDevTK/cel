# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import base64
import logging
import subprocess


class TestClient:

  def __init__(self, computeInstance, celCtlRunner):
    self._computeInstance = computeInstance
    self._celCtlRunner = celCtlRunner
    self.name = computeInstance.name

  def RunPowershell(self, script):
    # Ignore the Powershell progress stream - it prints garbage in our output.
    script = "$ProgressPreference = 'SilentlyContinue'\n" + script.strip()

    # PowerShell EncodedCommand takes a UTF-16 little endian b64encoded string
    logging.info("RunPowershell script: %s" % repr(script))
    b64Script = base64.b64encode(script.encode('utf-16-le'))
    return self.RunCommand('powershell', ['-EncodedCommand', b64Script])

  def RunPowershellNoThrow(self, script):
    """Run powershell script, without raising exception on failure.

    Returns: (retcode, output)"""
    try:
      output = self.RunPowershell(script)
      return 0, output
    except subprocess.CalledProcessError as e:
      return e.returncode, e.output

  def RunCommand(self, cmd, args=[]):
    # TODO: Escape arguments
    commandString = '%s %s' % (cmd, ' '.join(args))
    return self._celCtlRunner.RunCommand(self.name, commandString)

  def UploadFile(self, file, destination):
    return self._celCtlRunner.UploadFile(self.name, file, destination)

  def __repr__(self):
    return "<%s %s>" % (self.__class__.__name__, self.name)


class TestEnvironment:

  def __init__(self, computeProject, gsbucket, celCtlRunner):
    computeInstances = computeProject.GetComputeInstances()

    self.clients = {}
    for instance in computeInstances:
      testClient = TestClient(instance, celCtlRunner)
      self.clients[instance.name] = testClient

    self.gsbucket = gsbucket

  def __repr__(self):
    return "<%s clients=%s gsbucket=%s>" % (self.__class__.__name__,
                                            self.clients, self.gsbucket)
