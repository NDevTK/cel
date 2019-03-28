# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    [Parameter(Mandatory=$true)][String] $adminName,
    [Parameter(Mandatory=$true)][String] $adminPassword,

    [Parameter(Mandatory=$true)][String] $collectionName,
    [Parameter(Mandatory=$true)][String] $collectionDescription
)

Configuration RemoteDesktopSessionHost
{
    param
    (
        [Parameter(Mandatory = $true)]
        [ValidateNotNullorEmpty()]
        [System.Management.Automation.PSCredential]
        $credential,

        [Parameter(Mandatory = $true)][String] $localhost,
        [Parameter(Mandatory = $true)][String] $collectionName,
        [Parameter(Mandatory = $true)][String] $collectionDescription
    )

    Import-DscResource -Module xRemoteDesktopSessionHost

    Node localhost
    {
        LocalConfigurationManager
        {
            RebootNodeIfNeeded = $false
        }

        WindowsFeature Remote-Desktop-Services
        {
            Ensure = "Present"
            Name = "Remote-Desktop-Services"
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

        xRDSessionDeployment Deployment
        {
            ConnectionBroker = $localhost
            WebAccessServer = $localhost
            SessionHost = $localhost
            PsDscRunAsCredential = $credential
            DependsOn = "[WindowsFeature]Remote-Desktop-Services", "[WindowsFeature]RDS-RD-Server", "[WindowsFeature]RDS-Connection-Broker"
        }

        xRDSessionCollection Collection
        {
            CollectionName = $collectionName
            CollectionDescription = $collectionDescription
            ConnectionBroker = $localhost
            SessionHost = $localhost
            PsDscRunAsCredential = $credential
            DependsOn = "[xRDSessionDeployment]Deployment"
        }

        xRDRemoteApp Calc
        {
            Alias = "Calc"
            DisplayName = "Calc"
            FilePath = "C:\Windows\System32\calc.exe"
            ShowInWebAccess = $true
            CollectionName = $collectionName
            PsDscRunAsCredential = $credential
            DependsOn = "[xRDSessionCollection]Collection"
        }
    }
}

$ConfigData = @{
    AllNodes = @(
        @{
            Nodename = "localhost"
            PsDscAllowPlainTextPassword = $true
            PSDscAllowDomainUser = $true
        }
    )
}

$localhost = [System.Net.Dns]::GetHostByName((hostname)).HostName

$cred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

RemoteDesktopSessionHost -ConfigurationData $ConfigData -Credential $cred -localhost $localhost -CollectionName $collectionName -CollectionDescription $collectionDescription

Set-DSCLocalConfigurationManager -Path .\RemoteDesktopSessionHost -Verbose

$errorCount = $error.Count
Start-DscConfiguration -wait -force -verbose -path .\RemoteDesktopSessionHost

$m = Get-DscLocalConfigurationManager
if ($error.Count -gt $errorCount)
{
    $errorCode = 100

    foreach ($err in $error[$errorCount..($error.Count-1)])
    {
        Write-Host "FullyQualifiedErrorId: $($err.FullyQualifiedErrorId)"
        Format-List -InputObject $err
    }

    # Exit with error code
    Write-Host "Error Occurred, returning $errorCode, LCMState : $($m.LCMState)"
    Exit $errorCode
}

Write-Host "LCMState : $($m.LCMState)"
if ($m.LCMState -eq "PendingReboot")
{
    # Exit with code 200 to indicate reboot is needed
    Exit 200
}
