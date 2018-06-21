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

# Download PackageManagement_x64.msi
$downloadUrl = 'https://download.microsoft.com/download/C/4/1/C41378D4-7F41-4BBE-9D0D-0E4F98585C61/PackageManagement_x64.msi'
$WC = New-Object System.Net.WebClient
$WC.DownloadFile($downloadUrl,"c:\cel\PackageManagement_x64.msi")
$WC.Dispose()

# Install PackageManagement_x64.msi.
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
Write-Host "The warning message from Set-ExecutionPolicy can be safely ignored."
Set-ExecutionPolicy RemoteSigned -Scope LocalMachine -Force

Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -force

# WinRM, which is needed by DSC, is not enabled by default on Win 2008 R2.
# Enable it.
Enable-PSRemoting -SkipNetworkProfileCheck -Force

# Install modules
install-module-if-not-installed -Name xComputerManagement
install-module-if-not-installed -Name xNetworking
install-module-if-not-installed -Name xRemoteDesktopSessionHost
install-module-if-not-installed -Name xWebAdministration