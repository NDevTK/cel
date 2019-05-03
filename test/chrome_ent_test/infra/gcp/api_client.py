# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import googleapiclient.discovery
import googleapiclient.http
import httplib2
import os
import subprocess


def build(api, version):
  # We can't use application default credentials on LUCI bots.
  if "LUCI_CONTEXT" in os.environ:
    return googleapiclient.discovery.build(
        api, version, http=httplib2.Http(), requestBuilder=LuciHttpRequest)

  # This will use default credentials for local runs.
  return googleapiclient.discovery.build(api, version)


class LuciHttpRequest(googleapiclient.http.HttpRequest):
  """This builds an HttpRequest that uses tokens from `luci-auth`."""

  def __init__(self, http, postproc, uri, **kwargs):
    if kwargs['headers'] == None:
      kwargs['headers'] = {}

    token = subprocess.check_output([
        'luci-auth', 'token',
        '--scopes=https://www.googleapis.com/auth/cloud-platform'
    ])

    kwargs['headers']["Authorization"] = "Bearer %s" % token
    super(LuciHttpRequest, self).__init__(http, postproc, uri, **kwargs)
