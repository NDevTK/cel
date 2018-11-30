# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging


class EnterpriseTestCase:
  """Base class for tests that provides test hooks and resources."""

  def __init__(self, environment):
    logging.info('Initialize Test=%s with %s' % (self.__class__, environment))
    self.clients = environment.clients

  @staticmethod
  def GetTestMethods(_class):
    test_methods = []
    for _, elem in _class.__dict__.items():
      if hasattr(elem, 'IS_TEST_METHOD'):
        test_methods.append(elem)
    return test_methods
