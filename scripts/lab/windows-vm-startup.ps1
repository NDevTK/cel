#Requires -Version 2.0
#Requires -RunAsAdministrator

# Startup script for all VMs deployed on the lab.
# 
# Make sure Win-RM is installed and running.
# Create and install a certificate if one is missing.
# Copy over files from the correct GCS bucket.
# Signal successful start via serial port.

# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
# Part of Chromium's Enterprise Lab.
# See https://g3doc/company/teams/chrome/chrome_authentication_lab.md
#
# This script should be run on a freshly provisioned Windows machine in order to
# set it up for use within the enterprise lab:

Set-Location -Path C:\

$BootstrapPath = "gs://chrome-auth-lab-staging/bootstrap"

$ToolsPath = "C:\tools"
$LabToolsPath = "C:\tools\lab"
$KeysPath = "C:\keys"

$GceApiHeaders = @{
  "Metadata-Flavor" = "Google"
}

# EnsureDirectoryExists creates the directory specified by the Path argument if
# it doesn't already exist.
function EnsureDirectoryExists {
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

# AddToPathString safely appends |source| to |target|. The latter is considered
# to be a semicolon delimited string, such as PATH. Does nothing if |source|
# already exists in |target|.
# 
# Returned the new string. Does not modify |source| or |target|.
# 
# E.g.:
#    AddToPathString -source "bar" -target "foo"
#    # Returns "foo;bar"
#    
#    AddToPathString -source "bar" -target "foo;bar"
#    # Returns "foo;bar"
function AddToPathString {
    [CmdletBinding()]
    param([string]$target, [string]$source)

    if ( ( $target -split ";" ) -contains $source ) {
        return $target
    }
    return @( $target, $source ) -join ";"
}

# EnsureSearchPathContainsTarget makes sure that $env:Path contains the
# $pathToCheck. If it is not already there, then the cmdlet will add the path to
# the system Path environment variable as a suffix.
function EnsureSearchPathContainsTarget {
    [CmdletBinding()]
    param([string]$pathToCheck)

    $env:Path = AddToPathString -target $env:Path -source $pathToCheck

    $sysPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
    $newSysPath = AddToPathString -target $sysPath -source $pathToCheck
    if ( $newSysPath -eq $sysPath ) {
        Write-Output "path already contains $pathToCheck"
    } else {
        try {
            [System.Environment]::SetEnvironmentVariable("Path", $newSysPath, [System.EnvironmentVariableTarget]::Machine)
            Write-Output "System path updated to contain $pathToCheck"
        } catch [System.Security.SecurityException] {
            Write-Output "Re-run this script as administrator."
            exit
        }
    }
}

# QueryInstanceMetadata returns either the instance or project metdata with the
# specified name.
# 
# If $projectScoped is $true and the metadata attribute named by $name is not
# defined for the instance, then attribute is looked up in the project
# metadata.
function QueryInstanceMetadata {
	[CmdletBinding()]
	param([string]$name, [bool]$projectScoped)

	$url = @(
	  "http://metadata.google.internal/computeMetadata/v1/instance/attributes",
	  $name ) -join "/"

	try {
	  return Invoke-RestMethod -Uri $url -Headers $GceApiHeaders
	} catch [System.Net.WebException] {
	  # Fallthrough
	}

	if ( ! $projectScoped ) {
	  Write-Output "Failed to lookup instance metadata for $name"
	  return ""
	}

	$url = @(
	  "http://metadata.google.internal/computeMetadata/v1/project/attributes",
	  $name
	) -join "/"

	try {
	  return Invoke-RestMethod -Uri $url -Headers $GceApiHeaders
	} catch [System.Net.WebException] {
	  Write-Output "Failed to lookup instance and project metadata for $name"
	  return ""
	}
}

function UpdateGoogleCloudTools {
    Write-Host -ForegroundColor Cyan "==================== Google Cloud SDK"
    # Google Cloud tools should already be installed. Update if necessary.
    # Note that this step may fail if it requires interaction. If that happens,
    # the user will need to run 'gcloud components update' separately.
    #
    # Fortunately, gcloud will nag the user periodically if this is the case.
    & gcloud components update --quiet
}

function InstallLabTools {
    Write-Host -ForegroundColor Cyan "==================== Updating Lab Tools"
    EnsureDirectoryExists -Path $LabToolsPath
    EnsureSearchPathContainsTarget -pathToCheck $LabToolsPath

	& gsutil rsync -d -r $BootstrapPath $LabToolsPath
}

function EnsureWinRM {
    Write-Host -ForegroundColor Cyan "==================== Configuring WinRM"

	& winrm quickconfig

	Enable-PSRemoting -force
}

# EnsureAndDumpMachineCertificate creates a machine certificate if one doesn't
# exist. It then dumps the contents of the certificate to COM4.
function EnsureAndDumpMachineCertificate {
}

UpdateGoogleCloudTools
InstallLabTools
EnsureWinRM
EnsureAndDumpMachineCertificate

