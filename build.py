#!/usr/bin/env python

# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This script is for building the code in src/go. It will install required
# dependencies as a part of the build.
#
# The build process may change substantially. In particular, it might be worth
# setting it up so that users can check out the source under their `GOPATH` and
# just invoke `go install .`.
#
# Also, currently, in addition to PowerShell and Go, there's a build-time
# dependency on Python. This should be more-or-less reasonable, but one that we
# can eliminate fairly easily.
#
# Currently, full builds are only supported on Windows.

import argparse
import errno
import logging
import shutil
import os
import re
import subprocess
import sys
import textwrap

SOURCE_PATH = os.path.dirname(os.path.realpath(__file__))
BUILD_PATH = os.path.join(SOURCE_PATH, 'build')
STAMP_PATH = os.path.join(BUILD_PATH, 'stamps')

sys.path.append(BUILD_PATH)
from markdown_utils import FormatMarkdown

# NATIVE_GOOS is the GOOS that corresponds to the native platform. Any tool
# that needs to run on the host machine must be built for this OS regardless of
# the target GOOS.
NATIVE_GOOS = {
    "linux2": "linux",
    "linux": "linux",
    "win32": "windows",
    "cygwin": "windows",
    "darwin": "darwin"
}.get(sys.platform, "windows")


def _MergeEnv(args, target_host=False):
  go_env = {}

  effective_goos = NATIVE_GOOS
  if args is not None and args.goos and not target_host:
    effective_goos = args.goos
  go_env['GOOS'] = effective_goos
  environ_copy = os.environ.copy()
  environ_copy.update(go_env)
  return environ_copy


def _EnsureDir(path_to_dir):
  if not os.path.exists(path_to_dir):
    os.makedirs(path_to_dir)


def _RunCommand(args, **kwargs):
  logging.info("%s [CWD: %s, GOOS: %s]",
               ' '.join([(x if ' ' not in x else '"' + x + '"') for x in args]),
               kwargs.get('cwd', os.getcwd()),
               kwargs.get('env', os.environ).get('GOOS', NATIVE_GOOS))

  subprocess.check_call(args, **kwargs)


