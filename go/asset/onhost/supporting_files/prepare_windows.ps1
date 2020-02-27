# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

function install-package-management {
  param($requiredVersion)

  if ($PSVersionTable.PSVersion.Major -ge 5) {
    # Powershell version 5 and above have the Package Management cmdlets built
    # in. Skip installation.
    return
  }

  # Download the installation file from CIPD.
  # Set the ensure file encoding to ANSI. Cipd cannot interpret an ensure
  # file in Powershell's default unicode encoding.
  Write-Host "$(Get-Date): Downloading PackageManagement.msi from CIPD."
  $ensureFile = "c:\cel\supporting_files\ensure_file.txt"
  echo "infra/celab/third_party/powershell_modules/package_management $requiredVersion" | Out-File -Encoding Ascii $ensureFile
  cipd.bat ensure -ensure-file $ensureFile -root c:\cel

  # Install PackageManagement.msi.
  $arguments = @(
      "/i"
      "c:\cel\PackageManagement.msi"
      "/qn"
      "/norestart"
      "/l*v"
      "c:\cel\msi_log.txt"
  )
  $process = Start-Process -FilePath msiexec.exe -ArgumentList $arguments -Wait -PassThru
  if ($process.ExitCode -eq 0) {
      Write-Host "msiFile installation succeeded."
  } else {
      Write-Error "msiFile installation failed."
  }
}

function install-nuget-package-provider-if-not-installed {
  param($requiredVersion)

  if (Get-PackageProvider -ListAvailable -Name NuGet) {
    Write-Host "$(Get-Date): NuGet already exists"
  }
  else {
    Write-Host "$(Get-Date): Downloading NuGet from CIPD"
    $ensureFile = "c:\cel\supporting_files\ensure_file.txt"
    # Set the ensure file encoding to ANSI. Cipd cannot interpret an ensure
    # file in Powershell's default unicode encoding.
    echo "infra/celab/third_party/powershell_package_providers/nuget $requiredVersion" | Out-File -Encoding Ascii $ensureFile
    cipd.bat ensure -ensure-file $ensureFile -root c:\cel\supporting_files

    # Copy the nuget directory to the package provider directory.
    Copy-Item -Path "c:\cel\supporting_files\nuget" -Destination $env:ProgramFiles\PackageManagement\ProviderAssemblies -Recurse

    # Import the NuGet package provider.
    Import-PackageProvider -Name NuGet -RequiredVersion $requiredVersion
  }
}

function install-module-if-not-installed {
  param($name, $requiredVersion)

  if (Get-Module -ListAvailable -Name $name) {
    Write-Host "$(Get-Date): Module $name exists"
  }
  else {
    Write-Host "$(Get-Date): Downloading module $name from CIPD"
    $ensureFile = "c:\cel\supporting_files\ensure_file.txt"
    # Set the ensure file encoding to ANSI. Cipd cannot interpret an ensure
    # file in Powershell's default unicode encoding.
    echo "infra/celab/third_party/powershell_modules/$($name.ToLower()) $requiredVersion" | Out-File -Encoding Ascii $ensureFile
    cipd.bat ensure -ensure-file $ensureFile -root c:\cel\supporting_files

    Write-Host "$(Get-Date): Installing module $Name"
    Install-Module -Name $name -Force -Verbose -Repository "local_repo"
  }
}

# Install Chromium's depot tools bundle for windows.
# Depot tools contains the Chrome Infra Package Deployer binary (CIPD),
# which celab uses to access third-party binaries for setting up test
# client machines.
# For more details, please see
# https://commondatastorage.googleapis.com/chrome-infra-docs/flat/depot_tools/docs/html/depot_tools_tutorial.html#_setting_up
function install-depot-tools {
  $depotToolsDir = Join-Path $env:SYSTEMDRIVE "depot_tools"
  # Skip installation if the machine already has depot tools.
  if ([System.IO.Directory]::Exists($depotToolsDir) -and
      $($env:Path).ToLower().Contains($($depotToolsDir).ToLower())) {
    return
  }

  # Create a local directory to host depot tools
  if (![System.IO.Directory]::Exists($depotToolsDir)) {[System.IO.Directory]::CreateDirectory($depotToolsDir)}
  $depotToolsZip = Join-Path $depotToolsDir "depot_tools.zip"

  # Download Chromium's depot tools bundle for windows.
  $depotToolsDownloadUrl = 'https://storage.googleapis.com/chrome-infra/depot_tools.zip'
  $webClient = New-Object System.Net.WebClient
  $webClient.DownloadFile($depotToolsDownloadUrl, $depotToolsZip)
  $webClient.Dispose()

  # unzip depot_tools.zip.
  if ($PSVersionTable.PSVersion.Major -lt 5) {
    # Older version of Powershell do not have the Expand-Archive cmdlet.
    try {
      $shellApplication = new-object -com shell.application
      $zipPackage = $shellApplication.NameSpace($depotToolsZip)
      $destinationFolder = $shellApplication.NameSpace($depotToolsDir)
      $destinationFolder.CopyHere($zipPackage.Items(), 0x10)
    } catch {
      throw "Unable to unzip package using built-in compression. Error: `n $_"
    }
  } else {
    Expand-Archive -Path "$depotToolsZip" -DestinationPath "$depotToolsDir" -Force
  }

  # Include the depot tools directory in the $PATH environment variable.
  if ($($env:Path).ToLower().Contains($($depotToolsDir).ToLower()) -eq $false) {
    $env:Path += ";$depotToolsDir"
    [System.Environment]::SetEnvironmentVariable(
        'Path', $env:Path, [System.EnvironmentVariableTarget]::Machine)
  }
}

install-depot-tools
install-package-management -RequiredVersion latest

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
#Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -force -Verbose
install-nuget-package-provider-if-not-installed -RequiredVersion 2.8.5.208
Register-PSRepository -Name "local_repo" -SourceLocation c:\cel\supporting_files
install-module-if-not-installed -Name xActiveDirectory -RequiredVersion 2.19.0
install-module-if-not-installed -Name xComputerManagement -RequiredVersion 4.1.0
install-module-if-not-installed -Name xNetworking -RequiredVersion 5.7.0
install-module-if-not-installed -Name xRemoteDesktopSessionHost -RequiredVersion 1.8.0
install-module-if-not-installed -Name xWebAdministration -RequiredVersion 2.2.0
Unregister-PSRepository -Name "local_repo"