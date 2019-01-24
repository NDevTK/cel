# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import googleapiclient
import googleapiclient.discovery
import logging
import sys
import time


def ParseArgs():
  parser = argparse.ArgumentParser(
      description='Creates a GCP project linked to a given billing account.')

  parser.add_argument(
      '--project',
      metavar='<project>',
      dest="project",
      required=True,
      help='The id/name of the project to create.')

  parser.add_argument(
      '--folder',
      metavar='<folder>',
      dest="folder",
      required=True,
      help='The id of the folder in which to create the project.')

  parser.add_argument(
      '--billing',
      metavar='<billing>',
      dest="billing",
      required=True,
      help='The id of the billing account to link the project to.')

  return parser.parse_args()


def ConfigureLogging(args):
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


def CreateProject(args):
  service = googleapiclient.discovery.build('cloudresourcemanager', 'v1')

  request = service.projects().create(
      body={
          'project_id': args.project,
          'name': args.project,
          'parent': {
              'type': 'folder',
              'id': args.folder
          }
      })

  response = request.execute()
  logging.info("projects.create response: %s" % response)

  # Wait for the project creation to complete
  operation = response['name']
  while True:
    request = service.operations().get(name=operation)
    r = request.execute()
    logging.info("operations.get response: %s" % r)
    if 'done' in r and r['done']:
      break
    time.sleep(5)


def UpdateBilling(args):
  service = googleapiclient.discovery.build('cloudbilling', 'v1')

  fullProjectName = 'projects/' + args.project
  fullBillingAccount = 'billingAccounts/' + args.billing
  request = service.projects().updateBillingInfo(
      name=fullProjectName,
      body={
          "projectId": args.project,
          "billingAccountName": fullBillingAccount,
          "billingEnabled": True
      })

  response = request.execute()
  logging.info("projects.updateBillingInfo response: %s" % response)


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  try:
    CreateProject(args)
  except googleapiclient.errors.HttpError as e:
    if e.resp.status != 409:
      raise
    logging.warning("Project already exists. Moving to billing configuration.")

  UpdateBilling(args)

  sys.exit(0)
