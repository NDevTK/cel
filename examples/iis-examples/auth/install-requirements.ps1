# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
#
#Requires -RunAsAdministrator

# Run this script with elevated privileges on a freshly provisioned Windows
# Server instance to install all the components needed for running this
# example.
# 
# In this case, it'd be IIS along with the authentication modules for Basic and
# Windows authentication.
# 
# Note: 'Web-Mgmg-Console' is optional and only needed if you plan on poking at
# the IIS instance by hand.
$requiredFeatures = @(
  "Web-Common-HTTP",
  "Web-Security",
  "Web-Mgmt-Console",
  "Web-Windows-Auth",
  "Web-Basic-Auth" )

Install-WindowsFeature -Name $requiredFeatures

