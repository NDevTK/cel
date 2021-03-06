# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
from chrome_ent_test.infra.core import *


@category("core")
@environment(file="./assets/registry-keys.asset.textpb")
class MachineRegistryTest(EnterpriseTestCase):

  @test
  def VerifyRegistryKeys(self):
    """Verify that the registry keys were set correctly."""
    self.assertRegistryContains(
        "win2008-regkeys", "HKLM:\\System\\Foo", {
            'SomeStringKey': 'Some string value',
            'SomeExpandStringKey': 'Another string value',
            'SomeBinaryKey': '{1, 2, 3}',
            'SomeDwordKey': '123',
            'SomeQwordKey': '456',
            'SomeMultiStringKey': '{First, Second, Third}'
        })

    self.assertRegistryContains("win2008-regkeys", "HKLM:\\System\\Bar",
                                {'FooBar': '1'})

    self.assertRegistryContains("win10-regkeys", "HKLM:\\System\\Bar",
                                {'FooBar10': '1'})

  def assertRegistryContains(self, machine, path, expected):
    logging.debug("Checking %s for %s" % (path, expected))

    output = self.clients[machine].RunPowershell("""
      Get-ItemProperty "%s"
    """ % path)

    actual = self.parseRegistryValues(output)

    for key, value in expected.items():
      self.assertTrue(
          key in actual,
          "Expected key (%s) missing from Get-ItemProperty output." % key)
      self.assertEqual(actual[key], expected[key])

  def parseRegistryValues(self, output):
    """Parses the output of a Get-ItemProperty call for a registry key.

    Expected output example:
      SomeStringKey       : Some string value
      SomeExpandStringKey : Another string value
      SomeBinaryKey       : {1, 2, 3}
      SomeDwordKey        : 123
      SomeQwordKey        : 456
      SomeMultiStringKey  : {First, Second, Third}
      PSPath              : Microsoft.PowerShell.Core\Registry::HKEY_LOCAL_MACH
                            INE\System\Foo
      PSParentPath        : Microsoft.PowerShell.Core\Registry::HKEY_LOCAL_MACH
                            INE\System
      PSChildName         : Foo
      PSDrive             : HKLM
      PSProvider          : Microsoft.PowerShell.Core\Registry
    """
    values = {}

    lines = output.splitlines()
    for line in lines:
      parts = line.split(':')
      if len(parts) == 2:
        values[parts[0].strip()] = parts[1].strip()

    return values
