// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"context"
	"go.chromium.org/luci/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

const SCOPE string = "https://www.googleapis.com/auth/cloud-platform"

// GetDefaultClient returns an http.Client that authenticates with luci-auth
// if running on LUCI bots, or with Application Default Credentials otherwise.
// For more information, see:
// https://developers.google.com/accounts/docs/application-default-credentials
func GetDefaultClient(ctx context.Context) (*http.Client, error) {
	ts, err := GetDefaultTokenSource(ctx)
	if err != nil {
		return nil, err
	}

	return oauth2.NewClient(ctx, ts), nil
}

// This returns the token source to use for our GCP requests.
func GetDefaultTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	if _, ok := os.LookupEnv("LUCI_CONTEXT"); ok {
		opts := auth.Options{Scopes: []string{SCOPE}}
		authenticator := auth.NewAuthenticator(ctx, auth.SilentLogin, opts)
		return authenticator.TokenSource()
	}

	return google.DefaultTokenSource(ctx, SCOPE)
}
