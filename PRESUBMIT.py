# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import sys


def CheckIfFilesNeedFormatting(input_api, output_api):
  """CheckIfFilesNeedFormatting runs CheckFormatting from the build.py script
  to check if any modified files need formatting.

  It would be great if we can hook the code formatting stuff to 'git cl format'
  and have one of the canned presubmit checks verify this, but that's not
  currently possible with depot_tools.
  """

  files = input_api.AbsoluteLocalPaths()
  files = [f for f in files if input_api.os_path.exists(f)]
  if len(files) == 0:
    return []

  sys.path.append(input_api.PresubmitLocalPath())
  from build import CheckFormatting

  modified = CheckFormatting(files)

  if len(modified) == 0:
    return []

  return [
      output_api.PresubmitError(
          'Modified files require formatting. Please run "python build.py format"',
          modified)
  ]


def CheckChangeOnUpload(input_api, output_api):
  return CheckIfFilesNeedFormatting(input_api, output_api)


def CheckChangeOnCommit(input_api, output_api):
  return []
