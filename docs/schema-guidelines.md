# Asset and Host Environment Schema

[TOC]

## Introduction

This document specifies guidelines for authoring schema changes for the Chrome
Enterprise Lab.

The schema is defined by the set of `.proto` files in directories named `asset`
and `host` under the [schema][] directory.  See [Protocol Buffers][] for details
on the syntax for `.proto` files. These `.proto` files are used to generate Go
code which are included in the libraries and tools.

Refer to the [design document][Design] for overall design. Specifically, readers
are expected to be familiar with the [ASSET MANIFEST][] section and the [HOST
ENVIRONMENT][] sections of the design doc. For convenience, the Chrome Enterprise
Lab may be abbreviated as **CEL** in this document.

*** note
**The Schema Is a Living Specification.** The asset schema defines the types of
assets that are supported in the Chrome Enterprise Lab. It's expected that this
will evolve with the requirements of the Chromium project. Hence it'll never be
complete and must be designed in such a way that new asset types can be added
without too much churn.
***

## Concepts

### Asset Manifest

The **Asset Manifest** is a collection of assets that are recognized and can be
deployed by the lab. The canonical asset catalog is the list supported by the
toolchain at https://chromium.googlesource.com/enterprise/cel and described by
the `.proto` files in the [schema/asset][] directory.

Examples of things that could be included in an asset catalog include networks,
VM instances, AD domains, IIS servers, MCS servers, authenticated proxies, etc.

The Asset Manifest is rooted at the `AssetManifest` message
([source](../schema/asset/asset_manifest.proto)). All messages that appear in
`AssetManifest` are called top level assets.


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

Have a look at [an example](../examples/schema/ad/one-domain.asset.textpb) for
what a partial asset inventory looks like. Multiple assets of the same kind can
co-exist as long as they have unique identifiers and can be realized without
conflict.

### Host Environment

The **Host Environment** is the set of parameters that are required to define
how each of the assets in the Asset Inventory are realized. The Enterprise Lab
currently only supports hosting on Google Cloud Platform (GCP), so the host
environment is currently specific GCP.

All host environment parameters are rooted at the `HostEnvironment` message
defined in [`host_environment.proto`](../schema/host/host_environment.proto).

### Naming And References                                          {#references}

For the ease of identifying different objects, every object is given a name in
the conveniently named `name` field. You can see the names in the [asset][Asset
Example] and [host][Host Example] configuration examples.

The object being named can be referenced using familiar dot separated paths from
elsewhere in the configuration. The references are constructed as follows:

*   Field components are separated using `.`.

*   All references for assets start with the label `asset`. Host environment
    references start with the label `host`.

*   A field in a ProtoBuf message is referred-to using its field name. E.g. the
    collection of Windows machine objects in the asset manifest is
    `asset.windows_machine`.

*   An item in a collection is identified by the item's `name` value.

    E.g. The `machine_type` field of the following message in the asset manifest
    is `asset.windows_machine.dc.machine_type`.

    ``` conf
    # A Windows machine.
    windows_machine {
      name: 'dc'
      machine_type: 'win2012r2'
      network_interface { network: 'primary' }
    }
    ```

    This implies that an object in a collection must have a `name` field to be
    addressable.

### Validation                                                     {#validation}

The validation is performed based on the annotations that are specified in the
ProtoBuf schema. The validation rules are defined and explained in
[validation.proto][]. When defining schema, the
validation options can be specified as in the examples below:

From [active_directory.proto](/schema/asset/active_directory.proto):

``` proto
message WindowsContainer {
  oneof container {
    // Domain name.
    string ad_domain = 1 [(common.v).ref="asset.ad_domain"];

    // Machine name.
    string windows_machine = 2 [(common.v).ref="asset.windows_machine"];

    // Organizational unit.
    string ad_organizational_unit = 3 [(common.v).ref="asset.ad_organizational_unit"];
  }
}
```

In the above example, each field is annotated as being a reference to another
object in a collection. The value of the `ad_domain` field of a
`WindowsContainer` object must match the `name` value of a
`ActiveDirectoryDomain` object in `asset.ad_domain` colleciton.

There are additional parameters that you can use to restrict how a field is
validated. Refer to the [validation.proto][] file for complete list. The validator
recognizes several field types and attributes:

