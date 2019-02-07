# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from test.infra.core import *


@environment(file="./assets/domain-join.asset.textpb")
class DomainJoinTest(EnterpriseTestCase):

  @test
  def VerifyDomainJoined(self):
    for client in ['client2008', 'client2012', 'client2016']:
      ret, output = self.clients[client].RunCommand("systeminfo")

      domain = ADTestHelper._GetDomainFromSystemInfo(output)

      self.assertEqual(domain, "test1.com")


@environment(file="./assets/domain-tree.asset.textpb", timeout=7200)
class DomainTreeTest(EnterpriseTestCase):

  @test
  def VerifyDomainTreeExists(self):
    # There are 2 AD trees.
    #
    # The first one looks like this:
    #            a1-2012.com
    #         /             \
    #     b1.a1-2012.com   a2-2012.com
    #         |
    #    c1.b1.a1-2012.com

    # The second one looks like this:
    #            a1-2008.com
    #         /              \
    #     b1.a1-2008.com   a2-2008.com
    #         |
    #    c1.b1.a1-2008.com
    cases = []
    cases.append(('a1-2012dc', 'a1-2012.com'))
    cases.append(('a2-2012dc', 'a2-2012.com'))
    cases.append(('b1-2012dc', 'b1.a1-2012.com'))
    cases.append(('c1-2012dc', 'c1.b1.a1-2012.com'))

    cases.append(('a1-2008dc', 'a1-2008.com'))
    cases.append(('a2-2008dc', 'a2-2008.com'))
    cases.append(('b1-2008dc', 'b1.a1-2008.com'))
    cases.append(('c1-2008dc', 'c1.b1.a1-2008.com'))

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
