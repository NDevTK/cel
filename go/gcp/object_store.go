// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"bytes"
	"chromium.googlesource.com/enterprise/cel/go/common"
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	"cloud.google.com/go/storage"
	"context"
	"crypto/md5"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"strings"
)

type ObjectStore struct {
	context    context.Context
	client     *storage.Client
	bucket     *storage.BucketHandle
	projectID  string
	bucketName string
	prefix     string
}

const (
	DefaultPrefix                 = "cel-config"
	ContextIndexedObjectNamespace = "/o/"
)

func NewObjectStore(ctx context.Context, hc *http.Client, env *hostpb.HostEnvironment) (*ObjectStore, error) {
	if env.GetStorage() == nil {
		return nil, errors.New("host_environment does not specify storage parameters")
	}

	if env.GetStorage().GetBucket() == "" {
		return nil, errors.New("host_environment.storage.bucket is required")
	}

	client, err := storage.NewClient(ctx, option.WithHTTPClient(hc))
	if err != nil {
		return nil, err
	}

	prefix := env.GetStorage().GetPrefix()
	if prefix == "" {
		prefix = DefaultPrefix
	}
	prefix += ContextIndexedObjectNamespace

	return &ObjectStore{
		context:    ctx,
		client:     client,
		bucket:     client.Bucket(env.GetStorage().GetBucket()),
		projectID:  env.GetProject().GetName(),
		bucketName: env.GetStorage().GetBucket(),
		prefix:     prefix}, nil
}

func (o *ObjectStore) Close() error {
	return o.client.Close()
}

func (o *ObjectStore) Create() error {
	// TODO(asanka): Create should set reasonable retention policies.
	return o.bucket.Create(o.context, o.projectID, nil)
}

func (o *ObjectStore) PutObject(payload []byte) (string, error) {
	return o.PutNamedObject("", payload)
}

func (o *ObjectStore) PutNamedObject(baseName string, payload []byte) (ref string, err error) {
	objName := o.objectName(baseName, payload)
	h := o.bucket.Object(objName)

	// If the object is new, an error of type ErrObjectNotExist is expected.
	attrs, err := h.Attrs(o.context)
	if err != nil && err != storage.ErrObjectNotExist {
		return "", errors.Wrapf(err, "checking if object named \"%s\" exists in GCS bucket \"%s\"",
			objName, o.bucketName)
	}

	sum := md5.Sum(payload)

	// The object in storage is the same as the one we are about to upload, at
	// least as far as MD5 can tell apart. The consumer end of the API uses
	// subresource integrity with a better digest algorithm to make sure the
	// contents are what we expected them to be.
	if attrs != nil && bytes.Equal(sum[:], attrs.MD5) && attrs.Size == int64(len(payload)) {
		return objName, nil
	}

	// If the MD5 sum doesn't match, we are going to assume that a prior
	// write operation failed. Since we have the data, let's try to write
	// the object again.

	// TODO(asanka): Validate that ACLs are sane. We should refuse to use
	// an object that has a sketchy ACL.

	ctx, cancelFunc := context.WithCancel(o.context)
	defer func() {
		if err != nil {
			cancelFunc()
		}
	}()

	w := h.NewWriter(ctx)
	w.MD5 = sum[:]

	// TODO(asanka): Set object attributes. Specifically ACLs. These should be
	// based on the service accounts extracted from the HostEnvironment.

	_, err = w.Write(payload)
	if err != nil {
		return "", errors.Wrapf(err, "writing to object named \"%s\" in GCS bucket \"%s\"",
			objName, o.bucketName)
	}

	err = w.Close()
	if err != nil {
		return "", errors.Wrapf(err, "closing object named \"%s\" in GCS bucket \"%s\"", objName,
			o.bucketName)
	}

	return objName, nil
}

func (o *ObjectStore) GetObject(reference string) (name string, payload []byte, err error) {
	h := o.bucket.Object(reference)
	_, err = h.Attrs(o.context)
	if err != nil {
		return
	}
	name, sri, err := o.crackObjectName(reference)
	if err != nil {
		return
	}

	ctx, cancelFunc := context.WithCancel(o.context)
	defer func() {
		if err != nil && cancelFunc != nil {
			cancelFunc()
		}
	}()

	r, err := h.NewReader(ctx)
	if err != nil {
		return
	}

	payload, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	r.Close()

	err = common.ValidateIntegrity(payload, sri)
	cancelFunc = nil
	return
}

func (o *ObjectStore) objectName(baseName string, payload []byte) string {
	sri := common.IntegrityLabel(payload)
	suffix := ""
	if baseName != "" {
		suffix = "/" + baseName
	}
	return o.prefix + sri + suffix
}

func (o *ObjectStore) crackObjectName(name string) (suffix, sri string, err error) {
	if !strings.HasPrefix(name, o.prefix) {
		err = errors.Errorf("unexpected object prefix. Name \"%s\". Expected to start with \"%s\"",
			name, o.prefix)
		return
	}

	name = name[len(o.prefix):]
	components := strings.Split(name, "/")
	switch len(components) {
	case 2:
		suffix = components[1]
		fallthrough
	case 1:
		sri = components[0]
	default:
		err = errors.Errorf("invalid object name format: \"%s\"", name)
	}
	return
}

// AbsoluteReference returns the full GCS path based on a bucket name and an
// objectReference. The function assumes that both inputs are valud.
func AbsoluteReference(bucketName, objectReference string) string {
	return "gs://" + bucketName + "/" + objectReference
}
