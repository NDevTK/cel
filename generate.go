// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file should list all the generator steps that need to be done in this
// repository. For the purpose of making `go get` and friends work, the
// generator steps must be run each time any of the dependent files change.
//
// The following step requires that the vendored dependency for Google API Go
// Client is checked out at vendor/googel.golang.org/api/. This should be the
// case after running `deps ensure`.
//
//go:generate go run go/cmd/gen_api_proto/main.go -i vendor/google.golang.org/api/compute/v0.beta/compute-api.json -o schema/compute/compute-api.proto
//
// Keep all generated files under the go/gen/ directory for convenience. Keep
// this list sorted.
//
//go:generate protoc --go_out=go/gen schema/asset/active_directory.proto
//go:generate protoc --go_out=go/gen schema/asset/cert.proto
//go:generate protoc --go_out=go/gen schema/asset/dns.proto
//go:generate protoc --go_out=go/gen schema/asset/iis.proto
//go:generate protoc --go_out=go/gen schema/asset/network.proto
//go:generate protoc --go_out=go/gen schema/host/host_environment.proto
//go:generate protoc --go_out=go/gen schema/compute/compute-api.proto

package cel
