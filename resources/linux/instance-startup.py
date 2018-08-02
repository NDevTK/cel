#! /usr/bin/python

# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
import requests
import json
import subprocess
import os
import sys


def get_metadata(name):
  """Gets the metadata."""
  r = requests.get(
      'http://metadata.google.internal/computeMetadata/v1/project/attributes/' +
      name,
      headers={'Metadata-Flavor': 'Google'})
  return r.text


cel_agent = json.loads(get_metadata('cel-agent'))
print 'cel_agent is {0}'.format(cel_agent)
cel_agent_path = cel_agent['linux_agent_x64']['abs_path']

cel_manifest = get_metadata('cel-manifest')
print 'cel_manifest is {0}'.format(cel_manifest)
sys.stdout.flush()

cwd = os.getcwd()
dir = os.path.join(cwd, 'cel')
subprocess.call(['mkdir', '-p', dir])
subprocess.call(['gsutil', 'cp', cel_agent_path, 'cel_agent'], cwd=dir)
subprocess.call(['chmod', 'a+x', 'cel_agent'], cwd=dir)
subprocess.call(['gsutil', 'cp', cel_manifest, 'cel_manifest.textpb'], cwd=dir)
subprocess.call(
    [os.path.join(dir, 'cel_agent'),
     os.path.join(dir, 'cel_manifest.textpb')])
