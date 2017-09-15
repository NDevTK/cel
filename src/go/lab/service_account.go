// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"go/lab/config"
	iam "google.golang.org/api/iam/v1"
)

type ServiceAccount struct {
	BaseNamedAsset

	Config  *config.ServiceAccount
	Project *Project

	IamServiceAccount *iam.ServiceAccount
}

func (v *ServiceAccount) Resolve(s *Session) (err error) {
	a, ok := s.Config.Cloud.ServiceAccounts[v.Config.Id]
	if ok {
		v.IamServiceAccount = a
		return
	}

	defer Action(&err, "creating service account %s", v.Id())
	v.IamServiceAccount, err = s.GetIamService().Projects.ServiceAccounts.Create(v.Project.ResourcePath(),
		&iam.CreateServiceAccountRequest{
			AccountId: v.Config.Id,
			ServiceAccount: &iam.ServiceAccount{
				DisplayName: v.Config.DisplayName}}).Context(s.Context).Do()
	if err != nil {
		return
	}
	s.Config.Cloud.SetServiceAccountsChanged()
	return
}

func (v *ServiceAccount) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (v *ServiceAccount) Purge(s *Session) error {
	panic("not implemented")
}

func LookupServiceAccount(A *Assets, id string) *ServiceAccount {
	const kNamespace = "serviceAccounts"
	a := A.Get(kNamespace, id)
	if a == nil {
		panic("instance service account not found")
	}
	return a.(*ServiceAccount)
}

func ConstructServiceAccountAssets(A *Assets, c *Config) error {
	const kNamespace = "serviceAccounts"
	p := LookupProject(A, c.Project)
	for _, cs := range c.ServiceAccount {
		s := ServiceAccount{BaseNamedAsset{kNamespace, cs.Id, []Asset{p}}, cs, p, nil}
		err := A.Add(&s)
		if err != nil {
			return err
		}
	}
	return nil
}
