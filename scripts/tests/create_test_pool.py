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
      description='Configures a project pool for our test infra.')

  parser.add_argument(
      '--name',
      metavar='<name>',
      dest="name",
      required=True,
      help='The name of the project pool. ' +
      'GCP projects will be named: <name>-001, <name>-002, ...')

  parser.add_argument(
      '--folder',
      metavar='<folder>',
      dest="folder",
      required=True,
      help='The id of the folder in which to create the projects.')

  parser.add_argument(
      '--billing',
      metavar='<billing>',
      dest="billing",
      required=True,
      help='The id of the billing account to link the projects to.')

  parser.add_argument(
      '--size',
      metavar='<size>',
      dest="size",
      required=True,
      help='The size of the test pool (number of projects).')

  parser.add_argument(
      '--service_accounts',
      metavar='<acc1;acc2>',
      dest="accounts",
      required=True,
      help='The service accounts that will use this pool.')

  return parser.parse_args()


def ConfigureLogging(args):
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


def RunPython(script, args):
  cmd = ["python", script] + args

  logging.info("Running: %s" % cmd)
  output = subprocess.check_output(cmd)
  logging.info("Output: %s" % output)


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  pool_size = int(args.size)
  if pool_size <= 0:
    raise ValueError('Pool size must be bigger than 0: %s' % args.size)

  # There's no obvious reason why a bigger pool wouldn't work, but we haven't
  # verified it. Let's stick to sizes that what we know will work and prevent
  # accidentally creating huge numbers or GCP projects.
  if pool_size > 10:
    raise ValueError('Pool size must be smaller than 10: %s' % args.size)

  current_dir = os.path.dirname(os.path.realpath(__file__))
  create_project_py = os.path.join(current_dir, 'create_gcp_project.py')
  create_storage_py = os.path.join(current_dir, 'create_gcp_storage.py')
  config_project_py = os.path.join(current_dir, 'configure_test_project.py')
  config_cronjob_py = os.path.join(current_dir,
                                   'configure_cron_stop_instances.py')
  storage_assets = "%s-assets" % args.name
  storage_logs = "%s-logs" % args.name

  # Create all the GCP projects
  for i in xrange(0, pool_size):
    project_id = "%s-%03d" % (args.name, i + 1)

    cmd_args = [
        '--project', project_id, '--folder', args.folder, '--billing',
        args.billing
    ]
    RunPython(create_project_py, cmd_args)

    # Create the shared storages in the first project
    if i == 0:
      cmd_args = [
          '--project', project_id, '--storage', storage_assets, '--delete', '1'
      ]
      cmd_args = [
          '--project', project_id, '--storage', storage_logs, '--delete', '30'
      ]
      RunPython(create_storage_py, cmd_args)

    # Configure the project + ACLs
    cmd_args = [
        '--project', project_id, '--storage', storage_assets,
        '--service_accounts', args.accounts
    ]
    RunPython(config_project_py, cmd_args)

    # Configure scheduler to stop old compute instances (for fatal failures).
    cmd_args = ['--project', project_id]
    RunPython(config_cronjob_py, cmd_args)

  sys.exit(0)
