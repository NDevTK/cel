# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import googleapiclient.discovery
import logging
import sys


def ParseArgs():
  parser = argparse.ArgumentParser(
      description='Creates a storage bucket in a given project.')

  parser.add_argument(
      '--project',
      metavar='<project>',
      dest="project",
      required=True,
      help='The id of the project in which to create the storage.')

  parser.add_argument(
      '--storage',
      metavar='<storage>',
      dest="storage",
      required=True,
      help='The id of the storage to create.')

  parser.add_argument(
      '--deleteafter',
      dest="deleteafter",
      required=False,
      default=None,
      help='The number of days after which to delete objects.')

  parser.add_argument(
      '--domainreader',
      dest="domainreader",
      required=False,
      default=None,
      help='An additional IAM domain member to add read permissions to.')

  return parser.parse_args()


def ConfigureLogging(args):
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


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


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  service = googleapiclient.discovery.build('storage', 'v1')

  body = {'name': args.storage}

  if args.deleteafter:
    action = {'type': 'Delete'}
    condition = {'age': int(args.deleteafter)}
    rule = {"action": action, "condition": condition}
    body['lifecycle'] = {"rule": [rule]}

  request = service.buckets().insert(project=args.project, body=body)

  response = request.execute()
  logging.info("buckets.insert response: %s" % response)

  if args.domainreader:
    AddStorageBinding(args.storage, "roles/storage.objectViewer",
                      "domain:%s" % args.domainreader)
    AddStorageBinding(args.storage, "roles/storage.legacyObjectReader",
                      "domain:%s" % args.domainreader)

  sys.exit(0)
