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

# Enable WinRM, which is needed by DSC. This is not needed on GCE images,
# but for nested VM, we need this.
Enable-PSRemoting -SkipNetworkProfileCheck -Force

# Again, this statement is not needed on GCE images, but we need it for nested VM.
Write-Host "The warning message from Set-ExecutionPolicy can be safely ignored."
Set-ExecutionPolicy RemoteSigned -Scope LocalMachine -Force

# Install modules
Register-PSRepository -Default
Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -force -Verbose
install-module-if-not-installed -Name xActiveDirectory -RequiredVersion 2.19.0.0
install-module-if-not-installed -Name xComputerManagement -RequiredVersion 4.1.0.0
install-module-if-not-installed -Name xNetworking -RequiredVersion 5.7.0.0
install-module-if-not-installed -Name xRemoteDesktopSessionHost -RequiredVersion 1.8.0.0
install-module-if-not-installed -Name xWebAdministration -RequiredVersion 2.2.0.0