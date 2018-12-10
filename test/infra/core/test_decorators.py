# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from test_case import EnterpriseTestCase


class environment(object):
  """@environment annotation for test case classes."""

  def __init__(self, file):
    self.file = file

  def __call__(self, _class):
    if not issubclass(_class, EnterpriseTestCase):
      error = '@environment can only be used on EnterpriseTestCase subclasses!'
      raise Exception(error)

    _class.ASSET_FILE = self.file
    return _class


def test(_method):
  """@test annotation for test methods."""
  _method.IS_TEST_METHOD = True
  return _method