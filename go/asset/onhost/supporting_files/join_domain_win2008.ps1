# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Create a user in a domain.
param(
    # the FQDN name of the domain.    
    [Parameter(Mandatory=$true)] [String] $domainName,

    # the password of the domain administrator.    
    [Parameter(Mandatory=$true)] [String] $adminName,

    # the name of the user to be created.    
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    # the name of the dns server
    [Parameter(Mandatory=$true)] [String] $dnsServer
  )  
  
Configuration JoinDomain
{
    param
    (
        [Parameter(Mandatory = $true)]
        [ValidateNotNullorEmpty()]
        [System.Management.Automation.PSCredential]
        $credential
    )

    Import-DscResource -Module xComputerManagement
    Import-DscResource -Module xNetworking

    Node localhost
    {
        xComputer JoinDomain        
        {
            Name       = $env:COMPUTERNAME
            DomainName = $domainName
            Credential = $credential # Credential to join to domain
        }
    }
}

# Configuration Data for AD
$ConfigData = @{
    AllNodes = @(
        @{
            Nodename = "localhost"
            PsDscAllowPlainTextPassword = $true
            PSDscAllowDomainUser = $true
        }
    )
}

$domainCred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

$dnsServerAddress = [Net.DNS]::GetHostEntry($dnsServer).AddressList.IPAddressToString
netsh interface ip add dnsserver "Local Area Connection" $dnsServerAddress

JoinDomain -ConfigurationData $ConfigData -credential $domainCred

# Join the domain
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\JoinDomain -Verbose

if ($error.Count -gt $errorCount)
{
    $errorCode = 100

    foreach ($err in $error[$errorCount..($error.Count-1)])
    {
        # Look for retryable errors
        if ($err.FullyQualifiedErrorId -match "FailToJoinDomainFromWorkgroup")
        {
                $errorCode = 150
        }
    }

    # Exit with error code
    Write-Host "Error Occurred, returning $errorCode"
    Exit $errorCode
}

$m = Get-DscLocalConfigurationManager
Write-Host "LCMState : $($m.LCMState)"
if ($m.LCMState -eq "PendingReboot")
{
    # Exit with code 200 to indicate reboot is needed
    Exit 200 
}
