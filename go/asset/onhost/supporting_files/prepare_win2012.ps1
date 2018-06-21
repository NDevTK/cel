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

# Install modules
Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -force -Verbose
install-module-if-not-installed -Name xActiveDirectory
install-module-if-not-installed -Name xComputerManagement
install-module-if-not-installed -Name xNetworking
install-module-if-not-installed -Name xRemoteDesktopSessionHost
install-module-if-not-installed -Name xWebAdministration