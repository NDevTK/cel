# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from absl import app
from absl import flags
import tempfile
import logging
import shutil
import subprocess
import os
import re
import zipfile

FLAGS = flags.FLAGS

flags.DEFINE_string('input_file', None, 'The path of the celab build')

flags.mark_flag_as_required('input_file')

flags.DEFINE_enum('platform', None, ['windows-amd64', 'linux-amd64'],
                  'The platform of the build')
flags.mark_flag_as_required('platform')


def _runCommand(args, **kwargs):
  logging.info('Run command: %s', ' '.join(args))
  return subprocess.check_output(args, **kwargs)


def main(argv):
  package = "infra/celab/celab/{}".format(FLAGS.platform)

  yaml_template = """package: {}
description: CELab release
install_mode: copy
root: {}
data:
  - dir: .
"""

  temp_dir = tempfile.mkdtemp()
  bin_dir = os.path.join(temp_dir, "binaries")
  try:
    _runCommand(['gsutil', 'cp', FLAGS.input_file, temp_dir])
    filename = os.path.basename(FLAGS.input_file)
    zip_file = zipfile.ZipFile(os.path.join(temp_dir, filename), 'r')
    zip_file.extractall(bin_dir)

    yaml_file = tempfile.NamedTemporaryFile(delete=False)
    yaml_file.write(yaml_template.format(package, bin_dir))
    yaml_file.close()
    try:
      # run command 'create' to create & upload the package.
      output = _runCommand(['cipd', 'create', '-pkg-def=' + yaml_file.name])

      # get the instance id from the output.
      # the pattern we're looking for looks like this:
      # infra/celab/celab/windows-amd64:eMXP1ODJ6X2xxxx
      pattern = re.compile(r'{}:([^\s]+)'.format(package))
      m = pattern.search(output)

      if m is None:
        raise "Cannot find the pattern"

      # set ref "latest" to be the one just uploaded.
      version = m.group(1)
      _runCommand(
          ['cipd', 'set-ref', package, '-ref', 'latest', '-version', version])

    finally:
      os.remove(yaml_file.name)
  finally:
    shutil.rmtree(temp_dir)


if __name__ == '__main__':
  app.run(main)