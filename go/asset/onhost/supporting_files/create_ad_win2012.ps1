# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    [Parameter(Mandatory=$true)] [String] $domainName,
    [Parameter(Mandatory=$true)] [String] $netbiosName,
    [Parameter(Mandatory=$true)] [String] $adminName,
    [Parameter(Mandatory=$true)] [String] $adminPassword
)

# AD cannot be created if local admin password is empty.
# Set the password here.
$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
& $scriptDir\reset_local_admin_password.ps1 -newPassword $adminPassword

Configuration NewDomain
{
   param
    (
        [Parameter(Mandatory)]
        [pscredential]$safemodeAdministratorCred,
        [Parameter(Mandatory)]
        [pscredential]$domainCred
    )

    Import-DscResource -ModuleName xActiveDirectory

    Node localhost
    {
        LocalConfigurationManager
        {
            ActionAfterReboot = 'ContinueConfiguration'
            ConfigurationMode = 'ApplyOnly'
            RebootNodeIfNeeded = $false
        }

        WindowsFeature ADDSInstall
        {
            Ensure = "Present"
            Name = "AD-Domain-Services"
        }

        # Optional GUI tools
        WindowsFeature ADDSTools
        {
            Ensure = "Present"
            Name = "RSAT-ADDS"
        }

        xADDomain FirstDS
        {
            DomainName = $domainName
            DomainNetBiosName = $netbiosName
            DomainAdministratorCredential = $domainCred
            SafemodeAdministratorPassword = $safemodeAdministratorCred
            DependsOn = "[WindowsFeature]ADDSInstall"
        }
    }
}

# Configuration Data for AD
$ConfigData = @{
    AllNodes = @(
        @{
            Nodename = "localhost"
            PsDscAllowPlainTextPassword = $true
        }
    )
}

$safeModeAdminCred = New-Object System.Management.Automation.PSCredential ("(Password Only)", (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

# Credentials used to query for domain existence.
# Note: These are NOT used during domain creation.
$domainCred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

NewDomain -ConfigurationData $ConfigData `
    -safemodeAdministratorCred $safeModeAdminCred `
    -domainCred $domainCred

# Make sure that LCM is set to continue configuration after reboot
Set-DSCLocalConfigurationManager -Path .\NewDomain -Verbose

# Build the domain
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\NewDomain -Verbose

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