# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import googleapiclient.discovery


class ComputeProject:

  def __init__(self, name, zone):
    self.name = name
    self.zone = zone

  def GetComputeInstances(self):
    service = googleapiclient.discovery.build('compute', 'v1')

    request = service.instances().list(project=self.name, zone=self.zone)

    response = request.execute()

    instances = []
    for instance in response['items']:
      instances.append(ComputeInstance(instance['name'], self))

    return instances


class ComputeInstance:

  def __init__(self, name, project):
    self.name = name
    self._project = project
    self._consoleOutput = ''
    self._next = 0

  def GetLatestConsoleOutput(self):
    service = googleapiclient.discovery.build('compute', 'v1')

    request = service.instances().getSerialPortOutput(
        instance=self.name,
        project=self._project.name,
        zone=self._project.zone,
        start=self._next)

    response = request.execute()

    self._consoleOutput += response['contents']
    self._next = response['next']

    return self._consoleOutput

  def RunCommand(self, cmd, args):
    # TODO: Implement RunCommand
    pass
