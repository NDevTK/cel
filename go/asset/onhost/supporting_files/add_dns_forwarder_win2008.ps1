# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Add DNS conditional forwarder
param(
    # name of the forwarder
    [Parameter(Mandatory=$true)] [String] $name,

    # name of the instance whose IP address is to be added as forwarder
    [Parameter(Mandatory=$true)] [String] $masterServer
)

$address = [Net.DNS]::GetHostEntry($masterServer).AddressList.IPAddressToString
dnscmd /zoneadd $name /forwarder $address
