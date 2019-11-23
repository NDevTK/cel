# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from chrome_ent_test.infra.core import *


@environment(file="./assets/rds-examples.asset.textpb", timeout=5400)
class RemoteDesktopTest(EnterpriseTestCase):

  @test
  def VerifyRDS(self):
    for rds_host in ['win2008-rds', 'win2012-rds', 'win2016-rds']:
      script = "query termserver %s" % rds_host
      ret, output = self.clients['win2008-dc'].RunPowershellNoThrow(script)

      # This doesn't return 0 when it finds the RDS host. Parse the output.
      successMsg = "Known Remote Desktop Session Host servers"
      self.assertTrue(successMsg in output)
      self.assertTrue(rds_host.upper() in output)

      errorMsg = "This Remote Desktop Sesion Host server was not found"
      self.assertTrue(errorMsg not in output)
