#!/usr/bin/env python

# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
'''
Installs all dependencies required to run test.py.
'''

import subprocess

# TODO: Add Windows support
subprocess.check_call(['apt-get', 'install', 'python-pip'])
subprocess.check_call(['pip', 'install', 'google-api-python-client'])
subprocess.check_call(['pip', 'install', 'grpc-google-iam-admin-v1'])
subprocess.check_call(['pip', 'install', 'grpc-google-iam-v1'])
