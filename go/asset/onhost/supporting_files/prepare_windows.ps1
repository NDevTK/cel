# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Install module if it is not installed yet
function install-module-if-not-installed {
  param($name, $requiredVersion)

  if (Get-Module -ListAvailable -Name $name) {
    Write-Host "$(Get-Date): Module $name exists"
  }
  else {
    Write-Host "$(Get-Date): Installing module $Name"
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

# Enable WinRM, needed for DSC on all nested VMs and WinServer2008.
$retries = 0
while ($true) {
  try {
    Write-Host "$(Get-Date): Enable WinRM"
    Enable-PSRemoting -SkipNetworkProfileCheck -Force
    break
  }
  catch [System.InvalidOperationException]
  {
    Write-Host "$(Get-Date): Exception during Enable-PSRemoting:"
    Write-Host $_.Exception.Message

    $retries++
    if ($retries -gt 3) {
      Write-Host "$(Get-Date): Have already retried enough times. Aborting."
      throw $_.Exception
    }

    Write-Host "$(Get-Date): Will sleep a bit and try again."
    Start-Sleep 30
  }
}

# This statement is needed on nested VMs and WinServer2008.
Write-Host "$(Get-Date): The warning message from Set-ExecutionPolicy can be safely ignored."
Set-ExecutionPolicy RemoteSigned -Scope LocalMachine -Force

# Install modules
Register-PSRepository -Default
Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -force -Verbose
install-module-if-not-installed -Name xActiveDirectory -RequiredVersion 2.19.0.0
install-module-if-not-installed -Name xComputerManagement -RequiredVersion 4.1.0.0
install-module-if-not-installed -Name xNetworking -RequiredVersion 5.7.0.0
install-module-if-not-installed -Name xRemoteDesktopSessionHost -RequiredVersion 1.8.0.0
install-module-if-not-installed -Name xWebAdministration -RequiredVersion 2.2.0.0