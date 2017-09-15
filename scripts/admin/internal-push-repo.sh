#!/bin/bash

# This script is used to populate the GCS bucket containing the Chrome
# Enterprise Lab tools that need to be deployed to the lab VMs.

set -e
pushd .

cd "$(dirname "${BASH_SOURCE[0]}" )"/../..

./build.py build
gsutil -m rsync -r -d -x \(\\.git.\*\|src.\*\|pkg.\*) . gs://chrome-auth-lab-staging/tools

popd

