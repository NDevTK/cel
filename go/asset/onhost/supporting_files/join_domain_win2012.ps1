# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param(
    [Parameter(Mandatory=$true)] [String] $domainName,
    [Parameter(Mandatory=$true)] [String] $adminName,
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    # the address the dns server
    [Parameter(Mandatory=$true)] [String] $dnsServerAddress
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

JoinDomain -ConfigurationData $ConfigData -credential $domainCred

# Join the domain
$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\JoinDomain -Verbose

$m = Get-DscLocalConfigurationManager
if ($error.Count -gt $errorCount)
{
    $errorCode = 100

    foreach ($err in $error[$errorCount..($error.Count-1)])
    {
        Write-Host "FullyQualifiedErrorId: $($err.FullyQualifiedErrorId)"
        Format-List -InputObject $err

        # Look for retryable errors
        if ($err.FullyQualifiedErrorId -match "FailToJoinDomainFromWorkgroup")
        {
            $errorCode = 150
        }
        elseif ($err.FullyQualifiedErrorId -match "PathNotFound,Microsoft.PowerShell.Commands.GetItemPropertyCommand")
        {
            $errorCode=150
        }
        elseif ($err.FullyQualifiedErrorId -match "HRESULT 0x80338012")
        {
            # The error is "The client cannot connect to the destination specified in the request."
            # This can happen when the DC is not ready.
            $errorCode=150
        }
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