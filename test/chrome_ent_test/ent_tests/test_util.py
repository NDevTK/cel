# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This file contains the utility methods that can be used by python
# tests running on Windows.

import time

import win32con
import win32gui


def _window_enum_handler(hwnd, window_list):
  win_title = win32gui.GetWindowText(hwnd)
  if 'Google Chrome' in win_title:
    window_list.append(hwnd)


def _get_chrome_windows():
  """Gets the list of hwnd of Chrome windows."""
  window_list = []
  win32gui.EnumWindows(_window_enum_handler, window_list)
  return window_list


def shutdown_chrome():
  """Shutdown Chrome cleanly.

    Surprisingly there is no easy way in chromedriver to shutdown Chrome
    cleanly on Windows. So we have to use win32 API to do that: we find
    the chrome window first, then send WM_CLOSE message to it.
  """
  window_list = _get_chrome_windows()
  for win in window_list:
    win32gui.SendMessage(win, win32con.WM_CLOSE, 0, 0)

  # wait a little bit for chrome processes to end.
  # TODO: the right way is to wait until there are no chrome.exe processes.
  time.sleep(2)
