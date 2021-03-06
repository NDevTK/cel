#!/usr/bin/env python

# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
'''
Build and manage Chrome Enterprise Lab tools.

This script is for building the code in src/go. It will install required
dependencies as a part of the build.

If this is the first time you are building the toolchain, then you likely need
to do the following:

    build.py deps --install

This will install the dependencies that are required for building the
toolchain. Once you statisfy the dependencies, you can build the toolchain for
the host platform by:

    build.py build

Or you can invoke tests by:

    build.py test

Use "build.py build --help" for more information about how the build tool works
and instructions for setting up the build to work with "go build"/"go test".

See CONTRIBUTING.md for details for contributing code upstream.
'''

import argparse
import ast
import datetime
import errno
import itertools
import logging
import multiprocessing
import os
import re
import shutil
import subprocess
import sys
import textwrap
from distutils.version import LooseVersion

# Root of the source tree.
SOURCE_PATH = os.path.dirname(os.path.realpath(__file__))

# OUT_PATH is the root of the output tree where build artifacts are placed.
OUT_PATH = os.path.join(SOURCE_PATH, 'out')

# STAMP_PATH is a directory that contains timestamp files that are used during
# the build process to detect stale build artifacts.
STAMP_PATH = os.path.join(OUT_PATH, 'stamps')

# Go package root for the CEL toolchain.
PACKAGE_ROOT = "chromium.googlesource.com/enterprise/cel/go"

# Path containing the Go package corresponding to PACKAGE_ROOT.
ROOT_GO_PATH = os.path.join(SOURCE_PATH, "go")

# Path containing third party dependencies that are not managed via 'dep'.
THIRD_PARTY_DIR = os.path.join(SOURCE_PATH, "third_party")

# Path containing github.com/googleapis/googleapis
GOOGLEAPIS_DIR = os.path.join(THIRD_PARTY_DIR, "googleapis")

sys.path.append(os.path.join(SOURCE_PATH, 'build'))
from markdown_utils import FormatMarkdown

# HOST_GOOS is the GOOS that corresponds to the host platform. Any tool that
# needs to run on the host machine must be built for this OS regardless of the
# target GOOS.
HOST_GOOS = {
    "cygwin": "windows",
    "darwin": "darwin",
    "linux": "linux",
    "linux2": "linux",
    "win32": "windows",
}.get(sys.platform, "windows")

# Used by _GetCustomBuildEnv to cache the generated build environment.
CACHED_BUILD_ENV = None

# Supported target environments for cel_agent & cel_ui_agent which run on GCE
# instances. We will build all of them under `bin/resources/`. The architecture
# used for cel_ctl is the one used to build the project.Tuple of GOOS / GOARCH.
TARGET_ARCHS_AGENT = [
    # This list should include all our supported target platforms. For
    # example, once we start supporting 32-bit Windows environments, we'd
    # add something like this:
    #
    # Note that you might need to modify the backend_prep.go file to
    # include all the platforms.
    #   ("windows", "386"),
    ("windows", "amd64"),
    ("linux", "amd64"),
]


def _GetCustomBuildEnv():
  global CACHED_BUILD_ENV

  if CACHED_BUILD_ENV is not None:
    return CACHED_BUILD_ENV

  custom_env_file = os.path.join(SOURCE_PATH, '.build.environment')
  if not os.path.exists(custom_env_file):
    CACHED_BUILD_ENV = {}
    return CACHED_BUILD_ENV

  with open(custom_env_file, 'r') as f:
    contents = f.read()
    CACHED_BUILD_ENV = ast.literal_eval(contents)
    if not isinstance(CACHED_BUILD_ENV, dict):
      raise Exception(
          textwrap.dedent('''\
                    .build.environment must be a Python literal that evaluates
                    to a dictionary. See 'build.py format --help' for more
                    details.
                    '''))

  return CACHED_BUILD_ENV


def _MergeEnv(args, target_host=False):
  go_env = {}

  effective_goos = HOST_GOOS
  if args is not None and args.goos and not target_host:
    effective_goos = args.goos
  go_env['GOOS'] = effective_goos
  if args is not None and args.goarch:
    go_env['GOARCH'] = args.goarch
  environ_copy = os.environ.copy()
  environ_copy.update(go_env)
  environ_copy.update(_GetCustomBuildEnv())
  return environ_copy


def _EnsureDir(path_to_dir):
  if not os.path.exists(path_to_dir):
    os.makedirs(path_to_dir)


def _RunCommand(args, **kwargs):
  logging.info("%s [CWD: %s, GOOS: %s]",
               ' '.join([(x if ' ' not in x else '"' + x + '"') for x in args]),
               kwargs.get('cwd', os.getcwd()),
               kwargs.get('env', os.environ).get('GOOS', HOST_GOOS))

  subprocess.check_call(args, **kwargs)


def _RunCommandOutput(args, **kwargs):
  logging.info("%s [CWD: %s, GOOS: %s]",
               ' '.join([(x if ' ' not in x else '"' + x + '"') for x in args]),
               kwargs.get('cwd', os.getcwd()),
               kwargs.get('env', os.environ).get('GOOS', HOST_GOOS))

  return subprocess.check_output(args, encoding='UTF-8', **kwargs)


