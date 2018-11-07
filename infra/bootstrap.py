#!/usr/bin/env python
# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Go installation from go/bootstrap.py in Chromium infra's repository.
#
# https://chromium.googlesource.com/infra/infra/
"""Prepares a local hermetic Go & Protoc installation.

- Downloads and unpacks the Go toolset.
- Downloads and unpacks the Protoc library.
"""

import contextlib
import logging
import os
import platform
import shutil
import stat
import subprocess
import sys
import tarfile
import tempfile
import urllib
import zipfile

LOGGER = logging.getLogger(__name__)

# Platform depended suffix for executable files.
EXE_SFX = '.exe' if sys.platform == 'win32' else ''

# Pinned version of Go toolset and Protoc to download.
GO_TOOLSET_VERSION = 'go1.11'
PROTOC_VERSION = 'protoc-3.6.1'

# Platform dependent portion of a download URL.
# See http://golang.org/dl/ & https://github.com/google/protobuf/releases/.
GO_TOOLSET_VARIANTS = {
    ('linux2', 'x86-64'): '%s.linux-amd64.tar.gz' % GO_TOOLSET_VERSION,
    ('win32', 'x86-64'): '%s.windows-amd64.zip' % GO_TOOLSET_VERSION,
}
PROTOC_VARIANTS = {
    ('linux2', 'x86-64'): '%s-linux-x86_64.zip' % PROTOC_VERSION,
    ('win32', 'x86-64'): '%s-win32.zip' % PROTOC_VERSION,
}

# Download URL roots.
GO_DOWNLOAD_URL_PREFIX = 'https://storage.googleapis.com/golang'
PB_DOWNLOAD_URL_PREFIX = 'https://github.com/google/protobuf/releases/download/v3.6.1'


class Failure(Exception):
  """Bootstrap failed."""


def get_variant_url(variants, url_prefix):
  """URL of a platform specific archive."""
  arch = {
      'amd64': 'x86-64',
      'x86_64': 'x86-64',
  }.get(platform.machine().lower())
  variant = variants.get((sys.platform, arch))
  if not variant:
    raise Failure('Unrecognized platform')
  return '%s/%s' % (url_prefix, variant)


def get_go_toolset_url():
  """URL of a platform specific Go toolset archive."""
  return get_variant_url(GO_TOOLSET_VARIANTS, GO_DOWNLOAD_URL_PREFIX)


def get_protoc_url():
  """URL of a platform specific protoc archive."""
  return get_variant_url(PROTOC_VARIANTS, PB_DOWNLOAD_URL_PREFIX)


def read_file(path):
  """Returns contents of a given file or None if not readable."""
  assert isinstance(path, (list, tuple))
  try:
    with open(os.path.join(*path), 'r') as f:
      return f.read()
  except IOError:
    return None


def write_file(path, data):
  """Writes |data| to a file."""
  assert isinstance(path, (list, tuple))
  with open(os.path.join(*path), 'w') as f:
    f.write(data)


def remove_directory(path):
  """Recursively removes a directory."""
  assert isinstance(path, (list, tuple))
  p = os.path.join(*path)
  if not os.path.exists(p):
    return
  LOGGER.info('Removing %s', p)

  # Crutch to remove read-only file (.git/* in particular) on Windows.
  def onerror(func, path, _exc_info):
    if not os.access(path, os.W_OK):
      os.chmod(path, stat.S_IWUSR)
      func(path)
    else:
      raise

  shutil.rmtree(p, onerror=onerror if sys.platform == 'win32' else None)


def install_package(install_root, url):
  """Downloads and installs a package."""
  if not os.path.exists(install_root):
    os.makedirs(install_root)
  pkg_path = os.path.join(install_root, url[url.rfind('/') + 1:])

  LOGGER.info('Downloading %s...', url)
  download_file(url, pkg_path)

  LOGGER.info('Extracting...')
  if pkg_path.endswith('.zip'):
    with zipfile.ZipFile(pkg_path, 'r') as f:
      f.extractall(install_root)
  elif pkg_path.endswith('.tar.gz'):
    with tarfile.open(pkg_path, 'r:gz') as f:
      f.extractall(install_root)
  else:
    raise Failure('Unrecognized archive format')


