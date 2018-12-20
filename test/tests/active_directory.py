# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from test.infra.core import *


@environment(file="./examples/schema/ad/domain-join.asset.textpb")
class DomainJoinTest(EnterpriseTestCase):

  @test
  def VerifyDomainJoined(self):
    for client in ['client2008', 'client2012']:
      ret, output = self.clients[client].RunCommand("systeminfo")

      domain = ADTestHelper._GetDomainFromSystemInfo(output)

      self.assertEqual(domain, "test1.com")


@environment(file="./examples/schema/ad/domain-tree.asset.textpb", timeout=7200)
class DomainTreeTest(EnterpriseTestCase):

  @test
  def VerifyDomainTreeExists(self):
    # The AD Tree looks like this:
    #      a1.com
    #         |
    #     b1.a1.com
    #         |
    #    c1.b1.a1.com
    cases = []
    cases.append(('a1dc', 'a1.com'))
    cases.append(('b1dc', 'b1.a1.com'))
    cases.append(('c1dc', 'c1.b1.a1.com'))

    for client, expectedDomain in cases:
      ret, output = self.clients[client].RunCommand("systeminfo")

      domain = ADTestHelper._GetDomainFromSystemInfo(output)

      self.assertEqual(domain, expectedDomain)


class ADTestHelper:

  @staticmethod
  def _GetDomainFromSystemInfo(output):
    for line in output.splitlines():
      if line.startswith("Domain:"):
        parts = line.split(":")
        if len(parts) != 2:
          raise Exception("Unexpected format for domain line: %s" % repr(line))

        return parts[1].strip()

    return None
