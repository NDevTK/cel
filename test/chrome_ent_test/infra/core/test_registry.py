# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging


class TestRegistry:
  """Class that registers EnterpriseTestCases tagged with @environment."""
  _allTestCases = {}

  @staticmethod
  def Register(name, _class):
    """Registers a TestCase.

    Args:
      name: the qualified name of the test case class to register.
            ex: chrome_ent_test.tests.active_directory.DomainJoinTest
      _class: the class object of the test case.
    """
    if name in TestRegistry._allTestCases:
      error = '%s can not be registered twice (%s, %s)!' % (
          name, TestRegistry._allTestCases[className], _class)
      raise Exception(error)

    TestRegistry._allTestCases[name] = _class

  @staticmethod
  def Find(name):
    """Finds a class that matches a name.

    Returns:
      The full identifier for the class if found. None if not found.
    """
    for c in TestRegistry._allTestCases:
      if name == c:
        return c

    return None

  @staticmethod
  def FindAll(prefix):
    """Find all classes that start with a certain prefix.

    Returns:
      A list of the classes that start with that prefix. Empty list if none.
    """
    results = []

    for name in TestRegistry._allTestCases:
      if name.startswith(prefix):
        results.append(name)

    return results

  @staticmethod
  def Filter(tests, includes=[], excludes=[]):
    """Filter tests by including/excluding specific @category annotations.

    Args:
      tests: list of qualified class names of tests registered in the registry.
      includes: list of categories to include.
                if it's an empty list, we imply include-all.
      excludes: list of categories to exclude.

    Notes:
     - A non-empty `includes` acts like a union.
     - An exclude has priority over every include.

    Examples:
      Given tests=[
              TestA(categories=(A)),
              TestAB(categories=(A,B)),
              TestB(categories=(B))
            ]
      Filter for include=[A] would return TestA and TestAB.
      Filter for include=[A,B] would return all three tests.
      Filter for include=[] would return all three tests.
      Filter for exclude=[B] would return TestA only.

    Returns:
      A list of the test cases that satisfies include and exclude filters.
    """
    results = []

    for name in tests:
      # Shouldn't happen if `tests` comes from Find/FindAll, but who knows.
      if name not in TestRegistry._allTestCases:
        error = '%s could not be found in the registry.' % name
        raise Exception(error)

      categories = []
      if hasattr(TestRegistry._allTestCases[name], 'CATEGORIES'):
        categories = TestRegistry._allTestCases[name].CATEGORIES

      hasInclude = False
      hasExclude = False
      for category in categories:
        if category in includes:
          hasInclude = True

        if category in excludes:
          hasExclude = True

      if hasExclude:
        logging.info("Skipping %s because one of %s is in `excludes`." %
                     (name, categories))
        continue

      if includes:
        if not hasInclude:
          logging.info("Ignoring %s because none of %s are in `includes`." %
                       (name, categories))
          continue

      results.append(name)

    return results

  @staticmethod
  def AllTests():
    """Return the list of all available tests."""
    return TestRegistry._allTestCases.keys()
