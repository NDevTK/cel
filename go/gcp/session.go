// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
	iam "cloud.google.com/go/iam/admin/apiv1"
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"github.com/pkg/errors"
	cloudkms "google.golang.org/api/cloudkms/v1"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	compute "google.golang.org/api/compute/v1"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
	"log"
	"net/http"
	"sync"
)

type Session struct {
	ctx             context.Context
	client          *http.Client
	HostEnvironment *host.HostEnvironment

	computeService *compute.Service
	computeOnce    sync.Once
	computeResult  error

	cloudkmsService     *cloudkms.Service
	cloudkmsServiceOnce sync.Once
	cloudkmsResult      error

	iamClient     *iam.IamClient
	iamClientOnce sync.Once
	iamResult     error

	cloudresourcemanagerService *cloudresourcemanager.Service
	cloudresourcemanagerOnce    sync.Once
	cloudresourcemanagerResult  error

	deploymentmanagerService *deploymentmanager.Service
	deploymentmanagerOnce    sync.Once
	deploymentmanagerResult  error

	logClient *logging.Client
	logger    *logging.Logger

	Logger common.Logger
}

type gcpSessionKeyType int

var gcpSessionKey = gcpSessionKeyType(1)

func NewSession(ctx context.Context, client *http.Client, env *host.HostEnvironment, genId string) (s *Session, err error) {
	s = &Session{ctx: nil, client: client, HostEnvironment: env}
	s.ctx = context.WithValue(ctx, gcpSessionKey, s)

	if env.LogSettings.GetAdminLog() != "" {
		s.logClient, err = logging.NewClient(s.ctx, env.Project.Name)
		if err != nil {
			return
		}
		s.logger = s.logClient.Logger(env.LogSettings.AdminLog)
	}
	s.Logger = &loggerAdapter{gcpLogger: s.logger, generationId: genId}

	return
}

func SessionFromContext(ctx context.Context) (*Session, error) {
	if s, ok := ctx.Value(gcpSessionKey).(*Session); ok {
		return s, nil
	}
	return nil, errors.New("context does not have a gcp Session associated with it")
}

func (s *Session) GetHttpClient() *http.Client {
	return s.client
}

func (s *Session) GetContext() context.Context {
	return s.ctx
}

func (s *Session) GetProject() string {
	return s.HostEnvironment.Project.Name
}

func (s *Session) GetComputeService() (*compute.Service, error) {
	s.computeOnce.Do(func() {
		s.computeService, s.computeResult = compute.New(s.client)
	})
	return s.computeService, s.computeResult
}

func (s *Session) GetCloudKmsService() (*cloudkms.Service, error) {
	s.cloudkmsServiceOnce.Do(func() {
		s.cloudkmsService, s.cloudkmsResult = cloudkms.New(s.client)
	})
	return s.cloudkmsService, s.cloudkmsResult
}

func (s *Session) GetIamClient() (*iam.IamClient, error) {
	s.iamClientOnce.Do(func() {
		s.iamClient, s.iamResult = iam.NewIamClient(s.ctx)
	})
	return s.iamClient, s.iamResult
}

func (s *Session) GetCloudResourceManagerService() (*cloudresourcemanager.Service, error) {
	s.cloudresourcemanagerOnce.Do(func() {
		s.cloudresourcemanagerService, s.cloudresourcemanagerResult = cloudresourcemanager.New(s.client)
	})
	return s.cloudresourcemanagerService, s.cloudresourcemanagerResult
}

func (s *Session) GetDeploymentManagerService() (*deploymentmanager.Service, error) {
	s.deploymentmanagerOnce.Do(func() {
		s.deploymentmanagerService, s.deploymentmanagerResult = deploymentmanager.New(s.client)
	})
	return s.deploymentmanagerService, s.deploymentmanagerResult
}

func (s *Session) GetLogger() (common.Logger, error) {
	if s.Logger == nil {
		return nil, errors.New("Logging not configured")
	}

	return s.Logger, nil
}

type loggerAdapter struct {
	gcpLogger    *logging.Logger
	generationId string
}

func (l *loggerAdapter) getEntry(s logging.Severity, v fmt.Stringer) logging.Entry {
	return logging.Entry{
		Severity: s,
		Payload:  v,
		Labels:   map[string]string{"generation-id": l.generationId},
	}
}

func (l *loggerAdapter) Debug(v fmt.Stringer) {
	log.Println(v.String())
}

func (l *loggerAdapter) Info(v fmt.Stringer) {
	if l.gcpLogger != nil {
		l.gcpLogger.Log(l.getEntry(logging.Info, v))
	}
	log.Println(v.String())
}

func (l *loggerAdapter) Warning(v fmt.Stringer) {
	if l.gcpLogger != nil {
		l.gcpLogger.Log(l.getEntry(logging.Warning, v))
	}
	log.Println(v.String())
}

func (l *loggerAdapter) Error(v fmt.Stringer) {
	if l.gcpLogger != nil {
		l.gcpLogger.Log(l.getEntry(logging.Error, v))
	}
	log.Println(v.String())
}
