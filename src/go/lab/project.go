// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
)

type Project struct {
	BaseNamedAsset
	Project *cloudresourcemanager.Project
}

func (p *Project) PermanentAsset() {}

func (p *Project) Check(s *Session) (err error) {
	defer Action(&err, "querying project %s", p.id)
	p.Project, err = s.GetCloudResourceManagerService().Projects.Get(p.id).Context(s.Context).Do()
	return
}

func (p *Project) ResourcePath() string {
	return "projects/" + p.id
}

func LookupProject(A *Assets, id string) *Project {
	const kNamespace = "projects"
	if a := A.Get(kNamespace, id); a != nil {
		return a.(*Project)
	}
	panic("project not found")
}

func ConstructProjectAsset(A *Assets, c *Config) error {
	const kNamespace = "projects"
	p := Project{BaseNamedAsset{kNamespace, c.Project, nil}, nil}
	return A.Add(&p)
}
