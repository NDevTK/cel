# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

param (
    # The path to the api key to use.
    [Parameter(Mandatory=$true)] [String] $Path
)

Write-Host "Reading key at $Path."

$key = gsutil cat $Path
if ($LastExitCode -ne 0)
{
    Write-Host "Error occurred while getting Chromium API keys."
    Write-Host "Some features might be unavailable. [Error = $LastExitCode]"
    Exit 0
}

# Source: https://www.chromium.org/developers/how-tos/api-keys
[Environment]::SetEnvironmentVariable("GOOGLE_API_KEY", $key, [System.EnvironmentVariableTarget]::Machine)
[Environment]::SetEnvironmentVariable("GOOGLE_DEFAULT_CLIENT_ID", "celab", [System.EnvironmentVariableTarget]::Machine)
[Environment]::SetEnvironmentVariable("GOOGLE_DEFAULT_CLIENT_SECRET", "unused", [System.EnvironmentVariableTarget]::Machine)
