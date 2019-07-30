# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import googleapiclient
import googleapiclient.discovery
import logging
import os
import subprocess
import sys
import StringIO

# Dictionary of API names to Display names.
REQUIRED_GOOGLE_APIS = {
    'translateelement': 'Chrome Translate Element',
    'safebrowsing': 'Safe Browsing API'
}


def ParseArgs():
  parser = argparse.ArgumentParser(
      description='Configures Creds for Google APIs in Chromium E2E tests.')

  parser.add_argument(
      '--project',
      metavar='<project>',
      dest="project",
      required=True,
      help='The project id to configure.')

  parser.add_argument(
      '--storage',
      metavar='<storage>',
      dest="storage",
      required=True,
      help='The id of the storage to store the key in.')

  return parser.parse_args()


def ConfigureLogging(args):
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


def CreateCredentials(project):
  # This is not supported yet.
  # b/116182848 tracks the feature request for this.
  # Once there is an API to create Credentials, this method should:
  #   1) Create the Credentials
  #   2) Restrict it to the set of APIs in REQUIRED_GOOGLE_APIS
  #   3) Read and return the API Key value

  consoleUrl = 'https://cloud.google.com/console/apis'
  keyName = 'CELab Chromium E2E Tests API Key'
  apis = REQUIRED_GOOGLE_APIS.values()
  print('\nThe process to create the API key is manual:')
  print('  1. Visit %s/credentials/key?project=%s.' % (consoleUrl, project))
  print('  2. Name the key "%s"' % keyName)
  print('  3. Click on API restrictions and add: %s' % apis)
  print('  4. Click Create.\n')

  return raw_input('Paste the API Key value here: ')


def StoreApiKey(apiKey, storage):
  service = googleapiclient.discovery.build('storage', 'v1')

  objectName = 'api/key'
  content = StringIO.StringIO()
  content.write(apiKey)
  request = service.objects().insert(
      bucket=storage,
      body={
          'name': objectName,
          'eventBasedHold': True
      },
      media_body=googleapiclient.http.MediaIoBaseUpload(
          content, 'application/octet-stream'))

  response = request.execute()

  logging.info("objects.insert response: %s" % response)


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  # Enable the required service APIs.
  for service in REQUIRED_GOOGLE_APIS.keys():
    p = subprocess.Popen(
        ['gcloud', 'services', 'enable', service + ".googleapis.com"],
        env=dict(os.environ, CLOUDSDK_CORE_PROJECT=args.project))
    output, err = p.communicate()
    logging.info("Enable %s: %s -- %s" % (service, output, err))

  # Create Credentials (manual process) and store the key in our bucket.
  apiKey = CreateCredentials(args.project)
  StoreApiKey(apiKey, args.storage)

  sys.exit(0)
