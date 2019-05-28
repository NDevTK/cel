# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import time
import test_util
import json
from absl import app, flags
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

FLAGS = flags.FLAGS

flags.DEFINE_enum(
    'action', None, ['open_urls', 'start_chrome'], """The action to take.

    - open_urls: start chrome, then open urls passed thru --urls in tabs.
    - start_chrome: start chrome.
    """)

flags.DEFINE_multi_string('urls', None, "List of urls to open")
flags.DEFINE_string('user_data_dir', None,
                    "The user data directory used by chrome")


def _create_driver():
  chrome_options = Options()
  chrome_options.add_argument(r'user-data-dir=%s' % FLAGS.user_data_dir)
  driver = webdriver.Chrome(
      executable_path=
      "C:/ProgramData/chocolatey/lib/chromedriver/tools/chromedriver.exe",
      chrome_options=chrome_options,
      service_args=["--verbose", r"--log-path=c:\temp\chromedriver.log"])
  return driver


def _get_urls(driver):
  """Returns the list of URLs in tabs."""
  urls = []
  for w in driver.window_handles:
    driver.switch_to.window(w)
    urls.append(driver.current_url)
  list.sort(urls)
  return urls


def open_urls():
  driver = _create_driver()

  # open the first url in the current New Tab tab
  driver.get(FLAGS.urls[0])

  # open the rest of urls in new tabs
  for url in FLAGS.urls[1:]:
    driver.execute_script("window.open('%s');" % url)

  # give chrome some time to load everything
  time.sleep(2)

  print json.dumps(_get_urls(driver))
  test_util.shutdown_chrome()


def start_chrome():
  """Start chrome.

  Write the list of URLs in tabs to stdout.
  """
  driver = _create_driver()

  # give chrome some time to load everything. This is less than ideal, but
  # currently there's no statisfactory solution.
  time.sleep(10)

  print json.dumps(_get_urls(driver))
  test_util.shutdown_chrome()


def main(argv):
  if FLAGS.action == 'open_urls':
    open_urls()
  elif FLAGS.action == 'start_chrome':
    start_chrome()


if __name__ == '__main__':
  app.run(main)