// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type DeployerSession struct {
	ctx     common.Context
	client  *http.Client
	config  *Configuration
	backend *gcp.Session
}

func NewDeployerSession(ctx context.Context, client *http.Client, inputs []string) (*DeployerSession, error) {
	gen, err := createGenerationId()
	if err != nil {
		return nil, err
	}

	c := &Configuration{}
	for _, f := range inputs {
		err := c.Merge(f)
		if err != nil {
			return nil, err
		}
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	b, err := gcp.NewSession(ctx, client, &c.HostEnvironment, gen)
	if err != nil {
		return nil, err
	}

	o, err := gcp.NewObjectStore(b.GetContext(), client, &c.HostEnvironment)
	if err != nil {
		return nil, err
	}

	err = c.references.Publish(&c.Lab, "generation_id", gen)
	if err != nil {
		return nil, err
	}

	l, err := b.GetLogger()
	if err != nil {
		return nil, err
	}

	return &DeployerSession{ctx: &deployerContext{
		ctx:         b.GetContext(),
		objectStore: o,
		publisher:   c.GetNamespace(),
		logger:      l,
	}, client: client, config: c, backend: b}, nil
}

func (d *DeployerSession) GetContext() common.Context {
	return d.ctx
}

func (d *DeployerSession) GetConfiguration() *Configuration {
	return d.config
}

func (d *DeployerSession) GetBackend() *gcp.Session {
	return d.backend
}

// createGenerationId creates a random generation ID. It doesn't take history
// or content of the configuration into account. Instead we create a large
// enough random number that it's unlikely to collide with a previously used
// ID.
func createGenerationId() (string, error) {
	r := make([]byte, 16) // 16 bytes == 128 bits
	_, err := rand.Read(r)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(r), nil
}
