# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import unittest

from chrome_ent_test.infra.core.test_registry import TestRegistry


class TestRegistryFilterTests(unittest.TestCase):
  """Unit tests for TestRegistry methods."""

  @classmethod
  def setUpClass(_class):
    TestRegistry.Register("classA", TestClass(['a']))
    TestRegistry.Register("classB", TestClass(['b']))
    TestRegistry.Register("classAB", TestClass(['a', 'b']))
    TestRegistry.Register("classNone", object())

  def test_FindAllIncludeA(self):
    classes = TestRegistry.AllTests()
    classes = TestRegistry.Filter(classes, ['a'], [])

    self.assertEqual(2, len(classes))
    self.assertIn('classA', classes)
    self.assertIn('classAB', classes)

  def test_FindAllIncludeB(self):
    classes = TestRegistry.AllTests()
    classes = TestRegistry.Filter(classes, ['b'], [])

    self.assertEqual(2, len(classes))
    self.assertIn('classB', classes)
    self.assertIn('classAB', classes)

  def test_FindAllIncludeC(self):
    classes = TestRegistry.AllTests()
    classes = TestRegistry.Filter(classes, ['c'], [])

    self.assertEqual(0, len(classes))

  def test_FindAllExcludeB(self):
    classes = TestRegistry.AllTests()
    classes = TestRegistry.Filter(classes, [], ['b'])

    self.assertEqual(2, len(classes))
    self.assertIn('classA', classes)
    self.assertIn('classNone', classes)

  def test_FindAllIncludeAExcludeB(self):
    classes = TestRegistry.AllTests()
    classes = TestRegistry.Filter(classes, ['a'], ['b'])

    self.assertEqual(1, len(classes))
    self.assertIn('classA', classes)

  def test_FindAllAIncludeB(self):
    classes = TestRegistry.FindAll("classA")
    classes = TestRegistry.Filter(classes, ['b'])

    self.assertEqual(1, len(classes))
    self.assertIn('classAB', classes)

  def test_FindNoneIncludeAndExcludeNoOp(self):
    classes = TestRegistry.FindAll("bar")
    self.assertEqual(0, len(classes))

    classes = TestRegistry.Filter(classes, ['a'], ['b'])
    self.assertEqual(0, len(classes))


class TestClass:

  def __init__(self, categories):
    self.CATEGORIES = categories