*   **`REQUIRED`**: The field must not be empty. Additionally, if this field type
    is applied to a `repeated`, `oneof`, or `map`, requires that at least one
    instance be specified.

*   **`LABEL`**: The field must be a `string` that matches the `<label>`
    production in [RFC 1035][].

*   **`FQDN`**: The field must be a `string` that matches `<subdomains>`
    production in [RFC 1035][].

*   **`ORGLABEL`**: THe field must be a `string` that matches the production
    `〈<subdomains> ':'〉<label>`. E.g.: `foo.example.com:my-project`.

In addition, any `name` field is assumed by default to be of type `LABEL`,
though they can be explicitly qualified to be of a different type like `FQDN` or
`ORGLABEL`.

[validation.proto]: /schema/common/validation.proto
[RFC 1035]: https://www.ietf.org/rfc/rfc1035.txt

### Inline References                                       {#inline-references}

String values in configuration files may contain inline references to other
string fields in the configuration. E.g.: A name of a source disk image may
refer to the project name of the host environment as follows:

``` conf
image {
  name: 'win2012r2'
  latest {
    project: '${host.project.name}'
    family: 'my-custom-family'
  }
}
```

In this case the special form `${host.project.name}` expands to the value of the
string field at `host.project.name`. The contents inside the `${...}` is treated
the same as what's described in [Schema References][] with the additional
constraint that the field being referenced *must* be a `string` field.

An additional constraint with inline references is that they can refer to output
fields. The latter are only available after the corresponding asset has been
resolved. For example, the `Image` host resource contains a `url` field that's
annotated as `OUTPUT`:

``` proto
message Image {
  ...
  // Output. Will contain the resolved base image URL on success.
  string url = 5 [(common.v).type=OUTPUT];
}
```

Once the `Image` object has been resolved, then the `url` field will contain the
partial GCP resource name of the base image to use. If the `url` field is
referenced from elsewhere in the configuration, say in a `MachineType`
definition, then The [DEPLOYER][] creates an implicit dependency from the
`MachineType` object to the `Image` object. This dependency prevents the
`MachineType` from being resolved prior to the `Image` being resolved.

See [Deployment Overview][] for details on *when* inline string references are
resolved.


## Guidelines For Authoring Schema

### A Few DOs and DON'Ts                                          {#do-and-dont}

*   **Take future growth into consideration.** Some of the guidelines below
    are the result of such considerations. Growth in the asset catalog or any
    asset inventory shouldn't require inordinate refactoring.

*   **Keep the asset schema independent of the hosting environment.** I.e. it
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

*   **Be consistent with developer nomenclature**. The asset schema should
    map comfortably to how testers and developers think of networks.

*   **Be minimalistic.**  The Asset Schema *should only include* properties
    that are **material** to the tests being considered, and *should exclude
    anything else*. Any additional required properties that are not material to
    the test can be specified via the Host Environment Schema. Avoid over
    specifying assets. Don't add knobs we don't plan on turning. I.e.  don't add
    asset attributes that don't have a known use case.

*   **Don't expose implementation details of the deployment process or the
    toolchain.** Avoid introducing asset definitions for intermediate objects that
    are only used during deployment. In other words, once again, the schema
    should allow the test developer to describe their requirements in a minimal
    fashion.

*   **Keep dependencies flowing in one direction only.** In a parent-child
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

*   **Stick to [Protobuf Style Guide][] for style** with the following
    exceptions:

    *   Naming convention for repeated elements is to use the **singular** form
	instead of the *plural* form encouraged in the style guide. Using the
	singular form results in more readable `textpb` files which is how the
	assets will be specified in Chromium and other labs.
    *   Ignore the stuff about Google3 and `BUILD` files. These are not
	applicable to the CEL project.

*   **Keep proto files manageable.** Group asset schemas sensibly and move
    them into separate files. E.g. Keep Active Directory specific asset schemas
    in a single file and import that file from other protos.

*   **Avoid nesting messages.** Nested message make the protos less readable and
    produces unreasonably large identifiers when generating stub code in Go.

*   **Anchor all imports at the root of the source tree.** See existing files
    for more examples:

    ``` proto
    import "schema/common/file_reference.proto";
    ```