def _InstallDep(args):
  if (not hasattr(args, 'install')) or not args.install:
    raise Exception(
        textwrap.dedent('''\
            "dep" command not found.

            The CEL project uses "deps" to manage dependencies. You can get it by following
            the instructions at :

                https://github.com/golang/dep#setup

            A quick and portable way to get it is to run the following:

                go get -u github.com/golang/dep/cmd/dep

            Rerun as 'build.py deps --install' to install dependencies automatically. If
            you've already done so, it may be that the GOBIN path is not in the system
            PATH.
            '''))

  verbose_flag = []
  if hasattr(args, 'verbose') and args.verbose:
    verbose_flag += ["-v"]

  _RunCommand(
      ['go', 'get', '-u'] + verbose_flag + ['github.com/golang/dep/cmd/dep'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)


def _InstallGoProtoc(args):
  if (not hasattr(args, 'install')) or not args.install:
    raise Exception(
        textwrap.dedent('''\
            "protoc-gen-go" not found.

            The CEL project relies on generating Go code for Google ProtoBuf files. In
            addition to the Protocol Buffers Compiler (protoc), Go support requires
            protoc-gen-go which generates Go code. More information can be found including
            installation instructions at https://github.com/golang/protobuf.

            A quick and portable way to get it is to run the following:

                go get -u github.com/golang/protobuf/protoc-gen-go

            Rerun this script as 'build.py deps --install' to install "protoc-gen-go"
            automatically. You you've already done so, it may be that the GOBIN path is not
            in the system.
            '''))

  verbose_flag = []
  if hasattr(args, 'verbose') and args.verbose:
    verbose_flag += ["-v"]

  _RunCommand(
      ['go', 'get', '-u'] + verbose_flag +
      ['github.com/golang/protobuf/protoc-gen-go'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)


def _InstallProtoc(args):
  raise Exception(
      textwrap.dedent('''\
          "protoc" not found.

          The CEL project relies on generating Go code for Google ProtoBuf files. This
          requires having the ProtoBuf compiler in the PATH.

          Instructions for installing "protoc" can be found at
          https://developers.google.com/protocol-buffers/docs/downloads

          Unfortunately, protoc can't be installed automatically. So you'll need to
          install it manually. If you've arleady installed it, it's possible that the
          installed location is not in the system PATH.
          '''))


def _IsSentinelNewer(sentinel_path, *sources):
  if not os.path.exists(sentinel_path):
    return False
  basetime = os.path.getmtime(sentinel_path)
  for source in sources:
    if os.path.getmtime(source) > basetime:
      return False
  return True


def _Deps(args):
  '''Ensures dependencies are present.'''

  # Max number of times we are going to retry if a component installation fails.
  MAX_RETRY_COUNT = 3

  def _CheckAndInstall(command, installer, **kwargs):
    for x in range(MAX_RETRY_COUNT):
      try:
        _RunCommand(command, **kwargs)
      except OSError as e:
        if e.errno == errno.ENOENT:
          installer(args)
          continue
        raise e
      except subprocess.CalledProcessError:
        break
      break

  verbose_flag = []
  if hasattr(args, 'verbose') and args.verbose:
    verbose_flag += ["-v"]

  with open(os.devnull, 'r+') as f:
    _CheckAndInstall(
        ['protoc-gen-go'],
        _InstallGoProtoc,
        env=_MergeEnv(args, target_host=True),
        cwd=SOURCE_PATH,
        stdin=f,
        stdout=f,
        stderr=f)
    _CheckAndInstall(
        ['protoc', '-h'],
        _InstallProtoc,
        env=_MergeEnv(args, target_host=True),
        cwd=SOURCE_PATH,
        stdin=f,
        stdout=f,
        stderr=f)

  # Using a sentinel file to make sure we only run 'dep' if either the last run
  # was unsuccessful or if there has been a change to Gopkg.* files.
  _EnsureDir(STAMP_PATH)
  if not os.path.exists(os.path.join(STAMP_PATH, 'README')):
    with open(os.path.join(STAMP_PATH, 'README'), 'w') as f:
      f.write(
          textwrap.dedent('''\
                  This directory contains timestamp files.

                  Feel free to delete these. The only visible effect would be
                  that the build might take a bit longer to run.'''))

  sentinel = os.path.join(STAMP_PATH, 'deps.stamp')
  if _IsSentinelNewer(sentinel, os.path.join(SOURCE_PATH, 'Gopkg.toml'),
                      os.path.join(SOURCE_PATH, 'Gopkg.lock')):
    return

  _CheckAndInstall(
      ['dep', 'ensure'] + verbose_flag,
      _InstallDep,
      env=_MergeEnv(args),
      cwd=SOURCE_PATH)

  with open(sentinel, 'w') as f:
    pass


def _Generate(args):
  '''Generates Go code based on .proto files.'''

  _EnsureDir(os.path.join(SOURCE_PATH, 'schema', 'gcp', 'compute'))
  _EnsureDir(os.path.join(SOURCE_PATH, 'go', 'gcp', 'compute'))

  _RunCommand(
      [
          'go', 'run', 'go/cmd/gen_api_proto/main.go', '-i',
          'vendor/google.golang.org/api/compute/v0.beta/compute-api.json', '-o',
          'schema/gcp/compute/compute-api.proto', '-p',
          'chromium.googlesource.com/enterprise/cel/go/gcp', '-g',
          'go/gcp/compute/validate.go'
      ],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _RunCommand(
      [
          'protoc', '--go_out=../../../', 'schema/common/options.proto',
          'schema/common/fileref.proto', 'go/common/testdata/testmsgs.proto'
      ],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _RunCommand(
      [
          'protoc', '--go_out=../../../', 'schema/asset/active_directory.proto',
          'schema/asset/cert.proto', 'schema/asset/dns.proto',
          'schema/asset/iis.proto', 'schema/asset/network.proto',
          'schema/asset/asset_manifest.proto', 'schema/asset/machine.proto'
      ],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _RunCommand(
      ['protoc', '--go_out=../../../', 'schema/host/host_environment.proto'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _RunCommand(
      [
          'protoc', '--go_out=../../../', 'schema/meta/files.proto',
          'schema/meta/command.proto', 'schema/meta/reference.proto',
          'schema/meta/status.proto'
      ],
      env=_MergeEnv(args),
      cwd=SOURCE_PATH)

  _RunCommand(
      ['protoc', '--go_out=../../../', 'schema/gcp/compute/compute-api.proto'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)


def BuildCommand(args):
  '''Builds native binaries using Go.

When building for the native platform (i.e. no --goos option is given),
binaries are placed under GOPATH/bin. This is the default location binaries are
placed in when invoking 'go install'.

It is possible to invoke a cross compilation by specifying --goos. In that
case, the resulting build artifacts are placed under build/$GOOS/bin.'''

  if not args.fast:
    _Deps(args)
    _Generate(args)

  flags = []
  if args.verbose:
    flags += ['-v', '-x']

  build_env = _MergeEnv(args)
  _RunCommand(
      ['go', 'install'] + flags + ['./go/...'], env=build_env, cwd=SOURCE_PATH)


def TestCommand(args):
  '''Run Go tests.

Invokes 'go test' to run tests. Any additional arguments are passed down to 'go
test'.

'go test' is invoked for each known build target.

Note: Tests can only be run when GOOS == GOHOSTOS. Hence there's no command
line option to set GOOS.

About Network Requests During Tests
-----------------------------------

The test framework will intercept network requests and attempt to simulate
Google Cloud API calls. This requires a set of cached responses per each test
case. As tests evolve, this set of requests may need to be refreshed. Hence
there is a --reset option to the 'test' command. This option causes all the
cached response data to be discarded. The tests are then executed in 'record'
mode where network requests make it out into the network and all network
traffic is logged. Assuming this causes the tests to pass, you can then 'git
add' and 'git commit' the new or updated cached network requests.

TL;DR: After changing a test that makes network requests, you may see errors like:

        foo_test.go|...| Get https://www.googleapis.com/... : not implemented

    That means that the test is attempting to make a new network request.
    Assuming this is correct, do the following:

        ./build.py test --reset

    This will cause the tests to be run in 'record' mode. If the tests pass,
    then:

        git add src/go/lab/testdata
        git commit -m <...>

    You should now be able to run the tests normally again.

Note: When using 'reset and record' mode, don't specify any special 'go test'
options.

Note: Running tests in 'reset and record' mode requires access to the
chrome-auth-lab-dev Google Cloud project.
'''

  test_env = _MergeEnv(args, target_host=True)

  if args.reset:
    if len(args.gotest_args) != 0:
      sys.stderr.write(
          '--reset option cannot be used with additional \'go test\' arguments\n'
      )
      return

    testdata_dir = os.path.join(SOURCE_PATH, 'src', 'go', 'lab', 'testdata')
    for p in os.listdir(testdata_dir):
      testcase_dir = os.path.join(testdata_dir, p)
      if not os.path.isdir(testcase_dir):
        continue

      for q in ['requests', 'responses']:
        d = os.path.join(testcase_dir, q)
        if not os.path.isdir(d):
          continue

        for f in os.listdir(d):
          filepath = os.path.join(d, f)
          if not os.path.isfile(filepath):
            continue

          os.remove(filepath)

    test_env['LAB_RECORD'] = 'yes'

  _RunCommand(['go', 'test', './go/...'], env=test_env, cwd=SOURCE_PATH)


def GenCommand(args):
  '''Generate protobuf code.

Should be run after changing any of the *.proto files. This re-generates the Go
protobuf code based on the .proto files.
'''
  _Deps(args)
  _Generate(args)


def DepsCommand(args):
  '''Check for and ensure build dependencies.

Ensures that required build tools and Go packages are installed and ready to
use. Use the '--install' option to attempt to install missing build tools.
'''
  _Deps(args)


def ShowEnvCommand(args):
  '''Show the Go environment used for building.

Use the --goos option to see the Go environment used for cross compiling.
'''
  _RunCommand(['go', 'env'], env=_MergeEnv(args))


def RunCommand(args):
  '''Run a command under the build environment.

The specified command will be executed with environment variables configured
for 'go build'. If the command requires passing commandline arguments, preface
the entire command with '--' to prevent the arguments from being interpreted as
arguments for this script.
'''
  _RunCommand(args.prog, env=_MergeEnv(args))


def _FormatMarkdownFiles(args):

  o = subprocess.check_output(
      ['git', 'ls-files', '--exclude-standard', '--', '*.md'], cwd=SOURCE_PATH)
  md_files = o.splitlines()

  for f in md_files:
    FormatMarkdown(os.path.join(SOURCE_PATH, f))


def _FormatGoFiles(args):
  _RunCommand(['go', 'fmt', './go/...'], env=_MergeEnv(args), cwd=SOURCE_PATH)


def _FormatProtoFiles(args):
  o = subprocess.check_output(
      ['git', 'ls-files', '--exclude-standard', '--', '*.proto'],
      cwd=SOURCE_PATH)
  proto_files = o.splitlines()

  try:
    _RunCommand(['clang-format', '-i', '-style=Chromium'] + proto_files)
  except OSError as e:
    if e.errno == errno.ENOENT:
      sys.stderr.write(
          textwrap.dedent('''\
                    clang-format not found.

                    clang-format is used for formatting ProtoBuf files. Without
                    it, this script can't correctly format ProtoBuf files.'''))
    else:
      raise e


def _FormatPythonFiles(args):
  try:
    _RunCommand(['yapf', '-i', '-r', '.'], env=_MergeEnv(args), cwd=SOURCE_PATH)
  except OSError as e:
    if e.errno == errno.ENOENT:
      sys.stderr.write(
          textwrap.dedent('''\
              YAPF not found.

              YAPF is used for formatting Python files. See https://github.com/google/yapf
              for more information on how to install YAPF. Without it, this script can't
              correctly format Python source files.

              You can still land code if your change doesn't touch any Python files. If you
              do modify Python files, it's likely that someone will have to reformat the
              files later.
              '''))
    else:
      raise e


def FormatCommand(args):
  '''Reformat code and prepare for a code commit.

This command performs the following:

    1. Update all Markdown documentation with the latest set of tags from
       `docs/index.md`.

    2. Format Go code in the tree using Gofmt.

    3. Format Python files using YAPF. This project uses the Chromium Python
       coding style [1]. See https://github.com/google/yapf for information on
       installing YAPF.

[1]: https://chromium.googlesource.com/chromium/src/+/master/styleguide/styleguide.md
'''
  _FormatMarkdownFiles(args)
  _FormatGoFiles(args)
  _FormatProtoFiles(args)
  _FormatPythonFiles(args)


def CleanCommand(args):
  '''Remove build artifacts.
'''
  force_option = ['-f' if args.force else '-n']
  _RunCommand(
      ['git', 'clean', '-X'] + force_option,
      env=_MergeEnv(args),
      cwd=SOURCE_PATH)
  build_dir = os.path.join(SOURCE_PATH, 'build')
  if not args.force:
    print('Would remove {}'.format(build_dir))
    return

  if os.path.exists(build_dir):
    print('Removing {}'.format(build_dir))
    shutil.rmtree(build_dir)


def main():
  common_options = argparse.ArgumentParser(add_help=False)
  common_options.add_argument(
      '--goos', '-O', help='Set GOOS', choices=['windows', 'darwin', 'linux'])
  common_options.add_argument(
      '--verbose', '-v', help='Verbose output', action='store_true')

  parser = argparse.ArgumentParser(
      description='build and manage Chrome Enterprise Lab tools',
      formatter_class=argparse.RawDescriptionHelpFormatter)
  subparsers = parser.add_subparsers(help='sub-command help')

  # ----------------------------------------------------------------------------
  # build
  # ----------------------------------------------------------------------------
  build_command = subparsers.add_parser(
      'build',
      help=BuildCommand.__doc__.splitlines()[0],
      description=BuildCommand.__doc__,
      parents=[common_options],
      formatter_class=argparse.RawDescriptionHelpFormatter)
  build_command.add_argument(
      '--fast',
      '-F',
      action='store_true',
      help='''Fast build. Skips dependency and generator steps''')
  build_command.set_defaults(closure=BuildCommand)

  # ----------------------------------------------------------------------------
  # test
  # ----------------------------------------------------------------------------
  test_command = subparsers.add_parser(
      'test',
      help=TestCommand.__doc__.splitlines()[0],
      description=TestCommand.__doc__,
      parents=[common_options],
      formatter_class=argparse.RawDescriptionHelpFormatter)
  test_command.add_argument(
      '--reset',
      '-R',
      action='store_true',
      help=
      'remove all cached network request data and record all new network requests'
  )
  test_command.add_argument(
      'gotest_args',
      metavar='ARGS',
      type=str,
      help='''Aruments to pass down to "go test".
      Preface with "--" to disambiguate from arguments passed in to this build tool.''',
      nargs='*')
  test_command.set_defaults(closure=TestCommand)

  # ----------------------------------------------------------------------------
  # gen
  # ----------------------------------------------------------------------------
  gen_command = subparsers.add_parser(
      'gen',
      help=GenCommand.__doc__.splitlines()[0],
      description=GenCommand.__doc__,
      parents=[common_options],
      formatter_class=argparse.RawDescriptionHelpFormatter)
  gen_command.set_defaults(closure=GenCommand)

  # ----------------------------------------------------------------------------
  # clean
  # ----------------------------------------------------------------------------
  clean_command = subparsers.add_parser(
      'clean',
      help=CleanCommand.__doc__,
      parents=[common_options],
      formatter_class=argparse.RawDescriptionHelpFormatter)
  clean_command.add_argument(
      '--force',
      '-f',
      action='store_true',
      help='Force. Without this option this command doesn\'t do anything.')
  clean_command.set_defaults(closure=CleanCommand)

  # ----------------------------------------------------------------------------
  # deps
  # ----------------------------------------------------------------------------
  deps_command = subparsers.add_parser(
      'deps',
      help=DepsCommand.__doc__.splitlines()[0],
      description=DepsCommand.__doc__,
      parents=[common_options],
      formatter_class=argparse.RawDescriptionHelpFormatter)
  deps_command.add_argument(
      '--install',
      '-I',
      action='store_true',
      help='Install additional dependencies')
  deps_command.set_defaults(closure=DepsCommand)

  # ----------------------------------------------------------------------------
  # env
  # ----------------------------------------------------------------------------
  env_command = subparsers.add_parser(
      'env',
      help=ShowEnvCommand.__doc__.splitlines()[0],
      description=ShowEnvCommand.__doc__,
      formatter_class=argparse.RawDescriptionHelpFormatter,
      parents=[common_options])
  env_command.set_defaults(closure=ShowEnvCommand)

  # ----------------------------------------------------------------------------
  # format
  # ----------------------------------------------------------------------------
  format_command = subparsers.add_parser(
      'format',
      help=FormatCommand.__doc__.splitlines()[0],
      description=FormatCommand.__doc__,
      formatter_class=argparse.RawDescriptionHelpFormatter,
      parents=[common_options])
  format_command.set_defaults(closure=FormatCommand)

  # ----------------------------------------------------------------------------
  # run
  # ----------------------------------------------------------------------------
  run_command = subparsers.add_parser(
      'run',
      help=RunCommand.__doc__.splitlines()[0],
      description=RunCommand.__doc__,
      formatter_class=argparse.RawDescriptionHelpFormatter,
      parents=[common_options])
  run_command.add_argument(
      'prog', metavar='ARG', type=str, help='Program and arguments', nargs='+')
  run_command.set_defaults(closure=RunCommand)

  args = parser.parse_args()
  if hasattr(args, 'verbose') and args.verbose:
    logging.basicConfig(
        level=logging.INFO,
        format=('%(asctime)s %(levelname)s %(filename)s:'
                '%(lineno)s] %(message)s '))

  try:
    args.closure(args)
  except subprocess.CalledProcessError:
    sys.exit(1)


if __name__ == '__main__':
  main()
