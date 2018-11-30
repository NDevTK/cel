# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.


class TestClient:

  def __init__(self, computeInstance):
    self._computeInstance = computeInstance
    self.name = computeInstance.name

  def RunCommand(self, cmd, args):
    # TODO: Implement RunCommand
    pass

  def __repr__(self):
    return "<%s %s>" % (self.__class__.__name__, self.name)


class TestEnvironment:

  def __init__(self, computeProject):
    computeInstances = computeProject.GetComputeInstances()

    self.clients = {}
    for instance in computeInstances:
      testClient = TestClient(instance)
      self.clients[instance.name] = testClient

  def __repr__(self):
    return "<%s clients=%s>" % (self.__class__.__name__, self.clients)