*   **Top level messages must have a `name` field**: Every top level asset
    must have a name so that they can be referred to from other assets or tests.
    If the scope of the name isn't spelled out explicitly, it should be assumed
    that the name has global scope and case insensitive.

*   **Use RFC 1035 <label>s for names that are used in cross-references**. This
    is a hard and fast rule for GCP, but we are extending it to all messages
    where a `name` field needs to be referred from another message. This
    includes all top-level messages. The identifiers used for cross referencing
    should be eye-ball verifiable and easy to type. Hence the rule.

    It is possible that this restriction may affect the usefulness of some
    features. E.g. it should be possible to give a user a name that has unicode
    characters in it, or specify a FQDN for an ActiveDirectory domain. FQDNs
    don't comply with the `<label>` production since '.' is disallowed. In such
    a case, introduce a new field that doesn't need to be typed by hand each
    time. This second field is usually named `full_name`.

*   **Use `description` for longer descriptions**: Where a longer description is
    needed, use the field name `description` instead of `details` or some other
    similar name. Don't use `description` to refer to identifiers (See
    `full_name` above).

*   **Use the top level field name when referring to top level assets**: Fields
    that refer to top level assets must use the same name as the field used to
    define the top level asset.

    E.g.: Let's say we want to associate an IIS site with an IIS server. These
    are both top level assets defined in `asset_manifest.proto` as follows:

    ``` proto
    message AssetManifest {
      ...
      repeated IISServer iis_server =301;
      repeated IISSite iis_site = 302;
      ...
    }
    ```

    The `IISServer` message, being a top level asset, has a `name` attribute
    identifies it.

    ``` proto
    message IISServer {
      ...
      string name = 1;
      ...
    }
    ```

    When referring to a particualr server from the site asset, we use the same
    field name as that is used in `AssetManifest`. In this case it would be
    `iis_server`.

    ``` proto
    message IISSite {
      ...
      string name = 1;
      string iis_server = 2;
      ...
    }
    ```

    During deployment, the `iis_server` field of an `IISSite` asset must match a
    `name` of an `IISServer` asset.

    See? Simple. The same convention applies to references from assets to host
    environment except the names should be consistent with the field names in
    `HostEnvironment` ([source](../schema/host/host_environment.proto)). E.g.: A
    reference to a machine type from an asset should always be named
    `machine_type`.

*   **Use a `based_on` field for templating** : In order to minimize boilerplate
    when defining a group of assets with similar attributes, it should be
    possible to base one asset definition on another. This is done by adding a
    `based_on` attribute which names another asset with the same type and in the
    same namespace.

*** promo
Overall, be internally consistent.
***


### Boilerplate                                                   {#boilerplate}

All schemas use Protocol Buffers Version 3, with a suitable package name.  I.e.
Every `.proto` file should start with:

``` proto
// Copyright $YEAR The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

syntax = "proto3";
package $ASSET_OR_HOST;
option go_package="chromium.googlesource.com/enterprise/cel/go/$ASSET_OR_HOST";
```

Where `$YEAR` should be the year you introduce the proto file.  `$ASSET_OR_HOST`
is either `asset` or `host` depending on whether you are authoring a proto for
the asset manifest or the host manifest. Don't forget to import any dependencies
and properly document your messages.

For ease of copy&paste, use the following for asset protos:

``` proto
// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

syntax = "proto3";
package asset;
option go_package="chromium.googlesource.com/enterprise/cel/go/asset";
```

... and the following for host environment protos:

``` proto
// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

syntax = "proto3";
package host;
option go_package="chromium.googlesource.com/enterprise/cel/go/host";
```

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

## Asset Validation

Many fields in the asset schema take string fields and contain internal and
exteranl references. Constraints for these values come from many sources
including hosting environment. Hence, when adding new assets it's important to
also add a validator method that fulfils the `Validator` interface
([source](../go/common/validator.go)). See the `Validator` interface
documentation for how to write a validator.

Don't worry if you miss one. The default tests will fail if any `proto.Message`
classes are missing a `Validate()` method.

[Protobuf Style Guide]: https://developers.google.com/protocol-buffers/docs/style
[Protocol Buffers]: https://developers.google.com/protocol-buffers/
[schema]: ../schema/
[schema/asset]: ../schema/asset
[schema/host]: ../schema/host


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
