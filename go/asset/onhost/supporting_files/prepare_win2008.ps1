# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Install module if it is not installed yet
function install-module-if-not-installed {
    param($name, $requiredVersion)

    if (Get-Module -ListAvailable -Name $name) {
      Write-Host "Module $name exists"
    }
    else {
        Install-Module -Name $name -Force -Verbose -RequiredVersion $requiredVersion
    }
}

# Download PackageManagement_x64.msi
$downloadUrl = 'https://download.microsoft.com/download/C/4/1/C41378D4-7F41-4BBE-9D0D-0E4F98585C61/PackageManagement_x64.msi'
$WC = New-Object System.Net.WebClient
$WC.DownloadFile($downloadUrl, "c:\cel\PackageManagement_x64.msi")
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
install-module-if-not-installed -Name xActiveDirectory -RequiredVersion 2.19.0.0
install-module-if-not-installed -Name xComputerManagement -RequiredVersion 4.1.0.0
install-module-if-not-installed -Name xNetworking -RequiredVersion 5.7.0.0
install-module-if-not-installed -Name xRemoteDesktopSessionHost -RequiredVersion 1.6.0.0
install-module-if-not-installed -Name xWebAdministration -RequiredVersion 2.2.0.0

# WinServer2008 has scheduled tasks for consistency that can't be disabled (removed in WMF 5.0).
# Calling `Start-DscConfiguration` during those tasks fails and can't be bypassed with `-Force`.
# This will remove Pending DSC configs. We cycle through all onhost scripts after restarts so
# incomplete configurations will continue when we get back to it.
function Get-ConfigurationFilesCsv {
  $items = Get-ChildItem "C:\Windows\System32\Configuration\" | % { $_.FullName }
  return $items -join ", "
}

Write-Host "Configuration files before: " + (Get-ConfigurationFilesCsv)

Remove-Item C:\Windows\System32\Configuration\*.mof
Get-Process *WMI* | ? { $_.modules.ModuleName -like "*DSC*" } | Stop-Process -Force

Write-Host "Configuration files after: " + (Get-ConfigurationFilesCsv)
