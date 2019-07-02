# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Create a local user, and enable autologon for that user
param(
    [Parameter(Mandatory=$true)] [String] $userName,
    [Parameter(Mandatory=$true)] [String] $password
)

# enable autologon
net user $userName $password /add
if ($LASTEXITCODE -eq 2) {
    # the user already exists. Then change its password
    Write-Host "User already exists, changing password."
    net user $userName $password
}

c:\ProgramData\chocolatey\bin\autologon /accepteula $userName $env:computername $password