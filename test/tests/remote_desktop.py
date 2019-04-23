# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from test.infra.core import *


@environment(file="./assets/rds-examples.asset.textpb", timeout=5400)
class RemoteDesktopTest(EnterpriseTestCase):

  @test
  def VerifyRDS(self):
    # TODO(crbug.com/951596): Intermittent RDS deployment failures on Win2012.
    for version in ['win2008']:
      for _from, to in [('dc', 'client'), ('client', 'dc')]:
        clientName = "%s-%s" % (version, _from)
        targetName = "%s-%s" % (version, to)

        script = "query termserver %s" % targetName
        ret, output = self.clients[clientName].RunPowershellNoThrow(script)

        # This doesn't return 0 when it finds the RDS host. Parse the output.
        successMsg = "Known Remote Desktop Session Host servers"
        self.assertTrue(successMsg in output)
        self.assertTrue(targetName.upper() in output)

        errorMsg = "This Remote Desktop Sesion Host server was not found"
        self.assertTrue(errorMsg not in output)
