# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Create a local user, and enable autologon for that user
param(
    [Parameter(Mandatory=$true)] [String] $userName,
    [Parameter(Mandatory=$true)] [String] $password
)

# Disable sppsvc service so that the "Activate Windows" poup window won't show up.
# Otherwise, the UI is locked and the user won't get logged in normally.
# In this case, set_ui_agent.ps1 will not work.
Set-Itemproperty -path 'HKLM:\SYSTEM\CurrentControlSet\services\sppsvc' -Name 'Start' -value 4

# enable autologon for $userName
net user $userName $password /add
if ($LASTEXITCODE -eq 2) {
    # the user already exists. Then change its password
    Write-Host "User already exists, changing password."
    net user $userName $password
}

# We use a file to track if autologon is run for the first time or not,
# so that we can determine if the flag "/accepteula" should be passed or not.
# Otherwise, autologon will not correctly.
$semaphore_file = 'autologon_executed'
if (Test-Path $semaphore_file) {
    # autologon has been used. So no flag "/accepteula"
    c:\ProgramData\chocolatey\bin\autologon $userName $env:computername $password
} else {
    # autologon is run for the first time, so "/accepteula" is needed
    c:\ProgramData\chocolatey\bin\autologon /accepteula $userName $env:computername $password
    New-Item $semaphore_file -Type File
}