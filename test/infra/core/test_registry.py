# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging


class TestRegistry:
  """Class that registers EnterpriseTestCases tagged with @environment."""
  _allTestCases = []

  @staticmethod
  def Register(_class):
    """Registers a TestCase."""
    TestRegistry._allTestCases.append(_class)

  @staticmethod
  def Find(_class):
    """Finds a class that matches a name.

    Returns:
      The full identifier for the class if found. None if not found.
    """
    for c in TestRegistry._allTestCases:
      if _class == c:
        return c

    return None

  @staticmethod
  def FindAll(prefix):
    """Find all classes that start with a certain prefix.

    Returns:
      A list of the classes that start with that prefix. Empty list if none.
    """
    results = []

    for _class in TestRegistry._allTestCases:
      if _class.startswith(prefix):
        results.append(_class)

    return results
