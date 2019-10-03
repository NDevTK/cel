# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import googleapiclient.discovery
import logging
import os
import subprocess
import sys


def ParseArgs():
  parser = argparse.ArgumentParser(
      description='Configures a GCP project to work with our test infra.')

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
      help='The id of the storage to link this project to.')

  parser.add_argument(
      '--service_accounts',
      metavar='<acc1;acc2>',
      dest="accounts",
      required=True,
      help='The service accounts that will use this project with cel_ctl.')

  return parser.parse_args()


def ConfigureLogging(args):
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


def AddProjectBinding(project, role, member):
  # Get current bindings
  service = googleapiclient.discovery.build('cloudresourcemanager', 'v1')
  request = service.projects().getIamPolicy(resource=project, body={})
  response = request.execute()

  # Add Binding to existing binding
  response['bindings'].append({"role": role, "members": [member]})

  # Update bindings
  request = service.projects().setIamPolicy(
      resource=project, body={'policy': response})
  response = request.execute()

  logging.info("projects.setIamPolicy response: %s" % response)


def AddStorageBinding(storage, role, member):
  # Get current bindings
  service = googleapiclient.discovery.build('storage', 'v1')
  request = service.buckets().getIamPolicy(bucket=storage)
  response = request.execute()

  # Add Binding to existing binding
  response['bindings'].append({"role": role, "members": [member]})

  # Update bindings
  request = service.buckets().setIamPolicy(bucket=storage, body=response)
  response = request.execute()

  logging.info("buckets.setIamPolicy response: %s" % response)


def CreateServiceAccount(project, name, displayName):
  # Create a new service account.
  service = googleapiclient.discovery.build('iam', 'v1')

  body = {'accountId': name, 'serviceAccount': {'displayName': displayName}}

  request = service.projects().serviceAccounts().create(
      name='projects/' + project, body=body)
  response = request.execute()

  return response['email']


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  # Enable the required services to "warm up" projects before the first test.
  services = [
      'compute', 'cloudresourcemanager', 'iam', 'runtimeconfig', 'monitoring',
      'cloudkms', 'deploymentmanager'
  ]
  for service in services:
    p = subprocess.Popen(
        ['gcloud', 'services', 'enable', service + ".googleapis.com"],
        env=dict(os.environ, CLOUDSDK_CORE_PROJECT=args.project))
    output, err = p.communicate()
    logging.info("Enable %s: %s -- %s" % (service, output, err))

  # Give project permissions to service accounts
  for account in args.accounts.split(';'):
    AddProjectBinding(args.project, "roles/editor",
                      "serviceAccount:%s" % account)
    AddProjectBinding(args.project,
                      "roles/cloudkms.cryptoKeyEncrypterDecrypter",
                      "serviceAccount:%s" % account)

  # Create the account used in our VMs and setup access to the shared storage.
  email = CreateServiceAccount(args.project, "cel-instance-service",
                               "CEL Instance Service Account")
  AddProjectBinding(args.project, "roles/editor", "serviceAccount:%s" % email)

  # Give storage permissions to editors of this project.
  AddStorageBinding(args.storage, "roles/storage.objectAdmin",
                    "projectEditor:%s" % args.project)

  sys.exit(0)
