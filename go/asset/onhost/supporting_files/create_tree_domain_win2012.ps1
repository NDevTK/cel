# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    [Parameter(Mandatory=$true)] [String] $domainName,
    [Parameter(Mandatory=$true)] [String] $netbiosName,
    [Parameter(Mandatory=$true)] [String] $rootDomainName,
    [Parameter(Mandatory=$true)] [String] $dnsServer,
    [Parameter(Mandatory=$true)] [String] $adminName,
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    # The password of the administrator in this new domain
    [Parameter(Mandatory=$true)] [String] $localAdminPassword
  )

# DSC doesn't support tree domain, so we just use DSC to install
# AD services. Then we call powershell cmd to install the tree domain.

# AD cannot be created if local admin password is empty.
# Set the password here.
$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
& $scriptDir\reset_local_admin_password.ps1 -newPassword $localAdminPassword

configuration NewDomain
{
    Import-DscResource -ModuleName xActiveDirectory
    Import-DscResource -Module xNetworking

    Node $AllNodes.Where{$_.Role -eq "dc"}.Nodename
    {
        xDnsServerAddress DnsServerAddress
        {
            Address        = $dnsServerAddress
            InterfaceAlias = 'Ethernet'
            AddressFamily  = 'IPv4'
            Validate       = $false
        }

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
    }
}

# Configuration Data for AD
$ConfigData = @{
    AllNodes = @(
        @{
            Nodename = "localhost"
            Role = "dc"
            DomainName = $domainName
            RetryCount = 3
            RetryIntervalSec = 30
            PsDscAllowPlainTextPassword = $true
        }
    )
}

& $scriptDir\reset_local_admin_password.ps1 -newPassword $localAdminPassword

$safeModeAdmiCred = New-Object System.Management.Automation.PSCredential ("(Password Only)", (ConvertTo-SecureString $localAdminPassword -AsPlainText -Force))

$domainCred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

$dnsServerAddress = (Resolve-DNSName -Name $dnsServer -Type A).IPAddress

NewDomain -ConfigurationData $ConfigData

# Make sure that LCM is set to continue configuration after reboot
Set-DSCLocalConfigurationManager -Path .\NewDomain -Verbose

# Install domain services
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\NewDomain -Verbose

Import-Module ADDSDeployment
Install-AddsDomain -DomainType TreeDomain -ParentDomainName $rootDomainName `
    -NewDomainNetBiosName $netbiosName `
    -NewDomainName $domainName -Credential $domainCred `
    -SafeModeAdministratorPassword $safeModeAdmiCred.Password -Force `
    -Verbose

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