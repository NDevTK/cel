# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    # This is FQDN
    [Parameter(Mandatory=$true)] [String] $domainName,

    [Parameter(Mandatory=$true)] [String] $netbiosName,

    # FQDN of the root domain
    [Parameter(Mandatory=$true)] [String] $rootDomainName,

    [Parameter(Mandatory=$true)] [String] $dnsServer,
    [Parameter(Mandatory=$true)] [String] $adminName,
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    # The password of the administrator in this new domain
    [Parameter(Mandatory=$true)] [String] $localAdminPassword
  )

# DSC doesn't support tree domain, so we just use DSC to install
# AD services.


# AD cannot be created if local admin password is empty.
# Set the password here.
$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
& $scriptDir\reset_local_admin_password.ps1 -newPassword $adminPassword

configuration NewDomain
{
    Import-DscResource -ModuleName xActiveDirectory

    Node $AllNodes.Where{$_.Role -eq "dc"}.Nodename
    {
        LocalConfigurationManager
        {
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

# This is needed on Win 2008 R2
Enable-PSRemoting -SkipNetworkProfileCheck -Force

# first check if the domain is already created
$domain = Get-ADDomain -Identity $domainName -ErrorAction Stop
if ($domain.DnsRoot -eq $domainName)
{
    Exit 0
}

& $scriptDir\reset_local_admin_password.ps1 -newPassword $localAdminPassword

$safeModeAdmiCred = New-Object System.Management.Automation.PSCredential ("(Password Only)", (ConvertTo-SecureString $localAdminPassword -AsPlainText -Force))

$domainCred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

$dnsServerAddress = [Net.DNS]::GetHostEntry($dnsServer).AddressList.IPAddressToString
netsh interface ip add dnsserver "Local Area Connection" $dnsServerAddress

NewDomain -ConfigurationData $ConfigData

# Make sure that LCM is set to continue configuration after reboot
Set-DSCLocalConfigurationManager -Path .\NewDomain -Verbose

# Install domain services
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\NewDomain -Verbose

# Now call dcpromo to config AD. See https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2008-R2-and-2008/cc732887(v=ws.10)
dcpromo /unattend /ReplicaOrNewDomain:Domain /SafeModeAdminPassword:$localAdminPassword /RebootOnCompletion:no /NewDomain:Tree /InstallDNS:yes /ParentDomainDNSName:$parentDomainName /DomainNetBiosName:$netbiosName /NewDomainDNSName:$domainName /Password:$adminPassword /UserName:$adminName /UserDomain:$rootDomainName

Write-Host "Last exit code: $LastExitCode"

# See https://docs.microsoft.com/en-us/windows-server/identity/ad-ds/deploy/troubleshooting-domain-controller-deployment for the list of exit codes

# exit code [1, 4] means sucess
if (($LastExitCode -ge 1) -and ($LastExitCode -le 4))
{
    Write-Host "Reboot needed"
    Exit 200
}
else
{
    Write-Host "Error occurred"
    Exit 100
}