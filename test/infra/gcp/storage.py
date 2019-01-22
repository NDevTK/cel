# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import api_client
import googleapiclient
import StringIO


class StorageBucket:
  """Provides basic GCP Storage Objects functions (write, delete and list)."""

  def __init__(self, name):
    self.name = name

  def DeleteObject(self, objectName, expectedGeneration):
    """Deletes a GCP Storage Object.

    Args:
      objectName: The full path to an object to delete
      expectedGeneration: The expected generation id of the object.
        This ensures that we only delete the object we expect and prevents
        deleting an object that has been replaced since we last queried it.
    """
    service = api_client.build("storage", 'v1')

    req = service.objects().delete(
        bucket=self.name,
        object=objectName,
        ifGenerationMatch=expectedGeneration)

    resp = req.execute()

    return resp

  def WriteObject(self, objectName, content, expectedGeneration):
    """Writes or create a GCP Storage Object.

    Args:
      objectName: The full path to an object to delete
      content: The content of the storage object.
      expectedGeneration: The expected generation id of the object.
        Should be 0 if we don't expect that object to exist.
        This ensures that we only write the object we expect and prevents
        writing over object that has been written/replaced already.

    Returns:
      The generation id of the new object. The generation id is used when
      calling DeleteObject to ensure we're deleting the object we created
      and not an object created by someone else.
    """
    service = api_client.build("storage", 'v1')

    content = StringIO.StringIO()
    content.write(content)

    req = service.objects().insert(
        bucket=self.name,
        body={'name': objectName},
        ifGenerationMatch=expectedGeneration,
        media_body=googleapiclient.http.MediaIoBaseUpload(
            content, 'application/octet-stream'))

    resp = req.execute()

    return resp['generation']

  def ListObjects(self, prefix):
    """Lists GCP Storage Object with a given prefix

    Returns:
      A list of GCP Storage Objects as described here:
      https://cloud.google.com/storage/docs/json_api/v1/objects#resource
    """
    service = api_client.build("storage", 'v1')

    req = service.objects().list(bucket=self.name, prefix=prefix)

    resp = req.execute()

    if 'items' in resp:
      return resp['items']

    return []
