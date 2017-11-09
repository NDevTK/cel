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
import logging
import os
import re
import subprocess
import sys

SOURCE_PATH = os.path.dirname(os.path.realpath(__file__))
GOPATH = SOURCE_PATH
GOOS = {
  "linux2": "linux",
  "linux": "linux",
  "win32": "windows",
  "cygwin": "windows",
  "darwin": "darwin"
}.get(sys.platform, "windows")

GO_LAB_PACKAGE = 'go/lab'

GO_TOOLS = ['go/cmd/cel_admin']

GO_DEPENDENCIES = [
  'cloud.google.com/go/logging', 'github.com/maruel/subcommands',
  'golang.org/x/net/context', 'golang.org/x/oauth2/google',
  'google.golang.org/api/cloudkms/v1', 'google.golang.org/api/compute/v1',
  'google.golang.org/api/iam/v1',
  'google.golang.org/api/cloudresourcemanager/v1'
]


def MergeEnv(args):
  if 'GOPATH' in os.environ:
    if sys.platform == 'win32':
      gopath = '{};{}'.format(GOPATH, os.environ['GOPATH'])
    else:
      gopath = '{}:{}'.format(GOPATH, os.environ['GOPATH'])
  else:
    gopath = GOPATH

  go_env = {'GOPATH': gopath}

  if args is not None and args.goos:
    go_env['GOOS'] = args.goos

  environ_copy = os.environ.copy()
  environ_copy.update(go_env)
  return environ_copy


def GenerateProtobufCode():
  '''Invokes 'go generate' which generates the protobuf code.

  For convenience of building, the generated code is checked into the source
  tree. This is also the norm for 'go generate' where we want to avoid creating
  a dependency on protoc being available on the build machine.'''
  subprocess.check_call(
      ['go', 'generate', '.'],
      env=MergeEnv(None),
      cwd=os.path.abspath(os.path.join(SOURCE_PATH, 'src', 'go', 'lab')))


def BuildCommand(args):
  '''Builds native binaries using Go.

When building for the native platform (i.e. no --goos option is given),
binaries are placed under GOPATH/bin. This is the default location binaries are
placed in when invoking 'go install'.

It is possible to invoke a cross compilation by specifying --goos. In that
case, the resulting build artifacts are placed under $GOPATH/bin/$GOOS/.'''

  GenerateProtobufCode()

  verb = 'install'
  bin_dir = os.path.join(SOURCE_PATH, 'bin')
  cross_compile = False

  # Cross compiling
  if args.goos and args.goos != GOOS:
    verb = 'build'
    bin_dir = os.path.abspath(os.path.join(SOURCE_PATH, 'bin', args.goos))
    cross_compile = True

  if not os.path.exists(bin_dir):
    os.mkdir(bin_dir)

  if not cross_compile:
    subprocess.check_call(
        ['go', verb, GO_LAB_PACKAGE], env=MergeEnv(args), cwd=bin_dir)

  for tool in GO_TOOLS:
    subprocess.check_call(['go', verb, tool], env=MergeEnv(args), cwd=bin_dir)


def _RunCommand(args, **kwargs):
  logging.info("Run command '%s' in directory %s", ' '.join(args),
               kwargs.get('cwd', None))

  subprocess.check_call(
      args,
      **kwargs)


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

  test_env = MergeEnv(None)

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

  for target in [GO_LAB_PACKAGE] + GO_TOOLS:
    _RunCommand(['go', 'test', target] + args.gotest_args,
                env=test_env,
                cwd=os.path.abspath(
                    os.path.join(SOURCE_PATH, 'src', 'go', 'lab')))


def GenCommand(args):
  '''Generate protobuf code.

Should be run after changing any of the *.proto files. This re-generates the Go
protobuf code based on the .proto files.
'''
  GenerateProtobufCode()


def DepsCommand(args):
  '''Update Go dependencies.

Fetches and installs the latest versions of the Go dependencies required for
building.
'''
  flags = []
  if args.v:
    flags += ['-v']
  subprocess.check_call(
      ['go', 'get', '-u'] + flags + GO_DEPENDENCIES, env=MergeEnv(None))
  print """Successfully updated Go dependencies.

Git submodule dependencies were not updated.

Fetch Git submodules that are required by this repository with:

    git submodule update --init --recursive

To roll upstream changes to all submodule dependencies, use:

    git submodule update --remote

Note: You need to manually verify all new changes to submodules."""


def ShowEnvCommand(args):
  '''Show the Go environment used for building.

Use the --goos option to see the Go environment used for cross compiling.
'''
  subprocess.check_call(['go', 'env'], env=MergeEnv(args))


def RunCommand(args):
  '''Run a command under the build environment.

The specified command will be executed with environment variables configured
for 'go build'. If the command requires passing commandline arguments, preface
the entire command with '--' to prevent the arguments from being interpreted as
arguments for this script.
'''
  subprocess.check_call(args.prog, env=MergeEnv(args))


def main():
  common_options = argparse.ArgumentParser(add_help=False)
  common_options.add_argument(
      '--goos', '-O', help='Set GOOS', choices=['windows', 'darwin', 'linux'])

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
  build_command.set_defaults(closure=BuildCommand)

  # ----------------------------------------------------------------------------
  # test
  # ----------------------------------------------------------------------------
  test_command = subparsers.add_parser(
      'test',
      help=TestCommand.__doc__.splitlines()[0],
      description=TestCommand.__doc__,
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
      formatter_class=argparse.RawDescriptionHelpFormatter)
  gen_command.set_defaults(closure=GenCommand)

  # ----------------------------------------------------------------------------
  # deps
  # ----------------------------------------------------------------------------
  deps_command = subparsers.add_parser(
      'deps',
      help=DepsCommand.__doc__.splitlines()[0],
      description=DepsCommand.__doc__,
      formatter_class=argparse.RawDescriptionHelpFormatter)
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

  parser.add_argument('-v', action='store_true', help='verbose output')
  args = parser.parse_args()
  if args.v:
    logging.basicConfig(level=logging.INFO,
                        format=('%(asctime)s %(levelname)s %(filename)s:'
                                '%(lineno)s] %(message)s '))

  try:
    args.closure(args)
  except subprocess.CalledProcessError:
    sys.exit(1)


if __name__ == '__main__':
  main()
