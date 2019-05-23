# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import time
from selenium import webdriver
import os

os.environ["CHROME_LOG_FILE"] = r"c:\temp\chrome_log.txt"

driver = webdriver.Chrome(
    "C:/ProgramData/chocolatey/lib/chromedriver/tools/chromedriver.exe",
    service_args=["--verbose", r"--log-path=c:\temp\chromedriver.log"])
driver.get('http://www.google.com/xhtml')
time.sleep(5)  # wait for page to be loaded
search_box = driver.find_element_by_name('q')
search_box.send_keys('searchTerm')
search_box.submit()
time.sleep(5)  # wait for the page to be loaded

print driver.current_url
driver.quit()
