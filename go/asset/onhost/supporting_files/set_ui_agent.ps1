# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Create link in Startup folder so that cel_ui_agent.exe is started
# automatically

param(
    [Parameter(Mandatory=$true)] [String] $userName
)

$objShell = New-Object -ComObject("WScript.Shell")
$startUpPath = "c:\Users\$userName\Start Menu\Programs\Startup"

# It takes some time for user log on to finish, so
# we need to wait until the startup path is created
while ($true) {
    $created = Test-Path $startUpPath
    if ($created) {
        Write-Host "startup folder is created. Create link to cel_ui_agent.exe"
        break
    } else {
        Write-Host "startup folder is not created yet. Waiting..."
        Start-Sleep -Seconds 5
    }
}

$shortCut = $objShell.CreateShortcut($startUpPath + "\ui_agent.lnk")
$shortCut.TargetPath = "c:\cel\cel_ui_agent.exe"
$shortCut.Save()