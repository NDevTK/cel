# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This script is used by nested VM to fix its DNS entry. It adds an DNS
# entry on $dnsServerName for $computerName in $domainName, setting its
# address to $ipAddress.
# The purpose of this script is enable users to access the nested VM using host
# names in addition to IP. Unfortunately, after some time, the client will
# refresh its DNS entry on the DNS server. We'll see if we can find a solution.
param(
    [Parameter(Mandatory=$true)] [String] $adminName,
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    [Parameter(Mandatory=$true)] [String] $dnsServerName,
    [Parameter(Mandatory=$true)] [String] $domainName,
    [Parameter(Mandatory=$true)] [String] $computerName,
    [Parameter(Mandatory=$true)] [String] $ipAddress
)

$cred = New-Object System.Management.Automation.PSCredential ($adminName, (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

$scriptBlock = [scriptBlock]::Create("dnscmd . /RecordAdd $domainName $computerName A $ipAddress")
Invoke-Command -ComputerName $dnsServerName -ScriptBlock $scriptBlock  -credential $cred
