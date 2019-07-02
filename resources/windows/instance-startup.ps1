#Requires -Version 2.0
#Requires -RunAsAdministrator

# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Part of Chromium's Enterprise Lab. See
# https://chromium.googlesource.com/enterprise/cel/+/HEAD/README.md
#
# Startup script for all Windows VMs deployed in CEL.

Set-Location -Path C:\

# Additional headers to be included when invoking a REST API call to GCP.
$GceApiHeaders = @{
  "Metadata-Flavor" = "Google"
}

# Assert-Directory creates the directory specified by the Path argument if
# it doesn't already exist.
function Assert-Directory {
    [CmdletBinding()]
    param ([string]$Path)

    if ( Test-Path -Path $Path ) {
        Write-Output "$Path already exists"
        return
    }

    try {
        New-Item -Path $Path -ItemType directory
    } catch {
        Write-Error "Failed to create $Path"
        exit
    }
}

# Join-PathComponents safely appends |Source| to |Target|. The latter is
# considered to be a semicolon delimited string, such as the contents of
# $env:Path. Does nothing if |Source| already exists in |Target|.
# 
# Returned the new string. Does not modify |Source| or |Target|.
# 
# E.g.:
#    Join-PathComponents -source "bar" -target "foo"
#    # Returns "foo;bar"
#    
#    Join-PathComponents -source "bar" -target "foo;bar"
#    # Returns "foo;bar"
function Join-PathComponents {
    [CmdletBinding()]
    param([string]$Target, [string]$Source)

    if ( ( $Target -split ";" ) -contains $Source ) {
        return $Target
    }
    return @( $Target, $Source ) -join ";"
}

# Assert-DirectoryIsInPath makes sure that $env:Path contains the
# $path. If it is not already there, then the cmdlet will add the path to
# the system Path environment variable as a suffix.
function Assert-DirectoryIsInPath {
    [CmdletBinding()]
    param(
      [Parameter(Mandatory=$true,ValueFromPipeline=$true)]
      [string]
      $Path
    )
    $env:Path = Join-PathComponents -target $env:Path -source $Path

    $sysPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
    $newSysPath = Join-PathComponents -target $sysPath -source $Path
    if ( $newSysPath -eq $sysPath ) {
        Write-Output "path already contains $Path"
    } else {
        try {
            [System.Environment]::SetEnvironmentVariable("Path", $newSysPath, [System.EnvironmentVariableTarget]::Machine)
            Write-Output "System path updated to contain $Path"
        } catch [System.Security.SecurityException] {
            Write-Output "Re-run this script as administrator."
            exit
        }
    }
}

<#
.SYNOPSIS

Get-GcpMetadata returns either the instance or project metdata with the
specified name.

.PARAMETER Name

Name of attribute to query.

.PARAMETER ProjectScoped

If specified, the metadata attribute named by $name is queried in the project
metadata *if* it is not defined in the instance metadata.

.INPUTS

[string] name.

.OUTPUTS

The objects resulting from querying the metadata.

.EXAMPLE

  C:\PS> Get-GcpMetadata -name "foo"
  bar

#>
function Get-GcpMetadata {
  [CmdletBinding()]
  param(
    [Parameter(Mandatory=$true,ValueFromPipeline=$true)] [string] $Name,
    [Parameter()] [switch]$ProjectScoped
  )

  process {
    $url = ""
    if ( $ProjectScoped ) {
      $url = @(
          "http://metadata.google.internal/computeMetadata/v1/project/attributes",
          $Name
          ) -join "/"
    } else {
      $url = @(
          "http://metadata.google.internal/computeMetadata/v1/instance/attributes",
          $Name ) -join "/"
    }
    $url += "?alt=json"

    try {
      Invoke-RestMethod -Uri $url -Headers $GceApiHeaders
    } catch [System.Net.WebException] {
      throw "Failed to lookup instance and project metadata for $Name"
    }
  }
}

$metadata = Get-GcpMetadata -Name "cel-agent" -ProjectScoped
Write-Host "metadata cel-agent is $metadata"

$cel_agent = $metadata | ConvertFrom-Json
$cel_agent_path = $cel_agent.win_agent_x64.abs_path
$cel_ui_agent_path = $cel_agent.win_ui_agent_x64.abs_path

$cel_manifest = Get-GcpMetadata -Name "cel-manifest" -ProjectScoped
Write-Host "metadata cel-manifest is $cel_manifest"

Assert-Directory c:\cel
cd c:\cel
gsutil cp $cel_agent_path cel_agent.exe
gsutil cp $cel_ui_agent_path cel_ui_agent.exe
gsutil cp $cel_manifest cel_manifest.textpb
c:\cel\cel_agent.exe cel_manifest.textpb