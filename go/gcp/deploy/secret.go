// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"encoding/base64"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

type secret struct{}

func (*secret) ResolveIndexedObjects(ctx common.Context, secret *commonpb.Secret) error {
	s, err := gcp.SessionFromContext(ctx)
	if err != nil {
		return err
	}

	ks, err := s.GetCloudKmsService()
	if err != nil {
		return err
	}

	keyName := gcp.CryptoKeyResource(s.GetProject(), gcp.KmsLocationId, gcp.KmsKeyRingId,
		gcp.KmsCryptoKeyId)

	resp, err := ks.Projects.Locations.KeyRings.CryptoKeys.
		Encrypt(keyName, &cloudkms.EncryptRequest{
			Plaintext: string(base64.StdEncoding.EncodeToString(secret.Final)),
		}).
		Context(ctx).
		Do()
	if err != nil {
		return err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(resp.Ciphertext)
	if err != nil {
		return err
	}

	ref, err := ctx.GetObjectStore().PutObject(ciphertext)
	if err != nil {
		return err
	}

	return ctx.Publish(secret, "object_reference", ref)
}

func init() {
	common.RegisterResolverClass(&secret{})
}
