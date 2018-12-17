# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from test.infra.core import *


@environment(file="./examples/schema/ad/domain-join.asset.textpb")
class DomainJoinTest(EnterpriseTestCase):

  @test
  def VerifyDomainJoined(self):
    pass


@environment(file="./examples/schema/ad/domain-tree.asset.textpb", timeout=7200)
class DomainTreeTest(EnterpriseTestCase):

  @test
  def VerifyDomainTreeExists(self):
    pass


# TODO: Fix bug: This asset file makes cel_ctl crash.
@environment(file="./examples/schema/ad/two-domains.asset.textpb")
class TwoDomainsTest(EnterpriseTestCase):

  @test
  def VerifyTwoDomainsExist(self):
    pass
