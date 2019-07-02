# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    [Parameter(Mandatory=$true)] [String] $userName
)

# create link in Startup folder so that cel_ui_agent.exe is started
$objShell = New-Object -ComObject("WScript.Shell")
$startUpPath = "c:\Users\$userName\Start Menu\Programs\Startup"
$shortCut = $objShell.CreateShortcut($startUpPath + "\ui_agent.lnk")
$shortCut.TargetPath = "c:\cel\cel_ui_agent.exe"
$shortCut.Save()