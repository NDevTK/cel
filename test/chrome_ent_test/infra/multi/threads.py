# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime
import logging
import os
import subprocess
import sys
import threading
import time


class TestWorkerThread(threading.Thread):
  """Runs ./test.py for a given (test, host) pair."""

  def __init__(self, test, host, testPy, testPyArgs, errorLogsDir, callback):
    if callback is None:
      raise TypeError("`callback` cannot be None.")

    threading.Thread.__init__(self)
    self.daemon = True

    self.test = test
    self.host = host

    self._testPy = testPy
    self._testPyArgs = testPyArgs
    self._errorLogsDir = errorLogsDir
    self._callback = callback

  def run(self):
    cmd = [
        'python', self._testPy, '--test', self.test, '--host', self.host,
        '--cleanup'
    ]

    if self._testPyArgs != None:
      cmd += self._testPyArgs.split()

    if self._errorLogsDir != None:
      cmd += ['--error_logs_dir', os.path.join(self._errorLogsDir, self.test)]

    # Pass down the current verbosity to ./test.py
    level = logging.getLogger().getEffectiveLevel()
    if level == logging.WARNING:
      cmd += ['-v', '-1']
    elif level == logging.INFO:
      cmd += ['-v', '0']
    elif level == logging.DEBUG:
      cmd += ['-v', '1']

    logging.info("Running command %s" % cmd)
    success = False
    details = []
    try:
      output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
      success = True

      # Relevant output will be formatted by MultiTestController._PrintResult.
      logging.debug("Command %s succeeded." % cmd)
      details = self._ParseOutputForTestDetails(output)
    except subprocess.CalledProcessError as e:
      # Full output will be formatted by MultiTestController._PrintResult.
      logging.debug("Failed %s failed." % cmd)
      details = e.output.splitlines()
    finally:
      self._callback(self, success, details)

  def _ParseOutputForTestDetails(self, output):
    details = []

    for line in output.splitlines():
      if line.startswith("PASSED") or line.startswith("FAILED"):
        details.append(line)

    return details


class DisplayProgressThread(threading.Thread):
  """Prints and refreshed the current progress of a test run."""

  def __init__(self, controller, interval=2):
    """Initializes the DisplayProgress thread.

    Args:
      controller: The controller from which to gather progress information.
      interval: The refresh interval at which the thread will update stdout.
    """
    threading.Thread.__init__(self)
    self.daemon = True

    self._controller = controller
    self._interval = interval

    self._abort = False
    self._completedEvent = threading.Event()

  def run(self):
    c = self._controller
    while not self._abort:
      # It's possible to temporarily double-count a test that still has an
      # active thread and just updated the _totalX value. This is not critical
      # and just an FYI for users so it doesn't matter.
      completed = c._totalTestsFailed + c._totalTestsPassed
      running = len(c._activeTestThreads)
      left = len(c._testsToRun) - completed - running

      message = "%s tests running, %s left." % (running, left)
      now = datetime.datetime.now().strftime('%H:%M:%S')

      with c._printLock:
        sys.stdout.write("[%s] %s\r" % (now, message))
        sys.stdout.flush()

      time.sleep(self._interval)

    self._completedEvent.set()

  def Stop(self):
    """Stops the run() loop and clears the last progress trace."""
    self._abort = True

    # Wait for the final run() iteration to finish
    self._completedEvent.wait(self._interval * 2)

    # Clear the current line (only has text if it's the progress line)
    with self._controller._printLock:
      sys.stdout.write("%s\r" % (' ' * 80))
      sys.stdout.flush()
