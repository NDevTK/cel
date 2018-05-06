// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	computepb "chromium.googlesource.com/enterprise/cel/go/gcp/compute"
	"chromium.googlesource.com/enterprise/cel/go/host"
)

// Resolvers

type ProjectResolver struct{}

func (*ProjectResolver) ResolveImmediate(ctx common.Context, p *host.Project) (err error) {
	session, err := gcp.SessionFromContext(ctx)
	if err != nil {
		return err
	}

	defer gcp.GcpLoggedServiceAction(session, gcp.ComputeServiceName, &err,
		"Resolving metadata for Project \"%s\"", p.GetName())()

	svc, err := session.GetComputeService()
	if err != nil {
		return err
	}

	proj, err := svc.Projects.Get(p.GetName()).Context(ctx).Do()
	if err != nil {
		return err
	}

	var pbProj *computepb.Project
	err = common.HomomorphicCopy(&proj, &pbProj)

	return ctx.Publish(p, "project", pbProj)
}

func init() {
	common.RegisterResolverClass(&ProjectResolver{})
}
