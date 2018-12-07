# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import base64


class TestClient:

  def __init__(self, computeInstance, celCtlRunner):
    self._computeInstance = computeInstance
    self._celCtlRunner = celCtlRunner
    self.name = computeInstance.name

  def RunPowershell(self, script):
    # PowerShell EncodedCommand takes a UTF-16 little endian b64encoded string
    b64Script = base64.b64encode(script.encode('utf-16-le'))
    return self.RunCommand('powershell', ['-EncodedCommand', b64Script])

  def RunCommand(self, cmd, args):
    # TODO: Escape arguments
    commandString = '%s %s' % (cmd, ' '.join(args))
    return self._celCtlRunner.RunCommand(self.name, commandString)

  def __repr__(self):
    return "<%s %s>" % (self.__class__.__name__, self.name)


class TestEnvironment:

  def __init__(self, computeProject, celCtlRunner):
    computeInstances = computeProject.GetComputeInstances()

    self.clients = {}
    for instance in computeInstances:
      testClient = TestClient(instance, celCtlRunner)
      self.clients[instance.name] = testClient

  def __repr__(self):
    return "<%s clients=%s>" % (self.__class__.__name__, self.clients)