def download_file(url, path):
  """Fetches |url| to |path|."""
  last_progress = [0]

  def report(a, b, c):
    progress = int(a * b * 100.0 / c)
    if progress != last_progress[0]:
      print >> sys.stderr, 'Downloading... %d%%' % progress
      last_progress[0] = progress

  # TODO(vadimsh): Use something less crippled, something that validates SSL.
  urllib.urlretrieve(url, path, reporthook=report)


def ensure_installed(name, install_root, url, version):
  """Installs or updates a package if necessary.

  Returns True if a new package was installed.
  """
  installed = read_file([install_root, 'INSTALLED'])
  if installed == url:
    LOGGER.debug('%s is up-to-date: %s', name, version)
    return False
  LOGGER.info('Installing %s.', name)
  LOGGER.info('  Old package is %s', installed)
  LOGGER.info('  New package is %s', url)
  remove_directory([install_root])
  install_package(install_root, url)
  LOGGER.info('%s installed: %s', name, version)
  write_file([install_root, 'INSTALLED'], url)
  return True


@contextlib.contextmanager
def temp_dir(path):
  """Creates a temporary directory, then deletes it."""
  tmp = tempfile.mkdtemp(dir=path)
  try:
    yield tmp
  finally:
    remove_directory([tmp])


def check_hello_world(toolset_root):
  """Compiles and runs 'hello world' program to verify that toolset works."""
  with temp_dir(toolset_root) as tmp:
    path = os.path.join(tmp, 'hello.go')
    write_file([path], r"""
        package main
        func main() { println("hello, world\n") }
    """)

    go_exe = os.path.join(toolset_root, 'go', 'bin', 'go' + EXE_SFX)

    out = subprocess.check_output([go_exe, 'run', path],
                                  env=get_go_environ(toolset_root, tmp),
                                  stderr=subprocess.STDOUT)
    if out.strip() != 'hello, world':
      LOGGER.error('Failed to run sample program:\n%s', out)
      return False
    return True


def get_go_environ(toolset_root, workspace):
  """Returns a copy of os.environ with added GO* environment variables.

  Overrides GOROOT, GOPATH and GOBIN. Keeps everything else. Idempotent.

  Args:
    toolset_root: GOROOT would be <toolset_root>/go.
    workspace: main workspace directory or None if compiling in GOROOT.
  """
  env = os.environ.copy()
  env['GOROOT'] = os.path.join(toolset_root, 'go')
  if workspace:
    env['GOBIN'] = os.path.join(workspace, 'bin')
  else:
    env.pop('GOBIN', None)

  all_go_paths = []
  if workspace:
    all_go_paths.append(workspace)
  env['GOPATH'] = os.pathsep.join(all_go_paths)

  # New PATH entries.
  paths_to_add = [
      os.path.join(env['GOROOT'], 'bin'),
      env.get('GOBIN'),
  ]

  # Make sure not to add duplicates entries to PATH over and over again when
  # get_go_environ is invoked multiple times.
  path = env['PATH'].split(os.pathsep)
  paths_to_add = [p for p in paths_to_add if p and p not in path]
  env['PATH'] = os.pathsep.join(paths_to_add + path)

  return env


def main(argv):
  if len(argv) != 2:
    return -1

  LOGGER.info('Installing go...')
  go_dir = os.path.join(argv[1], "golang")
  url = get_go_toolset_url()
  ensure_installed("Go toolset", go_dir, url, GO_TOOLSET_VERSION)

  LOGGER.info('Validating go installation...')
  if not check_hello_world(go_dir):
    raise Failure('Something is not right, test program doesn\'t work')

  LOGGER.info('Installing protoc...')
  protoc_dir = os.path.join(argv[1], "protoc")
  url = get_protoc_url()
  ensure_installed("Protoc", protoc_dir, url, PROTOC_VERSION)

  # Ensure that we can execute the protoc binary
  if sys.platform == 'linux2':
    protoc_bin = os.path.join(protoc_dir, "bin", "protoc")
    mode = os.stat(protoc_bin).st_mode
    os.chmod(protoc_bin, mode | stat.S_IEXEC)

  return 0


if __name__ == '__main__':
  sys.exit(main(sys.argv))
