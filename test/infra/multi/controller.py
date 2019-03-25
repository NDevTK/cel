# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""This is the main module for a multi-test run.

It invokes ./test.py for different (test, host) pairs

Here are the different threads running in a multi-test run:
  Main loop (ExecuteTestCases):
    This loop runs on the main thread and is the entry point for run_tests.py.
    It iterates through all TestCases we need to run and starts one Test thread
    for every test as hosts become available in the HostPool. It blocks when
    it's waiting for an available host and until all tests are done.

  Test threads (TestWorkerThread):
    One TestWorkerThread is created per TestCase. Those threads are responsible
    for starting ./test.py with the appropriate arguments and reporting the
    result. The number of test threads that can run at the same time is limited
    by the HostPool max size or the HostProvider's capacity to provide hosts.
    When a test thread completes, it invokes a callback put its host back in
    the pool for the main loop to use.

  Progress thread (DisplayProgressThread):
    This thread is an optional thread (--noprogress) that prints test progress.
    It's meant to give interactive users a sense of the general state of the
    test run as it can run for a long time without printing anything.
    That thread doesn't access anything significant and only outputs to stdout
    using the main print lock that other threads use.
"""

from test.infra.multi.host_pool import HostPool
from test.infra.multi.threads import TestWorkerThread, DisplayProgressThread
import json
import logging
import os
import threading


class MultiTestController:

  def __init__(self, tests, hostProvider, errorLogsDir=None):
    """Initializes the controller for a multi-test run.

    Args:
      tests: A list of existing test classes that will be passed to ./test.py.
      hostProvider: A HostProvider that implements TryAcquire & Release.
    """
    self._testsToRun = tests
    self._hostPool = HostPool(hostProvider)
    self._errorLogsDir = errorLogsDir

    self._activeTestThreads = {}

    self._testsSummary = {}
    self._totalTestsPassed = 0
    self._totalTestsFailed = 0

    # All threads should use this lock when printing something to stdout
    self._printLock = threading.Lock()

  def ExecuteTestCases(self, showProgress=False):
    try:
      return self._ExecuteTestCases(showProgress)
    except:
      raise
    finally:
      # Save a summary of test results for LUCI to consume.
      if self._errorLogsDir != None:
        summaryFilePath = os.path.join(self._errorLogsDir, "summary.json")
        with open(summaryFilePath, "w") as summaryFile:
          json.dump(self._testsSummary, summaryFile)

  def _ExecuteTestCases(self, showProgress):
    """This is our main loop and entry point for ./run_tests.py."""
    progressThread = None
    if showProgress:
      progressThread = DisplayProgressThread(self)
      progressThread.start()

    # Get an available host for each test and start test threads.
    for test in self._testsToRun:
      host = self._hostPool.TakeOrWaitForHost()

      # Start a TestWorker thread to run this test
      callback = self._OnTestWorkerThreadCompleted
      thread = TestWorkerThread(test, host, self._errorLogsDir, callback)
      thread.start()
      self._activeTestThreads[thread.ident] = thread

    # We don't have more tests to run so we don't need new hosts. This releases
    # any available hosts and also flips a flag to release returning ones.
    self._hostPool.ReleaseAllAvailableHosts()

    # Wait for running TestWorkerThreads to complete
    for testThread in self._activeTestThreads.values():
      testThread.join()

    if progressThread != None:
      progressThread.Stop()

    # At this point, all test threads should've invoked our callback.
    for threadID in self._activeTestThreads:
      message = "Thread %s never invoked its callback." % threadID
      message += " [test=%s]" % self._activeTestThreads[threadID].test
      logging.warning(message)

    success = (self._totalTestsPassed == len(self._testsToRun))

    # Print summary
    results = (self._totalTestsPassed, len(self._testsToRun))
    summary = "\n%s/%s test cases passed.\n" % results
    with self._printLock:
      print(summary)

    return success

  def _OnTestWorkerThreadCompleted(self, thread, success, details):
    """This is called by a TestWorkerThread when it completes."""
    self._PrintResult(thread.test, success, details)
    self._hostPool.AddAvailableHost(thread.host)

    # Save the test.py logs for LUCI to consume.
    if self._errorLogsDir != None:
      outputFilePath = os.path.join(self._errorLogsDir,
                                    thread.test + ".output.txt")
      with open(outputFilePath, "w") as outputFile:
        for line in details:
          outputFile.write(line + "\n")

      self._testsSummary[thread.test] = {
          'success': success,
          'output': outputFilePath
      }

    if success:
      self._totalTestsPassed += 1
    else:
      self._totalTestsFailed += 1

    if thread.ident in self._activeTestThreads:
      del self._activeTestThreads[thread.ident]
    else:
      logging.warning("TestWorkerThreadCompleted called from unknown thread!")

  def _PrintResult(self, test, success, details):
    """Prints the results of a test.

    This function can be called from multiple threads and guarantees that each
    result is printed together on dedicated lines, including the final newline.
    `print` doesn't offer this guarantee which can produce weird output.

    This might not always play nicely with --verbosity.
    """
    with self._printLock:
      if success:
        print("PASSED  %s" % test)
      else:
        print("FAILED  %s" % test)

      for line in details:
        print("  %s" % line)
