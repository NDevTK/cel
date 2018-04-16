// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
)

// Resolvers

type ProjectResolver struct{}

func (ProjectResolver) ResolveImmediate(ctx common.Context, p *host.Project) (err error) {
	session, err := SessionFromContext(ctx)
	if err != nil {
		return err
	}

	defer GcpLoggedServiceAction(session, cloudResourceManagerServiceName, &err,
		"Resolving metadata for Project \"%s\"", p.GetName())()

	svc, err := session.GetCloudResourceManagerService()
	if err != nil {
		return err
	}

	cp, err := svc.Projects.Get(p.GetName()).Context(ctx).Do()

	if err != nil {
		return err
	}

	return ctx.Publish(p, "project_number", cp.ProjectNumber)
}

func init() {
	common.RegisterResolverClass(&ProjectResolver{})
}
