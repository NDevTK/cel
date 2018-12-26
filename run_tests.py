# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import logging
import sys
from test.infra.multi import *
import traceback

# Import all known tests
from test.tests import *


def ParseArgs():
  example = '%s --hosts ./hosts/' % sys.argv[0]

  parser = argparse.ArgumentParser(
      description='Test suite runner for CELab', epilog='example: %s' % example)

  parser.add_argument(
      '--hosts',
      required=True,
      metavar='<host_file;...>',
      help='Full paths to *.host.textpb files to use for tests (or directory)')
  parser.add_argument(
      '--tests',
      metavar='<test_class;...>',
      default='*',
      help='Fully qualified names of TestCases to run (supports my.package.*)')
  parser.add_argument(
      '--noprogress',
      dest='show_progress',
      default=True,
      action='store_false',
      help='Don\'t show progress while running tests')
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


def ConfigureLogging(args):
  level = logging.WARNING
  if args.verbose:
    level = logging.INFO
  if args.debug:
    level = logging.DEBUG

  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=level, format=logfmt, datefmt=datefmt)


if __name__ == '__main__':
  args = ParseArgs()

  ConfigureLogging(args)

  logging.info("Arguments: %s" % args)

  tests = ArgsParser.ParseTestsArg(args.tests)
  logging.debug('Found tests: %s', tests)

  hostFiles = ArgsParser.ParseHostsArg(args.hosts)
  logging.debug('Found hosts: %s', hostFiles)

  c = MultiTestController(tests, SimpleHostProvider(hostFiles))

  success = False
  try:
    success = c.ExecuteTestCases(args.error_logs_dir, args.show_progress)
  except KeyboardInterrupt:
    logging.error('Test run aborted.')
    should_write_logs = False
  except:
    print(traceback.format_exc())
    logging.error('Test run failed.')

  sys.exit(0 if success else 1)