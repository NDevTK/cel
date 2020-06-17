# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# A super simple UI test using pywinauto
from pywinauto.application import Application
import pyperclip

app = Application(backend="uia").start("notepad.exe")

try:
  # Send Atl+F then X to exit notepad
  app.UntitledNotepad.type_keys("%FX")
  print("SUCCESS")
except:
  print("FAILED: cannot run UI test")

try:
  # Set text in system clipboard
  pyperclip.copy('PASTE')
  print(pyperclip.paste())
except:
  print("FAILED: cannot paste in UI test")