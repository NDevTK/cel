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

configuration InstallFeatures
{
    param
    (
    )

    Import-DscResource -ModuleName xActiveDirectory

    Node localhost
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

# first check if the domain is already created
$domainFullName = "$domainName.$parentDomainName"
$domain = Get-ADDomain -Identity $domainFullName -ErrorAction Stop
if ($domain.DnsRoot -eq $domainFullName)
{
    Exit 0
}

$dnsServerAddress = [Net.DNS]::GetHostEntry($dnsServer).AddressList.IPAddressToString
netsh interface ip add dnsserver "Local Area Connection" $dnsServerAddress

InstallFeatures -ConfigurationData $ConfigData

Set-DSCLocalConfigurationManager -Path .\InstallFeatures -Verbose

# Install needed features
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\InstallFeatures -Verbose

if ($error.Count -gt $errorCount)
{
    # Exit with error code
    Write-Host "Error Occurred"
    Exit 100
}

# Now call dcpromo to config AD. See https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2008-R2-and-2008/cc732887(v=ws.10)
dcpromo /unattend /ReplicaOrNewDomain:Domain /SafeModeAdminPassword:$localAdminPassword /RebootOnCompletion:no /NewDomain:Child /InstallDNS:yes /ParentDomainDNSName:$parentDomainName /DomainNetBiosName:$netbiosName /ChildName:$domainName /Password:$adminPassword /UserName:$adminName /UserDomain:$parentDomainName

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