# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This PowerShell script sets a number of performance optimizations that will make nested win10 VM instances less sluggish to work with.
# The measures taken are from this article:
# https://thegeekpage.com/windows-10-high-cpu-and-ram-usage-problem-fix/

Write-Host "Config clear Page File at shutdown..."
REG ADD "HKLM\System\CurrentControlSet\Control\Session Manager\Memory Management" /f /v "ClearPageFileAtShutDown" /t REG_DWORD /d 1

Write-Host "Disabling the SysMain service..."
Stop-Service SysMain
Set-Service SysMain -StartupType Disabled

Write-Host "Disabling Background Apps..."
REG ADD "HKLM\Software\Policies\Microsoft\Windows\AppPrivacy" /f /v "LetAppsRunInBackground" /t REG_DWORD /d 2

Write-Host "Disabling Runtime Broker..."
REG ADD "HKLM\System\CurrentControlSet\Services\TimeBrokerSvc" /f /v "Start" /t REG_DWORD /d 4

Write-Host "Disabling Action Center..."
If (!(Test-Path "HKCU:\Software\Policies\Microsoft\Windows\Explorer")) {
  New-Item -Path "HKCU:\Software\Policies\Microsoft\Windows\Explorer" | Out-Null
}

REG ADD "HKCU\Software\Policies\Microsoft\Windows\Explorer" /f /v "DisableNotificationCenter" /t REG_DWORD /d 1
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\PushNotifications" /f /v "ToastEnabled" /t REG_DWORD /d 0
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\ContentDeliveryManager" /f /v "SubscribedContent-338389Enabled" /t REG_DWORD /d 0
REG ADD "HKLM\Software\Policies\Microsoft\Windows\CloudContent" /f /v "UseActionCenterExperience" /t REG_DWORD /d 0

Write-Host "Adjusting system settings for best performance..."
# The collection of registry settings that accounts for the 'best performance' system setting is taken from
# https://social.technet.microsoft.com/Forums/windowsserver/en-US/73d72328-38ed-4abe-a65d-83aaad0f9047/adjust-for-best-performance?forum=winserverpowershell
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\VisualEffects" /f /v "VisualFXSetting" /t REG_DWORD /d 2
REG ADD "HKCU\Control Panel\Desktop\WindowMetrics" /f /v "MinAnimate" /t REG_SZ /d "0"
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" /f /v "TaskbarAnimations" /t REG_SZ /d "0"
REG ADD "HKLM\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" /f /v "TaskbarAnimations" /t REG_SZ /d "-"
REG ADD "HKCU\Software\Microsoft\Windows\DWM" /f /v "CompositionPolicy" /t REG_DWORD /d 0
REG ADD "HKCU\Software\Microsoft\Windows\DWM" /f /v "ColorizationOpaqueBlend" /t REG_DWORD /d 0
REG ADD "HKCU\Software\Microsoft\Windows\DWM" /f /v "AlwaysHibernateThumbnails" /t REG_DWORD /d 0
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\Policies\Explorer" /f /v "DisableThumbnails" /t REG_DWORD /d 1
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" /f /v "ListviewAlphaSelect" /t REG_DWORD /d 0
REG ADD "HKCU\Control Panel\Desktop" /f /v "DragFullWindows" /t REG_SZ /d "0"
REG ADD "HKCU\Control Panel\Desktop" /f /v "FontSmoothing" /t REG_SZ /d "0"
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" /f /v "ListviewShadow" /t REG_DWORD /d 0
REG ADD "HKCU\Software\Microsoft\Windows\CurrentVersion\ThemeManager" /f /v "ThemeActive" /t REG_SZ /d "0"
REG ADD "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\ThemeManager" /f /v "ThemeActive" /t REG_SZ /d "-"
REG ADD "HKCU\Control Panel\Desktop" /f /v "UserPreferencesMask" /t REG_BINARY /d 9012018010000000
