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


def RunGoTests(input_api, output_api):
  results = []
  p = input_api.subprocess.Popen(
      [
          'go', 'test', '-p={}'.format(input_api.cpu_count), '-json', '-vet',
          'off', './go/...'
      ],
      cwd=input_api.PresubmitLocalPath(),
      stdout=input_api.subprocess.PIPE,
      stderr=input_api.subprocess.PIPE)
  output, err = p.communicate()
  failures = []

  for o in output.splitlines():
    d = input_api.json.loads(o)
    if "Action" not in d or "Test" not in d:
      continue
    if d["Action"] != "fail":
      continue
    failures.append(d)

  if len(failures) != 0:
    results.append(
        output_api.PresubmitError(
            message="Go tests failed",
            items=[
                'Test {} in {}'.format(d["Test"], d["Package"])
                for d in failures
            ]))

  return results


def RunGoVet(input_api, output_api):
  results = []
  p = input_api.subprocess.Popen(
      ['go', 'vet', './go/...'],
      cwd=input_api.PresubmitLocalPath(),
      stdout=input_api.subprocess.PIPE,
      stderr=input_api.subprocess.PIPE)
  output, err = p.communicate()

  failures = []
  for o in err.splitlines():
    o = o.strip()
    if o == "" or o.startswith("# "):
      continue
    failures.append(o)

  if len(failures) != 0:
    results.append(
        output_api.PresubmitError(message='Go vet failed', items=failures))
  return results


def CommonChecks(input_api, output_api):
  """CommonChecks runs common validation steps that apply to both upload and
  commit."""
  results = []
  results.extend(CheckIfFilesNeedFormatting(input_api, output_api))
  results.extend(RunGoTests(input_api, output_api))
  results.extend(RunGoVet(input_api, output_api))
  results.extend(
      input_api.canned_checks.PanProjectChecks(input_api, output_api))
  return results


def CheckChangeOnUpload(input_api, output_api):
  results = []
  results.extend(CommonChecks(input_api, output_api))
  return results


def CheckChangeOnCommit(input_api, output_api):
  results = []
  results.extend(CommonChecks(input_api, output_api))
  return results
