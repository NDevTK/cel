# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    [Parameter(Mandatory=$true)][String] $adminName,
    [Parameter(Mandatory=$true)][String] $adminPassword,

    # ExtraArgs catches extra RDS arguments that aren't used in this script.
    # e.g. RDS Collection name/description (unsupported in Windows Server 2008)
    [Parameter(ValueFromRemainingArguments = $true)]$extraArgs
)

Configuration RemoteDesktopSessionHost
{
    Node "localhost"
    {
        LocalConfigurationManager
        {
            RebootNodeIfNeeded = $false
        }

        WindowsFeature RDS-RD-Server
        {
            Ensure = "Present"
            Name = "RDS-RD-Server"
        }

        WindowsFeature Desktop-Experience
        {
            Ensure = "Present"
            Name = "Desktop-Experience"
        }

        WindowsFeature RDS-Connection-Broker
        {
            Ensure = "Present"
            Name = "RDS-Connection-Broker"
        }

        WindowsFeature RDS-Web-Access
        {
            Ensure = "Present"
            Name = "RDS-Web-Access"
        }

        WindowsFeature RDS-Licensing
        {
            Ensure = "Present"
            Name = "RDS-Licensing"
        }
    }
}

RemoteDesktopSessionHost

Set-DSCLocalConfigurationManager -Path .\RemoteDesktopSessionHost -Verbose

$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\RemoteDesktopSessionHost -Verbose

$pendingReboot = Test-Path "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Component Based Servicing\RebootPending"
Write-Host "PendingReboot : $pendingReboot"
if ($pendingReboot)
{
    # Exit with code 200 to indicate reboot is needed
    Exit 200
}

if ($error.Count -gt $errorCount)
{
    # Exit with error code
    Write-Host "Error Occurred"
    Exit 100
}

import-module RemoteDesktopServices

$RemoteAppProgramsPath = "RDS:\RemoteApp\RemoteAppPrograms"

# Check if we've already created the RemoteApp program
if (Test-Path $RemoteAppProgramsPath\Calc)
{
    Write-Host "Calc is already registered as a RemoteApp program."
    Exit 0
}

$cred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

# Allow all domain users to access the remote apps
[Microsoft.TerminalServices.PSEngine.UserGroupHelper]::AddMember("Remote Desktop Users", "Domain Users")
Set-Item RDS:\RDSConfiguration\SessionSettings\AllowConnections 1 -Credential $cred
Set-Item RDS:\RDSConfiguration\Connections\RDP-Tcp\SecuritySettings\UserAuthenticationRequired 1 -Credential $cred

# Add the Calc.exe app to the list of available remote apps
$params = @{
    Name = "Calc"
    ApplicationPath = "C:\Windows\System32\calc.exe"
    ShowInWebAccess = 1
    Credential = $cred
    Verbose = $true
}
Write-Host "New-Item in RemoteAppPrograms. Params:`n $($params | Out-String)"
New-Item $RemoteAppProgramsPath @params

if ($error.Count -gt $errorCount)
{
    # Exit with error code
    Write-Host "Error Occurred during remote desktop configuration"
    Exit 100
}