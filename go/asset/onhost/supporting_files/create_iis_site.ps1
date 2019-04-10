# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    # website name
    [Parameter(Mandatory=$true)] [String] $name,

    # "HTTP" or "HTTPS"
    [Parameter(Mandatory=$true)] [String] $Protocol,

    # The port to bind the website on
    [Parameter(Mandatory=$true)] [String] $Port,

    # "NONE", "NTLM", "KERBEROS" or "KERBEROS_NEGOTIABLE2"
    [Parameter(Mandatory=$true)] [String] $Authentication
)

Configuration CreateWebSite
{
    param
    (
        [Parameter(Mandatory=$true)] [String] $PhysicalPath
    )

    Import-DscResource -Module xWebAdministration

    Node localhost
    {
        # To print auth information on the test page.
        WindowsFeature AspNet
        {
            Ensure          = "Present"
            Name            = "Web-Asp-Net"
        }

        # Create the new Website with HTTP
        xWebsite NewWebsite
        {
            Ensure          = "Present"
            Name            = $name
            State           = "Started"
            PhysicalPath    = $PhysicalPath

            BindingInfo = MSFT_xWebBindingInformation
            {
                Protocol = $Protocol
                Port = $Port
            }

            AuthenticationInfo = MSFT_xWebAuthenticationInformation
            {
                Anonymous = ($Authentication -eq "NONE")
                Basic = $false
                Windows = ($Authentication -ne "NONE")
                Digest = $false
            }

            ApplicationPool = "Classic .NET AppPool"
            DependsOn = "[WindowsFeature]AspNet"
        }

        xIisFeatureDelegation UnlockWindowsAuthConfig {
            Path = "MACHINE/WEBROOT/APPHOST"
            Filter = "system.webServer/security/authentication/windowsAuthentication"
            OverrideMode = "Allow"
        }
    }
}

$websiteRoot = "C:\inetpub\wwwroot_$name"

CreateWebSite -PhysicalPath $websiteRoot

if (-not (Test-Path $websiteRoot))
{
    New-Item -Type Directory $websiteRoot
}

$text = @"
<html>
  <%
    Dim server_auth_type As String = Request.ServerVariables("AUTH_TYPE")
    Dim server_http_auth As String = Request.ServerVariables("HTTP_AUTHORIZATION")
    Dim auth_type As String = ""
    If (Len(server_http_auth) > Len(server_auth_type))
      Dim auth_token As String = server_http_auth.Substring(Len(server_auth_type) + 1)
      Dim decoded_auth_token As String = System.Text.ASCIIEncoding.ASCII.GetString(System.Convert.FromBase64String(auth_token))
      If (decoded_auth_token.StartsWith("NTLMSSP"))
        auth_type = "NTLM"
      ElseIf (auth_token.Substring(0, 1) = "Y")
        auth_type = "SPNEGO"
      Else
        auth_type = "Unknown"
      End If
    End If

    Dim username As String = ""
    If (User.Identity.IsAuthenticated)
      username = User.Identity.Name
    Else
      auth_type = "Anonymous"
    End If
  %>
  <head>
    <title>[<%= auth_type %>]<%= username %></title>
  </head>
  <body>
    <h1>Test Page - $name</h1>

    This is a test page for the '$name' site. [Protocol=$Protocol] [Port=$Port] [Authentication=$Authentication]

    <ul>
      <li><strong>AUTH_TYPE: </strong><%= server_auth_type %></li>
      <li><strong>HTTP_AUTHORIZATION: </strong><%= server_http_auth %></li>
      <ul>
        <li><strong>Looks like: </strong><%= auth_type %> (Security Event Log has more details)</li>
      </ul>
    </ul>
  </body>
</html>
"@ | Out-File "$websiteRoot\test.aspx" -Verbose

$errorCount = $error.Count
Start-DscConfiguration -Wait -Force -Path .\CreateWebSite -Verbose
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

if ($Authentication -ne "NONE")
{
    Write-Host "Setting up authentication config for IIS Site '$name'"

    # Clear the default auth providers list and fill it back according to $Providers
    $sitePath = ('IIS:\Sites\' + $name);
    $windowsAuthPath = "system.webServer/security/authentication/windowsAuthentication";
    $providersPath = "$windowsAuthPath/providers";

    Remove-WebConfigurationProperty -PSPath $sitePath -Location $sitePath -Filter $providersPath -Name "."

    if ($Authentication -eq "KERBEROS")
    {
        Add-WebConfigurationProperty -PSPath $sitePath -Location $sitePath  -Filter $providersPath -Name "." -Value @{value='Negotiate'}
    }
    elseif ($Authentication -eq "KERBEROS_NEGOTIABLE2")
    {
        # Set up Explicit Kerberos Authentication (without NTLM fallback).

        # Negotiable 2 providers can't be used with Kernel-mode authentication.
        # https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2008-R2-and-2008/cc771945(v=ws.11)
        Set-WebConfigurationProperty -PSPath $sitePath -Location $sitePath -Filter $windowsAuthPath -Name "useKernelMode" -Value "False"

        Add-WebConfigurationProperty -PSPath $sitePath -Location $sitePath  -Filter $providersPath -Name "." -Value @{value='Negotiate:Kerberos'}
    }
    else # NTLM
    {
        Add-WebConfigurationProperty -PSPath $sitePath -Location $sitePath -Filter $providersPath -Name "." -Value @{value='NTLM'}
    }

    # Setting AuthPersistSingleRequest to True means that IIS resets the auth at the end of each request
    # to force re-authentication. This allows us to display reliable auth information on the Test Page.
    Set-WebConfigurationProperty -PSPath $sitePath -Location $sitePath -Filter $windowsAuthPath -Name "authPersistSingleRequest" -Value "True"

    if ($error.Count -gt $errorCount)
    {
      Write-Host "Error Occurred"
      Exit 100
    }
}

# Add a firewall rule to allow other machines on the network to connect to the site
$ruleName = "IIS Site - $name (Port $Port)"

# If the firewall rule already exists, don't add it again
netsh advfirewall firewall show rule name=$ruleName
if ($LastExitCode -eq 0)
{
    Exit 0
}

Write-Host "Setting up firewall rule '$ruleName'"

netsh advfirewall firewall add rule name=$ruleName dir=in protocol=tcp localport=$Port action=allow

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

# The user can now access http://localhost/test.aspx
