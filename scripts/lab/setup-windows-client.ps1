#Requires -RunAsAdministrator

# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
# Part of Chromium's Enterprise Lab.
# See https://g3doc/company/teams/chrome/chrome_authentication_lab.md
#
# This script should be run on a freshly provisioned Windows machine in order to
# set it up for use within the AD.CHROME domain. It will create our directory
# structure, which looks like:
#
# C:
# |
# +- tools
# |    +-- chrome-auth-lab
# |    |     A git checkout of the 'chrome-auth-lab-tools' repository from
# |    |     google.com:chrome-auth-lab gcloud project
# |    |
# |    +-- depot_tools
# |    |     A git checkout of https://chromium.googlesource.com/chromium/tools/depot_tools.git
# |    |
# |    +-- luci-py
# |          A git clone of github.com/luci/luci-py
# |
# +- keys
# |    Google Cloud keys for accessing the isolate server
# |
# +- staging
#      Directory where isolated builds will be placed by the chrome-auth-lab tools.
# 
# In addition, the script installs Windows Debugging Tools and Microsoft Message
# Analyzer.

Set-Location -Path C:\

$ToolsPath = "C:\tools"
$DepotToolsPath = "C:\tools\depot_tools"
$LabToolsPath = "C:\tools\chrome-auth-lab-tools"
$LuciPath = "C:\tools\luci-py"
$KeysPath = "C:\keys"

$BootstrapPath = "gs://chrome-auth-lab-staging/bootstrap"

