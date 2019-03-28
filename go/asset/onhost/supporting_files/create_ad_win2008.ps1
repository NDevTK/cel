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

Configuration InstallFeatures
{
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
Write-Host "Get-ADDomain args: Identity=$domainName"
$domain = Get-ADDomain -Identity $domainName -ErrorAction Stop
if ($domain.DnsRoot -eq $domainName)
{
    Exit 0
}

# Install needed features
InstallFeatures
Set-DSCLocalConfigurationManager -Path .\InstallFeatures -Verbose

$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\InstallFeatures -Verbose

$pendingReboot = Test-Path "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Component Based Servicing\RebootPending"
Write-Host "PendingReboot : $pendingReboot"
if ($pendingReboot)
{
    # Exit with code 200 to indicate reboot is needed
    Exit 200
}

if ($error.Count -gt $errorCount)
{
    $errorCode = 100

    foreach ($err in $error[$errorCount..($error.Count-1)])
    {
        Write-Host "FullyQualifiedErrorId: $($err.FullyQualifiedErrorId)"
        Format-List -InputObject $err
    }

    # Exit with error code
    Write-Host "Error Occurred, returning $errorCode"
    Exit $errorCode
}

# Now call dcpromo to config AD
Write-Host "dcpromo args: SafeModeAdminPassword=$adminPassword, DomainNetBiosName=$netbiosName, NewDomainDNSName=$domainName, Password=$adminPassword, UserName=$adminName"
dcpromo /unattend /ReplicaOrNewDomain:Domain /SafeModeAdminPassword:$adminPassword /RebootOnCompletion:no /NewDomain:Forest /InstallDNS:yes /DomainNetBiosName:$netbiosName /NewDomainDNSName:$domainName /Password:$adminPassword /UserName:$adminName

Write-Host "Last exit code: $LastExitCode"

# See https://docs.microsoft.com/en-us/windows-server/identity/ad-ds/deploy/troubleshooting-domain-controller-deployment for the list of exit codes

# exit code [1, 4] means sucess
if (($LastExitCode -ge 1) -and ($LastExitCode -le 4))
{
    Write-Host "Reboot needed"
    Exit 200
}

# exit code 77: The specified argument 'DomainNetBiosName' was not recognized.
# exit code 98: Component binaries are being installed or uninstalled. Try again later.
if (($LastExitCode -eq 77) -or ($LastExitCode -eq 98))
{
    # Exit with code 150 to indicate this is a retryable error
    Exit 150
}

Write-Host "Error occurred"
Exit 100
