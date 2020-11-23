# WHERE'S THE CODE

[TOC]

## Background

This document gives you some idea of where to find which code, and also a high
level overview of how we've organized the code.

The toolchain is written primarily in Go. Hence the organization of the
directories follows what's required by Go.

## Top Level Directories                                       {#top-level-dirs}

* **`go`**: All Go sources, test files, immediate dependencies of tests.

* **`schema`**: ProtoBuf schema in the form of `.proto` files. These are used to
  generate Go code stubs that are then placed under the `go/` tree.

* **`docs`**: Documentation.

* **`examples`**: Examples. Not used by the build, but are there for your
  perusal.

* **`resources`**: Resources that are embedded into the toolchain.

* **`build`**: Code used by `build.py`.

* **`scripts`**: Example and miscellaneous scripts that are not directly used by
  the build system nor the toolchain.

In addition to the above, at after a building you'll find a few additional
directories as follows:

* **`out`**: Directory containing build artifacts.

* **`vendor`**: Directory containing external dependencies that were pulled down
  by the [Go Dep](https://golang.github.io/dep/) tool.

* **`third_party`**: Additional dependencies that are not managed by the Go
  toolchain.

## Go Code Organization                                               {#go-code}

The entire purpose of the codebase is to produce two binaries and their
supporting documentation.

The two binaries are:

* `cel_ctl`: The commandline control utility. This is the primary (at the
  moment) interface for users to interact with the toolchain.

* `cel_agent`: This is the agent binary that is installed and run on every VM
  that's created within the lab. Multiple of these may be built during a single
  build to account for the different architectures of the VMs.

### Separation of Abstration                                  {#go-abstractions}

* **Assets** vs **Host Environment**: These two concepts are described in detail
  in [ASSET MANIFEST][] and [HOST ENVIRONMENT][] respectively. In code we strive
  to maintain a package-wise separation between the two.

  The toolchain currently only supports Google Cloud Platform as a backend for
  hosting labs. This may change in the near future (already changed and the
  documentation is out-of-date? Let us know! File a bug
  [here](https://new.crbug.com)). For mobility in this space and for general
  code health, we want to maintain a separation between the `ASSET MANIFEST` and
  the `HOST ENVIRONMENT` code.

  *** note
  **What does "package-wise separation" mean?**

  Package-wise separation means that the code pertaining to HOST ENVIRONMENT
  must need to access code pertainng to ASSET MANIFEST. Since Go imports can't
  be cyclical, this implies that code pertaining to ASSET MANIFEST must
  necessarily be independent of the HOST ENVIRONMENT. This way the Go toolchain
  will help us maintain code health.
  ***

* **Deployer** vs **Agent**: The deployer is the primary task of the `cel_ctl`
  tool. The Agent is primarily the `cel_agent` binary. Code that's specific to
  one should be kept in separate Go packages from code that's specific to the
  other.

  In code, we maintain deployer/agent separation via subdirectories named
  `deploy` and `agent`.

* **GCP** vs **Not-GCP**: The goal here is to isolate GCP code into its own
  packages such that only those packages can import GCP API code. This, once
  again, is for maintaining code health by preventing unnecessary dependencies
  on GCP by packages that shouldn't.

  In code, this manifests itself as the `go/gcp` directory which contains all
  the GCP specific code.

### Source Directories                                                {#go-dirs}

The subdirectories under the top level `go/` directory are:

* **`cmd`**: Commands.

  * **`cmd/cel_ctl`**: The CEL controller tool. This directory should only
    contain the logic needed to understand commandline options. Code in the
    contained `main` package can import any other Go package other than those
    under `tools`, and other subdirectories of `cmd`.

    One of the primary duties of the `cel_ctl` command is to manage the
    deployement of assets into a lab. The code for doing that is in
    `go/cel/deploy` as described below.

  * **`cmd/cel_agent`**: The CEL agent tool. This directory should only contain
    the logic needed to understand commandline options. Code in the contained
    `main` package can import any other Go package other than those under
    `tools`, and other subdirectories of `cmd`.

* **`common`**: Common logic that's shared between many of the components of the
  CEL toolchain. There are `deploy` and `agent` subdirectories under here for
  deployer and agent specific code.

* **`cel`**: High level workflow. Once again `deploy` and `agent` subdirectories
  further separate out the workflow logic.

* **`asset`**: Asset Manifest specific code.

* **`host`**: Host Environment specific code.

* **`gcp`**: Go code that depends directly on Google Cloud Platform APIs and
  abstractions. Subdirectories therein have generated code for dealing with GCP
  API specific schemas (e.g.: `go/gcp/compute` contains generate code for
  managing data structures needed to talk to the Google Compute Engine API.

  There are `deploy` and `agent` subdirectories under here for deployer and
  agent specific code.

* **`tools`**: These tools are indirect dependencies for `cel_ctl` and
  `cel_agent`.


<!-- INCLUDE index.md (55 lines) -->
<!--
Index of tags used throughout the documentation. This list lives in
/docs/index.md and is included in all documents that depend on these tags.

In order to update the tags:

   1. Update `/docs/index.md`
   2. Run the following command from the root of the source tree:

         ./build.py format

Keep the tags below sorted.
-->

[ASSET MANIFEST]: design-summary.md#asset-manifest
[Additional Considerations]: background.md#additional-considerations
[Asset Description Schema]: schema-guidelines.md
[Asset Example]: /examples/schema/ad/one-domain.asset.textpb
[Asset Schema]: /schema/asset/
[Background]: background.md
[Bootstrapping]: bootstrapping.md
[Coding Patterns for Resolvers]: deployment.md#coding-patterns-for-resolvers
[Completed Asset Manifest]: deployment.md#completed-asset-manifest
[Concepts]: design-summary.md#concepts
[DEPLOYER]: design-summary.md#deployer
[Deploying Scripted Assets]: deployment.md#deploying-scripted-assets
[Deployment Details]: deployment.md
[Deployment Overview]: deployment.md#overview
[Design]: design-summary.md
[Frameworks/Tools Used]: background.md#tools-used
[GREETER]: design-summary.md#greeter
[Google Services]: google-services.md
[HOST ENVIRONMENT]: design-summary.md#host-environment
[HOST TEST RUNNER]: design-summary.md#host-test-runner
[Host Environment Schema]: /schema/host/
[Host Example]: /examples/schema/ad/one-domain.host.textpb
[ISOLATE]: design-summary.md#isolate
[Inline References]: deployment.md#inline-references
[Integration With Chromium Waterfall]: chrome-ci-integration.md
[Key Management]: key-management.md
[Objective]: design-summary.md#objective
[On-Premise Fixtures]: on-premise-fixtures.md
[Private Google Compute Images]: private-images.md
[SYSTEM TEST RUNNER]: design-summary.md#system-test-runner
[Scalability]: scalability.md
[Schema References]: schema-guidelines.md#references
[Schema Validation]: schema-guidelines.md#validation
[Source Locations]: source-locations.md
[TEST HOST]: design-summary.md#test-host
[TEST]: design-summary.md#test
[The Product]: design-summary.md#the-product
[Use Cases]: background.md#use-cases
[Workflows]: workflows.md
[cel_bot]: design-summary.md#cel_bot
[cel_py]: design-summary.md#cel_py