def _GetDependents(fn):
  '''\
_GetDependents returns a list of strings representing the full path to the
known direct depedents of the file at |fn|.

Currently only works for .proto files.
'''

  if not fn.endswith('.proto'):
    return []

  import_re = re.compile('\s*import\s+"([^"]*)"\s*;')
  deps = []

  with open(fn, 'r') as f:
    for line in f:
      m = import_re.match(line)
      if m is None:
        continue
      p = _SourcePath(m.group(1))
      if os.path.exists(p):
        deps.append(p)
  return deps


def _SourcePath(f):
  return os.path.join(SOURCE_PATH, f)


def _ExpandArg(a, **kwargs):
  if a == '$^':
    return kwargs['inp']
  return [a.format(**kwargs)]


def _BuildStep(args, inp=[], **kwargs):
  '''\
_BuildStep takes as input a list of input files and runs a build command
if the output file or a stamp file is found to be out of date.

In other words, it acts as a mini build step which only runs if the inputs are
newer than the outputs. As a special case, it attempts to determine the imports
of a '.proto' file and also takes into account the timestamps of the dependent
files.

Recognized keyword arguments are:

    inp: List[string]
        The list of input files. Can be paths relative to SOURCE_PATH.

    out: string
        A single output file. If specified, the timestamps of the input files
        as well as their discovered depents are compared against the modified
        time of the file at |out|. If |out| is missing, then the behavior is
        equivalent to |out| being older than the inputs.

    outpy: string
        A python proto root directory. If specified, the timestamps of the
        input files are compared against the modified time of the mapped file
        in |outpy|. If |outpy| is missing, then the behavior is equivalent to
        |outpy| being older than the inputs.

    stamp: string
        A stamp file. The behavior is equivalent to setting |out| except that
        the timestamp of the file at |stamp| is updated to the current time if
        the build step was successful.

All remaining keyword arguments are passed into subprocess.check_call().

The build command specified as a List[string] in  the |args| argument can
contain str.format() style references to keyword arguments. The special
argument string '$^' expands to |inp|.
'''

  deps = list(
      set(
          itertools.chain.from_iterable(
              [_GetDependents(_SourcePath(f)) for f in inp])))
  deps.extend([_SourcePath(f) for f in inp])

  if 'stamp' in kwargs and _IsTimestampNewer(
      _SourcePath(kwargs['stamp']), *deps):
    logging.info('Skipping _BuildStep [stamp: %s]', kwargs['stamp'])
    return

  # Verify if output files are more recent than the input files (and skip gen).
  if 'out' in kwargs and _IsTimestampNewer(_SourcePath(kwargs['out']), *deps):
    skip_generate = True

    if 'outpy' in kwargs:
      for f in inp:
        pyProto = f.replace('.proto', '_pb2.py').replace('-', '_')
        pyProtoPath = os.path.join(kwargs['outpy'], pyProto)
        if not _IsTimestampNewer(pyProtoPath, _SourcePath(f)):
          skip_generate = False

    if skip_generate:
      logging.info('Skipping _BuildStep [out: %s]', kwargs['out'])
      return

  kwargs['inp'] = inp
  args = list(
      itertools.chain.from_iterable([_ExpandArg(a, **kwargs) for a in args]))
  stamp = kwargs['stamp'] if 'stamp' in kwargs else None

  del kwargs['inp']
  for key in ['out', 'outpy', 'stamp']:
    if key in kwargs:
      del kwargs[key]

  logging.info("%s [CWD: %s, GOOS: %s]",
               ' '.join([(x if ' ' not in x else '"' + x + '"') for x in args]),
               kwargs.get('cwd', os.getcwd()),
               kwargs.get('env', os.environ).get('GOOS', HOST_GOOS))

  subprocess.check_call(args, **kwargs)

  if stamp is not None:
    with open(stamp, 'w') as f:
      pass


def _InitializeGoModules(args):
  # Print go env for debugging.
  _RunCommand(['go', 'env'])

  if (not hasattr(args, 'install')) or not args.install:
    _RunCommand(['ls', '-l', 'go.mod'], cwd=SOURCE_PATH)
    _RunCommand(['cat', 'go.mod'], cwd=SOURCE_PATH)
    raise Exception(
        textwrap.dedent('''\
            Go modules is not initialized.

            Go uses Go modules to manage dependencies. The first time
            compiling a project, one needs to initialize Go modules
            by running 'go mod init' at the <CEL src>/go directory.

            Rerun as 'build.py deps --install' to initialize Go modules
            automatically. If you've already done so, it may be that the GOBIN
            path is not in the system PATH.
            '''))

  _RunCommand(['go', 'env', '-w', 'GO111MODULE=on'])
  _RunCommand(['go', 'mod', 'init'],
              env=_MergeEnv(args, target_host=True),
              cwd=SOURCE_PATH)
  _RunCommand(['cat', 'go.mod'], cwd=SOURCE_PATH)


