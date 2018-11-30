# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
from test.infra.core import *


@environment(file="./examples/schema/iis/iis-various-auth.asset.textpb")
class IISSitesTest(EnterpriseTestCase):

  @test
  def VerifyAnonymousSite(self):
    # TODO: Verify that the anonymous site is up
    logging.info("Executing VerifyAnonymousSite (NYI)")
    pass

  @test
  def VerifyNTLMSite(self):
    # TODO: Verify that the NTLM site is up
    logging.info("Executing VerifyNTLMSite (NYI)")
    pass

  # TODO: Kerberos, NTLMv1 vs NTLMv2, ...
