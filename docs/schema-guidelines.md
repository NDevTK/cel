# Asset and Host Environment Schema

[TOC]

## Introduction

This document specifies guidelines for authoring schema changes for the Chrome
Enterprise Lab.

The schema is defined by the set of `.proto` files in the [schema] directory.
See [Protocol Buffers] for details on the syntax. These `.proto`
files are used to generate Go code which are included in the libraries and
tools.

Refer to the [design document] for overall design. Specifically, readers are
expected to be familiar with the [ASSET MANIFEST] section and the [HOST
ENVIRONMENT] sections of the design doc. For convenience, the Chrome Enterprise
Lab may be abbreviated as **CEL** in this document.

**The Schema Is a Living Specification.** The asset schema defines the types of
assets that are supported in the Chrome Enterprise Lab. It's expected that this
will evolve with the requirements of the Chromium project. Hence it'll never be
complete and must be designed in such a way that new asset types can be added
without too much churn.

## Nomenclature

### Asset Catalog

The **Asset Catalog** is a collection of assets that are recognized and can be
deployed by the lab. The canonical asset catalog is the list supported by the
toolchain at https://chromium.googlesource.com/enterprise/cel and described by
the `.proto` files in the [schema] directory.

E.g.: Examples of things that could be included in an asset catalog include
networks, VM instances, AD domains, IIS servers, MCS servers, authenticated
proxies, etc.


### Asset Inventory

Given an asset catalog, each team that wishes to use the CEL may define
instantiations of those assets in a reusable manner. This list of objects is
called the **Asset Inventory**. The schema should have the necessary facilities
to organize an asset inventory such that new tests or lab instantiations can
conveniently select a subset of assets from an existing inventory.

The CEL repository can maintain a sample asset inventory so that people have a
convenient set of examples to work with. However, the toolchain isn't bound by a
single canonical inventory.

E.g.: A network with specific properties, a preconfigured AD domain, VM
instances with specific properties.


## Guidelines For Authoring Schema

### A Few DOs and DON'Ts                                          {#do-and-dont}

*   **DO** take future growth into consideration. Some of the guidelines below
    are the result of such considerations. Growth in the asset catalog or any
    asset inventory shouldn't require inordinate refactoring.

*   **DO** keep the asset schema independent of the hosting environment. I.e. it
    should be possible for the same asset schema to be used if the hosting
    environment is changed from GCP to AWS or Azure. Hosting environment
    specific details should go in the host environment schema.

    While there are no plans currently to support a hosting environment other
    than GCP, having this rule allows us to evaluate whether the abstraction is
    correct whenever we are splitting the properties of an asset across the
    asset/host-environment boundary.

    If this split is kept clean, then any other team or organization would be
    able to instantiate a lab based on a asset manifest by providing the
    additional hosting information *without* modifying the asset manifest
    itself. The hope is that this would lead to easier sharing of assets, and
    easier integration of tests based on such assets into the mainline Chromium
    repository.

|||---|||
#### Yes

``` proto
message Machine {
string name = 1;
string os = 2;
string description = 3;
...
}
```

Descriptions and fields are generic. The host environment can describe how
different operating systems map to specific base image types.

#### No

``` proto
message Instance {
string base_image_url = 1;
...
}
```

Base image URLs are specific to hosting environment, and may even be
specific to a GCP project.
|||---|||

*   **DO** be consistent with developer nomenclature. The asset schema should
    map comfortably to how testers and developers think of networks.
   
*   **DO** be minimalistic.  The Asset Schema *should only include* properties
    that are **material** to the tests being considered, and *should exclude
    anything else*. Any additional required properties that are not material to
    the test can be specified via the Host Environment Schema. Avoid over
    specifying assets. Don't add knobs we don't plan on turning. I.e.  don't add
    asset attributes that don't have a known use case.

*   **DONT** expose implementation details of the deployment process or the
    toolchain. Avoid introducing asset definitions for intermediate objects that
    are only used during deployment. In other words, once again, the schema
    should allow the test developer to describe their requirements in a minimal
    fashion.

*   **DO** keep dependencies flowing in one direction only. In a parent-child
    relationship, the child is responsible for indicating its relationship to
    the parent, but not vice-versa. This allows new child assets to be defined
    without changing the parent asset.

|||---|||
#### Yes

``` proto
message Network {
  ...
}

message Machine {
  ...
  string network = 1;
}
```

#### No

``` proto
message Network {
  ...
  repeated string machines = 1;
}

message Machine {
  ...
}
```
|||---|||

*   **DO** stick to [Protobuf Style Guide] for style with the following
    exceptions:

    *   Naming convention for repeated elements is to use the **singular** form
	instead of the *plural* form encouraged in the style guide. Using the
	singular form results in more readable `textpb` files which is how the
	assets will be specified in Chromium and other labs.

    *   Ignore the stuff about Google3 and `BUILD` files. These are not
	applicable to the CEL project.

