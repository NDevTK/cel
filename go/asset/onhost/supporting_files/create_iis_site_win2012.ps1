# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    # website name
    [Parameter(Mandatory=$true)] [String] $name,

    # "HTTP" or "HTTPS"
    [Parameter(Mandatory=$true)] [String] $Protocol
)

Configuration CreateWebSite
{
    Import-DscResource -Module xWebAdministration

    Node 'localhost'
    {
        # Create the new Website with HTTP
        xWebsite NewWebsite
        {
            Ensure          = "Present"
            Name            = $name
            State           = "Started"
            PhysicalPath    = "C:\inetpub\wwwroot"
            AuthenticationInfo = MSFT_xWebAuthenticationInformation
            {
                Anonymous = $false
                Basic = $false
                Windows = $true
                Digest = $true
            }
        }
    }
}

CreateWebSite

$text = @"
<html>
  <h1>Test Page</h1>

  This is a test page
</html>
"@ | Out-File "C:\inetpub\wwwroot\test.html" -Verbose

$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\CreateWebSite -Verbose
if ($error.Count -gt $errorCount)
{
  # Exit with error code
  Write-Host "Error Occurred"
  Exit 100
}

# The user can now access http://localhost/test.html