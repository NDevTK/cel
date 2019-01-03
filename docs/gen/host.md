# Schema `host` {#host}



Messages that are valid in package `host` are as follows:

*** note
Note that this document uses the term "message" to refer to the same concept as
a "message" in Protocol Buffers. Hence every asset and host resource
description is a *message*. So is their embedded structures.
***

## Message `AddressPool` {#AddressPool}

Describes an external address pool.


### Inputs for `AddressPool`

* `string` [`name`](#AddressPool.name) = 1 (**Required**)
* `repeated` [`asset.Address`](asset.md#Address) [`fixed_address`](#AddressPool.fixed_address) = 2
* `repeated` `string` [`reserved_address`](#AddressPool.reserved_address) = 3
* `repeated` `string` [`reserved_address_regex`](#AddressPool.reserved_address_regex) = 4

### `name` {#AddressPool.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of address pool. Used for matching incoming references from an asset
description.

### `fixed_address` {#AddressPool.fixed_address}

| Property | Comments |
|----------|----------|
| Field Name | `fixed_address` |
| Type | [`asset.Address`](asset.md#Address) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Static literal addresses.

### `reserved_address` {#AddressPool.reserved_address}

| Property | Comments |
|----------|----------|
| Field Name | `reserved_address` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

GCE reserved addresses by exact address name. Should refer to external
addresses.

### `reserved_address_regex` {#AddressPool.reserved_address_regex}

| Property | Comments |
|----------|----------|
| Field Name | `reserved_address_regex` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

GCE reserved addresses by regex. The provided regular expression must
match the entire name. The reserved address must be an external address.

E.g.:
    reserved_address_regex: "foo.*"
... matches "foobar", but not "egfoox"

## Message `Family` {#Image.Family}



### Inputs for `Family`

* `string` [`project`](#Image.Family.project) = 1
* `string` [`family`](#Image.Family.family) = 2 (**Required**)

### `project` {#Image.Family.project}

| Property | Comments |
|----------|----------|
| Field Name | `project` |
| Type | `string` |

The GCP project providing the image. Not a foreign key into Project
though since it is legal to refer to images provided by external
projects.

### `family` {#Image.Family.family}

| Property | Comments |
|----------|----------|
| Field Name | `family` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The GCP Image family name. Combined with the |project| field, the
|family| is used to locate the GCP image family. If the |url| is not
specified, then the resolver will pick the latest available image from
this family and populate the URL with it.

## Message `HostEnvironment` {#HostEnvironment}



### Inputs for `HostEnvironment`

* [`host.Project`](host.md#Project) [`project`](#HostEnvironment.project) = 1 (**Required**)
* [`host.Storage`](host.md#Storage) [`storage`](#HostEnvironment.storage) = 2 (**Required**)
* [`host.LogSettings`](host.md#LogSettings) [`log_settings`](#HostEnvironment.log_settings) = 3 (**Required**)
* `repeated` [`host.MachineType`](host.md#MachineType) [`machine_type`](#HostEnvironment.machine_type) = 10
* `repeated` [`host.AddressPool`](host.md#AddressPool) [`address_pool`](#HostEnvironment.address_pool) = 11
* `repeated` [`host.Image`](host.md#Image) [`image`](#HostEnvironment.image) = 12

### Outputs for `HostEnvironment`

* [`host.RuntimeSupport`](host.md#RuntimeSupport) [`resources`](#HostEnvironment.resources) = 100

### `project` {#HostEnvironment.project}

| Property | Comments |
|----------|----------|
| Field Name | `project` |
| Type | [`host.Project`](host.md#Project) |
| Required | This field is required. It is an error to omit this field. |

### `storage` {#HostEnvironment.storage}

| Property | Comments |
|----------|----------|
| Field Name | `storage` |
| Type | [`host.Storage`](host.md#Storage) |
| Required | This field is required. It is an error to omit this field. |

### `log_settings` {#HostEnvironment.log_settings}

| Property | Comments |
|----------|----------|
| Field Name | `log_settings` |
| Type | [`host.LogSettings`](host.md#LogSettings) |
| Required | This field is required. It is an error to omit this field. |

### `machine_type` {#HostEnvironment.machine_type}

| Property | Comments |
|----------|----------|
| Field Name | `machine_type` |
| Type | [`host.MachineType`](host.md#MachineType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `address_pool` {#HostEnvironment.address_pool}

| Property | Comments |
|----------|----------|
| Field Name | `address_pool` |
| Type | [`host.AddressPool`](host.md#AddressPool) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `image` {#HostEnvironment.image}

| Property | Comments |
|----------|----------|
| Field Name | `image` |
| Type | [`host.Image`](host.md#Image) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `resources` {#HostEnvironment.resources}

| Property | Comments |
|----------|----------|
| Field Name | `resources` |
| Type | [`host.RuntimeSupport`](host.md#RuntimeSupport) |

Determined at deployment time.

## Message `Image` {#Image}

Describes a GCE source disk image. See
https://cloud.google.com/compute/docs/images#image_families

The image can be specified either using the project/family pair or using a
direct URL. During asset resolution, a project/family pair will be resolved
into a URL. I.e. a resolved Image{} resource always contains the URL of the
selected image.


### Inputs for `Image`

* `string` [`name`](#Image.name) = 1 (**Required**)
* [`host.Image.Family`](host.md#Image.Family) [`latest`](#Image.latest) = 2
* `string` [`fixed`](#Image.fixed) = 3
* `repeated` `string` [`package`](#Image.package) = 4

### Outputs for `Image`

* `string` [`url`](#Image.url) = 5

### `name` {#Image.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the image. Used within the host environment schema to refer to
this image.

### `latest` {#Image.latest}

| Property | Comments |
|----------|----------|
| Field Name | `latest` |
| Type | [`host.Image.Family`](host.md#Image.Family) |

If specified, indicates that the latest image matching project/family
should be used as the base.

### `fixed` {#Image.fixed}

| Property | Comments |
|----------|----------|
| Field Name | `fixed` |
| Type | `string` |

The full or partial URL to the disk image. If this is specified, then
the |project| and |family| values are ignored.

### `package` {#Image.package}

| Property | Comments |
|----------|----------|
| Field Name | `package` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Additional packages that should be applied to this image. The
interpretation of this field is dependent on the type of machine being
built. For example when building a WindowsMachine, these entries are
prepended to the list of 'windows_feature' values for the machine.

### `url` {#Image.url}

| Property | Comments |
|----------|----------|
| Field Name | `url` |
| Type | `string` |

Output. Will contain the resolved base image URL on success.

## Message `LogSettings` {#LogSettings}

Stackdriver log settings.


### Inputs for `LogSettings`

* `string` [`admin_log`](#LogSettings.admin_log) = 1 (**Required**)

### `admin_log` {#LogSettings.admin_log}

| Property | Comments |
|----------|----------|
| Field Name | `admin_log` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of log used for administrative messages. These are logged during the
deployment process and also when an instance is started / configured /
shutdown.

## Message `MachineType` {#MachineType}

Describes a type of machine (virtual or otherwise).

Not to be confused with GCE machine type. I.e.
These are not https://cloud.google.com/compute/docs/machine-types


### Inputs for `MachineType`

* `string` [`name`](#MachineType.name) = 1 (**Required**)
* [`compute.InstanceProperties`](gcp_compute.md#InstanceProperties) [`instance_properties`](#MachineType.instance_properties) = 5
* `string` [`instance_template`](#MachineType.instance_template) = 6 (**Required**)
* [`host.NestedVM`](host.md#NestedVM) [`nested_vm`](#MachineType.nested_vm) = 7

### `name` {#MachineType.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name is used to match machine type name from asset description.

### `instance_properties` {#MachineType.instance_properties}

| Property | Comments |
|----------|----------|
| Field Name | `instance_properties` |
| Type | [`compute.InstanceProperties`](gcp_compute.md#InstanceProperties) |

InstanceProperties are used for constructing a new GCE instance.

The instance_properties.AttachedDisk.Source can refer to an image name
using the syntax "${host.image.<imagename>.url}" where <imagename> is
the name of an Image object in the enclosing HostEnvironment.

### `instance_template` {#MachineType.instance_template}

| Property | Comments |
|----------|----------|
| Field Name | `instance_template` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

An instance template to use for constructing a new GCE instance. Should
be a full or a partial URL.

### `nested_vm` {#MachineType.nested_vm}

| Property | Comments |
|----------|----------|
| Field Name | `nested_vm` |
| Type | [`host.NestedVM`](host.md#NestedVM) |

## Message `NestedVM` {#NestedVM}



### Inputs for `NestedVM`

* `string` [`image`](#NestedVM.image) = 1 (**Required**)
* `string` [`user_name`](#NestedVM.user_name) = 2
* `string` [`password`](#NestedVM.password) = 3
* `string` [`machineType`](#NestedVM.machineType) = 4

### `image` {#NestedVM.image}

| Property | Comments |
|----------|----------|
| Field Name | `image` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The image to use. It's a gs://path.

### `user_name` {#NestedVM.user_name}

| Property | Comments |
|----------|----------|
| Field Name | `user_name` |
| Type | `string` |

The user name & password used to log in through ssh.

### `password` {#NestedVM.password}

| Property | Comments |
|----------|----------|
| Field Name | `password` |
| Type | `string` |

### `machineType` {#NestedVM.machineType}

| Property | Comments |
|----------|----------|
| Field Name | `machineType` |
| Type | `string` |

The GCE machine type used for the host. Default is n1-standard-2.

## Message `Project` {#Project}

GCP project hosting.


### Inputs for `Project`

* `string` [`name`](#Project.name) = 1
* `string` [`zone`](#Project.zone) = 2 (**Required**)

### Outputs for `Project`

* [`compute.Project`](gcp_compute.md#Project) [`project`](#Project.project) = 100

### `name` {#Project.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |

The project name. In GCP, this serves as the project ID.

E.g. `my-project-123`.

### `zone` {#Project.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The default zone to use when constructing resources. The |zone| also
implicitly defines the region. See
https://cloud.google.com/compute/docs/regions-zones/ for an up-to-date
list of regions and zones.

Only the zone identifier should be entered here. E.g. "us-west1-a"

### `project` {#Project.project}

| Property | Comments |
|----------|----------|
| Field Name | `project` |
| Type | [`compute.Project`](gcp_compute.md#Project) |

Project

## Message `RuntimeSupport` {#RuntimeSupport}

Base support resources that must exist at the time the lab assets are
deployed.

These values are output only and are determined at deployment time. They are
part of the schema so that on-host logic can refer to these values at
runtime and so that these values will appear in the completed asset
manifest.


### Inputs for `RuntimeSupport`

* [`host.Startup`](host.md#Startup) [`startup`](#RuntimeSupport.startup) = 102 (**Required**)

### Outputs for `RuntimeSupport`

* `.google.iam.admin.v1.ServiceAccount` [`service_account`](#RuntimeSupport.service_account) = 100
* `.cloudkms.CryptoKey` [`crypto_key`](#RuntimeSupport.crypto_key) = 101

### `service_account` {#RuntimeSupport.service_account}

| Property | Comments |
|----------|----------|
| Field Name | `service_account` |
| Type | `.google.iam.admin.v1.ServiceAccount` |

The GCP service account that will be used on all lab VMs.

### `crypto_key` {#RuntimeSupport.crypto_key}

| Property | Comments |
|----------|----------|
| Field Name | `crypto_key` |
| Type | `.cloudkms.CryptoKey` |

CryptoKey used for encrypting/decrypting privileged information between
the deployer and the instance VMs.

### `startup` {#RuntimeSupport.startup}

| Property | Comments |
|----------|----------|
| Field Name | `startup` |
| Type | [`host.Startup`](host.md#Startup) |
| Required | This field is required. It is an error to omit this field. |

Startup dependencies.

## Message `Startup` {#Startup}



### Outputs for `Startup`

* [`common.FileReference`](common.md#FileReference) [`win_startup`](#Startup.win_startup) = 100
* [`common.FileReference`](common.md#FileReference) [`win_agent_x64`](#Startup.win_agent_x64) = 101
* [`common.FileReference`](common.md#FileReference) [`linux_startup`](#Startup.linux_startup) = 102
* [`common.FileReference`](common.md#FileReference) [`linux_agent_x64`](#Startup.linux_agent_x64) = 103

### `win_startup` {#Startup.win_startup}

| Property | Comments |
|----------|----------|
| Field Name | `win_startup` |
| Type | [`common.FileReference`](common.md#FileReference) |

Windows startup file. Must be a Powershell ps1 file.

### `win_agent_x64` {#Startup.win_agent_x64}

| Property | Comments |
|----------|----------|
| Field Name | `win_agent_x64` |
| Type | [`common.FileReference`](common.md#FileReference) |

Windows CEL Agent executable.

### `linux_startup` {#Startup.linux_startup}

| Property | Comments |
|----------|----------|
| Field Name | `linux_startup` |
| Type | [`common.FileReference`](common.md#FileReference) |

Linux startup file.

### `linux_agent_x64` {#Startup.linux_agent_x64}

| Property | Comments |
|----------|----------|
| Field Name | `linux_agent_x64` |
| Type | [`common.FileReference`](common.md#FileReference) |

Linux CEL Agent executable.

## Message `Storage` {#Storage}

Google Cloud Storage bucket to use for deployment and metadata management
purposes.

The storage configured here must exist throughout the entire lifetime of the
lab. In addition, the bucket should ideally belong to the same project as
the lab.

TODO(asanka): Should include ways to configure retention of objects, and
perhaps customize ACLs.


### Inputs for `Storage`

* `string` [`bucket`](#Storage.bucket) = 1 (**Required**)
* `string` [`prefix`](#Storage.prefix) = 2

### Outputs for `Storage`

* `string` [`created_on`](#Storage.created_on) = 10

### `bucket` {#Storage.bucket}

| Property | Comments |
|----------|----------|
| Field Name | `bucket` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the GCS bucket. No `gs:` prefix. Just the name.

### `prefix` {#Storage.prefix}

| Property | Comments |
|----------|----------|
| Field Name | `prefix` |
| Type | `string` |

A prefix to attach to every named object stored in the GCS bucket. If left
empty, the default prefix of cel-config will be used.

The prefix must not end in a forward slash, though it can contain forward
slashes.

A lab will not access objects outside of the prefix. Hence the prefix used
at deployment time must be passed along to the client instances. All
object references are validated against the prefix. The CEL toolchain will
handle this for you. However, it does mean that object references created
against one prefix will not work with another.

### `created_on` {#Storage.created_on}

| Property | Comments |
|----------|----------|
| Field Name | `created_on` |
| Type | `string` |

The time at which the storage bucket was created as a RFC 3339 formatted
string.



# Enumerations


---
Generated from `schema/host/host_environment.proto`.
