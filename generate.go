// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file should list all the generator steps that need to be done in this
// repository. For the purpose of making `go get` and friends work, the
// generator steps must be run each time any of the dependent files change.

// The following step requires that the vendored dependency for Google API Go
// Client is checked out at vendor/googel.golang.org/api/. This should be the
// case after running `deps ensure`. See the CONTRIBUTING.md file for details.
//
//go:generate go run go/cmd/gen_api_proto/main.go -i vendor/google.golang.org/api/compute/v0.beta/compute-api.json -o schema/gcp/compute/compute-api.proto -g chromium.googlesource.com/enterprise/cel/go/gcp

// Note that all the .proto files that go into a single package should be
// specified on the same protoc invocation.
//
//go:generate protoc --go_out=../../../ schema/asset/active_directory.proto schema/asset/cert.proto schema/asset/dns.proto schema/asset/iis.proto schema/asset/network.proto schema/asset/common.proto
//go:generate protoc --go_out=../../../ schema/host/host_environment.proto
//go:generate protoc --go_out=../../../ schema/gcp/compute/compute-api.proto

package cel
