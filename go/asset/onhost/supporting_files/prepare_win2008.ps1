# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Install module if it is not installed yet
function install-module-if-not-installed {
    param($name)

    if (Get-Module -ListAvailable -Name $name) {
      Write-Host "Module $name exists"
    }
    else {
        Install-Module -Name $name -Force -Verbose
    }
  }

# install PackageManagement_x64.msi if we're running on Win2008R2.
$arguments = @(
    "/i"
    "c:\cel\PackageManagement_x64.msi"
    "/qn"
    "/norestart"
    "/l*v"
    "c:\cel\msi_log.txt"
)
$process = Start-Process -FilePath msiexec.exe -ArgumentList $arguments -Wait -PassThru
if ($process.ExitCode -eq 0) {
    Write-Host "msiFile has been successfully installed"
} else {
    Write-Error "msiFile installation faield"
}

# Somehow this is needed on Win2008. There will be a warning message
# that can be safely ignored.
Set-ExecutionPolicy RemoteSigned -Scope LocalMachine -Force

Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -force

# install the module that contains Expand-Archive cmdlet
install-module-if-not-installed -Name Microsoft.PowerShell.Archive