// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"chromium.googlesource.com/enterprise/cel/go/gcp/compute"
	"fmt"
	"github.com/pkg/errors"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
)

const celDeploymentName = "cel-lab"
const generationIdKey = "generation-id"

func DeleteObsoleteDeployments(ctx common.Context, s *gcp.Session) error {
	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return err
	}

	_, err = dm.Deployments.Get(s.GetProject(), celDeploymentName).Context(ctx).Do()
	if err != nil {
		if gcp.IsNotFoundError(err) {
			// nothing to delete
			return nil
		}
		return err
	}

	err = cleanupFirewallRules(ctx, s)
	if err != nil {
		return err
	}

	o, err := dm.Deployments.Delete(s.GetProject(), celDeploymentName).Context(ctx).Do()
	if err != nil {
		return err
	}

	return gcp.JoinOperation(s, o, "Deleting obsolete deployment")
}

func InvokeDeploymentManager(ctx common.Context, s *gcp.Session) (err error) {
	defer func() {
		if err != nil && gcp.IsDuplicateResourceError(err) {
			err = errors.Wrapf(err, `failed to deploy lab due to conflicting resources.

This error may be due to there being an obsolete deployment in the project.
While the CEL toolchain attempts to remove such deployments, such attempts may
fail silently under some circumstances. This is not ideal and we are working on
addressing this. In the meantime, here's what you can try:

* Look at the error listed below and manually remove the conflicting resource
  from the project.

* Let us know the circumstances under which you ran into this issue. If the
  offending resource was one that was deployed by the CEL toolchain, then the
  toolchain should do a better job of cleaning up. Please help us fix this.

  It is possible that the offending resource was placed in the project by
  something external. We also try to address these so that the toolchain can
  co-exist and deal nicely with such additions. This is specially true for
  changes made to corporate owned GCP projects. Let us know about these as
  well.

Apologies and thank you for your understanding.

Underlying error:`)
		}
	}()

	m := GetDeploymentManifest()

	y, err := m.GetYaml()
	if err != nil {
		return err
	}

	d := deploymentmanager.Deployment{
		Name:        celDeploymentName,
		Description: "CEL Lab",
		Labels: []*deploymentmanager.DeploymentLabelEntry{
			{
				Key:   generationIdKey,
				Value: common.Must(ctx.Get(common.RefPathMust("lab.generation_id"))).(string),
			},
		},
		Target: &deploymentmanager.TargetConfiguration{
			Config: &deploymentmanager.ConfigFile{
				Content: string(y),
			},
		},
	}

	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return err
	}

	o, err := dm.Deployments.Insert(s.GetProject(), &d).Context(ctx).Do()
	if err != nil {
		return err
	}

	return gcp.JoinOperation(s, o, "Deploying lab")
}

func cleanupFirewallRules(ctx common.Context, s *gcp.Session) error {
	networks := make(map[string]bool)
	ntype, _ := getResourceType(&compute.Network{})

	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return err
	}

	for token := ""; ; {
		rr := dm.Resources.List(s.GetProject(), celDeploymentName).
			Context(ctx).Filter(fmt.Sprintf("type eq %s", ntype))
		if token != "" {
			rr = rr.PageToken(token)
		}
		rl, err := rr.Do()
		if err != nil {
			break
		}

		for _, r := range rl.Resources {
			networks[gcp.PartialUrl(r.Url)] = true
		}

		if rl.NextPageToken == "" || len(rl.Resources) == 0 {
			break
		}
		token = rl.NextPageToken
	}

	if len(networks) == 0 {
		return nil
	}

	cs, err := s.GetComputeService()
	if err != nil {
		return err
	}

	var firewalls []string

	for token := ""; ; {
		fr := cs.Firewalls.List(s.GetProject()).Context(ctx)
		if token != "" {
			fr = fr.PageToken(token)
		}
		fl, err := fr.Do()
		if err != nil {
			return err
		}
		for _, f := range fl.Items {
			pu := gcp.PartialUrl(f.Network)
			if _, ok := networks[pu]; ok {
				firewalls = append(firewalls, f.Name)
			}
		}
		if fl.NextPageToken == "" {
			break
		}
		token = fl.NextPageToken
	}

	t := common.NewTasks(nil)
	for _, f := range firewalls {
		ff := f
		t.Go(func() error {
			o, err := cs.Firewalls.Delete(s.GetProject(), ff).Context(ctx).Do()
			if err != nil {
				return err
			}
			return gcp.JoinOperation(s, o, fmt.Sprintf("Deleting firewall %s", ff))
		})
	}
	return t.Join()
}

func GetCurrentDeployment(ctx common.Context, s *gcp.Session) (d *deploymentmanager.Deployment, err error) {
	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return nil, err
	}

	return dm.Deployments.Get(s.GetProject(), celDeploymentName).Context(ctx).Do()
}

func GetManifest(ctx common.Context, s *gcp.Session, mLink string) (m *deploymentmanager.Manifest, err error) {
	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return nil, err
	}
	return dm.Manifests.Get(s.GetProject(), celDeploymentName, gcp.LastPathComponent(mLink)).Context(ctx).Do()
}

func GetResources(ctx common.Context, s *gcp.Session) (r []*deploymentmanager.Resource, err error) {
	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return
	}

	for token := ""; ; {
		rr := dm.Resources.List(s.GetProject(), celDeploymentName).Context(ctx)
		if token != "" {
			rr = rr.PageToken(token)
		}
		rl, err := rr.Do()
		if err != nil {
			break
		}

		r = append(r, rl.Resources...)
		if rl.NextPageToken == "" || len(rl.Resources) == 0 {
			break
		}
		token = rl.NextPageToken
	}

	return
}
