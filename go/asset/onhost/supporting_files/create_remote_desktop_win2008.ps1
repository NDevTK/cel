# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    [Parameter(Mandatory=$true)][String] $collectionName,
    [Parameter(Mandatory=$true)][String] $collectionDescription
)

Configuration RemoteDesktopSessionHost
{
    Import-DscResource -Module xRemoteDesktopSessionHost

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

        WindowsFeature RSAT-RDS-Tools
        {
            Ensure = "Present"
            Name = "RSAT-RDS-Tools"
            IncludeAllSubFeature = $true
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
Start-DscConfiguration -wait -force -verbose -path .\RemoteDesktopSessionHost
if ($error.Count -gt $errorCount)
{
  # Exit with error code
  Write-Host "Error Occurred"
  Exit 100
}

$m = Get-DscLocalConfigurationManager
Write-Host "LCMState : $($m.LCMState)"
if ($m.LCMState -eq "PendingReboot")
{
    # Exit with code 200 to indicate reboot is needed
    Exit 200
}

# Now all the required windows features are installed, time to configure Remote
# Desktop Host. This part needs to run using the administrator credential.

import-module RemoteDesktop
$localhost = [System.Net.Dns]::GetHostByName((hostname)).HostName
$errorCount = $error.Count

$adminPassword ="2,b)n^!q:x~VuXck"
$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
& $scriptDir\reset_local_admin_password.ps1 -newPassword $adminPassword
$cred = New-Object System.Management.Automation.PSCredential ("Administrator", (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

Invoke-Command -Credential $cred -ComputerName localhost -ScriptBlock {
    param($localhost, $collectionName, $collectionDescription)
    $params = @{
        ConnectionBroker = $localhost
        WebAccessServer = $localhost
        SessionHost = $localhost
        Verbose = $true
    }
    Write-Host "New-RDSessionDeployment. Params:`n $($params | Out-String)"
    New-RDSessionDeployment @params

    $params = @{
        CollectionName=$collectionName
        SessionHost=$localhost
        CollectionDescription = $collectionDescription
        ConnectionBroker = $localhost
        Verbose = $true
    }
    Write-Host "New-RDSessionCollection. Params:`n $($params | Out-String)"
    New-RDSessionCollection @params

    $params = @{
        Alias = "Calc"
        DisplayName="Calc"
        FilePath="C:\Windows\System32\calc.exe"
        ShowInWebAccess=$true
        CollectionName=$collectionName
        ConnectionBroker=$localhost
        Verbose=$true
    }
    Write-Host "New-RDRemoteapp. Params:`n $($params | Out-String)"
    New-RDRemoteapp @params 
} -ArgumentList $localhost,$collectionName,$collectionDescription

if ($error.Count -gt $errorCount)
{
  # Exit with error code
  Write-Host "Error Occurred during remote desktop configuration"
  Exit 100
}