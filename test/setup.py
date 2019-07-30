# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import setuptools

setuptools.setup(
    name="chrome-ent-test",
    version="0.1",
    description="Chrome enterprise test framework",
    packages=setuptools.find_packages(exclude=["tests"]),
    install_requires=[
        "absl-py", "google-api-python-client", "protobuf",
        "grpc-google-iam-admin-v1", "grpc-google-iam-v1"
    ],
)
