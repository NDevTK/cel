# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

Configuration CreateIISServer
{
    Import-DscResource -Module xWebAdministration

    Node localhost
    {
        # Install the IIS role
        WindowsFeature IIS
        {
            Ensure          = "Present"
            Name            = "Web-Server"
        }

        WindowsFeature ManagementTools
        {
            Ensure          = "Present"
            Name            = "Web-Mgmt-Tools"
        }

        # Needed for Kerberos
        WindowsFeature DigestAuth
        {
            Ensure          = "Present"
            Name            = "Web-Digest-Auth"
        }

        # NTLM
        WindowsFeature WindowsAuth
        {
            Ensure          = "Present"
            Name            = "Web-Windows-Auth"
        }

        # Stop the default website
        xWebsite DefaultSite
        {
            Ensure          = "Present"
            Name            = "Default Web Site"
            State           = "Stopped"
            PhysicalPath    = "C:\inetpub\wwwroot"
            DependsOn       = "[WindowsFeature]IIS"
        }
    }
}

CreateIISServer

Start-DscConfiguration -Wait -Force -Path .\CreateIISServer -Verbose
