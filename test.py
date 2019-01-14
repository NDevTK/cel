# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import logging
import os
import sys
import traceback
import warnings

try:
  import test.infra.controller as controller
except ImportError as e:
  nonDefaultModules = ['googleapiclient.discovery', 'google.iam.admin.v1']
  for module in nonDefaultModules:
    if module in e.message:
      print('Failed to import %s in test.py.' % module)
      print('Try running  `python ./scripts/tests/setup.py`.\n')
  raise


def ParseArgs():
  example = '%s test.tests.IISSitesTest test.host.textpb' % sys.argv[0]

  parser = argparse.ArgumentParser(
      description='Test runner for CELab', epilog='example: %s' % example)

  parser.add_argument(
      '--test',
      required=True,
      metavar='<test_class>',
      help='The full class name of the EnterpriseTestCase class (w/ package)')
  parser.add_argument(
      '--host',
      required=True,
      metavar='<host_file>',
      help='The full path to the *.host.textpb file to use')
  parser.add_argument(
      '--cel_ctl',
      metavar='<path>',
      dest='cel_ctl',
      default=GetDefaultCelCtl(),
      action='store',
      help='Which binary to use to deploy the environment')
  parser.add_argument(
      '--noprogress',
      dest='show_progress',
      default=True,
      action='store_false',
      help='Don\'t show progress during the deployment phase')
  parser.add_argument(
      '--nodeploy',
      dest='deploy',
      default=True,
      action='store_false',
      help='Skip the deployment phase and go straight to tests')
  parser.add_argument(
      '--cleanup',
      dest='cleanup',
      default=False,
      action='store_true',
      help='Clean up the host environment after the test')
  parser.add_argument(
      '--error_logs_dir',
      metavar='<path>',
      dest='error_logs_dir',
      default=None,
      action='store',
      help='Where to collect extra logs on test failures')
  parser.add_argument(
      '-v',
      '--verbose',
      dest='verbose',
      action='store_true',
      help='Show info logs')
  parser.add_argument(
      '-vv',
      '--debug',
      dest='debug',
      action='store_true',
      help='Show debug and info logs')

  return parser.parse_args()


def GetDefaultCelCtl():
  # TODO: Add Windows support
  return os.path.join(sys.path[0], 'out/linux_amd64/bin/cel_ctl')


def ConfigureLogging(args):
  level = logging.WARNING
  if args.verbose:
    level = logging.INFO
  if args.debug:
    level = logging.DEBUG

  # Filter out logs from low level loggers
  errorOnlyLoggers = [
      'googleapiclient.discovery_cache', 'google.auth', 'google_auth_httplib2'
  ]
  for logger in errorOnlyLoggers:
    logging.getLogger(logger).setLevel(logging.ERROR)
  message = 'We recommend that most server applications use service accounts.'
  warnings.filterwarnings('ignore', '.*%s' % message)

  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=level, format=logfmt, datefmt=datefmt)


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  c = controller.SingleTestController(args.test, args.host, args.cel_ctl)

  success = False
  should_write_logs = (args.error_logs_dir != None)
  try:
    if args.deploy:
      c.DeployNewEnvironment(args.show_progress)

    success = c.ExecuteTestCase()
  except controller.CelCtlError:
    logging.error('Test failed to run cel_ctl.')
    should_write_logs = False
  except KeyboardInterrupt:
    logging.error('Test aborted.')
    should_write_logs = False
  except:
    print(traceback.format_exc())
    logging.error('Test failed.')
  finally:
    if not success and should_write_logs:
      print('Writing Compute logs to "%s"...' % args.error_logs_dir)
      c.TryWriteComputeLogsTo(args.error_logs_dir)

    if args.cleanup:
      print('Cleaning up host environment...')
      c.TryCleanHostEnvironment()

  sys.exit(0 if success else 1)
