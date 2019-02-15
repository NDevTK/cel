# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
from test.infra.core import *


@environment(file="./assets/nested-vms.asset.textpb")
class NestedVMsTest(EnterpriseTestCase):

  @test
  def VerifyNestedVMsAreUp(self):
    """Verify that the nested VMs are up and able to run commands."""
    ret, output = self.clients["win7-client"].RunPowershell("whoami")
    self.assertEqual(ret, 0)
    self.assertEqual(output.strip(), "win7-client\\win7")

    ret, output = self.clients["win10-client"].RunPowershell("whoami")
    self.assertEqual(ret, 0)
    self.assertEqual(output.strip(), "win10-client\\win10")
