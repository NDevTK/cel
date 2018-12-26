# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import api_client
import datetime
import logging
import sys
import time


class CloudRuntimeConfig:

  def __init__(self, name, project):
    self._name = name
    self._project = project
    self._configPath = 'projects/%s/configs/%s' % (project.name, name)
    self._assets = {}

  def WaitForAllAssetsReady(self, showProgress, timeout=3600, interval=30):
    """Returns when assets are ready (or failed).

    We poll the runtime config variables every `interval` seconds to find out.

    # TODO: Use variables.watch for low numbers of in-progress assets is small.
    # TODO: Implement a second timeout value based on updateTime.

    Args:
      showProgress: If we should print 'in-progress' status.
      timeout: Timeout value in seconds.
      interval: How many seconds to wait between each poll.

    Raises:
      AssetError: We timed out
      TimeoutError: All assets were not ready after `timeout`
    """
    logging.info('Entering WaitForAllAssetsReady(%s)' % timeout)

    api = api_client.build('runtimeconfig', 'v1beta1')

    startTime = time.time()
    assetStates = {}
    while startTime + timeout > time.time():
      variables = self._FetchVariables(api)

      assetStates = self._GetAssetStatesFromVariables(variables)

      if 'error' in assetStates:
        raise AssetError('Deployment failed: %s' % assetStates['error'])
      elif 'in-progress' in assetStates or "init" in assetStates:
        if showProgress:
          total = 0
          if 'init' in assetStates:
            total += len(assetStates['init'])
          if 'in-progress' in assetStates:
            total += len(assetStates['in-progress'])

          now = datetime.datetime.now().strftime('%H:%M:%S')
          message = '%s assets still deploying...' % total
          sys.stdout.write("[%s] %s\r" % (now, message))
          sys.stdout.flush()
        time.sleep(interval)
      elif len(assetStates) == 1 and 'ready' in assetStates:
        message = 'All (%s) assets are ready' % len(assetStates['ready'])
        if showProgress:
          now = datetime.datetime.now().strftime('%H:%M:%S')
          print('[%s] %s: %s' % (now, message, assetStates))
        else:
          logging.info('%s: %s' % (message, assetStates))
        return
      else:
        raise Exception('Failed to parse asset states: %s' % assetStates)

    message = 'Timed out waiting for assets to be ready'
    message += ': (timeout=%s, assets=%s)' % (timeout, assetStates)
    raise TimeoutError(message)

  def _GetAssetStatesFromVariables(self, variables):
    """Parses runtime config variables into asset information.

    Returns:
      A dictionary mapping status to assets in this format:
        { status: [assetA, assetB, ...], ... }
      status=(init, in-progress, error, ready)
    """
    assetStates = {}
    for variable in variables:
      name, status = self._ParseAssetInfoFromVariable(variable)
      if name == None:
        logging.warning("Couldn't parse asset state from '%s'" % variable)
        continue

      if status not in assetStates:
        assetStates[status] = []

      assetStates[status].append(name)

    return assetStates

  def _FetchVariables(self, api, maxRetries=3):
    """Fetches the variables for this config."""
    request = api.projects().configs().variables().list(
        parent=self._configPath, returnValues=True)

    retries = 0
    while True:
      try:
        response = request.execute()

        return response['variables']
      except Exception as e:
        logging.warning('Exception while fetching config variables: %s' % e)
        retries += 1
        if retries > maxRetries:
          raise

  def _ParseAssetInfoFromVariable(self, variable):
    """Tries to parse asset information from a config variable.

    Args:
      A variable structure as defined in a Cloud Runtime Config API.

    Returns:
      (name, status) if success
      (None, None) if failed
    """
    name = variable['name']
    assetPrefix = '%s/variables/asset/' % self._configPath
    statusSuffix = '/status'
    if not (name.startswith(assetPrefix) and name.endswith(statusSuffix)):
      return None, None

    name = name[len(assetPrefix):-len(statusSuffix)]
    status = variable['text'] if 'text' in variable else 'init'

    return name, status


class AssetError(Exception):
  pass


class TimeoutError(Exception):
  pass
