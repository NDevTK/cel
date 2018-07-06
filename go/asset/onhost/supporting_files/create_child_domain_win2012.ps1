# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    # name of the child domain. It is not the FQDN.
    [Parameter(Mandatory=$true)] [String] $domainName,

    # netbios name of the child domain
    [Parameter(Mandatory=$true)] [String] $netbiosName,

    # FQDN of the parent domain
    [Parameter(Mandatory=$true)] [String] $parentDomainName,

    # The name of the administrator, in the format domain\user
    [Parameter(Mandatory=$true)] [String] $adminName,

    # The password of the administrator
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    # The password of the administrator in this new domain
    [Parameter(Mandatory=$true)] [String] $localAdminPassword,

    # the dns server
    [Parameter(Mandatory=$true)] [String] $dnsServer
  )  

$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
& $scriptDir\reset_local_admin_password.ps1 -newPassword $localAdminPassword

configuration NewDomain
{
   param
    (
        [Parameter(Mandatory)]
        [pscredential]$safemodeAdministratorCred,
        [Parameter(Mandatory)]
        [pscredential]$domainCred
    )

    Import-DscResource -ModuleName xActiveDirectory
    Import-DscResource -Module xNetworking

    Node localhost
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

        xADDomain FirstDS
        {
            DomainName = $domainName
            DomainNetBiosName = $netbiosName
            ParentDomainName = $parentDomainName
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

$dnsServerAddress = (Resolve-DNSName -Name $dnsServer -Type A).IPAddress

$safeModeAdmiCred = New-Object System.Management.Automation.PSCredential ("(Password Only)", (ConvertTo-SecureString $localAdminPassword -AsPlainText -Force))

$domainCred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

NewDomain -ConfigurationData $ConfigData `
    -safemodeAdministratorCred $safeModeAdmiCred `
    -domainCred $domainCred

Set-DSCLocalConfigurationManager -Path .\NewDomain -Verbose

# Build the domain
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\NewDomain -Verbose
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