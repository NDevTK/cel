# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""
Script to run a UI test.

Sample usage:
  run_ui_test.py python c:\temp\ui_test.py

The script will call the cel_ui_agent to run command:
  python c:\temp\ui_test.py
"""

import json
import logging
import subprocess
import sys
import time

from absl import app, flags

import requests

FLAGS = flags.FLAGS
flags.DEFINE_integer('timeout', None, 'timeout, in seconds')
flags.mark_flag_as_required('timeout')


def ConfigureLogging():
  logfmt = '%(asctime)s %(filename)s:%(lineno)s: [%(levelname)s] %(message)s'
  datefmt = '%Y/%m/%d %H:%M:%S'

  logging.basicConfig(level=logging.INFO, format=logfmt, datefmt=datefmt)


def RunUITest(cmd):
  """Runs a UI test.

  Args:
     cmd: the command to run.
  """
  request = {"command": cmd, "timeout": FLAGS.timeout}
  logging.info("Request: %s", request)
  response = requests.post(
      'http://localhost:9000/Run', data=json.dumps(request))
  if not response.ok:
    response.raise_for_status()

  r = json.loads(response.content)
  if r['Status'] != 0:
    raise RuntimeError('Error starting commmand: {}'.format(r))

  # wait for command to finish
  while True:
    response = requests.get('http://localhost:9000/RunStatus')
    if not response.ok:
      response.raise_for_status()

    r = json.loads(response.content)
    status = r['Status']
    if status == 0:
      # finish without error
      print(r['Output'].encode('utf-8'))
      break
    elif status == 2:
      # command is still running. Wait for a while before rechecking
      # the status.
      logging.info('Command running')
      time.sleep(30)
    else:
      # error occurred
      raise RuntimeError('Error occurred: {}'.format(r))


def main(argv):
  ConfigureLogging()
  RunUITest(subprocess.list2cmdline(argv[1:]))


if __name__ == '__main__':
  app.run(main)