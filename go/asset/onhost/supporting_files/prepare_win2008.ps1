# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# WinServer2008 has scheduled tasks for consistency that can't be disabled (removed in WMF 5.0).
# Calling `Start-DscConfiguration` during those tasks fails and can't be bypassed with `-Force`.
# This will remove Pending DSC configs. We cycle through all onhost scripts after restarts so
# incomplete configurations will continue when we get back to it.
function Get-ConfigurationFilesCsv {
  $items = Get-ChildItem "C:\Windows\System32\Configuration\" | % { $_.FullName }
  return $items -join ", "
}

Write-Host "Configuration files before: " + (Get-ConfigurationFilesCsv)

Remove-Item C:\Windows\System32\Configuration\*.mof
Get-Process *WMI* | ? { $_.modules.ModuleName -like "*DSC*" } | Stop-Process -Force

Write-Host "Configuration files after: " + (Get-ConfigurationFilesCsv)
