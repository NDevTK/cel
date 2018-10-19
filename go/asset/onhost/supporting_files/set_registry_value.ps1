# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    # The path to the registry key to edit
    [Parameter(Mandatory=$true)] [String] $Path,

    # The name of the value to set
    [Parameter(Mandatory=$true)] [String] $Name,

    # The type of the value to set
    [Parameter(Mandatory=$true)] [String] $Type,

    # The data to set the value to
    [Parameter(Mandatory=$true)] [String] $Data
)

Write-Host "Setting registry key $Path : $Name = $Data ($Type)"

reg add $Path /v $Name /t $Type /d $Data /f

if ($LastExitCode -ne 0)
{
    Write-Host "Error occurred"
    Exit 100
}