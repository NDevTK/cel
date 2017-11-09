// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"net/http"
)

// GetDefaultClient returns an http.Client that authenticates with the
// "Application Default Credentials" for Google Compute Engine.  For more
// information, see:
// https://developers.google.com/accounts/docs/application-default-credentials
func GetDefaultClient(ctx context.Context) (*http.Client, error) {
	return google.DefaultClient(ctx, "https://www.googleapis.com/auth/compute")
}
