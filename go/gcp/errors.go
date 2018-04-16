// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"bytes"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	googleapi "google.golang.org/api/googleapi"
	"net/url"
	"text/template"
)

type gcpErrorResponse struct {
	Error string `json:"error"`
}

func HandleGcpError(err error, project, service string) error {
	if err == nil {
		return nil
	}

	if IsBadCredentialsError(err) {
		return &BadCredentialsError{InnerError: err}
	}

	if project != "" && service != "" && IsServiceNotEnabledError(err) {
		return &ServiceNotEnabledError{
			Project:    project,
			Service:    service,
			InnerError: err,
		}
	}

	return err
}

// IsNotFoundError returns true if |err| indicates that a requested Google
// cloud resource was not found.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	inner := errors.Cause(err)
	if e, ok := inner.(*googleapi.Error); ok {
		return e.Code == 404
	}
	return false
}

func IsBadCredentialsError(err error) bool {
	if err == nil {
		return false
	}
	err = errors.Cause(err)
	if u, ok := err.(*url.Error); ok {
		if oa, ok := u.Err.(*oauth2.RetrieveError); ok {
			if oa.Response.StatusCode != 400 {
				return false
			}

			var r gcpErrorResponse
			if err = json.Unmarshal(oa.Body, &r); err != nil {
				return false
			}

			return r.Error == "invalid_grant"
		}
	}
	return false
}

type BadCredentialsError struct {
	InnerError error
}

func (b *BadCredentialsError) Error() string {
	return fmt.Sprintf(`invalid credentials.

The Application Default Credentials[1] that were used to access the Google
Cloud APIs were not valid. Note that the these credentials can be different
from the credentials that are used by the 'gcloud' tool.

To check your Application Default Credentials, use the following command:

    gcloud auth application-default print-access-token

If the token is reported as invalid, or you need to re-obtain your credentials:

    gcloud auth application-default login

References:

  [1] Application Default Credentials: https://cloud.google.com/docs/authentication/production
`)
}

// IsServiceNotEnabledError returns true if it's /likely/ that the error
// indicates that the service is not enabled.
//
// The input error is assumed to be an error that's returned from a call to a
// Google Cloud REST endpoint.
func IsServiceNotEnabledError(err error) bool {
	// TODO(asanka): Implement this.
	return false
}

type ServiceNotEnabledError struct {
	Project    string
	Service    string
	InnerError error
}

func (s *ServiceNotEnabledError) Error() string {
	var b bytes.Buffer
	err := template.Must(template.New("msg").Parse(`the service "{{ .Service }}" is not enabled for project "{{ .Project }}".

The CEL toolchain attempts to enable some services, but there are those that
need to be enabled manually. In order to address this, try the following:

  1) Make sure that the default credentials that you are using has the
	 necessary privileges to manage services and APIs. You can verify your
	 identity using the following gcloud command:

	   gcloud info

  2) Manually enable the desired service.

	   gcloud services enable --project {{ .Project }} {{ .Service }}

Underlying error: {{ .InnerError }}

`)).Execute(&b, s)
	if err != nil {
		return s.InnerError.Error()
	}

	return b.String()
}

func (s *ServiceNotEnabledError) Cause() error {
	return s.InnerError
}

func GcpLoggedAction(s *Session, err *error, format string, v ...interface{}) func() {
	return GcpLoggedServiceAction(s, "", err, format, v...)
}

func GcpLoggedServiceAction(s *Session, service string, err *error, format string, v ...interface{}) func() {
	inner := common.LoggedAction(s.Logger, err, format, v...)
	return func() {
		if *err != nil {
			*err = HandleGcpError(*err, s.GetProject(), service)
		}
		inner()
	}
}