*   **DO** keep proto files manageable. Group asset schemas sensibly and move
    them into separate files. E.g. Keep Active Directory specific asset schemas
    in a single file and import that file from other protos.

*   **DON'T** nest messages. Nested message make the protos less readable and
    produces unreasonably large identifiers when generating stub code in Go.

*   **DO** anchor all imports at the root of the source tree. See existing files
    as an example.


### Boilerplate                                                   {#boilerplate}

All schemas use Protocol Buffers Version 3, with a suitable package name.  I.e.
Every `.proto` file should start with:

``` proto
// Copyright $YEAR The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

syntax = "proto3";

package $SUITABLE_PACKAGE_NAME;
option go_package="$ASSET_OR_HOST";
```

Where `$YEAR` should be the year you introduce the proto file, and
`$SUITABLE_PACAKGE_NAME` should be a single token identifying the package. Don't
forget to import any dependencies and properly document your messages.
`$ASSET_OR_HOST` is either `asset` or `host` depending on whether you are
authoring a proto for the asset manifest or the host manifest.

The Go package has to match the name of the directory. Until we have too much
schema that we need to organize by subdirectories, we are going to stick with
one Go package for all of it.


### Common Field Names

*   `name` : Every asset should have a name so that they can be referred to
    from other assets or tests. If the scope of the name isn't spelled out
    explicitly, it should be assumed that the name has global scope.

    Global in this case refers to a scope covering an entire asset inventory.

*   `based_on` : In order to minimize boilerplate when defining a group of
    assets with similar attributes, it should be possible to base one asset
    definition on another. This is done by adding a `based_on` attribute which
    names another asset with the same type and in the same namespace.

*   `description`: Where a longer description is needed, use the field name
    `description` instead of `details` or similar.

Overall, be internally consistent.


## Asset Types

There are three main types of assets (See design doc for details):

1.  Permanent
2.  On Demand
3.  Script

The considerations for designing schema for each type are as follows:

### 1. Permanent Assets

Asset Schema includes:

*   Name
*   _Verifiable properties_

Host Environment Schema includes:

*   Name
*   _Concrete properties_

E.g.: The following schema can describe a fixed service such as a web server
that runs on a specific host.

|||---|||
#### Asset Schema

``` proto
message FixedService {
  string name;
}
```

#### Host Environment Schema

``` proto
message FixedServiceInstance {
  string name;
  repeated DNSRecord dns_record;
}
```
|||---|||

### 2. On Demand Assets

These are assets that are fully deployed by the toolchain. As such, the
toolchain must be able to construct something in GCP based on the information it
has. A suggested pattern for splitting asset properties between asset schema and
host environment is as follows:

Asset Schema

*   Name
*   _Properties that are material to the test_
*   _Name of template from the host environment_

Host Environment Schema

*   Template-name
*   _Remainder of properties_

E.g.:

|||---|||
#### Asset Schema

``` proto
message Machine {
  string name;
  string machine_type;
  repeated NetworkInterface interface;
}
```

#### Host Environment Schema

``` proto
message MachineType {
  string name;
  GCEInstanceOptions instance_options;
}
```
|||---|||


Details like base image URL, zone, disk size, CPUs etc. will all go into the
Host Environment Schema. This way, each additional machine that's added to the
asset inventory can conveniently refer to a machine type via its name without
having to list out all the properties understood by GCP.


### 3. Script Assets

These are pretty much the same as On-Demand Assets, with the following
exceptions:

* A Script Asset must be independent of deployment strategy.

* A Script Asset must not directly depend on a Host Environment Schema
  component.


## Basic Schemas

*   Network
    *   Peers
*   Active Directory Domain
    *   AD DCs : Need at least one of these
    *   AD Containers
    *   ADMX Central Store (`%systemroot%\PolicyDefinitions`) : Don't need this
	unless people are manually interacting with the domain controller. See
	[ADMX Technology
	Review](https://technet.microsoft.com/en-us/library/cc749513(v=ws.10).aspx).
    *   AD GPOs
    *   AD Users
*   DNS Server
*   Certificate
*   IIS Server (version depends on the underlying OS)
    *   HTTP website
	*   CNAME
	*   Files
    *   HTTPS website
	*   CNAME
	*   Certificate
	*   Files
*   Client machine
    *   AD Member : Refers to a container
    *   Local Users
    *   Auto Login User : Required for Test Hosts. Can name a domain user.
    *   Test Host
*   Proxy server
    *   Type




[Protobuf Style Guide]: https://developers.google.com/protocol-buffers/docs/style
[Protocol Buffers]: https://developers.google.com/protocol-buffers/

[design document]: ../docs/START-HERE.md
[ASSET MANIFEST]: ../docs/asset-manifest.md
[HOST ENVIRONMENT]: ../docs/host-environment.md
[schema]: ../schema/


