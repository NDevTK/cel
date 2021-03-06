# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from __future__ import print_function
import argparse
import logging
import os
import sys
import traceback
import warnings
from absl import app
from absl import flags

FLAGS = flags.FLAGS

flags.DEFINE_string(
    'test', None,
    'The full class name of the EnterpriseTestCase class (w/ package)')
flags.mark_flag_as_required('test')

flags.DEFINE_string('test_filter', None,
                    'The name of the test to run in the test class')

flags.DEFINE_string('host', None,
                    'The full path to the *.host.textpb file to use')
flags.mark_flag_as_required('host')

flags.DEFINE_string('cel_ctl', None,
                    'Which binary to use to deploy the environment')
flags.DEFINE_bool(
    'deploy', True, 'Depoly the test environment. '
    'Set to false to skip the deployment phase and go straight to tests')
flags.DEFINE_bool(
    'skip_before_all', False, 'True to skip @before_all methods. '
    'Like --nodeploy, this is used to skip set up steps. '
    'Useful when developing new tests.')
flags.DEFINE_bool(
    'no_external_access', False, 'True to skip creating RDP/SSH firewall '
    'rules during deployment. Should be used in automated test runs.')
flags.DEFINE_bool('cleanup', False,
                  'Clean up the host environment after the test')
flags.DEFINE_string('error_logs_dir', None,
                    'Where to collect extra logs on test failures')
flags.DEFINE_multi_string('test_arg', None, 'Flags passed to tests')

try:
  import chrome_ent_test.infra.controller as controller
except ImportError as e:
  nonDefaultModules = ['googleapiclient.discovery', 'google.iam.admin.v1']
  for module in nonDefaultModules:
    if module in str(e):
      print('Failed to import %s in test.py.' % module)
      print('Try running  `python ./scripts/tests/setup.py`.\n')
  raise


def GetDefaultCelCtl():
  if sys.platform == 'win32':
    return os.path.join(sys.path[0], '../out/windows_amd64/bin/cel_ctl.exe')
  else:
    return os.path.join(sys.path[0], '../out/linux_amd64/bin/cel_ctl')


def ConfigureLogging():
  # Filter out logs from low level loggers
  errorOnlyLoggers = [
      'googleapiclient.discovery_cache', 'google.auth', 'google_auth_httplib2'
  ]
  for logger in errorOnlyLoggers:
    logging.getLogger(logger).setLevel(logging.ERROR)
  message = 'We recommend that most server applications use service accounts.'
  warnings.filterwarnings('ignore', '.*%s' % message)

  logging.error("%s: Logging level error is visible." % __file__)
  logging.warning("%s: Logging level warning is visible." % __file__)
  logging.info("%s: Logging level info is visible." % __file__)
  logging.debug("%s: Logging level debug is visible." % __file__)


def main(argv):
  if FLAGS.cel_ctl is None:
    FLAGS.cel_ctl = GetDefaultCelCtl()

  ConfigureLogging()

  c = controller.SingleTestController(
      FLAGS.test,
      FLAGS.host,
      FLAGS.cel_ctl,
      test_filter=FLAGS.test_filter,
      skip_before_all=FLAGS.skip_before_all,
      no_external_access=FLAGS.no_external_access)

  # Parse test specific flags. Note that we need to use a dummy element
  # as the first element of the list since absl.flags ignores the first element
  # during parsing.
  if FLAGS.test_arg is not None:
    FLAGS([''] + FLAGS.test_arg)

  success = False
  should_write_logs = (FLAGS.error_logs_dir != None)
  try:
    if FLAGS.deploy:
      c.DeployNewEnvironment()

    success = c.ExecuteTestCase()
  except KeyboardInterrupt:
    logging.error('Test aborted.')
    should_write_logs = False
  except:
    print(traceback.format_exc())
    logging.error('Test failed.')
  finally:
    if not success and should_write_logs:
      print('Writing Compute logs to "%s"...' % FLAGS.error_logs_dir)
      c.TryWriteComputeLogsTo(FLAGS.error_logs_dir)

    if FLAGS.cleanup:
      print('Cleaning up host environment...')
      c.TryCleanHostEnvironment()

  sys.exit(0 if success else 1)


if __name__ == '__main__':
  app.run(main)
