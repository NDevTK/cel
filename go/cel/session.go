// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
	cloudkms "google.golang.org/api/cloudkms/v1"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	compute "google.golang.org/api/compute/v1"
	iam "google.golang.org/api/iam/v1"
	"net/http"
	"sync"
)

type Session struct {
	Context context.Context
	Client  *http.Client
	Cloud   *CloudState

	compute_service *compute.Service
	compute_once    sync.Once

	cloudkms_service      *cloudkms.Service
	cloudkms_service_once sync.Once

	iam_service      *iam.Service
	iam_service_once sync.Once

	cloudresourcemanager_service *cloudresourcemanager.Service
	cloudresourcemanager_once    sync.Once

	// Logging. Use one of the Log* methods.
	log_client *logging.Client
	logger     *logging.Logger
}

func NewSession(ctx context.Context, client *http.Client, host_config string, asset_configs []string) (s *Session, err error) {
	s = &Session{Context: ctx, Client: client}

	s.Config, err = LoadConfigFiles(host_config, asset_configs)
	if err != nil {
		return
	}

	s.log_client, err = logging.NewClient(s.Context, s.Config.Project)
	if err != nil {
		return
	}
	s.logger = s.log_client.Logger(kLabAdminLogName)

	s.Cloud, err = QueryCloudState(s.Context, s.Client, s.Config.Project, s.Config.images)
	return
}

func (s *Session) GetComputeService() *compute.Service {
	s.compute_once.Do(func() {
		var err error
		s.compute_service, err = compute.New(s.Client)
		if err != nil {
			panic(err)
		}
	})
	return s.compute_service
}

func (s *Session) GetCloudKmsService() *cloudkms.Service {
	s.cloudkms_service_once.Do(func() {
		var err error
		s.cloudkms_service, err = cloudkms.New(s.Client)
		if err != nil {
			panic(err)
		}
	})
	return s.cloudkms_service
}

func (s *Session) GetIamService() *iam.Service {
	s.iam_service_once.Do(func() {
		var err error
		s.iam_service, err = iam.New(s.Client)
		if err != nil {
			panic(err)
		}
	})
	return s.iam_service
}

func (s *Session) GetCloudResourceManagerService() *cloudresourcemanager.Service {
	s.cloudresourcemanager_once.Do(func() {
		var err error
		s.cloudresourcemanager_service, err = cloudresourcemanager.New(s.Client)
		if err != nil {
			panic(err)
		}
	})
	return s.cloudresourcemanager_service
}

func (s *Session) LogInfo(e LogEntrySource) {
	s.logger.Log(e.Entry(logging.Info))
}

func (s *Session) LogError(e LogEntrySource) {
	s.logger.Log(e.Entry(logging.Error))
}
