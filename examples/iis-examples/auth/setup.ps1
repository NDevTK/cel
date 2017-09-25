# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
#
#Requires -RunAsAdministrator

# Run this script to configure an existing IIS server to serve the example
# websites. The script is meant to be idempotent in that running it multiple
# times is harmless.
# 
# Note: Only use this script on a throwaway IIS server that's not exposed to
# the internet. The settings that are being made here are not recommended for a
# production server.

# Current path is $PSScriptRoot

# The data for the websites happen to exist in the same directory tree as this
# script.
$siteData = $PSScriptRoot

# These changes allow lower level 'web.config' files to configure the specified
# configuration paths. Without this, a default IIS 7 instance won't allow our
# web.config files to specify authentication parameters on a per directory
# basis. By default these settings are only configurable on a per server basis.
Set-WebConfiguration //system.webServer/security/authentication/windowsAuthentication `
   -metadata overrideMode -value Allow -PSPath IIS:/
Set-WebConfiguration //system.webServer/security/authentication/basicAuthentication `
   -metadata overrideMode -value Allow -PSPath IIS:/
Set-WebConfiguration //system.webServer/security/authentication/anonymousAuthentication `
   -metadata overrideMode -value Allow -PSPath IIS:/

# The Default Web Site on a fresh install of IIS is using port 80. We are going
# to reroute port 80 to our website.
# 
# Note: if this IIS server is serving something other than the default website
# on port 80, then the next steps will fail. You'll need to manually disable
# whatever is using port 80 in order for the example to work.
Stop-Website "Default Web Site"


# The first webiste is going to bind to port 80.
$websiteName = "webauth"
$websitePath = "IIS:\Sites\" + $websiteName
if ( Test-Path($websitePath) ) {
  Remove-Website $websiteName
}
New-Website -name $websiteName -PhysicalPath $siteData -port 80
New-WebApplication -name "ntlm" -Site $websiteName -PhysicalPath ( $siteData + "\ntlm" )
Start-Website $websiteName

# Set up a local user for testing.
Remove-LocalUser -name foo

# This is just a dumb complex looking password that's meant to pass the default
# password complexity requirements. We can reconfigure those as well so that we
# can use something even simpler. But I'm not going to bother.
# 
# Note: We are creating a user with a known password on an IIS server. This is
# why you shouldn't run this script on a server that's exposed to the internet.
$password = "tH1s1zaP4zzw0rd" 
New-LocalUser foo -Password ( ConvertTo-SecureString $password -AsPlainText -Force ) -description "Test user"
Write-Host "Password for user foo is " $password

