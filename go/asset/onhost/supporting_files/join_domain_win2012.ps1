# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    [Parameter(Mandatory=$true)] [String] $domainName,
    [Parameter(Mandatory=$true)] [String] $adminName,
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
        xDnsServerAddress DnsServerAddress
        {
            Address        = $dnsServerAddress
            InterfaceAlias = 'Ethernet'
            AddressFamily  = 'IPv4'
            Validate       = $false
        }

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

$errorCount = $error.Count
$dnsServerAddress = (Resolve-DNSName -Name $dnsServer -Type A).IPAddress

if ($error.Count -gt $errorCount)
{
    if ($error[-1].FullyQualifiedErrorId -match "ERROR_TIMEOUT,Microsoft.DnsClient.Commands.ResolveDnsName")
    {
        # The DomainController might be temporarily down - return 150 to try again later.
        Write-Host "Transient Error Occurred, returning 150"
        Exit 150
    }

    Write-Host "Error Occurred, returning 100"
    Exit 100
}

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