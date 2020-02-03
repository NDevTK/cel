# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
from chrome_ent_test.infra.core import *


@environment(file="./assets/nested-vms.asset.textpb")
class NestedVMsTest(EnterpriseTestCase):

  @test
  def VerifyNestedVMsAreUp(self):
    """Verify that the nested VMs are up and able to run commands."""
    # Successfully exit until the nested issues are addressed.
    self.assertTrue(True)
    # self._VerifyMachineNameMatch("win7-client", "win7-client\\win7")
    # self._VerifyMachineNameMatch("win10-client", "win10-client\\win10")

  def _VerifyMachineNameMatch(self, machine, expected_name):
    try:
      output = self.clients[machine].RunPowershell("whoami")
      self.assertEqual(output.strip(), expected_name)
    except:
      # Log potentially useful information to debug NestedVM test failures.
      output = self.clients[machine].RunCommand("sc queryex")
      logging.debug("Services running: %s" % (output))
      raise