def _InstallGoProtoc(args):
  if (not hasattr(args, 'install')) or not args.install:
    raise Exception(
        textwrap.dedent('''\
            "protoc-gen-go" not found.

            The CEL project relies on generating Go code for Google ProtoBuf
            files. In addition to the Protocol Buffers Compiler (protoc), Go
            support requires protoc-gen-go which generates Go code. More
            information can be found including installation instructions at
            https://github.com/golang/protobuf.

            Rerun this script as 'build.py deps --install' to install
            "protoc-gen-go" automatically. If you've already done so,
            it may be that the GOBIN path is not in the system.
            '''))

  verbose_flag = []
  if hasattr(args, 'verbose') and args.verbose:
    verbose_flag += ["-v"]

  _RunCommand(
      ['go', 'install'] + verbose_flag +
      ['./vendor/github.com/golang/protobuf/protoc-gen-go'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)


def _InstallProtoc(args):
  raise Exception(
      textwrap.dedent('''\
          "protoc" not found or is too old. The version should be at least
          3.5.1 as reported by 'protoc --version'.

          The CEL project relies on generating Go code for Google ProtoBuf
          files. This requires having the ProtoBuf compiler in the PATH.

          Instructions for installing "protoc" can be found at
          https://developers.google.com/protocol-buffers/docs/downloads

          Unfortunately, protoc can't be installed automatically. So you'll
          need to install it manually. If you've arleady installed it, it's
          possible that the installed location is not in the system PATH.
          '''))


def _IsTimestampNewer(sentinel_path, *sources):
  '''\
Returns true if any of the `sources` has a timestamp that's newer than
`sentinel_path`.

All of `sources` and `sentinel_path` are full paths to files.
'''
  if not os.path.exists(sentinel_path):
    logging.info('  %s does not exist', sentinel_path)
    return False
  basetime = os.path.getmtime(sentinel_path)
  for source in sources:
    sourcetime = os.path.getmtime(source)
    if sourcetime > basetime:
      logging.info('  %s is newer than %s [%s ; %s]', source, sentinel_path,
                   sourcetime, basetime)
      return False
    else:
      logging.info('  %s is older than %s [%s ; %s]', source, sentinel_path,
                   sourcetime, basetime)
  return True


def _Deps(args):
  '''Ensures dependencies are present.'''

  # Max number of times we are going to retry if a component installation fails.
  MAX_RETRY_COUNT = 3

  def _CheckAndInstall(command, installer, trigger_installer_exceptions,
                       ignored_exceptions, **kwargs):
    succeeded = False
    for x in range(MAX_RETRY_COUNT):
      try:
        _RunCommand(command, **kwargs)
      except OSError as e:
        if e.errno == errno.ENOENT:
          installer(args)
          continue
        raise e
      except trigger_installer_exceptions:
        installer(args)
        continue
      except ignored_exceptions:
        pass
      succeeded = True
      break

    if not succeeded:
      raise Exception(
          textwrap.dedent('''\
              Failed _CheckAndInstall for `{}`.
              It may be that the GOBIN path is not in the system PATH.
              '''.format(command)))

  verbose_flag = []
  if hasattr(args, 'verbose') and args.verbose:
    verbose_flag += ["-v"]

  with open(os.devnull, 'r+') as f:
    _CheckAndInstall(
        ['protoc', '--version'],
        _InstallProtoc,
        (),
        # protoc-gen-go can fail with 'error:no files to generate'
        # which we consider a success
        subprocess.CalledProcessError,
        env=_MergeEnv(args, target_host=True),
        cwd=SOURCE_PATH,
        stdin=f,
        stdout=f,
        stderr=f)

  o = subprocess.check_output(['protoc', '--version']).decode().strip()
  if o.startswith('libprotoc '):
    if LooseVersion(o[len('libprotoc '):]) < LooseVersion("3.5.1"):
      raise Exception(
          textwrap.dedent('''\
          The version of ProtoBuf compiler installed on this machine is too
          old.  The version as reported by protoc is "{}". It should be at
          least 3.5.1 to build the CEL toolchain.

          Instructions for installing "protoc" can be found at
          https://developers.google.com/protocol-buffers/docs/downloads

          Unfortunately, protoc can't be installed automatically. So you'll
          need to install it manually. If you've arleady installed it, it's
          possible that the installed location is not in the system PATH.
          '''.format(o)))
  else:
    raise Exception(
        textwrap.dedent('''\
    "protoc --version" returned an unexpected string. Returned string was:
       "{}"

    Expected something like "libprotoc 1.2.3"
    '''.format(o)))

  # Using a sentinel file to make sure we only run 'go mod ...' if either the
  # last run was unsuccessful or if there has been a change to the go.mod or
  # go.sum files.
  _EnsureDir(STAMP_PATH)
  if not os.path.exists(os.path.join(STAMP_PATH, 'README')):
    with open(os.path.join(STAMP_PATH, 'README'), 'w') as f:
      f.write(
          textwrap.dedent('''\
                  This directory contains timestamp files.

                  Feel free to delete these. The only visible effect would be
                  that the build might take a bit longer to run.'''))

  update_modules = hasattr(args, 'update') and args.update

  # Constructs a directory named vendor in the main module's root directory
  # that contains copies of all packages needed to support builds and tests
  # of packages in the main module.
  _RunCommand(['go', 'env'], cwd=SOURCE_PATH)

  _CheckAndInstall(
      ['go', 'mod', 'vendor', '-e'] + verbose_flag,
      _InitializeGoModules,
      # If one calls 'go mod vendor' the first time before initializing
      # Go Modules (one-time only per repo), go will fail, which should
      # trigger the script to initialize Go Modules.
      subprocess.CalledProcessError,
      (),
      env=_MergeEnv(args),
      cwd=SOURCE_PATH)

  sentinel = os.path.join(STAMP_PATH, 'deps.stamp')
  if not update_modules and _IsTimestampNewer(
      sentinel, os.path.join(SOURCE_PATH, 'go.mod'),
      os.path.join(SOURCE_PATH, 'go.sum')):
    return

  if update_modules:
    # Adds any missing module requirements necessary to build the current
    # module's packages and dependencies, and removes requirements on modules
    # that don't provide any relevant packages.
    _RunCommand(['go', 'mod', 'tidy'], cwd=SOURCE_PATH)
    # Re-sync the dependencies.
    _RunCommand(['go', 'mod', 'vendor'], cwd=SOURCE_PATH)

  with open(os.devnull, 'r+') as f:
    _CheckAndInstall(['protoc-gen-go'],
                     _InstallGoProtoc, (), (),
                     env=_MergeEnv(args, target_host=True),
                     cwd=SOURCE_PATH,
                     stdin=f,
                     stdout=f,
                     stderr=f)

  with open(sentinel, 'w') as f:
    pass

  # Thirdparty Protos
  _EnsureDir(THIRD_PARTY_DIR)

  if not os.path.exists(GOOGLEAPIS_DIR):
    subprocess.check_call(
        ['git', 'clone', 'https://github.com/googleapis/googleapis.git'],
        cwd=THIRD_PARTY_DIR)
  if update_modules:
    subprocess.check_call(['git', 'pull', 'origin', 'master'],
                          cwd=GOOGLEAPIS_DIR)


def _Generate(args):
  '''\
Generates Go code based on .proto files.

Requires `protoc` be present on PATH. Use _Deps() to install `protoc` if its
missing.
'''

  _EnsureDir(STAMP_PATH)

  descriptor_path = os.path.join(OUT_PATH, 'schema')
  _EnsureDir(descriptor_path)
  _EnsureDir(os.path.join(SOURCE_PATH, 'go', 'schema'))

  gen_api_command = _BuildCommand('gen_api_proto', './go/tools/gen_api_proto',
                                  _MergeEnv(args, target_host=True))
  gen_api_invocation = [
      gen_api_command,
      '-i',
      '{inp[0]}',
      '-o',
      '{out}',
      '-s',
      'chromium.googlesource.com/enterprise/cel/go/schema',
      '-p',
      'chromium.googlesource.com/enterprise/cel/go/schema/gcp',
  ]

  _EnsureDir(os.path.join(SOURCE_PATH, 'schema', 'gcp', 'compute'))
  _EnsureDir(os.path.join(SOURCE_PATH, 'go', 'gcp', 'compute'))
  _BuildStep(
      gen_api_invocation + ['-g', 'go/gcp/compute/validate.go'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH,
      inp=['vendor/google.golang.org/api/compute/v1/compute-api.json'],
      out='schema/gcp/compute/compute-api.proto')

  _EnsureDir(os.path.join(SOURCE_PATH, 'go', 'gcp', 'cloudkms'))
  _EnsureDir(os.path.join(SOURCE_PATH, 'schema', 'gcp', 'cloudkms'))
  _BuildStep(
      gen_api_invocation + ['-g', 'go/gcp/cloudkms/validate.go'],
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH,
      inp=['vendor/google.golang.org/api/cloudkms/v1/cloudkms-api.json'],
      out='schema/gcp/cloudkms/cloudkms-api.proto')

  python_proto_path = os.path.join(SOURCE_PATH, "test", "chrome_ent_test",
                                   "infra", "proto")
  protoc_command = [
      'protoc', '--go_out=../../../', '--descriptor_set_out={out}',
      '--include_source_info', '--proto_path=.',
      '--python_out={}'.format(python_proto_path),
      '--proto_path={}'.format(GOOGLEAPIS_DIR), '$^'
  ]

  _BuildStep(
      protoc_command,
      inp=[
          'schema/common/validation.proto',
          'schema/common/file_reference.proto',
          'schema/common/secret.proto',
          'schema/common/testmsgs.proto',
      ],
      out=os.path.join(descriptor_path, 'common.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=[
          'schema/asset/active_directory.proto', 'schema/asset/cert.proto',
          'schema/asset/dns.proto', 'schema/asset/iis.proto',
          'schema/asset/network.proto', 'schema/asset/asset_manifest.proto',
          'schema/asset/machine.proto', 'schema/asset/remote_desktop.proto'
      ],
      out=os.path.join(descriptor_path, 'asset.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['schema/host/host_environment.proto'],
      out=os.path.join(descriptor_path, 'host.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['schema/lab/lab.proto'],
      out=os.path.join(descriptor_path, 'lab.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['schema/gcp/agent_metadata.proto'],
      out=os.path.join(descriptor_path, 'gcp_agent_metadata.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['schema/gcp/compute/compute-api.proto'],
      out=os.path.join(descriptor_path, 'gcp_compute.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['schema/gcp/cloudkms/cloudkms-api.proto'],
      out=os.path.join(descriptor_path, 'gcp_cloudkms.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['schema/gcp/runtimeconfig/runtimeconfig-api.proto'],
      out=os.path.join(descriptor_path, 'gcp_runtimeconfig.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      protoc_command,
      inp=['go/tools/gen_doc_proto/testdata/test.proto'],
      out=os.path.join(SOURCE_PATH, 'go', 'tools', 'gen_doc_proto', 'testdata',
                       'test.pb'),
      outpy=python_proto_path,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  esc_command = _BuildCommand('esc', './vendor/github.com/mjibson/esc',
                              _MergeEnv(args, target_host=True))
  _BuildStep([
      esc_command, '-pkg', 'onhost', '-prefix', 'go/asset/onhost/', '-o',
      'go/asset/onhost/static.go', 'go/asset/onhost/supporting_files/'
  ])

  esc_invocation = [
      esc_command, '-pkg', 'deploy', '-prefix', 'resources', '-o', '{out}',
      '-private', '$^'
  ]
  _BuildStep(
      esc_invocation,
      inp=[
          'resources/deployment/cel-base.yaml',
          'resources/deployment/gcp-builtins.host.textpb',
          'resources/windows/instance-startup.ps1',
          'resources/linux/instance-startup.py'
      ],
      out=os.path.join(SOURCE_PATH, 'go', 'gcp', 'deploy', 'resources.gen.go'),
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  gen_doc_command = _BuildCommand('gen_doc_proto', './go/tools/gen_doc_proto',
                                  _MergeEnv(args, target_host=True))
  gen_doc_invocation = [gen_doc_command, '-out', '{out}', '$^']
  doc_path = os.path.join(SOURCE_PATH, 'docs', 'gen')
  _BuildStep(
      gen_doc_invocation,
      inp=[os.path.join(descriptor_path, 'common.pb')],
      out=os.path.join(doc_path, 'common.md'),
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      gen_doc_invocation,
      inp=[os.path.join(descriptor_path, 'asset.pb')],
      out=os.path.join(doc_path, 'asset.md'),
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      gen_doc_invocation,
      inp=[os.path.join(descriptor_path, 'host.pb')],
      out=os.path.join(doc_path, 'host.md'),
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  _BuildStep(
      gen_doc_invocation,
      inp=[os.path.join(descriptor_path, 'gcp_compute.pb')],
      out=os.path.join(doc_path, 'gcp_compute.md'),
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  env = _MergeEnv(args)
  out_dir = os.path.join(_GetBuildDir(env), 'resources')
  for goos, goarch in TARGET_ARCHS_AGENT:
    env['GOOS'] = goos
    env['GOARCH'] = goarch
    _BuildCommand('cel_agent', './go/cmd/cel_agent', env, out_dir=out_dir)
    _BuildCommand('cel_ui_agent', './go/cmd/cel_ui_agent', env, out_dir=out_dir)


def _GetBuildDir(build_env):
  '''\
Return the build directory.

This is $SOURCE_PATH/$GOOS_$GOARCH/bin.
'''
  goos = subprocess.check_output(['go', 'env', 'GOOS'],
                                 env=build_env,
                                 cwd=SOURCE_PATH,
                                 encoding='UTF-8').strip()
  goarch = subprocess.check_output(['go', 'env', 'GOARCH'],
                                   env=build_env,
                                   cwd=SOURCE_PATH,
                                   encoding='UTF-8').strip()
  return os.path.join(OUT_PATH, '{}_{}'.format(goos, goarch), 'bin')


def _BuildCommand(command,
                  package,
                  build_env,
                  build_version=None,
                  out_dir=None,
                  verbose=False):
  '''\
  _BuildCommand builds a Go command.

  '''
  flags = []
  if verbose:
    flags += ['-v', '-x']

  if build_version:
    now = datetime.datetime.utcnow().strftime("%Y-%m-%d %H:%M")
    version_string = "%s, built on %s" % (build_version, now)
    flags += ['-ldflags', '-X "main.version=%s"' % version_string]

  if out_dir is None:
    out_dir = _GetBuildDir(build_env)
  _EnsureDir(out_dir)
  suffix = '.exe' if build_env['GOOS'] == 'windows' else ''
  out = os.path.join(out_dir, command + suffix)
  _RunCommand(
      ['go', 'build'] + flags + ['-o', out, package],
      env=build_env,
      cwd=SOURCE_PATH)
  return out


def BuildCommand(args):
  '''\
Build all non-test Go source files.

Build artifacts can be found in the out/$GOOS_$GOARCH/bin directory after a
successful build.  Does not attempt to install any packages by default.

The build step also checks if the dependencies are up-to-date. It also
generates files that are needed by the build. These additional steps happen
prior to the build, and only if the dependencies have changed.

Why not just run "go build" ?

    The CEL repository doesn't include generated sources. In particular this
    includes:

    * Code generated by the Protocol Buffers compiler.
    * Prtocol buffer definitions of Google Cloud Platform REST objects and
      their corresponding generated Go code.
    * Vendored dependencies.

    The "build.py build" invocation ensures that these generated and vendored
    source files are present. It also places the resulting executables in
    platform specific directories. The latter makes it possible to do cross
    compilation.

    If you'd like to be able to invoke "go build" manually, then invoke
    "build.py deps" first. This provides the same assurances with regard to
    dependencies as running "build.py build".
'''

  if not args.fast:
    _Deps(args)

  # Generate should do minimal work if nothing has changed.
  _Generate(args)

  build_env = _MergeEnv(args)

  flags = []
  if args.verbose:
    flags += ['-v', '-x']

  if not args.fast:
    # Do a (redundant) full build. This catches build errors that don't affect
    # the go/cmd/ build that's done next.
    _RunCommand(
        ['go', 'build'] + flags + ['./go/...'], env=build_env, cwd=SOURCE_PATH)

  commands = os.listdir(os.path.join(SOURCE_PATH, 'go', 'cmd'))

  build_version = None
  with open("VERSION", 'r') as file:
    build_version = file.read()

  # Build cel_ctl with the current architecture
  for command in commands:
    _BuildCommand(
        command,
        './go/cmd/' + command,
        build_env,
        build_version=build_version,
        verbose=args.verbose)


def _GetGoPackages(root_package, root_path):
  has_go_files = False
  packages = []
  for d in os.listdir(root_path):
    this_path = os.path.join(root_path, d)
    if os.path.islink(this_path):
      continue

    if os.path.isdir(this_path):
      packages.extend(_GetGoPackages(root_package + "/" + d, this_path))
      continue

    if d.endswith(".go"):
      has_go_files = True

  if has_go_files:
    packages.append(root_package)

  return packages


def TestCommand(args):
  '''\
Run Go tests.

Ensures dependencies are present and invokes 'go test' to run tests. Any
additional arguments are passed down to 'go test'.

'build.py test' is basically equivalent to 'go test ...'. It's primarily here
for convenience when running tests on all the go packages contained herein. If
test filtering is to be performed, or you'd like to specify individual packages
to be tested, use 'go test' directly.

During development, you can invoke "build.py deps" separately and then manually
invoke "go test <...>" as you see fit.

Note: Tests can only be run when GOOS == GOHOSTOS. Hence there's no command
line option to set GOOS.
'''

  for test_arg in args.gotest_args:
    if not test_arg.startswith('-'):
      raise (Exception(
          textwrap.dedent('''\
              It looks like you are passing in package names.
              Please invoke 'go test' directly.
              ''')))

  if not args.fast:
    _Deps(args)
    _Generate(args)

  test_env = _MergeEnv(args, target_host=True)
  packages = _GetGoPackages(PACKAGE_ROOT, ROOT_GO_PATH)

  # If no coverage information is needed, parallelize the test invocation.
  if not args.coverage:
    par_arg = ['-p={}'.format(multiprocessing.cpu_count())]
    _RunCommand(
        ['go', 'test'] + par_arg + args.gotest_args + packages,
        env=test_env,
        cwd=SOURCE_PATH)
    return

  for p in packages:
    cover_flags = []

    if args.coverage:
      rel_package_name = p[len(PACKAGE_ROOT) + 1:]
      cover_profile = os.path.join(
          OUT_PATH,
          ''.join('_' if x == '/' else x for x in rel_package_name) + ".cover")
      cover_flags = [
          '-cover', '-covermode', 'atomic', '-coverprofile', cover_profile
      ]
      print('''\

Use 'go tool cover -http %s' to view coverage information in HTML.''' %
            (cover_profile))

    _RunCommand(
        ['go', 'test'] + args.gotest_args + cover_flags + [p],
        env=test_env,
        cwd=SOURCE_PATH)


def GenCommand(args):
  '''\
Generate protobuf code.

Should be run after changing any of the *.proto files. This re-generates the Go
protobuf code based on the .proto files.
'''
  _Deps(args)
  _Generate(args)


def DepsCommand(args):
  '''\
Check for and ensure build dependencies.

Ensures that required build tools and Go packages are installed and ready to
use. Use the '--install' option to attempt to install missing build tools.

Developers can use the '--update' option as shorthand for invoking 'go mod tidy
&& go mod vendor'.
'''
  _Deps(args)


def ShowEnvCommand(args):
  '''\
Show the Go environment used for building.

Use the --goos option to see the Go environment used for cross compiling.
'''
  _RunCommand(['go', 'env'], env=_MergeEnv(args))


def RunCommand(args):
  '''\
Run a command under the build environment.

The specified command will be executed with environment variables configured
for 'go build'. If the command requires passing commandline arguments, preface
the entire command with '--' to prevent the arguments from being interpreted as
arguments for this script.
'''

  build_env = _MergeEnv(args)
  run_args = {'env': build_env}
  if args.build_dir:
    run_args['cwd'] = _GetBuildDir(build_env)
  _RunCommand(args.prog, **run_args)


def _FormatMarkdownFiles(args, md_files):
  if len(md_files) == 0:
    return []

  modified = []
  for f in md_files:
    m = FormatMarkdown(os.path.join(SOURCE_PATH, f), dry_run=args.check)
    if m:
      modified.append(f)

  return modified


def _FormatGoFiles(args, go_files):
  if len(go_files) == 0:
    return []
  if args.check:
    o = _RunCommandOutput(
        ['gofmt', '-l'] + go_files, cwd=SOURCE_PATH, env=_MergeEnv(args))
    return o.splitlines()

  _RunCommand(
      ['gofmt', '-l', '-w'] + go_files,
      cwd=SOURCE_PATH,
      env=_MergeEnv(args, target_host=True))


def _CheckClangFormat(files, args):
  env = _MergeEnv(args)
  modified = []
  for f in files:
    o = _RunCommandOutput(
        ['clang-format', '-output-replacements-xml', '-style=Chromium', f],
        cwd=SOURCE_PATH,
        env=env)
    lines = o.splitlines()
    for line in lines:
      if line.startswith('<replacement '):
        modified.append(f)
        break
  return modified


def _FormatProtoFiles(args, proto_files):
  if len(proto_files) == 0:
    return []

  try:
    if args.check:
      return _CheckClangFormat(proto_files, args)

    _RunCommand(
        ['clang-format', '-i', '-style=Chromium'] + proto_files,
        env=_MergeEnv(args, target_host=True))

  except OSError as e:
    if e.errno == errno.ENOENT:
      sys.stderr.write(
          textwrap.dedent('''\
                    clang-format not found.

                    clang-format is used for formatting ProtoBuf files. Without
                    it, this script can't correctly format ProtoBuf files.'''))
    raise e

  except subprocess.CalledProcessError as e:
    if e.returncode == 1:
      sys.stderr.write(
          textwrap.dedent('''\

                      See 'build.py format --help' for more details on how to
                      configure a depot_tools provided clang_format tool to
                      work with a CEL build tree.
                      '''))
    raise e


def _FormatPythonFiles(args, py_files):
  if len(py_files) == 0:
    return []

  yapf_executable = 'yapf'
  if os.name == 'nt':
    yapf_executable = 'yapf.bat'

  try:
    if args.check:
      try:
        o = _RunCommandOutput(
            [yapf_executable, '-r', '-d'] + py_files,
            env=_MergeEnv(args, target_host=True),
            cwd=SOURCE_PATH)
      except subprocess.CalledProcessError as e:
        o = e.output

      lines = o.splitlines()
      modified = []
      for line in lines:
        if not line.startswith('--- '):
          continue
        fields = line.split()
        if len(fields) < 3:
          continue
        modified.append(fields[1])
      return modified

    _RunCommand(
        [yapf_executable, '-i', '-r'] + py_files,
        env=_MergeEnv(args, target_host=True),
        cwd=SOURCE_PATH)

  except OSError as e:
    if e.errno == errno.ENOENT:
      sys.stderr.write(
          textwrap.dedent('''\
              YAPF not found.

              YAPF is used for formatting Python files. See https://github.com/google/yapf
              for more information on how to install YAPF. Without it, this
              script can't correctly format Python source files.

              You can still land code if your change doesn't touch any Python
              files. If you do modify Python files, it's likely that someone
              will have to reformat the files later.
              '''))
    else:
      raise e


def FormatCommand(args):
  '''\
Reformat code and prepare for a code commit.

This command performs the following:

    1. Resolve imports and verify links in Markdown documents.

    2. Format Go code in the tree using Gofmt.

    3. Format Python files using YAPF. This project uses the Chromium Python
       coding style [1]. See https://github.com/google/yapf for information on
       installing YAPF.

    4. Format ProtoBuf files and textpb files using clang-format.

Problems with 'clang-format'?

  You may encounter an error which looks like the following when invoking
  'build.py format':

        Problem while looking for clang-format in Chromium source tree:
        Could not find checkout in any parent of the current path.
        Set CHROMIUM_BUILDTOOLS_PATH to use outside of a chromium checkout.

  This is due to the 'depot_tools' provided 'clang-format' script being in your
  path. It attempts to locate the 'buildtools' folder from a Chromium checkout,
  which doesn't work when you are working inside the CEL codebase.

  If this happens, you can resolve the issue using one of the following methods:

     1. Adjust your PATH variable so that a non-depot_tools clang-format binary
        is found first. -- or --

     2. If you have a Chromium checkout handy, set the CHROMIUM_BUILDTOOLS_PATH
        environment variable to point to the 'buildtools' directory. E.g. if
        your Chromium checkout is in /chromium, then:

          CHROMIUM_BUILDTOOLS_PATH=/chromium/src/buildtools ./build.py format

     3. Create a .build.environment file at the root of the CEL checkout to set
        the CHROMIUM_BUILDTOOLS_PATH: environment variable. The environment
        variables defined in .build.environment are applied to all binaries
        invoked by build.py.

        The .build.environment file consists of a Python literal in text form
        defining a dictionary whose keys are environment variable names to be
        set. The values are, of course, the value of the environment variable.
        
        E.g.: Using the same paths as the previous option:

          echo '{ "CHROMIUM_BUILDTOOLS_PATH": "/chromium/src/buildtools" }' \
            > .build.environment

        Now you should be able to invoke 'build.py' directly without having to
        set the environment variable each time.

[1]: https://chromium.googlesource.com/chromium/src/+/master/styleguide/styleguide.md
'''

  logging.info("checking annotations")
  vet_annotations_cmd = _BuildCommand('vet_annotations',
                                      './go/tools/vet_annotations',
                                      _MergeEnv(args, target_host=True))

  broken_calls = _RunCommandOutput([vet_annotations_cmd] + [
      os.path.join(SOURCE_PATH, 'go', d)
      for d in os.listdir(os.path.join(SOURCE_PATH, 'go'))
      if d != 'tools'
  ])

  if broken_calls != "":
    print(broken_calls)
    sys.exit(1)

  windows = sys.platform == 'win32'
  git_command = 'git'
  if windows:
    git_command = 'git.bat'

  o = subprocess.check_output([git_command, 'ls-files'],
                              cwd=SOURCE_PATH,
                              env=_MergeEnv(args),
                              encoding='UTF-8')
  all_files = [os.path.join(SOURCE_PATH, p) for p in o.splitlines()]

  logging.info("checking .proto files")
  pr = _FormatProtoFiles(args, [f for f in all_files if f.endswith('.proto')])

  logging.info("checking .md files")
  md = _FormatMarkdownFiles(args, [f for f in all_files if f.endswith('.md')])

  logging.info("checking .go files")
  go = _FormatGoFiles(args, [f for f in all_files if f.endswith('.go')])

  logging.info("checking .py files")
  py_files_filter = lambda f: f.endswith('.py') and not f.endswith('_pb2.py')
  py = _FormatPythonFiles(args, [f for f in all_files if py_files_filter(f)])

  if args.check:
    modified_files = [
        os.path.relpath(f, SOURCE_PATH) for f in (pr + md + go + py)
    ]
    if len(modified_files) == 0:
      return

    print("The following files need reformatting. " +
          "Use 'python build.py format' to fix:\n")

    for f in sorted(modified_files):
      print(f)
    sys.exit(1)


def CheckFormatting(files):
  '''\
CheckFormatting returns a list of files within our source tree that are
incorrectly formatted.

This function is used by our PRESUBMIT.py script to block commits of
incorrectly formatted code.
'''

  class fakeargs(object):

    def __init__(self):
      self.check = True
      self.verbose = False
      self.goos = ''
      self.goarch = ''

  args = fakeargs()
  pr = _FormatProtoFiles(args, [f for f in files if f.endswith('.proto')])
  md = _FormatMarkdownFiles(args, [f for f in files if f.endswith('.md')])
  go = _FormatGoFiles(args, [f for f in files if f.endswith('.go')])
  py_files_filter = lambda f: f.endswith('.py') and not f.endswith('_pb2.py')
  py = _FormatPythonFiles(args, [f for f in files if py_files_filter(f)])
  modified_files = [
      os.path.relpath(f, SOURCE_PATH) for f in (pr + md + go + py)
  ]
  return modified_files


def CleanCommand(args):
  '''Remove build artifacts.'''

  force_option = ['-f' if args.force else '-n']
  _RunCommand(
      ['git', 'clean', '-X'] + force_option,
      env=_MergeEnv(args, target_host=True),
      cwd=SOURCE_PATH)

  if os.path.exists(OUT_PATH):
    if not args.force:
      print('Would remove {}/'.format(OUT_PATH))
      return

    print('Removing {}/'.format(OUT_PATH))
    shutil.rmtree(OUT_PATH)


def CreatePackageCommand(args):
  '''Create package.'''
  _RunCommand([
      'python3', 'setup.py', 'bdist_wheel', '--universal',
      '--dist-dir=%s' % OUT_PATH
  ],
              cwd=os.path.join(SOURCE_PATH, 'test'))


def main():
  common_options = argparse.ArgumentParser(add_help=False)
  common_options.add_argument(
      '--goos', '-O', help='set GOOS', choices=['windows', 'darwin', 'linux'])
  common_options.add_argument('--goarch', '-A', help='set GOARCH')
  common_options.add_argument(
      '--verbose', '-v', help='verbose output', action='store_true')

  parser = argparse.ArgumentParser(
      description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
  subparsers = parser.add_subparsers(help='Subcommands')

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
      '-f',
      action='store_true',
      help='''fast build. Skips dependency and generator steps''')
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
      '--fast',
      '-F',
      action='store_true',
      help='''fast build. Skips dependency and generator steps''')
  test_command.add_argument(
      '--coverage',
      '-c',
      action='store_true',
      help='''generate test coverage info''')
  test_command.add_argument(
      'gotest_args',
      metavar='ARGS',
      type=str,
      help='''aruments to pass down to "go test".
      Preface with "--" to disambiguate arguments passed in to
      this build tool.''',
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
      help='force. Without this option "clean" command doesn\'t do anything.')
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
      help='install additional dependencies')
  deps_command.add_argument(
      '--update', '-U', action='store_true', help='update dependencies')
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
  format_command.add_argument(
      '--check',
      '-n',
      action='store_true',
      help=("check if files are correctly formatted, but don\'t modify files "
            "on disk"))
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
      '--build_dir',
      '-b',
      action='store_true',
      help='resolve paths relative to build directory')
  run_command.add_argument(
      'prog', metavar='ARG', type=str, help='Program and arguments', nargs='+')
  run_command.set_defaults(closure=RunCommand)

  # ----------------------------------------------------------------------------
  # create package
  # ----------------------------------------------------------------------------
  create_package_command = subparsers.add_parser(
      'create_package',
      help=CreatePackageCommand.__doc__.splitlines()[0],
      description=CreatePackageCommand.__doc__,
      formatter_class=argparse.RawDescriptionHelpFormatter,
      parents=[common_options])
  create_package_command.set_defaults(closure=CreatePackageCommand)

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
