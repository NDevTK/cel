# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import re
import time
from chrome_ent_test.infra.core import *


@environment(file="./assets/iis-various-auth.asset.textpb")
class IISSitesTest(EnterpriseTestCase):

  @test
  def VerifyAnonymousSite(self):
    cases = self._GetDomainTestCases(8082)

    # Add non-domain cases for anonymous sites
    cases.append(
        IISTestHelper._TestCase('no-domain-2012',
                                'http://no-domain-2012:8082/test.aspx'))
    cases.append(
        IISTestHelper._TestCase('no-domain-2016',
                                'http://no-domain-2016:8082/test.aspx'))

    for case in cases:
      logging.info("Verify %s can reach %s" % (case.client, case.target))
      script = 'Invoke-WebRequest %s -UseBasicParsing' % case.target
      script = '(%s).Content' % script
      output = self.clients[case.client].RunPowershell(script)
      IISTestHelper._AssertTitleEquals(self, case, output, "[Anonymous]")

  @test
  def VerifyNTLMSite(self):
    cases = self._GetDomainTestCases(8081)

    for case in cases:
      logging.info("Verify %s with NTLM" % case)
      IISTestHelper._VerifyDefaultCredentialAccess(self, case,
                                                   "[NTLM]%s" % case.username)

    case = cases[0]
    logging.info("Verify %s can't use anonymous access" % case)
    IISTestHelper._VerifyAnonymousAccessFails(self, case)

  @test
  def VerifyKerberosSite(self):
    cases = self._GetDomainTestCases(80)

    for case in cases:
      logging.info("Verify %s with Kerberos" % case)
      IISTestHelper._VerifyDefaultCredentialAccess(self, case,
                                                   "[SPNEGO]%s" % case.username)

    case = cases[0]
    logging.info("Verify %s can't use anonymous access" % case)
    IISTestHelper._VerifyAnonymousAccessFails(self, case)

  @test
  def VerifyKerberosNego2Site(self):
    cases = self._GetDomainTestCases(8080)

    for case in cases:
      logging.info("Verify %s with Kerberos" % case)
      IISTestHelper._VerifyDefaultCredentialAccess(self, case,
                                                   "[SPNEGO]%s" % case.username)

    case = cases[0]
    logging.info("Verify %s can't use anonymous access" % case)
    IISTestHelper._VerifyAnonymousAccessFails(self, case)

  def _GetDomainTestCases(self, port):
    """Get a list of _TestCases to verify for IISSites on 2016/2012 domains."""
    cases = []

    for version in ['2012', '2016']:
      domainPrefix = 'domain%s' % version
      for clientSuffix in ['dc', 'web', 'clt']:
        clientName = '%s-%s' % (domainPrefix, clientSuffix)
        hostname = '%s-web.%s.com' % (domainPrefix, domainPrefix)
        target = 'http://%s:%d/test.aspx' % (hostname, port)
        case = IISTestHelper._TestCase(clientName, target)

        username = "%s\\joe%s" % (domainPrefix, version)
        password = 'AAAaaa111!!!'
        case.SetCredential(username, password)

        cases.append(case)

    return cases


MAX_RETRIES_EVENTLOG = 2
DELAY_RETRY_EVENTLOG = 5  # In seconds


@category("core")
@environment(file="./assets/iis-ntlm-v1.asset.textpb")
class IISNTLMTest(EnterpriseTestCase):

  @test
  def VerifyNTLM1Site(self):
    for version in ["win2012", "win2016"]:
      case = IISTestHelper._TestCase('%s-ntlm1' % version,
                                     'http://website.test.com/test.aspx')
      case.SetCredential('joe', 'AAAaaa111!!!')
      self._VerifyNTLMSite(case, "NTLM V1")

  @test
  def VerifyNTLM2Site(self):
    for version in ["win2012", "win2016"]:
      case = IISTestHelper._TestCase('%s-ntlm2' % version,
                                     'http://website.test.com/test.aspx')
      case.SetCredential('joe', 'AAAaaa111!!!')
      self._VerifyNTLMSite(case, "NTLM V2")

  def _VerifyNTLMSite(self, case, expectedVersion):
    # Get the latest security EventLog Index. We'll use it later.
    script = "Get-EventLog Security -newest 1 | % { $_.Index }"
    output = self.clients['website'].RunPowershell(script)

    lastEventLogIndex = int(output)

    # Make sure the website is up & serves NTLM requests.
    logging.info("Verify %s with NTLM" % case)
    IISTestHelper._VerifyDefaultCredentialAccess(
        self, case, "[NTLM]test\\%s" % case.username)

    # Look at the EventLogs and assert on the NTLM version used.
    script = '''
    $logs = Get-EventLog Security -newest 500
    $logs = $logs | ? {{ $_.Index -gt {index} -and $_.EventID -eq 4624 }}
    $logs = $logs | ? {{ $_.Message.Contains("NtLmSsp") }}

    $logs | % {{ $_.Message }}
    '''.format(index=lastEventLogIndex)

    output = self.clients['website'].RunPowershell(script)

    packageUsed = "Package Name (NTLM only):\t%s" % expectedVersion

    # Event logs can take a few seconds to be available in Get-EventLog.
    i = 0
    while i < MAX_RETRIES_EVENTLOG and packageUsed not in output:
      logging.debug("No NTLM event found. Will retry in a few seconds.")
      time.sleep(DELAY_RETRY_EVENTLOG)

      ret, output = self.clients['website'].RunPowershell(script)
      self.assertEqual(ret, 0, 'Get-EventLog failed for %s.' % case)
      i += 1

    message = "Couldn't find %s in output" % packageUsed
    self.assertTrue(packageUsed in output, message)


class IISTestHelper:

  @staticmethod
  def _VerifyDefaultCredentialAccess(test, case, expectedTitle):
    """Verifies that an Invoke-WebRequest with default credential succeeds."""
    script = '''
    $password = ConvertTo-SecureString "{password}" -AsPlainText -Force
    $cred = New-Object PSCredential("{username}", $password)
    $result = Invoke-WebRequest {target} -UseBasicParsing -Credential $cred
    $result.Content
    '''.format(
        target=case.target, username=case.username, password=case.password)

    output = test.clients[case.client].RunPowershell(script)

    IISTestHelper._AssertTitleEquals(test, case, output, expectedTitle)

  @staticmethod
  def _VerifyAnonymousAccessFails(test, case):
    """Verifies that an Invoke-WebReguest with no credential fails w/ 401."""
    script = '(Invoke-WebRequest %s -UseBasicParsing).StatusCode' % case.target
    ret, output = test.clients[case.client].RunPowershellNoThrow(script)
    test.assertEqual(ret, 1)
    test.assertTrue("401 - Unauthorized" in output)

  @staticmethod
  def _AssertTitleEquals(test, case, outputHTML, expectedTitle):
    """Verifies the title of a page is what we'd expect."""
    match = re.search('\<title\>(.*)\</title\>', outputHTML)

    if not match:
      message = 'Could not find title in HTML output for %s' % case
      test.assertTrue(False, message)

    actualTitle = match.groups()[0]
    message = 'Invoke-WebRequest returned unexpected title for %s' % case
    test.assertEqual(expectedTitle.lower(), actualTitle.lower(), message)

  class _TestCase:

    def __init__(self, client, target):
      self.client = client
      self.target = target
      self.username = None
      self.password = None

    def SetCredential(self, username, password):
      self.username = username
      self.password = password

    def __repr__(self):
      return "(%s, %s)" % (self.client, self.target)
