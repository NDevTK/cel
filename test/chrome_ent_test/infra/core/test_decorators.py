# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from __future__ import absolute_import
from .test_case import EnterpriseTestCase
from .test_registry import TestRegistry


class environment(object):
  """@environment annotation for test case classes."""

  def __init__(self, file, timeout=None):
    self.file = file
    self.timeout = timeout

  def __call__(self, _class):
    if not issubclass(_class, EnterpriseTestCase):
      error = '@environment can only be used on EnterpriseTestCase subclasses!'
      raise Exception(error)

    className = "%s.%s" % (_class.__module__, _class.__name__)
    TestRegistry.Register(className, _class)
    _class.ASSET_FILE = self.file
    _class.DEPLOY_TIMEOUT = self.timeout
    return _class


class category(object):
  """@category annotation for test classes."""

  def __init__(self, category):
    self.category = category

  def __call__(self, _class):
    if not issubclass(_class, EnterpriseTestCase):
      error = '@category can only be used on EnterpriseTestCase subclasses!'
      raise Exception(error)

    if not hasattr(_class, 'CATEGORIES'):
      _class.CATEGORIES = []

    _class.CATEGORIES.append(self.category)
    return _class


def test(_method):
  """@test annotation for test methods."""
  _method.IS_TEST_METHOD = True
  return _method


def before_all(_method):
  """@before_all annotation.

  Such method is run right before any test methods are run."""
  _method.IS_BEFORE_ALL = True
  return _method