# EnsureDirectoryExists creates the directory specified by the Path argument if
# it doesn't already exist.
Function EnsureDirectoryExists {
    [CmdletBinding()]
    Param ([string]$Path)

    If ( Test-Path -Path $Path ) {
        Write-Output "$Path already exists"
        Return
    }

    Try {
        New-Item -Path $Path -ItemType directory
    } Catch {
        Write-Error "Failed to create $Path"
        Exit
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
Function AddToPathString {
    [CmdletBinding()]
    Param([string]$target, [string]$source)

    If ( ( $target -split ";" ) -contains $source ) {
        Return $target
    }
    Return @( $target, $source ) -join ";"
}

Function EnsureSearchPathContainsTarget {
    [CmdletBinding()]
    Param([string]$pathToCheck)

    $env:Path = AddToPathString -target $env:Path -source $pathToCheck

    $sysPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
    $newSysPath = AddToPathString -target $sysPath -source $pathToCheck
    if ( $newSysPath -eq $sysPath ) {
        Write-Output "System path already contains $pathToCheck"
    } else {
        Try {
            [System.Environment]::SetEnvironmentVariable("Path", $newSysPath, [System.EnvironmentVariableTarget]::Machine)
            Write-Output "System path updated to contain $pathToCheck"
        } Catch [System.Security.SecurityException] {
            Write-Output "Re-run this script as administrator."
            Exit
        }
    }
}

Function EnsureSearchPathContainsGit {
    $GitDir = (Get-Item -Path "${DepotToolsPath}\git*bin").Name
    $GitBinDir = "${DepotToolsPath}\${GitDir}\bin"
    EnsureSearchPathContainsTarget -pathToCheck $GitBinDir
}

Function InstallDepotTools {
    Write-Host -ForegroundColor Cyan "==================== Depot Tools"
    EnsureDirectoryExists -Path $DepotToolsPath
    EnsureSearchPathContainsTarget -pathToCheck $DepotToolsPath

    If ( Test-Path -Path "${DepotToolsPath}\.git" -PathType Container ) {
        Write-Host -ForegroundColor Green "depot_tools already exists. Updating"
	Start-Process -FilePath "gclient.bat" -WorkingDirectory $DepotToolsPath -Wait -NoNewWindow | Out-Null
	EnsureSearchPathContainsGit
        Return
    }

    $DepotToolsArchive = "depot_tools.zip"
    $LocalArchiveCopy = "${ToolsPath}\${DepotToolsArchive}"

    Write-Host -ForegroundColor Green "Copying and extracting depot_tools from $DepotToolsSource"
    & gsutil cp "${BootstrapPath}/$DepotToolsArchive" $LocalArchiveCopy
    Expand-Archive -Path $LocalArchiveCopy -DestinationPath $DepotToolsPath -Force
    Remove-Item -Path $LocalArchiveCopy

    # Run gclient once to update it. gclient should be on the path already.
    Start-Process -FilePath "gclient.bat" -WorkingDirectory $DepotToolsPath -Wait -NoNewWindow | Out-Null
    EnsureSearchPathContainsGit
}

Function InstallLabTools {
    Write-Host -ForegroundColor Cyan "==================== Lab Tools"
    EnsureDirectoryExists -Path $LabToolsPath
    EnsureSearchPathContainsTarget -pathToCheck $LabToolsPath

    If ( Test-Path -Path "${LabToolsPath}\.git" -PathType Container ) {
        Write-Host -ForegroundColor Green "$LabToolsPath already exists. Updating."
        Push-Location -Path $LabToolsPath
        & git pull origin master | Out-Null
        Pop-Location
        Return
    }

    Write-Host -ForegroundColor Green "Cloning chrome-auth-lab-tools ..."
    Push-Location -Path $ToolsPath
    & gcloud source repos clone chrome-auth-lab-tools --project=google.com:chrome-auth-lab | Out-Null
    Pop-Location
}

Function InstallLuci {
    Write-Host -ForegroundColor Cyan "==================== Luci"
    EnsureDirectoryExists -Path $LuciPath
    EnsureDirectoryExists -Path $KeysPath

    # Copy the isolate server keys. These authenticate the current host using the GCE
    # service account to the isolate server.
    Write-Host -ForegroundColor Green "Syncing lab VM keys ..."
    & gsutil rsync -d -r -C gs://chrome-auth-lab-staging/keys $KeysPath

    [System.Environment]::SetEnvironmentVariable(
        "SWARMING_AUTH_SERVICE_ACCOUNT_JSON",
        "${KeysPath}\chrome-auth-lab-isolate-labrat.json",
        [System.EnvironmentVariableTarget]::Machine)

    [System.Environment]::SetEnvironmentVariable(
        "ISOLATE_SERVER",
        "https://chrome-auth-lab-isolate.appspot.com",
        [System.EnvironmentVariableTarget]::Machine)

    If ( Test-Path -Path "${LuciPath}\.git" -PathType Container ) {
	Write-Host -ForegroundColor Green "Luci already exists. Updating ..."
        Push-Location -Path $LuciPath
        git pull origin
        Pop-Location
        Return
    }

    Write-Host -ForegroundColor Green "Cloing Luci-py from github..."
    Push-Location -Path $ToolsPath
    & git clone https://github.com/luci/luci-py.git $LuciPath
    Pop-Location
}

Function UpdateGoogleCloudTools {
    Write-Host -ForegroundColor Cyan "==================== Google Cloud SDK"
    # Google Cloud tools should already be installed. Update if necessary.
    # Note that this step may fail if it requires interaction. If that happens,
    # the user will need to run 'gcloud components update' separately.
    #
    # Fortunately, gcloud will nag the user periodically if this is the case.
    & gcloud components update --quiet
}

Function InstallChocolatey {
    Write-Host -ForegroundColor Cyan "==================== Chocolatey"

    if ( Test-Path -Path "C:\ProgramData\Chocolatey\choco.exe" ) {
	Write-Host -ForegroundColor Green "Chocolatey is already installed. Updating ..."
	& choco upgrade chocolatey --confirm
    }

    Write-Host -ForegroundColor Green "Installing chocolatey ..."
    Invoke-WebRequest https://chocolatey.org/install.ps1 -UseBasicParsing | Invoke-Expression
}

Function InstallDebuggingTools {
    Write-Host -ForegroundColor Cyan "==================== Microsoft Windows Debgging Tools"
    & choco upgrade windbg --confirm
}

Function InstallMessageAnalyzer {
    Write-Host -ForegroundColor Cyan "==================== Microsoft Message Analyzer"
    & choco upgrade microsoft-message-analyzer --confirm
}

Function InstallWireshark {
    Write-Host -ForegroundColor Cyan "==================== Wireshark"
    & choco upgrade wireshark --confirm
}

Function InstallNotepadPlusPlus {
    Write-Host -ForegroundColor Cyan "==================== Notepad++"
    & choco upgrade notepadplusplus --confirm
}

Function UpdateDNS {
    Write-Host -ForegroundColor Cyan "==================== Updating DNS"
    # Set the DNS server to point to the AD.
    Set-DnsClientServerAddress -InterfaceAlias "Ethernet" -ServerAddresses ("10.240.0.0")
}

Function AddRemoteDesktopUsers {
    Write-Host -ForegroundColor Cyan "==================== Adding Remote Desktop Users"
    $Members = @()

		# Only works if the language is set to english.
    $RDGroup = [ADSI]"WinNT://${env:COMPUTERNAME}/Remote Desktop Users,group"
    $RDMembers = $RDGroup.Invoke("Members")
    $RDMembers | ForEach {
        $ADSPath = $_.GetType().InvokeMember("ADSPath", "GetProperty", $null, $_, $null)
	$Members += $ADSPath
    }

    If ( ( $Members -contains "WinNT://AD/Domain Users" ) -and ( $Members -contains "WinNT://AD/Domain Admins" ) ) {
    	Write-Host -ForegroundColor Green "'AD\Domain Users' and 'AD\Domain Admins' are already Remote Desktop users"
	Return
    }

    & net localgroup "Remote Desktop Users" /add "AD\Domain Users"
    & net localgroup "Remote Desktop Users" /add "AD\Domain Admins"
}

UpdateGoogleCloudTools

# Software installation.

InstallDepotTools
InstallLabTools
InstallLuci
InstallChocolatey
InstallDebuggingTools
InstallMessageAnalyzer
InstallWireshark
InstallNotepadPlusPlus

UpdateDNS

Get-ComputerInfo -Property "CsDomain" -OutVariable ComputerDomain

If ( $ComputerDomain.CsDomain -ieq "ad.chrome" ) {
	# Adding Remote Desktop users will only succeed after the computer is
	# joined to the domain.
	AddRemoteDesktopUsers
} Else {
        Write-Host -ForegroundColor Cyan "==================== Moving VM to AD.CHROME Domain"
        Write-Host -ForegroundColor Red "Don't forget to re-run this script after the machine restarts..."

	# The computer is not already part of the domain, then we need to add it
	# first. This script should then be re-run to complete the installation.
	Add-Computer -DomainName "AD.CHROME" -OUPath "OU=Clients,DC=ad,DC=chrome" -Restart -Confirm
}


Write-Host -ForegroundColor Cyan "==================== Done"

