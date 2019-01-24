# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import logging
import os
import tempfile
import shutil
import subprocess

TOPIC_NAME = 'cron-stop-instances'

FUNCTION_SOURCE = """
import base64
import datetime
from dateutil.parser import parse
import googleapiclient.discovery
import os

def StopOldInstances(data, context):
  if 'data' not in data or not data['data']:
    print('StopAllInstances MUST receive the Compute zone in `data`')
    return

  project = os.environ['GCP_PROJECT']
  zone = base64.b64decode(data['data']).decode()
  print('Executing StopOldInstances: project=%s ; zone=%s' % (project, zone))

  service = googleapiclient.discovery.build('compute', 'v1')

  request = service.instances().list(project=project, zone=zone)
  response = request.execute()

  now = datetime.datetime.now(datetime.timezone.utc)
  if 'items' in response:
    for instance in response['items']:
      name = instance['name']
      age = now - parse(instance['creationTimestamp'])
      if age > datetime.timedelta(days=1):
        print("Stopping %s because it's old enough: %s" % (name, age))
        r = service.instances().stop(project=project, zone=zone, instance=name)
        r.execute()
      else:
        print("Skipping %s because it's not old enough: %s" % (name, age))
"""


def ParseArgs():
  parser = argparse.ArgumentParser(
      description=
      'Configures a GCP project to stop old compute instances automatically.')

  parser.add_argument(
      '--project',
      metavar='<project>',
      dest="project",
      required=True,
      help='The id of the project which to configure.')

  return parser.parse_args()


def ConfigureLogging(args):
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  env = dict(os.environ, CLOUDSDK_CORE_PROJECT=args.project)

  # Enable required APIs: Scheduler, Pub/Sub & Functions
  services = ['appengine', 'cloudfunctions', 'cloudscheduler', 'pubsub']
  for service in services:
    p = subprocess.Popen(
        ['gcloud', 'services', 'enable', service + ".googleapis.com"], env=env)
    output, err = p.communicate()
    logging.info("Enable %s: %s -- %s" % (service, output, err))

  # Create a Pub/Sub topic to listen to
  p = subprocess.Popen(['gcloud', 'pubsub', 'topics', 'create', TOPIC_NAME],
                       env=env)
  output, err = p.communicate()
  logging.info("Topics Create: %s -- %s" % (output, err))

  # Create the StopOldInstance function to execute
  tempDir = tempfile.mkdtemp()
  with open(os.path.join(tempDir, 'main.py'), 'w') as f:
    f.write(FUNCTION_SOURCE)

  with open(os.path.join(tempDir, 'requirements.txt'), 'w') as f:
    f.write("python-dateutil==2.7.5")

  p = subprocess.Popen([
      'gcloud', 'functions', 'deploy', 'StopOldInstances', '--trigger-topic',
      'cron-stop-instances', '--runtime', 'python37', '--source', tempDir
  ],
                       env=env)
  output, err = p.communicate()
  logging.info("Functions Deploy: %s -- %s" % (output, err))

  shutil.rmtree(tempDir)

  # Create a job to execute this function once a day
  p = subprocess.Popen(
      ['gcloud', 'beta', 'app', 'create', '--region', 'us-east1'], env=env)
  output, err = p.communicate()
  logging.info("Create App for the Scheduler job: %s -- %s" % (output, err))

  # Cron to execute daily at midnight
  daily = '0 0 * * *'
  p = subprocess.Popen([
      'gcloud', 'beta', 'scheduler', 'jobs', 'create', 'pubsub',
      'StopAllInstances', '--message-body', 'us-east1-b', '--topic',
      'projects/%s/topics/%s' % (args.project, TOPIC_NAME), '--schedule', daily
  ],
                       env=env)
  output, err = p.communicate()
  logging.info("Scheduler Jobs Create: %s -- %s" % (output, err))
