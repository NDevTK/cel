# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.


# Create a user in a domain.
param(
    # the FQDN name of the domain.
    [Parameter(Mandatory=$true)] [String] $domainName,

    # the password of the domain administrator.
    [Parameter(Mandatory=$true)] [String] $adminPassword,

    # the name of the user to be created.
    [Parameter(Mandatory=$true)] [String] $userName,

    # the password of the user.
    [Parameter(Mandatory=$true)] [String] $password,

    [Parameter()] [String] $description = ""
  )  
  
Configuration AddUser
{
    Import-DscResource -Module xActiveDirectory

    Node localhost
    {
        xADUser AddUser
        {
            DomainName = $domainName
            DomainAdministratorCredential = $domainCred
            UserName = $userName
            Password = $userCred
            Description = $description
            Ensure = "Present"
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

$domainCred = New-Object System.Management.Automation.PSCredential ("$domainName\administrator", (ConvertTo-SecureString $adminPassword -AsPlainText -Force))

$userCred = New-Object System.Management.Automation.PSCredential ("$domainName\$userName", (ConvertTo-SecureString $password -AsPlainText -Force))

AddUser -ConfigurationData $ConfigData

# Add the user
Start-DscConfiguration -Wait -Force -Path .\AddUser -Verbose
