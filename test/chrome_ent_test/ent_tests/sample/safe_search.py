# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import os
import logging
from chrome_ent_test.infra.core import *
from chrome_ent_test.ent_tests import ChromeEnterpriseTestCase
from absl import flags

FLAGS = flags.FLAGS


@environment(file="./safe_search.asset.textpb")
class SafeSearchTest(ChromeEnterpriseTestCase):

  @test
  def Test(self):
    self.InstallChrome('client2012')
    self.InstallWebDriver('client2012')

    # disable policy ForceGoogleSafeSearch
    self.SetPolicy('win2012-dc', 'ForceGoogleSafeSearch', 0, 'DWORD')
    self.RunCommand('client2012', 'gpupdate /force')
    dir = os.path.dirname(os.path.abspath(__file__))
    logging.info('ForceGoogleSafeSearch DISABLED')
    output = self.RunWebDriverTest(
        'client2012', os.path.join(dir, 'safe_search_webdriver_test.py'))
    logging.info('url used: %s', output)

    # assert that safe search is NOT enabled
    self.assertTrue('safe=active' not in output)
    self.assertTrue('ssui=on' not in output)

    # enable policy ForceGoogleSafeSearch
    self.SetPolicy('win2012-dc', 'ForceGoogleSafeSearch', 1, 'DWORD')
    self.RunCommand('client2012', 'gpupdate')
    logging.info('ForceGoogleSafeSearch ENABLED')
    output = self.RunWebDriverTest(
        'client2012', os.path.join(dir, 'safe_search_webdriver_test.py'))
    logging.info('url used: %s', output)

    # assert that safe search is enabled
    self.assertTrue('safe=active' in output)
    self.assertTrue('ssui=on' in output)
