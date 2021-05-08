# Schema `compute` {#compute}

Generated protobuf for compute

--- Skip validation ---


Messages that are valid in package `compute` are as follows:

*** note
Note that this document uses the term "message" to refer to the same concept as
a "message" in Protocol Buffers. Hence every asset and host resource
description is a *message*. So is their embedded structures.
***

## Message `AcceleratorConfig` {#AcceleratorConfig}

A specification of the type and number of accelerator cards attached to the
instance.


### Inputs for `AcceleratorConfig`

* `int32` [`acceleratorCount`](#AcceleratorConfig.acceleratorCount) = 1
* `string` [`acceleratorType`](#AcceleratorConfig.acceleratorType) = 2

### `acceleratorCount` {#AcceleratorConfig.acceleratorCount}

| Property | Comments |
|----------|----------|
| Field Name | `acceleratorCount` |
| Type | `int32` |

The number of the guest accelerator cards exposed to this instance.

### `acceleratorType` {#AcceleratorConfig.acceleratorType}

| Property | Comments |
|----------|----------|
| Field Name | `acceleratorType` |
| Type | `string` |

Full or partial URL of the accelerator type resource to attach to this
instance. If you are creating an instance template, specify only the
accelerator name.

## Message `AcceleratorType` {#AcceleratorType}

An Accelerator Type resource. (== resource_for beta.acceleratorTypes ==) (==
resource_for v1.acceleratorTypes ==)


### Inputs for `AcceleratorType`

* `string` [`creationTimestamp`](#AcceleratorType.creationTimestamp) = 1
* [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) [`deprecated`](#AcceleratorType.deprecated) = 2
* `string` [`description`](#AcceleratorType.description) = 3
* `string` [`id`](#AcceleratorType.id) = 4
* `string` [`kind`](#AcceleratorType.kind) = 5
* `int32` [`maximumCardsPerInstance`](#AcceleratorType.maximumCardsPerInstance) = 6
* `string` [`name`](#AcceleratorType.name) = 7 (**Required**)
* `string` [`selfLink`](#AcceleratorType.selfLink) = 8
* `string` [`zone`](#AcceleratorType.zone) = 9

### `creationTimestamp` {#AcceleratorType.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `deprecated` {#AcceleratorType.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) |

[Output Only] The deprecation status associated with this accelerator type.

### `description` {#AcceleratorType.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] An optional textual description of the resource.

### `id` {#AcceleratorType.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#AcceleratorType.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The type of the resource. Always compute#acceleratorType for
accelerator types.

### `maximumCardsPerInstance` {#AcceleratorType.maximumCardsPerInstance}

| Property | Comments |
|----------|----------|
| Field Name | `maximumCardsPerInstance` |
| Type | `int32` |

[Output Only] Maximum accelerator cards allowed per instance.

### `name` {#AcceleratorType.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `selfLink` {#AcceleratorType.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined fully-qualified URL for this resource.

### `zone` {#AcceleratorType.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] The name of the zone where the accelerator type resides, such
as us-central1-a. You must specify this field as part of the HTTP request
URL. It is not settable as a field in the request body.

## Message `AcceleratorTypeAggregatedList` {#AcceleratorTypeAggregatedList}



### Inputs for `AcceleratorTypeAggregatedList`

* `string` [`id`](#AcceleratorTypeAggregatedList.id) = 1
* `repeated` [`compute.AcceleratorTypeAggregatedList.ItemsEntry`](gcp_compute.md#AcceleratorTypeAggregatedList.ItemsEntry) [`items`](#AcceleratorTypeAggregatedList.items) = 2
* `string` [`kind`](#AcceleratorTypeAggregatedList.kind) = 3
* `string` [`nextPageToken`](#AcceleratorTypeAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#AcceleratorTypeAggregatedList.selfLink) = 5
* [`compute.AcceleratorTypeAggregatedList.Warning`](gcp_compute.md#AcceleratorTypeAggregatedList.Warning) [`warning`](#AcceleratorTypeAggregatedList.warning) = 6

### `id` {#AcceleratorTypeAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#AcceleratorTypeAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.AcceleratorTypeAggregatedList.ItemsEntry`](gcp_compute.md#AcceleratorTypeAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of AcceleratorTypesScopedList resources.

### `kind` {#AcceleratorTypeAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#acceleratorTypeAggregatedList
for aggregated lists of accelerator types.

### `nextPageToken` {#AcceleratorTypeAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#AcceleratorTypeAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#AcceleratorTypeAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AcceleratorTypeAggregatedList.Warning`](gcp_compute.md#AcceleratorTypeAggregatedList.Warning) |

## Message `AcceleratorTypeList` {#AcceleratorTypeList}

Contains a list of accelerator types.


### Inputs for `AcceleratorTypeList`

* `string` [`id`](#AcceleratorTypeList.id) = 1
* `repeated` [`compute.AcceleratorType`](gcp_compute.md#AcceleratorType) [`items`](#AcceleratorTypeList.items) = 2
* `string` [`kind`](#AcceleratorTypeList.kind) = 3
* `string` [`nextPageToken`](#AcceleratorTypeList.nextPageToken) = 4
* `string` [`selfLink`](#AcceleratorTypeList.selfLink) = 5
* [`compute.AcceleratorTypeList.Warning`](gcp_compute.md#AcceleratorTypeList.Warning) [`warning`](#AcceleratorTypeList.warning) = 6

### `id` {#AcceleratorTypeList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#AcceleratorTypeList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.AcceleratorType`](gcp_compute.md#AcceleratorType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of AcceleratorType resources.

### `kind` {#AcceleratorTypeList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#acceleratorTypeList for lists
of accelerator types.

### `nextPageToken` {#AcceleratorTypeList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#AcceleratorTypeList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#AcceleratorTypeList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AcceleratorTypeList.Warning`](gcp_compute.md#AcceleratorTypeList.Warning) |

## Message `AcceleratorTypesScopedList` {#AcceleratorTypesScopedList}



### Inputs for `AcceleratorTypesScopedList`

* `repeated` [`compute.AcceleratorType`](gcp_compute.md#AcceleratorType) [`acceleratorTypes`](#AcceleratorTypesScopedList.acceleratorTypes) = 1
* [`compute.AcceleratorTypesScopedList.Warning`](gcp_compute.md#AcceleratorTypesScopedList.Warning) [`warning`](#AcceleratorTypesScopedList.warning) = 2

### `acceleratorTypes` {#AcceleratorTypesScopedList.acceleratorTypes}

| Property | Comments |
|----------|----------|
| Field Name | `acceleratorTypes` |
| Type | [`compute.AcceleratorType`](gcp_compute.md#AcceleratorType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of accelerator types contained in this scope.

### `warning` {#AcceleratorTypesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AcceleratorTypesScopedList.Warning`](gcp_compute.md#AcceleratorTypesScopedList.Warning) |

## Message `AccessConfig` {#AccessConfig}

An access configuration attached to an instance's network interface. Only
one access config per instance is supported.


### Inputs for `AccessConfig`

* `string` [`kind`](#AccessConfig.kind) = 1
* `string` [`name`](#AccessConfig.name) = 2 (**Required**)
* `string` [`natIP`](#AccessConfig.natIP) = 3
* `string` [`publicPtrDomainName`](#AccessConfig.publicPtrDomainName) = 4
* `bool` [`setPublicPtr`](#AccessConfig.setPublicPtr) = 5
* `string` [`type`](#AccessConfig.type) = 6

### `kind` {#AccessConfig.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#accessConfig for access
configs.

### `name` {#AccessConfig.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name of this access configuration. The default and recommended name is
External NAT but you can use any arbitrary string you would like. For
example, My external IP or Network Access.

### `natIP` {#AccessConfig.natIP}

| Property | Comments |
|----------|----------|
| Field Name | `natIP` |
| Type | `string` |

An external IP address associated with this instance. Specify an unused
static external IP address available to the project or leave this field
undefined to use an IP from a shared ephemeral IP address pool. If you
specify a static external IP address, it must live in the same region as the
zone of the instance.

### `publicPtrDomainName` {#AccessConfig.publicPtrDomainName}

| Property | Comments |
|----------|----------|
| Field Name | `publicPtrDomainName` |
| Type | `string` |

The DNS domain name for the public PTR record. This field can only be set
when the set_public_ptr field is enabled.

### `setPublicPtr` {#AccessConfig.setPublicPtr}

| Property | Comments |
|----------|----------|
| Field Name | `setPublicPtr` |
| Type | `bool` |

Specifies whether a public DNS ?PTR? record should be created to map the
external IP address of the instance to a DNS domain name.

### `type` {#AccessConfig.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

The type of configuration. The default and only option is ONE_TO_ONE_NAT.
Valid values:
    ONE_TO_ONE_NAT

## Message `Address` {#Address}

A reserved address resource. (== resource_for beta.addresses ==) (==
resource_for v1.addresses ==) (== resource_for beta.globalAddresses ==) (==
resource_for v1.globalAddresses ==)


### Inputs for `Address`

* `string` [`address`](#Address.address) = 1
* `string` [`addressType`](#Address.addressType) = 2
* `string` [`creationTimestamp`](#Address.creationTimestamp) = 3
* `string` [`description`](#Address.description) = 4
* `string` [`id`](#Address.id) = 5
* `string` [`ipVersion`](#Address.ipVersion) = 6
* `string` [`kind`](#Address.kind) = 7
* `string` [`name`](#Address.name) = 8 (**Required**)
* `string` [`region`](#Address.region) = 9
* `string` [`selfLink`](#Address.selfLink) = 10
* `string` [`status`](#Address.status) = 11
* `string` [`subnetwork`](#Address.subnetwork) = 12
* `repeated` `string` [`users`](#Address.users) = 13

### `address` {#Address.address}

| Property | Comments |
|----------|----------|
| Field Name | `address` |
| Type | `string` |

The static IP address represented by this resource.

### `addressType` {#Address.addressType}

| Property | Comments |
|----------|----------|
| Field Name | `addressType` |
| Type | `string` |

The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified,
defaults to EXTERNAL.
Valid values:
    EXTERNAL
    INTERNAL
    UNSPECIFIED_TYPE

### `creationTimestamp` {#Address.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Address.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#Address.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `ipVersion` {#Address.ipVersion}

| Property | Comments |
|----------|----------|
| Field Name | `ipVersion` |
| Type | `string` |

The IP Version that will be used by this address. Valid options are IPV4 or
IPV6. This can only be specified for a global address.
Valid values:
    IPV4
    IPV6
    UNSPECIFIED_VERSION

### `kind` {#Address.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#address for addresses.

### `name` {#Address.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `region` {#Address.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the regional address resides. This
field is not applicable to global addresses. You must specify this field as
part of the HTTP request URL. You cannot set this field in the request body.

### `selfLink` {#Address.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `status` {#Address.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the address, which can be one of RESERVING,
RESERVED, or IN_USE. An address that is RESERVING is currently in the
process of being reserved. A RESERVED address is currently reserved and
available to use. An IN_USE address is currently being used by another
resource and is not available.
Valid values:
    IN_USE
    RESERVED
    RESERVING

### `subnetwork` {#Address.subnetwork}

| Property | Comments |
|----------|----------|
| Field Name | `subnetwork` |
| Type | `string` |

The URL of the subnetwork in which to reserve the address. If an IP address
is specified, it must be within the subnetwork's IP range. This field can
only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.

### `users` {#Address.users}

| Property | Comments |
|----------|----------|
| Field Name | `users` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] The URLs of the resources that are using this address.

## Message `AddressAggregatedList` {#AddressAggregatedList}



### Inputs for `AddressAggregatedList`

* `string` [`id`](#AddressAggregatedList.id) = 1
* `repeated` [`compute.AddressAggregatedList.ItemsEntry`](gcp_compute.md#AddressAggregatedList.ItemsEntry) [`items`](#AddressAggregatedList.items) = 2
* `string` [`kind`](#AddressAggregatedList.kind) = 3
* `string` [`nextPageToken`](#AddressAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#AddressAggregatedList.selfLink) = 5
* [`compute.AddressAggregatedList.Warning`](gcp_compute.md#AddressAggregatedList.Warning) [`warning`](#AddressAggregatedList.warning) = 6

### `id` {#AddressAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#AddressAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.AddressAggregatedList.ItemsEntry`](gcp_compute.md#AddressAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of AddressesScopedList resources.

### `kind` {#AddressAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#addressAggregatedList for
aggregated lists of addresses.

### `nextPageToken` {#AddressAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#AddressAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#AddressAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AddressAggregatedList.Warning`](gcp_compute.md#AddressAggregatedList.Warning) |

## Message `AddressList` {#AddressList}

Contains a list of addresses.


### Inputs for `AddressList`

* `string` [`id`](#AddressList.id) = 1
* `repeated` [`compute.Address`](gcp_compute.md#Address) [`items`](#AddressList.items) = 2
* `string` [`kind`](#AddressList.kind) = 3
* `string` [`nextPageToken`](#AddressList.nextPageToken) = 4
* `string` [`selfLink`](#AddressList.selfLink) = 5
* [`compute.AddressList.Warning`](gcp_compute.md#AddressList.Warning) [`warning`](#AddressList.warning) = 6

### `id` {#AddressList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#AddressList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Address`](gcp_compute.md#Address) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Address resources.

### `kind` {#AddressList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#addressList for lists of
addresses.

### `nextPageToken` {#AddressList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#AddressList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#AddressList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AddressList.Warning`](gcp_compute.md#AddressList.Warning) |

## Message `AddressesScopedList` {#AddressesScopedList}



### Inputs for `AddressesScopedList`

* `repeated` [`compute.Address`](gcp_compute.md#Address) [`addresses`](#AddressesScopedList.addresses) = 1
* [`compute.AddressesScopedList.Warning`](gcp_compute.md#AddressesScopedList.Warning) [`warning`](#AddressesScopedList.warning) = 2

### `addresses` {#AddressesScopedList.addresses}

| Property | Comments |
|----------|----------|
| Field Name | `addresses` |
| Type | [`compute.Address`](gcp_compute.md#Address) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of addresses contained in this scope.

### `warning` {#AddressesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AddressesScopedList.Warning`](gcp_compute.md#AddressesScopedList.Warning) |

## Message `AliasIpRange` {#AliasIpRange}

An alias IP range attached to an instance's network interface.


### Inputs for `AliasIpRange`

* `string` [`ipCidrRange`](#AliasIpRange.ipCidrRange) = 1
* `string` [`subnetworkRangeName`](#AliasIpRange.subnetworkRangeName) = 2

### `ipCidrRange` {#AliasIpRange.ipCidrRange}

| Property | Comments |
|----------|----------|
| Field Name | `ipCidrRange` |
| Type | `string` |

The IP CIDR range represented by this alias IP range. This IP CIDR range
must belong to the specified subnetwork and cannot contain IP addresses
reserved by system or used by other network interfaces. This range may be a
single IP address (e.g. 10.2.3.4), a netmask (e.g. /24) or a CIDR format
string (e.g. 10.1.2.0/24).

### `subnetworkRangeName` {#AliasIpRange.subnetworkRangeName}

| Property | Comments |
|----------|----------|
| Field Name | `subnetworkRangeName` |
| Type | `string` |

Optional subnetwork secondary range name specifying the secondary range from
which to allocate the IP CIDR range for this alias IP range. If left
unspecified, the primary range of the subnetwork will be used.

## Message `Allowed` {#Firewall.Allowed}

The list of ALLOW rules specified by this firewall. Each rule specifies a
protocol and port-range tuple that describes a permitted connection.


### Inputs for `Allowed`

* `string` [`IPProtocol`](#Firewall.Allowed.IPProtocol) = 1
* `repeated` `string` [`ports`](#Firewall.Allowed.ports) = 2

### `IPProtocol` {#Firewall.Allowed.IPProtocol}

| Property | Comments |
|----------|----------|
| Field Name | `IPProtocol` |
| Type | `string` |

The IP protocol to which this rule applies. The protocol type is required
when creating a firewall rule. This value can either be one of the following
well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP
protocol number.

### `ports` {#Firewall.Allowed.ports}

| Property | Comments |
|----------|----------|
| Field Name | `ports` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

An optional list of ports to which this rule applies. This field is only
applicable for UDP or TCP protocol. Each entry must be either an integer or
a range. If not specified, this rule applies to connections through any
port.

Example inputs include: ["22"], ["80","443"], and ["12345-12349"].

## Message `AttachedDisk` {#AttachedDisk}

An instance-attached disk resource.


### Inputs for `AttachedDisk`

* `bool` [`autoDelete`](#AttachedDisk.autoDelete) = 1
* `bool` [`boot`](#AttachedDisk.boot) = 2
* `string` [`deviceName`](#AttachedDisk.deviceName) = 3
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`diskEncryptionKey`](#AttachedDisk.diskEncryptionKey) = 4
* `repeated` [`compute.GuestOsFeature`](gcp_compute.md#GuestOsFeature) [`guestOsFeatures`](#AttachedDisk.guestOsFeatures) = 5
* `int32` [`index`](#AttachedDisk.index) = 6
* [`compute.AttachedDiskInitializeParams`](gcp_compute.md#AttachedDiskInitializeParams) [`initializeParams`](#AttachedDisk.initializeParams) = 7
* `string` [`interface`](#AttachedDisk.interface) = 8
* `string` [`kind`](#AttachedDisk.kind) = 9
* `repeated` `string` [`licenses`](#AttachedDisk.licenses) = 10
* `string` [`mode`](#AttachedDisk.mode) = 11
* `string` [`source`](#AttachedDisk.source) = 12
* `string` [`type`](#AttachedDisk.type) = 13

### `autoDelete` {#AttachedDisk.autoDelete}

| Property | Comments |
|----------|----------|
| Field Name | `autoDelete` |
| Type | `bool` |

Specifies whether the disk will be auto-deleted when the instance is deleted
(but not when the disk is detached from the instance).

### `boot` {#AttachedDisk.boot}

| Property | Comments |
|----------|----------|
| Field Name | `boot` |
| Type | `bool` |

Indicates that this is a boot disk. The virtual machine will use the first
partition of the disk for its root filesystem.

### `deviceName` {#AttachedDisk.deviceName}

| Property | Comments |
|----------|----------|
| Field Name | `deviceName` |
| Type | `string` |

Specifies a unique device name of your choice that is reflected into the
/dev/disk/by-id/google-* tree of a Linux operating system running within the
instance. This name can be used to reference the device for mounting,
resizing, and so on, from within the instance.

If not specified, the server chooses a default device name to apply to this
disk, in the form persistent-disks-x, where x is a number assigned by Google
Compute Engine. This field is only applicable for persistent disks.

### `diskEncryptionKey` {#AttachedDisk.diskEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `diskEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

Encrypts or decrypts a disk using a customer-supplied encryption key.

If you are creating a new disk, this field encrypts the new disk using an
encryption key that you provide. If you are attaching an existing disk that
is already encrypted, this field decrypts the disk using the
customer-supplied encryption key.

If you encrypt a disk using a customer-supplied key, you must provide the
same key again when you attempt to use this resource at a later time. For
example, you must provide the key when you create a snapshot or an image
from the disk or when you attach the disk to a virtual machine instance.

If you do not provide an encryption key, then the disk will be encrypted
using an automatically generated key and you do not need to provide a key to
use the disk later.

Instance templates do not store customer-supplied encryption keys, so you
cannot use your own keys to encrypt disks in a managed instance group.

### `guestOsFeatures` {#AttachedDisk.guestOsFeatures}

| Property | Comments |
|----------|----------|
| Field Name | `guestOsFeatures` |
| Type | [`compute.GuestOsFeature`](gcp_compute.md#GuestOsFeature) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of features to enable on the guest operating system. Applicable only
for bootable images. Read  Enabling guest operating system features to see a
list of available options.

### `index` {#AttachedDisk.index}

| Property | Comments |
|----------|----------|
| Field Name | `index` |
| Type | `int32` |

[Output Only] A zero-based index to this disk, where 0 is reserved for the
boot disk. If you have many disks attached to an instance, each disk would
have a unique index number.

### `initializeParams` {#AttachedDisk.initializeParams}

| Property | Comments |
|----------|----------|
| Field Name | `initializeParams` |
| Type | [`compute.AttachedDiskInitializeParams`](gcp_compute.md#AttachedDiskInitializeParams) |

[Input Only] Specifies the parameters for a new disk that will be created
alongside the new instance. Use initialization parameters to create boot
disks or local SSDs attached to the new instance.

This property is mutually exclusive with the source property; you can only
define one or the other, but not both.

### `interface` {#AttachedDisk.interface}

| Property | Comments |
|----------|----------|
| Field Name | `interface` |
| Type | `string` |

Specifies the disk interface to use for attaching this disk, which is either
SCSI or NVME. The default is SCSI. Persistent disks must always use SCSI and
the request will fail if you attempt to attach a persistent disk in any
other format than SCSI. Local SSDs can use either NVME or SCSI. For
performance characteristics of SCSI over NVMe, see Local SSD performance.
Valid values:
    NVME
    SCSI

### `kind` {#AttachedDisk.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#attachedDisk for attached
disks.

### `licenses` {#AttachedDisk.licenses}

| Property | Comments |
|----------|----------|
| Field Name | `licenses` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Any valid publicly visible licenses.

### `mode` {#AttachedDisk.mode}

| Property | Comments |
|----------|----------|
| Field Name | `mode` |
| Type | `string` |

The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If
not specified, the default is to attach the disk in READ_WRITE mode.
Valid values:
    READ_ONLY
    READ_WRITE

### `source` {#AttachedDisk.source}

| Property | Comments |
|----------|----------|
| Field Name | `source` |
| Type | `string` |

Specifies a valid partial or full URL to an existing Persistent Disk
resource. When creating a new instance, one of initializeParams.sourceImage
or disks.source is required except for local SSD.

If desired, you can also attach existing non-root persistent disks using
this property. This field is only applicable for persistent disks.

Note that for InstanceTemplate, specify the disk name, not the URL for the
disk.

### `type` {#AttachedDisk.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

Specifies the type of the disk, either SCRATCH or PERSISTENT. If not
specified, the default is PERSISTENT.
Valid values:
    PERSISTENT
    SCRATCH

## Message `AttachedDiskInitializeParams` {#AttachedDiskInitializeParams}

[Input Only] Specifies the parameters for a new disk that will be created
alongside the new instance. Use initialization parameters to create boot
disks or local SSDs attached to the new instance.

This property is mutually exclusive with the source property; you can only
define one or the other, but not both.


### Inputs for `AttachedDiskInitializeParams`

* `string` [`diskName`](#AttachedDiskInitializeParams.diskName) = 1
* `string` [`diskSizeGb`](#AttachedDiskInitializeParams.diskSizeGb) = 2
* `string` [`diskType`](#AttachedDiskInitializeParams.diskType) = 3
* `repeated` [`compute.AttachedDiskInitializeParams.LabelsEntry`](gcp_compute.md#AttachedDiskInitializeParams.LabelsEntry) [`labels`](#AttachedDiskInitializeParams.labels) = 4
* `string` [`sourceImage`](#AttachedDiskInitializeParams.sourceImage) = 5
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceImageEncryptionKey`](#AttachedDiskInitializeParams.sourceImageEncryptionKey) = 6

### `diskName` {#AttachedDiskInitializeParams.diskName}

| Property | Comments |
|----------|----------|
| Field Name | `diskName` |
| Type | `string` |

Specifies the disk name. If not specified, the default is to use the name of
the instance.

### `diskSizeGb` {#AttachedDiskInitializeParams.diskSizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `diskSizeGb` |
| Type | `string` |

Specifies the size of the disk in base-2 GB.

### `diskType` {#AttachedDiskInitializeParams.diskType}

| Property | Comments |
|----------|----------|
| Field Name | `diskType` |
| Type | `string` |

Specifies the disk type to use to create the instance. If not specified, the
default is pd-standard, specified using the full URL. For example:
https://www.googleapis.com/compute/v1/projects/project/zones/zone/diskTypes/pd-standard


Other values include pd-ssd and local-ssd. If you define this field, you can
provide either the full or partial URL. For example, the following are valid
values:
-
https://www.googleapis.com/compute/v1/projects/project/zones/zone/diskTypes/diskType
- projects/project/zones/zone/diskTypes/diskType
- zones/zone/diskTypes/diskType  Note that for InstanceTemplate, this is the
name of the disk type, not URL.

### `labels` {#AttachedDiskInitializeParams.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.AttachedDiskInitializeParams.LabelsEntry`](gcp_compute.md#AttachedDiskInitializeParams.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Labels to apply to this disk. These can be later modified by the
disks.setLabels method. This field is only applicable for persistent disks.

### `sourceImage` {#AttachedDiskInitializeParams.sourceImage}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImage` |
| Type | `string` |

The source image to create this disk. When creating a new instance, one of
initializeParams.sourceImage or disks.source is required except for local
SSD.

To create a disk with one of the public operating system images, specify the
image by its family name. For example, specify family/debian-8 to use the
latest Debian 8 image:
projects/debian-cloud/global/images/family/debian-8


Alternatively, use a specific version of a public operating system image:
projects/debian-cloud/global/images/debian-8-jessie-vYYYYMMDD


To create a disk with a custom image that you created, specify the image
name in the following format:
global/images/my-custom-image


You can also specify a custom image by its image family, which returns the
latest version of the image in that family. Replace the image name with
family/family-name:
global/images/family/my-image-family


If the source image is deleted later, this field will not be set.

### `sourceImageEncryptionKey` {#AttachedDiskInitializeParams.sourceImageEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImageEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source image. Required if the
source image is protected by a customer-supplied encryption key.

Instance templates do not store customer-supplied encryption keys, so you
cannot create disks for instances in a managed instance group if the source
images are encrypted with your own keys.

## Message `Autoscaler` {#Autoscaler}

Represents an Autoscaler resource. Autoscalers allow you to automatically
scale virtual machine instances in managed instance groups according to an
autoscaling policy that you define. For more information, read Autoscaling
Groups of Instances. (== resource_for beta.autoscalers ==) (== resource_for
v1.autoscalers ==) (== resource_for beta.regionAutoscalers ==) (==
resource_for v1.regionAutoscalers ==)


### Inputs for `Autoscaler`

* [`compute.AutoscalingPolicy`](gcp_compute.md#AutoscalingPolicy) [`autoscalingPolicy`](#Autoscaler.autoscalingPolicy) = 1
* `string` [`creationTimestamp`](#Autoscaler.creationTimestamp) = 2
* `string` [`description`](#Autoscaler.description) = 3
* `string` [`id`](#Autoscaler.id) = 4
* `string` [`kind`](#Autoscaler.kind) = 5
* `string` [`name`](#Autoscaler.name) = 6 (**Required**)
* `string` [`region`](#Autoscaler.region) = 7
* `string` [`selfLink`](#Autoscaler.selfLink) = 8
* `string` [`status`](#Autoscaler.status) = 9
* `repeated` [`compute.AutoscalerStatusDetails`](gcp_compute.md#AutoscalerStatusDetails) [`statusDetails`](#Autoscaler.statusDetails) = 10
* `string` [`target`](#Autoscaler.target) = 11
* `string` [`zone`](#Autoscaler.zone) = 12

### `autoscalingPolicy` {#Autoscaler.autoscalingPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `autoscalingPolicy` |
| Type | [`compute.AutoscalingPolicy`](gcp_compute.md#AutoscalingPolicy) |

The configuration parameters for the autoscaling algorithm. You can define
one or more of the policies for an autoscaler: cpuUtilization,
customMetricUtilizations, and loadBalancingUtilization.

If none of these are specified, the default will be to autoscale based on
cpuUtilization to 0.6 or 60%.

### `creationTimestamp` {#Autoscaler.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Autoscaler.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#Autoscaler.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Autoscaler.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#autoscaler for
autoscalers.

### `name` {#Autoscaler.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `region` {#Autoscaler.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the instance group resides (for
autoscalers living in regional scope).

### `selfLink` {#Autoscaler.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `status` {#Autoscaler.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the autoscaler configuration.
Valid values:
    ACTIVE
    DELETING
    ERROR
    PENDING

### `statusDetails` {#Autoscaler.statusDetails}

| Property | Comments |
|----------|----------|
| Field Name | `statusDetails` |
| Type | [`compute.AutoscalerStatusDetails`](gcp_compute.md#AutoscalerStatusDetails) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Human-readable details about the current state of the
autoscaler. Read the documentation for Commonly returned status messages for
examples of status messages you might encounter.

### `target` {#Autoscaler.target}

| Property | Comments |
|----------|----------|
| Field Name | `target` |
| Type | `string` |

URL of the managed instance group that this autoscaler will scale.

### `zone` {#Autoscaler.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] URL of the zone where the instance group resides (for
autoscalers living in zonal scope).

## Message `AutoscalerAggregatedList` {#AutoscalerAggregatedList}



### Inputs for `AutoscalerAggregatedList`

* `string` [`id`](#AutoscalerAggregatedList.id) = 1
* `repeated` [`compute.AutoscalerAggregatedList.ItemsEntry`](gcp_compute.md#AutoscalerAggregatedList.ItemsEntry) [`items`](#AutoscalerAggregatedList.items) = 2
* `string` [`kind`](#AutoscalerAggregatedList.kind) = 3
* `string` [`nextPageToken`](#AutoscalerAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#AutoscalerAggregatedList.selfLink) = 5
* [`compute.AutoscalerAggregatedList.Warning`](gcp_compute.md#AutoscalerAggregatedList.Warning) [`warning`](#AutoscalerAggregatedList.warning) = 6

### `id` {#AutoscalerAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#AutoscalerAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.AutoscalerAggregatedList.ItemsEntry`](gcp_compute.md#AutoscalerAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of AutoscalersScopedList resources.

### `kind` {#AutoscalerAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#autoscalerAggregatedList for
aggregated lists of autoscalers.

### `nextPageToken` {#AutoscalerAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#AutoscalerAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#AutoscalerAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AutoscalerAggregatedList.Warning`](gcp_compute.md#AutoscalerAggregatedList.Warning) |

## Message `AutoscalerList` {#AutoscalerList}

Contains a list of Autoscaler resources.


### Inputs for `AutoscalerList`

* `string` [`id`](#AutoscalerList.id) = 1
* `repeated` [`compute.Autoscaler`](gcp_compute.md#Autoscaler) [`items`](#AutoscalerList.items) = 2
* `string` [`kind`](#AutoscalerList.kind) = 3
* `string` [`nextPageToken`](#AutoscalerList.nextPageToken) = 4
* `string` [`selfLink`](#AutoscalerList.selfLink) = 5
* [`compute.AutoscalerList.Warning`](gcp_compute.md#AutoscalerList.Warning) [`warning`](#AutoscalerList.warning) = 6

### `id` {#AutoscalerList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#AutoscalerList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Autoscaler`](gcp_compute.md#Autoscaler) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Autoscaler resources.

### `kind` {#AutoscalerList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#autoscalerList for lists of
autoscalers.

### `nextPageToken` {#AutoscalerList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#AutoscalerList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#AutoscalerList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AutoscalerList.Warning`](gcp_compute.md#AutoscalerList.Warning) |

## Message `AutoscalerStatusDetails` {#AutoscalerStatusDetails}



### Inputs for `AutoscalerStatusDetails`

* `string` [`message`](#AutoscalerStatusDetails.message) = 1
* `string` [`type`](#AutoscalerStatusDetails.type) = 2

### `message` {#AutoscalerStatusDetails.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

The status message.

### `type` {#AutoscalerStatusDetails.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

The type of error returned.
Valid values:
    ALL_INSTANCES_UNHEALTHY
    BACKEND_SERVICE_DOES_NOT_EXIST
    CAPPED_AT_MAX_NUM_REPLICAS
    CUSTOM_METRIC_DATA_POINTS_TOO_SPARSE
    CUSTOM_METRIC_INVALID
    MIN_EQUALS_MAX
    MISSING_CUSTOM_METRIC_DATA_POINTS
    MISSING_LOAD_BALANCING_DATA_POINTS
    MORE_THAN_ONE_BACKEND_SERVICE
    NOT_ENOUGH_QUOTA_AVAILABLE
    REGION_RESOURCE_STOCKOUT
    SCALING_TARGET_DOES_NOT_EXIST
    UNKNOWN
    UNSUPPORTED_MAX_RATE_LOAD_BALANCING_CONFIGURATION
    ZONE_RESOURCE_STOCKOUT

## Message `AutoscalersScopedList` {#AutoscalersScopedList}



### Inputs for `AutoscalersScopedList`

* `repeated` [`compute.Autoscaler`](gcp_compute.md#Autoscaler) [`autoscalers`](#AutoscalersScopedList.autoscalers) = 1
* [`compute.AutoscalersScopedList.Warning`](gcp_compute.md#AutoscalersScopedList.Warning) [`warning`](#AutoscalersScopedList.warning) = 2

### `autoscalers` {#AutoscalersScopedList.autoscalers}

| Property | Comments |
|----------|----------|
| Field Name | `autoscalers` |
| Type | [`compute.Autoscaler`](gcp_compute.md#Autoscaler) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of autoscalers contained in this scope.

### `warning` {#AutoscalersScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.AutoscalersScopedList.Warning`](gcp_compute.md#AutoscalersScopedList.Warning) |

## Message `AutoscalingPolicy` {#AutoscalingPolicy}

Cloud Autoscaler policy.


### Inputs for `AutoscalingPolicy`

* `int32` [`coolDownPeriodSec`](#AutoscalingPolicy.coolDownPeriodSec) = 1
* [`compute.AutoscalingPolicyCpuUtilization`](gcp_compute.md#AutoscalingPolicyCpuUtilization) [`cpuUtilization`](#AutoscalingPolicy.cpuUtilization) = 2
* `repeated` [`compute.AutoscalingPolicyCustomMetricUtilization`](gcp_compute.md#AutoscalingPolicyCustomMetricUtilization) [`customMetricUtilizations`](#AutoscalingPolicy.customMetricUtilizations) = 3
* [`compute.AutoscalingPolicyLoadBalancingUtilization`](gcp_compute.md#AutoscalingPolicyLoadBalancingUtilization) [`loadBalancingUtilization`](#AutoscalingPolicy.loadBalancingUtilization) = 4
* `int32` [`maxNumReplicas`](#AutoscalingPolicy.maxNumReplicas) = 5
* `int32` [`minNumReplicas`](#AutoscalingPolicy.minNumReplicas) = 6

### `coolDownPeriodSec` {#AutoscalingPolicy.coolDownPeriodSec}

| Property | Comments |
|----------|----------|
| Field Name | `coolDownPeriodSec` |
| Type | `int32` |

The number of seconds that the autoscaler should wait before it starts
collecting information from a new instance. This prevents the autoscaler
from collecting information when the instance is initializing, during which
the collected usage would not be reliable. The default time autoscaler waits
is 60 seconds.

Virtual machine initialization times might vary because of numerous factors.
We recommend that you test how long an instance may take to initialize. To
do this, create an instance and time the startup process.

### `cpuUtilization` {#AutoscalingPolicy.cpuUtilization}

| Property | Comments |
|----------|----------|
| Field Name | `cpuUtilization` |
| Type | [`compute.AutoscalingPolicyCpuUtilization`](gcp_compute.md#AutoscalingPolicyCpuUtilization) |

Defines the CPU utilization policy that allows the autoscaler to scale based
on the average CPU utilization of a managed instance group.

### `customMetricUtilizations` {#AutoscalingPolicy.customMetricUtilizations}

| Property | Comments |
|----------|----------|
| Field Name | `customMetricUtilizations` |
| Type | [`compute.AutoscalingPolicyCustomMetricUtilization`](gcp_compute.md#AutoscalingPolicyCustomMetricUtilization) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Configuration parameters of autoscaling based on a custom metric.

### `loadBalancingUtilization` {#AutoscalingPolicy.loadBalancingUtilization}

| Property | Comments |
|----------|----------|
| Field Name | `loadBalancingUtilization` |
| Type | [`compute.AutoscalingPolicyLoadBalancingUtilization`](gcp_compute.md#AutoscalingPolicyLoadBalancingUtilization) |

Configuration parameters of autoscaling based on load balancer.

### `maxNumReplicas` {#AutoscalingPolicy.maxNumReplicas}

| Property | Comments |
|----------|----------|
| Field Name | `maxNumReplicas` |
| Type | `int32` |

The maximum number of instances that the autoscaler can scale up to. This is
required when creating or updating an autoscaler. The maximum number of
replicas should not be lower than minimal number of replicas.

### `minNumReplicas` {#AutoscalingPolicy.minNumReplicas}

| Property | Comments |
|----------|----------|
| Field Name | `minNumReplicas` |
| Type | `int32` |

The minimum number of replicas that the autoscaler can scale down to. This
cannot be less than 0. If not provided, autoscaler will choose a default
value depending on maximum number of instances allowed.

## Message `AutoscalingPolicyCpuUtilization` {#AutoscalingPolicyCpuUtilization}

CPU utilization policy.


### Inputs for `AutoscalingPolicyCpuUtilization`

* `TYPE_DOUBLE` [`utilizationTarget`](#AutoscalingPolicyCpuUtilization.utilizationTarget) = 1

### `utilizationTarget` {#AutoscalingPolicyCpuUtilization.utilizationTarget}

| Property | Comments |
|----------|----------|
| Field Name | `utilizationTarget` |
| Type | `TYPE_DOUBLE` |

The target CPU utilization that the autoscaler should maintain. Must be a
float value in the range (0, 1]. If not specified, the default is 0.6.

If the CPU level is below the target utilization, the autoscaler scales down
the number of instances until it reaches the minimum number of instances you
specified or until the average CPU of your instances reaches the target
utilization.

If the average CPU is above the target utilization, the autoscaler scales up
until it reaches the maximum number of instances you specified or until the
average utilization reaches the target utilization.

## Message `AutoscalingPolicyCustomMetricUtilization` {#AutoscalingPolicyCustomMetricUtilization}

Custom utilization metric policy.


### Inputs for `AutoscalingPolicyCustomMetricUtilization`

* `string` [`metric`](#AutoscalingPolicyCustomMetricUtilization.metric) = 1
* `TYPE_DOUBLE` [`utilizationTarget`](#AutoscalingPolicyCustomMetricUtilization.utilizationTarget) = 2
* `string` [`utilizationTargetType`](#AutoscalingPolicyCustomMetricUtilization.utilizationTargetType) = 3

### `metric` {#AutoscalingPolicyCustomMetricUtilization.metric}

| Property | Comments |
|----------|----------|
| Field Name | `metric` |
| Type | `string` |

The identifier (type) of the Stackdriver Monitoring metric. The metric
cannot have negative values.

The metric must have a value type of INT64 or DOUBLE.

### `utilizationTarget` {#AutoscalingPolicyCustomMetricUtilization.utilizationTarget}

| Property | Comments |
|----------|----------|
| Field Name | `utilizationTarget` |
| Type | `TYPE_DOUBLE` |

The target value of the metric that autoscaler should maintain. This must be
a positive value. A utilization metric scales number of virtual machines
handling requests to increase or decrease proportionally to the metric.

For example, a good metric to use as a utilization_target is
compute.googleapis.com/instance/network/received_bytes_count. The autoscaler
will work to keep this value constant for each of the instances.

### `utilizationTargetType` {#AutoscalingPolicyCustomMetricUtilization.utilizationTargetType}

| Property | Comments |
|----------|----------|
| Field Name | `utilizationTargetType` |
| Type | `string` |

Defines how target utilization value is expressed for a Stackdriver
Monitoring metric. Either GAUGE, DELTA_PER_SECOND, or DELTA_PER_MINUTE. If
not specified, the default is GAUGE.
Valid values:
    DELTA_PER_MINUTE
    DELTA_PER_SECOND
    GAUGE

## Message `AutoscalingPolicyLoadBalancingUtilization` {#AutoscalingPolicyLoadBalancingUtilization}

Configuration parameters of autoscaling based on load balancing.


### Inputs for `AutoscalingPolicyLoadBalancingUtilization`

* `TYPE_DOUBLE` [`utilizationTarget`](#AutoscalingPolicyLoadBalancingUtilization.utilizationTarget) = 1

### `utilizationTarget` {#AutoscalingPolicyLoadBalancingUtilization.utilizationTarget}

| Property | Comments |
|----------|----------|
| Field Name | `utilizationTarget` |
| Type | `TYPE_DOUBLE` |

Fraction of backend capacity utilization (set in HTTP(s) load balancing
configuration) that autoscaler should maintain. Must be a positive float
value. If not defined, the default is 0.8.

## Message `Backend` {#Backend}

Message containing information of one individual backend.


### Inputs for `Backend`

* `string` [`balancingMode`](#Backend.balancingMode) = 1
* `TYPE_DOUBLE` [`capacityScaler`](#Backend.capacityScaler) = 2
* `string` [`description`](#Backend.description) = 3
* `string` [`group`](#Backend.group) = 4
* `int32` [`maxConnections`](#Backend.maxConnections) = 5
* `int32` [`maxConnectionsPerInstance`](#Backend.maxConnectionsPerInstance) = 6
* `int32` [`maxRate`](#Backend.maxRate) = 7
* `TYPE_DOUBLE` [`maxRatePerInstance`](#Backend.maxRatePerInstance) = 8
* `TYPE_DOUBLE` [`maxUtilization`](#Backend.maxUtilization) = 9

### `balancingMode` {#Backend.balancingMode}

| Property | Comments |
|----------|----------|
| Field Name | `balancingMode` |
| Type | `string` |

Specifies the balancing mode for this backend. For global HTTP(S) or TCP/SSL
load balancing, the default is UTILIZATION. Valid values are UTILIZATION,
RATE (for HTTP(S)) and CONNECTION (for TCP/SSL).

For Internal Load Balancing, the default and only supported mode is
CONNECTION.
Valid values:
    CONNECTION
    RATE
    UTILIZATION

### `capacityScaler` {#Backend.capacityScaler}

| Property | Comments |
|----------|----------|
| Field Name | `capacityScaler` |
| Type | `TYPE_DOUBLE` |

A multiplier applied to the group's maximum servicing capacity (based on
UTILIZATION, RATE or CONNECTION). Default value is 1, which means the group
will serve up to 100% of its configured capacity (depending on
balancingMode). A setting of 0 means the group is completely drained,
offering 0% of its available Capacity. Valid range is [0.0,1.0].

This cannot be used for internal load balancing.

### `description` {#Backend.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `group` {#Backend.group}

| Property | Comments |
|----------|----------|
| Field Name | `group` |
| Type | `string` |

The fully-qualified URL of a Instance Group resource. This instance group
defines the list of instances that serve traffic. Member virtual machine
instances from each instance group must live in the same zone as the
instance group itself. No two backends in a backend service are allowed to
use same Instance Group resource.

Note that you must specify an Instance Group resource using the
fully-qualified URL, rather than a partial URL.

When the BackendService has load balancing scheme INTERNAL, the instance
group must be within the same region as the BackendService.

### `maxConnections` {#Backend.maxConnections}

| Property | Comments |
|----------|----------|
| Field Name | `maxConnections` |
| Type | `int32` |

The max number of simultaneous connections for the group. Can be used with
either CONNECTION or UTILIZATION balancing modes. For CONNECTION mode,
either maxConnections or maxConnectionsPerInstance must be set.

This cannot be used for internal load balancing.

### `maxConnectionsPerInstance` {#Backend.maxConnectionsPerInstance}

| Property | Comments |
|----------|----------|
| Field Name | `maxConnectionsPerInstance` |
| Type | `int32` |

The max number of simultaneous connections that a single backend instance
can handle. This is used to calculate the capacity of the group. Can be used
in either CONNECTION or UTILIZATION balancing modes. For CONNECTION mode,
either maxConnections or maxConnectionsPerInstance must be set.

This cannot be used for internal load balancing.

### `maxRate` {#Backend.maxRate}

| Property | Comments |
|----------|----------|
| Field Name | `maxRate` |
| Type | `int32` |

The max requests per second (RPS) of the group. Can be used with either RATE
or UTILIZATION balancing modes, but required if RATE mode. For RATE mode,
either maxRate or maxRatePerInstance must be set.

This cannot be used for internal load balancing.

### `maxRatePerInstance` {#Backend.maxRatePerInstance}

| Property | Comments |
|----------|----------|
| Field Name | `maxRatePerInstance` |
| Type | `TYPE_DOUBLE` |

The max requests per second (RPS) that a single backend instance can handle.
This is used to calculate the capacity of the group. Can be used in either
balancing mode. For RATE mode, either maxRate or maxRatePerInstance must be
set.

This cannot be used for internal load balancing.

### `maxUtilization` {#Backend.maxUtilization}

| Property | Comments |
|----------|----------|
| Field Name | `maxUtilization` |
| Type | `TYPE_DOUBLE` |

Used when balancingMode is UTILIZATION. This ratio defines the CPU
utilization target for the group. The default is 0.8. Valid range is [0.0,
1.0].

This cannot be used for internal load balancing.

## Message `BackendBucket` {#BackendBucket}

A BackendBucket resource. This resource defines a Cloud Storage bucket.


### Inputs for `BackendBucket`

* `string` [`bucketName`](#BackendBucket.bucketName) = 1
* `string` [`creationTimestamp`](#BackendBucket.creationTimestamp) = 2
* `string` [`description`](#BackendBucket.description) = 3
* `bool` [`enableCdn`](#BackendBucket.enableCdn) = 4
* `string` [`id`](#BackendBucket.id) = 5
* `string` [`kind`](#BackendBucket.kind) = 6
* `string` [`name`](#BackendBucket.name) = 7 (**Required**)
* `string` [`selfLink`](#BackendBucket.selfLink) = 8

### `bucketName` {#BackendBucket.bucketName}

| Property | Comments |
|----------|----------|
| Field Name | `bucketName` |
| Type | `string` |

Cloud Storage bucket name.

### `creationTimestamp` {#BackendBucket.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#BackendBucket.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional textual description of the resource; provided by the client when
the resource is created.

### `enableCdn` {#BackendBucket.enableCdn}

| Property | Comments |
|----------|----------|
| Field Name | `enableCdn` |
| Type | `bool` |

If true, enable Cloud CDN for this BackendBucket.

### `id` {#BackendBucket.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `kind` {#BackendBucket.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of the resource.

### `name` {#BackendBucket.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `selfLink` {#BackendBucket.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

## Message `BackendBucketList` {#BackendBucketList}

Contains a list of BackendBucket resources.


### Inputs for `BackendBucketList`

* `string` [`id`](#BackendBucketList.id) = 1
* `repeated` [`compute.BackendBucket`](gcp_compute.md#BackendBucket) [`items`](#BackendBucketList.items) = 2
* `string` [`kind`](#BackendBucketList.kind) = 3
* `string` [`nextPageToken`](#BackendBucketList.nextPageToken) = 4
* `string` [`selfLink`](#BackendBucketList.selfLink) = 5
* [`compute.BackendBucketList.Warning`](gcp_compute.md#BackendBucketList.Warning) [`warning`](#BackendBucketList.warning) = 6

### `id` {#BackendBucketList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#BackendBucketList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.BackendBucket`](gcp_compute.md#BackendBucket) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of BackendBucket resources.

### `kind` {#BackendBucketList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#BackendBucketList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#BackendBucketList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#BackendBucketList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.BackendBucketList.Warning`](gcp_compute.md#BackendBucketList.Warning) |

## Message `BackendService` {#BackendService}

A BackendService resource. This resource defines a group of backend virtual
machines and their serving capacity. (== resource_for v1.backendService ==)
(== resource_for beta.backendService ==)


### Inputs for `BackendService`

* `int32` [`affinityCookieTtlSec`](#BackendService.affinityCookieTtlSec) = 1
* `repeated` [`compute.Backend`](gcp_compute.md#Backend) [`backends`](#BackendService.backends) = 2
* [`compute.BackendServiceCdnPolicy`](gcp_compute.md#BackendServiceCdnPolicy) [`cdnPolicy`](#BackendService.cdnPolicy) = 3
* [`compute.ConnectionDraining`](gcp_compute.md#ConnectionDraining) [`connectionDraining`](#BackendService.connectionDraining) = 4
* `string` [`creationTimestamp`](#BackendService.creationTimestamp) = 5
* `string` [`description`](#BackendService.description) = 6
* `bool` [`enableCDN`](#BackendService.enableCDN) = 7
* `string` [`fingerprint`](#BackendService.fingerprint) = 8
* `repeated` `string` [`healthChecks`](#BackendService.healthChecks) = 9
* [`compute.BackendServiceIAP`](gcp_compute.md#BackendServiceIAP) [`iap`](#BackendService.iap) = 10
* `string` [`id`](#BackendService.id) = 11
* `string` [`kind`](#BackendService.kind) = 12
* `string` [`loadBalancingScheme`](#BackendService.loadBalancingScheme) = 13
* `string` [`name`](#BackendService.name) = 14 (**Required**)
* `int32` [`port`](#BackendService.port) = 15
* `string` [`portName`](#BackendService.portName) = 16
* `string` [`protocol`](#BackendService.protocol) = 17
* `string` [`region`](#BackendService.region) = 18
* `string` [`selfLink`](#BackendService.selfLink) = 19
* `string` [`sessionAffinity`](#BackendService.sessionAffinity) = 20
* `int32` [`timeoutSec`](#BackendService.timeoutSec) = 21

### `affinityCookieTtlSec` {#BackendService.affinityCookieTtlSec}

| Property | Comments |
|----------|----------|
| Field Name | `affinityCookieTtlSec` |
| Type | `int32` |

Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If
set to 0, the cookie is non-persistent and lasts only until the end of the
browser session (or equivalent). The maximum allowed value for TTL is one
day.

When the load balancing scheme is INTERNAL, this field is not used.

### `backends` {#BackendService.backends}

| Property | Comments |
|----------|----------|
| Field Name | `backends` |
| Type | [`compute.Backend`](gcp_compute.md#Backend) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of backends that serve this BackendService.

### `cdnPolicy` {#BackendService.cdnPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `cdnPolicy` |
| Type | [`compute.BackendServiceCdnPolicy`](gcp_compute.md#BackendServiceCdnPolicy) |

Cloud CDN configuration for this BackendService.

### `connectionDraining` {#BackendService.connectionDraining}

| Property | Comments |
|----------|----------|
| Field Name | `connectionDraining` |
| Type | [`compute.ConnectionDraining`](gcp_compute.md#ConnectionDraining) |

### `creationTimestamp` {#BackendService.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#BackendService.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `enableCDN` {#BackendService.enableCDN}

| Property | Comments |
|----------|----------|
| Field Name | `enableCDN` |
| Type | `bool` |

If true, enable Cloud CDN for this BackendService.

When the load balancing scheme is INTERNAL, this field is not used.

### `fingerprint` {#BackendService.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint of this resource. A hash of the contents stored in this object.
This field is used in optimistic locking. This field will be ignored when
inserting a BackendService. An up-to-date fingerprint must be provided in
order to update the BackendService.

### `healthChecks` {#BackendService.healthChecks}

| Property | Comments |
|----------|----------|
| Field Name | `healthChecks` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of URLs to the HttpHealthCheck or HttpsHealthCheck resource for
health checking this BackendService. Currently at most one health check can
be specified, and a health check is required for Compute Engine backend
services. A health check must not be specified for App Engine backend and
Cloud Function backend.

For internal load balancing, a URL to a HealthCheck resource must be
specified instead.

### `iap` {#BackendService.iap}

| Property | Comments |
|----------|----------|
| Field Name | `iap` |
| Type | [`compute.BackendServiceIAP`](gcp_compute.md#BackendServiceIAP) |

### `id` {#BackendService.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#BackendService.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#backendService for backend
services.

### `loadBalancingScheme` {#BackendService.loadBalancingScheme}

| Property | Comments |
|----------|----------|
| Field Name | `loadBalancingScheme` |
| Type | `string` |

Indicates whether the backend service will be used with internal or external
load balancing. A backend service created for one type of load balancing
cannot be used with the other. Possible values are INTERNAL and EXTERNAL.
Valid values:
    EXTERNAL
    INTERNAL
    INVALID_LOAD_BALANCING_SCHEME

### `name` {#BackendService.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `port` {#BackendService.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

Deprecated in favor of portName. The TCP port to connect on the backend. The
default value is 80.

This cannot be used for internal load balancing.

### `portName` {#BackendService.portName}

| Property | Comments |
|----------|----------|
| Field Name | `portName` |
| Type | `string` |

Name of backend port. The same name should appear in the instance groups
referenced by this service. Required when the load balancing scheme is
EXTERNAL.

When the load balancing scheme is INTERNAL, this field is not used.

### `protocol` {#BackendService.protocol}

| Property | Comments |
|----------|----------|
| Field Name | `protocol` |
| Type | `string` |

The protocol this BackendService uses to communicate with backends.

Possible values are HTTP, HTTPS, TCP, and SSL. The default is HTTP.

For internal load balancing, the possible values are TCP and UDP, and the
default is TCP.
Valid values:
    HTTP
    HTTPS
    SSL
    TCP
    UDP

### `region` {#BackendService.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the regional backend service resides.
This field is not applicable to global backend services. You must specify
this field as part of the HTTP request URL. It is not settable as a field in
the request body.

### `selfLink` {#BackendService.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sessionAffinity` {#BackendService.sessionAffinity}

| Property | Comments |
|----------|----------|
| Field Name | `sessionAffinity` |
| Type | `string` |

Type of session affinity to use. The default is NONE.

When the load balancing scheme is EXTERNAL, can be NONE, CLIENT_IP, or
GENERATED_COOKIE.

When the load balancing scheme is INTERNAL, can be NONE, CLIENT_IP,
CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO.

When the protocol is UDP, this field is not used.
Valid values:
    CLIENT_IP
    CLIENT_IP_PORT_PROTO
    CLIENT_IP_PROTO
    GENERATED_COOKIE
    NONE

### `timeoutSec` {#BackendService.timeoutSec}

| Property | Comments |
|----------|----------|
| Field Name | `timeoutSec` |
| Type | `int32` |

How many seconds to wait for the backend before considering it a failed
request. Default is 30 seconds.

## Message `BackendServiceAggregatedList` {#BackendServiceAggregatedList}

Contains a list of BackendServicesScopedList.


### Inputs for `BackendServiceAggregatedList`

* `string` [`id`](#BackendServiceAggregatedList.id) = 1
* `repeated` [`compute.BackendServiceAggregatedList.ItemsEntry`](gcp_compute.md#BackendServiceAggregatedList.ItemsEntry) [`items`](#BackendServiceAggregatedList.items) = 2
* `string` [`kind`](#BackendServiceAggregatedList.kind) = 3
* `string` [`nextPageToken`](#BackendServiceAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#BackendServiceAggregatedList.selfLink) = 5
* [`compute.BackendServiceAggregatedList.Warning`](gcp_compute.md#BackendServiceAggregatedList.Warning) [`warning`](#BackendServiceAggregatedList.warning) = 6

### `id` {#BackendServiceAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#BackendServiceAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.BackendServiceAggregatedList.ItemsEntry`](gcp_compute.md#BackendServiceAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of BackendServicesScopedList resources.

### `kind` {#BackendServiceAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#BackendServiceAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#BackendServiceAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#BackendServiceAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.BackendServiceAggregatedList.Warning`](gcp_compute.md#BackendServiceAggregatedList.Warning) |

## Message `BackendServiceCdnPolicy` {#BackendServiceCdnPolicy}

Message containing Cloud CDN configuration for a backend service.


### Inputs for `BackendServiceCdnPolicy`

* [`compute.CacheKeyPolicy`](gcp_compute.md#CacheKeyPolicy) [`cacheKeyPolicy`](#BackendServiceCdnPolicy.cacheKeyPolicy) = 1

### `cacheKeyPolicy` {#BackendServiceCdnPolicy.cacheKeyPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `cacheKeyPolicy` |
| Type | [`compute.CacheKeyPolicy`](gcp_compute.md#CacheKeyPolicy) |

The CacheKeyPolicy for this CdnPolicy.

## Message `BackendServiceGroupHealth` {#BackendServiceGroupHealth}



### Inputs for `BackendServiceGroupHealth`

* `repeated` [`compute.HealthStatus`](gcp_compute.md#HealthStatus) [`healthStatus`](#BackendServiceGroupHealth.healthStatus) = 1
* `string` [`kind`](#BackendServiceGroupHealth.kind) = 2

### `healthStatus` {#BackendServiceGroupHealth.healthStatus}

| Property | Comments |
|----------|----------|
| Field Name | `healthStatus` |
| Type | [`compute.HealthStatus`](gcp_compute.md#HealthStatus) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `kind` {#BackendServiceGroupHealth.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#backendServiceGroupHealth for
the health of backend services.

## Message `BackendServiceIAP` {#BackendServiceIAP}

Identity-Aware Proxy


### Inputs for `BackendServiceIAP`

* `bool` [`enabled`](#BackendServiceIAP.enabled) = 1
* `string` [`oauth2ClientId`](#BackendServiceIAP.oauth2ClientId) = 2
* `string` [`oauth2ClientSecret`](#BackendServiceIAP.oauth2ClientSecret) = 3
* `string` [`oauth2ClientSecretSha256`](#BackendServiceIAP.oauth2ClientSecretSha256) = 4

### `enabled` {#BackendServiceIAP.enabled}

| Property | Comments |
|----------|----------|
| Field Name | `enabled` |
| Type | `bool` |

### `oauth2ClientId` {#BackendServiceIAP.oauth2ClientId}

| Property | Comments |
|----------|----------|
| Field Name | `oauth2ClientId` |
| Type | `string` |

### `oauth2ClientSecret` {#BackendServiceIAP.oauth2ClientSecret}

| Property | Comments |
|----------|----------|
| Field Name | `oauth2ClientSecret` |
| Type | `string` |

### `oauth2ClientSecretSha256` {#BackendServiceIAP.oauth2ClientSecretSha256}

| Property | Comments |
|----------|----------|
| Field Name | `oauth2ClientSecretSha256` |
| Type | `string` |

[Output Only] SHA256 hash value for the field oauth2_client_secret above.

## Message `BackendServiceList` {#BackendServiceList}

Contains a list of BackendService resources.


### Inputs for `BackendServiceList`

* `string` [`id`](#BackendServiceList.id) = 1
* `repeated` [`compute.BackendService`](gcp_compute.md#BackendService) [`items`](#BackendServiceList.items) = 2
* `string` [`kind`](#BackendServiceList.kind) = 3
* `string` [`nextPageToken`](#BackendServiceList.nextPageToken) = 4
* `string` [`selfLink`](#BackendServiceList.selfLink) = 5
* [`compute.BackendServiceList.Warning`](gcp_compute.md#BackendServiceList.Warning) [`warning`](#BackendServiceList.warning) = 6

### `id` {#BackendServiceList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#BackendServiceList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.BackendService`](gcp_compute.md#BackendService) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of BackendService resources.

### `kind` {#BackendServiceList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#backendServiceList for lists
of backend services.

### `nextPageToken` {#BackendServiceList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#BackendServiceList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#BackendServiceList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.BackendServiceList.Warning`](gcp_compute.md#BackendServiceList.Warning) |

## Message `BackendServicesScopedList` {#BackendServicesScopedList}



### Inputs for `BackendServicesScopedList`

* `repeated` [`compute.BackendService`](gcp_compute.md#BackendService) [`backendServices`](#BackendServicesScopedList.backendServices) = 1
* [`compute.BackendServicesScopedList.Warning`](gcp_compute.md#BackendServicesScopedList.Warning) [`warning`](#BackendServicesScopedList.warning) = 2

### `backendServices` {#BackendServicesScopedList.backendServices}

| Property | Comments |
|----------|----------|
| Field Name | `backendServices` |
| Type | [`compute.BackendService`](gcp_compute.md#BackendService) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of BackendServices contained in this scope.

### `warning` {#BackendServicesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.BackendServicesScopedList.Warning`](gcp_compute.md#BackendServicesScopedList.Warning) |

## Message `CacheInvalidationRule` {#CacheInvalidationRule}



### Inputs for `CacheInvalidationRule`

* `string` [`host`](#CacheInvalidationRule.host) = 1
* `string` [`path`](#CacheInvalidationRule.path) = 2

### `host` {#CacheInvalidationRule.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

If set, this invalidation rule will only apply to requests with a Host
header matching host.

### `path` {#CacheInvalidationRule.path}

| Property | Comments |
|----------|----------|
| Field Name | `path` |
| Type | `string` |

## Message `CacheKeyPolicy` {#CacheKeyPolicy}

Message containing what to include in the cache key for a request for Cloud
CDN.


### Inputs for `CacheKeyPolicy`

* `bool` [`includeHost`](#CacheKeyPolicy.includeHost) = 1
* `bool` [`includeProtocol`](#CacheKeyPolicy.includeProtocol) = 2
* `bool` [`includeQueryString`](#CacheKeyPolicy.includeQueryString) = 3
* `repeated` `string` [`queryStringBlacklist`](#CacheKeyPolicy.queryStringBlacklist) = 4
* `repeated` `string` [`queryStringWhitelist`](#CacheKeyPolicy.queryStringWhitelist) = 5

### `includeHost` {#CacheKeyPolicy.includeHost}

| Property | Comments |
|----------|----------|
| Field Name | `includeHost` |
| Type | `bool` |

If true, requests to different hosts will be cached separately.

### `includeProtocol` {#CacheKeyPolicy.includeProtocol}

| Property | Comments |
|----------|----------|
| Field Name | `includeProtocol` |
| Type | `bool` |

If true, http and https requests will be cached separately.

### `includeQueryString` {#CacheKeyPolicy.includeQueryString}

| Property | Comments |
|----------|----------|
| Field Name | `includeQueryString` |
| Type | `bool` |

If true, include query string parameters in the cache key according to
query_string_whitelist and query_string_blacklist. If neither is set, the
entire query string will be included. If false, the query string will be
excluded from the cache key entirely.

### `queryStringBlacklist` {#CacheKeyPolicy.queryStringBlacklist}

| Property | Comments |
|----------|----------|
| Field Name | `queryStringBlacklist` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Names of query string parameters to exclude in cache keys. All other
parameters will be included. Either specify query_string_whitelist or
query_string_blacklist, not both. '&' and '=' will be percent encoded and
not treated as delimiters.

### `queryStringWhitelist` {#CacheKeyPolicy.queryStringWhitelist}

| Property | Comments |
|----------|----------|
| Field Name | `queryStringWhitelist` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Names of query string parameters to include in cache keys. All other
parameters will be excluded. Either specify query_string_whitelist or
query_string_blacklist, not both. '&' and '=' will be percent encoded and
not treated as delimiters.

## Message `Commitment` {#Commitment}

Represents a Commitment resource. Creating a Commitment resource means that
you are purchasing a committed use contract with an explicit start and end
time. You can create commitments based on vCPUs and memory usage and receive
discounted rates. For full details, read Signing Up for Committed Use
Discounts.

Committed use discounts are subject to Google Cloud Platform's Service
Specific Terms. By purchasing a committed use discount, you agree to these
terms. Committed use discounts will not renew, so you must purchase a new
commitment to continue receiving discounts. (== resource_for
beta.commitments ==) (== resource_for v1.commitments ==)


### Inputs for `Commitment`

* `string` [`creationTimestamp`](#Commitment.creationTimestamp) = 1
* `string` [`description`](#Commitment.description) = 2
* `string` [`endTimestamp`](#Commitment.endTimestamp) = 3
* `string` [`id`](#Commitment.id) = 4
* `string` [`kind`](#Commitment.kind) = 5
* `string` [`name`](#Commitment.name) = 6 (**Required**)
* `string` [`plan`](#Commitment.plan) = 7
* `string` [`region`](#Commitment.region) = 8
* `repeated` [`compute.ResourceCommitment`](gcp_compute.md#ResourceCommitment) [`resources`](#Commitment.resources) = 9
* `string` [`selfLink`](#Commitment.selfLink) = 10
* `string` [`startTimestamp`](#Commitment.startTimestamp) = 11
* `string` [`status`](#Commitment.status) = 12
* `string` [`statusMessage`](#Commitment.statusMessage) = 13

### `creationTimestamp` {#Commitment.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Commitment.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `endTimestamp` {#Commitment.endTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `endTimestamp` |
| Type | `string` |

[Output Only] Commitment end time in RFC3339 text format.

### `id` {#Commitment.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Commitment.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#commitment for
commitments.

### `name` {#Commitment.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `plan` {#Commitment.plan}

| Property | Comments |
|----------|----------|
| Field Name | `plan` |
| Type | `string` |

The plan for this commitment, which determines duration and discount rate.
The currently supported plans are TWELVE_MONTH (1 year), and
THIRTY_SIX_MONTH (3 years).
Valid values:
    INVALID
    THIRTY_SIX_MONTH
    TWELVE_MONTH

### `region` {#Commitment.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where this commitment may be used.

### `resources` {#Commitment.resources}

| Property | Comments |
|----------|----------|
| Field Name | `resources` |
| Type | [`compute.ResourceCommitment`](gcp_compute.md#ResourceCommitment) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of commitment amounts for particular resources. Note that VCPU and
MEMORY resource commitments must occur together.

### `selfLink` {#Commitment.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `startTimestamp` {#Commitment.startTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `startTimestamp` |
| Type | `string` |

[Output Only] Commitment start time in RFC3339 text format.

### `status` {#Commitment.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] Status of the commitment with regards to eventual expiration
(each commitment has an end date defined). One of the following values:
NOT_YET_ACTIVE, ACTIVE, EXPIRED.
Valid values:
    ACTIVE
    CREATING
    EXPIRED
    NOT_YET_ACTIVE

### `statusMessage` {#Commitment.statusMessage}

| Property | Comments |
|----------|----------|
| Field Name | `statusMessage` |
| Type | `string` |

[Output Only] An optional, human-readable explanation of the status.

## Message `CommitmentAggregatedList` {#CommitmentAggregatedList}



### Inputs for `CommitmentAggregatedList`

* `string` [`id`](#CommitmentAggregatedList.id) = 1
* `repeated` [`compute.CommitmentAggregatedList.ItemsEntry`](gcp_compute.md#CommitmentAggregatedList.ItemsEntry) [`items`](#CommitmentAggregatedList.items) = 2
* `string` [`kind`](#CommitmentAggregatedList.kind) = 3
* `string` [`nextPageToken`](#CommitmentAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#CommitmentAggregatedList.selfLink) = 5
* [`compute.CommitmentAggregatedList.Warning`](gcp_compute.md#CommitmentAggregatedList.Warning) [`warning`](#CommitmentAggregatedList.warning) = 6

### `id` {#CommitmentAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#CommitmentAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.CommitmentAggregatedList.ItemsEntry`](gcp_compute.md#CommitmentAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of CommitmentsScopedList resources.

### `kind` {#CommitmentAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#commitmentAggregatedList for
aggregated lists of commitments.

### `nextPageToken` {#CommitmentAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#CommitmentAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#CommitmentAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.CommitmentAggregatedList.Warning`](gcp_compute.md#CommitmentAggregatedList.Warning) |

## Message `CommitmentList` {#CommitmentList}

Contains a list of Commitment resources.


### Inputs for `CommitmentList`

* `string` [`id`](#CommitmentList.id) = 1
* `repeated` [`compute.Commitment`](gcp_compute.md#Commitment) [`items`](#CommitmentList.items) = 2
* `string` [`kind`](#CommitmentList.kind) = 3
* `string` [`nextPageToken`](#CommitmentList.nextPageToken) = 4
* `string` [`selfLink`](#CommitmentList.selfLink) = 5
* [`compute.CommitmentList.Warning`](gcp_compute.md#CommitmentList.Warning) [`warning`](#CommitmentList.warning) = 6

### `id` {#CommitmentList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#CommitmentList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Commitment`](gcp_compute.md#Commitment) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Commitment resources.

### `kind` {#CommitmentList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#commitmentList for lists of
commitments.

### `nextPageToken` {#CommitmentList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#CommitmentList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#CommitmentList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.CommitmentList.Warning`](gcp_compute.md#CommitmentList.Warning) |

## Message `CommitmentsScopedList` {#CommitmentsScopedList}



### Inputs for `CommitmentsScopedList`

* `repeated` [`compute.Commitment`](gcp_compute.md#Commitment) [`commitments`](#CommitmentsScopedList.commitments) = 1
* [`compute.CommitmentsScopedList.Warning`](gcp_compute.md#CommitmentsScopedList.Warning) [`warning`](#CommitmentsScopedList.warning) = 2

### `commitments` {#CommitmentsScopedList.commitments}

| Property | Comments |
|----------|----------|
| Field Name | `commitments` |
| Type | [`compute.Commitment`](gcp_compute.md#Commitment) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of commitments contained in this scope.

### `warning` {#CommitmentsScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.CommitmentsScopedList.Warning`](gcp_compute.md#CommitmentsScopedList.Warning) |

## Message `ConnectionDraining` {#ConnectionDraining}

Message containing connection draining configuration.


### Inputs for `ConnectionDraining`

* `int32` [`drainingTimeoutSec`](#ConnectionDraining.drainingTimeoutSec) = 1

### `drainingTimeoutSec` {#ConnectionDraining.drainingTimeoutSec}

| Property | Comments |
|----------|----------|
| Field Name | `drainingTimeoutSec` |
| Type | `int32` |

Time for which instance will be drained (not accept new connections, but
still work to finish started).

## Message `CustomerEncryptionKey` {#CustomerEncryptionKey}

Represents a customer-supplied encryption key


### Inputs for `CustomerEncryptionKey`

* `string` [`rawKey`](#CustomerEncryptionKey.rawKey) = 1
* `string` [`sha256`](#CustomerEncryptionKey.sha256) = 2

### `rawKey` {#CustomerEncryptionKey.rawKey}

| Property | Comments |
|----------|----------|
| Field Name | `rawKey` |
| Type | `string` |

Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648
base64 to either encrypt or decrypt this resource.

### `sha256` {#CustomerEncryptionKey.sha256}

| Property | Comments |
|----------|----------|
| Field Name | `sha256` |
| Type | `string` |

[Output only] The RFC 4648 base64 encoded SHA-256 hash of the
customer-supplied encryption key that protects this resource.

## Message `CustomerEncryptionKeyProtectedDisk` {#CustomerEncryptionKeyProtectedDisk}



### Inputs for `CustomerEncryptionKeyProtectedDisk`

* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`diskEncryptionKey`](#CustomerEncryptionKeyProtectedDisk.diskEncryptionKey) = 1
* `string` [`source`](#CustomerEncryptionKeyProtectedDisk.source) = 2

### `diskEncryptionKey` {#CustomerEncryptionKeyProtectedDisk.diskEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `diskEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

Decrypts data associated with the disk with a customer-supplied encryption
key.

### `source` {#CustomerEncryptionKeyProtectedDisk.source}

| Property | Comments |
|----------|----------|
| Field Name | `source` |
| Type | `string` |

Specifies a valid partial or full URL to an existing Persistent Disk
resource. This field is only applicable for persistent disks.

## Message `Data` {#InterconnectAttachmentsScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InterconnectAttachmentsScopedList.Warning.Data.key) = 1
* `string` [`value`](#InterconnectAttachmentsScopedList.Warning.Data.value) = 2

### `key` {#InterconnectAttachmentsScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InterconnectAttachmentsScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AutoscalersScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AutoscalersScopedList.Warning.Data.key) = 1
* `string` [`value`](#AutoscalersScopedList.Warning.Data.value) = 2

### `key` {#AutoscalersScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AutoscalersScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetVpnGatewayList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetVpnGatewayList.Warning.Data.key) = 1
* `string` [`value`](#TargetVpnGatewayList.Warning.Data.value) = 2

### `key` {#TargetVpnGatewayList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetVpnGatewayList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AutoscalerList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AutoscalerList.Warning.Data.key) = 1
* `string` [`value`](#AutoscalerList.Warning.Data.value) = 2

### `key` {#AutoscalerList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AutoscalerList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupsScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupsScopedList.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupsScopedList.Warning.Data.value) = 2

### `key` {#InstanceGroupsScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupsScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetVpnGatewayAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetVpnGatewayAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#TargetVpnGatewayAggregatedList.Warning.Data.value) = 2

### `key` {#TargetVpnGatewayAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetVpnGatewayAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#BackendBucketList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#BackendBucketList.Warning.Data.key) = 1
* `string` [`value`](#BackendBucketList.Warning.Data.value) = 2

### `key` {#BackendBucketList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#BackendBucketList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetVpnGatewaysScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetVpnGatewaysScopedList.Warning.Data.key) = 1
* `string` [`value`](#TargetVpnGatewaysScopedList.Warning.Data.value) = 2

### `key` {#TargetVpnGatewaysScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetVpnGatewaysScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AutoscalerAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AutoscalerAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#AutoscalerAggregatedList.Warning.Data.value) = 2

### `key` {#AutoscalerAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AutoscalerAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceList.Warning.Data.key) = 1
* `string` [`value`](#InstanceList.Warning.Data.value) = 2

### `key` {#InstanceList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#NetworkList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#NetworkList.Warning.Data.key) = 1
* `string` [`value`](#NetworkList.Warning.Data.value) = 2

### `key` {#NetworkList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#NetworkList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#BackendServiceAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#BackendServiceAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#BackendServiceAggregatedList.Warning.Data.value) = 2

### `key` {#BackendServiceAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#BackendServiceAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#UrlMapList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#UrlMapList.Warning.Data.key) = 1
* `string` [`value`](#UrlMapList.Warning.Data.value) = 2

### `key` {#UrlMapList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#UrlMapList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#Route.Warnings.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#Route.Warnings.Data.key) = 1
* `string` [`value`](#Route.Warnings.Data.value) = 2

### `key` {#Route.Warnings.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#Route.Warnings.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#Operation.Warnings.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#Operation.Warnings.Data.key) = 1
* `string` [`value`](#Operation.Warnings.Data.value) = 2

### `key` {#Operation.Warnings.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#Operation.Warnings.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#OperationAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#OperationAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#OperationAggregatedList.Warning.Data.value) = 2

### `key` {#OperationAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#OperationAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetTcpProxyList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetTcpProxyList.Warning.Data.key) = 1
* `string` [`value`](#TargetTcpProxyList.Warning.Data.value) = 2

### `key` {#TargetTcpProxyList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetTcpProxyList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#BackendServiceList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#BackendServiceList.Warning.Data.key) = 1
* `string` [`value`](#BackendServiceList.Warning.Data.value) = 2

### `key` {#BackendServiceList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#BackendServiceList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AddressList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AddressList.Warning.Data.key) = 1
* `string` [`value`](#AddressList.Warning.Data.value) = 2

### `key` {#AddressList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AddressList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#VpnTunnelAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#VpnTunnelAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#VpnTunnelAggregatedList.Warning.Data.value) = 2

### `key` {#VpnTunnelAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#VpnTunnelAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#BackendServicesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#BackendServicesScopedList.Warning.Data.key) = 1
* `string` [`value`](#BackendServicesScopedList.Warning.Data.value) = 2

### `key` {#BackendServicesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#BackendServicesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AddressAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AddressAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#AddressAggregatedList.Warning.Data.value) = 2

### `key` {#AddressAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AddressAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#VpnTunnelList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#VpnTunnelList.Warning.Data.key) = 1
* `string` [`value`](#VpnTunnelList.Warning.Data.value) = 2

### `key` {#VpnTunnelList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#VpnTunnelList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#OperationList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#OperationList.Warning.Data.key) = 1
* `string` [`value`](#OperationList.Warning.Data.value) = 2

### `key` {#OperationList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#OperationList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AcceleratorTypesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AcceleratorTypesScopedList.Warning.Data.key) = 1
* `string` [`value`](#AcceleratorTypesScopedList.Warning.Data.value) = 2

### `key` {#AcceleratorTypesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AcceleratorTypesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceListReferrers.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceListReferrers.Warning.Data.key) = 1
* `string` [`value`](#InstanceListReferrers.Warning.Data.value) = 2

### `key` {#InstanceListReferrers.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceListReferrers.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetSslProxyList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetSslProxyList.Warning.Data.key) = 1
* `string` [`value`](#TargetSslProxyList.Warning.Data.value) = 2

### `key` {#TargetSslProxyList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetSslProxyList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#CommitmentAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#CommitmentAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#CommitmentAggregatedList.Warning.Data.value) = 2

### `key` {#CommitmentAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#CommitmentAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#VpnTunnelsScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#VpnTunnelsScopedList.Warning.Data.key) = 1
* `string` [`value`](#VpnTunnelsScopedList.Warning.Data.value) = 2

### `key` {#VpnTunnelsScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#VpnTunnelsScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetPoolsScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetPoolsScopedList.Warning.Data.key) = 1
* `string` [`value`](#TargetPoolsScopedList.Warning.Data.value) = 2

### `key` {#TargetPoolsScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetPoolsScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupsListInstances.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupsListInstances.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupsListInstances.Warning.Data.value) = 2

### `key` {#InstanceGroupsListInstances.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupsListInstances.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AcceleratorTypeList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AcceleratorTypeList.Warning.Data.key) = 1
* `string` [`value`](#AcceleratorTypeList.Warning.Data.value) = 2

### `key` {#AcceleratorTypeList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AcceleratorTypeList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#MachineTypesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#MachineTypesScopedList.Warning.Data.key) = 1
* `string` [`value`](#MachineTypesScopedList.Warning.Data.value) = 2

### `key` {#MachineTypesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#MachineTypesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#CommitmentsScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#CommitmentsScopedList.Warning.Data.key) = 1
* `string` [`value`](#CommitmentsScopedList.Warning.Data.value) = 2

### `key` {#CommitmentsScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#CommitmentsScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#XpnHostList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#XpnHostList.Warning.Data.key) = 1
* `string` [`value`](#XpnHostList.Warning.Data.value) = 2

### `key` {#XpnHostList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#XpnHostList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AcceleratorTypeAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AcceleratorTypeAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#AcceleratorTypeAggregatedList.Warning.Data.value) = 2

### `key` {#AcceleratorTypeAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AcceleratorTypeAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#ZoneList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#ZoneList.Warning.Data.key) = 1
* `string` [`value`](#ZoneList.Warning.Data.value) = 2

### `key` {#ZoneList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#ZoneList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RegionList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RegionList.Warning.Data.key) = 1
* `string` [`value`](#RegionList.Warning.Data.value) = 2

### `key` {#RegionList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RegionList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceTemplateList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceTemplateList.Warning.Data.key) = 1
* `string` [`value`](#InstanceTemplateList.Warning.Data.value) = 2

### `key` {#InstanceTemplateList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceTemplateList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#MachineTypeList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#MachineTypeList.Warning.Data.key) = 1
* `string` [`value`](#MachineTypeList.Warning.Data.value) = 2

### `key` {#MachineTypeList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#MachineTypeList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RegionInstanceGroupsListInstances.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RegionInstanceGroupsListInstances.Warning.Data.key) = 1
* `string` [`value`](#RegionInstanceGroupsListInstances.Warning.Data.value) = 2

### `key` {#RegionInstanceGroupsListInstances.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RegionInstanceGroupsListInstances.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#MachineTypeAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#MachineTypeAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#MachineTypeAggregatedList.Warning.Data.value) = 2

### `key` {#MachineTypeAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#MachineTypeAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetPoolList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetPoolList.Warning.Data.key) = 1
* `string` [`value`](#TargetPoolList.Warning.Data.value) = 2

### `key` {#TargetPoolList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetPoolList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#DiskAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#DiskAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#DiskAggregatedList.Warning.Data.value) = 2

### `key` {#DiskAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#DiskAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupManagersScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupManagersScopedList.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupManagersScopedList.Warning.Data.value) = 2

### `key` {#InstanceGroupManagersScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupManagersScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RouterAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RouterAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#RouterAggregatedList.Warning.Data.value) = 2

### `key` {#RouterAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RouterAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetPoolAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetPoolAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#TargetPoolAggregatedList.Warning.Data.value) = 2

### `key` {#TargetPoolAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetPoolAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#DiskList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#DiskList.Warning.Data.key) = 1
* `string` [`value`](#DiskList.Warning.Data.value) = 2

### `key` {#DiskList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#DiskList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupManagerList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupManagerList.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupManagerList.Warning.Data.value) = 2

### `key` {#InstanceGroupManagerList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupManagerList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RouterList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RouterList.Warning.Data.key) = 1
* `string` [`value`](#RouterList.Warning.Data.value) = 2

### `key` {#RouterList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RouterList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupManagerAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupManagerAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupManagerAggregatedList.Warning.Data.value) = 2

### `key` {#InstanceGroupManagerAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupManagerAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#LicensesListResponse.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#LicensesListResponse.Warning.Data.key) = 1
* `string` [`value`](#LicensesListResponse.Warning.Data.value) = 2

### `key` {#LicensesListResponse.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#LicensesListResponse.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#ForwardingRulesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#ForwardingRulesScopedList.Warning.Data.key) = 1
* `string` [`value`](#ForwardingRulesScopedList.Warning.Data.value) = 2

### `key` {#ForwardingRulesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#ForwardingRulesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#DiskTypeAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#DiskTypeAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#DiskTypeAggregatedList.Warning.Data.value) = 2

### `key` {#DiskTypeAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#DiskTypeAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RoutersScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RoutersScopedList.Warning.Data.key) = 1
* `string` [`value`](#RoutersScopedList.Warning.Data.value) = 2

### `key` {#RoutersScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RoutersScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetInstancesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetInstancesScopedList.Warning.Data.key) = 1
* `string` [`value`](#TargetInstancesScopedList.Warning.Data.value) = 2

### `key` {#TargetInstancesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetInstancesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#DiskTypeList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#DiskTypeList.Warning.Data.key) = 1
* `string` [`value`](#DiskTypeList.Warning.Data.value) = 2

### `key` {#DiskTypeList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#DiskTypeList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstancesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstancesScopedList.Warning.Data.key) = 1
* `string` [`value`](#InstancesScopedList.Warning.Data.value) = 2

### `key` {#InstancesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstancesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetInstanceList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetInstanceList.Warning.Data.key) = 1
* `string` [`value`](#TargetInstanceList.Warning.Data.value) = 2

### `key` {#TargetInstanceList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetInstanceList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#DiskTypesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#DiskTypesScopedList.Warning.Data.key) = 1
* `string` [`value`](#DiskTypesScopedList.Warning.Data.value) = 2

### `key` {#DiskTypesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#DiskTypesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RouteList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RouteList.Warning.Data.key) = 1
* `string` [`value`](#RouteList.Warning.Data.value) = 2

### `key` {#RouteList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RouteList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupList.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupList.Warning.Data.value) = 2

### `key` {#InstanceGroupList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetInstanceAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetInstanceAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#TargetInstanceAggregatedList.Warning.Data.value) = 2

### `key` {#TargetInstanceAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetInstanceAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#DisksScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#DisksScopedList.Warning.Data.key) = 1
* `string` [`value`](#DisksScopedList.Warning.Data.value) = 2

### `key` {#DisksScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#DisksScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RegionInstanceGroupManagerList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RegionInstanceGroupManagerList.Warning.Data.key) = 1
* `string` [`value`](#RegionInstanceGroupManagerList.Warning.Data.value) = 2

### `key` {#RegionInstanceGroupManagerList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RegionInstanceGroupManagerList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#OperationsScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#OperationsScopedList.Warning.Data.key) = 1
* `string` [`value`](#OperationsScopedList.Warning.Data.value) = 2

### `key` {#OperationsScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#OperationsScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceGroupAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceGroupAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#InstanceGroupAggregatedList.Warning.Data.value) = 2

### `key` {#InstanceGroupAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceGroupAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SnapshotList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SnapshotList.Warning.Data.key) = 1
* `string` [`value`](#SnapshotList.Warning.Data.value) = 2

### `key` {#SnapshotList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SnapshotList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InterconnectLocationList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InterconnectLocationList.Warning.Data.key) = 1
* `string` [`value`](#InterconnectLocationList.Warning.Data.value) = 2

### `key` {#InterconnectLocationList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InterconnectLocationList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#FirewallList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#FirewallList.Warning.Data.key) = 1
* `string` [`value`](#FirewallList.Warning.Data.value) = 2

### `key` {#FirewallList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#FirewallList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InterconnectAttachmentAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InterconnectAttachmentAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#InterconnectAttachmentAggregatedList.Warning.Data.value) = 2

### `key` {#InterconnectAttachmentAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InterconnectAttachmentAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InstanceAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InstanceAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#InstanceAggregatedList.Warning.Data.value) = 2

### `key` {#InstanceAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InstanceAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RegionAutoscalerList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RegionAutoscalerList.Warning.Data.key) = 1
* `string` [`value`](#RegionAutoscalerList.Warning.Data.value) = 2

### `key` {#RegionAutoscalerList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RegionAutoscalerList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetHttpsProxyList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetHttpsProxyList.Warning.Data.key) = 1
* `string` [`value`](#TargetHttpsProxyList.Warning.Data.value) = 2

### `key` {#TargetHttpsProxyList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetHttpsProxyList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#ForwardingRuleAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#ForwardingRuleAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#ForwardingRuleAggregatedList.Warning.Data.value) = 2

### `key` {#ForwardingRuleAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#ForwardingRuleAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SslCertificateList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SslCertificateList.Warning.Data.key) = 1
* `string` [`value`](#SslCertificateList.Warning.Data.value) = 2

### `key` {#SslCertificateList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SslCertificateList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#TargetHttpProxyList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#TargetHttpProxyList.Warning.Data.key) = 1
* `string` [`value`](#TargetHttpProxyList.Warning.Data.value) = 2

### `key` {#TargetHttpProxyList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#TargetHttpProxyList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#ForwardingRuleList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#ForwardingRuleList.Warning.Data.key) = 1
* `string` [`value`](#ForwardingRuleList.Warning.Data.value) = 2

### `key` {#ForwardingRuleList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#ForwardingRuleList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RegionInstanceGroupList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RegionInstanceGroupList.Warning.Data.key) = 1
* `string` [`value`](#RegionInstanceGroupList.Warning.Data.value) = 2

### `key` {#RegionInstanceGroupList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RegionInstanceGroupList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SubnetworksScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SubnetworksScopedList.Warning.Data.key) = 1
* `string` [`value`](#SubnetworksScopedList.Warning.Data.value) = 2

### `key` {#SubnetworksScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SubnetworksScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SubnetworkList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SubnetworkList.Warning.Data.key) = 1
* `string` [`value`](#SubnetworkList.Warning.Data.value) = 2

### `key` {#SubnetworkList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SubnetworkList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InterconnectAttachmentList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InterconnectAttachmentList.Warning.Data.key) = 1
* `string` [`value`](#InterconnectAttachmentList.Warning.Data.value) = 2

### `key` {#InterconnectAttachmentList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InterconnectAttachmentList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#InterconnectList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#InterconnectList.Warning.Data.key) = 1
* `string` [`value`](#InterconnectList.Warning.Data.value) = 2

### `key` {#InterconnectList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#InterconnectList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#ImageList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#ImageList.Warning.Data.key) = 1
* `string` [`value`](#ImageList.Warning.Data.value) = 2

### `key` {#ImageList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#ImageList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SslPoliciesList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SslPoliciesList.Warning.Data.key) = 1
* `string` [`value`](#SslPoliciesList.Warning.Data.value) = 2

### `key` {#SslPoliciesList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SslPoliciesList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#CommitmentList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#CommitmentList.Warning.Data.key) = 1
* `string` [`value`](#CommitmentList.Warning.Data.value) = 2

### `key` {#CommitmentList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#CommitmentList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#RegionDiskTypeList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#RegionDiskTypeList.Warning.Data.key) = 1
* `string` [`value`](#RegionDiskTypeList.Warning.Data.value) = 2

### `key` {#RegionDiskTypeList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#RegionDiskTypeList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SslPolicy.Warnings.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SslPolicy.Warnings.Data.key) = 1
* `string` [`value`](#SslPolicy.Warnings.Data.value) = 2

### `key` {#SslPolicy.Warnings.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SslPolicy.Warnings.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#SubnetworkAggregatedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#SubnetworkAggregatedList.Warning.Data.key) = 1
* `string` [`value`](#SubnetworkAggregatedList.Warning.Data.value) = 2

### `key` {#SubnetworkAggregatedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#SubnetworkAggregatedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#HealthCheckList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#HealthCheckList.Warning.Data.key) = 1
* `string` [`value`](#HealthCheckList.Warning.Data.value) = 2

### `key` {#HealthCheckList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#HealthCheckList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#HttpHealthCheckList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#HttpHealthCheckList.Warning.Data.key) = 1
* `string` [`value`](#HttpHealthCheckList.Warning.Data.value) = 2

### `key` {#HttpHealthCheckList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#HttpHealthCheckList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#AddressesScopedList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#AddressesScopedList.Warning.Data.key) = 1
* `string` [`value`](#AddressesScopedList.Warning.Data.value) = 2

### `key` {#AddressesScopedList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#AddressesScopedList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Data` {#HttpsHealthCheckList.Warning.Data}

[Output Only] Metadata about this warning in key: value format. For example:
"data": [ { "key": "scope", "value": "zones/us-east1-d" }


### Inputs for `Data`

* `string` [`key`](#HttpsHealthCheckList.Warning.Data.key) = 1
* `string` [`value`](#HttpsHealthCheckList.Warning.Data.value) = 2

### `key` {#HttpsHealthCheckList.Warning.Data.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

[Output Only] A key that provides more detail on the warning being returned.
For example, for warnings where there are no results in a list request for a
particular zone, this key might be scope and the key value might be the zone
name. Other examples might be a key indicating a deprecated resource and a
suggested replacement, or a warning about invalid network settings (for
example, if an instance attempts to perform IP forwarding but is not enabled
for IP forwarding).

### `value` {#HttpsHealthCheckList.Warning.Data.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

[Output Only] A warning data value corresponding to the key.

## Message `Denied` {#Firewall.Denied}

The list of DENY rules specified by this firewall. Each rule specifies a
protocol and port-range tuple that describes a denied connection.


### Inputs for `Denied`

* `string` [`IPProtocol`](#Firewall.Denied.IPProtocol) = 1
* `repeated` `string` [`ports`](#Firewall.Denied.ports) = 2

### `IPProtocol` {#Firewall.Denied.IPProtocol}

| Property | Comments |
|----------|----------|
| Field Name | `IPProtocol` |
| Type | `string` |

The IP protocol to which this rule applies. The protocol type is required
when creating a firewall rule. This value can either be one of the following
well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP
protocol number.

### `ports` {#Firewall.Denied.ports}

| Property | Comments |
|----------|----------|
| Field Name | `ports` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

An optional list of ports to which this rule applies. This field is only
applicable for UDP or TCP protocol. Each entry must be either an integer or
a range. If not specified, this rule applies to connections through any
port.

Example inputs include: ["22"], ["80","443"], and ["12345-12349"].

## Message `DeprecationStatus` {#DeprecationStatus}

Deprecation status for a public resource.


### Inputs for `DeprecationStatus`

* `string` [`deleted`](#DeprecationStatus.deleted) = 1
* `string` [`deprecated`](#DeprecationStatus.deprecated) = 2
* `string` [`obsolete`](#DeprecationStatus.obsolete) = 3
* `string` [`replacement`](#DeprecationStatus.replacement) = 4
* `string` [`state`](#DeprecationStatus.state) = 5

### `deleted` {#DeprecationStatus.deleted}

| Property | Comments |
|----------|----------|
| Field Name | `deleted` |
| Type | `string` |

An optional RFC3339 timestamp on or after which the state of this resource
is intended to change to DELETED. This is only informational and the status
will not change unless the client explicitly changes it.

### `deprecated` {#DeprecationStatus.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | `string` |

An optional RFC3339 timestamp on or after which the state of this resource
is intended to change to DEPRECATED. This is only informational and the
status will not change unless the client explicitly changes it.

### `obsolete` {#DeprecationStatus.obsolete}

| Property | Comments |
|----------|----------|
| Field Name | `obsolete` |
| Type | `string` |

An optional RFC3339 timestamp on or after which the state of this resource
is intended to change to OBSOLETE. This is only informational and the status
will not change unless the client explicitly changes it.

### `replacement` {#DeprecationStatus.replacement}

| Property | Comments |
|----------|----------|
| Field Name | `replacement` |
| Type | `string` |

The URL of the suggested replacement for a deprecated resource. The
suggested replacement resource must be the same kind of resource as the
deprecated resource.

### `state` {#DeprecationStatus.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

The deprecation state of this resource. This can be DEPRECATED, OBSOLETE, or
DELETED. Operations which create a new resource using a DEPRECATED resource
will return successfully, but with a warning indicating the deprecated
resource and recommending its replacement. Operations which use OBSOLETE or
DELETED resources will be rejected and result in an error.
Valid values:
    DELETED
    DEPRECATED
    OBSOLETE

## Message `Disk` {#Disk}

A Disk resource. (== resource_for beta.disks ==) (== resource_for v1.disks
==)


### Inputs for `Disk`

* `string` [`creationTimestamp`](#Disk.creationTimestamp) = 1
* `string` [`description`](#Disk.description) = 2
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`diskEncryptionKey`](#Disk.diskEncryptionKey) = 3
* `repeated` [`compute.GuestOsFeature`](gcp_compute.md#GuestOsFeature) [`guestOsFeatures`](#Disk.guestOsFeatures) = 4
* `string` [`id`](#Disk.id) = 5
* `string` [`kind`](#Disk.kind) = 6
* `string` [`labelFingerprint`](#Disk.labelFingerprint) = 7
* `repeated` [`compute.Disk.LabelsEntry`](gcp_compute.md#Disk.LabelsEntry) [`labels`](#Disk.labels) = 8
* `string` [`lastAttachTimestamp`](#Disk.lastAttachTimestamp) = 9
* `string` [`lastDetachTimestamp`](#Disk.lastDetachTimestamp) = 10
* `repeated` `string` [`licenseCodes`](#Disk.licenseCodes) = 11
* `repeated` `string` [`licenses`](#Disk.licenses) = 12
* `string` [`name`](#Disk.name) = 13 (**Required**)
* `string` [`options`](#Disk.options) = 14
* `string` [`region`](#Disk.region) = 15
* `repeated` `string` [`replicaZones`](#Disk.replicaZones) = 16
* `string` [`selfLink`](#Disk.selfLink) = 17
* `string` [`sizeGb`](#Disk.sizeGb) = 18
* `string` [`sourceImage`](#Disk.sourceImage) = 19
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceImageEncryptionKey`](#Disk.sourceImageEncryptionKey) = 20
* `string` [`sourceImageId`](#Disk.sourceImageId) = 21
* `string` [`sourceSnapshot`](#Disk.sourceSnapshot) = 22
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceSnapshotEncryptionKey`](#Disk.sourceSnapshotEncryptionKey) = 23
* `string` [`sourceSnapshotId`](#Disk.sourceSnapshotId) = 24
* `string` [`status`](#Disk.status) = 25
* `string` [`type`](#Disk.type) = 26
* `repeated` `string` [`users`](#Disk.users) = 27
* `string` [`zone`](#Disk.zone) = 28

### `creationTimestamp` {#Disk.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Disk.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `diskEncryptionKey` {#Disk.diskEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `diskEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

Encrypts the disk using a customer-supplied encryption key.

After you encrypt a disk with a customer-supplied key, you must provide the
same key if you use the disk later (e.g. to create a disk snapshot or an
image, or to attach the disk to a virtual machine).

Customer-supplied encryption keys do not protect access to metadata of the
disk.

If you do not provide an encryption key when creating the disk, then the
disk will be encrypted using an automatically generated key and you do not
need to provide a key to use the disk later.

### `guestOsFeatures` {#Disk.guestOsFeatures}

| Property | Comments |
|----------|----------|
| Field Name | `guestOsFeatures` |
| Type | [`compute.GuestOsFeature`](gcp_compute.md#GuestOsFeature) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of features to enable on the guest operating system. Applicable only
for bootable images. Read  Enabling guest operating system features to see a
list of available options.

### `id` {#Disk.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Disk.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#disk for disks.

### `labelFingerprint` {#Disk.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

A fingerprint for the labels being applied to this disk, which is
essentially a hash of the labels set used for optimistic locking. The
fingerprint is initially generated by Compute Engine and changes after every
request to modify or update labels. You must always provide an up-to-date
fingerprint hash in order to update or change labels.

To see the latest fingerprint, make a get() request to retrieve a disk.

### `labels` {#Disk.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.Disk.LabelsEntry`](gcp_compute.md#Disk.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Labels to apply to this disk. These can be later modified by the setLabels
method.

### `lastAttachTimestamp` {#Disk.lastAttachTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `lastAttachTimestamp` |
| Type | `string` |

[Output Only] Last attach timestamp in RFC3339 text format.

### `lastDetachTimestamp` {#Disk.lastDetachTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `lastDetachTimestamp` |
| Type | `string` |

[Output Only] Last detach timestamp in RFC3339 text format.

### `licenseCodes` {#Disk.licenseCodes}

| Property | Comments |
|----------|----------|
| Field Name | `licenseCodes` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Integer license codes indicating which licenses are attached to this disk.

### `licenses` {#Disk.licenses}

| Property | Comments |
|----------|----------|
| Field Name | `licenses` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Any applicable publicly visible licenses.

### `name` {#Disk.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `options` {#Disk.options}

| Property | Comments |
|----------|----------|
| Field Name | `options` |
| Type | `string` |

Internal use only.

### `region` {#Disk.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the disk resides. Only applicable for
regional resources. You must specify this field as part of the HTTP request
URL. It is not settable as a field in the request body.

### `replicaZones` {#Disk.replicaZones}

| Property | Comments |
|----------|----------|
| Field Name | `replicaZones` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

URLs of the zones where the disk should be replicated to. Only applicable
for regional resources.

### `selfLink` {#Disk.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined fully-qualified URL for this resource.

### `sizeGb` {#Disk.sizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `sizeGb` |
| Type | `string` |

Size of the persistent disk, specified in GB. You can specify this field
when creating a persistent disk using the sourceImage or sourceSnapshot
parameter, or specify it alone to create an empty persistent disk.

If you specify this field along with sourceImage or sourceSnapshot, the
value of sizeGb must not be less than the size of the sourceImage or the
size of the snapshot. Acceptable values are 1 to 65536, inclusive.

### `sourceImage` {#Disk.sourceImage}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImage` |
| Type | `string` |

The source image used to create this disk. If the source image is deleted,
this field will not be set.

To create a disk with one of the public operating system images, specify the
image by its family name. For example, specify family/debian-8 to use the
latest Debian 8 image:
projects/debian-cloud/global/images/family/debian-8


Alternatively, use a specific version of a public operating system image:
projects/debian-cloud/global/images/debian-8-jessie-vYYYYMMDD


To create a disk with a custom image that you created, specify the image
name in the following format:
global/images/my-custom-image


You can also specify a custom image by its image family, which returns the
latest version of the image in that family. Replace the image name with
family/family-name:
global/images/family/my-image-family

### `sourceImageEncryptionKey` {#Disk.sourceImageEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImageEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source image. Required if the
source image is protected by a customer-supplied encryption key.

### `sourceImageId` {#Disk.sourceImageId}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImageId` |
| Type | `string` |

[Output Only] The ID value of the image used to create this disk. This value
identifies the exact image that was used to create this persistent disk. For
example, if you created the persistent disk from an image that was later
deleted and recreated under the same name, the source image ID would
identify the exact version of the image that was used.

### `sourceSnapshot` {#Disk.sourceSnapshot}

| Property | Comments |
|----------|----------|
| Field Name | `sourceSnapshot` |
| Type | `string` |

The source snapshot used to create this disk. You can provide this as a
partial or full URL to the resource. For example, the following are valid
values:
-
https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot
- projects/project/global/snapshots/snapshot
- global/snapshots/snapshot

### `sourceSnapshotEncryptionKey` {#Disk.sourceSnapshotEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceSnapshotEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source snapshot. Required if the
source snapshot is protected by a customer-supplied encryption key.

### `sourceSnapshotId` {#Disk.sourceSnapshotId}

| Property | Comments |
|----------|----------|
| Field Name | `sourceSnapshotId` |
| Type | `string` |

[Output Only] The unique ID of the snapshot used to create this disk. This
value identifies the exact snapshot that was used to create this persistent
disk. For example, if you created the persistent disk from a snapshot that
was later deleted and recreated under the same name, the source snapshot ID
would identify the exact version of the snapshot that was used.

### `status` {#Disk.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of disk creation.
Valid values:
    CREATING
    FAILED
    READY
    RESTORING

### `type` {#Disk.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

URL of the disk type resource describing which disk type to use to create
the disk. Provide this when creating the disk. For example:
project/zones/zone/diskTypes/pd-standard or pd-ssd

### `users` {#Disk.users}

| Property | Comments |
|----------|----------|
| Field Name | `users` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Links to the users of the disk (attached instances) in form:
project/zones/zone/instances/instance

### `zone` {#Disk.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] URL of the zone where the disk resides. You must specify this
field as part of the HTTP request URL. It is not settable as a field in the
request body.

## Message `DiskAggregatedList` {#DiskAggregatedList}



### Inputs for `DiskAggregatedList`

* `string` [`id`](#DiskAggregatedList.id) = 1
* `repeated` [`compute.DiskAggregatedList.ItemsEntry`](gcp_compute.md#DiskAggregatedList.ItemsEntry) [`items`](#DiskAggregatedList.items) = 2
* `string` [`kind`](#DiskAggregatedList.kind) = 3
* `string` [`nextPageToken`](#DiskAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#DiskAggregatedList.selfLink) = 5
* [`compute.DiskAggregatedList.Warning`](gcp_compute.md#DiskAggregatedList.Warning) [`warning`](#DiskAggregatedList.warning) = 6

### `id` {#DiskAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#DiskAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.DiskAggregatedList.ItemsEntry`](gcp_compute.md#DiskAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of DisksScopedList resources.

### `kind` {#DiskAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#diskAggregatedList for
aggregated lists of persistent disks.

### `nextPageToken` {#DiskAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#DiskAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#DiskAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.DiskAggregatedList.Warning`](gcp_compute.md#DiskAggregatedList.Warning) |

## Message `DiskInstantiationConfig` {#DiskInstantiationConfig}

A specification of the desired way to instantiate a disk in the instance
template when its created from a source instance.


### Inputs for `DiskInstantiationConfig`

* `bool` [`autoDelete`](#DiskInstantiationConfig.autoDelete) = 1
* `string` [`customImage`](#DiskInstantiationConfig.customImage) = 2
* `string` [`deviceName`](#DiskInstantiationConfig.deviceName) = 3
* `string` [`instantiateFrom`](#DiskInstantiationConfig.instantiateFrom) = 4

### `autoDelete` {#DiskInstantiationConfig.autoDelete}

| Property | Comments |
|----------|----------|
| Field Name | `autoDelete` |
| Type | `bool` |

Specifies whether the disk will be auto-deleted when the instance is deleted
(but not when the disk is detached from the instance).

### `customImage` {#DiskInstantiationConfig.customImage}

| Property | Comments |
|----------|----------|
| Field Name | `customImage` |
| Type | `string` |

The custom source image to be used to restore this disk when instantiating
this instance template.

### `deviceName` {#DiskInstantiationConfig.deviceName}

| Property | Comments |
|----------|----------|
| Field Name | `deviceName` |
| Type | `string` |

Specifies the device name of the disk to which the configurations apply to.

### `instantiateFrom` {#DiskInstantiationConfig.instantiateFrom}

| Property | Comments |
|----------|----------|
| Field Name | `instantiateFrom` |
| Type | `string` |

Specifies whether to include the disk and what image to use. Possible values
are:
- source-image: to use the same image that was used to create the source
instance's corresponding disk. Applicable to the boot disk and additional
read-write disks.
- source-image-family: to use the same image family that was used to create
the source instance's corresponding disk. Applicable to the boot disk and
additional read-write disks.
- custom-image: to use a user-provided image url for disk creation.
Applicable to the boot disk and additional read-write disks.
- attach-read-only: to attach a read-only disk. Applicable to read-only
disks.
- do-not-include: to exclude a disk from the template. Applicable to
additional read-write disks, local SSDs, and read-only disks.
Valid values:
    ATTACH_READ_ONLY
    BLANK
    CUSTOM_IMAGE
    DEFAULT
    DO_NOT_INCLUDE
    SOURCE_IMAGE
    SOURCE_IMAGE_FAMILY

## Message `DiskList` {#DiskList}

A list of Disk resources.


### Inputs for `DiskList`

* `string` [`id`](#DiskList.id) = 1
* `repeated` [`compute.Disk`](gcp_compute.md#Disk) [`items`](#DiskList.items) = 2
* `string` [`kind`](#DiskList.kind) = 3
* `string` [`nextPageToken`](#DiskList.nextPageToken) = 4
* `string` [`selfLink`](#DiskList.selfLink) = 5
* [`compute.DiskList.Warning`](gcp_compute.md#DiskList.Warning) [`warning`](#DiskList.warning) = 6

### `id` {#DiskList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#DiskList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Disk`](gcp_compute.md#Disk) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Disk resources.

### `kind` {#DiskList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#diskList for lists of disks.

### `nextPageToken` {#DiskList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#DiskList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#DiskList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.DiskList.Warning`](gcp_compute.md#DiskList.Warning) |

## Message `DiskMoveRequest` {#DiskMoveRequest}



### Inputs for `DiskMoveRequest`

* `string` [`destinationZone`](#DiskMoveRequest.destinationZone) = 1
* `string` [`targetDisk`](#DiskMoveRequest.targetDisk) = 2

### `destinationZone` {#DiskMoveRequest.destinationZone}

| Property | Comments |
|----------|----------|
| Field Name | `destinationZone` |
| Type | `string` |

The URL of the destination zone to move the disk. This can be a full or
partial URL. For example, the following are all valid URLs to a zone:
- https://www.googleapis.com/compute/v1/projects/project/zones/zone
- projects/project/zones/zone
- zones/zone

### `targetDisk` {#DiskMoveRequest.targetDisk}

| Property | Comments |
|----------|----------|
| Field Name | `targetDisk` |
| Type | `string` |

The URL of the target disk to move. This can be a full or partial URL. For
example, the following are all valid URLs to a disk:
-
https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
- projects/project/zones/zone/disks/disk
- zones/zone/disks/disk

## Message `DiskType` {#DiskType}

A DiskType resource. (== resource_for beta.diskTypes ==) (== resource_for
v1.diskTypes ==)


### Inputs for `DiskType`

* `string` [`creationTimestamp`](#DiskType.creationTimestamp) = 1
* `string` [`defaultDiskSizeGb`](#DiskType.defaultDiskSizeGb) = 2
* [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) [`deprecated`](#DiskType.deprecated) = 3
* `string` [`description`](#DiskType.description) = 4
* `string` [`id`](#DiskType.id) = 5
* `string` [`kind`](#DiskType.kind) = 6
* `string` [`name`](#DiskType.name) = 7 (**Required**)
* `string` [`region`](#DiskType.region) = 8
* `string` [`selfLink`](#DiskType.selfLink) = 9
* `string` [`validDiskSize`](#DiskType.validDiskSize) = 10
* `string` [`zone`](#DiskType.zone) = 11

### `creationTimestamp` {#DiskType.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `defaultDiskSizeGb` {#DiskType.defaultDiskSizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `defaultDiskSizeGb` |
| Type | `string` |

[Output Only] Server-defined default disk size in GB.

### `deprecated` {#DiskType.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) |

[Output Only] The deprecation status associated with this disk type.

### `description` {#DiskType.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] An optional description of this resource.

### `id` {#DiskType.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#DiskType.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#diskType for disk types.

### `name` {#DiskType.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `region` {#DiskType.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the disk type resides. Only applicable
for regional resources. You must specify this field as part of the HTTP
request URL. It is not settable as a field in the request body.

### `selfLink` {#DiskType.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `validDiskSize` {#DiskType.validDiskSize}

| Property | Comments |
|----------|----------|
| Field Name | `validDiskSize` |
| Type | `string` |

[Output Only] An optional textual description of the valid disk size, such
as "10GB-10TB".

### `zone` {#DiskType.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] URL of the zone where the disk type resides. You must specify
this field as part of the HTTP request URL. It is not settable as a field in
the request body.

## Message `DiskTypeAggregatedList` {#DiskTypeAggregatedList}



### Inputs for `DiskTypeAggregatedList`

* `string` [`id`](#DiskTypeAggregatedList.id) = 1
* `repeated` [`compute.DiskTypeAggregatedList.ItemsEntry`](gcp_compute.md#DiskTypeAggregatedList.ItemsEntry) [`items`](#DiskTypeAggregatedList.items) = 2
* `string` [`kind`](#DiskTypeAggregatedList.kind) = 3
* `string` [`nextPageToken`](#DiskTypeAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#DiskTypeAggregatedList.selfLink) = 5
* [`compute.DiskTypeAggregatedList.Warning`](gcp_compute.md#DiskTypeAggregatedList.Warning) [`warning`](#DiskTypeAggregatedList.warning) = 6

### `id` {#DiskTypeAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#DiskTypeAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.DiskTypeAggregatedList.ItemsEntry`](gcp_compute.md#DiskTypeAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of DiskTypesScopedList resources.

### `kind` {#DiskTypeAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#diskTypeAggregatedList.

### `nextPageToken` {#DiskTypeAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#DiskTypeAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#DiskTypeAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.DiskTypeAggregatedList.Warning`](gcp_compute.md#DiskTypeAggregatedList.Warning) |

## Message `DiskTypeList` {#DiskTypeList}

Contains a list of disk types.


### Inputs for `DiskTypeList`

* `string` [`id`](#DiskTypeList.id) = 1
* `repeated` [`compute.DiskType`](gcp_compute.md#DiskType) [`items`](#DiskTypeList.items) = 2
* `string` [`kind`](#DiskTypeList.kind) = 3
* `string` [`nextPageToken`](#DiskTypeList.nextPageToken) = 4
* `string` [`selfLink`](#DiskTypeList.selfLink) = 5
* [`compute.DiskTypeList.Warning`](gcp_compute.md#DiskTypeList.Warning) [`warning`](#DiskTypeList.warning) = 6

### `id` {#DiskTypeList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#DiskTypeList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.DiskType`](gcp_compute.md#DiskType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of DiskType resources.

### `kind` {#DiskTypeList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#diskTypeList for disk types.

### `nextPageToken` {#DiskTypeList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#DiskTypeList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#DiskTypeList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.DiskTypeList.Warning`](gcp_compute.md#DiskTypeList.Warning) |

## Message `DiskTypesScopedList` {#DiskTypesScopedList}



### Inputs for `DiskTypesScopedList`

* `repeated` [`compute.DiskType`](gcp_compute.md#DiskType) [`diskTypes`](#DiskTypesScopedList.diskTypes) = 1
* [`compute.DiskTypesScopedList.Warning`](gcp_compute.md#DiskTypesScopedList.Warning) [`warning`](#DiskTypesScopedList.warning) = 2

### `diskTypes` {#DiskTypesScopedList.diskTypes}

| Property | Comments |
|----------|----------|
| Field Name | `diskTypes` |
| Type | [`compute.DiskType`](gcp_compute.md#DiskType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of disk types contained in this scope.

### `warning` {#DiskTypesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.DiskTypesScopedList.Warning`](gcp_compute.md#DiskTypesScopedList.Warning) |

## Message `DisksResizeRequest` {#DisksResizeRequest}



### Inputs for `DisksResizeRequest`

* `string` [`sizeGb`](#DisksResizeRequest.sizeGb) = 1

### `sizeGb` {#DisksResizeRequest.sizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `sizeGb` |
| Type | `string` |

The new size of the persistent disk, which is specified in GB.

## Message `DisksScopedList` {#DisksScopedList}



### Inputs for `DisksScopedList`

* `repeated` [`compute.Disk`](gcp_compute.md#Disk) [`disks`](#DisksScopedList.disks) = 1
* [`compute.DisksScopedList.Warning`](gcp_compute.md#DisksScopedList.Warning) [`warning`](#DisksScopedList.warning) = 2

### `disks` {#DisksScopedList.disks}

| Property | Comments |
|----------|----------|
| Field Name | `disks` |
| Type | [`compute.Disk`](gcp_compute.md#Disk) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of disks contained in this scope.

### `warning` {#DisksScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.DisksScopedList.Warning`](gcp_compute.md#DisksScopedList.Warning) |

## Message `Error` {#Operation.Error}

[Output Only] If errors are generated during processing of the operation,
this field will be populated.
[Output Only] If errors are generated during processing of the operation,
this field will be populated.


### Inputs for `Error`

* `repeated` [`compute.Operation.Error.Errors`](gcp_compute.md#Operation.Error.Errors) [`errors`](#Operation.Error.errors) = 1

### `errors` {#Operation.Error.errors}

| Property | Comments |
|----------|----------|
| Field Name | `errors` |
| Type | [`compute.Operation.Error.Errors`](gcp_compute.md#Operation.Error.Errors) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `Errors` {#ManagedInstanceLastAttempt.Errors}

[Output Only] Encountered errors during the last attempt to create or delete
the instance.
[Output Only] Encountered errors during the last attempt to create or delete
the instance.


### Inputs for `Errors`

* `repeated` [`compute.ManagedInstanceLastAttempt.Errors.Errors`](gcp_compute.md#ManagedInstanceLastAttempt.Errors.Errors) [`errors`](#ManagedInstanceLastAttempt.Errors.errors) = 1

### `errors` {#ManagedInstanceLastAttempt.Errors.errors}

| Property | Comments |
|----------|----------|
| Field Name | `errors` |
| Type | [`compute.ManagedInstanceLastAttempt.Errors.Errors`](gcp_compute.md#ManagedInstanceLastAttempt.Errors.Errors) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `Errors` {#ManagedInstanceLastAttempt.Errors.Errors}

[Output Only] The array of errors encountered while processing this
operation.


### Inputs for `Errors`

* `string` [`code`](#ManagedInstanceLastAttempt.Errors.Errors.code) = 1
* `string` [`location`](#ManagedInstanceLastAttempt.Errors.Errors.location) = 2
* `string` [`message`](#ManagedInstanceLastAttempt.Errors.Errors.message) = 3

### `code` {#ManagedInstanceLastAttempt.Errors.Errors.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] The error type identifier for this error.

### `location` {#ManagedInstanceLastAttempt.Errors.Errors.location}

| Property | Comments |
|----------|----------|
| Field Name | `location` |
| Type | `string` |

[Output Only] Indicates the field in the request that caused the error. This
property is optional.

### `message` {#ManagedInstanceLastAttempt.Errors.Errors.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] An optional, human-readable error message.

## Message `Errors` {#Operation.Error.Errors}

[Output Only] The array of errors encountered while processing this
operation.


### Inputs for `Errors`

* `string` [`code`](#Operation.Error.Errors.code) = 1
* `string` [`location`](#Operation.Error.Errors.location) = 2
* `string` [`message`](#Operation.Error.Errors.message) = 3

### `code` {#Operation.Error.Errors.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] The error type identifier for this error.

### `location` {#Operation.Error.Errors.location}

| Property | Comments |
|----------|----------|
| Field Name | `location` |
| Type | `string` |

[Output Only] Indicates the field in the request that caused the error. This
property is optional.

### `message` {#Operation.Error.Errors.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] An optional, human-readable error message.

## Message `Firewall` {#Firewall}

Represents a Firewall resource.


### Inputs for `Firewall`

* `repeated` [`compute.Firewall.Allowed`](gcp_compute.md#Firewall.Allowed) [`allowed`](#Firewall.allowed) = 1
* `string` [`creationTimestamp`](#Firewall.creationTimestamp) = 2
* `repeated` [`compute.Firewall.Denied`](gcp_compute.md#Firewall.Denied) [`denied`](#Firewall.denied) = 3
* `string` [`description`](#Firewall.description) = 4
* `repeated` `string` [`destinationRanges`](#Firewall.destinationRanges) = 5
* `string` [`direction`](#Firewall.direction) = 6
* `string` [`id`](#Firewall.id) = 7
* `string` [`kind`](#Firewall.kind) = 8
* `string` [`name`](#Firewall.name) = 9 (**Required**)
* `string` [`network`](#Firewall.network) = 10
* `int32` [`priority`](#Firewall.priority) = 11
* `string` [`selfLink`](#Firewall.selfLink) = 12
* `repeated` `string` [`sourceRanges`](#Firewall.sourceRanges) = 13
* `repeated` `string` [`sourceServiceAccounts`](#Firewall.sourceServiceAccounts) = 14
* `repeated` `string` [`sourceTags`](#Firewall.sourceTags) = 15
* `repeated` `string` [`targetServiceAccounts`](#Firewall.targetServiceAccounts) = 16
* `repeated` `string` [`targetTags`](#Firewall.targetTags) = 17

### `allowed` {#Firewall.allowed}

| Property | Comments |
|----------|----------|
| Field Name | `allowed` |
| Type | [`compute.Firewall.Allowed`](gcp_compute.md#Firewall.Allowed) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `creationTimestamp` {#Firewall.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `denied` {#Firewall.denied}

| Property | Comments |
|----------|----------|
| Field Name | `denied` |
| Type | [`compute.Firewall.Denied`](gcp_compute.md#Firewall.Denied) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `description` {#Firewall.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `destinationRanges` {#Firewall.destinationRanges}

| Property | Comments |
|----------|----------|
| Field Name | `destinationRanges` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

If destination ranges are specified, the firewall will apply only to traffic
that has destination IP address in these ranges. These ranges must be
expressed in CIDR format. Only IPv4 is supported.

### `direction` {#Firewall.direction}

| Property | Comments |
|----------|----------|
| Field Name | `direction` |
| Type | `string` |

Direction of traffic to which this firewall applies; default is INGRESS.
Note: For INGRESS traffic, it is NOT supported to specify destinationRanges;
For EGRESS traffic, it is NOT supported to specify sourceRanges OR
sourceTags.
Valid values:
    EGRESS
    INGRESS

### `id` {#Firewall.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Firewall.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#firewall for firewall
rules.

### `name` {#Firewall.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource; provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `network` {#Firewall.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

URL of the network resource for this firewall rule. If not specified when
creating a firewall rule, the default network is used:
global/networks/default
If you choose to specify this property, you can specify the network as a
full or partial URL. For example, the following are all valid URLs:
-
https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network
- projects/myproject/global/networks/my-network
- global/networks/default

### `priority` {#Firewall.priority}

| Property | Comments |
|----------|----------|
| Field Name | `priority` |
| Type | `int32` |

Priority for this rule. This is an integer between 0 and 65535, both
inclusive. When not specified, the value assumed is 1000. Relative
priorities determine precedence of conflicting rules. Lower value of
priority implies higher precedence (eg, a rule with priority 0 has higher
precedence than a rule with priority 1). DENY rules take precedence over
ALLOW rules having equal priority.

### `selfLink` {#Firewall.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sourceRanges` {#Firewall.sourceRanges}

| Property | Comments |
|----------|----------|
| Field Name | `sourceRanges` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

If source ranges are specified, the firewall will apply only to traffic that
has source IP address in these ranges. These ranges must be expressed in
CIDR format. One or both of sourceRanges and sourceTags may be set. If both
properties are set, the firewall will apply to traffic that has source IP
address within sourceRanges OR the source IP that belongs to a tag listed in
the sourceTags property. The connection does not need to match both
properties for the firewall to apply. Only IPv4 is supported.

### `sourceServiceAccounts` {#Firewall.sourceServiceAccounts}

| Property | Comments |
|----------|----------|
| Field Name | `sourceServiceAccounts` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

If source service accounts are specified, the firewall will apply only to
traffic originating from an instance with a service account in this list.
Source service accounts cannot be used to control traffic to an instance's
external IP address because service accounts are associated with an
instance, not an IP address. sourceRanges can be set at the same time as
sourceServiceAccounts. If both are set, the firewall will apply to traffic
that has source IP address within sourceRanges OR the source IP belongs to
an instance with service account listed in sourceServiceAccount. The
connection does not need to match both properties for the firewall to apply.
sourceServiceAccounts cannot be used at the same time as sourceTags or
targetTags.

### `sourceTags` {#Firewall.sourceTags}

| Property | Comments |
|----------|----------|
| Field Name | `sourceTags` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

If source tags are specified, the firewall rule applies only to traffic with
source IPs that match the primary network interfaces of VM instances that
have the tag and are in the same VPC network. Source tags cannot be used to
control traffic to an instance's external IP address, it only applies to
traffic between instances in the same virtual network. Because tags are
associated with instances, not IP addresses. One or both of sourceRanges and
sourceTags may be set. If both properties are set, the firewall will apply
to traffic that has source IP address within sourceRanges OR the source IP
that belongs to a tag listed in the sourceTags property. The connection does
not need to match both properties for the firewall to apply.

### `targetServiceAccounts` {#Firewall.targetServiceAccounts}

| Property | Comments |
|----------|----------|
| Field Name | `targetServiceAccounts` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of service accounts indicating sets of instances located in the
network that may make network connections as specified in allowed[].
targetServiceAccounts cannot be used at the same time as targetTags or
sourceTags. If neither targetServiceAccounts nor targetTags are specified,
the firewall rule applies to all instances on the specified network.

### `targetTags` {#Firewall.targetTags}

| Property | Comments |
|----------|----------|
| Field Name | `targetTags` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of tags that controls which instances the firewall rule applies to.
If targetTags are specified, then the firewall rule applies only to
instances in the VPC network that have one of those tags. If no targetTags
are specified, the firewall rule applies to all instances on the specified
network.

## Message `FirewallList` {#FirewallList}

Contains a list of firewalls.


### Inputs for `FirewallList`

* `string` [`id`](#FirewallList.id) = 1
* `repeated` [`compute.Firewall`](gcp_compute.md#Firewall) [`items`](#FirewallList.items) = 2
* `string` [`kind`](#FirewallList.kind) = 3
* `string` [`nextPageToken`](#FirewallList.nextPageToken) = 4
* `string` [`selfLink`](#FirewallList.selfLink) = 5
* [`compute.FirewallList.Warning`](gcp_compute.md#FirewallList.Warning) [`warning`](#FirewallList.warning) = 6

### `id` {#FirewallList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#FirewallList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Firewall`](gcp_compute.md#Firewall) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Firewall resources.

### `kind` {#FirewallList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#firewallList for lists of
firewalls.

### `nextPageToken` {#FirewallList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#FirewallList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#FirewallList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.FirewallList.Warning`](gcp_compute.md#FirewallList.Warning) |

## Message `ForwardingRule` {#ForwardingRule}

A ForwardingRule resource. A ForwardingRule resource specifies which pool of
target virtual machines to forward a packet to if it matches the given
[IPAddress, IPProtocol, ports] tuple. (== resource_for beta.forwardingRules
==) (== resource_for v1.forwardingRules ==) (== resource_for
beta.globalForwardingRules ==) (== resource_for v1.globalForwardingRules ==)
(== resource_for beta.regionForwardingRules ==) (== resource_for
v1.regionForwardingRules ==)


### Inputs for `ForwardingRule`

* `string` [`IPAddress`](#ForwardingRule.IPAddress) = 1
* `string` [`IPProtocol`](#ForwardingRule.IPProtocol) = 2
* `string` [`backendService`](#ForwardingRule.backendService) = 3
* `string` [`creationTimestamp`](#ForwardingRule.creationTimestamp) = 4
* `string` [`description`](#ForwardingRule.description) = 5
* `string` [`id`](#ForwardingRule.id) = 6
* `string` [`ipVersion`](#ForwardingRule.ipVersion) = 7
* `string` [`kind`](#ForwardingRule.kind) = 8
* `string` [`loadBalancingScheme`](#ForwardingRule.loadBalancingScheme) = 9
* `string` [`name`](#ForwardingRule.name) = 10 (**Required**)
* `string` [`network`](#ForwardingRule.network) = 11
* `string` [`portRange`](#ForwardingRule.portRange) = 12
* `repeated` `string` [`ports`](#ForwardingRule.ports) = 13
* `string` [`region`](#ForwardingRule.region) = 14
* `string` [`selfLink`](#ForwardingRule.selfLink) = 15
* `string` [`subnetwork`](#ForwardingRule.subnetwork) = 16
* `string` [`target`](#ForwardingRule.target) = 17

### `IPAddress` {#ForwardingRule.IPAddress}

| Property | Comments |
|----------|----------|
| Field Name | `IPAddress` |
| Type | `string` |

The IP address that this forwarding rule is serving on behalf of.

Addresses are restricted based on the forwarding rule's load balancing
scheme (EXTERNAL or INTERNAL) and scope (global or regional).

When the load balancing scheme is EXTERNAL, for global forwarding rules, the
address must be a global IP, and for regional forwarding rules, the address
must live in the same region as the forwarding rule. If this field is empty,
an ephemeral IPv4 address from the same scope (global or regional) will be
assigned. A regional forwarding rule supports IPv4 only. A global forwarding
rule supports either IPv4 or IPv6.

When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP
address belonging to the network/subnet configured for the forwarding rule.
By default, if this field is empty, an ephemeral internal IP address will be
automatically allocated from the IP range of the subnet or network
configured for this forwarding rule.

An address can be specified either by a literal IP address or a URL
reference to an existing Address resource. The following examples are all
valid:
- 100.1.2.3
-
https://www.googleapis.com/compute/v1/projects/project/regions/region/addresses/address
- projects/project/regions/region/addresses/address
- regions/region/addresses/address
- global/addresses/address
- address

### `IPProtocol` {#ForwardingRule.IPProtocol}

| Property | Comments |
|----------|----------|
| Field Name | `IPProtocol` |
| Type | `string` |

The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP,
AH, SCTP or ICMP.

When the load balancing scheme is INTERNAL, only TCP and UDP are valid.
Valid values:
    AH
    ESP
    ICMP
    SCTP
    TCP
    UDP

### `backendService` {#ForwardingRule.backendService}

| Property | Comments |
|----------|----------|
| Field Name | `backendService` |
| Type | `string` |

This field is not used for external load balancing.

For internal load balancing, this field identifies the BackendService
resource to receive the matched traffic.

### `creationTimestamp` {#ForwardingRule.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#ForwardingRule.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#ForwardingRule.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `ipVersion` {#ForwardingRule.ipVersion}

| Property | Comments |
|----------|----------|
| Field Name | `ipVersion` |
| Type | `string` |

The IP Version that will be used by this forwarding rule. Valid options are
IPV4 or IPV6. This can only be specified for a global forwarding rule.
Valid values:
    IPV4
    IPV6
    UNSPECIFIED_VERSION

### `kind` {#ForwardingRule.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#forwardingRule for
Forwarding Rule resources.

### `loadBalancingScheme` {#ForwardingRule.loadBalancingScheme}

| Property | Comments |
|----------|----------|
| Field Name | `loadBalancingScheme` |
| Type | `string` |

This signifies what the ForwardingRule will be used for and can only take
the following values: INTERNAL, EXTERNAL The value of INTERNAL means that
this will be used for Internal Network Load Balancing (TCP, UDP). The value
of EXTERNAL means that this will be used for External Load Balancing
(HTTP(S) LB, External TCP/UDP LB, SSL Proxy)
Valid values:
    EXTERNAL
    INTERNAL
    INVALID

### `name` {#ForwardingRule.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource; provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `network` {#ForwardingRule.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

This field is not used for external load balancing.

For internal load balancing, this field identifies the network that the load
balanced IP should belong to for this Forwarding Rule. If this field is not
specified, the default network will be used.

### `portRange` {#ForwardingRule.portRange}

| Property | Comments |
|----------|----------|
| Field Name | `portRange` |
| Type | `string` |

This field is used along with the target field for TargetHttpProxy,
TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
TargetPool, TargetInstance.

Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets addressed
to ports in the specified range will be forwarded to target. Forwarding
rules with the same [IPAddress, IPProtocol] pair must have disjoint port
ranges.

Some types of forwarding target have constraints on the acceptable ports:
- TargetHttpProxy: 80, 8080
- TargetHttpsProxy: 443
- TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688,
1883, 5222
- TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688,
1883, 5222
- TargetVpnGateway: 500, 4500

### `ports` {#ForwardingRule.ports}

| Property | Comments |
|----------|----------|
| Field Name | `ports` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

This field is used along with the backend_service field for internal load
balancing.

When the load balancing scheme is INTERNAL, a single port or a comma
separated list of ports can be configured. Only packets addressed to these
ports will be forwarded to the backends configured with this forwarding
rule.

You may specify a maximum of up to 5 ports.

### `region` {#ForwardingRule.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the regional forwarding rule resides.
This field is not applicable to global forwarding rules. You must specify
this field as part of the HTTP request URL. It is not settable as a field in
the request body.

### `selfLink` {#ForwardingRule.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `subnetwork` {#ForwardingRule.subnetwork}

| Property | Comments |
|----------|----------|
| Field Name | `subnetwork` |
| Type | `string` |

This field is not used for external load balancing.

For internal load balancing, this field identifies the subnetwork that the
load balanced IP should belong to for this Forwarding Rule.

If the network specified is in auto subnet mode, this field is optional.
However, if the network is in custom subnet mode, a subnetwork must be
specified.

### `target` {#ForwardingRule.target}

| Property | Comments |
|----------|----------|
| Field Name | `target` |
| Type | `string` |

The URL of the target resource to receive the matched traffic. For regional
forwarding rules, this target must live in the same region as the forwarding
rule. For global forwarding rules, this target must be a global load
balancing resource. The forwarded traffic must be of a type appropriate to
the target object.

## Message `ForwardingRuleAggregatedList` {#ForwardingRuleAggregatedList}



### Inputs for `ForwardingRuleAggregatedList`

* `string` [`id`](#ForwardingRuleAggregatedList.id) = 1
* `repeated` [`compute.ForwardingRuleAggregatedList.ItemsEntry`](gcp_compute.md#ForwardingRuleAggregatedList.ItemsEntry) [`items`](#ForwardingRuleAggregatedList.items) = 2
* `string` [`kind`](#ForwardingRuleAggregatedList.kind) = 3
* `string` [`nextPageToken`](#ForwardingRuleAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#ForwardingRuleAggregatedList.selfLink) = 5
* [`compute.ForwardingRuleAggregatedList.Warning`](gcp_compute.md#ForwardingRuleAggregatedList.Warning) [`warning`](#ForwardingRuleAggregatedList.warning) = 6

### `id` {#ForwardingRuleAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#ForwardingRuleAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.ForwardingRuleAggregatedList.ItemsEntry`](gcp_compute.md#ForwardingRuleAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of ForwardingRulesScopedList resources.

### `kind` {#ForwardingRuleAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#forwardingRuleAggregatedList
for lists of forwarding rules.

### `nextPageToken` {#ForwardingRuleAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#ForwardingRuleAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#ForwardingRuleAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.ForwardingRuleAggregatedList.Warning`](gcp_compute.md#ForwardingRuleAggregatedList.Warning) |

## Message `ForwardingRuleList` {#ForwardingRuleList}

Contains a list of ForwardingRule resources.


### Inputs for `ForwardingRuleList`

* `string` [`id`](#ForwardingRuleList.id) = 1
* `repeated` [`compute.ForwardingRule`](gcp_compute.md#ForwardingRule) [`items`](#ForwardingRuleList.items) = 2
* `string` [`kind`](#ForwardingRuleList.kind) = 3
* `string` [`nextPageToken`](#ForwardingRuleList.nextPageToken) = 4
* `string` [`selfLink`](#ForwardingRuleList.selfLink) = 5
* [`compute.ForwardingRuleList.Warning`](gcp_compute.md#ForwardingRuleList.Warning) [`warning`](#ForwardingRuleList.warning) = 6

### `id` {#ForwardingRuleList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#ForwardingRuleList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.ForwardingRule`](gcp_compute.md#ForwardingRule) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of ForwardingRule resources.

### `kind` {#ForwardingRuleList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#ForwardingRuleList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#ForwardingRuleList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#ForwardingRuleList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.ForwardingRuleList.Warning`](gcp_compute.md#ForwardingRuleList.Warning) |

## Message `ForwardingRulesScopedList` {#ForwardingRulesScopedList}



### Inputs for `ForwardingRulesScopedList`

* `repeated` [`compute.ForwardingRule`](gcp_compute.md#ForwardingRule) [`forwardingRules`](#ForwardingRulesScopedList.forwardingRules) = 1
* [`compute.ForwardingRulesScopedList.Warning`](gcp_compute.md#ForwardingRulesScopedList.Warning) [`warning`](#ForwardingRulesScopedList.warning) = 2

### `forwardingRules` {#ForwardingRulesScopedList.forwardingRules}

| Property | Comments |
|----------|----------|
| Field Name | `forwardingRules` |
| Type | [`compute.ForwardingRule`](gcp_compute.md#ForwardingRule) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of forwarding rules contained in this scope.

### `warning` {#ForwardingRulesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.ForwardingRulesScopedList.Warning`](gcp_compute.md#ForwardingRulesScopedList.Warning) |

## Message `GlobalSetLabelsRequest` {#GlobalSetLabelsRequest}



### Inputs for `GlobalSetLabelsRequest`

* `string` [`labelFingerprint`](#GlobalSetLabelsRequest.labelFingerprint) = 1
* `repeated` [`compute.GlobalSetLabelsRequest.LabelsEntry`](gcp_compute.md#GlobalSetLabelsRequest.LabelsEntry) [`labels`](#GlobalSetLabelsRequest.labels) = 2

### `labelFingerprint` {#GlobalSetLabelsRequest.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

The fingerprint of the previous set of labels for this resource, used to
detect conflicts. The fingerprint is initially generated by Compute Engine
and changes after every request to modify or update labels. You must always
provide an up-to-date fingerprint hash when updating or changing labels.
Make a get() request to the resource to get the latest fingerprint.

### `labels` {#GlobalSetLabelsRequest.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.GlobalSetLabelsRequest.LabelsEntry`](gcp_compute.md#GlobalSetLabelsRequest.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of labels to apply for this resource. Each label key & value must
comply with RFC1035. Specifically, the name must be 1-63 characters long and
match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the
first character must be a lowercase letter, and all following characters
must be a dash, lowercase letter, or digit, except the last character, which
cannot be a dash. For example, "webserver-frontend": "images". A label value
can also be empty (e.g. "my-label": "").

## Message `GuestOsFeature` {#GuestOsFeature}

Guest OS features.


### Inputs for `GuestOsFeature`

* `string` [`type`](#GuestOsFeature.type) = 1

### `type` {#GuestOsFeature.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

The ID of a supported feature. Read  Enabling guest operating system
features to see a list of available options.
Valid values:
    FEATURE_TYPE_UNSPECIFIED
    MULTI_IP_SUBNET
    SECURE_BOOT
    UEFI_COMPATIBLE
    VIRTIO_SCSI_MULTIQUEUE
    WINDOWS

## Message `HTTPHealthCheck` {#HTTPHealthCheck}



### Inputs for `HTTPHealthCheck`

* `string` [`host`](#HTTPHealthCheck.host) = 1
* `int32` [`port`](#HTTPHealthCheck.port) = 2
* `string` [`portName`](#HTTPHealthCheck.portName) = 3
* `string` [`proxyHeader`](#HTTPHealthCheck.proxyHeader) = 4
* `string` [`requestPath`](#HTTPHealthCheck.requestPath) = 5

### `host` {#HTTPHealthCheck.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

The value of the host header in the HTTP health check request. If left empty
(default value), the IP on behalf of which this health check is performed
will be used.

### `port` {#HTTPHealthCheck.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The TCP port number for the health check request. The default value is 80.
Valid values are 1 through 65535.

### `portName` {#HTTPHealthCheck.portName}

| Property | Comments |
|----------|----------|
| Field Name | `portName` |
| Type | `string` |

Port name as defined in InstanceGroup#NamedPort#name. If both port and
port_name are defined, port takes precedence.

### `proxyHeader` {#HTTPHealthCheck.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

Specifies the type of proxy header to append before sending data to the
backend, either NONE or PROXY_V1. The default is NONE.
Valid values:
    NONE
    PROXY_V1

### `requestPath` {#HTTPHealthCheck.requestPath}

| Property | Comments |
|----------|----------|
| Field Name | `requestPath` |
| Type | `string` |

The request path of the HTTP health check request. The default value is /.

## Message `HTTPSHealthCheck` {#HTTPSHealthCheck}



### Inputs for `HTTPSHealthCheck`

* `string` [`host`](#HTTPSHealthCheck.host) = 1
* `int32` [`port`](#HTTPSHealthCheck.port) = 2
* `string` [`portName`](#HTTPSHealthCheck.portName) = 3
* `string` [`proxyHeader`](#HTTPSHealthCheck.proxyHeader) = 4
* `string` [`requestPath`](#HTTPSHealthCheck.requestPath) = 5

### `host` {#HTTPSHealthCheck.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

The value of the host header in the HTTPS health check request. If left
empty (default value), the IP on behalf of which this health check is
performed will be used.

### `port` {#HTTPSHealthCheck.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The TCP port number for the health check request. The default value is 443.
Valid values are 1 through 65535.

### `portName` {#HTTPSHealthCheck.portName}

| Property | Comments |
|----------|----------|
| Field Name | `portName` |
| Type | `string` |

Port name as defined in InstanceGroup#NamedPort#name. If both port and
port_name are defined, port takes precedence.

### `proxyHeader` {#HTTPSHealthCheck.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

Specifies the type of proxy header to append before sending data to the
backend, either NONE or PROXY_V1. The default is NONE.
Valid values:
    NONE
    PROXY_V1

### `requestPath` {#HTTPSHealthCheck.requestPath}

| Property | Comments |
|----------|----------|
| Field Name | `requestPath` |
| Type | `string` |

The request path of the HTTPS health check request. The default value is /.

## Message `HealthCheck` {#HealthCheck}

An HealthCheck resource. This resource defines a template for how individual
virtual machines should be checked for health, via one of the supported
protocols.


### Inputs for `HealthCheck`

* `int32` [`checkIntervalSec`](#HealthCheck.checkIntervalSec) = 1
* `string` [`creationTimestamp`](#HealthCheck.creationTimestamp) = 2
* `string` [`description`](#HealthCheck.description) = 3
* `int32` [`healthyThreshold`](#HealthCheck.healthyThreshold) = 4
* [`compute.HTTPHealthCheck`](gcp_compute.md#HTTPHealthCheck) [`httpHealthCheck`](#HealthCheck.httpHealthCheck) = 5
* [`compute.HTTPSHealthCheck`](gcp_compute.md#HTTPSHealthCheck) [`httpsHealthCheck`](#HealthCheck.httpsHealthCheck) = 6
* `string` [`id`](#HealthCheck.id) = 7
* `string` [`kind`](#HealthCheck.kind) = 8
* `string` [`name`](#HealthCheck.name) = 9 (**Required**)
* `string` [`selfLink`](#HealthCheck.selfLink) = 10
* [`compute.SSLHealthCheck`](gcp_compute.md#SSLHealthCheck) [`sslHealthCheck`](#HealthCheck.sslHealthCheck) = 11
* [`compute.TCPHealthCheck`](gcp_compute.md#TCPHealthCheck) [`tcpHealthCheck`](#HealthCheck.tcpHealthCheck) = 12
* `int32` [`timeoutSec`](#HealthCheck.timeoutSec) = 13
* `string` [`type`](#HealthCheck.type) = 14
* `int32` [`unhealthyThreshold`](#HealthCheck.unhealthyThreshold) = 15

### `checkIntervalSec` {#HealthCheck.checkIntervalSec}

| Property | Comments |
|----------|----------|
| Field Name | `checkIntervalSec` |
| Type | `int32` |

How often (in seconds) to send a health check. The default value is 5
seconds.

### `creationTimestamp` {#HealthCheck.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in 3339 text format.

### `description` {#HealthCheck.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `healthyThreshold` {#HealthCheck.healthyThreshold}

| Property | Comments |
|----------|----------|
| Field Name | `healthyThreshold` |
| Type | `int32` |

A so-far unhealthy instance will be marked healthy after this many
consecutive successes. The default value is 2.

### `httpHealthCheck` {#HealthCheck.httpHealthCheck}

| Property | Comments |
|----------|----------|
| Field Name | `httpHealthCheck` |
| Type | [`compute.HTTPHealthCheck`](gcp_compute.md#HTTPHealthCheck) |

### `httpsHealthCheck` {#HealthCheck.httpsHealthCheck}

| Property | Comments |
|----------|----------|
| Field Name | `httpsHealthCheck` |
| Type | [`compute.HTTPSHealthCheck`](gcp_compute.md#HTTPSHealthCheck) |

### `id` {#HealthCheck.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#HealthCheck.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of the resource.

### `name` {#HealthCheck.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `selfLink` {#HealthCheck.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sslHealthCheck` {#HealthCheck.sslHealthCheck}

| Property | Comments |
|----------|----------|
| Field Name | `sslHealthCheck` |
| Type | [`compute.SSLHealthCheck`](gcp_compute.md#SSLHealthCheck) |

### `tcpHealthCheck` {#HealthCheck.tcpHealthCheck}

| Property | Comments |
|----------|----------|
| Field Name | `tcpHealthCheck` |
| Type | [`compute.TCPHealthCheck`](gcp_compute.md#TCPHealthCheck) |

### `timeoutSec` {#HealthCheck.timeoutSec}

| Property | Comments |
|----------|----------|
| Field Name | `timeoutSec` |
| Type | `int32` |

How long (in seconds) to wait before claiming failure. The default value is
5 seconds. It is invalid for timeoutSec to have greater value than
checkIntervalSec.

### `type` {#HealthCheck.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

Specifies the type of the healthCheck, either TCP, SSL, HTTP or HTTPS. If
not specified, the default is TCP. Exactly one of the protocol-specific
health check field must be specified, which must match type field.
Valid values:
    HTTP
    HTTPS
    INVALID
    SSL
    TCP

### `unhealthyThreshold` {#HealthCheck.unhealthyThreshold}

| Property | Comments |
|----------|----------|
| Field Name | `unhealthyThreshold` |
| Type | `int32` |

A so-far healthy instance will be marked unhealthy after this many
consecutive failures. The default value is 2.

## Message `HealthCheckList` {#HealthCheckList}

Contains a list of HealthCheck resources.


### Inputs for `HealthCheckList`

* `string` [`id`](#HealthCheckList.id) = 1
* `repeated` [`compute.HealthCheck`](gcp_compute.md#HealthCheck) [`items`](#HealthCheckList.items) = 2
* `string` [`kind`](#HealthCheckList.kind) = 3
* `string` [`nextPageToken`](#HealthCheckList.nextPageToken) = 4
* `string` [`selfLink`](#HealthCheckList.selfLink) = 5
* [`compute.HealthCheckList.Warning`](gcp_compute.md#HealthCheckList.Warning) [`warning`](#HealthCheckList.warning) = 6

### `id` {#HealthCheckList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#HealthCheckList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.HealthCheck`](gcp_compute.md#HealthCheck) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of HealthCheck resources.

### `kind` {#HealthCheckList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#HealthCheckList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#HealthCheckList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#HealthCheckList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.HealthCheckList.Warning`](gcp_compute.md#HealthCheckList.Warning) |

## Message `HealthCheckReference` {#HealthCheckReference}

A full or valid partial URL to a health check. For example, the following
are valid URLs:
-
https://www.googleapis.com/compute/beta/projects/project-id/global/httpHealthChecks/health-check
- projects/project-id/global/httpHealthChecks/health-check
- global/httpHealthChecks/health-check


### Inputs for `HealthCheckReference`

* `string` [`healthCheck`](#HealthCheckReference.healthCheck) = 1

### `healthCheck` {#HealthCheckReference.healthCheck}

| Property | Comments |
|----------|----------|
| Field Name | `healthCheck` |
| Type | `string` |

## Message `HealthStatus` {#HealthStatus}



### Inputs for `HealthStatus`

* `string` [`healthState`](#HealthStatus.healthState) = 1
* `string` [`instance`](#HealthStatus.instance) = 2
* `string` [`ipAddress`](#HealthStatus.ipAddress) = 3
* `int32` [`port`](#HealthStatus.port) = 4

### `healthState` {#HealthStatus.healthState}

| Property | Comments |
|----------|----------|
| Field Name | `healthState` |
| Type | `string` |

Health state of the instance.
Valid values:
    HEALTHY
    UNHEALTHY

### `instance` {#HealthStatus.instance}

| Property | Comments |
|----------|----------|
| Field Name | `instance` |
| Type | `string` |

URL of the instance resource.

### `ipAddress` {#HealthStatus.ipAddress}

| Property | Comments |
|----------|----------|
| Field Name | `ipAddress` |
| Type | `string` |

The IP address represented by this resource.

### `port` {#HealthStatus.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The port on the instance.

## Message `HostRule` {#HostRule}

UrlMaps A host-matching rule for a URL. If matched, will use the named
PathMatcher to select the BackendService.


### Inputs for `HostRule`

* `string` [`description`](#HostRule.description) = 1
* `repeated` `string` [`hosts`](#HostRule.hosts) = 2
* `string` [`pathMatcher`](#HostRule.pathMatcher) = 3

### `description` {#HostRule.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `hosts` {#HostRule.hosts}

| Property | Comments |
|----------|----------|
| Field Name | `hosts` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of host patterns to match. They must be valid hostnames, except *
will match any string of ([a-z0-9-.]*). In that case, * must be the first
character and must be followed in the pattern by either - or ..

### `pathMatcher` {#HostRule.pathMatcher}

| Property | Comments |
|----------|----------|
| Field Name | `pathMatcher` |
| Type | `string` |

The name of the PathMatcher to use to match the path portion of the URL if
the hostRule matches the URL's host portion.

## Message `HttpHealthCheck` {#HttpHealthCheck}

An HttpHealthCheck resource. This resource defines a template for how
individual instances should be checked for health, via HTTP.


### Inputs for `HttpHealthCheck`

* `int32` [`checkIntervalSec`](#HttpHealthCheck.checkIntervalSec) = 1
* `string` [`creationTimestamp`](#HttpHealthCheck.creationTimestamp) = 2
* `string` [`description`](#HttpHealthCheck.description) = 3
* `int32` [`healthyThreshold`](#HttpHealthCheck.healthyThreshold) = 4
* `string` [`host`](#HttpHealthCheck.host) = 5
* `string` [`id`](#HttpHealthCheck.id) = 6
* `string` [`kind`](#HttpHealthCheck.kind) = 7
* `string` [`name`](#HttpHealthCheck.name) = 8 (**Required**)
* `int32` [`port`](#HttpHealthCheck.port) = 9
* `string` [`requestPath`](#HttpHealthCheck.requestPath) = 10
* `string` [`selfLink`](#HttpHealthCheck.selfLink) = 11
* `int32` [`timeoutSec`](#HttpHealthCheck.timeoutSec) = 12
* `int32` [`unhealthyThreshold`](#HttpHealthCheck.unhealthyThreshold) = 13

### `checkIntervalSec` {#HttpHealthCheck.checkIntervalSec}

| Property | Comments |
|----------|----------|
| Field Name | `checkIntervalSec` |
| Type | `int32` |

How often (in seconds) to send a health check. The default value is 5
seconds.

### `creationTimestamp` {#HttpHealthCheck.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#HttpHealthCheck.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `healthyThreshold` {#HttpHealthCheck.healthyThreshold}

| Property | Comments |
|----------|----------|
| Field Name | `healthyThreshold` |
| Type | `int32` |

A so-far unhealthy instance will be marked healthy after this many
consecutive successes. The default value is 2.

### `host` {#HttpHealthCheck.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

The value of the host header in the HTTP health check request. If left empty
(default value), the public IP on behalf of which this health check is
performed will be used.

### `id` {#HttpHealthCheck.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#HttpHealthCheck.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#httpHealthCheck for HTTP
health checks.

### `name` {#HttpHealthCheck.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `port` {#HttpHealthCheck.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The TCP port number for the HTTP health check request. The default value is
80.

### `requestPath` {#HttpHealthCheck.requestPath}

| Property | Comments |
|----------|----------|
| Field Name | `requestPath` |
| Type | `string` |

The request path of the HTTP health check request. The default value is /.

### `selfLink` {#HttpHealthCheck.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `timeoutSec` {#HttpHealthCheck.timeoutSec}

| Property | Comments |
|----------|----------|
| Field Name | `timeoutSec` |
| Type | `int32` |

How long (in seconds) to wait before claiming failure. The default value is
5 seconds. It is invalid for timeoutSec to have greater value than
checkIntervalSec.

### `unhealthyThreshold` {#HttpHealthCheck.unhealthyThreshold}

| Property | Comments |
|----------|----------|
| Field Name | `unhealthyThreshold` |
| Type | `int32` |

A so-far healthy instance will be marked unhealthy after this many
consecutive failures. The default value is 2.

## Message `HttpHealthCheckList` {#HttpHealthCheckList}

Contains a list of HttpHealthCheck resources.


### Inputs for `HttpHealthCheckList`

* `string` [`id`](#HttpHealthCheckList.id) = 1
* `repeated` [`compute.HttpHealthCheck`](gcp_compute.md#HttpHealthCheck) [`items`](#HttpHealthCheckList.items) = 2
* `string` [`kind`](#HttpHealthCheckList.kind) = 3
* `string` [`nextPageToken`](#HttpHealthCheckList.nextPageToken) = 4
* `string` [`selfLink`](#HttpHealthCheckList.selfLink) = 5
* [`compute.HttpHealthCheckList.Warning`](gcp_compute.md#HttpHealthCheckList.Warning) [`warning`](#HttpHealthCheckList.warning) = 6

### `id` {#HttpHealthCheckList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#HttpHealthCheckList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.HttpHealthCheck`](gcp_compute.md#HttpHealthCheck) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of HttpHealthCheck resources.

### `kind` {#HttpHealthCheckList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#HttpHealthCheckList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#HttpHealthCheckList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#HttpHealthCheckList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.HttpHealthCheckList.Warning`](gcp_compute.md#HttpHealthCheckList.Warning) |

## Message `HttpsHealthCheck` {#HttpsHealthCheck}

An HttpsHealthCheck resource. This resource defines a template for how
individual instances should be checked for health, via HTTPS.


### Inputs for `HttpsHealthCheck`

* `int32` [`checkIntervalSec`](#HttpsHealthCheck.checkIntervalSec) = 1
* `string` [`creationTimestamp`](#HttpsHealthCheck.creationTimestamp) = 2
* `string` [`description`](#HttpsHealthCheck.description) = 3
* `int32` [`healthyThreshold`](#HttpsHealthCheck.healthyThreshold) = 4
* `string` [`host`](#HttpsHealthCheck.host) = 5
* `string` [`id`](#HttpsHealthCheck.id) = 6
* `string` [`kind`](#HttpsHealthCheck.kind) = 7
* `string` [`name`](#HttpsHealthCheck.name) = 8 (**Required**)
* `int32` [`port`](#HttpsHealthCheck.port) = 9
* `string` [`requestPath`](#HttpsHealthCheck.requestPath) = 10
* `string` [`selfLink`](#HttpsHealthCheck.selfLink) = 11
* `int32` [`timeoutSec`](#HttpsHealthCheck.timeoutSec) = 12
* `int32` [`unhealthyThreshold`](#HttpsHealthCheck.unhealthyThreshold) = 13

### `checkIntervalSec` {#HttpsHealthCheck.checkIntervalSec}

| Property | Comments |
|----------|----------|
| Field Name | `checkIntervalSec` |
| Type | `int32` |

How often (in seconds) to send a health check. The default value is 5
seconds.

### `creationTimestamp` {#HttpsHealthCheck.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#HttpsHealthCheck.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `healthyThreshold` {#HttpsHealthCheck.healthyThreshold}

| Property | Comments |
|----------|----------|
| Field Name | `healthyThreshold` |
| Type | `int32` |

A so-far unhealthy instance will be marked healthy after this many
consecutive successes. The default value is 2.

### `host` {#HttpsHealthCheck.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

The value of the host header in the HTTPS health check request. If left
empty (default value), the public IP on behalf of which this health check is
performed will be used.

### `id` {#HttpsHealthCheck.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#HttpsHealthCheck.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of the resource.

### `name` {#HttpsHealthCheck.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `port` {#HttpsHealthCheck.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The TCP port number for the HTTPS health check request. The default value is
443.

### `requestPath` {#HttpsHealthCheck.requestPath}

| Property | Comments |
|----------|----------|
| Field Name | `requestPath` |
| Type | `string` |

The request path of the HTTPS health check request. The default value is
"/".

### `selfLink` {#HttpsHealthCheck.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `timeoutSec` {#HttpsHealthCheck.timeoutSec}

| Property | Comments |
|----------|----------|
| Field Name | `timeoutSec` |
| Type | `int32` |

How long (in seconds) to wait before claiming failure. The default value is
5 seconds. It is invalid for timeoutSec to have a greater value than
checkIntervalSec.

### `unhealthyThreshold` {#HttpsHealthCheck.unhealthyThreshold}

| Property | Comments |
|----------|----------|
| Field Name | `unhealthyThreshold` |
| Type | `int32` |

A so-far healthy instance will be marked unhealthy after this many
consecutive failures. The default value is 2.

## Message `HttpsHealthCheckList` {#HttpsHealthCheckList}

Contains a list of HttpsHealthCheck resources.


### Inputs for `HttpsHealthCheckList`

* `string` [`id`](#HttpsHealthCheckList.id) = 1
* `repeated` [`compute.HttpsHealthCheck`](gcp_compute.md#HttpsHealthCheck) [`items`](#HttpsHealthCheckList.items) = 2
* `string` [`kind`](#HttpsHealthCheckList.kind) = 3
* `string` [`nextPageToken`](#HttpsHealthCheckList.nextPageToken) = 4
* `string` [`selfLink`](#HttpsHealthCheckList.selfLink) = 5
* [`compute.HttpsHealthCheckList.Warning`](gcp_compute.md#HttpsHealthCheckList.Warning) [`warning`](#HttpsHealthCheckList.warning) = 6

### `id` {#HttpsHealthCheckList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#HttpsHealthCheckList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.HttpsHealthCheck`](gcp_compute.md#HttpsHealthCheck) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of HttpsHealthCheck resources.

### `kind` {#HttpsHealthCheckList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#HttpsHealthCheckList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#HttpsHealthCheckList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#HttpsHealthCheckList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.HttpsHealthCheckList.Warning`](gcp_compute.md#HttpsHealthCheckList.Warning) |

## Message `Image` {#Image}

An Image resource. (== resource_for beta.images ==) (== resource_for
v1.images ==)


### Inputs for `Image`

* `string` [`archiveSizeBytes`](#Image.archiveSizeBytes) = 1
* `string` [`creationTimestamp`](#Image.creationTimestamp) = 2
* [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) [`deprecated`](#Image.deprecated) = 3
* `string` [`description`](#Image.description) = 4
* `string` [`diskSizeGb`](#Image.diskSizeGb) = 5
* `string` [`family`](#Image.family) = 6
* `repeated` [`compute.GuestOsFeature`](gcp_compute.md#GuestOsFeature) [`guestOsFeatures`](#Image.guestOsFeatures) = 7
* `string` [`id`](#Image.id) = 8
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`imageEncryptionKey`](#Image.imageEncryptionKey) = 9
* `string` [`kind`](#Image.kind) = 10
* `string` [`labelFingerprint`](#Image.labelFingerprint) = 11
* `repeated` [`compute.Image.LabelsEntry`](gcp_compute.md#Image.LabelsEntry) [`labels`](#Image.labels) = 12
* `repeated` `string` [`licenseCodes`](#Image.licenseCodes) = 13
* `repeated` `string` [`licenses`](#Image.licenses) = 14
* `string` [`name`](#Image.name) = 15 (**Required**)
* [`compute.Image.RawDisk`](gcp_compute.md#Image.RawDisk) [`rawDisk`](#Image.rawDisk) = 16
* `string` [`selfLink`](#Image.selfLink) = 17
* `string` [`sourceDisk`](#Image.sourceDisk) = 18
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceDiskEncryptionKey`](#Image.sourceDiskEncryptionKey) = 19
* `string` [`sourceDiskId`](#Image.sourceDiskId) = 20
* `string` [`sourceImage`](#Image.sourceImage) = 21
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceImageEncryptionKey`](#Image.sourceImageEncryptionKey) = 22
* `string` [`sourceImageId`](#Image.sourceImageId) = 23
* `string` [`sourceSnapshot`](#Image.sourceSnapshot) = 24
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceSnapshotEncryptionKey`](#Image.sourceSnapshotEncryptionKey) = 25
* `string` [`sourceSnapshotId`](#Image.sourceSnapshotId) = 26
* `string` [`sourceType`](#Image.sourceType) = 27
* `string` [`status`](#Image.status) = 28

### `archiveSizeBytes` {#Image.archiveSizeBytes}

| Property | Comments |
|----------|----------|
| Field Name | `archiveSizeBytes` |
| Type | `string` |

Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).

### `creationTimestamp` {#Image.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `deprecated` {#Image.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) |

The deprecation status associated with this image.

### `description` {#Image.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `diskSizeGb` {#Image.diskSizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `diskSizeGb` |
| Type | `string` |

Size of the image when restored onto a persistent disk (in GB).

### `family` {#Image.family}

| Property | Comments |
|----------|----------|
| Field Name | `family` |
| Type | `string` |

The name of the image family to which this image belongs. You can create
disks by specifying an image family instead of a specific image name. The
image family always returns its latest image that is not deprecated. The
name of the image family must comply with RFC1035.

### `guestOsFeatures` {#Image.guestOsFeatures}

| Property | Comments |
|----------|----------|
| Field Name | `guestOsFeatures` |
| Type | [`compute.GuestOsFeature`](gcp_compute.md#GuestOsFeature) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of features to enable on the guest operating system. Applicable only
for bootable images. Read  Enabling guest operating system features to see a
list of available options.

### `id` {#Image.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `imageEncryptionKey` {#Image.imageEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `imageEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

Encrypts the image using a customer-supplied encryption key.

After you encrypt an image with a customer-supplied key, you must provide
the same key if you use the image later (e.g. to create a disk from the
image).

Customer-supplied encryption keys do not protect access to metadata of the
disk.

If you do not provide an encryption key when creating the image, then the
disk will be encrypted using an automatically generated key and you do not
need to provide a key to use the image later.

### `kind` {#Image.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#image for images.

### `labelFingerprint` {#Image.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

A fingerprint for the labels being applied to this image, which is
essentially a hash of the labels used for optimistic locking. The
fingerprint is initially generated by Compute Engine and changes after every
request to modify or update labels. You must always provide an up-to-date
fingerprint hash in order to update or change labels.

To see the latest fingerprint, make a get() request to retrieve an image.

### `labels` {#Image.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.Image.LabelsEntry`](gcp_compute.md#Image.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Labels to apply to this image. These can be later modified by the setLabels
method.

### `licenseCodes` {#Image.licenseCodes}

| Property | Comments |
|----------|----------|
| Field Name | `licenseCodes` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Integer license codes indicating which licenses are attached to this image.

### `licenses` {#Image.licenses}

| Property | Comments |
|----------|----------|
| Field Name | `licenses` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Any applicable license URI.

### `name` {#Image.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource; provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `rawDisk` {#Image.rawDisk}

| Property | Comments |
|----------|----------|
| Field Name | `rawDisk` |
| Type | [`compute.Image.RawDisk`](gcp_compute.md#Image.RawDisk) |

### `selfLink` {#Image.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sourceDisk` {#Image.sourceDisk}

| Property | Comments |
|----------|----------|
| Field Name | `sourceDisk` |
| Type | `string` |

URL of the source disk used to create this image. This can be a full or
valid partial URL. You must provide either this property or the
rawDisk.source property but not both to create an image. For example, the
following are valid values:
-
https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
- projects/project/zones/zone/disks/disk
- zones/zone/disks/disk

### `sourceDiskEncryptionKey` {#Image.sourceDiskEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceDiskEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source disk. Required if the
source disk is protected by a customer-supplied encryption key.

### `sourceDiskId` {#Image.sourceDiskId}

| Property | Comments |
|----------|----------|
| Field Name | `sourceDiskId` |
| Type | `string` |

The ID value of the disk used to create this image. This value may be used
to determine whether the image was taken from the current or a previous
instance of a given disk name.

### `sourceImage` {#Image.sourceImage}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImage` |
| Type | `string` |

URL of the source image used to create this image. This can be a full or
valid partial URL. You must provide exactly one of:
- this property, or
- the rawDisk.source property, or
- the sourceDisk property   in order to create an image.

### `sourceImageEncryptionKey` {#Image.sourceImageEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImageEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source image. Required if the
source image is protected by a customer-supplied encryption key.

### `sourceImageId` {#Image.sourceImageId}

| Property | Comments |
|----------|----------|
| Field Name | `sourceImageId` |
| Type | `string` |

[Output Only] The ID value of the image used to create this image. This
value may be used to determine whether the image was taken from the current
or a previous instance of a given image name.

### `sourceSnapshot` {#Image.sourceSnapshot}

| Property | Comments |
|----------|----------|
| Field Name | `sourceSnapshot` |
| Type | `string` |

URL of the source snapshot used to create this image. This can be a full or
valid partial URL. You must provide exactly one of:
- this property, or
- the sourceImage property, or
- the rawDisk.source property, or
- the sourceDisk property   in order to create an image.

### `sourceSnapshotEncryptionKey` {#Image.sourceSnapshotEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceSnapshotEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source snapshot. Required if the
source snapshot is protected by a customer-supplied encryption key.

### `sourceSnapshotId` {#Image.sourceSnapshotId}

| Property | Comments |
|----------|----------|
| Field Name | `sourceSnapshotId` |
| Type | `string` |

[Output Only] The ID value of the snapshot used to create this image. This
value may be used to determine whether the snapshot was taken from the
current or a previous instance of a given snapshot name.

### `sourceType` {#Image.sourceType}

| Property | Comments |
|----------|----------|
| Field Name | `sourceType` |
| Type | `string` |

The type of the image used to create this disk. The default and only value
is RAW
Valid values:
    RAW

### `status` {#Image.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the image. An image can be used to create other
resources, such as instances, only after the image has been successfully
created and the status is set to READY. Possible values are FAILED, PENDING,
or READY.
Valid values:
    FAILED
    PENDING
    READY

## Message `ImageList` {#ImageList}

Contains a list of images.


### Inputs for `ImageList`

* `string` [`id`](#ImageList.id) = 1
* `repeated` [`compute.Image`](gcp_compute.md#Image) [`items`](#ImageList.items) = 2
* `string` [`kind`](#ImageList.kind) = 3
* `string` [`nextPageToken`](#ImageList.nextPageToken) = 4
* `string` [`selfLink`](#ImageList.selfLink) = 5
* [`compute.ImageList.Warning`](gcp_compute.md#ImageList.Warning) [`warning`](#ImageList.warning) = 6

### `id` {#ImageList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#ImageList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Image`](gcp_compute.md#Image) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Image resources.

### `kind` {#ImageList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#ImageList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#ImageList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#ImageList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.ImageList.Warning`](gcp_compute.md#ImageList.Warning) |

## Message `Instance` {#Instance}

An Instance resource. (== resource_for beta.instances ==) (== resource_for
v1.instances ==)


### Inputs for `Instance`

* `bool` [`canIpForward`](#Instance.canIpForward) = 1
* `string` [`cpuPlatform`](#Instance.cpuPlatform) = 2
* `string` [`creationTimestamp`](#Instance.creationTimestamp) = 3
* `bool` [`deletionProtection`](#Instance.deletionProtection) = 4
* `string` [`description`](#Instance.description) = 5
* `repeated` [`compute.AttachedDisk`](gcp_compute.md#AttachedDisk) [`disks`](#Instance.disks) = 6
* `repeated` [`compute.AcceleratorConfig`](gcp_compute.md#AcceleratorConfig) [`guestAccelerators`](#Instance.guestAccelerators) = 7
* `string` [`id`](#Instance.id) = 8
* `string` [`kind`](#Instance.kind) = 9
* `string` [`labelFingerprint`](#Instance.labelFingerprint) = 10
* `repeated` [`compute.Instance.LabelsEntry`](gcp_compute.md#Instance.LabelsEntry) [`labels`](#Instance.labels) = 11
* `string` [`machineType`](#Instance.machineType) = 12
* [`compute.Metadata`](gcp_compute.md#Metadata) [`metadata`](#Instance.metadata) = 13
* `string` [`minCpuPlatform`](#Instance.minCpuPlatform) = 14
* `string` [`name`](#Instance.name) = 15 (**Required**)
* `repeated` [`compute.NetworkInterface`](gcp_compute.md#NetworkInterface) [`networkInterfaces`](#Instance.networkInterfaces) = 16
* [`compute.Scheduling`](gcp_compute.md#Scheduling) [`scheduling`](#Instance.scheduling) = 17
* `string` [`selfLink`](#Instance.selfLink) = 18
* `repeated` [`compute.ServiceAccount`](gcp_compute.md#ServiceAccount) [`serviceAccounts`](#Instance.serviceAccounts) = 19
* `bool` [`startRestricted`](#Instance.startRestricted) = 20
* `string` [`status`](#Instance.status) = 21
* `string` [`statusMessage`](#Instance.statusMessage) = 22
* [`compute.Tags`](gcp_compute.md#Tags) [`tags`](#Instance.tags) = 23
* `string` [`zone`](#Instance.zone) = 24

### `canIpForward` {#Instance.canIpForward}

| Property | Comments |
|----------|----------|
| Field Name | `canIpForward` |
| Type | `bool` |

Allows this instance to send and receive packets with non-matching
destination or source IPs. This is required if you plan to use this instance
to forward routes. For more information, see Enabling IP Forwarding.

### `cpuPlatform` {#Instance.cpuPlatform}

| Property | Comments |
|----------|----------|
| Field Name | `cpuPlatform` |
| Type | `string` |

[Output Only] The CPU platform used by this instance.

### `creationTimestamp` {#Instance.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `deletionProtection` {#Instance.deletionProtection}

| Property | Comments |
|----------|----------|
| Field Name | `deletionProtection` |
| Type | `bool` |

Whether the resource should be protected against deletion.

### `description` {#Instance.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `disks` {#Instance.disks}

| Property | Comments |
|----------|----------|
| Field Name | `disks` |
| Type | [`compute.AttachedDisk`](gcp_compute.md#AttachedDisk) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Array of disks associated with this instance. Persistent disks must be
created before you can assign them.

### `guestAccelerators` {#Instance.guestAccelerators}

| Property | Comments |
|----------|----------|
| Field Name | `guestAccelerators` |
| Type | [`compute.AcceleratorConfig`](gcp_compute.md#AcceleratorConfig) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of the type and count of accelerator cards attached to the instance.

### `id` {#Instance.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Instance.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#instance for instances.

### `labelFingerprint` {#Instance.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

A fingerprint for this request, which is essentially a hash of the label's
contents and used for optimistic locking. The fingerprint is initially
generated by Compute Engine and changes after every request to modify or
update labels. You must always provide an up-to-date fingerprint hash in
order to update or change labels.

To see the latest fingerprint, make get() request to the instance.

### `labels` {#Instance.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.Instance.LabelsEntry`](gcp_compute.md#Instance.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Labels to apply to this instance. These can be later modified by the
setLabels method.

### `machineType` {#Instance.machineType}

| Property | Comments |
|----------|----------|
| Field Name | `machineType` |
| Type | `string` |

Full or partial URL of the machine type resource to use for this instance,
in the format: zones/zone/machineTypes/machine-type. This is provided by the
client when the instance is created. For example, the following is a valid
partial url to a predefined machine type:
zones/us-central1-f/machineTypes/n1-standard-1


To create a custom machine type, provide a URL to a machine type in the
following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ...
24, etc), and MEMORY is the total memory for this instance. Memory must be a
multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120
MB):
zones/zone/machineTypes/custom-CPUS-MEMORY


For example: zones/us-central1-f/machineTypes/custom-4-5120

For a full list of restrictions, read the Specifications for custom machine
types.

### `metadata` {#Instance.metadata}

| Property | Comments |
|----------|----------|
| Field Name | `metadata` |
| Type | [`compute.Metadata`](gcp_compute.md#Metadata) |

The metadata key/value pairs assigned to this instance. This includes custom
metadata and predefined keys.

### `minCpuPlatform` {#Instance.minCpuPlatform}

| Property | Comments |
|----------|----------|
| Field Name | `minCpuPlatform` |
| Type | `string` |

Specifies a minimum CPU platform for the VM instance. Applicable values are
the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell"
or minCpuPlatform: "Intel Sandy Bridge".

### `name` {#Instance.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name of the resource, provided by the client when initially creating the
resource. The resource name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match the
regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first
character must be a lowercase letter, and all following characters must be a
dash, lowercase letter, or digit, except the last character, which cannot be
a dash.

### `networkInterfaces` {#Instance.networkInterfaces}

| Property | Comments |
|----------|----------|
| Field Name | `networkInterfaces` |
| Type | [`compute.NetworkInterface`](gcp_compute.md#NetworkInterface) |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of network configurations for this instance. These specify how
interfaces are configured to interact with other network services, such as
connecting to the internet. Multiple interfaces are supported per instance.

### `scheduling` {#Instance.scheduling}

| Property | Comments |
|----------|----------|
| Field Name | `scheduling` |
| Type | [`compute.Scheduling`](gcp_compute.md#Scheduling) |

Sets the scheduling options for this instance.

### `selfLink` {#Instance.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `serviceAccounts` {#Instance.serviceAccounts}

| Property | Comments |
|----------|----------|
| Field Name | `serviceAccounts` |
| Type | [`compute.ServiceAccount`](gcp_compute.md#ServiceAccount) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of service accounts, with their specified scopes, authorized for this
instance. Only one service account per VM instance is supported.

Service accounts generate access tokens that can be accessed through the
metadata server and used to authenticate applications on the instance. See
Service Accounts for more information.

### `startRestricted` {#Instance.startRestricted}

| Property | Comments |
|----------|----------|
| Field Name | `startRestricted` |
| Type | `bool` |

[Output Only] Whether a VM has been restricted for start because Compute
Engine has detected suspicious activity.

### `status` {#Instance.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the instance. One of the following values:
PROVISIONING, STAGING, RUNNING, STOPPING, STOPPED, SUSPENDING, SUSPENDED,
and TERMINATED.
Valid values:
    PROVISIONING
    RUNNING
    STAGING
    STOPPED
    STOPPING
    SUSPENDED
    SUSPENDING
    TERMINATED

### `statusMessage` {#Instance.statusMessage}

| Property | Comments |
|----------|----------|
| Field Name | `statusMessage` |
| Type | `string` |

[Output Only] An optional, human-readable explanation of the status.

### `tags` {#Instance.tags}

| Property | Comments |
|----------|----------|
| Field Name | `tags` |
| Type | [`compute.Tags`](gcp_compute.md#Tags) |

A list of tags to apply to this instance. Tags are used to identify valid
sources or targets for network firewalls and are specified by the client
during instance creation. The tags can be later modified by the setTags
method. Each tag within the list must comply with RFC1035.

### `zone` {#Instance.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] URL of the zone where the instance resides. You must specify
this field as part of the HTTP request URL. It is not settable as a field in
the request body.

## Message `InstanceAggregatedList` {#InstanceAggregatedList}



### Inputs for `InstanceAggregatedList`

* `string` [`id`](#InstanceAggregatedList.id) = 1
* `repeated` [`compute.InstanceAggregatedList.ItemsEntry`](gcp_compute.md#InstanceAggregatedList.ItemsEntry) [`items`](#InstanceAggregatedList.items) = 2
* `string` [`kind`](#InstanceAggregatedList.kind) = 3
* `string` [`nextPageToken`](#InstanceAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceAggregatedList.selfLink) = 5
* [`compute.InstanceAggregatedList.Warning`](gcp_compute.md#InstanceAggregatedList.Warning) [`warning`](#InstanceAggregatedList.warning) = 6

### `id` {#InstanceAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceAggregatedList.ItemsEntry`](gcp_compute.md#InstanceAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstancesScopedList resources.

### `kind` {#InstanceAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#instanceAggregatedList for
aggregated lists of Instance resources.

### `nextPageToken` {#InstanceAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceAggregatedList.Warning`](gcp_compute.md#InstanceAggregatedList.Warning) |

## Message `InstanceGroup` {#InstanceGroup}

InstanceGroups (== resource_for beta.instanceGroups ==) (== resource_for
v1.instanceGroups ==) (== resource_for beta.regionInstanceGroups ==) (==
resource_for v1.regionInstanceGroups ==)


### Inputs for `InstanceGroup`

* `string` [`creationTimestamp`](#InstanceGroup.creationTimestamp) = 1
* `string` [`description`](#InstanceGroup.description) = 2
* `string` [`fingerprint`](#InstanceGroup.fingerprint) = 3
* `string` [`id`](#InstanceGroup.id) = 4
* `string` [`kind`](#InstanceGroup.kind) = 5
* `string` [`name`](#InstanceGroup.name) = 6 (**Required**)
* `repeated` [`compute.NamedPort`](gcp_compute.md#NamedPort) [`namedPorts`](#InstanceGroup.namedPorts) = 7
* `string` [`network`](#InstanceGroup.network) = 8
* `string` [`region`](#InstanceGroup.region) = 9
* `string` [`selfLink`](#InstanceGroup.selfLink) = 10
* `int32` [`size`](#InstanceGroup.size) = 11
* `string` [`subnetwork`](#InstanceGroup.subnetwork) = 12
* `string` [`zone`](#InstanceGroup.zone) = 13

### `creationTimestamp` {#InstanceGroup.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] The creation timestamp for this instance group in RFC3339 text
format.

### `description` {#InstanceGroup.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `fingerprint` {#InstanceGroup.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

[Output Only] The fingerprint of the named ports. The system uses this
fingerprint to detect conflicts when multiple users change the named ports
concurrently.

### `id` {#InstanceGroup.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] A unique identifier for this instance group, generated by the
server.

### `kind` {#InstanceGroup.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always compute#instanceGroup for
instance groups.

### `name` {#InstanceGroup.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name of the instance group. The name must be 1-63 characters long, and
comply with RFC1035.

### `namedPorts` {#InstanceGroup.namedPorts}

| Property | Comments |
|----------|----------|
| Field Name | `namedPorts` |
| Type | [`compute.NamedPort`](gcp_compute.md#NamedPort) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Assigns a name to a port number. For example: {name: "http", port: 80}

This allows the system to reference ports by the assigned name instead of a
port number. Named ports can also contain multiple ports. For example:
[{name: "http", port: 80},{name: "http", port: 8080}]

Named ports apply to all instances in this instance group.

### `network` {#InstanceGroup.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

The URL of the network to which all instances in the instance group belong.

### `region` {#InstanceGroup.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] The URL of the region where the instance group is located (for
regional resources).

### `selfLink` {#InstanceGroup.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] The URL for this instance group. The server generates this
URL.

### `size` {#InstanceGroup.size}

| Property | Comments |
|----------|----------|
| Field Name | `size` |
| Type | `int32` |

[Output Only] The total number of instances in the instance group.

### `subnetwork` {#InstanceGroup.subnetwork}

| Property | Comments |
|----------|----------|
| Field Name | `subnetwork` |
| Type | `string` |

[Output Only] The URL of the subnetwork to which all instances in the
instance group belong.

### `zone` {#InstanceGroup.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] The URL of the zone where the instance group is located (for
zonal resources).

## Message `InstanceGroupAggregatedList` {#InstanceGroupAggregatedList}



### Inputs for `InstanceGroupAggregatedList`

* `string` [`id`](#InstanceGroupAggregatedList.id) = 1
* `repeated` [`compute.InstanceGroupAggregatedList.ItemsEntry`](gcp_compute.md#InstanceGroupAggregatedList.ItemsEntry) [`items`](#InstanceGroupAggregatedList.items) = 2
* `string` [`kind`](#InstanceGroupAggregatedList.kind) = 3
* `string` [`nextPageToken`](#InstanceGroupAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceGroupAggregatedList.selfLink) = 5
* [`compute.InstanceGroupAggregatedList.Warning`](gcp_compute.md#InstanceGroupAggregatedList.Warning) [`warning`](#InstanceGroupAggregatedList.warning) = 6

### `id` {#InstanceGroupAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceGroupAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceGroupAggregatedList.ItemsEntry`](gcp_compute.md#InstanceGroupAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceGroupsScopedList resources.

### `kind` {#InstanceGroupAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceGroupAggregatedList for aggregated lists of instance groups.

### `nextPageToken` {#InstanceGroupAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceGroupAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceGroupAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupAggregatedList.Warning`](gcp_compute.md#InstanceGroupAggregatedList.Warning) |

## Message `InstanceGroupList` {#InstanceGroupList}

A list of InstanceGroup resources.


### Inputs for `InstanceGroupList`

* `string` [`id`](#InstanceGroupList.id) = 1
* `repeated` [`compute.InstanceGroup`](gcp_compute.md#InstanceGroup) [`items`](#InstanceGroupList.items) = 2
* `string` [`kind`](#InstanceGroupList.kind) = 3
* `string` [`nextPageToken`](#InstanceGroupList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceGroupList.selfLink) = 5
* [`compute.InstanceGroupList.Warning`](gcp_compute.md#InstanceGroupList.Warning) [`warning`](#InstanceGroupList.warning) = 6

### `id` {#InstanceGroupList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceGroupList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceGroup`](gcp_compute.md#InstanceGroup) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceGroup resources.

### `kind` {#InstanceGroupList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always compute#instanceGroupList
for instance group lists.

### `nextPageToken` {#InstanceGroupList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceGroupList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceGroupList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupList.Warning`](gcp_compute.md#InstanceGroupList.Warning) |

## Message `InstanceGroupManager` {#InstanceGroupManager}

An Instance Group Manager resource. (== resource_for
beta.instanceGroupManagers ==) (== resource_for v1.instanceGroupManagers ==)
(== resource_for beta.regionInstanceGroupManagers ==) (== resource_for
v1.regionInstanceGroupManagers ==)


### Inputs for `InstanceGroupManager`

* `string` [`baseInstanceName`](#InstanceGroupManager.baseInstanceName) = 1
* `string` [`creationTimestamp`](#InstanceGroupManager.creationTimestamp) = 2
* [`compute.InstanceGroupManagerActionsSummary`](gcp_compute.md#InstanceGroupManagerActionsSummary) [`currentActions`](#InstanceGroupManager.currentActions) = 3
* `string` [`description`](#InstanceGroupManager.description) = 4
* `string` [`fingerprint`](#InstanceGroupManager.fingerprint) = 5
* `string` [`id`](#InstanceGroupManager.id) = 6
* `string` [`instanceGroup`](#InstanceGroupManager.instanceGroup) = 7
* `string` [`instanceTemplate`](#InstanceGroupManager.instanceTemplate) = 8
* `string` [`kind`](#InstanceGroupManager.kind) = 9
* `string` [`name`](#InstanceGroupManager.name) = 10 (**Required**)
* `repeated` [`compute.NamedPort`](gcp_compute.md#NamedPort) [`namedPorts`](#InstanceGroupManager.namedPorts) = 11
* `string` [`region`](#InstanceGroupManager.region) = 12
* `string` [`selfLink`](#InstanceGroupManager.selfLink) = 13
* `repeated` `string` [`targetPools`](#InstanceGroupManager.targetPools) = 14
* `int32` [`targetSize`](#InstanceGroupManager.targetSize) = 15
* `string` [`zone`](#InstanceGroupManager.zone) = 16

### `baseInstanceName` {#InstanceGroupManager.baseInstanceName}

| Property | Comments |
|----------|----------|
| Field Name | `baseInstanceName` |
| Type | `string` |

The base instance name to use for instances in this group. The value must be
1-58 characters long. Instances are named by appending a hyphen and a random
four-character string to the base instance name. The base instance name must
comply with RFC1035.

### `creationTimestamp` {#InstanceGroupManager.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] The creation timestamp for this managed instance group in
RFC3339 text format.

### `currentActions` {#InstanceGroupManager.currentActions}

| Property | Comments |
|----------|----------|
| Field Name | `currentActions` |
| Type | [`compute.InstanceGroupManagerActionsSummary`](gcp_compute.md#InstanceGroupManagerActionsSummary) |

[Output Only] The list of instance actions and the number of instances in
this managed instance group that are scheduled for each of those actions.

### `description` {#InstanceGroupManager.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `fingerprint` {#InstanceGroupManager.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint of this resource. This field may be used in optimistic locking.
It will be ignored when inserting an InstanceGroupManager. An up-to-date
fingerprint must be provided in order to update the InstanceGroupManager.

### `id` {#InstanceGroupManager.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] A unique identifier for this resource type. The server
generates this identifier.

### `instanceGroup` {#InstanceGroupManager.instanceGroup}

| Property | Comments |
|----------|----------|
| Field Name | `instanceGroup` |
| Type | `string` |

[Output Only] The URL of the Instance Group resource.

### `instanceTemplate` {#InstanceGroupManager.instanceTemplate}

| Property | Comments |
|----------|----------|
| Field Name | `instanceTemplate` |
| Type | `string` |

The URL of the instance template that is specified for this managed instance
group. The group uses this template to create all new instances in the
managed instance group.

### `kind` {#InstanceGroupManager.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceGroupManager for managed instance groups.

### `name` {#InstanceGroupManager.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name of the managed instance group. The name must be 1-63 characters
long, and comply with RFC1035.

### `namedPorts` {#InstanceGroupManager.namedPorts}

| Property | Comments |
|----------|----------|
| Field Name | `namedPorts` |
| Type | [`compute.NamedPort`](gcp_compute.md#NamedPort) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Named ports configured for the Instance Groups complementary to this
Instance Group Manager.

### `region` {#InstanceGroupManager.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] The URL of the region where the managed instance group resides
(for regional resources).

### `selfLink` {#InstanceGroupManager.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] The URL for this managed instance group. The server defines
this URL.

### `targetPools` {#InstanceGroupManager.targetPools}

| Property | Comments |
|----------|----------|
| Field Name | `targetPools` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs for all TargetPool resources to which instances in the
instanceGroup field are added. The target pools automatically apply to all
of the instances in the managed instance group.

### `targetSize` {#InstanceGroupManager.targetSize}

| Property | Comments |
|----------|----------|
| Field Name | `targetSize` |
| Type | `int32` |

The target number of running instances for this managed instance group.
Deleting or abandoning instances reduces this number. Resizing the group
changes this number.

### `zone` {#InstanceGroupManager.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] The URL of the zone where the managed instance group is
located (for zonal resources).

## Message `InstanceGroupManagerActionsSummary` {#InstanceGroupManagerActionsSummary}



### Inputs for `InstanceGroupManagerActionsSummary`

* `int32` [`abandoning`](#InstanceGroupManagerActionsSummary.abandoning) = 1
* `int32` [`creating`](#InstanceGroupManagerActionsSummary.creating) = 2
* `int32` [`creatingWithoutRetries`](#InstanceGroupManagerActionsSummary.creatingWithoutRetries) = 3
* `int32` [`deleting`](#InstanceGroupManagerActionsSummary.deleting) = 4
* `int32` [`none`](#InstanceGroupManagerActionsSummary.none) = 5
* `int32` [`recreating`](#InstanceGroupManagerActionsSummary.recreating) = 6
* `int32` [`refreshing`](#InstanceGroupManagerActionsSummary.refreshing) = 7
* `int32` [`restarting`](#InstanceGroupManagerActionsSummary.restarting) = 8

### `abandoning` {#InstanceGroupManagerActionsSummary.abandoning}

| Property | Comments |
|----------|----------|
| Field Name | `abandoning` |
| Type | `int32` |

[Output Only] The total number of instances in the managed instance group
that are scheduled to be abandoned. Abandoning an instance removes it from
the managed instance group without deleting it.

### `creating` {#InstanceGroupManagerActionsSummary.creating}

| Property | Comments |
|----------|----------|
| Field Name | `creating` |
| Type | `int32` |

[Output Only] The number of instances in the managed instance group that are
scheduled to be created or are currently being created. If the group fails
to create any of these instances, it tries again until it creates the
instance successfully.

If you have disabled creation retries, this field will not be populated;
instead, the creatingWithoutRetries field will be populated.

### `creatingWithoutRetries` {#InstanceGroupManagerActionsSummary.creatingWithoutRetries}

| Property | Comments |
|----------|----------|
| Field Name | `creatingWithoutRetries` |
| Type | `int32` |

[Output Only] The number of instances that the managed instance group will
attempt to create. The group attempts to create each instance only once. If
the group fails to create any of these instances, it decreases the group's
targetSize value accordingly.

### `deleting` {#InstanceGroupManagerActionsSummary.deleting}

| Property | Comments |
|----------|----------|
| Field Name | `deleting` |
| Type | `int32` |

[Output Only] The number of instances in the managed instance group that are
scheduled to be deleted or are currently being deleted.

### `none` {#InstanceGroupManagerActionsSummary.none}

| Property | Comments |
|----------|----------|
| Field Name | `none` |
| Type | `int32` |

[Output Only] The number of instances in the managed instance group that are
running and have no scheduled actions.

### `recreating` {#InstanceGroupManagerActionsSummary.recreating}

| Property | Comments |
|----------|----------|
| Field Name | `recreating` |
| Type | `int32` |

[Output Only] The number of instances in the managed instance group that are
scheduled to be recreated or are currently being being recreated. Recreating
an instance deletes the existing root persistent disk and creates a new disk
from the image that is defined in the instance template.

### `refreshing` {#InstanceGroupManagerActionsSummary.refreshing}

| Property | Comments |
|----------|----------|
| Field Name | `refreshing` |
| Type | `int32` |

[Output Only] The number of instances in the managed instance group that are
being reconfigured with properties that do not require a restart or a
recreate action. For example, setting or removing target pools for the
instance.

### `restarting` {#InstanceGroupManagerActionsSummary.restarting}

| Property | Comments |
|----------|----------|
| Field Name | `restarting` |
| Type | `int32` |

[Output Only] The number of instances in the managed instance group that are
scheduled to be restarted or are currently being restarted.

## Message `InstanceGroupManagerAggregatedList` {#InstanceGroupManagerAggregatedList}



### Inputs for `InstanceGroupManagerAggregatedList`

* `string` [`id`](#InstanceGroupManagerAggregatedList.id) = 1
* `repeated` [`compute.InstanceGroupManagerAggregatedList.ItemsEntry`](gcp_compute.md#InstanceGroupManagerAggregatedList.ItemsEntry) [`items`](#InstanceGroupManagerAggregatedList.items) = 2
* `string` [`kind`](#InstanceGroupManagerAggregatedList.kind) = 3
* `string` [`nextPageToken`](#InstanceGroupManagerAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceGroupManagerAggregatedList.selfLink) = 5
* [`compute.InstanceGroupManagerAggregatedList.Warning`](gcp_compute.md#InstanceGroupManagerAggregatedList.Warning) [`warning`](#InstanceGroupManagerAggregatedList.warning) = 6

### `id` {#InstanceGroupManagerAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceGroupManagerAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceGroupManagerAggregatedList.ItemsEntry`](gcp_compute.md#InstanceGroupManagerAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceGroupManagersScopedList resources.

### `kind` {#InstanceGroupManagerAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceGroupManagerAggregatedList for an aggregated list of managed
instance groups.

### `nextPageToken` {#InstanceGroupManagerAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceGroupManagerAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceGroupManagerAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupManagerAggregatedList.Warning`](gcp_compute.md#InstanceGroupManagerAggregatedList.Warning) |

## Message `InstanceGroupManagerList` {#InstanceGroupManagerList}

[Output Only] A list of managed instance groups.


### Inputs for `InstanceGroupManagerList`

* `string` [`id`](#InstanceGroupManagerList.id) = 1
* `repeated` [`compute.InstanceGroupManager`](gcp_compute.md#InstanceGroupManager) [`items`](#InstanceGroupManagerList.items) = 2
* `string` [`kind`](#InstanceGroupManagerList.kind) = 3
* `string` [`nextPageToken`](#InstanceGroupManagerList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceGroupManagerList.selfLink) = 5
* [`compute.InstanceGroupManagerList.Warning`](gcp_compute.md#InstanceGroupManagerList.Warning) [`warning`](#InstanceGroupManagerList.warning) = 6

### `id` {#InstanceGroupManagerList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceGroupManagerList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceGroupManager`](gcp_compute.md#InstanceGroupManager) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceGroupManager resources.

### `kind` {#InstanceGroupManagerList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceGroupManagerList for a list of managed instance groups.

### `nextPageToken` {#InstanceGroupManagerList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceGroupManagerList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceGroupManagerList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupManagerList.Warning`](gcp_compute.md#InstanceGroupManagerList.Warning) |

## Message `InstanceGroupManagersAbandonInstancesRequest` {#InstanceGroupManagersAbandonInstancesRequest}



### Inputs for `InstanceGroupManagersAbandonInstancesRequest`

* `repeated` `string` [`instances`](#InstanceGroupManagersAbandonInstancesRequest.instances) = 1

### `instances` {#InstanceGroupManagersAbandonInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs of one or more instances to abandon. This can be a full URL or a
partial URL, such as zones/[ZONE]/instances/[INSTANCE_NAME].

## Message `InstanceGroupManagersDeleteInstancesRequest` {#InstanceGroupManagersDeleteInstancesRequest}



### Inputs for `InstanceGroupManagersDeleteInstancesRequest`

* `repeated` `string` [`instances`](#InstanceGroupManagersDeleteInstancesRequest.instances) = 1

### `instances` {#InstanceGroupManagersDeleteInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs of one or more instances to delete. This can be a full URL or a
partial URL, such as zones/[ZONE]/instances/[INSTANCE_NAME].

## Message `InstanceGroupManagersListManagedInstancesResponse` {#InstanceGroupManagersListManagedInstancesResponse}



### Inputs for `InstanceGroupManagersListManagedInstancesResponse`

* `repeated` [`compute.ManagedInstance`](gcp_compute.md#ManagedInstance) [`managedInstances`](#InstanceGroupManagersListManagedInstancesResponse.managedInstances) = 1

### `managedInstances` {#InstanceGroupManagersListManagedInstancesResponse.managedInstances}

| Property | Comments |
|----------|----------|
| Field Name | `managedInstances` |
| Type | [`compute.ManagedInstance`](gcp_compute.md#ManagedInstance) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] The list of instances in the managed instance group.

## Message `InstanceGroupManagersRecreateInstancesRequest` {#InstanceGroupManagersRecreateInstancesRequest}



### Inputs for `InstanceGroupManagersRecreateInstancesRequest`

* `repeated` `string` [`instances`](#InstanceGroupManagersRecreateInstancesRequest.instances) = 1

### `instances` {#InstanceGroupManagersRecreateInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs of one or more instances to recreate. This can be a full URL or a
partial URL, such as zones/[ZONE]/instances/[INSTANCE_NAME].

## Message `InstanceGroupManagersScopedList` {#InstanceGroupManagersScopedList}



### Inputs for `InstanceGroupManagersScopedList`

* `repeated` [`compute.InstanceGroupManager`](gcp_compute.md#InstanceGroupManager) [`instanceGroupManagers`](#InstanceGroupManagersScopedList.instanceGroupManagers) = 1
* [`compute.InstanceGroupManagersScopedList.Warning`](gcp_compute.md#InstanceGroupManagersScopedList.Warning) [`warning`](#InstanceGroupManagersScopedList.warning) = 2

### `instanceGroupManagers` {#InstanceGroupManagersScopedList.instanceGroupManagers}

| Property | Comments |
|----------|----------|
| Field Name | `instanceGroupManagers` |
| Type | [`compute.InstanceGroupManager`](gcp_compute.md#InstanceGroupManager) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] The list of managed instance groups that are contained in the
specified project and zone.

### `warning` {#InstanceGroupManagersScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupManagersScopedList.Warning`](gcp_compute.md#InstanceGroupManagersScopedList.Warning) |

## Message `InstanceGroupManagersSetInstanceTemplateRequest` {#InstanceGroupManagersSetInstanceTemplateRequest}



### Inputs for `InstanceGroupManagersSetInstanceTemplateRequest`

* `string` [`instanceTemplate`](#InstanceGroupManagersSetInstanceTemplateRequest.instanceTemplate) = 1

### `instanceTemplate` {#InstanceGroupManagersSetInstanceTemplateRequest.instanceTemplate}

| Property | Comments |
|----------|----------|
| Field Name | `instanceTemplate` |
| Type | `string` |

The URL of the instance template that is specified for this managed instance
group. The group uses this template to create all new instances in the
managed instance group.

## Message `InstanceGroupManagersSetTargetPoolsRequest` {#InstanceGroupManagersSetTargetPoolsRequest}



### Inputs for `InstanceGroupManagersSetTargetPoolsRequest`

* `string` [`fingerprint`](#InstanceGroupManagersSetTargetPoolsRequest.fingerprint) = 1
* `repeated` `string` [`targetPools`](#InstanceGroupManagersSetTargetPoolsRequest.targetPools) = 2

### `fingerprint` {#InstanceGroupManagersSetTargetPoolsRequest.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

The fingerprint of the target pools information. Use this optional property
to prevent conflicts when multiple users change the target pools settings
concurrently. Obtain the fingerprint with the instanceGroupManagers.get
method. Then, include the fingerprint in your request to ensure that you do
not overwrite changes that were applied from another concurrent request.

### `targetPools` {#InstanceGroupManagersSetTargetPoolsRequest.targetPools}

| Property | Comments |
|----------|----------|
| Field Name | `targetPools` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of target pool URLs that instances in this managed instance group
belong to. The managed instance group applies these target pools to all of
the instances in the group. Existing instances and new instances in the
group all receive these target pool settings.

## Message `InstanceGroupsAddInstancesRequest` {#InstanceGroupsAddInstancesRequest}



### Inputs for `InstanceGroupsAddInstancesRequest`

* `repeated` [`compute.InstanceReference`](gcp_compute.md#InstanceReference) [`instances`](#InstanceGroupsAddInstancesRequest.instances) = 1

### `instances` {#InstanceGroupsAddInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | [`compute.InstanceReference`](gcp_compute.md#InstanceReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of instances to add to the instance group.

## Message `InstanceGroupsListInstances` {#InstanceGroupsListInstances}



### Inputs for `InstanceGroupsListInstances`

* `string` [`id`](#InstanceGroupsListInstances.id) = 1
* `repeated` [`compute.InstanceWithNamedPorts`](gcp_compute.md#InstanceWithNamedPorts) [`items`](#InstanceGroupsListInstances.items) = 2
* `string` [`kind`](#InstanceGroupsListInstances.kind) = 3
* `string` [`nextPageToken`](#InstanceGroupsListInstances.nextPageToken) = 4
* `string` [`selfLink`](#InstanceGroupsListInstances.selfLink) = 5
* [`compute.InstanceGroupsListInstances.Warning`](gcp_compute.md#InstanceGroupsListInstances.Warning) [`warning`](#InstanceGroupsListInstances.warning) = 6

### `id` {#InstanceGroupsListInstances.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceGroupsListInstances.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceWithNamedPorts`](gcp_compute.md#InstanceWithNamedPorts) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceWithNamedPorts resources.

### `kind` {#InstanceGroupsListInstances.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceGroupsListInstances for the list of instances in the
specified instance group.

### `nextPageToken` {#InstanceGroupsListInstances.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceGroupsListInstances.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceGroupsListInstances.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupsListInstances.Warning`](gcp_compute.md#InstanceGroupsListInstances.Warning) |

## Message `InstanceGroupsListInstancesRequest` {#InstanceGroupsListInstancesRequest}



### Inputs for `InstanceGroupsListInstancesRequest`

* `string` [`instanceState`](#InstanceGroupsListInstancesRequest.instanceState) = 1

### `instanceState` {#InstanceGroupsListInstancesRequest.instanceState}

| Property | Comments |
|----------|----------|
| Field Name | `instanceState` |
| Type | `string` |

A filter for the state of the instances in the instance group. Valid options
are ALL or RUNNING. If you do not specify this parameter the list includes
all instances regardless of their state.
Valid values:
    ALL
    RUNNING

## Message `InstanceGroupsRemoveInstancesRequest` {#InstanceGroupsRemoveInstancesRequest}



### Inputs for `InstanceGroupsRemoveInstancesRequest`

* `repeated` [`compute.InstanceReference`](gcp_compute.md#InstanceReference) [`instances`](#InstanceGroupsRemoveInstancesRequest.instances) = 1

### `instances` {#InstanceGroupsRemoveInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | [`compute.InstanceReference`](gcp_compute.md#InstanceReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of instances to remove from the instance group.

## Message `InstanceGroupsScopedList` {#InstanceGroupsScopedList}



### Inputs for `InstanceGroupsScopedList`

* `repeated` [`compute.InstanceGroup`](gcp_compute.md#InstanceGroup) [`instanceGroups`](#InstanceGroupsScopedList.instanceGroups) = 1
* [`compute.InstanceGroupsScopedList.Warning`](gcp_compute.md#InstanceGroupsScopedList.Warning) [`warning`](#InstanceGroupsScopedList.warning) = 2

### `instanceGroups` {#InstanceGroupsScopedList.instanceGroups}

| Property | Comments |
|----------|----------|
| Field Name | `instanceGroups` |
| Type | [`compute.InstanceGroup`](gcp_compute.md#InstanceGroup) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] The list of instance groups that are contained in this scope.

### `warning` {#InstanceGroupsScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceGroupsScopedList.Warning`](gcp_compute.md#InstanceGroupsScopedList.Warning) |

## Message `InstanceGroupsSetNamedPortsRequest` {#InstanceGroupsSetNamedPortsRequest}



### Inputs for `InstanceGroupsSetNamedPortsRequest`

* `string` [`fingerprint`](#InstanceGroupsSetNamedPortsRequest.fingerprint) = 1
* `repeated` [`compute.NamedPort`](gcp_compute.md#NamedPort) [`namedPorts`](#InstanceGroupsSetNamedPortsRequest.namedPorts) = 2

### `fingerprint` {#InstanceGroupsSetNamedPortsRequest.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

The fingerprint of the named ports information for this instance group. Use
this optional property to prevent conflicts when multiple users change the
named ports settings concurrently. Obtain the fingerprint with the
instanceGroups.get method. Then, include the fingerprint in your request to
ensure that you do not overwrite changes that were applied from another
concurrent request.

### `namedPorts` {#InstanceGroupsSetNamedPortsRequest.namedPorts}

| Property | Comments |
|----------|----------|
| Field Name | `namedPorts` |
| Type | [`compute.NamedPort`](gcp_compute.md#NamedPort) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of named ports to set for this instance group.

## Message `InstanceList` {#InstanceList}

Contains a list of instances.


### Inputs for `InstanceList`

* `string` [`id`](#InstanceList.id) = 1
* `repeated` [`compute.Instance`](gcp_compute.md#Instance) [`items`](#InstanceList.items) = 2
* `string` [`kind`](#InstanceList.kind) = 3
* `string` [`nextPageToken`](#InstanceList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceList.selfLink) = 5
* [`compute.InstanceList.Warning`](gcp_compute.md#InstanceList.Warning) [`warning`](#InstanceList.warning) = 6

### `id` {#InstanceList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Instance`](gcp_compute.md#Instance) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Instance resources.

### `kind` {#InstanceList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#instanceList for lists of
Instance resources.

### `nextPageToken` {#InstanceList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceList.Warning`](gcp_compute.md#InstanceList.Warning) |

## Message `InstanceListReferrers` {#InstanceListReferrers}

Contains a list of instance referrers.


### Inputs for `InstanceListReferrers`

* `string` [`id`](#InstanceListReferrers.id) = 1
* `repeated` [`compute.Reference`](gcp_compute.md#Reference) [`items`](#InstanceListReferrers.items) = 2
* `string` [`kind`](#InstanceListReferrers.kind) = 3
* `string` [`nextPageToken`](#InstanceListReferrers.nextPageToken) = 4
* `string` [`selfLink`](#InstanceListReferrers.selfLink) = 5
* [`compute.InstanceListReferrers.Warning`](gcp_compute.md#InstanceListReferrers.Warning) [`warning`](#InstanceListReferrers.warning) = 6

### `id` {#InstanceListReferrers.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceListReferrers.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Reference`](gcp_compute.md#Reference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Reference resources.

### `kind` {#InstanceListReferrers.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#instanceListReferrers for
lists of Instance referrers.

### `nextPageToken` {#InstanceListReferrers.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceListReferrers.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceListReferrers.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceListReferrers.Warning`](gcp_compute.md#InstanceListReferrers.Warning) |

## Message `InstanceMoveRequest` {#InstanceMoveRequest}



### Inputs for `InstanceMoveRequest`

* `string` [`destinationZone`](#InstanceMoveRequest.destinationZone) = 1
* `string` [`targetInstance`](#InstanceMoveRequest.targetInstance) = 2

### `destinationZone` {#InstanceMoveRequest.destinationZone}

| Property | Comments |
|----------|----------|
| Field Name | `destinationZone` |
| Type | `string` |

The URL of the destination zone to move the instance. This can be a full or
partial URL. For example, the following are all valid URLs to a zone:
- https://www.googleapis.com/compute/v1/projects/project/zones/zone
- projects/project/zones/zone
- zones/zone

### `targetInstance` {#InstanceMoveRequest.targetInstance}

| Property | Comments |
|----------|----------|
| Field Name | `targetInstance` |
| Type | `string` |

The URL of the target instance to move. This can be a full or partial URL.
For example, the following are all valid URLs to an instance:
-
https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
- projects/project/zones/zone/instances/instance
- zones/zone/instances/instance

## Message `InstanceProperties` {#InstanceProperties}



### Inputs for `InstanceProperties`

* `bool` [`canIpForward`](#InstanceProperties.canIpForward) = 1
* `string` [`description`](#InstanceProperties.description) = 2
* `repeated` [`compute.AttachedDisk`](gcp_compute.md#AttachedDisk) [`disks`](#InstanceProperties.disks) = 3
* `repeated` [`compute.AcceleratorConfig`](gcp_compute.md#AcceleratorConfig) [`guestAccelerators`](#InstanceProperties.guestAccelerators) = 4
* `repeated` [`compute.InstanceProperties.LabelsEntry`](gcp_compute.md#InstanceProperties.LabelsEntry) [`labels`](#InstanceProperties.labels) = 5
* `string` [`machineType`](#InstanceProperties.machineType) = 6
* [`compute.Metadata`](gcp_compute.md#Metadata) [`metadata`](#InstanceProperties.metadata) = 7
* `string` [`minCpuPlatform`](#InstanceProperties.minCpuPlatform) = 8
* `repeated` [`compute.NetworkInterface`](gcp_compute.md#NetworkInterface) [`networkInterfaces`](#InstanceProperties.networkInterfaces) = 9
* [`compute.Scheduling`](gcp_compute.md#Scheduling) [`scheduling`](#InstanceProperties.scheduling) = 10
* `repeated` [`compute.ServiceAccount`](gcp_compute.md#ServiceAccount) [`serviceAccounts`](#InstanceProperties.serviceAccounts) = 11
* [`compute.Tags`](gcp_compute.md#Tags) [`tags`](#InstanceProperties.tags) = 12

### `canIpForward` {#InstanceProperties.canIpForward}

| Property | Comments |
|----------|----------|
| Field Name | `canIpForward` |
| Type | `bool` |

Enables instances created based on this template to send packets with source
IP addresses other than their own and receive packets with destination IP
addresses other than their own. If these instances will be used as an IP
gateway or it will be set as the next-hop in a Route resource, specify true.
If unsure, leave this set to false. See the Enable IP forwarding
documentation for more information.

### `description` {#InstanceProperties.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional text description for the instances that are created from this
instance template.

### `disks` {#InstanceProperties.disks}

| Property | Comments |
|----------|----------|
| Field Name | `disks` |
| Type | [`compute.AttachedDisk`](gcp_compute.md#AttachedDisk) |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of disks that are associated with the instances that are created
from this template.

### `guestAccelerators` {#InstanceProperties.guestAccelerators}

| Property | Comments |
|----------|----------|
| Field Name | `guestAccelerators` |
| Type | [`compute.AcceleratorConfig`](gcp_compute.md#AcceleratorConfig) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of guest accelerator cards' type and count to use for instances
created from the instance template.

### `labels` {#InstanceProperties.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.InstanceProperties.LabelsEntry`](gcp_compute.md#InstanceProperties.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Labels to apply to instances that are created from this template.

### `machineType` {#InstanceProperties.machineType}

| Property | Comments |
|----------|----------|
| Field Name | `machineType` |
| Type | `string` |

The machine type to use for instances that are created from this template.

### `metadata` {#InstanceProperties.metadata}

| Property | Comments |
|----------|----------|
| Field Name | `metadata` |
| Type | [`compute.Metadata`](gcp_compute.md#Metadata) |

The metadata key/value pairs to assign to instances that are created from
this template. These pairs can consist of custom metadata or predefined
keys. See Project and instance metadata for more information.

### `minCpuPlatform` {#InstanceProperties.minCpuPlatform}

| Property | Comments |
|----------|----------|
| Field Name | `minCpuPlatform` |
| Type | `string` |

Minimum cpu/platform to be used by this instance. The instance may be
scheduled on the specified or newer cpu/platform. Applicable values are the
friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or
minCpuPlatform: "Intel Sandy Bridge". For more information, read Specifying
a Minimum CPU Platform.

### `networkInterfaces` {#InstanceProperties.networkInterfaces}

| Property | Comments |
|----------|----------|
| Field Name | `networkInterfaces` |
| Type | [`compute.NetworkInterface`](gcp_compute.md#NetworkInterface) |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of network access configurations for this interface.

### `scheduling` {#InstanceProperties.scheduling}

| Property | Comments |
|----------|----------|
| Field Name | `scheduling` |
| Type | [`compute.Scheduling`](gcp_compute.md#Scheduling) |

Specifies the scheduling options for the instances that are created from
this template.

### `serviceAccounts` {#InstanceProperties.serviceAccounts}

| Property | Comments |
|----------|----------|
| Field Name | `serviceAccounts` |
| Type | [`compute.ServiceAccount`](gcp_compute.md#ServiceAccount) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of service accounts with specified scopes. Access tokens for these
service accounts are available to the instances that are created from this
template. Use metadata queries to obtain the access tokens for these
instances.

### `tags` {#InstanceProperties.tags}

| Property | Comments |
|----------|----------|
| Field Name | `tags` |
| Type | [`compute.Tags`](gcp_compute.md#Tags) |

A list of tags to apply to the instances that are created from this
template. The tags identify valid sources or targets for network firewalls.
The setTags method can modify this list of tags. Each tag within the list
must comply with RFC1035.

## Message `InstanceReference` {#InstanceReference}



### Inputs for `InstanceReference`

* `string` [`instance`](#InstanceReference.instance) = 1

### `instance` {#InstanceReference.instance}

| Property | Comments |
|----------|----------|
| Field Name | `instance` |
| Type | `string` |

The URL for a specific instance.

## Message `InstanceTemplate` {#InstanceTemplate}

An Instance Template resource. (== resource_for beta.instanceTemplates ==)
(== resource_for v1.instanceTemplates ==)


### Inputs for `InstanceTemplate`

* `string` [`creationTimestamp`](#InstanceTemplate.creationTimestamp) = 1
* `string` [`description`](#InstanceTemplate.description) = 2
* `string` [`id`](#InstanceTemplate.id) = 3
* `string` [`kind`](#InstanceTemplate.kind) = 4
* `string` [`name`](#InstanceTemplate.name) = 5 (**Required**)
* [`compute.InstanceProperties`](gcp_compute.md#InstanceProperties) [`properties`](#InstanceTemplate.properties) = 6
* `string` [`selfLink`](#InstanceTemplate.selfLink) = 7
* `string` [`sourceInstance`](#InstanceTemplate.sourceInstance) = 8
* [`compute.SourceInstanceParams`](gcp_compute.md#SourceInstanceParams) [`sourceInstanceParams`](#InstanceTemplate.sourceInstanceParams) = 9

### `creationTimestamp` {#InstanceTemplate.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] The creation timestamp for this instance template in RFC3339
text format.

### `description` {#InstanceTemplate.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#InstanceTemplate.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] A unique identifier for this instance template. The server
defines this identifier.

### `kind` {#InstanceTemplate.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always compute#instanceTemplate
for instance templates.

### `name` {#InstanceTemplate.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource; provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `properties` {#InstanceTemplate.properties}

| Property | Comments |
|----------|----------|
| Field Name | `properties` |
| Type | [`compute.InstanceProperties`](gcp_compute.md#InstanceProperties) |

The instance properties for this instance template.

### `selfLink` {#InstanceTemplate.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] The URL for this instance template. The server defines this
URL.

### `sourceInstance` {#InstanceTemplate.sourceInstance}

| Property | Comments |
|----------|----------|
| Field Name | `sourceInstance` |
| Type | `string` |

The source instance used to create the template. You can provide this as a
partial or full URL to the resource. For example, the following are valid
values:
-
https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
- projects/project/zones/zone/instances/instance

### `sourceInstanceParams` {#InstanceTemplate.sourceInstanceParams}

| Property | Comments |
|----------|----------|
| Field Name | `sourceInstanceParams` |
| Type | [`compute.SourceInstanceParams`](gcp_compute.md#SourceInstanceParams) |

The source instance params to use to create this instance template.

## Message `InstanceTemplateList` {#InstanceTemplateList}

A list of instance templates.


### Inputs for `InstanceTemplateList`

* `string` [`id`](#InstanceTemplateList.id) = 1
* `repeated` [`compute.InstanceTemplate`](gcp_compute.md#InstanceTemplate) [`items`](#InstanceTemplateList.items) = 2
* `string` [`kind`](#InstanceTemplateList.kind) = 3
* `string` [`nextPageToken`](#InstanceTemplateList.nextPageToken) = 4
* `string` [`selfLink`](#InstanceTemplateList.selfLink) = 5
* [`compute.InstanceTemplateList.Warning`](gcp_compute.md#InstanceTemplateList.Warning) [`warning`](#InstanceTemplateList.warning) = 6

### `id` {#InstanceTemplateList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InstanceTemplateList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceTemplate`](gcp_compute.md#InstanceTemplate) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceTemplate resources.

### `kind` {#InstanceTemplateList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceTemplatesListResponse for instance template lists.

### `nextPageToken` {#InstanceTemplateList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InstanceTemplateList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InstanceTemplateList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstanceTemplateList.Warning`](gcp_compute.md#InstanceTemplateList.Warning) |

## Message `InstanceWithNamedPorts` {#InstanceWithNamedPorts}



### Inputs for `InstanceWithNamedPorts`

* `string` [`instance`](#InstanceWithNamedPorts.instance) = 1
* `repeated` [`compute.NamedPort`](gcp_compute.md#NamedPort) [`namedPorts`](#InstanceWithNamedPorts.namedPorts) = 2
* `string` [`status`](#InstanceWithNamedPorts.status) = 3

### `instance` {#InstanceWithNamedPorts.instance}

| Property | Comments |
|----------|----------|
| Field Name | `instance` |
| Type | `string` |

[Output Only] The URL of the instance.

### `namedPorts` {#InstanceWithNamedPorts.namedPorts}

| Property | Comments |
|----------|----------|
| Field Name | `namedPorts` |
| Type | [`compute.NamedPort`](gcp_compute.md#NamedPort) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] The named ports that belong to this instance group.

### `status` {#InstanceWithNamedPorts.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the instance.
Valid values:
    PROVISIONING
    RUNNING
    STAGING
    STOPPED
    STOPPING
    SUSPENDED
    SUSPENDING
    TERMINATED

## Message `InstancesScopedList` {#InstancesScopedList}



### Inputs for `InstancesScopedList`

* `repeated` [`compute.Instance`](gcp_compute.md#Instance) [`instances`](#InstancesScopedList.instances) = 1
* [`compute.InstancesScopedList.Warning`](gcp_compute.md#InstancesScopedList.Warning) [`warning`](#InstancesScopedList.warning) = 2

### `instances` {#InstancesScopedList.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | [`compute.Instance`](gcp_compute.md#Instance) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of instances contained in this scope.

### `warning` {#InstancesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InstancesScopedList.Warning`](gcp_compute.md#InstancesScopedList.Warning) |

## Message `InstancesSetLabelsRequest` {#InstancesSetLabelsRequest}



### Inputs for `InstancesSetLabelsRequest`

* `string` [`labelFingerprint`](#InstancesSetLabelsRequest.labelFingerprint) = 1
* `repeated` [`compute.InstancesSetLabelsRequest.LabelsEntry`](gcp_compute.md#InstancesSetLabelsRequest.LabelsEntry) [`labels`](#InstancesSetLabelsRequest.labels) = 2

### `labelFingerprint` {#InstancesSetLabelsRequest.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

Fingerprint of the previous set of labels for this resource, used to prevent
conflicts. Provide the latest fingerprint value when making a request to add
or change labels.

### `labels` {#InstancesSetLabelsRequest.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.InstancesSetLabelsRequest.LabelsEntry`](gcp_compute.md#InstancesSetLabelsRequest.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `InstancesSetMachineResourcesRequest` {#InstancesSetMachineResourcesRequest}



### Inputs for `InstancesSetMachineResourcesRequest`

* `repeated` [`compute.AcceleratorConfig`](gcp_compute.md#AcceleratorConfig) [`guestAccelerators`](#InstancesSetMachineResourcesRequest.guestAccelerators) = 1

### `guestAccelerators` {#InstancesSetMachineResourcesRequest.guestAccelerators}

| Property | Comments |
|----------|----------|
| Field Name | `guestAccelerators` |
| Type | [`compute.AcceleratorConfig`](gcp_compute.md#AcceleratorConfig) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of the type and count of accelerator cards attached to the instance.

## Message `InstancesSetMachineTypeRequest` {#InstancesSetMachineTypeRequest}



### Inputs for `InstancesSetMachineTypeRequest`

* `string` [`machineType`](#InstancesSetMachineTypeRequest.machineType) = 1

### `machineType` {#InstancesSetMachineTypeRequest.machineType}

| Property | Comments |
|----------|----------|
| Field Name | `machineType` |
| Type | `string` |

Full or partial URL of the machine type resource. See Machine Types for a
full list of machine types. For example:
zones/us-central1-f/machineTypes/n1-standard-1

## Message `InstancesSetMinCpuPlatformRequest` {#InstancesSetMinCpuPlatformRequest}



### Inputs for `InstancesSetMinCpuPlatformRequest`

* `string` [`minCpuPlatform`](#InstancesSetMinCpuPlatformRequest.minCpuPlatform) = 1

### `minCpuPlatform` {#InstancesSetMinCpuPlatformRequest.minCpuPlatform}

| Property | Comments |
|----------|----------|
| Field Name | `minCpuPlatform` |
| Type | `string` |

Minimum cpu/platform this instance should be started at.

## Message `InstancesSetServiceAccountRequest` {#InstancesSetServiceAccountRequest}



### Inputs for `InstancesSetServiceAccountRequest`

* `string` [`email`](#InstancesSetServiceAccountRequest.email) = 1
* `repeated` `string` [`scopes`](#InstancesSetServiceAccountRequest.scopes) = 2

### `email` {#InstancesSetServiceAccountRequest.email}

| Property | Comments |
|----------|----------|
| Field Name | `email` |
| Type | `string` |

Email address of the service account.

### `scopes` {#InstancesSetServiceAccountRequest.scopes}

| Property | Comments |
|----------|----------|
| Field Name | `scopes` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of scopes to be made available for this service account.

## Message `InstancesStartWithEncryptionKeyRequest` {#InstancesStartWithEncryptionKeyRequest}



### Inputs for `InstancesStartWithEncryptionKeyRequest`

* `repeated` [`compute.CustomerEncryptionKeyProtectedDisk`](gcp_compute.md#CustomerEncryptionKeyProtectedDisk) [`disks`](#InstancesStartWithEncryptionKeyRequest.disks) = 1

### `disks` {#InstancesStartWithEncryptionKeyRequest.disks}

| Property | Comments |
|----------|----------|
| Field Name | `disks` |
| Type | [`compute.CustomerEncryptionKeyProtectedDisk`](gcp_compute.md#CustomerEncryptionKeyProtectedDisk) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Array of disks associated with this instance that are protected with a
customer-supplied encryption key.

In order to start the instance, the disk url and its corresponding key must
be provided.

If the disk is not protected with a customer-supplied encryption key it
should not be specified.

## Message `Interconnect` {#Interconnect}

Represents an Interconnects resource. The Interconnects resource is a
dedicated connection between Google's network and your on-premises network.
For more information, see the  Dedicated overview page. (== resource_for
v1.interconnects ==) (== resource_for beta.interconnects ==)


### Inputs for `Interconnect`

* `bool` [`adminEnabled`](#Interconnect.adminEnabled) = 1
* `repeated` [`compute.InterconnectCircuitInfo`](gcp_compute.md#InterconnectCircuitInfo) [`circuitInfos`](#Interconnect.circuitInfos) = 2
* `string` [`creationTimestamp`](#Interconnect.creationTimestamp) = 3
* `string` [`customerName`](#Interconnect.customerName) = 4
* `string` [`description`](#Interconnect.description) = 5
* `repeated` [`compute.InterconnectOutageNotification`](gcp_compute.md#InterconnectOutageNotification) [`expectedOutages`](#Interconnect.expectedOutages) = 6
* `string` [`googleIpAddress`](#Interconnect.googleIpAddress) = 7
* `string` [`googleReferenceId`](#Interconnect.googleReferenceId) = 8
* `string` [`id`](#Interconnect.id) = 9
* `repeated` `string` [`interconnectAttachments`](#Interconnect.interconnectAttachments) = 10
* `string` [`interconnectType`](#Interconnect.interconnectType) = 11
* `string` [`kind`](#Interconnect.kind) = 12
* `string` [`linkType`](#Interconnect.linkType) = 13
* `string` [`location`](#Interconnect.location) = 14
* `string` [`name`](#Interconnect.name) = 15 (**Required**)
* `string` [`nocContactEmail`](#Interconnect.nocContactEmail) = 16
* `string` [`operationalStatus`](#Interconnect.operationalStatus) = 17
* `string` [`peerIpAddress`](#Interconnect.peerIpAddress) = 18
* `int32` [`provisionedLinkCount`](#Interconnect.provisionedLinkCount) = 19
* `int32` [`requestedLinkCount`](#Interconnect.requestedLinkCount) = 20
* `string` [`selfLink`](#Interconnect.selfLink) = 21
* `string` [`state`](#Interconnect.state) = 22

### `adminEnabled` {#Interconnect.adminEnabled}

| Property | Comments |
|----------|----------|
| Field Name | `adminEnabled` |
| Type | `bool` |

Administrative status of the interconnect. When this is set to true, the
Interconnect is functional and can carry traffic. When set to false, no
packets can be carried over the interconnect and no BGP routes are exchanged
over it. By default, the status is set to true.

### `circuitInfos` {#Interconnect.circuitInfos}

| Property | Comments |
|----------|----------|
| Field Name | `circuitInfos` |
| Type | [`compute.InterconnectCircuitInfo`](gcp_compute.md#InterconnectCircuitInfo) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of CircuitInfo objects, that describe the individual
circuits in this LAG.

### `creationTimestamp` {#Interconnect.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `customerName` {#Interconnect.customerName}

| Property | Comments |
|----------|----------|
| Field Name | `customerName` |
| Type | `string` |

Customer name, to put in the Letter of Authorization as the party authorized
to request a crossconnect.

### `description` {#Interconnect.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `expectedOutages` {#Interconnect.expectedOutages}

| Property | Comments |
|----------|----------|
| Field Name | `expectedOutages` |
| Type | [`compute.InterconnectOutageNotification`](gcp_compute.md#InterconnectOutageNotification) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of outages expected for this Interconnect.

### `googleIpAddress` {#Interconnect.googleIpAddress}

| Property | Comments |
|----------|----------|
| Field Name | `googleIpAddress` |
| Type | `string` |

[Output Only] IP address configured on the Google side of the Interconnect
link. This can be used only for ping tests.

### `googleReferenceId` {#Interconnect.googleReferenceId}

| Property | Comments |
|----------|----------|
| Field Name | `googleReferenceId` |
| Type | `string` |

[Output Only] Google reference ID; to be used when raising support tickets
with Google or otherwise to debug backend connectivity issues.

### `id` {#Interconnect.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `interconnectAttachments` {#Interconnect.interconnectAttachments}

| Property | Comments |
|----------|----------|
| Field Name | `interconnectAttachments` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of the URLs of all InterconnectAttachments configured
to use this Interconnect.

### `interconnectType` {#Interconnect.interconnectType}

| Property | Comments |
|----------|----------|
| Field Name | `interconnectType` |
| Type | `string` |

Type of interconnect. Note that "IT_PRIVATE" has been deprecated in favor of
"DEDICATED"
Valid values:
    DEDICATED
    IT_PRIVATE
    PARTNER

### `kind` {#Interconnect.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#interconnect for
interconnects.

### `linkType` {#Interconnect.linkType}

| Property | Comments |
|----------|----------|
| Field Name | `linkType` |
| Type | `string` |

Type of link requested. This field indicates speed of each of the links in
the bundle, not the entire bundle. Only 10G per link is allowed for a
dedicated interconnect. Options: Ethernet_10G_LR
Valid values:
    LINK_TYPE_ETHERNET_10G_LR

### `location` {#Interconnect.location}

| Property | Comments |
|----------|----------|
| Field Name | `location` |
| Type | `string` |

URL of the InterconnectLocation object that represents where this connection
is to be provisioned.

### `name` {#Interconnect.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `nocContactEmail` {#Interconnect.nocContactEmail}

| Property | Comments |
|----------|----------|
| Field Name | `nocContactEmail` |
| Type | `string` |

Email address to contact the customer NOC for operations and maintenance
notifications regarding this Interconnect. If specified, this will be used
for notifications in addition to all other forms described, such as
Stackdriver logs alerting and Cloud Notifications.

### `operationalStatus` {#Interconnect.operationalStatus}

| Property | Comments |
|----------|----------|
| Field Name | `operationalStatus` |
| Type | `string` |

[Output Only] The current status of whether or not this Interconnect is
functional.
Valid values:
    OS_ACTIVE
    OS_UNPROVISIONED

### `peerIpAddress` {#Interconnect.peerIpAddress}

| Property | Comments |
|----------|----------|
| Field Name | `peerIpAddress` |
| Type | `string` |

[Output Only] IP address configured on the customer side of the Interconnect
link. The customer should configure this IP address during turnup when
prompted by Google NOC. This can be used only for ping tests.

### `provisionedLinkCount` {#Interconnect.provisionedLinkCount}

| Property | Comments |
|----------|----------|
| Field Name | `provisionedLinkCount` |
| Type | `int32` |

[Output Only] Number of links actually provisioned in this interconnect.

### `requestedLinkCount` {#Interconnect.requestedLinkCount}

| Property | Comments |
|----------|----------|
| Field Name | `requestedLinkCount` |
| Type | `int32` |

Target number of physical links in the link bundle, as requested by the
customer.

### `selfLink` {#Interconnect.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `state` {#Interconnect.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

[Output Only] The current state of whether or not this Interconnect is
functional.
Valid values:
    ACTIVE
    UNPROVISIONED

## Message `InterconnectAttachment` {#InterconnectAttachment}

Represents an InterconnectAttachment (VLAN attachment) resource. For more
information, see  Creating VLAN Attachments. (== resource_for
beta.interconnectAttachments ==) (== resource_for v1.interconnectAttachments
==)


### Inputs for `InterconnectAttachment`

* `bool` [`adminEnabled`](#InterconnectAttachment.adminEnabled) = 1
* `string` [`bandwidth`](#InterconnectAttachment.bandwidth) = 2
* `repeated` `string` [`candidateSubnets`](#InterconnectAttachment.candidateSubnets) = 3
* `string` [`cloudRouterIpAddress`](#InterconnectAttachment.cloudRouterIpAddress) = 4
* `string` [`creationTimestamp`](#InterconnectAttachment.creationTimestamp) = 5
* `string` [`customerRouterIpAddress`](#InterconnectAttachment.customerRouterIpAddress) = 6
* `string` [`description`](#InterconnectAttachment.description) = 7
* `string` [`edgeAvailabilityDomain`](#InterconnectAttachment.edgeAvailabilityDomain) = 8
* `string` [`googleReferenceId`](#InterconnectAttachment.googleReferenceId) = 9
* `string` [`id`](#InterconnectAttachment.id) = 10
* `string` [`interconnect`](#InterconnectAttachment.interconnect) = 11
* `string` [`kind`](#InterconnectAttachment.kind) = 12
* `string` [`name`](#InterconnectAttachment.name) = 13 (**Required**)
* `string` [`operationalStatus`](#InterconnectAttachment.operationalStatus) = 14
* `string` [`pairingKey`](#InterconnectAttachment.pairingKey) = 15
* `string` [`partnerAsn`](#InterconnectAttachment.partnerAsn) = 16
* [`compute.InterconnectAttachmentPartnerMetadata`](gcp_compute.md#InterconnectAttachmentPartnerMetadata) [`partnerMetadata`](#InterconnectAttachment.partnerMetadata) = 17
* [`compute.InterconnectAttachmentPrivateInfo`](gcp_compute.md#InterconnectAttachmentPrivateInfo) [`privateInterconnectInfo`](#InterconnectAttachment.privateInterconnectInfo) = 18
* `string` [`region`](#InterconnectAttachment.region) = 19
* `string` [`router`](#InterconnectAttachment.router) = 20
* `string` [`selfLink`](#InterconnectAttachment.selfLink) = 21
* `string` [`state`](#InterconnectAttachment.state) = 22
* `string` [`type`](#InterconnectAttachment.type) = 23
* `int32` [`vlanTag8021q`](#InterconnectAttachment.vlanTag8021q) = 24

### `adminEnabled` {#InterconnectAttachment.adminEnabled}

| Property | Comments |
|----------|----------|
| Field Name | `adminEnabled` |
| Type | `bool` |

Determines whether this Attachment will carry packets. Not present for
PARTNER_PROVIDER.

### `bandwidth` {#InterconnectAttachment.bandwidth}

| Property | Comments |
|----------|----------|
| Field Name | `bandwidth` |
| Type | `string` |

Provisioned bandwidth capacity for the interconnectAttachment. Can be set by
the partner to update the customer's provisioned bandwidth. Output only for
for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
Valid values:
    BPS_100M
    BPS_10G
    BPS_1G
    BPS_200M
    BPS_2G
    BPS_300M
    BPS_400M
    BPS_500M
    BPS_50M
    BPS_5G

### `candidateSubnets` {#InterconnectAttachment.candidateSubnets}

| Property | Comments |
|----------|----------|
| Field Name | `candidateSubnets` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Up to 16 candidate prefixes that can be used to restrict the allocation of
cloudRouterIpAddress and customerRouterIpAddress for this attachment. All
prefixes must be within link-local address space (169.254.0.0/16) and must
be /29 or shorter (/28, /27, etc). Google will attempt to select an unused
/29 from the supplied candidate prefix(es). The request will fail if all
possible /29s are in use on Google?s edge. If not supplied, Google will
randomly select an unused /29 from all of link-local space.

### `cloudRouterIpAddress` {#InterconnectAttachment.cloudRouterIpAddress}

| Property | Comments |
|----------|----------|
| Field Name | `cloudRouterIpAddress` |
| Type | `string` |

[Output Only] IPv4 address + prefix length to be configured on Cloud Router
Interface for this interconnect attachment.

### `creationTimestamp` {#InterconnectAttachment.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `customerRouterIpAddress` {#InterconnectAttachment.customerRouterIpAddress}

| Property | Comments |
|----------|----------|
| Field Name | `customerRouterIpAddress` |
| Type | `string` |

[Output Only] IPv4 address + prefix length to be configured on the customer
router subinterface for this interconnect attachment.

### `description` {#InterconnectAttachment.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource.

### `edgeAvailabilityDomain` {#InterconnectAttachment.edgeAvailabilityDomain}

| Property | Comments |
|----------|----------|
| Field Name | `edgeAvailabilityDomain` |
| Type | `string` |

Desired availability domain for the attachment. Only available for type
PARTNER, at creation time. For improved reliability, customers should
configure a pair of attachments with one per availability domain. The
selected availability domain will be provided to the Partner via the pairing
key so that the provisioned circuit will lie in the specified domain. If not
specified, the value will default to AVAILABILITY_DOMAIN_ANY.
Valid values:
    AVAILABILITY_DOMAIN_1
    AVAILABILITY_DOMAIN_2
    AVAILABILITY_DOMAIN_ANY

### `googleReferenceId` {#InterconnectAttachment.googleReferenceId}

| Property | Comments |
|----------|----------|
| Field Name | `googleReferenceId` |
| Type | `string` |

[Output Only] Google reference ID, to be used when raising support tickets
with Google or otherwise to debug backend connectivity issues.

### `id` {#InterconnectAttachment.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `interconnect` {#InterconnectAttachment.interconnect}

| Property | Comments |
|----------|----------|
| Field Name | `interconnect` |
| Type | `string` |

URL of the underlying Interconnect object that this attachment's traffic
will traverse through.

### `kind` {#InterconnectAttachment.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#interconnectAttachment
for interconnect attachments.

### `name` {#InterconnectAttachment.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `operationalStatus` {#InterconnectAttachment.operationalStatus}

| Property | Comments |
|----------|----------|
| Field Name | `operationalStatus` |
| Type | `string` |

[Output Only] The current status of whether or not this interconnect
attachment is functional.
Valid values:
    OS_ACTIVE
    OS_UNPROVISIONED

### `pairingKey` {#InterconnectAttachment.pairingKey}

| Property | Comments |
|----------|----------|
| Field Name | `pairingKey` |
| Type | `string` |

[Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present
for DEDICATED]. The opaque identifier of an PARTNER attachment used to
initiate provisioning with a selected partner. Of the form
"XXXXX/region/domain"

### `partnerAsn` {#InterconnectAttachment.partnerAsn}

| Property | Comments |
|----------|----------|
| Field Name | `partnerAsn` |
| Type | `string` |

Optional BGP ASN for the router that should be supplied by a layer 3 Partner
if they configured BGP on behalf of the customer. Output only for PARTNER
type, input only for PARTNER_PROVIDER, not available for DEDICATED.

### `partnerMetadata` {#InterconnectAttachment.partnerMetadata}

| Property | Comments |
|----------|----------|
| Field Name | `partnerMetadata` |
| Type | [`compute.InterconnectAttachmentPartnerMetadata`](gcp_compute.md#InterconnectAttachmentPartnerMetadata) |

Informational metadata about Partner attachments from Partners to display to
customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER,
not available for DEDICATED.

### `privateInterconnectInfo` {#InterconnectAttachment.privateInterconnectInfo}

| Property | Comments |
|----------|----------|
| Field Name | `privateInterconnectInfo` |
| Type | [`compute.InterconnectAttachmentPrivateInfo`](gcp_compute.md#InterconnectAttachmentPrivateInfo) |

[Output Only] Information specific to an InterconnectAttachment. This
property is populated if the interconnect that this is attached to is of
type DEDICATED.

### `region` {#InterconnectAttachment.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the regional interconnect attachment
resides. You must specify this field as part of the HTTP request URL. It is
not settable as a field in the request body.

### `router` {#InterconnectAttachment.router}

| Property | Comments |
|----------|----------|
| Field Name | `router` |
| Type | `string` |

URL of the cloud router to be used for dynamic routing. This router must be
in the same region as this InterconnectAttachment. The
InterconnectAttachment will automatically connect the Interconnect to the
network & region within which the Cloud Router is configured.

### `selfLink` {#InterconnectAttachment.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `state` {#InterconnectAttachment.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

[Output Only] The current state of this attachment's functionality.
Valid values:
    ACTIVE
    DEFUNCT
    PARTNER_REQUEST_RECEIVED
    PENDING_CUSTOMER
    PENDING_PARTNER
    STATE_UNSPECIFIED
    UNPROVISIONED

### `type` {#InterconnectAttachment.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

Valid values:
    DEDICATED
    PARTNER
    PARTNER_PROVIDER

### `vlanTag8021q` {#InterconnectAttachment.vlanTag8021q}

| Property | Comments |
|----------|----------|
| Field Name | `vlanTag8021q` |
| Type | `int32` |

Available only for DEDICATED and PARTNER_PROVIDER. Desired VLAN tag for this
attachment, in the range 2-4094. This field refers to 802.1q VLAN tag, also
known as IEEE 802.1Q Only specified at creation time.

## Message `InterconnectAttachmentAggregatedList` {#InterconnectAttachmentAggregatedList}



### Inputs for `InterconnectAttachmentAggregatedList`

* `string` [`id`](#InterconnectAttachmentAggregatedList.id) = 1
* `repeated` [`compute.InterconnectAttachmentAggregatedList.ItemsEntry`](gcp_compute.md#InterconnectAttachmentAggregatedList.ItemsEntry) [`items`](#InterconnectAttachmentAggregatedList.items) = 2
* `string` [`kind`](#InterconnectAttachmentAggregatedList.kind) = 3
* `string` [`nextPageToken`](#InterconnectAttachmentAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#InterconnectAttachmentAggregatedList.selfLink) = 5
* [`compute.InterconnectAttachmentAggregatedList.Warning`](gcp_compute.md#InterconnectAttachmentAggregatedList.Warning) [`warning`](#InterconnectAttachmentAggregatedList.warning) = 6

### `id` {#InterconnectAttachmentAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InterconnectAttachmentAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InterconnectAttachmentAggregatedList.ItemsEntry`](gcp_compute.md#InterconnectAttachmentAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InterconnectAttachmentsScopedList resources.

### `kind` {#InterconnectAttachmentAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always
compute#interconnectAttachmentAggregatedList for aggregated lists of
interconnect attachments.

### `nextPageToken` {#InterconnectAttachmentAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InterconnectAttachmentAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InterconnectAttachmentAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InterconnectAttachmentAggregatedList.Warning`](gcp_compute.md#InterconnectAttachmentAggregatedList.Warning) |

## Message `InterconnectAttachmentList` {#InterconnectAttachmentList}

Response to the list request, and contains a list of interconnect
attachments.


### Inputs for `InterconnectAttachmentList`

* `string` [`id`](#InterconnectAttachmentList.id) = 1
* `repeated` [`compute.InterconnectAttachment`](gcp_compute.md#InterconnectAttachment) [`items`](#InterconnectAttachmentList.items) = 2
* `string` [`kind`](#InterconnectAttachmentList.kind) = 3
* `string` [`nextPageToken`](#InterconnectAttachmentList.nextPageToken) = 4
* `string` [`selfLink`](#InterconnectAttachmentList.selfLink) = 5
* [`compute.InterconnectAttachmentList.Warning`](gcp_compute.md#InterconnectAttachmentList.Warning) [`warning`](#InterconnectAttachmentList.warning) = 6

### `id` {#InterconnectAttachmentList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InterconnectAttachmentList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InterconnectAttachment`](gcp_compute.md#InterconnectAttachment) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InterconnectAttachment resources.

### `kind` {#InterconnectAttachmentList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#interconnectAttachmentList
for lists of interconnect attachments.

### `nextPageToken` {#InterconnectAttachmentList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InterconnectAttachmentList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InterconnectAttachmentList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InterconnectAttachmentList.Warning`](gcp_compute.md#InterconnectAttachmentList.Warning) |

## Message `InterconnectAttachmentPartnerMetadata` {#InterconnectAttachmentPartnerMetadata}

Informational metadata about Partner attachments from Partners to display to
customers. These fields are propagated from PARTNER_PROVIDER attachments to
their corresponding PARTNER attachments.


### Inputs for `InterconnectAttachmentPartnerMetadata`

* `string` [`interconnectName`](#InterconnectAttachmentPartnerMetadata.interconnectName) = 1
* `string` [`partnerName`](#InterconnectAttachmentPartnerMetadata.partnerName) = 2
* `string` [`portalUrl`](#InterconnectAttachmentPartnerMetadata.portalUrl) = 3

### `interconnectName` {#InterconnectAttachmentPartnerMetadata.interconnectName}

| Property | Comments |
|----------|----------|
| Field Name | `interconnectName` |
| Type | `string` |

Plain text name of the Interconnect this attachment is connected to, as
displayed in the Partner?s portal. For instance ?Chicago 1?. This value may
be validated to match approved Partner values.

### `partnerName` {#InterconnectAttachmentPartnerMetadata.partnerName}

| Property | Comments |
|----------|----------|
| Field Name | `partnerName` |
| Type | `string` |

Plain text name of the Partner providing this attachment. This value may be
validated to match approved Partner values.

### `portalUrl` {#InterconnectAttachmentPartnerMetadata.portalUrl}

| Property | Comments |
|----------|----------|
| Field Name | `portalUrl` |
| Type | `string` |

URL of the Partner?s portal for this Attachment. Partners may customise this
to be a deep-link to the specific resource on the Partner portal. This value
may be validated to match approved Partner values.

## Message `InterconnectAttachmentPrivateInfo` {#InterconnectAttachmentPrivateInfo}

Information for an interconnect attachment when this belongs to an
interconnect of type DEDICATED.


### Inputs for `InterconnectAttachmentPrivateInfo`

* `int32` [`tag8021q`](#InterconnectAttachmentPrivateInfo.tag8021q) = 1

### `tag8021q` {#InterconnectAttachmentPrivateInfo.tag8021q}

| Property | Comments |
|----------|----------|
| Field Name | `tag8021q` |
| Type | `int32` |

[Output Only] 802.1q encapsulation tag to be used for traffic between Google
and the customer, going to and from this network and region.

## Message `InterconnectAttachmentsScopedList` {#InterconnectAttachmentsScopedList}



### Inputs for `InterconnectAttachmentsScopedList`

* `repeated` [`compute.InterconnectAttachment`](gcp_compute.md#InterconnectAttachment) [`interconnectAttachments`](#InterconnectAttachmentsScopedList.interconnectAttachments) = 1
* [`compute.InterconnectAttachmentsScopedList.Warning`](gcp_compute.md#InterconnectAttachmentsScopedList.Warning) [`warning`](#InterconnectAttachmentsScopedList.warning) = 2

### `interconnectAttachments` {#InterconnectAttachmentsScopedList.interconnectAttachments}

| Property | Comments |
|----------|----------|
| Field Name | `interconnectAttachments` |
| Type | [`compute.InterconnectAttachment`](gcp_compute.md#InterconnectAttachment) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of interconnect attachments contained in this scope.

### `warning` {#InterconnectAttachmentsScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InterconnectAttachmentsScopedList.Warning`](gcp_compute.md#InterconnectAttachmentsScopedList.Warning) |

## Message `InterconnectCircuitInfo` {#InterconnectCircuitInfo}

Describes a single physical circuit between the Customer and Google.
CircuitInfo objects are created by Google, so all fields are output only.
Next id: 4


### Inputs for `InterconnectCircuitInfo`

* `string` [`customerDemarcId`](#InterconnectCircuitInfo.customerDemarcId) = 1
* `string` [`googleCircuitId`](#InterconnectCircuitInfo.googleCircuitId) = 2
* `string` [`googleDemarcId`](#InterconnectCircuitInfo.googleDemarcId) = 3

### `customerDemarcId` {#InterconnectCircuitInfo.customerDemarcId}

| Property | Comments |
|----------|----------|
| Field Name | `customerDemarcId` |
| Type | `string` |

Customer-side demarc ID for this circuit.

### `googleCircuitId` {#InterconnectCircuitInfo.googleCircuitId}

| Property | Comments |
|----------|----------|
| Field Name | `googleCircuitId` |
| Type | `string` |

Google-assigned unique ID for this circuit. Assigned at circuit turn-up.

### `googleDemarcId` {#InterconnectCircuitInfo.googleDemarcId}

| Property | Comments |
|----------|----------|
| Field Name | `googleDemarcId` |
| Type | `string` |

Google-side demarc ID for this circuit. Assigned at circuit turn-up and
provided by Google to the customer in the LOA.

## Message `InterconnectList` {#InterconnectList}

Response to the list request, and contains a list of interconnects.


### Inputs for `InterconnectList`

* `string` [`id`](#InterconnectList.id) = 1
* `repeated` [`compute.Interconnect`](gcp_compute.md#Interconnect) [`items`](#InterconnectList.items) = 2
* `string` [`kind`](#InterconnectList.kind) = 3
* `string` [`nextPageToken`](#InterconnectList.nextPageToken) = 4
* `string` [`selfLink`](#InterconnectList.selfLink) = 5
* [`compute.InterconnectList.Warning`](gcp_compute.md#InterconnectList.Warning) [`warning`](#InterconnectList.warning) = 6

### `id` {#InterconnectList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InterconnectList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Interconnect`](gcp_compute.md#Interconnect) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Interconnect resources.

### `kind` {#InterconnectList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#interconnectList for lists of
interconnects.

### `nextPageToken` {#InterconnectList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InterconnectList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InterconnectList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InterconnectList.Warning`](gcp_compute.md#InterconnectList.Warning) |

## Message `InterconnectLocation` {#InterconnectLocation}

Represents an InterconnectLocations resource. The InterconnectLocations
resource describes the locations where you can connect to Google's networks.
For more information, see  Colocation Facilities.


### Inputs for `InterconnectLocation`

* `string` [`address`](#InterconnectLocation.address) = 1
* `string` [`availabilityZone`](#InterconnectLocation.availabilityZone) = 2
* `string` [`city`](#InterconnectLocation.city) = 3
* `string` [`continent`](#InterconnectLocation.continent) = 4
* `string` [`creationTimestamp`](#InterconnectLocation.creationTimestamp) = 5
* `string` [`description`](#InterconnectLocation.description) = 6
* `string` [`facilityProvider`](#InterconnectLocation.facilityProvider) = 7
* `string` [`facilityProviderFacilityId`](#InterconnectLocation.facilityProviderFacilityId) = 8
* `string` [`id`](#InterconnectLocation.id) = 9
* `string` [`kind`](#InterconnectLocation.kind) = 10
* `string` [`name`](#InterconnectLocation.name) = 11 (**Required**)
* `string` [`peeringdbFacilityId`](#InterconnectLocation.peeringdbFacilityId) = 12
* `repeated` [`compute.InterconnectLocationRegionInfo`](gcp_compute.md#InterconnectLocationRegionInfo) [`regionInfos`](#InterconnectLocation.regionInfos) = 13
* `string` [`selfLink`](#InterconnectLocation.selfLink) = 14

### `address` {#InterconnectLocation.address}

| Property | Comments |
|----------|----------|
| Field Name | `address` |
| Type | `string` |

[Output Only] The postal address of the Point of Presence, each line in the
address is separated by a newline character.

### `availabilityZone` {#InterconnectLocation.availabilityZone}

| Property | Comments |
|----------|----------|
| Field Name | `availabilityZone` |
| Type | `string` |

[Output Only] Availability zone for this location. Within a metropolitan
area (metro), maintenance will not be simultaneously scheduled in more than
one availability zone. Example: "zone1" or "zone2".

### `city` {#InterconnectLocation.city}

| Property | Comments |
|----------|----------|
| Field Name | `city` |
| Type | `string` |

[Output Only] Metropolitan area designator that indicates which city an
interconnect is located. For example: "Chicago, IL", "Amsterdam,
Netherlands".

### `continent` {#InterconnectLocation.continent}

| Property | Comments |
|----------|----------|
| Field Name | `continent` |
| Type | `string` |

[Output Only] Continent for this location.
Valid values:
    AFRICA
    ASIA_PAC
    C_AFRICA
    C_ASIA_PAC
    C_EUROPE
    C_NORTH_AMERICA
    C_SOUTH_AMERICA
    EUROPE
    NORTH_AMERICA
    SOUTH_AMERICA

### `creationTimestamp` {#InterconnectLocation.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#InterconnectLocation.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] An optional description of the resource.

### `facilityProvider` {#InterconnectLocation.facilityProvider}

| Property | Comments |
|----------|----------|
| Field Name | `facilityProvider` |
| Type | `string` |

[Output Only] The name of the provider for this facility (e.g., EQUINIX).

### `facilityProviderFacilityId` {#InterconnectLocation.facilityProviderFacilityId}

| Property | Comments |
|----------|----------|
| Field Name | `facilityProviderFacilityId` |
| Type | `string` |

[Output Only] A provider-assigned Identifier for this facility (e.g.,
Ashburn-DC1).

### `id` {#InterconnectLocation.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#InterconnectLocation.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#interconnectLocation for
interconnect locations.

### `name` {#InterconnectLocation.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `peeringdbFacilityId` {#InterconnectLocation.peeringdbFacilityId}

| Property | Comments |
|----------|----------|
| Field Name | `peeringdbFacilityId` |
| Type | `string` |

[Output Only] The peeringdb identifier for this facility (corresponding with
a netfac type in peeringdb).

### `regionInfos` {#InterconnectLocation.regionInfos}

| Property | Comments |
|----------|----------|
| Field Name | `regionInfos` |
| Type | [`compute.InterconnectLocationRegionInfo`](gcp_compute.md#InterconnectLocationRegionInfo) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of InterconnectLocation.RegionInfo objects, that
describe parameters pertaining to the relation between this
InterconnectLocation and various Google Cloud regions.

### `selfLink` {#InterconnectLocation.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

## Message `InterconnectLocationList` {#InterconnectLocationList}

Response to the list request, and contains a list of interconnect locations.


### Inputs for `InterconnectLocationList`

* `string` [`id`](#InterconnectLocationList.id) = 1
* `repeated` [`compute.InterconnectLocation`](gcp_compute.md#InterconnectLocation) [`items`](#InterconnectLocationList.items) = 2
* `string` [`kind`](#InterconnectLocationList.kind) = 3
* `string` [`nextPageToken`](#InterconnectLocationList.nextPageToken) = 4
* `string` [`selfLink`](#InterconnectLocationList.selfLink) = 5
* [`compute.InterconnectLocationList.Warning`](gcp_compute.md#InterconnectLocationList.Warning) [`warning`](#InterconnectLocationList.warning) = 6

### `id` {#InterconnectLocationList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#InterconnectLocationList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InterconnectLocation`](gcp_compute.md#InterconnectLocation) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InterconnectLocation resources.

### `kind` {#InterconnectLocationList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#interconnectLocationList for
lists of interconnect locations.

### `nextPageToken` {#InterconnectLocationList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#InterconnectLocationList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#InterconnectLocationList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.InterconnectLocationList.Warning`](gcp_compute.md#InterconnectLocationList.Warning) |

## Message `InterconnectLocationRegionInfo` {#InterconnectLocationRegionInfo}

Information about any potential InterconnectAttachments between an
Interconnect at a specific InterconnectLocation, and a specific Cloud
Region.


### Inputs for `InterconnectLocationRegionInfo`

* `string` [`expectedRttMs`](#InterconnectLocationRegionInfo.expectedRttMs) = 1
* `string` [`locationPresence`](#InterconnectLocationRegionInfo.locationPresence) = 2
* `string` [`region`](#InterconnectLocationRegionInfo.region) = 3

### `expectedRttMs` {#InterconnectLocationRegionInfo.expectedRttMs}

| Property | Comments |
|----------|----------|
| Field Name | `expectedRttMs` |
| Type | `string` |

Expected round-trip time in milliseconds, from this InterconnectLocation to
a VM in this region.

### `locationPresence` {#InterconnectLocationRegionInfo.locationPresence}

| Property | Comments |
|----------|----------|
| Field Name | `locationPresence` |
| Type | `string` |

Identifies the network presence of this location.
Valid values:
    GLOBAL
    LOCAL_REGION
    LP_GLOBAL
    LP_LOCAL_REGION

### `region` {#InterconnectLocationRegionInfo.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

URL for the region of this location.

## Message `InterconnectOutageNotification` {#InterconnectOutageNotification}

Description of a planned outage on this Interconnect. Next id: 9


### Inputs for `InterconnectOutageNotification`

* `repeated` `string` [`affectedCircuits`](#InterconnectOutageNotification.affectedCircuits) = 1
* `string` [`description`](#InterconnectOutageNotification.description) = 2
* `string` [`endTime`](#InterconnectOutageNotification.endTime) = 3
* `string` [`issueType`](#InterconnectOutageNotification.issueType) = 4
* `string` [`name`](#InterconnectOutageNotification.name) = 5 (**Required**)
* `string` [`source`](#InterconnectOutageNotification.source) = 6
* `string` [`startTime`](#InterconnectOutageNotification.startTime) = 7
* `string` [`state`](#InterconnectOutageNotification.state) = 8

### `affectedCircuits` {#InterconnectOutageNotification.affectedCircuits}

| Property | Comments |
|----------|----------|
| Field Name | `affectedCircuits` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

If issue_type is IT_PARTIAL_OUTAGE, a list of the Google-side circuit IDs
that will be affected.

### `description` {#InterconnectOutageNotification.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

A description about the purpose of the outage.

### `endTime` {#InterconnectOutageNotification.endTime}

| Property | Comments |
|----------|----------|
| Field Name | `endTime` |
| Type | `string` |

Scheduled end time for the outage (milliseconds since Unix epoch).

### `issueType` {#InterconnectOutageNotification.issueType}

| Property | Comments |
|----------|----------|
| Field Name | `issueType` |
| Type | `string` |

Form this outage is expected to take. Note that the "IT_" versions of this
enum have been deprecated in favor of the unprefixed values.
Valid values:
    IT_OUTAGE
    IT_PARTIAL_OUTAGE
    OUTAGE
    PARTIAL_OUTAGE

### `name` {#InterconnectOutageNotification.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Unique identifier for this outage notification.

### `source` {#InterconnectOutageNotification.source}

| Property | Comments |
|----------|----------|
| Field Name | `source` |
| Type | `string` |

The party that generated this notification. Note that "NSRC_GOOGLE" has been
deprecated in favor of "GOOGLE"
Valid values:
    GOOGLE
    NSRC_GOOGLE

### `startTime` {#InterconnectOutageNotification.startTime}

| Property | Comments |
|----------|----------|
| Field Name | `startTime` |
| Type | `string` |

Scheduled start time for the outage (milliseconds since Unix epoch).

### `state` {#InterconnectOutageNotification.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

State of this notification. Note that the "NS_" versions of this enum have
been deprecated in favor of the unprefixed values.
Valid values:
    ACTIVE
    CANCELLED
    NS_ACTIVE
    NS_CANCELED

## Message `Items` {#Metadata.Items}

Array of key/value pairs. The total size of all keys and values must be less
than 512 KB.


### Inputs for `Items`

* `string` [`key`](#Metadata.Items.key) = 1
* `string` [`value`](#Metadata.Items.value) = 2

### `key` {#Metadata.Items.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

Key for the metadata entry. Keys must conform to the following regexp:
[a-zA-Z0-9-_]+, and be less than 128 bytes in length. This is reflected as
part of a URL in the metadata server. Additionally, to avoid ambiguity, keys
must not conflict with any other metadata keys for the project.

### `value` {#Metadata.Items.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

Value for the metadata entry. These are free-form strings, and only have
meaning as interpreted by the image running in the instance. The only
restriction placed on values is that their size must be less than or equal
to 262144 bytes (256 KiB).

## Message `ItemsEntry` {#InstanceGroupAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#InstanceGroupAggregatedList.ItemsEntry.key) = 1
* [`compute.InstanceGroupsScopedList`](gcp_compute.md#InstanceGroupsScopedList) [`value`](#InstanceGroupAggregatedList.ItemsEntry.value) = 2

### `key` {#InstanceGroupAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#InstanceGroupAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.InstanceGroupsScopedList`](gcp_compute.md#InstanceGroupsScopedList) |

## Message `ItemsEntry` {#DiskTypeAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#DiskTypeAggregatedList.ItemsEntry.key) = 1
* [`compute.DiskTypesScopedList`](gcp_compute.md#DiskTypesScopedList) [`value`](#DiskTypeAggregatedList.ItemsEntry.value) = 2

### `key` {#DiskTypeAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#DiskTypeAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.DiskTypesScopedList`](gcp_compute.md#DiskTypesScopedList) |

## Message `ItemsEntry` {#InstanceAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#InstanceAggregatedList.ItemsEntry.key) = 1
* [`compute.InstancesScopedList`](gcp_compute.md#InstancesScopedList) [`value`](#InstanceAggregatedList.ItemsEntry.value) = 2

### `key` {#InstanceAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#InstanceAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.InstancesScopedList`](gcp_compute.md#InstancesScopedList) |

## Message `ItemsEntry` {#TargetVpnGatewayAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#TargetVpnGatewayAggregatedList.ItemsEntry.key) = 1
* [`compute.TargetVpnGatewaysScopedList`](gcp_compute.md#TargetVpnGatewaysScopedList) [`value`](#TargetVpnGatewayAggregatedList.ItemsEntry.value) = 2

### `key` {#TargetVpnGatewayAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#TargetVpnGatewayAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.TargetVpnGatewaysScopedList`](gcp_compute.md#TargetVpnGatewaysScopedList) |

## Message `ItemsEntry` {#DiskAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#DiskAggregatedList.ItemsEntry.key) = 1
* [`compute.DisksScopedList`](gcp_compute.md#DisksScopedList) [`value`](#DiskAggregatedList.ItemsEntry.value) = 2

### `key` {#DiskAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#DiskAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.DisksScopedList`](gcp_compute.md#DisksScopedList) |

## Message `ItemsEntry` {#SubnetworkAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#SubnetworkAggregatedList.ItemsEntry.key) = 1
* [`compute.SubnetworksScopedList`](gcp_compute.md#SubnetworksScopedList) [`value`](#SubnetworkAggregatedList.ItemsEntry.value) = 2

### `key` {#SubnetworkAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#SubnetworkAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.SubnetworksScopedList`](gcp_compute.md#SubnetworksScopedList) |

## Message `ItemsEntry` {#InterconnectAttachmentAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#InterconnectAttachmentAggregatedList.ItemsEntry.key) = 1
* [`compute.InterconnectAttachmentsScopedList`](gcp_compute.md#InterconnectAttachmentsScopedList) [`value`](#InterconnectAttachmentAggregatedList.ItemsEntry.value) = 2

### `key` {#InterconnectAttachmentAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#InterconnectAttachmentAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.InterconnectAttachmentsScopedList`](gcp_compute.md#InterconnectAttachmentsScopedList) |

## Message `ItemsEntry` {#AddressAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#AddressAggregatedList.ItemsEntry.key) = 1
* [`compute.AddressesScopedList`](gcp_compute.md#AddressesScopedList) [`value`](#AddressAggregatedList.ItemsEntry.value) = 2

### `key` {#AddressAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#AddressAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.AddressesScopedList`](gcp_compute.md#AddressesScopedList) |

## Message `ItemsEntry` {#VpnTunnelAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#VpnTunnelAggregatedList.ItemsEntry.key) = 1
* [`compute.VpnTunnelsScopedList`](gcp_compute.md#VpnTunnelsScopedList) [`value`](#VpnTunnelAggregatedList.ItemsEntry.value) = 2

### `key` {#VpnTunnelAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#VpnTunnelAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.VpnTunnelsScopedList`](gcp_compute.md#VpnTunnelsScopedList) |

## Message `ItemsEntry` {#InstanceGroupManagerAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#InstanceGroupManagerAggregatedList.ItemsEntry.key) = 1
* [`compute.InstanceGroupManagersScopedList`](gcp_compute.md#InstanceGroupManagersScopedList) [`value`](#InstanceGroupManagerAggregatedList.ItemsEntry.value) = 2

### `key` {#InstanceGroupManagerAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#InstanceGroupManagerAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.InstanceGroupManagersScopedList`](gcp_compute.md#InstanceGroupManagersScopedList) |

## Message `ItemsEntry` {#AcceleratorTypeAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#AcceleratorTypeAggregatedList.ItemsEntry.key) = 1
* [`compute.AcceleratorTypesScopedList`](gcp_compute.md#AcceleratorTypesScopedList) [`value`](#AcceleratorTypeAggregatedList.ItemsEntry.value) = 2

### `key` {#AcceleratorTypeAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#AcceleratorTypeAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.AcceleratorTypesScopedList`](gcp_compute.md#AcceleratorTypesScopedList) |

## Message `ItemsEntry` {#RouterAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#RouterAggregatedList.ItemsEntry.key) = 1
* [`compute.RoutersScopedList`](gcp_compute.md#RoutersScopedList) [`value`](#RouterAggregatedList.ItemsEntry.value) = 2

### `key` {#RouterAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#RouterAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.RoutersScopedList`](gcp_compute.md#RoutersScopedList) |

## Message `ItemsEntry` {#ForwardingRuleAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#ForwardingRuleAggregatedList.ItemsEntry.key) = 1
* [`compute.ForwardingRulesScopedList`](gcp_compute.md#ForwardingRulesScopedList) [`value`](#ForwardingRuleAggregatedList.ItemsEntry.value) = 2

### `key` {#ForwardingRuleAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#ForwardingRuleAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.ForwardingRulesScopedList`](gcp_compute.md#ForwardingRulesScopedList) |

## Message `ItemsEntry` {#TargetInstanceAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#TargetInstanceAggregatedList.ItemsEntry.key) = 1
* [`compute.TargetInstancesScopedList`](gcp_compute.md#TargetInstancesScopedList) [`value`](#TargetInstanceAggregatedList.ItemsEntry.value) = 2

### `key` {#TargetInstanceAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#TargetInstanceAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.TargetInstancesScopedList`](gcp_compute.md#TargetInstancesScopedList) |

## Message `ItemsEntry` {#BackendServiceAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#BackendServiceAggregatedList.ItemsEntry.key) = 1
* [`compute.BackendServicesScopedList`](gcp_compute.md#BackendServicesScopedList) [`value`](#BackendServiceAggregatedList.ItemsEntry.value) = 2

### `key` {#BackendServiceAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#BackendServiceAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.BackendServicesScopedList`](gcp_compute.md#BackendServicesScopedList) |

## Message `ItemsEntry` {#CommitmentAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#CommitmentAggregatedList.ItemsEntry.key) = 1
* [`compute.CommitmentsScopedList`](gcp_compute.md#CommitmentsScopedList) [`value`](#CommitmentAggregatedList.ItemsEntry.value) = 2

### `key` {#CommitmentAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#CommitmentAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.CommitmentsScopedList`](gcp_compute.md#CommitmentsScopedList) |

## Message `ItemsEntry` {#MachineTypeAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#MachineTypeAggregatedList.ItemsEntry.key) = 1
* [`compute.MachineTypesScopedList`](gcp_compute.md#MachineTypesScopedList) [`value`](#MachineTypeAggregatedList.ItemsEntry.value) = 2

### `key` {#MachineTypeAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#MachineTypeAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.MachineTypesScopedList`](gcp_compute.md#MachineTypesScopedList) |

## Message `ItemsEntry` {#OperationAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#OperationAggregatedList.ItemsEntry.key) = 1
* [`compute.OperationsScopedList`](gcp_compute.md#OperationsScopedList) [`value`](#OperationAggregatedList.ItemsEntry.value) = 2

### `key` {#OperationAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#OperationAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.OperationsScopedList`](gcp_compute.md#OperationsScopedList) |

## Message `ItemsEntry` {#AutoscalerAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#AutoscalerAggregatedList.ItemsEntry.key) = 1
* [`compute.AutoscalersScopedList`](gcp_compute.md#AutoscalersScopedList) [`value`](#AutoscalerAggregatedList.ItemsEntry.value) = 2

### `key` {#AutoscalerAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#AutoscalerAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.AutoscalersScopedList`](gcp_compute.md#AutoscalersScopedList) |

## Message `ItemsEntry` {#TargetPoolAggregatedList.ItemsEntry}



### Inputs for `ItemsEntry`

* `string` [`key`](#TargetPoolAggregatedList.ItemsEntry.key) = 1
* [`compute.TargetPoolsScopedList`](gcp_compute.md#TargetPoolsScopedList) [`value`](#TargetPoolAggregatedList.ItemsEntry.value) = 2

### `key` {#TargetPoolAggregatedList.ItemsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#TargetPoolAggregatedList.ItemsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`compute.TargetPoolsScopedList`](gcp_compute.md#TargetPoolsScopedList) |

## Message `LabelsEntry` {#Disk.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#Disk.LabelsEntry.key) = 1
* `string` [`value`](#Disk.LabelsEntry.value) = 2

### `key` {#Disk.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#Disk.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#GlobalSetLabelsRequest.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#GlobalSetLabelsRequest.LabelsEntry.key) = 1
* `string` [`value`](#GlobalSetLabelsRequest.LabelsEntry.value) = 2

### `key` {#GlobalSetLabelsRequest.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#GlobalSetLabelsRequest.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#Instance.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#Instance.LabelsEntry.key) = 1
* `string` [`value`](#Instance.LabelsEntry.value) = 2

### `key` {#Instance.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#Instance.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#Snapshot.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#Snapshot.LabelsEntry.key) = 1
* `string` [`value`](#Snapshot.LabelsEntry.value) = 2

### `key` {#Snapshot.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#Snapshot.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#InstancesSetLabelsRequest.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#InstancesSetLabelsRequest.LabelsEntry.key) = 1
* `string` [`value`](#InstancesSetLabelsRequest.LabelsEntry.value) = 2

### `key` {#InstancesSetLabelsRequest.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#InstancesSetLabelsRequest.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#InstanceProperties.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#InstanceProperties.LabelsEntry.key) = 1
* `string` [`value`](#InstanceProperties.LabelsEntry.value) = 2

### `key` {#InstanceProperties.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#InstanceProperties.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#RegionSetLabelsRequest.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#RegionSetLabelsRequest.LabelsEntry.key) = 1
* `string` [`value`](#RegionSetLabelsRequest.LabelsEntry.value) = 2

### `key` {#RegionSetLabelsRequest.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#RegionSetLabelsRequest.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#ZoneSetLabelsRequest.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#ZoneSetLabelsRequest.LabelsEntry.key) = 1
* `string` [`value`](#ZoneSetLabelsRequest.LabelsEntry.value) = 2

### `key` {#ZoneSetLabelsRequest.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#ZoneSetLabelsRequest.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#AttachedDiskInitializeParams.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#AttachedDiskInitializeParams.LabelsEntry.key) = 1
* `string` [`value`](#AttachedDiskInitializeParams.LabelsEntry.value) = 2

### `key` {#AttachedDiskInitializeParams.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#AttachedDiskInitializeParams.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `LabelsEntry` {#Image.LabelsEntry}



### Inputs for `LabelsEntry`

* `string` [`key`](#Image.LabelsEntry.key) = 1
* `string` [`value`](#Image.LabelsEntry.value) = 2

### `key` {#Image.LabelsEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#Image.LabelsEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `License` {#License}

A license resource.


### Inputs for `License`

* `bool` [`chargesUseFee`](#License.chargesUseFee) = 1
* `string` [`creationTimestamp`](#License.creationTimestamp) = 2
* `string` [`description`](#License.description) = 3
* `string` [`id`](#License.id) = 4
* `string` [`kind`](#License.kind) = 5
* `string` [`licenseCode`](#License.licenseCode) = 6
* `string` [`name`](#License.name) = 7 (**Required**)
* [`compute.LicenseResourceRequirements`](gcp_compute.md#LicenseResourceRequirements) [`resourceRequirements`](#License.resourceRequirements) = 8
* `string` [`selfLink`](#License.selfLink) = 9
* `bool` [`transferable`](#License.transferable) = 10

### `chargesUseFee` {#License.chargesUseFee}

| Property | Comments |
|----------|----------|
| Field Name | `chargesUseFee` |
| Type | `bool` |

[Output Only] Deprecated. This field no longer reflects whether a license
charges a usage fee.

### `creationTimestamp` {#License.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#License.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional textual description of the resource; provided by the client when
the resource is created.

### `id` {#License.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#License.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#license for licenses.

### `licenseCode` {#License.licenseCode}

| Property | Comments |
|----------|----------|
| Field Name | `licenseCode` |
| Type | `string` |

[Output Only] The unique code used to attach this license to images,
snapshots, and disks.

### `name` {#License.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource. The name is 1-63 characters long and
complies with RFC1035.

### `resourceRequirements` {#License.resourceRequirements}

| Property | Comments |
|----------|----------|
| Field Name | `resourceRequirements` |
| Type | [`compute.LicenseResourceRequirements`](gcp_compute.md#LicenseResourceRequirements) |

### `selfLink` {#License.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `transferable` {#License.transferable}

| Property | Comments |
|----------|----------|
| Field Name | `transferable` |
| Type | `bool` |

If false, licenses will not be copied from the source resource when creating
an image from a disk, disk from snapshot, or snapshot from disk.

## Message `LicenseCode` {#LicenseCode}



### Inputs for `LicenseCode`

* `string` [`creationTimestamp`](#LicenseCode.creationTimestamp) = 1
* `string` [`description`](#LicenseCode.description) = 2
* `string` [`id`](#LicenseCode.id) = 3
* `string` [`kind`](#LicenseCode.kind) = 4
* `repeated` [`compute.LicenseCodeLicenseAlias`](gcp_compute.md#LicenseCodeLicenseAlias) [`licenseAlias`](#LicenseCode.licenseAlias) = 5
* `string` [`name`](#LicenseCode.name) = 6 (**Required**)
* `string` [`selfLink`](#LicenseCode.selfLink) = 7
* `string` [`state`](#LicenseCode.state) = 8
* `bool` [`transferable`](#LicenseCode.transferable) = 9

### `creationTimestamp` {#LicenseCode.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#LicenseCode.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] Description of this License Code.

### `id` {#LicenseCode.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#LicenseCode.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#licenseCode for licenses.

### `licenseAlias` {#LicenseCode.licenseAlias}

| Property | Comments |
|----------|----------|
| Field Name | `licenseAlias` |
| Type | [`compute.LicenseCodeLicenseAlias`](gcp_compute.md#LicenseCodeLicenseAlias) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] URL and description aliases of Licenses with the same License
Code.

### `name` {#LicenseCode.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource. The name is 1-20 characters long and
must be a valid 64 bit integer.

### `selfLink` {#LicenseCode.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `state` {#LicenseCode.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

[Output Only] Current state of this License Code.
Valid values:
    DISABLED
    ENABLED
    RESTRICTED
    STATE_UNSPECIFIED
    TERMINATED

### `transferable` {#LicenseCode.transferable}

| Property | Comments |
|----------|----------|
| Field Name | `transferable` |
| Type | `bool` |

[Output Only] If true, the license will remain attached when creating images
or snapshots from disks. Otherwise, the license is not transferred.

## Message `LicenseCodeLicenseAlias` {#LicenseCodeLicenseAlias}



### Inputs for `LicenseCodeLicenseAlias`

* `string` [`description`](#LicenseCodeLicenseAlias.description) = 1
* `string` [`selfLink`](#LicenseCodeLicenseAlias.selfLink) = 2

### `description` {#LicenseCodeLicenseAlias.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] Description of this License Code.

### `selfLink` {#LicenseCodeLicenseAlias.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] URL of license corresponding to this License Code.

## Message `LicenseResourceRequirements` {#LicenseResourceRequirements}



### Inputs for `LicenseResourceRequirements`

* `int32` [`minGuestCpuCount`](#LicenseResourceRequirements.minGuestCpuCount) = 1
* `int32` [`minMemoryMb`](#LicenseResourceRequirements.minMemoryMb) = 2

### `minGuestCpuCount` {#LicenseResourceRequirements.minGuestCpuCount}

| Property | Comments |
|----------|----------|
| Field Name | `minGuestCpuCount` |
| Type | `int32` |

Minimum number of guest cpus required to use the Instance. Enforced at
Instance creation and Instance start.

### `minMemoryMb` {#LicenseResourceRequirements.minMemoryMb}

| Property | Comments |
|----------|----------|
| Field Name | `minMemoryMb` |
| Type | `int32` |

Minimum memory required to use the Instance. Enforced at Instance creation
and Instance start.

## Message `LicensesListResponse` {#LicensesListResponse}



### Inputs for `LicensesListResponse`

* `string` [`id`](#LicensesListResponse.id) = 1
* `repeated` [`compute.License`](gcp_compute.md#License) [`items`](#LicensesListResponse.items) = 2
* `string` [`nextPageToken`](#LicensesListResponse.nextPageToken) = 3
* `string` [`selfLink`](#LicensesListResponse.selfLink) = 4
* [`compute.LicensesListResponse.Warning`](gcp_compute.md#LicensesListResponse.Warning) [`warning`](#LicensesListResponse.warning) = 5

### `id` {#LicensesListResponse.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#LicensesListResponse.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.License`](gcp_compute.md#License) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of License resources.

### `nextPageToken` {#LicensesListResponse.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#LicensesListResponse.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#LicensesListResponse.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.LicensesListResponse.Warning`](gcp_compute.md#LicensesListResponse.Warning) |

## Message `MachineType` {#MachineType}

A Machine Type resource. (== resource_for v1.machineTypes ==) (==
resource_for beta.machineTypes ==)


### Inputs for `MachineType`

* `string` [`creationTimestamp`](#MachineType.creationTimestamp) = 1
* [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) [`deprecated`](#MachineType.deprecated) = 2
* `string` [`description`](#MachineType.description) = 3
* `int32` [`guestCpus`](#MachineType.guestCpus) = 4
* `string` [`id`](#MachineType.id) = 5
* `int32` [`imageSpaceGb`](#MachineType.imageSpaceGb) = 6
* `bool` [`isSharedCpu`](#MachineType.isSharedCpu) = 7
* `string` [`kind`](#MachineType.kind) = 8
* `int32` [`maximumPersistentDisks`](#MachineType.maximumPersistentDisks) = 9
* `string` [`maximumPersistentDisksSizeGb`](#MachineType.maximumPersistentDisksSizeGb) = 10
* `int32` [`memoryMb`](#MachineType.memoryMb) = 11
* `string` [`name`](#MachineType.name) = 12 (**Required**)
* `repeated` [`compute.MachineType.ScratchDisks`](gcp_compute.md#MachineType.ScratchDisks) [`scratchDisks`](#MachineType.scratchDisks) = 13
* `string` [`selfLink`](#MachineType.selfLink) = 14
* `string` [`zone`](#MachineType.zone) = 15

### `creationTimestamp` {#MachineType.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `deprecated` {#MachineType.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) |

[Output Only] The deprecation status associated with this machine type.

### `description` {#MachineType.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] An optional textual description of the resource.

### `guestCpus` {#MachineType.guestCpus}

| Property | Comments |
|----------|----------|
| Field Name | `guestCpus` |
| Type | `int32` |

[Output Only] The number of virtual CPUs that are available to the instance.

### `id` {#MachineType.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `imageSpaceGb` {#MachineType.imageSpaceGb}

| Property | Comments |
|----------|----------|
| Field Name | `imageSpaceGb` |
| Type | `int32` |

[Deprecated] This property is deprecated and will never be populated with
any relevant values.

### `isSharedCpu` {#MachineType.isSharedCpu}

| Property | Comments |
|----------|----------|
| Field Name | `isSharedCpu` |
| Type | `bool` |

[Output Only] Whether this machine type has a shared CPU. See Shared-core
machine types for more information.

### `kind` {#MachineType.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The type of the resource. Always compute#machineType for
machine types.

### `maximumPersistentDisks` {#MachineType.maximumPersistentDisks}

| Property | Comments |
|----------|----------|
| Field Name | `maximumPersistentDisks` |
| Type | `int32` |

[Output Only] Maximum persistent disks allowed.

### `maximumPersistentDisksSizeGb` {#MachineType.maximumPersistentDisksSizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `maximumPersistentDisksSizeGb` |
| Type | `string` |

[Output Only] Maximum total persistent disks size (GB) allowed.

### `memoryMb` {#MachineType.memoryMb}

| Property | Comments |
|----------|----------|
| Field Name | `memoryMb` |
| Type | `int32` |

[Output Only] The amount of physical memory available to the instance,
defined in MB.

### `name` {#MachineType.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `scratchDisks` {#MachineType.scratchDisks}

| Property | Comments |
|----------|----------|
| Field Name | `scratchDisks` |
| Type | [`compute.MachineType.ScratchDisks`](gcp_compute.md#MachineType.ScratchDisks) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `selfLink` {#MachineType.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `zone` {#MachineType.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] The name of the zone where the machine type resides, such as
us-central1-a.

## Message `MachineTypeAggregatedList` {#MachineTypeAggregatedList}



### Inputs for `MachineTypeAggregatedList`

* `string` [`id`](#MachineTypeAggregatedList.id) = 1
* `repeated` [`compute.MachineTypeAggregatedList.ItemsEntry`](gcp_compute.md#MachineTypeAggregatedList.ItemsEntry) [`items`](#MachineTypeAggregatedList.items) = 2
* `string` [`kind`](#MachineTypeAggregatedList.kind) = 3
* `string` [`nextPageToken`](#MachineTypeAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#MachineTypeAggregatedList.selfLink) = 5
* [`compute.MachineTypeAggregatedList.Warning`](gcp_compute.md#MachineTypeAggregatedList.Warning) [`warning`](#MachineTypeAggregatedList.warning) = 6

### `id` {#MachineTypeAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#MachineTypeAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.MachineTypeAggregatedList.ItemsEntry`](gcp_compute.md#MachineTypeAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of MachineTypesScopedList resources.

### `kind` {#MachineTypeAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#machineTypeAggregatedList for
aggregated lists of machine types.

### `nextPageToken` {#MachineTypeAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#MachineTypeAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#MachineTypeAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.MachineTypeAggregatedList.Warning`](gcp_compute.md#MachineTypeAggregatedList.Warning) |

## Message `MachineTypeList` {#MachineTypeList}

Contains a list of machine types.


### Inputs for `MachineTypeList`

* `string` [`id`](#MachineTypeList.id) = 1
* `repeated` [`compute.MachineType`](gcp_compute.md#MachineType) [`items`](#MachineTypeList.items) = 2
* `string` [`kind`](#MachineTypeList.kind) = 3
* `string` [`nextPageToken`](#MachineTypeList.nextPageToken) = 4
* `string` [`selfLink`](#MachineTypeList.selfLink) = 5
* [`compute.MachineTypeList.Warning`](gcp_compute.md#MachineTypeList.Warning) [`warning`](#MachineTypeList.warning) = 6

### `id` {#MachineTypeList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#MachineTypeList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.MachineType`](gcp_compute.md#MachineType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of MachineType resources.

### `kind` {#MachineTypeList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#machineTypeList for lists of
machine types.

### `nextPageToken` {#MachineTypeList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#MachineTypeList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#MachineTypeList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.MachineTypeList.Warning`](gcp_compute.md#MachineTypeList.Warning) |

## Message `MachineTypesScopedList` {#MachineTypesScopedList}



### Inputs for `MachineTypesScopedList`

* `repeated` [`compute.MachineType`](gcp_compute.md#MachineType) [`machineTypes`](#MachineTypesScopedList.machineTypes) = 1
* [`compute.MachineTypesScopedList.Warning`](gcp_compute.md#MachineTypesScopedList.Warning) [`warning`](#MachineTypesScopedList.warning) = 2

### `machineTypes` {#MachineTypesScopedList.machineTypes}

| Property | Comments |
|----------|----------|
| Field Name | `machineTypes` |
| Type | [`compute.MachineType`](gcp_compute.md#MachineType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of machine types contained in this scope.

### `warning` {#MachineTypesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.MachineTypesScopedList.Warning`](gcp_compute.md#MachineTypesScopedList.Warning) |

## Message `ManagedInstance` {#ManagedInstance}

A Managed Instance resource.


### Inputs for `ManagedInstance`

* `string` [`currentAction`](#ManagedInstance.currentAction) = 1
* `string` [`id`](#ManagedInstance.id) = 2
* `string` [`instance`](#ManagedInstance.instance) = 3
* `string` [`instanceStatus`](#ManagedInstance.instanceStatus) = 4
* [`compute.ManagedInstanceLastAttempt`](gcp_compute.md#ManagedInstanceLastAttempt) [`lastAttempt`](#ManagedInstance.lastAttempt) = 5

### `currentAction` {#ManagedInstance.currentAction}

| Property | Comments |
|----------|----------|
| Field Name | `currentAction` |
| Type | `string` |

[Output Only] The current action that the managed instance group has
scheduled for the instance. Possible values:
- NONE The instance is running, and the managed instance group does not have
any scheduled actions for this instance.
- CREATING The managed instance group is creating this instance. If the
group fails to create this instance, it will try again until it is
successful.
- CREATING_WITHOUT_RETRIES The managed instance group is attempting to
create this instance only once. If the group fails to create this instance,
it does not try again and the group's targetSize value is decreased instead.
- RECREATING The managed instance group is recreating this instance.
- DELETING The managed instance group is permanently deleting this instance.
- ABANDONING The managed instance group is abandoning this instance. The
instance will be removed from the instance group and from any target pools
that are associated with this group.
- RESTARTING The managed instance group is restarting the instance.
- REFRESHING The managed instance group is applying configuration changes to
the instance without stopping it. For example, the group can update the
target pool list for an instance without stopping that instance.
- VERIFYING The managed instance group has created the instance and it is in
the process of being verified.
Valid values:
    ABANDONING
    CREATING
    CREATING_WITHOUT_RETRIES
    DELETING
    NONE
    RECREATING
    REFRESHING
    RESTARTING

### `id` {#ManagedInstance.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output only] The unique identifier for this resource. This field is empty
when instance does not exist.

### `instance` {#ManagedInstance.instance}

| Property | Comments |
|----------|----------|
| Field Name | `instance` |
| Type | `string` |

[Output Only] The URL of the instance. The URL can exist even if the
instance has not yet been created.

### `instanceStatus` {#ManagedInstance.instanceStatus}

| Property | Comments |
|----------|----------|
| Field Name | `instanceStatus` |
| Type | `string` |

[Output Only] The status of the instance. This field is empty when the
instance does not exist.
Valid values:
    PROVISIONING
    RUNNING
    STAGING
    STOPPED
    STOPPING
    SUSPENDED
    SUSPENDING
    TERMINATED

### `lastAttempt` {#ManagedInstance.lastAttempt}

| Property | Comments |
|----------|----------|
| Field Name | `lastAttempt` |
| Type | [`compute.ManagedInstanceLastAttempt`](gcp_compute.md#ManagedInstanceLastAttempt) |

[Output Only] Information about the last attempt to create or delete the
instance.

## Message `ManagedInstanceLastAttempt` {#ManagedInstanceLastAttempt}



### Inputs for `ManagedInstanceLastAttempt`

* [`compute.ManagedInstanceLastAttempt.Errors`](gcp_compute.md#ManagedInstanceLastAttempt.Errors) [`errors`](#ManagedInstanceLastAttempt.errors) = 1

### `errors` {#ManagedInstanceLastAttempt.errors}

| Property | Comments |
|----------|----------|
| Field Name | `errors` |
| Type | [`compute.ManagedInstanceLastAttempt.Errors`](gcp_compute.md#ManagedInstanceLastAttempt.Errors) |

## Message `Metadata` {#Metadata}

A metadata key/value entry.


### Inputs for `Metadata`

* `string` [`fingerprint`](#Metadata.fingerprint) = 1
* `repeated` [`compute.Metadata.Items`](gcp_compute.md#Metadata.Items) [`items`](#Metadata.items) = 2
* `string` [`kind`](#Metadata.kind) = 3

### `fingerprint` {#Metadata.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Specifies a fingerprint for this request, which is essentially a hash of the
metadata's contents and used for optimistic locking. The fingerprint is
initially generated by Compute Engine and changes after every request to
modify or update metadata. You must always provide an up-to-date fingerprint
hash in order to update or change metadata.

### `items` {#Metadata.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Metadata.Items`](gcp_compute.md#Metadata.Items) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `kind` {#Metadata.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#metadata for metadata.

## Message `NamedPort` {#NamedPort}

The named port. For example: .


### Inputs for `NamedPort`

* `string` [`name`](#NamedPort.name) = 1 (**Required**)
* `int32` [`port`](#NamedPort.port) = 2

### `name` {#NamedPort.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name for this named port. The name must be 1-63 characters long, and
comply with RFC1035.

### `port` {#NamedPort.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The port number, which can be a value between 1 and 65535.

## Message `Network` {#Network}

Represents a Network resource. Read Networks and Firewalls for more
information. (== resource_for v1.networks ==) (== resource_for beta.networks
==)


### Inputs for `Network`

* `string` [`IPv4Range`](#Network.IPv4Range) = 1
* `bool` [`autoCreateSubnetworks`](#Network.autoCreateSubnetworks) = 2
* `string` [`creationTimestamp`](#Network.creationTimestamp) = 3
* `string` [`description`](#Network.description) = 4
* `string` [`gatewayIPv4`](#Network.gatewayIPv4) = 5
* `string` [`id`](#Network.id) = 6
* `string` [`kind`](#Network.kind) = 7
* `string` [`name`](#Network.name) = 8 (**Required**)
* `repeated` [`compute.NetworkPeering`](gcp_compute.md#NetworkPeering) [`peerings`](#Network.peerings) = 9
* [`compute.NetworkRoutingConfig`](gcp_compute.md#NetworkRoutingConfig) [`routingConfig`](#Network.routingConfig) = 10
* `string` [`selfLink`](#Network.selfLink) = 11
* `repeated` `string` [`subnetworks`](#Network.subnetworks) = 12

### `IPv4Range` {#Network.IPv4Range}

| Property | Comments |
|----------|----------|
| Field Name | `IPv4Range` |
| Type | `string` |

The range of internal addresses that are legal on this network. This range
is a CIDR specification, for example: 192.168.0.0/16. Provided by the client
when the network is created.

### `autoCreateSubnetworks` {#Network.autoCreateSubnetworks}

| Property | Comments |
|----------|----------|
| Field Name | `autoCreateSubnetworks` |
| Type | `bool` |

When set to true, the network is created in "auto subnet mode". When set to
false, the network is in "custom subnet mode".

In "auto subnet mode", a newly created network is assigned the default CIDR
of 10.128.0.0/9 and it automatically creates one subnetwork per region.

### `creationTimestamp` {#Network.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Network.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `gatewayIPv4` {#Network.gatewayIPv4}

| Property | Comments |
|----------|----------|
| Field Name | `gatewayIPv4` |
| Type | `string` |

A gateway address for default routing to other networks. This value is read
only and is selected by the Google Compute Engine, typically as the first
usable address in the IPv4Range.

### `id` {#Network.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Network.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#network for networks.

### `name` {#Network.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `peerings` {#Network.peerings}

| Property | Comments |
|----------|----------|
| Field Name | `peerings` |
| Type | [`compute.NetworkPeering`](gcp_compute.md#NetworkPeering) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of network peerings for the resource.

### `routingConfig` {#Network.routingConfig}

| Property | Comments |
|----------|----------|
| Field Name | `routingConfig` |
| Type | [`compute.NetworkRoutingConfig`](gcp_compute.md#NetworkRoutingConfig) |

The network-level routing configuration for this network. Used by Cloud
Router to determine what type of network-wide routing behavior to enforce.

### `selfLink` {#Network.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `subnetworks` {#Network.subnetworks}

| Property | Comments |
|----------|----------|
| Field Name | `subnetworks` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Server-defined fully-qualified URLs for all subnetworks in
this network.

## Message `NetworkInterface` {#NetworkInterface}

A network interface resource attached to an instance.


### Inputs for `NetworkInterface`

* `repeated` [`compute.AccessConfig`](gcp_compute.md#AccessConfig) [`accessConfigs`](#NetworkInterface.accessConfigs) = 1
* `repeated` [`compute.AliasIpRange`](gcp_compute.md#AliasIpRange) [`aliasIpRanges`](#NetworkInterface.aliasIpRanges) = 2
* `string` [`fingerprint`](#NetworkInterface.fingerprint) = 3
* `string` [`kind`](#NetworkInterface.kind) = 4
* `string` [`name`](#NetworkInterface.name) = 5 (**Required**)
* `string` [`network`](#NetworkInterface.network) = 6
* `string` [`networkIP`](#NetworkInterface.networkIP) = 7
* `string` [`subnetwork`](#NetworkInterface.subnetwork) = 8

### `accessConfigs` {#NetworkInterface.accessConfigs}

| Property | Comments |
|----------|----------|
| Field Name | `accessConfigs` |
| Type | [`compute.AccessConfig`](gcp_compute.md#AccessConfig) |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of configurations for this interface. Currently, only one access
config, ONE_TO_ONE_NAT, is supported. If there are no accessConfigs
specified, then this instance will have no external internet access.

### `aliasIpRanges` {#NetworkInterface.aliasIpRanges}

| Property | Comments |
|----------|----------|
| Field Name | `aliasIpRanges` |
| Type | [`compute.AliasIpRange`](gcp_compute.md#AliasIpRange) |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of alias IP ranges for this network interface. Can only be
specified for network interfaces on subnet-mode networks.

### `fingerprint` {#NetworkInterface.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint hash of contents stored in this network interface. This field
will be ignored when inserting an Instance or adding a NetworkInterface. An
up-to-date fingerprint must be provided in order to update the
NetworkInterface.

### `kind` {#NetworkInterface.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#networkInterface for
network interfaces.

### `name` {#NetworkInterface.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] The name of the network interface, generated by the server.
For network devices, these are eth0, eth1, etc.

### `network` {#NetworkInterface.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

URL of the network resource for this instance. When creating an instance, if
neither the network nor the subnetwork is specified, the default network
global/networks/default is used; if the network is not specified but the
subnetwork is specified, the network is inferred.

This field is optional when creating a firewall rule. If not specified when
creating a firewall rule, the default network global/networks/default is
used.

If you specify this property, you can specify the network as a full or
partial URL. For example, the following are all valid URLs:
-
https://www.googleapis.com/compute/v1/projects/project/global/networks/network
- projects/project/global/networks/network
- global/networks/default

### `networkIP` {#NetworkInterface.networkIP}

| Property | Comments |
|----------|----------|
| Field Name | `networkIP` |
| Type | `string` |

An IPv4 internal network address to assign to the instance for this network
interface. If not specified by the user, an unused internal IP is assigned
by the system.

### `subnetwork` {#NetworkInterface.subnetwork}

| Property | Comments |
|----------|----------|
| Field Name | `subnetwork` |
| Type | `string` |

The URL of the Subnetwork resource for this instance. If the network
resource is in legacy mode, do not provide this property. If the network is
in auto subnet mode, providing the subnetwork is optional. If the network is
in custom subnet mode, then this field should be specified. If you specify
this property, you can specify the subnetwork as a full or partial URL. For
example, the following are all valid URLs:
-
https://www.googleapis.com/compute/v1/projects/project/regions/region/subnetworks/subnetwork
- regions/region/subnetworks/subnetwork

## Message `NetworkList` {#NetworkList}

Contains a list of networks.


### Inputs for `NetworkList`

* `string` [`id`](#NetworkList.id) = 1
* `repeated` [`compute.Network`](gcp_compute.md#Network) [`items`](#NetworkList.items) = 2
* `string` [`kind`](#NetworkList.kind) = 3
* `string` [`nextPageToken`](#NetworkList.nextPageToken) = 4
* `string` [`selfLink`](#NetworkList.selfLink) = 5
* [`compute.NetworkList.Warning`](gcp_compute.md#NetworkList.Warning) [`warning`](#NetworkList.warning) = 6

### `id` {#NetworkList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#NetworkList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Network`](gcp_compute.md#Network) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Network resources.

### `kind` {#NetworkList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#networkList for lists of
networks.

### `nextPageToken` {#NetworkList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#NetworkList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#NetworkList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.NetworkList.Warning`](gcp_compute.md#NetworkList.Warning) |

## Message `NetworkPeering` {#NetworkPeering}

A network peering attached to a network resource. The message includes the
peering name, peer network, peering state, and a flag indicating whether
Google Compute Engine should automatically create routes for the peering.


### Inputs for `NetworkPeering`

* `bool` [`autoCreateRoutes`](#NetworkPeering.autoCreateRoutes) = 1
* `string` [`name`](#NetworkPeering.name) = 2 (**Required**)
* `string` [`network`](#NetworkPeering.network) = 3
* `string` [`state`](#NetworkPeering.state) = 4
* `string` [`stateDetails`](#NetworkPeering.stateDetails) = 5

### `autoCreateRoutes` {#NetworkPeering.autoCreateRoutes}

| Property | Comments |
|----------|----------|
| Field Name | `autoCreateRoutes` |
| Type | `bool` |

Whether full mesh connectivity is created and managed automatically. When it
is set to true, Google Compute Engine will automatically create and manage
the routes between two networks when the state is ACTIVE. Otherwise, user
needs to create routes manually to route packets to peer network.

### `name` {#NetworkPeering.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of this peering. Provided by the client when the peering is created.
The name must comply with RFC1035. Specifically, the name must be 1-63
characters long and match regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`
which means the first character must be a lowercase letter, and all the
following characters must be a dash, lowercase letter, or digit, except the
last character, which cannot be a dash.

### `network` {#NetworkPeering.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

The URL of the peer network. It can be either full URL or partial URL. The
peer network may belong to a different project. If the partial URL does not
contain project, it is assumed that the peer network is in the same project
as the current network.

### `state` {#NetworkPeering.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

[Output Only] State for the peering.
Valid values:
    ACTIVE
    INACTIVE

### `stateDetails` {#NetworkPeering.stateDetails}

| Property | Comments |
|----------|----------|
| Field Name | `stateDetails` |
| Type | `string` |

[Output Only] Details about the current state of the peering.

## Message `NetworkRoutingConfig` {#NetworkRoutingConfig}

A routing configuration attached to a network resource. The message includes
the list of routers associated with the network, and a flag indicating the
type of routing behavior to enforce network-wide.


### Inputs for `NetworkRoutingConfig`

* `string` [`routingMode`](#NetworkRoutingConfig.routingMode) = 1

### `routingMode` {#NetworkRoutingConfig.routingMode}

| Property | Comments |
|----------|----------|
| Field Name | `routingMode` |
| Type | `string` |

The network-wide routing mode to use. If set to REGIONAL, this network's
cloud routers will only advertise routes with subnetworks of this network in
the same region as the router. If set to GLOBAL, this network's cloud
routers will advertise routes with all subnetworks of this network, across
regions.
Valid values:
    GLOBAL
    REGIONAL

## Message `NetworksAddPeeringRequest` {#NetworksAddPeeringRequest}



### Inputs for `NetworksAddPeeringRequest`

* `bool` [`autoCreateRoutes`](#NetworksAddPeeringRequest.autoCreateRoutes) = 1
* `string` [`name`](#NetworksAddPeeringRequest.name) = 2 (**Required**)
* `string` [`peerNetwork`](#NetworksAddPeeringRequest.peerNetwork) = 3

### `autoCreateRoutes` {#NetworksAddPeeringRequest.autoCreateRoutes}

| Property | Comments |
|----------|----------|
| Field Name | `autoCreateRoutes` |
| Type | `bool` |

Whether Google Compute Engine manages the routes automatically.

### `name` {#NetworksAddPeeringRequest.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the peering, which should conform to RFC1035.

### `peerNetwork` {#NetworksAddPeeringRequest.peerNetwork}

| Property | Comments |
|----------|----------|
| Field Name | `peerNetwork` |
| Type | `string` |

URL of the peer network. It can be either full URL or partial URL. The peer
network may belong to a different project. If the partial URL does not
contain project, it is assumed that the peer network is in the same project
as the current network.

## Message `NetworksRemovePeeringRequest` {#NetworksRemovePeeringRequest}



### Inputs for `NetworksRemovePeeringRequest`

* `string` [`name`](#NetworksRemovePeeringRequest.name) = 1 (**Required**)

### `name` {#NetworksRemovePeeringRequest.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the peering, which should conform to RFC1035.

## Message `Operation` {#Operation}

An Operation resource, used to manage asynchronous API requests. (==
resource_for v1.globalOperations ==) (== resource_for beta.globalOperations
==) (== resource_for v1.regionOperations ==) (== resource_for
beta.regionOperations ==) (== resource_for v1.zoneOperations ==) (==
resource_for beta.zoneOperations ==)


### Inputs for `Operation`

* `string` [`clientOperationId`](#Operation.clientOperationId) = 1
* `string` [`creationTimestamp`](#Operation.creationTimestamp) = 2
* `string` [`description`](#Operation.description) = 3
* `string` [`endTime`](#Operation.endTime) = 4
* [`compute.Operation.Error`](gcp_compute.md#Operation.Error) [`error`](#Operation.error) = 5
* `string` [`httpErrorMessage`](#Operation.httpErrorMessage) = 6
* `int32` [`httpErrorStatusCode`](#Operation.httpErrorStatusCode) = 7
* `string` [`id`](#Operation.id) = 8
* `string` [`insertTime`](#Operation.insertTime) = 9
* `string` [`kind`](#Operation.kind) = 10
* `string` [`name`](#Operation.name) = 11 (**Required**)
* `string` [`operationType`](#Operation.operationType) = 12
* `int32` [`progress`](#Operation.progress) = 13
* `string` [`region`](#Operation.region) = 14
* `string` [`selfLink`](#Operation.selfLink) = 15
* `string` [`startTime`](#Operation.startTime) = 16
* `string` [`status`](#Operation.status) = 17
* `string` [`statusMessage`](#Operation.statusMessage) = 18
* `string` [`targetId`](#Operation.targetId) = 19
* `string` [`targetLink`](#Operation.targetLink) = 20
* `string` [`user`](#Operation.user) = 21
* `repeated` [`compute.Operation.Warnings`](gcp_compute.md#Operation.Warnings) [`warnings`](#Operation.warnings) = 22
* `string` [`zone`](#Operation.zone) = 23

### `clientOperationId` {#Operation.clientOperationId}

| Property | Comments |
|----------|----------|
| Field Name | `clientOperationId` |
| Type | `string` |

[Output Only] The value of `requestId` if you provided it in the request.
Not present otherwise.

### `creationTimestamp` {#Operation.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Deprecated] This field is deprecated.

### `description` {#Operation.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] A textual description of the operation, which is set when the
operation is created.

### `endTime` {#Operation.endTime}

| Property | Comments |
|----------|----------|
| Field Name | `endTime` |
| Type | `string` |

[Output Only] The time that this operation was completed. This value is in
RFC3339 text format.

### `error` {#Operation.error}

| Property | Comments |
|----------|----------|
| Field Name | `error` |
| Type | [`compute.Operation.Error`](gcp_compute.md#Operation.Error) |

### `httpErrorMessage` {#Operation.httpErrorMessage}

| Property | Comments |
|----------|----------|
| Field Name | `httpErrorMessage` |
| Type | `string` |

[Output Only] If the operation fails, this field contains the HTTP error
message that was returned, such as NOT FOUND.

### `httpErrorStatusCode` {#Operation.httpErrorStatusCode}

| Property | Comments |
|----------|----------|
| Field Name | `httpErrorStatusCode` |
| Type | `int32` |

[Output Only] If the operation fails, this field contains the HTTP error
status code that was returned. For example, a 404 means the resource was not
found.

### `id` {#Operation.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `insertTime` {#Operation.insertTime}

| Property | Comments |
|----------|----------|
| Field Name | `insertTime` |
| Type | `string` |

[Output Only] The time that this operation was requested. This value is in
RFC3339 text format.

### `kind` {#Operation.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#operation for Operation
resources.

### `name` {#Operation.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `operationType` {#Operation.operationType}

| Property | Comments |
|----------|----------|
| Field Name | `operationType` |
| Type | `string` |

[Output Only] The type of operation, such as insert, update, or delete, and
so on.

### `progress` {#Operation.progress}

| Property | Comments |
|----------|----------|
| Field Name | `progress` |
| Type | `int32` |

[Output Only] An optional progress indicator that ranges from 0 to 100.
There is no requirement that this be linear or support any granularity of
operations. This should not be used to guess when the operation will be
complete. This number should monotonically increase as the operation
progresses.

### `region` {#Operation.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] The URL of the region where the operation resides. Only
available when performing regional operations. You must specify this field
as part of the HTTP request URL. It is not settable as a field in the
request body.

### `selfLink` {#Operation.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `startTime` {#Operation.startTime}

| Property | Comments |
|----------|----------|
| Field Name | `startTime` |
| Type | `string` |

[Output Only] The time that this operation was started by the server. This
value is in RFC3339 text format.

### `status` {#Operation.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the operation, which can be one of the
following: PENDING, RUNNING, or DONE.
Valid values:
    DONE
    PENDING
    RUNNING

### `statusMessage` {#Operation.statusMessage}

| Property | Comments |
|----------|----------|
| Field Name | `statusMessage` |
| Type | `string` |

[Output Only] An optional textual description of the current status of the
operation.

### `targetId` {#Operation.targetId}

| Property | Comments |
|----------|----------|
| Field Name | `targetId` |
| Type | `string` |

[Output Only] The unique target ID, which identifies a specific incarnation
of the target resource.

### `targetLink` {#Operation.targetLink}

| Property | Comments |
|----------|----------|
| Field Name | `targetLink` |
| Type | `string` |

[Output Only] The URL of the resource that the operation modifies. For
operations related to creating a snapshot, this points to the persistent
disk that the snapshot was created from.

### `user` {#Operation.user}

| Property | Comments |
|----------|----------|
| Field Name | `user` |
| Type | `string` |

[Output Only] User who requested the operation, for example:
user@example.com.

### `warnings` {#Operation.warnings}

| Property | Comments |
|----------|----------|
| Field Name | `warnings` |
| Type | [`compute.Operation.Warnings`](gcp_compute.md#Operation.Warnings) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `zone` {#Operation.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] The URL of the zone where the operation resides. Only
available when performing per-zone operations. You must specify this field
as part of the HTTP request URL. It is not settable as a field in the
request body.

## Message `OperationAggregatedList` {#OperationAggregatedList}



### Inputs for `OperationAggregatedList`

* `string` [`id`](#OperationAggregatedList.id) = 1
* `repeated` [`compute.OperationAggregatedList.ItemsEntry`](gcp_compute.md#OperationAggregatedList.ItemsEntry) [`items`](#OperationAggregatedList.items) = 2
* `string` [`kind`](#OperationAggregatedList.kind) = 3
* `string` [`nextPageToken`](#OperationAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#OperationAggregatedList.selfLink) = 5
* [`compute.OperationAggregatedList.Warning`](gcp_compute.md#OperationAggregatedList.Warning) [`warning`](#OperationAggregatedList.warning) = 6

### `id` {#OperationAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `items` {#OperationAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.OperationAggregatedList.ItemsEntry`](gcp_compute.md#OperationAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A map of scoped operation lists.

### `kind` {#OperationAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#operationAggregatedList for
aggregated lists of operations.

### `nextPageToken` {#OperationAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#OperationAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#OperationAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.OperationAggregatedList.Warning`](gcp_compute.md#OperationAggregatedList.Warning) |

## Message `OperationList` {#OperationList}

Contains a list of Operation resources.


### Inputs for `OperationList`

* `string` [`id`](#OperationList.id) = 1
* `repeated` [`compute.Operation`](gcp_compute.md#Operation) [`items`](#OperationList.items) = 2
* `string` [`kind`](#OperationList.kind) = 3
* `string` [`nextPageToken`](#OperationList.nextPageToken) = 4
* `string` [`selfLink`](#OperationList.selfLink) = 5
* [`compute.OperationList.Warning`](gcp_compute.md#OperationList.Warning) [`warning`](#OperationList.warning) = 6

### `id` {#OperationList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `items` {#OperationList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Operation`](gcp_compute.md#Operation) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of Operation resources.

### `kind` {#OperationList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#operations for Operations
resource.

### `nextPageToken` {#OperationList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#OperationList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#OperationList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.OperationList.Warning`](gcp_compute.md#OperationList.Warning) |

## Message `OperationsScopedList` {#OperationsScopedList}



### Inputs for `OperationsScopedList`

* `repeated` [`compute.Operation`](gcp_compute.md#Operation) [`operations`](#OperationsScopedList.operations) = 1
* [`compute.OperationsScopedList.Warning`](gcp_compute.md#OperationsScopedList.Warning) [`warning`](#OperationsScopedList.warning) = 2

### `operations` {#OperationsScopedList.operations}

| Property | Comments |
|----------|----------|
| Field Name | `operations` |
| Type | [`compute.Operation`](gcp_compute.md#Operation) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of operations contained in this scope.

### `warning` {#OperationsScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.OperationsScopedList.Warning`](gcp_compute.md#OperationsScopedList.Warning) |

## Message `PathMatcher` {#PathMatcher}

A matcher for the path portion of the URL. The BackendService from the
longest-matched rule will serve the URL. If no rule was matched, the default
service will be used.


### Inputs for `PathMatcher`

* `string` [`defaultService`](#PathMatcher.defaultService) = 1
* `string` [`description`](#PathMatcher.description) = 2
* `string` [`name`](#PathMatcher.name) = 3 (**Required**)
* `repeated` [`compute.PathRule`](gcp_compute.md#PathRule) [`pathRules`](#PathMatcher.pathRules) = 4

### `defaultService` {#PathMatcher.defaultService}

| Property | Comments |
|----------|----------|
| Field Name | `defaultService` |
| Type | `string` |

The full or partial URL to the BackendService resource. This will be used if
none of the pathRules defined by this PathMatcher is matched by the URL's
path portion. For example, the following are all valid URLs to a
BackendService resource:
-
https://www.googleapis.com/compute/v1/projects/project/global/backendServices/backendService
- compute/v1/projects/project/global/backendServices/backendService
- global/backendServices/backendService

### `description` {#PathMatcher.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `name` {#PathMatcher.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name to which this PathMatcher is referred by the HostRule.

### `pathRules` {#PathMatcher.pathRules}

| Property | Comments |
|----------|----------|
| Field Name | `pathRules` |
| Type | [`compute.PathRule`](gcp_compute.md#PathRule) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of path rules.

## Message `PathRule` {#PathRule}

A path-matching rule for a URL. If matched, will use the specified
BackendService to handle the traffic arriving at this URL.


### Inputs for `PathRule`

* `repeated` `string` [`paths`](#PathRule.paths) = 1
* `string` [`service`](#PathRule.service) = 2

### `paths` {#PathRule.paths}

| Property | Comments |
|----------|----------|
| Field Name | `paths` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of path patterns to match. Each must start with / and the only
place a * is allowed is at the end following a /. The string fed to the path
matcher does not include any text after the first ? or #, and those chars
are not allowed here.

### `service` {#PathRule.service}

| Property | Comments |
|----------|----------|
| Field Name | `service` |
| Type | `string` |

The URL of the BackendService resource if this rule is matched.

## Message `Project` {#Project}

A Project resource. For an overview of projects, see  Cloud Platform
Resource Hierarchy. (== resource_for v1.projects ==) (== resource_for
beta.projects ==)


### Inputs for `Project`

* [`compute.Metadata`](gcp_compute.md#Metadata) [`commonInstanceMetadata`](#Project.commonInstanceMetadata) = 1
* `string` [`creationTimestamp`](#Project.creationTimestamp) = 2
* `string` [`defaultServiceAccount`](#Project.defaultServiceAccount) = 3
* `string` [`description`](#Project.description) = 4
* `repeated` `string` [`enabledFeatures`](#Project.enabledFeatures) = 5
* `string` [`id`](#Project.id) = 6
* `string` [`kind`](#Project.kind) = 7
* `string` [`name`](#Project.name) = 8 (**Required**)
* `repeated` [`compute.Quota`](gcp_compute.md#Quota) [`quotas`](#Project.quotas) = 9
* `string` [`selfLink`](#Project.selfLink) = 10
* [`compute.UsageExportLocation`](gcp_compute.md#UsageExportLocation) [`usageExportLocation`](#Project.usageExportLocation) = 11
* `string` [`xpnProjectStatus`](#Project.xpnProjectStatus) = 12

### `commonInstanceMetadata` {#Project.commonInstanceMetadata}

| Property | Comments |
|----------|----------|
| Field Name | `commonInstanceMetadata` |
| Type | [`compute.Metadata`](gcp_compute.md#Metadata) |

Metadata key/value pairs available to all instances contained in this
project. See Custom metadata for more information.

### `creationTimestamp` {#Project.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `defaultServiceAccount` {#Project.defaultServiceAccount}

| Property | Comments |
|----------|----------|
| Field Name | `defaultServiceAccount` |
| Type | `string` |

[Output Only] Default service account used by VMs running in this project.

### `description` {#Project.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional textual description of the resource.

### `enabledFeatures` {#Project.enabledFeatures}

| Property | Comments |
|----------|----------|
| Field Name | `enabledFeatures` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Restricted features enabled for use on this project.

### `id` {#Project.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server. This is not the project ID, and is just a unique ID
used by Compute Engine to identify resources.

### `kind` {#Project.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#project for projects.

### `name` {#Project.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The project ID. For example: my-example-project. Use the project ID to make
requests to Compute Engine.

### `quotas` {#Project.quotas}

| Property | Comments |
|----------|----------|
| Field Name | `quotas` |
| Type | [`compute.Quota`](gcp_compute.md#Quota) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Quotas assigned to this project.

### `selfLink` {#Project.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `usageExportLocation` {#Project.usageExportLocation}

| Property | Comments |
|----------|----------|
| Field Name | `usageExportLocation` |
| Type | [`compute.UsageExportLocation`](gcp_compute.md#UsageExportLocation) |

The naming prefix for daily usage reports and the Google Cloud Storage
bucket where they are stored.

### `xpnProjectStatus` {#Project.xpnProjectStatus}

| Property | Comments |
|----------|----------|
| Field Name | `xpnProjectStatus` |
| Type | `string` |

[Output Only] The role this project has in a shared VPC configuration.
Currently only HOST projects are differentiated.
Valid values:
    HOST
    UNSPECIFIED_XPN_PROJECT_STATUS

## Message `ProjectsDisableXpnResourceRequest` {#ProjectsDisableXpnResourceRequest}



### Inputs for `ProjectsDisableXpnResourceRequest`

* [`compute.XpnResourceId`](gcp_compute.md#XpnResourceId) [`xpnResource`](#ProjectsDisableXpnResourceRequest.xpnResource) = 1

### `xpnResource` {#ProjectsDisableXpnResourceRequest.xpnResource}

| Property | Comments |
|----------|----------|
| Field Name | `xpnResource` |
| Type | [`compute.XpnResourceId`](gcp_compute.md#XpnResourceId) |

Service resource (a.k.a service project) ID.

## Message `ProjectsEnableXpnResourceRequest` {#ProjectsEnableXpnResourceRequest}



### Inputs for `ProjectsEnableXpnResourceRequest`

* [`compute.XpnResourceId`](gcp_compute.md#XpnResourceId) [`xpnResource`](#ProjectsEnableXpnResourceRequest.xpnResource) = 1

### `xpnResource` {#ProjectsEnableXpnResourceRequest.xpnResource}

| Property | Comments |
|----------|----------|
| Field Name | `xpnResource` |
| Type | [`compute.XpnResourceId`](gcp_compute.md#XpnResourceId) |

Service resource (a.k.a service project) ID.

## Message `ProjectsGetXpnResources` {#ProjectsGetXpnResources}



### Inputs for `ProjectsGetXpnResources`

* `string` [`kind`](#ProjectsGetXpnResources.kind) = 1
* `string` [`nextPageToken`](#ProjectsGetXpnResources.nextPageToken) = 2
* `repeated` [`compute.XpnResourceId`](gcp_compute.md#XpnResourceId) [`resources`](#ProjectsGetXpnResources.resources) = 3

### `kind` {#ProjectsGetXpnResources.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#projectsGetXpnResources for
lists of service resources (a.k.a service projects)

### `nextPageToken` {#ProjectsGetXpnResources.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `resources` {#ProjectsGetXpnResources.resources}

| Property | Comments |
|----------|----------|
| Field Name | `resources` |
| Type | [`compute.XpnResourceId`](gcp_compute.md#XpnResourceId) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Service resources (a.k.a service projects) attached to this project as their
shared VPC host.

## Message `ProjectsListXpnHostsRequest` {#ProjectsListXpnHostsRequest}



### Inputs for `ProjectsListXpnHostsRequest`

* `string` [`organization`](#ProjectsListXpnHostsRequest.organization) = 1

### `organization` {#ProjectsListXpnHostsRequest.organization}

| Property | Comments |
|----------|----------|
| Field Name | `organization` |
| Type | `string` |

Optional organization ID managed by Cloud Resource Manager, for which to
list shared VPC host projects. If not specified, the organization will be
inferred from the project.

## Message `Quota` {#Quota}

A quotas entry.


### Inputs for `Quota`

* `TYPE_DOUBLE` [`limit`](#Quota.limit) = 1
* `string` [`metric`](#Quota.metric) = 2
* `TYPE_DOUBLE` [`usage`](#Quota.usage) = 3

### `limit` {#Quota.limit}

| Property | Comments |
|----------|----------|
| Field Name | `limit` |
| Type | `TYPE_DOUBLE` |

[Output Only] Quota limit for this metric.

### `metric` {#Quota.metric}

| Property | Comments |
|----------|----------|
| Field Name | `metric` |
| Type | `string` |

[Output Only] Name of the quota metric.
Valid values:
    AUTOSCALERS
    BACKEND_BUCKETS
    BACKEND_SERVICES
    COMMITMENTS
    CPUS
    CPUS_ALL_REGIONS
    DISKS_TOTAL_GB
    FIREWALLS
    FORWARDING_RULES
    HEALTH_CHECKS
    IMAGES
    INSTANCES
    INSTANCE_GROUPS
    INSTANCE_GROUP_MANAGERS
    INSTANCE_TEMPLATES
    INTERCONNECTS
    INTERCONNECT_ATTACHMENTS_PER_REGION
    INTERCONNECT_ATTACHMENTS_TOTAL_MBPS
    INTERNAL_ADDRESSES
    IN_USE_ADDRESSES
    LOCAL_SSD_TOTAL_GB
    NETWORKS
    NVIDIA_K80_GPUS
    NVIDIA_P100_GPUS
    NVIDIA_V100_GPUS
    PREEMPTIBLE_CPUS
    PREEMPTIBLE_LOCAL_SSD_GB
    PREEMPTIBLE_NVIDIA_K80_GPUS
    PREEMPTIBLE_NVIDIA_P100_GPUS
    PREEMPTIBLE_NVIDIA_V100_GPUS
    REGIONAL_AUTOSCALERS
    REGIONAL_INSTANCE_GROUP_MANAGERS
    ROUTERS
    ROUTES
    SECURITY_POLICIES
    SECURITY_POLICY_RULES
    SNAPSHOTS
    SSD_TOTAL_GB
    SSL_CERTIFICATES
    STATIC_ADDRESSES
    SUBNETWORKS
    TARGET_HTTPS_PROXIES
    TARGET_HTTP_PROXIES
    TARGET_INSTANCES
    TARGET_POOLS
    TARGET_SSL_PROXIES
    TARGET_TCP_PROXIES
    TARGET_VPN_GATEWAYS
    URL_MAPS
    VPN_TUNNELS

### `usage` {#Quota.usage}

| Property | Comments |
|----------|----------|
| Field Name | `usage` |
| Type | `TYPE_DOUBLE` |

[Output Only] Current usage of this metric.

## Message `RawDisk` {#Image.RawDisk}

The parameters of the raw disk image.
The parameters of the raw disk image.


### Inputs for `RawDisk`

* `string` [`containerType`](#Image.RawDisk.containerType) = 1
* `string` [`sha1Checksum`](#Image.RawDisk.sha1Checksum) = 2
* `string` [`source`](#Image.RawDisk.source) = 3

### `containerType` {#Image.RawDisk.containerType}

| Property | Comments |
|----------|----------|
| Field Name | `containerType` |
| Type | `string` |

The format used to encode and transmit the block device, which should be
TAR. This is just a container and transmission format and not a runtime
format. Provided by the client when the disk image is created.
Valid values:
    TAR

### `sha1Checksum` {#Image.RawDisk.sha1Checksum}

| Property | Comments |
|----------|----------|
| Field Name | `sha1Checksum` |
| Type | `string` |

An optional SHA1 checksum of the disk image before unpackaging; provided by
the client when the disk image is created.

### `source` {#Image.RawDisk.source}

| Property | Comments |
|----------|----------|
| Field Name | `source` |
| Type | `string` |

The full Google Cloud Storage URL where the disk image is stored. You must
provide either this property or the sourceDisk property but not both.

## Message `Reference` {#Reference}

Represents a reference to a resource.


### Inputs for `Reference`

* `string` [`kind`](#Reference.kind) = 1
* `string` [`referenceType`](#Reference.referenceType) = 2
* `string` [`referrer`](#Reference.referrer) = 3
* `string` [`target`](#Reference.target) = 4

### `kind` {#Reference.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#reference for references.

### `referenceType` {#Reference.referenceType}

| Property | Comments |
|----------|----------|
| Field Name | `referenceType` |
| Type | `string` |

A description of the reference type with no implied semantics. Possible
values include:
- MEMBER_OF

### `referrer` {#Reference.referrer}

| Property | Comments |
|----------|----------|
| Field Name | `referrer` |
| Type | `string` |

URL of the resource which refers to the target.

### `target` {#Reference.target}

| Property | Comments |
|----------|----------|
| Field Name | `target` |
| Type | `string` |

URL of the resource to which this reference points.

## Message `Region` {#Region}

Region resource. (== resource_for beta.regions ==) (== resource_for
v1.regions ==)


### Inputs for `Region`

* `string` [`creationTimestamp`](#Region.creationTimestamp) = 1
* [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) [`deprecated`](#Region.deprecated) = 2
* `string` [`description`](#Region.description) = 3
* `string` [`id`](#Region.id) = 4
* `string` [`kind`](#Region.kind) = 5
* `string` [`name`](#Region.name) = 6 (**Required**)
* `repeated` [`compute.Quota`](gcp_compute.md#Quota) [`quotas`](#Region.quotas) = 7
* `string` [`selfLink`](#Region.selfLink) = 8
* `string` [`status`](#Region.status) = 9
* `repeated` `string` [`zones`](#Region.zones) = 10

### `creationTimestamp` {#Region.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `deprecated` {#Region.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) |

[Output Only] The deprecation status associated with this region.

### `description` {#Region.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] Textual description of the resource.

### `id` {#Region.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Region.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#region for regions.

### `name` {#Region.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `quotas` {#Region.quotas}

| Property | Comments |
|----------|----------|
| Field Name | `quotas` |
| Type | [`compute.Quota`](gcp_compute.md#Quota) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Quotas assigned to this region.

### `selfLink` {#Region.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `status` {#Region.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] Status of the region, either UP or DOWN.
Valid values:
    DOWN
    UP

### `zones` {#Region.zones}

| Property | Comments |
|----------|----------|
| Field Name | `zones` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of zones available in this region, in the form of
resource URLs.

## Message `RegionAutoscalerList` {#RegionAutoscalerList}

Contains a list of autoscalers.


### Inputs for `RegionAutoscalerList`

* `string` [`id`](#RegionAutoscalerList.id) = 1
* `repeated` [`compute.Autoscaler`](gcp_compute.md#Autoscaler) [`items`](#RegionAutoscalerList.items) = 2
* `string` [`kind`](#RegionAutoscalerList.kind) = 3
* `string` [`nextPageToken`](#RegionAutoscalerList.nextPageToken) = 4
* `string` [`selfLink`](#RegionAutoscalerList.selfLink) = 5
* [`compute.RegionAutoscalerList.Warning`](gcp_compute.md#RegionAutoscalerList.Warning) [`warning`](#RegionAutoscalerList.warning) = 6

### `id` {#RegionAutoscalerList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RegionAutoscalerList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Autoscaler`](gcp_compute.md#Autoscaler) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Autoscaler resources.

### `kind` {#RegionAutoscalerList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#RegionAutoscalerList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RegionAutoscalerList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RegionAutoscalerList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RegionAutoscalerList.Warning`](gcp_compute.md#RegionAutoscalerList.Warning) |

## Message `RegionDiskTypeList` {#RegionDiskTypeList}



### Inputs for `RegionDiskTypeList`

* `string` [`id`](#RegionDiskTypeList.id) = 1
* `repeated` [`compute.DiskType`](gcp_compute.md#DiskType) [`items`](#RegionDiskTypeList.items) = 2
* `string` [`kind`](#RegionDiskTypeList.kind) = 3
* `string` [`nextPageToken`](#RegionDiskTypeList.nextPageToken) = 4
* `string` [`selfLink`](#RegionDiskTypeList.selfLink) = 5
* [`compute.RegionDiskTypeList.Warning`](gcp_compute.md#RegionDiskTypeList.Warning) [`warning`](#RegionDiskTypeList.warning) = 6

### `id` {#RegionDiskTypeList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RegionDiskTypeList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.DiskType`](gcp_compute.md#DiskType) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of DiskType resources.

### `kind` {#RegionDiskTypeList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#regionDiskTypeList for region
disk types.

### `nextPageToken` {#RegionDiskTypeList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RegionDiskTypeList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RegionDiskTypeList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RegionDiskTypeList.Warning`](gcp_compute.md#RegionDiskTypeList.Warning) |

## Message `RegionDisksResizeRequest` {#RegionDisksResizeRequest}



### Inputs for `RegionDisksResizeRequest`

* `string` [`sizeGb`](#RegionDisksResizeRequest.sizeGb) = 1

### `sizeGb` {#RegionDisksResizeRequest.sizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `sizeGb` |
| Type | `string` |

The new size of the regional persistent disk, which is specified in GB.

## Message `RegionInstanceGroupList` {#RegionInstanceGroupList}

Contains a list of InstanceGroup resources.


### Inputs for `RegionInstanceGroupList`

* `string` [`id`](#RegionInstanceGroupList.id) = 1
* `repeated` [`compute.InstanceGroup`](gcp_compute.md#InstanceGroup) [`items`](#RegionInstanceGroupList.items) = 2
* `string` [`kind`](#RegionInstanceGroupList.kind) = 3
* `string` [`nextPageToken`](#RegionInstanceGroupList.nextPageToken) = 4
* `string` [`selfLink`](#RegionInstanceGroupList.selfLink) = 5
* [`compute.RegionInstanceGroupList.Warning`](gcp_compute.md#RegionInstanceGroupList.Warning) [`warning`](#RegionInstanceGroupList.warning) = 6

### `id` {#RegionInstanceGroupList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RegionInstanceGroupList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceGroup`](gcp_compute.md#InstanceGroup) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceGroup resources.

### `kind` {#RegionInstanceGroupList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

The resource type.

### `nextPageToken` {#RegionInstanceGroupList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RegionInstanceGroupList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RegionInstanceGroupList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RegionInstanceGroupList.Warning`](gcp_compute.md#RegionInstanceGroupList.Warning) |

## Message `RegionInstanceGroupManagerList` {#RegionInstanceGroupManagerList}

Contains a list of managed instance groups.


### Inputs for `RegionInstanceGroupManagerList`

* `string` [`id`](#RegionInstanceGroupManagerList.id) = 1
* `repeated` [`compute.InstanceGroupManager`](gcp_compute.md#InstanceGroupManager) [`items`](#RegionInstanceGroupManagerList.items) = 2
* `string` [`kind`](#RegionInstanceGroupManagerList.kind) = 3
* `string` [`nextPageToken`](#RegionInstanceGroupManagerList.nextPageToken) = 4
* `string` [`selfLink`](#RegionInstanceGroupManagerList.selfLink) = 5
* [`compute.RegionInstanceGroupManagerList.Warning`](gcp_compute.md#RegionInstanceGroupManagerList.Warning) [`warning`](#RegionInstanceGroupManagerList.warning) = 6

### `id` {#RegionInstanceGroupManagerList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RegionInstanceGroupManagerList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceGroupManager`](gcp_compute.md#InstanceGroupManager) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceGroupManager resources.

### `kind` {#RegionInstanceGroupManagerList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The resource type, which is always
compute#instanceGroupManagerList for a list of managed instance groups that
exist in th regional scope.

### `nextPageToken` {#RegionInstanceGroupManagerList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RegionInstanceGroupManagerList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RegionInstanceGroupManagerList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RegionInstanceGroupManagerList.Warning`](gcp_compute.md#RegionInstanceGroupManagerList.Warning) |

## Message `RegionInstanceGroupManagersAbandonInstancesRequest` {#RegionInstanceGroupManagersAbandonInstancesRequest}



### Inputs for `RegionInstanceGroupManagersAbandonInstancesRequest`

* `repeated` `string` [`instances`](#RegionInstanceGroupManagersAbandonInstancesRequest.instances) = 1

### `instances` {#RegionInstanceGroupManagersAbandonInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs of one or more instances to abandon. This can be a full URL or a
partial URL, such as zones/[ZONE]/instances/[INSTANCE_NAME].

## Message `RegionInstanceGroupManagersDeleteInstancesRequest` {#RegionInstanceGroupManagersDeleteInstancesRequest}



### Inputs for `RegionInstanceGroupManagersDeleteInstancesRequest`

* `repeated` `string` [`instances`](#RegionInstanceGroupManagersDeleteInstancesRequest.instances) = 1

### `instances` {#RegionInstanceGroupManagersDeleteInstancesRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs of one or more instances to delete. This can be a full URL or a
partial URL, such as zones/[ZONE]/instances/[INSTANCE_NAME].

## Message `RegionInstanceGroupManagersListInstancesResponse` {#RegionInstanceGroupManagersListInstancesResponse}



### Inputs for `RegionInstanceGroupManagersListInstancesResponse`

* `repeated` [`compute.ManagedInstance`](gcp_compute.md#ManagedInstance) [`managedInstances`](#RegionInstanceGroupManagersListInstancesResponse.managedInstances) = 1

### `managedInstances` {#RegionInstanceGroupManagersListInstancesResponse.managedInstances}

| Property | Comments |
|----------|----------|
| Field Name | `managedInstances` |
| Type | [`compute.ManagedInstance`](gcp_compute.md#ManagedInstance) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of managed instances.

## Message `RegionInstanceGroupManagersRecreateRequest` {#RegionInstanceGroupManagersRecreateRequest}



### Inputs for `RegionInstanceGroupManagersRecreateRequest`

* `repeated` `string` [`instances`](#RegionInstanceGroupManagersRecreateRequest.instances) = 1

### `instances` {#RegionInstanceGroupManagersRecreateRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URLs of one or more instances to recreate. This can be a full URL or a
partial URL, such as zones/[ZONE]/instances/[INSTANCE_NAME].

## Message `RegionInstanceGroupManagersSetTargetPoolsRequest` {#RegionInstanceGroupManagersSetTargetPoolsRequest}



### Inputs for `RegionInstanceGroupManagersSetTargetPoolsRequest`

* `string` [`fingerprint`](#RegionInstanceGroupManagersSetTargetPoolsRequest.fingerprint) = 1
* `repeated` `string` [`targetPools`](#RegionInstanceGroupManagersSetTargetPoolsRequest.targetPools) = 2

### `fingerprint` {#RegionInstanceGroupManagersSetTargetPoolsRequest.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint of the target pools information, which is a hash of the
contents. This field is used for optimistic locking when you update the
target pool entries. This field is optional.

### `targetPools` {#RegionInstanceGroupManagersSetTargetPoolsRequest.targetPools}

| Property | Comments |
|----------|----------|
| Field Name | `targetPools` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URL of all TargetPool resources to which instances in the instanceGroup
field are added. The target pools automatically apply to all of the
instances in the managed instance group.

## Message `RegionInstanceGroupManagersSetTemplateRequest` {#RegionInstanceGroupManagersSetTemplateRequest}



### Inputs for `RegionInstanceGroupManagersSetTemplateRequest`

* `string` [`instanceTemplate`](#RegionInstanceGroupManagersSetTemplateRequest.instanceTemplate) = 1

### `instanceTemplate` {#RegionInstanceGroupManagersSetTemplateRequest.instanceTemplate}

| Property | Comments |
|----------|----------|
| Field Name | `instanceTemplate` |
| Type | `string` |

URL of the InstanceTemplate resource from which all new instances will be
created.

## Message `RegionInstanceGroupsListInstances` {#RegionInstanceGroupsListInstances}



### Inputs for `RegionInstanceGroupsListInstances`

* `string` [`id`](#RegionInstanceGroupsListInstances.id) = 1
* `repeated` [`compute.InstanceWithNamedPorts`](gcp_compute.md#InstanceWithNamedPorts) [`items`](#RegionInstanceGroupsListInstances.items) = 2
* `string` [`kind`](#RegionInstanceGroupsListInstances.kind) = 3
* `string` [`nextPageToken`](#RegionInstanceGroupsListInstances.nextPageToken) = 4
* `string` [`selfLink`](#RegionInstanceGroupsListInstances.selfLink) = 5
* [`compute.RegionInstanceGroupsListInstances.Warning`](gcp_compute.md#RegionInstanceGroupsListInstances.Warning) [`warning`](#RegionInstanceGroupsListInstances.warning) = 6

### `id` {#RegionInstanceGroupsListInstances.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RegionInstanceGroupsListInstances.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.InstanceWithNamedPorts`](gcp_compute.md#InstanceWithNamedPorts) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of InstanceWithNamedPorts resources.

### `kind` {#RegionInstanceGroupsListInstances.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

The resource type.

### `nextPageToken` {#RegionInstanceGroupsListInstances.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RegionInstanceGroupsListInstances.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RegionInstanceGroupsListInstances.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RegionInstanceGroupsListInstances.Warning`](gcp_compute.md#RegionInstanceGroupsListInstances.Warning) |

## Message `RegionInstanceGroupsListInstancesRequest` {#RegionInstanceGroupsListInstancesRequest}



### Inputs for `RegionInstanceGroupsListInstancesRequest`

* `string` [`instanceState`](#RegionInstanceGroupsListInstancesRequest.instanceState) = 1
* `string` [`portName`](#RegionInstanceGroupsListInstancesRequest.portName) = 2

### `instanceState` {#RegionInstanceGroupsListInstancesRequest.instanceState}

| Property | Comments |
|----------|----------|
| Field Name | `instanceState` |
| Type | `string` |

Instances in which state should be returned. Valid options are: 'ALL',
'RUNNING'. By default, it lists all instances.
Valid values:
    ALL
    RUNNING

### `portName` {#RegionInstanceGroupsListInstancesRequest.portName}

| Property | Comments |
|----------|----------|
| Field Name | `portName` |
| Type | `string` |

Name of port user is interested in. It is optional. If it is set, only
information about this ports will be returned. If it is not set, all the
named ports will be returned. Always lists all instances.

## Message `RegionInstanceGroupsSetNamedPortsRequest` {#RegionInstanceGroupsSetNamedPortsRequest}



### Inputs for `RegionInstanceGroupsSetNamedPortsRequest`

* `string` [`fingerprint`](#RegionInstanceGroupsSetNamedPortsRequest.fingerprint) = 1
* `repeated` [`compute.NamedPort`](gcp_compute.md#NamedPort) [`namedPorts`](#RegionInstanceGroupsSetNamedPortsRequest.namedPorts) = 2

### `fingerprint` {#RegionInstanceGroupsSetNamedPortsRequest.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

The fingerprint of the named ports information for this instance group. Use
this optional property to prevent conflicts when multiple users change the
named ports settings concurrently. Obtain the fingerprint with the
instanceGroups.get method. Then, include the fingerprint in your request to
ensure that you do not overwrite changes that were applied from another
concurrent request.

### `namedPorts` {#RegionInstanceGroupsSetNamedPortsRequest.namedPorts}

| Property | Comments |
|----------|----------|
| Field Name | `namedPorts` |
| Type | [`compute.NamedPort`](gcp_compute.md#NamedPort) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of named ports to set for this instance group.

## Message `RegionList` {#RegionList}

Contains a list of region resources.


### Inputs for `RegionList`

* `string` [`id`](#RegionList.id) = 1
* `repeated` [`compute.Region`](gcp_compute.md#Region) [`items`](#RegionList.items) = 2
* `string` [`kind`](#RegionList.kind) = 3
* `string` [`nextPageToken`](#RegionList.nextPageToken) = 4
* `string` [`selfLink`](#RegionList.selfLink) = 5
* [`compute.RegionList.Warning`](gcp_compute.md#RegionList.Warning) [`warning`](#RegionList.warning) = 6

### `id` {#RegionList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RegionList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Region`](gcp_compute.md#Region) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Region resources.

### `kind` {#RegionList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#regionList for lists of
regions.

### `nextPageToken` {#RegionList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RegionList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RegionList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RegionList.Warning`](gcp_compute.md#RegionList.Warning) |

## Message `RegionSetLabelsRequest` {#RegionSetLabelsRequest}



### Inputs for `RegionSetLabelsRequest`

* `string` [`labelFingerprint`](#RegionSetLabelsRequest.labelFingerprint) = 1
* `repeated` [`compute.RegionSetLabelsRequest.LabelsEntry`](gcp_compute.md#RegionSetLabelsRequest.LabelsEntry) [`labels`](#RegionSetLabelsRequest.labels) = 2

### `labelFingerprint` {#RegionSetLabelsRequest.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

The fingerprint of the previous set of labels for this resource, used to
detect conflicts. The fingerprint is initially generated by Compute Engine
and changes after every request to modify or update labels. You must always
provide an up-to-date fingerprint hash in order to update or change labels.
Make a get() request to the resource to get the latest fingerprint.

### `labels` {#RegionSetLabelsRequest.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.RegionSetLabelsRequest.LabelsEntry`](gcp_compute.md#RegionSetLabelsRequest.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The labels to set for this resource.

## Message `ResourceCommitment` {#ResourceCommitment}

Commitment for a particular resource (a Commitment is composed of one or
more of these).


### Inputs for `ResourceCommitment`

* `string` [`amount`](#ResourceCommitment.amount) = 1
* `string` [`type`](#ResourceCommitment.type) = 2

### `amount` {#ResourceCommitment.amount}

| Property | Comments |
|----------|----------|
| Field Name | `amount` |
| Type | `string` |

The amount of the resource purchased (in a type-dependent unit, such as
bytes). For vCPUs, this can just be an integer. For memory, this must be
provided in MB. Memory must be a multiple of 256 MB, with up to 6.5GB of
memory per every vCPU.

### `type` {#ResourceCommitment.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

Type of resource for which this commitment applies. Possible values are VCPU
and MEMORY
Valid values:
    MEMORY
    UNSPECIFIED
    VCPU

## Message `ResourceGroupReference` {#ResourceGroupReference}



### Inputs for `ResourceGroupReference`

* `string` [`group`](#ResourceGroupReference.group) = 1

### `group` {#ResourceGroupReference.group}

| Property | Comments |
|----------|----------|
| Field Name | `group` |
| Type | `string` |

A URI referencing one of the instance groups listed in the backend service.

## Message `Route` {#Route}

Represents a Route resource. A route specifies how certain packets should be
handled by the network. Routes are associated with instances by tags and the
set of routes for a particular instance is called its routing table.

For each packet leaving an instance, the system searches that instance's
routing table for a single best matching route. Routes match packets by
destination IP address, preferring smaller or more specific ranges over
larger ones. If there is a tie, the system selects the route with the
smallest priority value. If there is still a tie, it uses the layer three
and four packet headers to select just one of the remaining matching routes.
The packet is then forwarded as specified by the nextHop field of the
winning route - either to another instance destination, an instance gateway,
or a Google Compute Engine-operated gateway.

Packets that do not match any route in the sending instance's routing table
are dropped. (== resource_for beta.routes ==) (== resource_for v1.routes ==)


### Inputs for `Route`

* `string` [`creationTimestamp`](#Route.creationTimestamp) = 1
* `string` [`description`](#Route.description) = 2
* `string` [`destRange`](#Route.destRange) = 3
* `string` [`id`](#Route.id) = 4
* `string` [`kind`](#Route.kind) = 5
* `string` [`name`](#Route.name) = 6 (**Required**)
* `string` [`network`](#Route.network) = 7
* `string` [`nextHopGateway`](#Route.nextHopGateway) = 8
* `string` [`nextHopInstance`](#Route.nextHopInstance) = 9
* `string` [`nextHopIp`](#Route.nextHopIp) = 10
* `string` [`nextHopNetwork`](#Route.nextHopNetwork) = 11
* `string` [`nextHopPeering`](#Route.nextHopPeering) = 12
* `string` [`nextHopVpnTunnel`](#Route.nextHopVpnTunnel) = 13
* `int32` [`priority`](#Route.priority) = 14
* `string` [`selfLink`](#Route.selfLink) = 15
* `repeated` `string` [`tags`](#Route.tags) = 16
* `repeated` [`compute.Route.Warnings`](gcp_compute.md#Route.Warnings) [`warnings`](#Route.warnings) = 17

### `creationTimestamp` {#Route.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Route.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `destRange` {#Route.destRange}

| Property | Comments |
|----------|----------|
| Field Name | `destRange` |
| Type | `string` |

The destination range of outgoing packets that this route applies to. Only
IPv4 is supported.

### `id` {#Route.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Route.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of this resource. Always compute#routes for Route
resources.

### `name` {#Route.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `network` {#Route.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

Fully-qualified URL of the network that this route applies to.

### `nextHopGateway` {#Route.nextHopGateway}

| Property | Comments |
|----------|----------|
| Field Name | `nextHopGateway` |
| Type | `string` |

The URL to a gateway that should handle matching packets. You can only
specify the internet gateway using a full or partial valid URL: 
projects/<project-id>/global/gateways/default-internet-gateway

### `nextHopInstance` {#Route.nextHopInstance}

| Property | Comments |
|----------|----------|
| Field Name | `nextHopInstance` |
| Type | `string` |

The URL to an instance that should handle matching packets. You can specify
this as a full or partial URL. For example:
https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/

### `nextHopIp` {#Route.nextHopIp}

| Property | Comments |
|----------|----------|
| Field Name | `nextHopIp` |
| Type | `string` |

The network IP address of an instance that should handle matching packets.
Only IPv4 is supported.

### `nextHopNetwork` {#Route.nextHopNetwork}

| Property | Comments |
|----------|----------|
| Field Name | `nextHopNetwork` |
| Type | `string` |

The URL of the local network if it should handle matching packets.

### `nextHopPeering` {#Route.nextHopPeering}

| Property | Comments |
|----------|----------|
| Field Name | `nextHopPeering` |
| Type | `string` |

[Output Only] The network peering name that should handle matching packets,
which should conform to RFC1035.

### `nextHopVpnTunnel` {#Route.nextHopVpnTunnel}

| Property | Comments |
|----------|----------|
| Field Name | `nextHopVpnTunnel` |
| Type | `string` |

The URL to a VpnTunnel that should handle matching packets.

### `priority` {#Route.priority}

| Property | Comments |
|----------|----------|
| Field Name | `priority` |
| Type | `int32` |

The priority of this route. Priority is used to break ties in cases where
there is more than one matching route of equal prefix length. In the case of
two routes with equal prefix length, the one with the lowest-numbered
priority value wins. Default value is 1000. Valid range is 0 through 65535.

### `selfLink` {#Route.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined fully-qualified URL for this resource.

### `tags` {#Route.tags}

| Property | Comments |
|----------|----------|
| Field Name | `tags` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of instance tags to which this route applies.

### `warnings` {#Route.warnings}

| Property | Comments |
|----------|----------|
| Field Name | `warnings` |
| Type | [`compute.Route.Warnings`](gcp_compute.md#Route.Warnings) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `RouteList` {#RouteList}

Contains a list of Route resources.


### Inputs for `RouteList`

* `string` [`id`](#RouteList.id) = 1
* `repeated` [`compute.Route`](gcp_compute.md#Route) [`items`](#RouteList.items) = 2
* `string` [`kind`](#RouteList.kind) = 3
* `string` [`nextPageToken`](#RouteList.nextPageToken) = 4
* `string` [`selfLink`](#RouteList.selfLink) = 5
* [`compute.RouteList.Warning`](gcp_compute.md#RouteList.Warning) [`warning`](#RouteList.warning) = 6

### `id` {#RouteList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RouteList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Route`](gcp_compute.md#Route) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Route resources.

### `kind` {#RouteList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#RouteList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RouteList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RouteList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RouteList.Warning`](gcp_compute.md#RouteList.Warning) |

## Message `Router` {#Router}

Router resource.


### Inputs for `Router`

* [`compute.RouterBgp`](gcp_compute.md#RouterBgp) [`bgp`](#Router.bgp) = 1
* `repeated` [`compute.RouterBgpPeer`](gcp_compute.md#RouterBgpPeer) [`bgpPeers`](#Router.bgpPeers) = 2
* `string` [`creationTimestamp`](#Router.creationTimestamp) = 3
* `string` [`description`](#Router.description) = 4
* `string` [`id`](#Router.id) = 5
* `repeated` [`compute.RouterInterface`](gcp_compute.md#RouterInterface) [`interfaces`](#Router.interfaces) = 6
* `string` [`kind`](#Router.kind) = 7
* `string` [`name`](#Router.name) = 8 (**Required**)
* `string` [`network`](#Router.network) = 9
* `string` [`region`](#Router.region) = 10
* `string` [`selfLink`](#Router.selfLink) = 11

### `bgp` {#Router.bgp}

| Property | Comments |
|----------|----------|
| Field Name | `bgp` |
| Type | [`compute.RouterBgp`](gcp_compute.md#RouterBgp) |

BGP information specific to this router.

### `bgpPeers` {#Router.bgpPeers}

| Property | Comments |
|----------|----------|
| Field Name | `bgpPeers` |
| Type | [`compute.RouterBgpPeer`](gcp_compute.md#RouterBgpPeer) |
| Repeated | Any number of instances of this type is allowed in the schema. |

BGP information that needs to be configured into the routing stack to
establish the BGP peering. It must specify peer ASN and either interface
name, IP, or peer IP. Please refer to RFC4273.

### `creationTimestamp` {#Router.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Router.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#Router.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `interfaces` {#Router.interfaces}

| Property | Comments |
|----------|----------|
| Field Name | `interfaces` |
| Type | [`compute.RouterInterface`](gcp_compute.md#RouterInterface) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Router interfaces. Each interface requires either one linked resource (e.g.
linkedVpnTunnel), or IP address and IP address range (e.g. ipRange), or
both.

### `kind` {#Router.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#router for routers.

### `name` {#Router.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `network` {#Router.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

URI of the network to which this router belongs.

### `region` {#Router.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URI of the region where the router resides. You must specify
this field as part of the HTTP request URL. It is not settable as a field in
the request body.

### `selfLink` {#Router.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

## Message `RouterAdvertisedIpRange` {#RouterAdvertisedIpRange}

Description-tagged IP ranges for the router to advertise.


### Inputs for `RouterAdvertisedIpRange`

* `string` [`description`](#RouterAdvertisedIpRange.description) = 1
* `string` [`range`](#RouterAdvertisedIpRange.range) = 2

### `description` {#RouterAdvertisedIpRange.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

User-specified description for the IP range.

### `range` {#RouterAdvertisedIpRange.range}

| Property | Comments |
|----------|----------|
| Field Name | `range` |
| Type | `string` |

The IP range to advertise. The value must be a CIDR-formatted string.

## Message `RouterAggregatedList` {#RouterAggregatedList}

Contains a list of routers.


### Inputs for `RouterAggregatedList`

* `string` [`id`](#RouterAggregatedList.id) = 1
* `repeated` [`compute.RouterAggregatedList.ItemsEntry`](gcp_compute.md#RouterAggregatedList.ItemsEntry) [`items`](#RouterAggregatedList.items) = 2
* `string` [`kind`](#RouterAggregatedList.kind) = 3
* `string` [`nextPageToken`](#RouterAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#RouterAggregatedList.selfLink) = 5
* [`compute.RouterAggregatedList.Warning`](gcp_compute.md#RouterAggregatedList.Warning) [`warning`](#RouterAggregatedList.warning) = 6

### `id` {#RouterAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RouterAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.RouterAggregatedList.ItemsEntry`](gcp_compute.md#RouterAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Router resources.

### `kind` {#RouterAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#RouterAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RouterAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RouterAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RouterAggregatedList.Warning`](gcp_compute.md#RouterAggregatedList.Warning) |

## Message `RouterBgp` {#RouterBgp}



### Inputs for `RouterBgp`

* `string` [`advertiseMode`](#RouterBgp.advertiseMode) = 1
* `repeated` `string` [`advertisedGroups`](#RouterBgp.advertisedGroups) = 2
* `repeated` [`compute.RouterAdvertisedIpRange`](gcp_compute.md#RouterAdvertisedIpRange) [`advertisedIpRanges`](#RouterBgp.advertisedIpRanges) = 3
* `int32` [`asn`](#RouterBgp.asn) = 4

### `advertiseMode` {#RouterBgp.advertiseMode}

| Property | Comments |
|----------|----------|
| Field Name | `advertiseMode` |
| Type | `string` |

User-specified flag to indicate which mode to use for advertisement.
Valid values:
    CUSTOM
    DEFAULT

### `advertisedGroups` {#RouterBgp.advertisedGroups}

| Property | Comments |
|----------|----------|
| Field Name | `advertisedGroups` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

User-specified list of prefix groups to advertise in custom mode. This field
can only be populated if advertise_mode is CUSTOM and is advertised to all
peers of the router. These groups will be advertised in addition to any
specified prefixes. Leave this field blank to advertise no custom groups.

### `advertisedIpRanges` {#RouterBgp.advertisedIpRanges}

| Property | Comments |
|----------|----------|
| Field Name | `advertisedIpRanges` |
| Type | [`compute.RouterAdvertisedIpRange`](gcp_compute.md#RouterAdvertisedIpRange) |
| Repeated | Any number of instances of this type is allowed in the schema. |

User-specified list of individual IP ranges to advertise in custom mode.
This field can only be populated if advertise_mode is CUSTOM and is
advertised to all peers of the router. These IP ranges will be advertised in
addition to any specified groups. Leave this field blank to advertise no
custom IP ranges.

### `asn` {#RouterBgp.asn}

| Property | Comments |
|----------|----------|
| Field Name | `asn` |
| Type | `int32` |

Local BGP Autonomous System Number (ASN). Must be an RFC6996 private ASN,
either 16-bit or 32-bit. The value will be fixed for this router resource.
All VPN tunnels that link to this router will have the same local ASN.

## Message `RouterBgpPeer` {#RouterBgpPeer}



### Inputs for `RouterBgpPeer`

* `string` [`advertiseMode`](#RouterBgpPeer.advertiseMode) = 1
* `repeated` `string` [`advertisedGroups`](#RouterBgpPeer.advertisedGroups) = 2
* `repeated` [`compute.RouterAdvertisedIpRange`](gcp_compute.md#RouterAdvertisedIpRange) [`advertisedIpRanges`](#RouterBgpPeer.advertisedIpRanges) = 3
* `int32` [`advertisedRoutePriority`](#RouterBgpPeer.advertisedRoutePriority) = 4
* `string` [`interfaceName`](#RouterBgpPeer.interfaceName) = 5
* `string` [`ipAddress`](#RouterBgpPeer.ipAddress) = 6
* `string` [`managementType`](#RouterBgpPeer.managementType) = 7
* `string` [`name`](#RouterBgpPeer.name) = 8 (**Required**)
* `int32` [`peerAsn`](#RouterBgpPeer.peerAsn) = 9
* `string` [`peerIpAddress`](#RouterBgpPeer.peerIpAddress) = 10

### `advertiseMode` {#RouterBgpPeer.advertiseMode}

| Property | Comments |
|----------|----------|
| Field Name | `advertiseMode` |
| Type | `string` |

User-specified flag to indicate which mode to use for advertisement.
Valid values:
    CUSTOM
    DEFAULT

### `advertisedGroups` {#RouterBgpPeer.advertisedGroups}

| Property | Comments |
|----------|----------|
| Field Name | `advertisedGroups` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

User-specified list of prefix groups to advertise in custom mode. This field
can only be populated if advertise_mode is CUSTOM and overrides the list
defined for the router (in Bgp message). These groups will be advertised in
addition to any specified prefixes. Leave this field blank to advertise no
custom groups.

### `advertisedIpRanges` {#RouterBgpPeer.advertisedIpRanges}

| Property | Comments |
|----------|----------|
| Field Name | `advertisedIpRanges` |
| Type | [`compute.RouterAdvertisedIpRange`](gcp_compute.md#RouterAdvertisedIpRange) |
| Repeated | Any number of instances of this type is allowed in the schema. |

User-specified list of individual IP ranges to advertise in custom mode.
This field can only be populated if advertise_mode is CUSTOM and overrides
the list defined for the router (in Bgp message). These IP ranges will be
advertised in addition to any specified groups. Leave this field blank to
advertise no custom IP ranges.

### `advertisedRoutePriority` {#RouterBgpPeer.advertisedRoutePriority}

| Property | Comments |
|----------|----------|
| Field Name | `advertisedRoutePriority` |
| Type | `int32` |

The priority of routes advertised to this BGP peer. In the case where there
is more than one matching route of maximum length, the routes with lowest
priority value win.

### `interfaceName` {#RouterBgpPeer.interfaceName}

| Property | Comments |
|----------|----------|
| Field Name | `interfaceName` |
| Type | `string` |

Name of the interface the BGP peer is associated with.

### `ipAddress` {#RouterBgpPeer.ipAddress}

| Property | Comments |
|----------|----------|
| Field Name | `ipAddress` |
| Type | `string` |

IP address of the interface inside Google Cloud Platform. Only IPv4 is
supported.

### `managementType` {#RouterBgpPeer.managementType}

| Property | Comments |
|----------|----------|
| Field Name | `managementType` |
| Type | `string` |

[Output Only] Type of how the resource/configuration of the BGP peer is
managed. MANAGED_BY_USER is the default value; MANAGED_BY_ATTACHMENT
represents an BGP peer that is automatically created for PARTNER
interconnectAttachment, Google will automatically create/delete this type of
BGP peer when the PARTNER interconnectAttachment is created/deleted.
Valid values:
    MANAGED_BY_ATTACHMENT
    MANAGED_BY_USER

### `name` {#RouterBgpPeer.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of this BGP peer. The name must be 1-63 characters long and comply with
RFC1035.

### `peerAsn` {#RouterBgpPeer.peerAsn}

| Property | Comments |
|----------|----------|
| Field Name | `peerAsn` |
| Type | `int32` |

Peer BGP Autonomous System Number (ASN). For VPN use case, this value can be
different for every tunnel.

### `peerIpAddress` {#RouterBgpPeer.peerIpAddress}

| Property | Comments |
|----------|----------|
| Field Name | `peerIpAddress` |
| Type | `string` |

IP address of the BGP interface outside Google cloud. Only IPv4 is
supported.

## Message `RouterInterface` {#RouterInterface}



### Inputs for `RouterInterface`

* `string` [`ipRange`](#RouterInterface.ipRange) = 1
* `string` [`linkedInterconnectAttachment`](#RouterInterface.linkedInterconnectAttachment) = 2
* `string` [`linkedVpnTunnel`](#RouterInterface.linkedVpnTunnel) = 3
* `string` [`managementType`](#RouterInterface.managementType) = 4
* `string` [`name`](#RouterInterface.name) = 5 (**Required**)

### `ipRange` {#RouterInterface.ipRange}

| Property | Comments |
|----------|----------|
| Field Name | `ipRange` |
| Type | `string` |

IP address and range of the interface. The IP range must be in the RFC3927
link-local IP space. The value must be a CIDR-formatted string, for example:
169.254.0.1/30. NOTE: Do not truncate the address as it represents the IP
address of the interface.

### `linkedInterconnectAttachment` {#RouterInterface.linkedInterconnectAttachment}

| Property | Comments |
|----------|----------|
| Field Name | `linkedInterconnectAttachment` |
| Type | `string` |

URI of the linked interconnect attachment. It must be in the same region as
the router. Each interface can have at most one linked resource and it could
either be a VPN Tunnel or an interconnect attachment.

### `linkedVpnTunnel` {#RouterInterface.linkedVpnTunnel}

| Property | Comments |
|----------|----------|
| Field Name | `linkedVpnTunnel` |
| Type | `string` |

URI of the linked VPN tunnel. It must be in the same region as the router.
Each interface can have at most one linked resource and it could either be a
VPN Tunnel or an interconnect attachment.

### `managementType` {#RouterInterface.managementType}

| Property | Comments |
|----------|----------|
| Field Name | `managementType` |
| Type | `string` |

[Output Only] Type of how the resource/configuration of the interface is
managed. MANAGED_BY_USER is the default value; MANAGED_BY_ATTACHMENT
represents an interface that is automatically created for PARTNER type
interconnectAttachment, Google will automatically create/update/delete this
type of interface when the PARTNER interconnectAttachment is
created/provisioned/deleted.
Valid values:
    MANAGED_BY_ATTACHMENT
    MANAGED_BY_USER

### `name` {#RouterInterface.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of this interface entry. The name must be 1-63 characters long and
comply with RFC1035.

## Message `RouterList` {#RouterList}

Contains a list of Router resources.


### Inputs for `RouterList`

* `string` [`id`](#RouterList.id) = 1
* `repeated` [`compute.Router`](gcp_compute.md#Router) [`items`](#RouterList.items) = 2
* `string` [`kind`](#RouterList.kind) = 3
* `string` [`nextPageToken`](#RouterList.nextPageToken) = 4
* `string` [`selfLink`](#RouterList.selfLink) = 5
* [`compute.RouterList.Warning`](gcp_compute.md#RouterList.Warning) [`warning`](#RouterList.warning) = 6

### `id` {#RouterList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#RouterList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Router`](gcp_compute.md#Router) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Router resources.

### `kind` {#RouterList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#router for routers.

### `nextPageToken` {#RouterList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#RouterList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#RouterList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RouterList.Warning`](gcp_compute.md#RouterList.Warning) |

## Message `RouterStatus` {#RouterStatus}



### Inputs for `RouterStatus`

* `repeated` [`compute.Route`](gcp_compute.md#Route) [`bestRoutes`](#RouterStatus.bestRoutes) = 1
* `repeated` [`compute.Route`](gcp_compute.md#Route) [`bestRoutesForRouter`](#RouterStatus.bestRoutesForRouter) = 2
* `repeated` [`compute.RouterStatusBgpPeerStatus`](gcp_compute.md#RouterStatusBgpPeerStatus) [`bgpPeerStatus`](#RouterStatus.bgpPeerStatus) = 3
* `string` [`network`](#RouterStatus.network) = 4

### `bestRoutes` {#RouterStatus.bestRoutes}

| Property | Comments |
|----------|----------|
| Field Name | `bestRoutes` |
| Type | [`compute.Route`](gcp_compute.md#Route) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Best routes for this router's network.

### `bestRoutesForRouter` {#RouterStatus.bestRoutesForRouter}

| Property | Comments |
|----------|----------|
| Field Name | `bestRoutesForRouter` |
| Type | [`compute.Route`](gcp_compute.md#Route) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Best routes learned by this router.

### `bgpPeerStatus` {#RouterStatus.bgpPeerStatus}

| Property | Comments |
|----------|----------|
| Field Name | `bgpPeerStatus` |
| Type | [`compute.RouterStatusBgpPeerStatus`](gcp_compute.md#RouterStatusBgpPeerStatus) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `network` {#RouterStatus.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

URI of the network to which this router belongs.

## Message `RouterStatusBgpPeerStatus` {#RouterStatusBgpPeerStatus}



### Inputs for `RouterStatusBgpPeerStatus`

* `repeated` [`compute.Route`](gcp_compute.md#Route) [`advertisedRoutes`](#RouterStatusBgpPeerStatus.advertisedRoutes) = 1
* `string` [`ipAddress`](#RouterStatusBgpPeerStatus.ipAddress) = 2
* `string` [`linkedVpnTunnel`](#RouterStatusBgpPeerStatus.linkedVpnTunnel) = 3
* `string` [`name`](#RouterStatusBgpPeerStatus.name) = 4 (**Required**)
* `int32` [`numLearnedRoutes`](#RouterStatusBgpPeerStatus.numLearnedRoutes) = 5
* `string` [`peerIpAddress`](#RouterStatusBgpPeerStatus.peerIpAddress) = 6
* `string` [`state`](#RouterStatusBgpPeerStatus.state) = 7
* `string` [`status`](#RouterStatusBgpPeerStatus.status) = 8
* `string` [`uptime`](#RouterStatusBgpPeerStatus.uptime) = 9
* `string` [`uptimeSeconds`](#RouterStatusBgpPeerStatus.uptimeSeconds) = 10

### `advertisedRoutes` {#RouterStatusBgpPeerStatus.advertisedRoutes}

| Property | Comments |
|----------|----------|
| Field Name | `advertisedRoutes` |
| Type | [`compute.Route`](gcp_compute.md#Route) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Routes that were advertised to the remote BGP peer

### `ipAddress` {#RouterStatusBgpPeerStatus.ipAddress}

| Property | Comments |
|----------|----------|
| Field Name | `ipAddress` |
| Type | `string` |

IP address of the local BGP interface.

### `linkedVpnTunnel` {#RouterStatusBgpPeerStatus.linkedVpnTunnel}

| Property | Comments |
|----------|----------|
| Field Name | `linkedVpnTunnel` |
| Type | `string` |

URL of the VPN tunnel that this BGP peer controls.

### `name` {#RouterStatusBgpPeerStatus.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of this BGP peer. Unique within the Routers resource.

### `numLearnedRoutes` {#RouterStatusBgpPeerStatus.numLearnedRoutes}

| Property | Comments |
|----------|----------|
| Field Name | `numLearnedRoutes` |
| Type | `int32` |

Number of routes learned from the remote BGP Peer.

### `peerIpAddress` {#RouterStatusBgpPeerStatus.peerIpAddress}

| Property | Comments |
|----------|----------|
| Field Name | `peerIpAddress` |
| Type | `string` |

IP address of the remote BGP interface.

### `state` {#RouterStatusBgpPeerStatus.state}

| Property | Comments |
|----------|----------|
| Field Name | `state` |
| Type | `string` |

BGP state as specified in RFC1771.

### `status` {#RouterStatusBgpPeerStatus.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

Status of the BGP peer: {UP, DOWN}
Valid values:
    DOWN
    UNKNOWN
    UP

### `uptime` {#RouterStatusBgpPeerStatus.uptime}

| Property | Comments |
|----------|----------|
| Field Name | `uptime` |
| Type | `string` |

Time this session has been up. Format: 14 years, 51 weeks, 6 days, 23 hours,
59 minutes, 59 seconds

### `uptimeSeconds` {#RouterStatusBgpPeerStatus.uptimeSeconds}

| Property | Comments |
|----------|----------|
| Field Name | `uptimeSeconds` |
| Type | `string` |

Time this session has been up, in seconds. Format: 145

## Message `RouterStatusResponse` {#RouterStatusResponse}



### Inputs for `RouterStatusResponse`

* `string` [`kind`](#RouterStatusResponse.kind) = 1
* [`compute.RouterStatus`](gcp_compute.md#RouterStatus) [`result`](#RouterStatusResponse.result) = 2

### `kind` {#RouterStatusResponse.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `result` {#RouterStatusResponse.result}

| Property | Comments |
|----------|----------|
| Field Name | `result` |
| Type | [`compute.RouterStatus`](gcp_compute.md#RouterStatus) |

## Message `RoutersPreviewResponse` {#RoutersPreviewResponse}



### Inputs for `RoutersPreviewResponse`

* [`compute.Router`](gcp_compute.md#Router) [`resource`](#RoutersPreviewResponse.resource) = 1

### `resource` {#RoutersPreviewResponse.resource}

| Property | Comments |
|----------|----------|
| Field Name | `resource` |
| Type | [`compute.Router`](gcp_compute.md#Router) |

Preview of given router.

## Message `RoutersScopedList` {#RoutersScopedList}



### Inputs for `RoutersScopedList`

* `repeated` [`compute.Router`](gcp_compute.md#Router) [`routers`](#RoutersScopedList.routers) = 1
* [`compute.RoutersScopedList.Warning`](gcp_compute.md#RoutersScopedList.Warning) [`warning`](#RoutersScopedList.warning) = 2

### `routers` {#RoutersScopedList.routers}

| Property | Comments |
|----------|----------|
| Field Name | `routers` |
| Type | [`compute.Router`](gcp_compute.md#Router) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of routers contained in this scope.

### `warning` {#RoutersScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.RoutersScopedList.Warning`](gcp_compute.md#RoutersScopedList.Warning) |

## Message `SSLHealthCheck` {#SSLHealthCheck}



### Inputs for `SSLHealthCheck`

* `int32` [`port`](#SSLHealthCheck.port) = 1
* `string` [`portName`](#SSLHealthCheck.portName) = 2
* `string` [`proxyHeader`](#SSLHealthCheck.proxyHeader) = 3
* `string` [`request`](#SSLHealthCheck.request) = 4
* `string` [`response`](#SSLHealthCheck.response) = 5

### `port` {#SSLHealthCheck.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The TCP port number for the health check request. The default value is 443.
Valid values are 1 through 65535.

### `portName` {#SSLHealthCheck.portName}

| Property | Comments |
|----------|----------|
| Field Name | `portName` |
| Type | `string` |

Port name as defined in InstanceGroup#NamedPort#name. If both port and
port_name are defined, port takes precedence.

### `proxyHeader` {#SSLHealthCheck.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

Specifies the type of proxy header to append before sending data to the
backend, either NONE or PROXY_V1. The default is NONE.
Valid values:
    NONE
    PROXY_V1

### `request` {#SSLHealthCheck.request}

| Property | Comments |
|----------|----------|
| Field Name | `request` |
| Type | `string` |

The application data to send once the SSL connection has been established
(default value is empty). If both request and response are empty, the
connection establishment alone will indicate health. The request data can
only be ASCII.

### `response` {#SSLHealthCheck.response}

| Property | Comments |
|----------|----------|
| Field Name | `response` |
| Type | `string` |

The bytes to match against the beginning of the response data. If left empty
(the default value), any response will indicate health. The response data
can only be ASCII.

## Message `Scheduling` {#Scheduling}

Sets the scheduling options for an Instance.


### Inputs for `Scheduling`

* `bool` [`automaticRestart`](#Scheduling.automaticRestart) = 1
* `string` [`onHostMaintenance`](#Scheduling.onHostMaintenance) = 2
* `bool` [`preemptible`](#Scheduling.preemptible) = 3

### `automaticRestart` {#Scheduling.automaticRestart}

| Property | Comments |
|----------|----------|
| Field Name | `automaticRestart` |
| Type | `bool` |

Specifies whether the instance should be automatically restarted if it is
terminated by Compute Engine (not terminated by a user). You can only set
the automatic restart option for standard instances. Preemptible instances
cannot be automatically restarted.

By default, this is set to true so an instance is automatically restarted if
it is terminated by Compute Engine.

### `onHostMaintenance` {#Scheduling.onHostMaintenance}

| Property | Comments |
|----------|----------|
| Field Name | `onHostMaintenance` |
| Type | `string` |

Defines the maintenance behavior for this instance. For standard instances,
the default behavior is MIGRATE. For preemptible instances, the default and
only possible behavior is TERMINATE. For more information, see Setting
Instance Scheduling Options.
Valid values:
    MIGRATE
    TERMINATE

### `preemptible` {#Scheduling.preemptible}

| Property | Comments |
|----------|----------|
| Field Name | `preemptible` |
| Type | `bool` |

Defines whether the instance is preemptible. This can only be set during
instance creation, it cannot be set or changed after the instance has been
created.

## Message `ScratchDisks` {#MachineType.ScratchDisks}

[Output Only] A list of extended scratch disks assigned to the instance.


### Inputs for `ScratchDisks`

* `int32` [`diskGb`](#MachineType.ScratchDisks.diskGb) = 1

### `diskGb` {#MachineType.ScratchDisks.diskGb}

| Property | Comments |
|----------|----------|
| Field Name | `diskGb` |
| Type | `int32` |

Size of the scratch disk, defined in GB.

## Message `SerialPortOutput` {#SerialPortOutput}

An instance's serial console output.


### Inputs for `SerialPortOutput`

* `string` [`contents`](#SerialPortOutput.contents) = 1
* `string` [`kind`](#SerialPortOutput.kind) = 2
* `string` [`next`](#SerialPortOutput.next) = 3
* `string` [`selfLink`](#SerialPortOutput.selfLink) = 4
* `string` [`start`](#SerialPortOutput.start) = 5

### `contents` {#SerialPortOutput.contents}

| Property | Comments |
|----------|----------|
| Field Name | `contents` |
| Type | `string` |

[Output Only] The contents of the console output.

### `kind` {#SerialPortOutput.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#serialPortOutput for
serial port output.

### `next` {#SerialPortOutput.next}

| Property | Comments |
|----------|----------|
| Field Name | `next` |
| Type | `string` |

[Output Only] The position of the next byte of content from the serial
console output. Use this value in the next request as the start parameter.

### `selfLink` {#SerialPortOutput.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `start` {#SerialPortOutput.start}

| Property | Comments |
|----------|----------|
| Field Name | `start` |
| Type | `string` |

The starting byte position of the output that was returned. This should
match the start parameter sent with the request. If the serial console
output exceeds the size of the buffer, older output will be overwritten by
newer content and the start values will be mismatched.

## Message `ServiceAccount` {#ServiceAccount}

A service account.


### Inputs for `ServiceAccount`

* `string` [`email`](#ServiceAccount.email) = 1
* `repeated` `string` [`scopes`](#ServiceAccount.scopes) = 2

### `email` {#ServiceAccount.email}

| Property | Comments |
|----------|----------|
| Field Name | `email` |
| Type | `string` |

Email address of the service account.

### `scopes` {#ServiceAccount.scopes}

| Property | Comments |
|----------|----------|
| Field Name | `scopes` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of scopes to be made available for this service account.

## Message `Snapshot` {#Snapshot}

A persistent disk snapshot resource. (== resource_for beta.snapshots ==) (==
resource_for v1.snapshots ==)


### Inputs for `Snapshot`

* `string` [`creationTimestamp`](#Snapshot.creationTimestamp) = 1
* `string` [`description`](#Snapshot.description) = 2
* `string` [`diskSizeGb`](#Snapshot.diskSizeGb) = 3
* `string` [`id`](#Snapshot.id) = 4
* `string` [`kind`](#Snapshot.kind) = 5
* `string` [`labelFingerprint`](#Snapshot.labelFingerprint) = 6
* `repeated` [`compute.Snapshot.LabelsEntry`](gcp_compute.md#Snapshot.LabelsEntry) [`labels`](#Snapshot.labels) = 7
* `repeated` `string` [`licenseCodes`](#Snapshot.licenseCodes) = 8
* `repeated` `string` [`licenses`](#Snapshot.licenses) = 9
* `string` [`name`](#Snapshot.name) = 10 (**Required**)
* `string` [`selfLink`](#Snapshot.selfLink) = 11
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`snapshotEncryptionKey`](#Snapshot.snapshotEncryptionKey) = 12
* `string` [`sourceDisk`](#Snapshot.sourceDisk) = 13
* [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) [`sourceDiskEncryptionKey`](#Snapshot.sourceDiskEncryptionKey) = 14
* `string` [`sourceDiskId`](#Snapshot.sourceDiskId) = 15
* `string` [`status`](#Snapshot.status) = 16
* `string` [`storageBytes`](#Snapshot.storageBytes) = 17
* `string` [`storageBytesStatus`](#Snapshot.storageBytesStatus) = 18

### `creationTimestamp` {#Snapshot.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Snapshot.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `diskSizeGb` {#Snapshot.diskSizeGb}

| Property | Comments |
|----------|----------|
| Field Name | `diskSizeGb` |
| Type | `string` |

[Output Only] Size of the snapshot, specified in GB.

### `id` {#Snapshot.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Snapshot.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#snapshot for Snapshot
resources.

### `labelFingerprint` {#Snapshot.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

A fingerprint for the labels being applied to this snapshot, which is
essentially a hash of the labels set used for optimistic locking. The
fingerprint is initially generated by Compute Engine and changes after every
request to modify or update labels. You must always provide an up-to-date
fingerprint hash in order to update or change labels.

To see the latest fingerprint, make a get() request to retrieve a snapshot.

### `labels` {#Snapshot.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.Snapshot.LabelsEntry`](gcp_compute.md#Snapshot.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Labels to apply to this snapshot. These can be later modified by the
setLabels method. Label values may be empty.

### `licenseCodes` {#Snapshot.licenseCodes}

| Property | Comments |
|----------|----------|
| Field Name | `licenseCodes` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Integer license codes indicating which licenses are attached
to this snapshot.

### `licenses` {#Snapshot.licenses}

| Property | Comments |
|----------|----------|
| Field Name | `licenses` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of public visible licenses that apply to this snapshot.
This can be because the original image had licenses attached (such as a
Windows image).

### `name` {#Snapshot.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource; provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `selfLink` {#Snapshot.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `snapshotEncryptionKey` {#Snapshot.snapshotEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `snapshotEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

Encrypts the snapshot using a customer-supplied encryption key.

After you encrypt a snapshot using a customer-supplied key, you must provide
the same key if you use the image later For example, you must provide the
encryption key when you create a disk from the encrypted snapshot in a
future request.

Customer-supplied encryption keys do not protect access to metadata of the
disk.

If you do not provide an encryption key when creating the snapshot, then the
snapshot will be encrypted using an automatically generated key and you do
not need to provide a key to use the snapshot later.

### `sourceDisk` {#Snapshot.sourceDisk}

| Property | Comments |
|----------|----------|
| Field Name | `sourceDisk` |
| Type | `string` |

[Output Only] The source disk used to create this snapshot.

### `sourceDiskEncryptionKey` {#Snapshot.sourceDiskEncryptionKey}

| Property | Comments |
|----------|----------|
| Field Name | `sourceDiskEncryptionKey` |
| Type | [`compute.CustomerEncryptionKey`](gcp_compute.md#CustomerEncryptionKey) |

The customer-supplied encryption key of the source disk. Required if the
source disk is protected by a customer-supplied encryption key.

### `sourceDiskId` {#Snapshot.sourceDiskId}

| Property | Comments |
|----------|----------|
| Field Name | `sourceDiskId` |
| Type | `string` |

[Output Only] The ID value of the disk used to create this snapshot. This
value may be used to determine whether the snapshot was taken from the
current or a previous instance of a given disk name.

### `status` {#Snapshot.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the snapshot. This can be CREATING, DELETING,
FAILED, READY, or UPLOADING.
Valid values:
    CREATING
    DELETING
    FAILED
    READY
    UPLOADING

### `storageBytes` {#Snapshot.storageBytes}

| Property | Comments |
|----------|----------|
| Field Name | `storageBytes` |
| Type | `string` |

[Output Only] A size of the storage used by the snapshot. As snapshots share
storage, this number is expected to change with snapshot creation/deletion.

### `storageBytesStatus` {#Snapshot.storageBytesStatus}

| Property | Comments |
|----------|----------|
| Field Name | `storageBytesStatus` |
| Type | `string` |

[Output Only] An indicator whether storageBytes is in a stable state or it
is being adjusted as a result of shared storage reallocation. This status
can either be UPDATING, meaning the size of the snapshot is being updated,
or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
Valid values:
    UPDATING
    UP_TO_DATE

## Message `SnapshotList` {#SnapshotList}

Contains a list of Snapshot resources.


### Inputs for `SnapshotList`

* `string` [`id`](#SnapshotList.id) = 1
* `repeated` [`compute.Snapshot`](gcp_compute.md#Snapshot) [`items`](#SnapshotList.items) = 2
* `string` [`kind`](#SnapshotList.kind) = 3
* `string` [`nextPageToken`](#SnapshotList.nextPageToken) = 4
* `string` [`selfLink`](#SnapshotList.selfLink) = 5
* [`compute.SnapshotList.Warning`](gcp_compute.md#SnapshotList.Warning) [`warning`](#SnapshotList.warning) = 6

### `id` {#SnapshotList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#SnapshotList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Snapshot`](gcp_compute.md#Snapshot) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Snapshot resources.

### `kind` {#SnapshotList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#SnapshotList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#SnapshotList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#SnapshotList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.SnapshotList.Warning`](gcp_compute.md#SnapshotList.Warning) |

## Message `SourceInstanceParams` {#SourceInstanceParams}

A specification of the parameters to use when creating the instance template
from a source instance.


### Inputs for `SourceInstanceParams`

* `repeated` [`compute.DiskInstantiationConfig`](gcp_compute.md#DiskInstantiationConfig) [`diskConfigs`](#SourceInstanceParams.diskConfigs) = 1

### `diskConfigs` {#SourceInstanceParams.diskConfigs}

| Property | Comments |
|----------|----------|
| Field Name | `diskConfigs` |
| Type | [`compute.DiskInstantiationConfig`](gcp_compute.md#DiskInstantiationConfig) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Attached disks configuration. If not provided, defaults are applied: For
boot disk and any other R/W disks, new custom images will be created from
each disk. For read-only disks, they will be attached in read-only mode.
Local SSD disks will be created as blank volumes.

## Message `SslCertificate` {#SslCertificate}

An SslCertificate resource. This resource provides a mechanism to upload an
SSL key and certificate to the load balancer to serve secure connections
from the user. (== resource_for beta.sslCertificates ==) (== resource_for
v1.sslCertificates ==)


### Inputs for `SslCertificate`

* `string` [`certificate`](#SslCertificate.certificate) = 1
* `string` [`creationTimestamp`](#SslCertificate.creationTimestamp) = 2
* `string` [`description`](#SslCertificate.description) = 3
* `string` [`id`](#SslCertificate.id) = 4
* `string` [`kind`](#SslCertificate.kind) = 5
* `string` [`name`](#SslCertificate.name) = 6 (**Required**)
* `string` [`privateKey`](#SslCertificate.privateKey) = 7
* `string` [`selfLink`](#SslCertificate.selfLink) = 8

### `certificate` {#SslCertificate.certificate}

| Property | Comments |
|----------|----------|
| Field Name | `certificate` |
| Type | `string` |

A local certificate file. The certificate must be in PEM format. The
certificate chain must be no greater than 5 certs long. The chain must
include at least one intermediate cert.

### `creationTimestamp` {#SslCertificate.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#SslCertificate.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#SslCertificate.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#SslCertificate.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#sslCertificate for SSL
certificates.

### `name` {#SslCertificate.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `privateKey` {#SslCertificate.privateKey}

| Property | Comments |
|----------|----------|
| Field Name | `privateKey` |
| Type | `string` |

A write-only private key in PEM format. Only insert requests will include
this field.

### `selfLink` {#SslCertificate.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output only] Server-defined URL for the resource.

## Message `SslCertificateList` {#SslCertificateList}

Contains a list of SslCertificate resources.


### Inputs for `SslCertificateList`

* `string` [`id`](#SslCertificateList.id) = 1
* `repeated` [`compute.SslCertificate`](gcp_compute.md#SslCertificate) [`items`](#SslCertificateList.items) = 2
* `string` [`kind`](#SslCertificateList.kind) = 3
* `string` [`nextPageToken`](#SslCertificateList.nextPageToken) = 4
* `string` [`selfLink`](#SslCertificateList.selfLink) = 5
* [`compute.SslCertificateList.Warning`](gcp_compute.md#SslCertificateList.Warning) [`warning`](#SslCertificateList.warning) = 6

### `id` {#SslCertificateList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#SslCertificateList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.SslCertificate`](gcp_compute.md#SslCertificate) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of SslCertificate resources.

### `kind` {#SslCertificateList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#SslCertificateList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#SslCertificateList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#SslCertificateList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.SslCertificateList.Warning`](gcp_compute.md#SslCertificateList.Warning) |

## Message `SslPoliciesList` {#SslPoliciesList}



### Inputs for `SslPoliciesList`

* `string` [`id`](#SslPoliciesList.id) = 1
* `repeated` [`compute.SslPolicy`](gcp_compute.md#SslPolicy) [`items`](#SslPoliciesList.items) = 2
* `string` [`kind`](#SslPoliciesList.kind) = 3
* `string` [`nextPageToken`](#SslPoliciesList.nextPageToken) = 4
* `string` [`selfLink`](#SslPoliciesList.selfLink) = 5
* [`compute.SslPoliciesList.Warning`](gcp_compute.md#SslPoliciesList.Warning) [`warning`](#SslPoliciesList.warning) = 6

### `id` {#SslPoliciesList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#SslPoliciesList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.SslPolicy`](gcp_compute.md#SslPolicy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of SslPolicy resources.

### `kind` {#SslPoliciesList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#sslPoliciesList for lists
of sslPolicies.

### `nextPageToken` {#SslPoliciesList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#SslPoliciesList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#SslPoliciesList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.SslPoliciesList.Warning`](gcp_compute.md#SslPoliciesList.Warning) |

## Message `SslPoliciesListAvailableFeaturesResponse` {#SslPoliciesListAvailableFeaturesResponse}



### Inputs for `SslPoliciesListAvailableFeaturesResponse`

* `repeated` `string` [`features`](#SslPoliciesListAvailableFeaturesResponse.features) = 1

### `features` {#SslPoliciesListAvailableFeaturesResponse.features}

| Property | Comments |
|----------|----------|
| Field Name | `features` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `SslPolicy` {#SslPolicy}

A SSL policy specifies the server-side support for SSL features. This can be
attached to a TargetHttpsProxy or a TargetSslProxy. This affects connections
between clients and the HTTPS or SSL proxy load balancer. They do not affect
the connection between the load balancers and the backends.


### Inputs for `SslPolicy`

* `string` [`creationTimestamp`](#SslPolicy.creationTimestamp) = 1
* `repeated` `string` [`customFeatures`](#SslPolicy.customFeatures) = 2
* `string` [`description`](#SslPolicy.description) = 3
* `repeated` `string` [`enabledFeatures`](#SslPolicy.enabledFeatures) = 4
* `string` [`fingerprint`](#SslPolicy.fingerprint) = 5
* `string` [`id`](#SslPolicy.id) = 6
* `string` [`kind`](#SslPolicy.kind) = 7
* `string` [`minTlsVersion`](#SslPolicy.minTlsVersion) = 8
* `string` [`name`](#SslPolicy.name) = 9 (**Required**)
* `string` [`profile`](#SslPolicy.profile) = 10
* `string` [`selfLink`](#SslPolicy.selfLink) = 11
* `repeated` [`compute.SslPolicy.Warnings`](gcp_compute.md#SslPolicy.Warnings) [`warnings`](#SslPolicy.warnings) = 12

### `creationTimestamp` {#SslPolicy.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `customFeatures` {#SslPolicy.customFeatures}

| Property | Comments |
|----------|----------|
| Field Name | `customFeatures` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of features enabled when the selected profile is CUSTOM. The
- method returns the set of features that can be specified in this list.
This field must be empty if the profile is not CUSTOM.

### `description` {#SslPolicy.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `enabledFeatures` {#SslPolicy.enabledFeatures}

| Property | Comments |
|----------|----------|
| Field Name | `enabledFeatures` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] The list of features enabled in the SSL policy.

### `fingerprint` {#SslPolicy.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint of this resource. A hash of the contents stored in this object.
This field is used in optimistic locking. This field will be ignored when
inserting a SslPolicy. An up-to-date fingerprint must be provided in order
to update the SslPolicy.

### `id` {#SslPolicy.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#SslPolicy.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output only] Type of the resource. Always compute#sslPolicyfor SSL
policies.

### `minTlsVersion` {#SslPolicy.minTlsVersion}

| Property | Comments |
|----------|----------|
| Field Name | `minTlsVersion` |
| Type | `string` |

The minimum version of SSL protocol that can be used by the clients to
establish a connection with the load balancer. This can be one of TLS_1_0,
TLS_1_1, TLS_1_2.
Valid values:
    TLS_1_0
    TLS_1_1
    TLS_1_2

### `name` {#SslPolicy.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match the
regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first
character must be a lowercase letter, and all following characters must be a
dash, lowercase letter, or digit, except the last character, which cannot be
a dash.

### `profile` {#SslPolicy.profile}

| Property | Comments |
|----------|----------|
| Field Name | `profile` |
| Type | `string` |

Profile specifies the set of SSL features that can be used by the load
balancer when negotiating SSL with clients. This can be one of COMPATIBLE,
MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to
enable must be specified in the customFeatures field.
Valid values:
    COMPATIBLE
    CUSTOM
    MODERN
    RESTRICTED

### `selfLink` {#SslPolicy.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `warnings` {#SslPolicy.warnings}

| Property | Comments |
|----------|----------|
| Field Name | `warnings` |
| Type | [`compute.SslPolicy.Warnings`](gcp_compute.md#SslPolicy.Warnings) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `SslPolicyReference` {#SslPolicyReference}



### Inputs for `SslPolicyReference`

* `string` [`sslPolicy`](#SslPolicyReference.sslPolicy) = 1

### `sslPolicy` {#SslPolicyReference.sslPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `sslPolicy` |
| Type | `string` |

URL of the SSL policy resource. Set this to empty string to clear any
existing SSL policy associated with the target proxy resource.

## Message `Subnetwork` {#Subnetwork}

A Subnetwork resource. (== resource_for beta.subnetworks ==) (==
resource_for v1.subnetworks ==)


### Inputs for `Subnetwork`

* `string` [`creationTimestamp`](#Subnetwork.creationTimestamp) = 1
* `string` [`description`](#Subnetwork.description) = 2
* `bool` [`enableFlowLogs`](#Subnetwork.enableFlowLogs) = 3
* `string` [`fingerprint`](#Subnetwork.fingerprint) = 4
* `string` [`gatewayAddress`](#Subnetwork.gatewayAddress) = 5
* `string` [`id`](#Subnetwork.id) = 6
* `string` [`ipCidrRange`](#Subnetwork.ipCidrRange) = 7
* `string` [`kind`](#Subnetwork.kind) = 8
* `string` [`name`](#Subnetwork.name) = 9 (**Required**)
* `string` [`network`](#Subnetwork.network) = 10
* `bool` [`privateIpGoogleAccess`](#Subnetwork.privateIpGoogleAccess) = 11
* `string` [`region`](#Subnetwork.region) = 12
* `repeated` [`compute.SubnetworkSecondaryRange`](gcp_compute.md#SubnetworkSecondaryRange) [`secondaryIpRanges`](#Subnetwork.secondaryIpRanges) = 13
* `string` [`selfLink`](#Subnetwork.selfLink) = 14

### `creationTimestamp` {#Subnetwork.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#Subnetwork.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource. This field can be set only at resource creation time.

### `enableFlowLogs` {#Subnetwork.enableFlowLogs}

| Property | Comments |
|----------|----------|
| Field Name | `enableFlowLogs` |
| Type | `bool` |

Whether to enable flow logging for this subnetwork.

### `fingerprint` {#Subnetwork.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint of this resource. A hash of the contents stored in this object.
This field is used in optimistic locking. This field will be ignored when
inserting a Subnetwork. An up-to-date fingerprint must be provided in order
to update the Subnetwork.

### `gatewayAddress` {#Subnetwork.gatewayAddress}

| Property | Comments |
|----------|----------|
| Field Name | `gatewayAddress` |
| Type | `string` |

[Output Only] The gateway address for default routes to reach destination
addresses outside this subnetwork.

### `id` {#Subnetwork.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `ipCidrRange` {#Subnetwork.ipCidrRange}

| Property | Comments |
|----------|----------|
| Field Name | `ipCidrRange` |
| Type | `string` |

The range of internal addresses that are owned by this subnetwork. Provide
this property when you create the subnetwork. For example, 10.0.0.0/8 or
192.168.0.0/16. Ranges must be unique and non-overlapping within a network.
Only IPv4 is supported. This field can be set only at resource creation
time.

### `kind` {#Subnetwork.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#subnetwork for Subnetwork
resources.

### `name` {#Subnetwork.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name of the resource, provided by the client when initially creating the
resource. The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `network` {#Subnetwork.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

The URL of the network to which this subnetwork belongs, provided by the
client when initially creating the subnetwork. Only networks that are in the
distributed mode can have subnetworks. This field can be set only at
resource creation time.

### `privateIpGoogleAccess` {#Subnetwork.privateIpGoogleAccess}

| Property | Comments |
|----------|----------|
| Field Name | `privateIpGoogleAccess` |
| Type | `bool` |

Whether the VMs in this subnet can access Google services without assigned
external IP addresses. This field can be both set at resource creation time
and updated using setPrivateIpGoogleAccess.

### `region` {#Subnetwork.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

URL of the region where the Subnetwork resides. This field can be set only
at resource creation time.

### `secondaryIpRanges` {#Subnetwork.secondaryIpRanges}

| Property | Comments |
|----------|----------|
| Field Name | `secondaryIpRanges` |
| Type | [`compute.SubnetworkSecondaryRange`](gcp_compute.md#SubnetworkSecondaryRange) |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of configurations for secondary IP ranges for VM instances
contained in this subnetwork. The primary IP of such VM must belong to the
primary ipCidrRange of the subnetwork. The alias IPs may belong to either
primary or secondary ranges.

### `selfLink` {#Subnetwork.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

## Message `SubnetworkAggregatedList` {#SubnetworkAggregatedList}



### Inputs for `SubnetworkAggregatedList`

* `string` [`id`](#SubnetworkAggregatedList.id) = 1
* `repeated` [`compute.SubnetworkAggregatedList.ItemsEntry`](gcp_compute.md#SubnetworkAggregatedList.ItemsEntry) [`items`](#SubnetworkAggregatedList.items) = 2
* `string` [`kind`](#SubnetworkAggregatedList.kind) = 3
* `string` [`nextPageToken`](#SubnetworkAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#SubnetworkAggregatedList.selfLink) = 5
* [`compute.SubnetworkAggregatedList.Warning`](gcp_compute.md#SubnetworkAggregatedList.Warning) [`warning`](#SubnetworkAggregatedList.warning) = 6

### `id` {#SubnetworkAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#SubnetworkAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.SubnetworkAggregatedList.ItemsEntry`](gcp_compute.md#SubnetworkAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of SubnetworksScopedList resources.

### `kind` {#SubnetworkAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#subnetworkAggregatedList for
aggregated lists of subnetworks.

### `nextPageToken` {#SubnetworkAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#SubnetworkAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#SubnetworkAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.SubnetworkAggregatedList.Warning`](gcp_compute.md#SubnetworkAggregatedList.Warning) |

## Message `SubnetworkList` {#SubnetworkList}

Contains a list of Subnetwork resources.


### Inputs for `SubnetworkList`

* `string` [`id`](#SubnetworkList.id) = 1
* `repeated` [`compute.Subnetwork`](gcp_compute.md#Subnetwork) [`items`](#SubnetworkList.items) = 2
* `string` [`kind`](#SubnetworkList.kind) = 3
* `string` [`nextPageToken`](#SubnetworkList.nextPageToken) = 4
* `string` [`selfLink`](#SubnetworkList.selfLink) = 5
* [`compute.SubnetworkList.Warning`](gcp_compute.md#SubnetworkList.Warning) [`warning`](#SubnetworkList.warning) = 6

### `id` {#SubnetworkList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#SubnetworkList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Subnetwork`](gcp_compute.md#Subnetwork) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Subnetwork resources.

### `kind` {#SubnetworkList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#subnetworkList for lists of
subnetworks.

### `nextPageToken` {#SubnetworkList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#SubnetworkList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#SubnetworkList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.SubnetworkList.Warning`](gcp_compute.md#SubnetworkList.Warning) |

## Message `SubnetworkSecondaryRange` {#SubnetworkSecondaryRange}

Represents a secondary IP range of a subnetwork.


### Inputs for `SubnetworkSecondaryRange`

* `string` [`ipCidrRange`](#SubnetworkSecondaryRange.ipCidrRange) = 1
* `string` [`rangeName`](#SubnetworkSecondaryRange.rangeName) = 2

### `ipCidrRange` {#SubnetworkSecondaryRange.ipCidrRange}

| Property | Comments |
|----------|----------|
| Field Name | `ipCidrRange` |
| Type | `string` |

The range of IP addresses belonging to this subnetwork secondary range.
Provide this property when you create the subnetwork. Ranges must be unique
and non-overlapping with all primary and secondary IP ranges within a
network. Only IPv4 is supported.

### `rangeName` {#SubnetworkSecondaryRange.rangeName}

| Property | Comments |
|----------|----------|
| Field Name | `rangeName` |
| Type | `string` |

The name associated with this subnetwork secondary range, used when adding
an alias IP range to a VM instance. The name must be 1-63 characters long,
and comply with RFC1035. The name must be unique within the subnetwork.

## Message `SubnetworksExpandIpCidrRangeRequest` {#SubnetworksExpandIpCidrRangeRequest}



### Inputs for `SubnetworksExpandIpCidrRangeRequest`

* `string` [`ipCidrRange`](#SubnetworksExpandIpCidrRangeRequest.ipCidrRange) = 1

### `ipCidrRange` {#SubnetworksExpandIpCidrRangeRequest.ipCidrRange}

| Property | Comments |
|----------|----------|
| Field Name | `ipCidrRange` |
| Type | `string` |

The IP (in CIDR format or netmask) of internal addresses that are legal on
this Subnetwork. This range should be disjoint from other subnetworks within
this network. This range can only be larger than (i.e. a superset of) the
range previously defined before the update.

## Message `SubnetworksScopedList` {#SubnetworksScopedList}



### Inputs for `SubnetworksScopedList`

* `repeated` [`compute.Subnetwork`](gcp_compute.md#Subnetwork) [`subnetworks`](#SubnetworksScopedList.subnetworks) = 1
* [`compute.SubnetworksScopedList.Warning`](gcp_compute.md#SubnetworksScopedList.Warning) [`warning`](#SubnetworksScopedList.warning) = 2

### `subnetworks` {#SubnetworksScopedList.subnetworks}

| Property | Comments |
|----------|----------|
| Field Name | `subnetworks` |
| Type | [`compute.Subnetwork`](gcp_compute.md#Subnetwork) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of subnetworks contained in this scope.

### `warning` {#SubnetworksScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.SubnetworksScopedList.Warning`](gcp_compute.md#SubnetworksScopedList.Warning) |

## Message `SubnetworksSetPrivateIpGoogleAccessRequest` {#SubnetworksSetPrivateIpGoogleAccessRequest}



### Inputs for `SubnetworksSetPrivateIpGoogleAccessRequest`

* `bool` [`privateIpGoogleAccess`](#SubnetworksSetPrivateIpGoogleAccessRequest.privateIpGoogleAccess) = 1

### `privateIpGoogleAccess` {#SubnetworksSetPrivateIpGoogleAccessRequest.privateIpGoogleAccess}

| Property | Comments |
|----------|----------|
| Field Name | `privateIpGoogleAccess` |
| Type | `bool` |

## Message `TCPHealthCheck` {#TCPHealthCheck}



### Inputs for `TCPHealthCheck`

* `int32` [`port`](#TCPHealthCheck.port) = 1
* `string` [`portName`](#TCPHealthCheck.portName) = 2
* `string` [`proxyHeader`](#TCPHealthCheck.proxyHeader) = 3
* `string` [`request`](#TCPHealthCheck.request) = 4
* `string` [`response`](#TCPHealthCheck.response) = 5

### `port` {#TCPHealthCheck.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `int32` |

The TCP port number for the health check request. The default value is 80.
Valid values are 1 through 65535.

### `portName` {#TCPHealthCheck.portName}

| Property | Comments |
|----------|----------|
| Field Name | `portName` |
| Type | `string` |

Port name as defined in InstanceGroup#NamedPort#name. If both port and
port_name are defined, port takes precedence.

### `proxyHeader` {#TCPHealthCheck.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

Specifies the type of proxy header to append before sending data to the
backend, either NONE or PROXY_V1. The default is NONE.
Valid values:
    NONE
    PROXY_V1

### `request` {#TCPHealthCheck.request}

| Property | Comments |
|----------|----------|
| Field Name | `request` |
| Type | `string` |

The application data to send once the TCP connection has been established
(default value is empty). If both request and response are empty, the
connection establishment alone will indicate health. The request data can
only be ASCII.

### `response` {#TCPHealthCheck.response}

| Property | Comments |
|----------|----------|
| Field Name | `response` |
| Type | `string` |

The bytes to match against the beginning of the response data. If left empty
(the default value), any response will indicate health. The response data
can only be ASCII.

## Message `Tags` {#Tags}

A set of instance tags.


### Inputs for `Tags`

* `string` [`fingerprint`](#Tags.fingerprint) = 1
* `repeated` `string` [`items`](#Tags.items) = 2

### `fingerprint` {#Tags.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Specifies a fingerprint for this request, which is essentially a hash of the
tags' contents and used for optimistic locking. The fingerprint is initially
generated by Compute Engine and changes after every request to modify or
update tags. You must always provide an up-to-date fingerprint hash in order
to update or change tags.

To see the latest fingerprint, make get() request to the instance.

### `items` {#Tags.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

An array of tags. Each tag must be 1-63 characters long, and comply with
RFC1035.

## Message `TargetHttpProxy` {#TargetHttpProxy}

A TargetHttpProxy resource. This resource defines an HTTP proxy. (==
resource_for beta.targetHttpProxies ==) (== resource_for
v1.targetHttpProxies ==)


### Inputs for `TargetHttpProxy`

* `string` [`creationTimestamp`](#TargetHttpProxy.creationTimestamp) = 1
* `string` [`description`](#TargetHttpProxy.description) = 2
* `string` [`id`](#TargetHttpProxy.id) = 3
* `string` [`kind`](#TargetHttpProxy.kind) = 4
* `string` [`name`](#TargetHttpProxy.name) = 5 (**Required**)
* `string` [`selfLink`](#TargetHttpProxy.selfLink) = 6
* `string` [`urlMap`](#TargetHttpProxy.urlMap) = 7

### `creationTimestamp` {#TargetHttpProxy.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetHttpProxy.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#TargetHttpProxy.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#TargetHttpProxy.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetHttpProxy for target
HTTP proxies.

### `name` {#TargetHttpProxy.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `selfLink` {#TargetHttpProxy.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `urlMap` {#TargetHttpProxy.urlMap}

| Property | Comments |
|----------|----------|
| Field Name | `urlMap` |
| Type | `string` |

URL to the UrlMap resource that defines the mapping from URL to the
BackendService.

## Message `TargetHttpProxyList` {#TargetHttpProxyList}

A list of TargetHttpProxy resources.


### Inputs for `TargetHttpProxyList`

* `string` [`id`](#TargetHttpProxyList.id) = 1
* `repeated` [`compute.TargetHttpProxy`](gcp_compute.md#TargetHttpProxy) [`items`](#TargetHttpProxyList.items) = 2
* `string` [`kind`](#TargetHttpProxyList.kind) = 3
* `string` [`nextPageToken`](#TargetHttpProxyList.nextPageToken) = 4
* `string` [`selfLink`](#TargetHttpProxyList.selfLink) = 5
* [`compute.TargetHttpProxyList.Warning`](gcp_compute.md#TargetHttpProxyList.Warning) [`warning`](#TargetHttpProxyList.warning) = 6

### `id` {#TargetHttpProxyList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetHttpProxyList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetHttpProxy`](gcp_compute.md#TargetHttpProxy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetHttpProxy resources.

### `kind` {#TargetHttpProxyList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource. Always compute#targetHttpProxyList for lists of target
HTTP proxies.

### `nextPageToken` {#TargetHttpProxyList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetHttpProxyList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetHttpProxyList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetHttpProxyList.Warning`](gcp_compute.md#TargetHttpProxyList.Warning) |

## Message `TargetHttpsProxiesSetQuicOverrideRequest` {#TargetHttpsProxiesSetQuicOverrideRequest}



### Inputs for `TargetHttpsProxiesSetQuicOverrideRequest`

* `string` [`quicOverride`](#TargetHttpsProxiesSetQuicOverrideRequest.quicOverride) = 1

### `quicOverride` {#TargetHttpsProxiesSetQuicOverrideRequest.quicOverride}

| Property | Comments |
|----------|----------|
| Field Name | `quicOverride` |
| Type | `string` |

QUIC policy for the TargetHttpsProxy resource.
Valid values:
    DISABLE
    ENABLE
    NONE

## Message `TargetHttpsProxiesSetSslCertificatesRequest` {#TargetHttpsProxiesSetSslCertificatesRequest}



### Inputs for `TargetHttpsProxiesSetSslCertificatesRequest`

* `repeated` `string` [`sslCertificates`](#TargetHttpsProxiesSetSslCertificatesRequest.sslCertificates) = 1

### `sslCertificates` {#TargetHttpsProxiesSetSslCertificatesRequest.sslCertificates}

| Property | Comments |
|----------|----------|
| Field Name | `sslCertificates` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

New set of SslCertificate resources to associate with this TargetHttpsProxy
resource. Currently exactly one SslCertificate resource must be specified.

## Message `TargetHttpsProxy` {#TargetHttpsProxy}

A TargetHttpsProxy resource. This resource defines an HTTPS proxy. (==
resource_for beta.targetHttpsProxies ==) (== resource_for
v1.targetHttpsProxies ==)


### Inputs for `TargetHttpsProxy`

* `string` [`creationTimestamp`](#TargetHttpsProxy.creationTimestamp) = 1
* `string` [`description`](#TargetHttpsProxy.description) = 2
* `string` [`id`](#TargetHttpsProxy.id) = 3
* `string` [`kind`](#TargetHttpsProxy.kind) = 4
* `string` [`name`](#TargetHttpsProxy.name) = 5 (**Required**)
* `string` [`quicOverride`](#TargetHttpsProxy.quicOverride) = 6
* `string` [`selfLink`](#TargetHttpsProxy.selfLink) = 7
* `repeated` `string` [`sslCertificates`](#TargetHttpsProxy.sslCertificates) = 8
* `string` [`sslPolicy`](#TargetHttpsProxy.sslPolicy) = 9
* `string` [`urlMap`](#TargetHttpsProxy.urlMap) = 10

### `creationTimestamp` {#TargetHttpsProxy.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetHttpsProxy.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#TargetHttpsProxy.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#TargetHttpsProxy.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetHttpsProxy for target
HTTPS proxies.

### `name` {#TargetHttpsProxy.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `quicOverride` {#TargetHttpsProxy.quicOverride}

| Property | Comments |
|----------|----------|
| Field Name | `quicOverride` |
| Type | `string` |

Specifies the QUIC override policy for this TargetHttpsProxy resource. This
determines whether the load balancer will attempt to negotiate QUIC with
clients or not. Can specify one of NONE, ENABLE, or DISABLE. Specify ENABLE
to always enable QUIC, Enables QUIC when set to ENABLE, and disables QUIC
when set to DISABLE. If NONE is specified, uses the QUIC policy with no user
overrides, which is equivalent to DISABLE. Not specifying this field is
equivalent to specifying NONE.
Valid values:
    DISABLE
    ENABLE
    NONE

### `selfLink` {#TargetHttpsProxy.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sslCertificates` {#TargetHttpsProxy.sslCertificates}

| Property | Comments |
|----------|----------|
| Field Name | `sslCertificates` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

URLs to SslCertificate resources that are used to authenticate connections
between users and the load balancer. Currently, exactly one SSL certificate
must be specified.

### `sslPolicy` {#TargetHttpsProxy.sslPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `sslPolicy` |
| Type | `string` |

URL of SslPolicy resource that will be associated with the TargetHttpsProxy
resource. If not set, the TargetHttpsProxy resource will not have any SSL
policy configured.

### `urlMap` {#TargetHttpsProxy.urlMap}

| Property | Comments |
|----------|----------|
| Field Name | `urlMap` |
| Type | `string` |

A fully-qualified or valid partial URL to the UrlMap resource that defines
the mapping from URL to the BackendService. For example, the following are
all valid URLs for specifying a URL map:
- https://www.googleapis.compute/v1/projects/project/global/urlMaps/url-map
- projects/project/global/urlMaps/url-map
- global/urlMaps/url-map

## Message `TargetHttpsProxyList` {#TargetHttpsProxyList}

Contains a list of TargetHttpsProxy resources.


### Inputs for `TargetHttpsProxyList`

* `string` [`id`](#TargetHttpsProxyList.id) = 1
* `repeated` [`compute.TargetHttpsProxy`](gcp_compute.md#TargetHttpsProxy) [`items`](#TargetHttpsProxyList.items) = 2
* `string` [`kind`](#TargetHttpsProxyList.kind) = 3
* `string` [`nextPageToken`](#TargetHttpsProxyList.nextPageToken) = 4
* `string` [`selfLink`](#TargetHttpsProxyList.selfLink) = 5
* [`compute.TargetHttpsProxyList.Warning`](gcp_compute.md#TargetHttpsProxyList.Warning) [`warning`](#TargetHttpsProxyList.warning) = 6

### `id` {#TargetHttpsProxyList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetHttpsProxyList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetHttpsProxy`](gcp_compute.md#TargetHttpsProxy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetHttpsProxy resources.

### `kind` {#TargetHttpsProxyList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource. Always compute#targetHttpsProxyList for lists of target
HTTPS proxies.

### `nextPageToken` {#TargetHttpsProxyList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetHttpsProxyList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetHttpsProxyList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetHttpsProxyList.Warning`](gcp_compute.md#TargetHttpsProxyList.Warning) |

## Message `TargetInstance` {#TargetInstance}

A TargetInstance resource. This resource defines an endpoint instance that
terminates traffic of certain protocols. (== resource_for
beta.targetInstances ==) (== resource_for v1.targetInstances ==)


### Inputs for `TargetInstance`

* `string` [`creationTimestamp`](#TargetInstance.creationTimestamp) = 1
* `string` [`description`](#TargetInstance.description) = 2
* `string` [`id`](#TargetInstance.id) = 3
* `string` [`instance`](#TargetInstance.instance) = 4
* `string` [`kind`](#TargetInstance.kind) = 5
* `string` [`name`](#TargetInstance.name) = 6 (**Required**)
* `string` [`natPolicy`](#TargetInstance.natPolicy) = 7
* `string` [`selfLink`](#TargetInstance.selfLink) = 8
* `string` [`zone`](#TargetInstance.zone) = 9

### `creationTimestamp` {#TargetInstance.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetInstance.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#TargetInstance.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `instance` {#TargetInstance.instance}

| Property | Comments |
|----------|----------|
| Field Name | `instance` |
| Type | `string` |

A URL to the virtual machine instance that handles traffic for this target
instance. When creating a target instance, you can provide the
fully-qualified URL or a valid partial URL to the desired virtual machine.
For example, the following are all valid URLs:
-
https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
- projects/project/zones/zone/instances/instance
- zones/zone/instances/instance

### `kind` {#TargetInstance.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] The type of the resource. Always compute#targetInstance for
target instances.

### `name` {#TargetInstance.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `natPolicy` {#TargetInstance.natPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `natPolicy` |
| Type | `string` |

NAT option controlling how IPs are NAT'ed to the instance. Currently only
NO_NAT (default value) is supported.
Valid values:
    NO_NAT

### `selfLink` {#TargetInstance.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `zone` {#TargetInstance.zone}

| Property | Comments |
|----------|----------|
| Field Name | `zone` |
| Type | `string` |

[Output Only] URL of the zone where the target instance resides. You must
specify this field as part of the HTTP request URL. It is not settable as a
field in the request body.

## Message `TargetInstanceAggregatedList` {#TargetInstanceAggregatedList}



### Inputs for `TargetInstanceAggregatedList`

* `string` [`id`](#TargetInstanceAggregatedList.id) = 1
* `repeated` [`compute.TargetInstanceAggregatedList.ItemsEntry`](gcp_compute.md#TargetInstanceAggregatedList.ItemsEntry) [`items`](#TargetInstanceAggregatedList.items) = 2
* `string` [`kind`](#TargetInstanceAggregatedList.kind) = 3
* `string` [`nextPageToken`](#TargetInstanceAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#TargetInstanceAggregatedList.selfLink) = 5
* [`compute.TargetInstanceAggregatedList.Warning`](gcp_compute.md#TargetInstanceAggregatedList.Warning) [`warning`](#TargetInstanceAggregatedList.warning) = 6

### `id` {#TargetInstanceAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetInstanceAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetInstanceAggregatedList.ItemsEntry`](gcp_compute.md#TargetInstanceAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetInstance resources.

### `kind` {#TargetInstanceAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#TargetInstanceAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetInstanceAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetInstanceAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetInstanceAggregatedList.Warning`](gcp_compute.md#TargetInstanceAggregatedList.Warning) |

## Message `TargetInstanceList` {#TargetInstanceList}

Contains a list of TargetInstance resources.


### Inputs for `TargetInstanceList`

* `string` [`id`](#TargetInstanceList.id) = 1
* `repeated` [`compute.TargetInstance`](gcp_compute.md#TargetInstance) [`items`](#TargetInstanceList.items) = 2
* `string` [`kind`](#TargetInstanceList.kind) = 3
* `string` [`nextPageToken`](#TargetInstanceList.nextPageToken) = 4
* `string` [`selfLink`](#TargetInstanceList.selfLink) = 5
* [`compute.TargetInstanceList.Warning`](gcp_compute.md#TargetInstanceList.Warning) [`warning`](#TargetInstanceList.warning) = 6

### `id` {#TargetInstanceList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetInstanceList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetInstance`](gcp_compute.md#TargetInstance) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetInstance resources.

### `kind` {#TargetInstanceList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#TargetInstanceList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetInstanceList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetInstanceList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetInstanceList.Warning`](gcp_compute.md#TargetInstanceList.Warning) |

## Message `TargetInstancesScopedList` {#TargetInstancesScopedList}



### Inputs for `TargetInstancesScopedList`

* `repeated` [`compute.TargetInstance`](gcp_compute.md#TargetInstance) [`targetInstances`](#TargetInstancesScopedList.targetInstances) = 1
* [`compute.TargetInstancesScopedList.Warning`](gcp_compute.md#TargetInstancesScopedList.Warning) [`warning`](#TargetInstancesScopedList.warning) = 2

### `targetInstances` {#TargetInstancesScopedList.targetInstances}

| Property | Comments |
|----------|----------|
| Field Name | `targetInstances` |
| Type | [`compute.TargetInstance`](gcp_compute.md#TargetInstance) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of target instances contained in this scope.

### `warning` {#TargetInstancesScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetInstancesScopedList.Warning`](gcp_compute.md#TargetInstancesScopedList.Warning) |

## Message `TargetPool` {#TargetPool}

A TargetPool resource. This resource defines a pool of instances, an
associated HttpHealthCheck resource, and the fallback target pool. (==
resource_for beta.targetPools ==) (== resource_for v1.targetPools ==)


### Inputs for `TargetPool`

* `string` [`backupPool`](#TargetPool.backupPool) = 1
* `string` [`creationTimestamp`](#TargetPool.creationTimestamp) = 2
* `string` [`description`](#TargetPool.description) = 3
* `TYPE_DOUBLE` [`failoverRatio`](#TargetPool.failoverRatio) = 4
* `repeated` `string` [`healthChecks`](#TargetPool.healthChecks) = 5
* `string` [`id`](#TargetPool.id) = 6
* `repeated` `string` [`instances`](#TargetPool.instances) = 7
* `string` [`kind`](#TargetPool.kind) = 8
* `string` [`name`](#TargetPool.name) = 9 (**Required**)
* `string` [`region`](#TargetPool.region) = 10
* `string` [`selfLink`](#TargetPool.selfLink) = 11
* `string` [`sessionAffinity`](#TargetPool.sessionAffinity) = 12

### `backupPool` {#TargetPool.backupPool}

| Property | Comments |
|----------|----------|
| Field Name | `backupPool` |
| Type | `string` |

This field is applicable only when the containing target pool is serving a
forwarding rule as the primary pool, and its failoverRatio field is properly
set to a value between [0, 1].

backupPool and failoverRatio together define the fallback behavior of the
primary target pool: if the ratio of the healthy instances in the primary
pool is at or below failoverRatio, traffic arriving at the load-balanced IP
will be directed to the backup pool.

In case where failoverRatio and backupPool are not set, or all the instances
in the backup pool are unhealthy, the traffic will be directed back to the
primary pool in the "force" mode, where traffic will be spread to the
healthy instances with the best effort, or to all instances when no instance
is healthy.

### `creationTimestamp` {#TargetPool.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetPool.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `failoverRatio` {#TargetPool.failoverRatio}

| Property | Comments |
|----------|----------|
| Field Name | `failoverRatio` |
| Type | `TYPE_DOUBLE` |

This field is applicable only when the containing target pool is serving a
forwarding rule as the primary pool (i.e., not as a backup pool to some
other target pool). The value of the field must be in [0, 1].

If set, backupPool must also be set. They together define the fallback
behavior of the primary target pool: if the ratio of the healthy instances
in the primary pool is at or below this number, traffic arriving at the
load-balanced IP will be directed to the backup pool.

In case where failoverRatio is not set or all the instances in the backup
pool are unhealthy, the traffic will be directed back to the primary pool in
the "force" mode, where traffic will be spread to the healthy instances with
the best effort, or to all instances when no instance is healthy.

### `healthChecks` {#TargetPool.healthChecks}

| Property | Comments |
|----------|----------|
| Field Name | `healthChecks` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The URL of the HttpHealthCheck resource. A member instance in this pool is
considered healthy if and only if the health checks pass. An empty list
means all member instances will be considered healthy at all times. Only
HttpHealthChecks are supported. Only one health check may be specified.

### `id` {#TargetPool.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `instances` {#TargetPool.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of resource URLs to the virtual machine instances serving this pool.
They must live in zones contained in the same region as this pool.

### `kind` {#TargetPool.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#targetPool for target
pools.

### `name` {#TargetPool.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `region` {#TargetPool.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the target pool resides.

### `selfLink` {#TargetPool.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sessionAffinity` {#TargetPool.sessionAffinity}

| Property | Comments |
|----------|----------|
| Field Name | `sessionAffinity` |
| Type | `string` |

Sesssion affinity option, must be one of the following values:
NONE: Connections from the same client IP may go to any instance in the
pool.
CLIENT_IP: Connections from the same client IP will go to the same instance
in the pool while that instance remains healthy.
CLIENT_IP_PROTO: Connections from the same client IP with the same IP
protocol will go to the same instance in the pool while that instance
remains healthy.
Valid values:
    CLIENT_IP
    CLIENT_IP_PORT_PROTO
    CLIENT_IP_PROTO
    GENERATED_COOKIE
    NONE

## Message `TargetPoolAggregatedList` {#TargetPoolAggregatedList}



### Inputs for `TargetPoolAggregatedList`

* `string` [`id`](#TargetPoolAggregatedList.id) = 1
* `repeated` [`compute.TargetPoolAggregatedList.ItemsEntry`](gcp_compute.md#TargetPoolAggregatedList.ItemsEntry) [`items`](#TargetPoolAggregatedList.items) = 2
* `string` [`kind`](#TargetPoolAggregatedList.kind) = 3
* `string` [`nextPageToken`](#TargetPoolAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#TargetPoolAggregatedList.selfLink) = 5
* [`compute.TargetPoolAggregatedList.Warning`](gcp_compute.md#TargetPoolAggregatedList.Warning) [`warning`](#TargetPoolAggregatedList.warning) = 6

### `id` {#TargetPoolAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetPoolAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetPoolAggregatedList.ItemsEntry`](gcp_compute.md#TargetPoolAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetPool resources.

### `kind` {#TargetPoolAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetPoolAggregatedList for
aggregated lists of target pools.

### `nextPageToken` {#TargetPoolAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetPoolAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetPoolAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetPoolAggregatedList.Warning`](gcp_compute.md#TargetPoolAggregatedList.Warning) |

## Message `TargetPoolInstanceHealth` {#TargetPoolInstanceHealth}



### Inputs for `TargetPoolInstanceHealth`

* `repeated` [`compute.HealthStatus`](gcp_compute.md#HealthStatus) [`healthStatus`](#TargetPoolInstanceHealth.healthStatus) = 1
* `string` [`kind`](#TargetPoolInstanceHealth.kind) = 2

### `healthStatus` {#TargetPoolInstanceHealth.healthStatus}

| Property | Comments |
|----------|----------|
| Field Name | `healthStatus` |
| Type | [`compute.HealthStatus`](gcp_compute.md#HealthStatus) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `kind` {#TargetPoolInstanceHealth.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetPoolInstanceHealth when
checking the health of an instance.

## Message `TargetPoolList` {#TargetPoolList}

Contains a list of TargetPool resources.


### Inputs for `TargetPoolList`

* `string` [`id`](#TargetPoolList.id) = 1
* `repeated` [`compute.TargetPool`](gcp_compute.md#TargetPool) [`items`](#TargetPoolList.items) = 2
* `string` [`kind`](#TargetPoolList.kind) = 3
* `string` [`nextPageToken`](#TargetPoolList.nextPageToken) = 4
* `string` [`selfLink`](#TargetPoolList.selfLink) = 5
* [`compute.TargetPoolList.Warning`](gcp_compute.md#TargetPoolList.Warning) [`warning`](#TargetPoolList.warning) = 6

### `id` {#TargetPoolList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetPoolList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetPool`](gcp_compute.md#TargetPool) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetPool resources.

### `kind` {#TargetPoolList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetPoolList for lists of
target pools.

### `nextPageToken` {#TargetPoolList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetPoolList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetPoolList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetPoolList.Warning`](gcp_compute.md#TargetPoolList.Warning) |

## Message `TargetPoolsAddHealthCheckRequest` {#TargetPoolsAddHealthCheckRequest}



### Inputs for `TargetPoolsAddHealthCheckRequest`

* `repeated` [`compute.HealthCheckReference`](gcp_compute.md#HealthCheckReference) [`healthChecks`](#TargetPoolsAddHealthCheckRequest.healthChecks) = 1

### `healthChecks` {#TargetPoolsAddHealthCheckRequest.healthChecks}

| Property | Comments |
|----------|----------|
| Field Name | `healthChecks` |
| Type | [`compute.HealthCheckReference`](gcp_compute.md#HealthCheckReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The HttpHealthCheck to add to the target pool.

## Message `TargetPoolsAddInstanceRequest` {#TargetPoolsAddInstanceRequest}



### Inputs for `TargetPoolsAddInstanceRequest`

* `repeated` [`compute.InstanceReference`](gcp_compute.md#InstanceReference) [`instances`](#TargetPoolsAddInstanceRequest.instances) = 1

### `instances` {#TargetPoolsAddInstanceRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | [`compute.InstanceReference`](gcp_compute.md#InstanceReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A full or partial URL to an instance to add to this target pool. This can be
a full or partial URL. For example, the following are valid URLs:
-
https://www.googleapis.com/compute/v1/projects/project-id/zones/zone/instances/instance-name
- projects/project-id/zones/zone/instances/instance-name
- zones/zone/instances/instance-name

## Message `TargetPoolsRemoveHealthCheckRequest` {#TargetPoolsRemoveHealthCheckRequest}



### Inputs for `TargetPoolsRemoveHealthCheckRequest`

* `repeated` [`compute.HealthCheckReference`](gcp_compute.md#HealthCheckReference) [`healthChecks`](#TargetPoolsRemoveHealthCheckRequest.healthChecks) = 1

### `healthChecks` {#TargetPoolsRemoveHealthCheckRequest.healthChecks}

| Property | Comments |
|----------|----------|
| Field Name | `healthChecks` |
| Type | [`compute.HealthCheckReference`](gcp_compute.md#HealthCheckReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Health check URL to be removed. This can be a full or valid partial URL. For
example, the following are valid URLs:
-
https://www.googleapis.com/compute/beta/projects/project/global/httpHealthChecks/health-check
- projects/project/global/httpHealthChecks/health-check
- global/httpHealthChecks/health-check

## Message `TargetPoolsRemoveInstanceRequest` {#TargetPoolsRemoveInstanceRequest}



### Inputs for `TargetPoolsRemoveInstanceRequest`

* `repeated` [`compute.InstanceReference`](gcp_compute.md#InstanceReference) [`instances`](#TargetPoolsRemoveInstanceRequest.instances) = 1

### `instances` {#TargetPoolsRemoveInstanceRequest.instances}

| Property | Comments |
|----------|----------|
| Field Name | `instances` |
| Type | [`compute.InstanceReference`](gcp_compute.md#InstanceReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

URLs of the instances to be removed from target pool.

## Message `TargetPoolsScopedList` {#TargetPoolsScopedList}



### Inputs for `TargetPoolsScopedList`

* `repeated` [`compute.TargetPool`](gcp_compute.md#TargetPool) [`targetPools`](#TargetPoolsScopedList.targetPools) = 1
* [`compute.TargetPoolsScopedList.Warning`](gcp_compute.md#TargetPoolsScopedList.Warning) [`warning`](#TargetPoolsScopedList.warning) = 2

### `targetPools` {#TargetPoolsScopedList.targetPools}

| Property | Comments |
|----------|----------|
| Field Name | `targetPools` |
| Type | [`compute.TargetPool`](gcp_compute.md#TargetPool) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of target pools contained in this scope.

### `warning` {#TargetPoolsScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetPoolsScopedList.Warning`](gcp_compute.md#TargetPoolsScopedList.Warning) |

## Message `TargetReference` {#TargetReference}



### Inputs for `TargetReference`

* `string` [`target`](#TargetReference.target) = 1

### `target` {#TargetReference.target}

| Property | Comments |
|----------|----------|
| Field Name | `target` |
| Type | `string` |

## Message `TargetSslProxiesSetBackendServiceRequest` {#TargetSslProxiesSetBackendServiceRequest}



### Inputs for `TargetSslProxiesSetBackendServiceRequest`

* `string` [`service`](#TargetSslProxiesSetBackendServiceRequest.service) = 1

### `service` {#TargetSslProxiesSetBackendServiceRequest.service}

| Property | Comments |
|----------|----------|
| Field Name | `service` |
| Type | `string` |

The URL of the new BackendService resource for the targetSslProxy.

## Message `TargetSslProxiesSetProxyHeaderRequest` {#TargetSslProxiesSetProxyHeaderRequest}



### Inputs for `TargetSslProxiesSetProxyHeaderRequest`

* `string` [`proxyHeader`](#TargetSslProxiesSetProxyHeaderRequest.proxyHeader) = 1

### `proxyHeader` {#TargetSslProxiesSetProxyHeaderRequest.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

The new type of proxy header to append before sending data to the backend.
NONE or PROXY_V1 are allowed.
Valid values:
    NONE
    PROXY_V1

## Message `TargetSslProxiesSetSslCertificatesRequest` {#TargetSslProxiesSetSslCertificatesRequest}



### Inputs for `TargetSslProxiesSetSslCertificatesRequest`

* `repeated` `string` [`sslCertificates`](#TargetSslProxiesSetSslCertificatesRequest.sslCertificates) = 1

### `sslCertificates` {#TargetSslProxiesSetSslCertificatesRequest.sslCertificates}

| Property | Comments |
|----------|----------|
| Field Name | `sslCertificates` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

New set of URLs to SslCertificate resources to associate with this
TargetSslProxy. Currently exactly one ssl certificate must be specified.

## Message `TargetSslProxy` {#TargetSslProxy}

A TargetSslProxy resource. This resource defines an SSL proxy. (==
resource_for beta.targetSslProxies ==) (== resource_for v1.targetSslProxies
==)


### Inputs for `TargetSslProxy`

* `string` [`creationTimestamp`](#TargetSslProxy.creationTimestamp) = 1
* `string` [`description`](#TargetSslProxy.description) = 2
* `string` [`id`](#TargetSslProxy.id) = 3
* `string` [`kind`](#TargetSslProxy.kind) = 4
* `string` [`name`](#TargetSslProxy.name) = 5 (**Required**)
* `string` [`proxyHeader`](#TargetSslProxy.proxyHeader) = 6
* `string` [`selfLink`](#TargetSslProxy.selfLink) = 7
* `string` [`service`](#TargetSslProxy.service) = 8
* `repeated` `string` [`sslCertificates`](#TargetSslProxy.sslCertificates) = 9
* `string` [`sslPolicy`](#TargetSslProxy.sslPolicy) = 10

### `creationTimestamp` {#TargetSslProxy.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetSslProxy.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#TargetSslProxy.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#TargetSslProxy.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#targetSslProxy for target
SSL proxies.

### `name` {#TargetSslProxy.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `proxyHeader` {#TargetSslProxy.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

Specifies the type of proxy header to append before sending data to the
backend, either NONE or PROXY_V1. The default is NONE.
Valid values:
    NONE
    PROXY_V1

### `selfLink` {#TargetSslProxy.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `service` {#TargetSslProxy.service}

| Property | Comments |
|----------|----------|
| Field Name | `service` |
| Type | `string` |

URL to the BackendService resource.

### `sslCertificates` {#TargetSslProxy.sslCertificates}

| Property | Comments |
|----------|----------|
| Field Name | `sslCertificates` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

URLs to SslCertificate resources that are used to authenticate connections
to Backends. Currently exactly one SSL certificate must be specified.

### `sslPolicy` {#TargetSslProxy.sslPolicy}

| Property | Comments |
|----------|----------|
| Field Name | `sslPolicy` |
| Type | `string` |

URL of SslPolicy resource that will be associated with the TargetSslProxy
resource. If not set, the TargetSslProxy resource will not have any SSL
policy configured.

## Message `TargetSslProxyList` {#TargetSslProxyList}

Contains a list of TargetSslProxy resources.


### Inputs for `TargetSslProxyList`

* `string` [`id`](#TargetSslProxyList.id) = 1
* `repeated` [`compute.TargetSslProxy`](gcp_compute.md#TargetSslProxy) [`items`](#TargetSslProxyList.items) = 2
* `string` [`kind`](#TargetSslProxyList.kind) = 3
* `string` [`nextPageToken`](#TargetSslProxyList.nextPageToken) = 4
* `string` [`selfLink`](#TargetSslProxyList.selfLink) = 5
* [`compute.TargetSslProxyList.Warning`](gcp_compute.md#TargetSslProxyList.Warning) [`warning`](#TargetSslProxyList.warning) = 6

### `id` {#TargetSslProxyList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetSslProxyList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetSslProxy`](gcp_compute.md#TargetSslProxy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetSslProxy resources.

### `kind` {#TargetSslProxyList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#TargetSslProxyList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetSslProxyList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetSslProxyList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetSslProxyList.Warning`](gcp_compute.md#TargetSslProxyList.Warning) |

## Message `TargetTcpProxiesSetBackendServiceRequest` {#TargetTcpProxiesSetBackendServiceRequest}



### Inputs for `TargetTcpProxiesSetBackendServiceRequest`

* `string` [`service`](#TargetTcpProxiesSetBackendServiceRequest.service) = 1

### `service` {#TargetTcpProxiesSetBackendServiceRequest.service}

| Property | Comments |
|----------|----------|
| Field Name | `service` |
| Type | `string` |

The URL of the new BackendService resource for the targetTcpProxy.

## Message `TargetTcpProxiesSetProxyHeaderRequest` {#TargetTcpProxiesSetProxyHeaderRequest}



### Inputs for `TargetTcpProxiesSetProxyHeaderRequest`

* `string` [`proxyHeader`](#TargetTcpProxiesSetProxyHeaderRequest.proxyHeader) = 1

### `proxyHeader` {#TargetTcpProxiesSetProxyHeaderRequest.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

The new type of proxy header to append before sending data to the backend.
NONE or PROXY_V1 are allowed.
Valid values:
    NONE
    PROXY_V1

## Message `TargetTcpProxy` {#TargetTcpProxy}

A TargetTcpProxy resource. This resource defines a TCP proxy. (==
resource_for beta.targetTcpProxies ==) (== resource_for v1.targetTcpProxies
==)


### Inputs for `TargetTcpProxy`

* `string` [`creationTimestamp`](#TargetTcpProxy.creationTimestamp) = 1
* `string` [`description`](#TargetTcpProxy.description) = 2
* `string` [`id`](#TargetTcpProxy.id) = 3
* `string` [`kind`](#TargetTcpProxy.kind) = 4
* `string` [`name`](#TargetTcpProxy.name) = 5 (**Required**)
* `string` [`proxyHeader`](#TargetTcpProxy.proxyHeader) = 6
* `string` [`selfLink`](#TargetTcpProxy.selfLink) = 7
* `string` [`service`](#TargetTcpProxy.service) = 8

### `creationTimestamp` {#TargetTcpProxy.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetTcpProxy.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `id` {#TargetTcpProxy.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#TargetTcpProxy.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#targetTcpProxy for target
TCP proxies.

### `name` {#TargetTcpProxy.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `proxyHeader` {#TargetTcpProxy.proxyHeader}

| Property | Comments |
|----------|----------|
| Field Name | `proxyHeader` |
| Type | `string` |

Specifies the type of proxy header to append before sending data to the
backend, either NONE or PROXY_V1. The default is NONE.
Valid values:
    NONE
    PROXY_V1

### `selfLink` {#TargetTcpProxy.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `service` {#TargetTcpProxy.service}

| Property | Comments |
|----------|----------|
| Field Name | `service` |
| Type | `string` |

URL to the BackendService resource.

## Message `TargetTcpProxyList` {#TargetTcpProxyList}

Contains a list of TargetTcpProxy resources.


### Inputs for `TargetTcpProxyList`

* `string` [`id`](#TargetTcpProxyList.id) = 1
* `repeated` [`compute.TargetTcpProxy`](gcp_compute.md#TargetTcpProxy) [`items`](#TargetTcpProxyList.items) = 2
* `string` [`kind`](#TargetTcpProxyList.kind) = 3
* `string` [`nextPageToken`](#TargetTcpProxyList.nextPageToken) = 4
* `string` [`selfLink`](#TargetTcpProxyList.selfLink) = 5
* [`compute.TargetTcpProxyList.Warning`](gcp_compute.md#TargetTcpProxyList.Warning) [`warning`](#TargetTcpProxyList.warning) = 6

### `id` {#TargetTcpProxyList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetTcpProxyList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetTcpProxy`](gcp_compute.md#TargetTcpProxy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetTcpProxy resources.

### `kind` {#TargetTcpProxyList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#TargetTcpProxyList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetTcpProxyList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetTcpProxyList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetTcpProxyList.Warning`](gcp_compute.md#TargetTcpProxyList.Warning) |

## Message `TargetVpnGateway` {#TargetVpnGateway}

Represents a Target VPN gateway resource. (== resource_for
beta.targetVpnGateways ==) (== resource_for v1.targetVpnGateways ==)


### Inputs for `TargetVpnGateway`

* `string` [`creationTimestamp`](#TargetVpnGateway.creationTimestamp) = 1
* `string` [`description`](#TargetVpnGateway.description) = 2
* `repeated` `string` [`forwardingRules`](#TargetVpnGateway.forwardingRules) = 3
* `string` [`id`](#TargetVpnGateway.id) = 4
* `string` [`kind`](#TargetVpnGateway.kind) = 5
* `string` [`name`](#TargetVpnGateway.name) = 6 (**Required**)
* `string` [`network`](#TargetVpnGateway.network) = 7
* `string` [`region`](#TargetVpnGateway.region) = 8
* `string` [`selfLink`](#TargetVpnGateway.selfLink) = 9
* `string` [`status`](#TargetVpnGateway.status) = 10
* `repeated` `string` [`tunnels`](#TargetVpnGateway.tunnels) = 11

### `creationTimestamp` {#TargetVpnGateway.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#TargetVpnGateway.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `forwardingRules` {#TargetVpnGateway.forwardingRules}

| Property | Comments |
|----------|----------|
| Field Name | `forwardingRules` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of URLs to the ForwardingRule resources.
ForwardingRules are created using compute.forwardingRules.insert and
associated to a VPN gateway.

### `id` {#TargetVpnGateway.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#TargetVpnGateway.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetVpnGateway for target
VPN gateways.

### `name` {#TargetVpnGateway.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `network` {#TargetVpnGateway.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |

URL of the network to which this VPN gateway is attached. Provided by the
client when the VPN gateway is created.

### `region` {#TargetVpnGateway.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the target VPN gateway resides. You
must specify this field as part of the HTTP request URL. It is not settable
as a field in the request body.

### `selfLink` {#TargetVpnGateway.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `status` {#TargetVpnGateway.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the VPN gateway.
Valid values:
    CREATING
    DELETING
    FAILED
    READY

### `tunnels` {#TargetVpnGateway.tunnels}

| Property | Comments |
|----------|----------|
| Field Name | `tunnels` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of URLs to VpnTunnel resources. VpnTunnels are created
using compute.vpntunnels.insert method and associated to a VPN gateway.

## Message `TargetVpnGatewayAggregatedList` {#TargetVpnGatewayAggregatedList}



### Inputs for `TargetVpnGatewayAggregatedList`

* `string` [`id`](#TargetVpnGatewayAggregatedList.id) = 1
* `repeated` [`compute.TargetVpnGatewayAggregatedList.ItemsEntry`](gcp_compute.md#TargetVpnGatewayAggregatedList.ItemsEntry) [`items`](#TargetVpnGatewayAggregatedList.items) = 2
* `string` [`kind`](#TargetVpnGatewayAggregatedList.kind) = 3
* `string` [`nextPageToken`](#TargetVpnGatewayAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#TargetVpnGatewayAggregatedList.selfLink) = 5
* [`compute.TargetVpnGatewayAggregatedList.Warning`](gcp_compute.md#TargetVpnGatewayAggregatedList.Warning) [`warning`](#TargetVpnGatewayAggregatedList.warning) = 6

### `id` {#TargetVpnGatewayAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetVpnGatewayAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetVpnGatewayAggregatedList.ItemsEntry`](gcp_compute.md#TargetVpnGatewayAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetVpnGateway resources.

### `kind` {#TargetVpnGatewayAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetVpnGateway for target
VPN gateways.

### `nextPageToken` {#TargetVpnGatewayAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetVpnGatewayAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetVpnGatewayAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetVpnGatewayAggregatedList.Warning`](gcp_compute.md#TargetVpnGatewayAggregatedList.Warning) |

## Message `TargetVpnGatewayList` {#TargetVpnGatewayList}

Contains a list of TargetVpnGateway resources.


### Inputs for `TargetVpnGatewayList`

* `string` [`id`](#TargetVpnGatewayList.id) = 1
* `repeated` [`compute.TargetVpnGateway`](gcp_compute.md#TargetVpnGateway) [`items`](#TargetVpnGatewayList.items) = 2
* `string` [`kind`](#TargetVpnGatewayList.kind) = 3
* `string` [`nextPageToken`](#TargetVpnGatewayList.nextPageToken) = 4
* `string` [`selfLink`](#TargetVpnGatewayList.selfLink) = 5
* [`compute.TargetVpnGatewayList.Warning`](gcp_compute.md#TargetVpnGatewayList.Warning) [`warning`](#TargetVpnGatewayList.warning) = 6

### `id` {#TargetVpnGatewayList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#TargetVpnGatewayList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.TargetVpnGateway`](gcp_compute.md#TargetVpnGateway) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of TargetVpnGateway resources.

### `kind` {#TargetVpnGatewayList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#targetVpnGateway for target
VPN gateways.

### `nextPageToken` {#TargetVpnGatewayList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#TargetVpnGatewayList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#TargetVpnGatewayList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetVpnGatewayList.Warning`](gcp_compute.md#TargetVpnGatewayList.Warning) |

## Message `TargetVpnGatewaysScopedList` {#TargetVpnGatewaysScopedList}



### Inputs for `TargetVpnGatewaysScopedList`

* `repeated` [`compute.TargetVpnGateway`](gcp_compute.md#TargetVpnGateway) [`targetVpnGateways`](#TargetVpnGatewaysScopedList.targetVpnGateways) = 1
* [`compute.TargetVpnGatewaysScopedList.Warning`](gcp_compute.md#TargetVpnGatewaysScopedList.Warning) [`warning`](#TargetVpnGatewaysScopedList.warning) = 2

### `targetVpnGateways` {#TargetVpnGatewaysScopedList.targetVpnGateways}

| Property | Comments |
|----------|----------|
| Field Name | `targetVpnGateways` |
| Type | [`compute.TargetVpnGateway`](gcp_compute.md#TargetVpnGateway) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of target vpn gateways contained in this scope.

### `warning` {#TargetVpnGatewaysScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.TargetVpnGatewaysScopedList.Warning`](gcp_compute.md#TargetVpnGatewaysScopedList.Warning) |

## Message `TestFailure` {#TestFailure}



### Inputs for `TestFailure`

* `string` [`actualService`](#TestFailure.actualService) = 1
* `string` [`expectedService`](#TestFailure.expectedService) = 2
* `string` [`host`](#TestFailure.host) = 3
* `string` [`path`](#TestFailure.path) = 4

### `actualService` {#TestFailure.actualService}

| Property | Comments |
|----------|----------|
| Field Name | `actualService` |
| Type | `string` |

### `expectedService` {#TestFailure.expectedService}

| Property | Comments |
|----------|----------|
| Field Name | `expectedService` |
| Type | `string` |

### `host` {#TestFailure.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

### `path` {#TestFailure.path}

| Property | Comments |
|----------|----------|
| Field Name | `path` |
| Type | `string` |

## Message `TestPermissionsRequest` {#TestPermissionsRequest}



### Inputs for `TestPermissionsRequest`

* `repeated` `string` [`permissions`](#TestPermissionsRequest.permissions) = 1

### `permissions` {#TestPermissionsRequest.permissions}

| Property | Comments |
|----------|----------|
| Field Name | `permissions` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

The set of permissions to check for the 'resource'. Permissions with
wildcards (such as '*' or 'storage.*') are not allowed.

## Message `TestPermissionsResponse` {#TestPermissionsResponse}



### Inputs for `TestPermissionsResponse`

* `repeated` `string` [`permissions`](#TestPermissionsResponse.permissions) = 1

### `permissions` {#TestPermissionsResponse.permissions}

| Property | Comments |
|----------|----------|
| Field Name | `permissions` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A subset of `TestPermissionsRequest.permissions` that the caller is allowed.

## Message `UrlMap` {#UrlMap}

A UrlMap resource. This resource defines the mapping from URL to the
BackendService resource, based on the "longest-match" of the URL's host and
path.


### Inputs for `UrlMap`

* `string` [`creationTimestamp`](#UrlMap.creationTimestamp) = 1
* `string` [`defaultService`](#UrlMap.defaultService) = 2
* `string` [`description`](#UrlMap.description) = 3
* `string` [`fingerprint`](#UrlMap.fingerprint) = 4
* `repeated` [`compute.HostRule`](gcp_compute.md#HostRule) [`hostRules`](#UrlMap.hostRules) = 5
* `string` [`id`](#UrlMap.id) = 6
* `string` [`kind`](#UrlMap.kind) = 7
* `string` [`name`](#UrlMap.name) = 8 (**Required**)
* `repeated` [`compute.PathMatcher`](gcp_compute.md#PathMatcher) [`pathMatchers`](#UrlMap.pathMatchers) = 9
* `string` [`selfLink`](#UrlMap.selfLink) = 10
* `repeated` [`compute.UrlMapTest`](gcp_compute.md#UrlMapTest) [`tests`](#UrlMap.tests) = 11

### `creationTimestamp` {#UrlMap.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `defaultService` {#UrlMap.defaultService}

| Property | Comments |
|----------|----------|
| Field Name | `defaultService` |
| Type | `string` |

The URL of the BackendService resource if none of the hostRules match.

### `description` {#UrlMap.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `fingerprint` {#UrlMap.fingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `fingerprint` |
| Type | `string` |

Fingerprint of this resource. A hash of the contents stored in this object.
This field is used in optimistic locking. This field will be ignored when
inserting a UrlMap. An up-to-date fingerprint must be provided in order to
update the UrlMap.

### `hostRules` {#UrlMap.hostRules}

| Property | Comments |
|----------|----------|
| Field Name | `hostRules` |
| Type | [`compute.HostRule`](gcp_compute.md#HostRule) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of HostRules to use against the URL.

### `id` {#UrlMap.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#UrlMap.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#urlMaps for url maps.

### `name` {#UrlMap.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `pathMatchers` {#UrlMap.pathMatchers}

| Property | Comments |
|----------|----------|
| Field Name | `pathMatchers` |
| Type | [`compute.PathMatcher`](gcp_compute.md#PathMatcher) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of named PathMatchers to use against the URL.

### `selfLink` {#UrlMap.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `tests` {#UrlMap.tests}

| Property | Comments |
|----------|----------|
| Field Name | `tests` |
| Type | [`compute.UrlMapTest`](gcp_compute.md#UrlMapTest) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of expected URL mapping tests. Request to update this UrlMap will
succeed only if all of the test cases pass. You can specify a maximum of 100
tests per UrlMap.

## Message `UrlMapList` {#UrlMapList}

Contains a list of UrlMap resources.


### Inputs for `UrlMapList`

* `string` [`id`](#UrlMapList.id) = 1
* `repeated` [`compute.UrlMap`](gcp_compute.md#UrlMap) [`items`](#UrlMapList.items) = 2
* `string` [`kind`](#UrlMapList.kind) = 3
* `string` [`nextPageToken`](#UrlMapList.nextPageToken) = 4
* `string` [`selfLink`](#UrlMapList.selfLink) = 5
* [`compute.UrlMapList.Warning`](gcp_compute.md#UrlMapList.Warning) [`warning`](#UrlMapList.warning) = 6

### `id` {#UrlMapList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#UrlMapList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.UrlMap`](gcp_compute.md#UrlMap) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of UrlMap resources.

### `kind` {#UrlMapList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#UrlMapList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#UrlMapList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#UrlMapList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.UrlMapList.Warning`](gcp_compute.md#UrlMapList.Warning) |

## Message `UrlMapReference` {#UrlMapReference}



### Inputs for `UrlMapReference`

* `string` [`urlMap`](#UrlMapReference.urlMap) = 1

### `urlMap` {#UrlMapReference.urlMap}

| Property | Comments |
|----------|----------|
| Field Name | `urlMap` |
| Type | `string` |

## Message `UrlMapTest` {#UrlMapTest}

Message for the expected URL mappings.


### Inputs for `UrlMapTest`

* `string` [`description`](#UrlMapTest.description) = 1
* `string` [`host`](#UrlMapTest.host) = 2
* `string` [`path`](#UrlMapTest.path) = 3
* `string` [`service`](#UrlMapTest.service) = 4

### `description` {#UrlMapTest.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

Description of this test case.

### `host` {#UrlMapTest.host}

| Property | Comments |
|----------|----------|
| Field Name | `host` |
| Type | `string` |

Host portion of the URL.

### `path` {#UrlMapTest.path}

| Property | Comments |
|----------|----------|
| Field Name | `path` |
| Type | `string` |

Path portion of the URL.

### `service` {#UrlMapTest.service}

| Property | Comments |
|----------|----------|
| Field Name | `service` |
| Type | `string` |

Expected BackendService resource the given URL should be mapped to.

## Message `UrlMapValidationResult` {#UrlMapValidationResult}

Message representing the validation result for a UrlMap.


### Inputs for `UrlMapValidationResult`

* `repeated` `string` [`loadErrors`](#UrlMapValidationResult.loadErrors) = 1
* `bool` [`loadSucceeded`](#UrlMapValidationResult.loadSucceeded) = 2
* `repeated` [`compute.TestFailure`](gcp_compute.md#TestFailure) [`testFailures`](#UrlMapValidationResult.testFailures) = 3
* `bool` [`testPassed`](#UrlMapValidationResult.testPassed) = 4

### `loadErrors` {#UrlMapValidationResult.loadErrors}

| Property | Comments |
|----------|----------|
| Field Name | `loadErrors` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `loadSucceeded` {#UrlMapValidationResult.loadSucceeded}

| Property | Comments |
|----------|----------|
| Field Name | `loadSucceeded` |
| Type | `bool` |

Whether the given UrlMap can be successfully loaded. If false, 'loadErrors'
indicates the reasons.

### `testFailures` {#UrlMapValidationResult.testFailures}

| Property | Comments |
|----------|----------|
| Field Name | `testFailures` |
| Type | [`compute.TestFailure`](gcp_compute.md#TestFailure) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `testPassed` {#UrlMapValidationResult.testPassed}

| Property | Comments |
|----------|----------|
| Field Name | `testPassed` |
| Type | `bool` |

If successfully loaded, this field indicates whether the test passed. If
false, 'testFailures's indicate the reason of failure.

## Message `UrlMapsValidateRequest` {#UrlMapsValidateRequest}



### Inputs for `UrlMapsValidateRequest`

* [`compute.UrlMap`](gcp_compute.md#UrlMap) [`resource`](#UrlMapsValidateRequest.resource) = 1

### `resource` {#UrlMapsValidateRequest.resource}

| Property | Comments |
|----------|----------|
| Field Name | `resource` |
| Type | [`compute.UrlMap`](gcp_compute.md#UrlMap) |

Content of the UrlMap to be validated.

## Message `UrlMapsValidateResponse` {#UrlMapsValidateResponse}



### Inputs for `UrlMapsValidateResponse`

* [`compute.UrlMapValidationResult`](gcp_compute.md#UrlMapValidationResult) [`result`](#UrlMapsValidateResponse.result) = 1

### `result` {#UrlMapsValidateResponse.result}

| Property | Comments |
|----------|----------|
| Field Name | `result` |
| Type | [`compute.UrlMapValidationResult`](gcp_compute.md#UrlMapValidationResult) |

## Message `UsageExportLocation` {#UsageExportLocation}

The location in Cloud Storage and naming method of the daily usage report.
Contains bucket_name and report_name prefix.


### Inputs for `UsageExportLocation`

* `string` [`bucketName`](#UsageExportLocation.bucketName) = 1
* `string` [`reportNamePrefix`](#UsageExportLocation.reportNamePrefix) = 2

### `bucketName` {#UsageExportLocation.bucketName}

| Property | Comments |
|----------|----------|
| Field Name | `bucketName` |
| Type | `string` |

The name of an existing bucket in Cloud Storage where the usage report
object is stored. The Google Service Account is granted write access to this
bucket. This can either be the bucket name by itself, such as
example-bucket, or the bucket name with gs:// or
https://storage.googleapis.com/ in front of it, such as gs://example-bucket.

### `reportNamePrefix` {#UsageExportLocation.reportNamePrefix}

| Property | Comments |
|----------|----------|
| Field Name | `reportNamePrefix` |
| Type | `string` |

An optional prefix for the name of the usage report object stored in
bucketName. If not supplied, defaults to usage. The report is stored as a
CSV file named report_name_prefix_gce_YYYYMMDD.csv where YYYYMMDD is the day
of the usage according to Pacific Time. If you supply a prefix, it should
conform to Cloud Storage object naming conventions.

## Message `VpnTunnel` {#VpnTunnel}

VPN tunnel resource. (== resource_for beta.vpnTunnels ==) (== resource_for
v1.vpnTunnels ==)


### Inputs for `VpnTunnel`

* `string` [`creationTimestamp`](#VpnTunnel.creationTimestamp) = 1
* `string` [`description`](#VpnTunnel.description) = 2
* `string` [`detailedStatus`](#VpnTunnel.detailedStatus) = 3
* `string` [`id`](#VpnTunnel.id) = 4
* `int32` [`ikeVersion`](#VpnTunnel.ikeVersion) = 5
* `string` [`kind`](#VpnTunnel.kind) = 6
* `repeated` `string` [`localTrafficSelector`](#VpnTunnel.localTrafficSelector) = 7
* `string` [`name`](#VpnTunnel.name) = 8 (**Required**)
* `string` [`peerIp`](#VpnTunnel.peerIp) = 9
* `string` [`region`](#VpnTunnel.region) = 10
* `repeated` `string` [`remoteTrafficSelector`](#VpnTunnel.remoteTrafficSelector) = 11
* `string` [`router`](#VpnTunnel.router) = 12
* `string` [`selfLink`](#VpnTunnel.selfLink) = 13
* `string` [`sharedSecret`](#VpnTunnel.sharedSecret) = 14
* `string` [`sharedSecretHash`](#VpnTunnel.sharedSecretHash) = 15
* `string` [`status`](#VpnTunnel.status) = 16
* `string` [`targetVpnGateway`](#VpnTunnel.targetVpnGateway) = 17

### `creationTimestamp` {#VpnTunnel.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `description` {#VpnTunnel.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

An optional description of this resource. Provide this property when you
create the resource.

### `detailedStatus` {#VpnTunnel.detailedStatus}

| Property | Comments |
|----------|----------|
| Field Name | `detailedStatus` |
| Type | `string` |

[Output Only] Detailed status message for the VPN tunnel.

### `id` {#VpnTunnel.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `ikeVersion` {#VpnTunnel.ikeVersion}

| Property | Comments |
|----------|----------|
| Field Name | `ikeVersion` |
| Type | `int32` |

IKE protocol version to use when establishing the VPN tunnel with peer VPN
gateway. Acceptable IKE versions are 1 or 2. Default version is 2.

### `kind` {#VpnTunnel.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#vpnTunnel for VPN tunnels.

### `localTrafficSelector` {#VpnTunnel.localTrafficSelector}

| Property | Comments |
|----------|----------|
| Field Name | `localTrafficSelector` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Local traffic selector to use when establishing the VPN tunnel with peer VPN
gateway. The value should be a CIDR formatted string, for example:
192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.

### `name` {#VpnTunnel.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the resource. Provided by the client when the resource is created.
The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular
expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must
be a lowercase letter, and all following characters must be a dash,
lowercase letter, or digit, except the last character, which cannot be a
dash.

### `peerIp` {#VpnTunnel.peerIp}

| Property | Comments |
|----------|----------|
| Field Name | `peerIp` |
| Type | `string` |

IP address of the peer VPN gateway. Only IPv4 is supported.

### `region` {#VpnTunnel.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] URL of the region where the VPN tunnel resides. You must
specify this field as part of the HTTP request URL. It is not settable as a
field in the request body.

### `remoteTrafficSelector` {#VpnTunnel.remoteTrafficSelector}

| Property | Comments |
|----------|----------|
| Field Name | `remoteTrafficSelector` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

Remote traffic selectors to use when establishing the VPN tunnel with peer
VPN gateway. The value should be a CIDR formatted string, for example:
192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.

### `router` {#VpnTunnel.router}

| Property | Comments |
|----------|----------|
| Field Name | `router` |
| Type | `string` |

URL of router resource to be used for dynamic routing.

### `selfLink` {#VpnTunnel.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `sharedSecret` {#VpnTunnel.sharedSecret}

| Property | Comments |
|----------|----------|
| Field Name | `sharedSecret` |
| Type | `string` |

Shared secret used to set the secure session between the Cloud VPN gateway
and the peer VPN gateway.

### `sharedSecretHash` {#VpnTunnel.sharedSecretHash}

| Property | Comments |
|----------|----------|
| Field Name | `sharedSecretHash` |
| Type | `string` |

Hash of the shared secret.

### `status` {#VpnTunnel.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] The status of the VPN tunnel.
Valid values:
    ALLOCATING_RESOURCES
    AUTHORIZATION_ERROR
    DEPROVISIONING
    ESTABLISHED
    FAILED
    FIRST_HANDSHAKE
    NEGOTIATION_FAILURE
    NETWORK_ERROR
    NO_INCOMING_PACKETS
    PROVISIONING
    REJECTED
    WAITING_FOR_FULL_CONFIG

### `targetVpnGateway` {#VpnTunnel.targetVpnGateway}

| Property | Comments |
|----------|----------|
| Field Name | `targetVpnGateway` |
| Type | `string` |

URL of the Target VPN gateway with which this VPN tunnel is associated.
Provided by the client when the VPN tunnel is created.

## Message `VpnTunnelAggregatedList` {#VpnTunnelAggregatedList}



### Inputs for `VpnTunnelAggregatedList`

* `string` [`id`](#VpnTunnelAggregatedList.id) = 1
* `repeated` [`compute.VpnTunnelAggregatedList.ItemsEntry`](gcp_compute.md#VpnTunnelAggregatedList.ItemsEntry) [`items`](#VpnTunnelAggregatedList.items) = 2
* `string` [`kind`](#VpnTunnelAggregatedList.kind) = 3
* `string` [`nextPageToken`](#VpnTunnelAggregatedList.nextPageToken) = 4
* `string` [`selfLink`](#VpnTunnelAggregatedList.selfLink) = 5
* [`compute.VpnTunnelAggregatedList.Warning`](gcp_compute.md#VpnTunnelAggregatedList.Warning) [`warning`](#VpnTunnelAggregatedList.warning) = 6

### `id` {#VpnTunnelAggregatedList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#VpnTunnelAggregatedList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.VpnTunnelAggregatedList.ItemsEntry`](gcp_compute.md#VpnTunnelAggregatedList.ItemsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of VpnTunnelsScopedList resources.

### `kind` {#VpnTunnelAggregatedList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#vpnTunnel for VPN tunnels.

### `nextPageToken` {#VpnTunnelAggregatedList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#VpnTunnelAggregatedList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#VpnTunnelAggregatedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.VpnTunnelAggregatedList.Warning`](gcp_compute.md#VpnTunnelAggregatedList.Warning) |

## Message `VpnTunnelList` {#VpnTunnelList}

Contains a list of VpnTunnel resources.


### Inputs for `VpnTunnelList`

* `string` [`id`](#VpnTunnelList.id) = 1
* `repeated` [`compute.VpnTunnel`](gcp_compute.md#VpnTunnel) [`items`](#VpnTunnelList.items) = 2
* `string` [`kind`](#VpnTunnelList.kind) = 3
* `string` [`nextPageToken`](#VpnTunnelList.nextPageToken) = 4
* `string` [`selfLink`](#VpnTunnelList.selfLink) = 5
* [`compute.VpnTunnelList.Warning`](gcp_compute.md#VpnTunnelList.Warning) [`warning`](#VpnTunnelList.warning) = 6

### `id` {#VpnTunnelList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#VpnTunnelList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.VpnTunnel`](gcp_compute.md#VpnTunnel) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of VpnTunnel resources.

### `kind` {#VpnTunnelList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#vpnTunnel for VPN tunnels.

### `nextPageToken` {#VpnTunnelList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#VpnTunnelList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#VpnTunnelList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.VpnTunnelList.Warning`](gcp_compute.md#VpnTunnelList.Warning) |

## Message `VpnTunnelsScopedList` {#VpnTunnelsScopedList}



### Inputs for `VpnTunnelsScopedList`

* `repeated` [`compute.VpnTunnel`](gcp_compute.md#VpnTunnel) [`vpnTunnels`](#VpnTunnelsScopedList.vpnTunnels) = 1
* [`compute.VpnTunnelsScopedList.Warning`](gcp_compute.md#VpnTunnelsScopedList.Warning) [`warning`](#VpnTunnelsScopedList.warning) = 2

### `vpnTunnels` {#VpnTunnelsScopedList.vpnTunnels}

| Property | Comments |
|----------|----------|
| Field Name | `vpnTunnels` |
| Type | [`compute.VpnTunnel`](gcp_compute.md#VpnTunnel) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of vpn tunnels contained in this scope.

### `warning` {#VpnTunnelsScopedList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.VpnTunnelsScopedList.Warning`](gcp_compute.md#VpnTunnelsScopedList.Warning) |

## Message `Warning` {#ImageList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#ImageList.Warning.code) = 1
* `repeated` [`compute.ImageList.Warning.Data`](gcp_compute.md#ImageList.Warning.Data) [`data`](#ImageList.Warning.data) = 2
* `string` [`message`](#ImageList.Warning.message) = 3

### `code` {#ImageList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#ImageList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.ImageList.Warning.Data`](gcp_compute.md#ImageList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#ImageList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RouterAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RouterAggregatedList.Warning.code) = 1
* `repeated` [`compute.RouterAggregatedList.Warning.Data`](gcp_compute.md#RouterAggregatedList.Warning.Data) [`data`](#RouterAggregatedList.Warning.data) = 2
* `string` [`message`](#RouterAggregatedList.Warning.message) = 3

### `code` {#RouterAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RouterAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RouterAggregatedList.Warning.Data`](gcp_compute.md#RouterAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RouterAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupsListInstances.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupsListInstances.Warning.code) = 1
* `repeated` [`compute.InstanceGroupsListInstances.Warning.Data`](gcp_compute.md#InstanceGroupsListInstances.Warning.Data) [`data`](#InstanceGroupsListInstances.Warning.data) = 2
* `string` [`message`](#InstanceGroupsListInstances.Warning.message) = 3

### `code` {#InstanceGroupsListInstances.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupsListInstances.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupsListInstances.Warning.Data`](gcp_compute.md#InstanceGroupsListInstances.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupsListInstances.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupsScopedList.Warning}

[Output Only] An informational warning that replaces the list of instance
groups when the list is empty.
[Output Only] An informational warning that replaces the list of instance
groups when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupsScopedList.Warning.code) = 1
* `repeated` [`compute.InstanceGroupsScopedList.Warning.Data`](gcp_compute.md#InstanceGroupsScopedList.Warning.Data) [`data`](#InstanceGroupsScopedList.Warning.data) = 2
* `string` [`message`](#InstanceGroupsScopedList.Warning.message) = 3

### `code` {#InstanceGroupsScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupsScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupsScopedList.Warning.Data`](gcp_compute.md#InstanceGroupsScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupsScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetHttpProxyList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetHttpProxyList.Warning.code) = 1
* `repeated` [`compute.TargetHttpProxyList.Warning.Data`](gcp_compute.md#TargetHttpProxyList.Warning.Data) [`data`](#TargetHttpProxyList.Warning.data) = 2
* `string` [`message`](#TargetHttpProxyList.Warning.message) = 3

### `code` {#TargetHttpProxyList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetHttpProxyList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetHttpProxyList.Warning.Data`](gcp_compute.md#TargetHttpProxyList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetHttpProxyList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#ForwardingRuleList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#ForwardingRuleList.Warning.code) = 1
* `repeated` [`compute.ForwardingRuleList.Warning.Data`](gcp_compute.md#ForwardingRuleList.Warning.Data) [`data`](#ForwardingRuleList.Warning.data) = 2
* `string` [`message`](#ForwardingRuleList.Warning.message) = 3

### `code` {#ForwardingRuleList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#ForwardingRuleList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.ForwardingRuleList.Warning.Data`](gcp_compute.md#ForwardingRuleList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#ForwardingRuleList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RouteList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RouteList.Warning.code) = 1
* `repeated` [`compute.RouteList.Warning.Data`](gcp_compute.md#RouteList.Warning.Data) [`data`](#RouteList.Warning.data) = 2
* `string` [`message`](#RouteList.Warning.message) = 3

### `code` {#RouteList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RouteList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RouteList.Warning.Data`](gcp_compute.md#RouteList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RouteList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceList.Warning.code) = 1
* `repeated` [`compute.InstanceList.Warning.Data`](gcp_compute.md#InstanceList.Warning.Data) [`data`](#InstanceList.Warning.data) = 2
* `string` [`message`](#InstanceList.Warning.message) = 3

### `code` {#InstanceList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceList.Warning.Data`](gcp_compute.md#InstanceList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#ZoneList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#ZoneList.Warning.code) = 1
* `repeated` [`compute.ZoneList.Warning.Data`](gcp_compute.md#ZoneList.Warning.Data) [`data`](#ZoneList.Warning.data) = 2
* `string` [`message`](#ZoneList.Warning.message) = 3

### `code` {#ZoneList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#ZoneList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.ZoneList.Warning.Data`](gcp_compute.md#ZoneList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#ZoneList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceListReferrers.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceListReferrers.Warning.code) = 1
* `repeated` [`compute.InstanceListReferrers.Warning.Data`](gcp_compute.md#InstanceListReferrers.Warning.Data) [`data`](#InstanceListReferrers.Warning.data) = 2
* `string` [`message`](#InstanceListReferrers.Warning.message) = 3

### `code` {#InstanceListReferrers.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceListReferrers.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceListReferrers.Warning.Data`](gcp_compute.md#InstanceListReferrers.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceListReferrers.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetHttpsProxyList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetHttpsProxyList.Warning.code) = 1
* `repeated` [`compute.TargetHttpsProxyList.Warning.Data`](gcp_compute.md#TargetHttpsProxyList.Warning.Data) [`data`](#TargetHttpsProxyList.Warning.data) = 2
* `string` [`message`](#TargetHttpsProxyList.Warning.message) = 3

### `code` {#TargetHttpsProxyList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetHttpsProxyList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetHttpsProxyList.Warning.Data`](gcp_compute.md#TargetHttpsProxyList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetHttpsProxyList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#ForwardingRuleAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#ForwardingRuleAggregatedList.Warning.code) = 1
* `repeated` [`compute.ForwardingRuleAggregatedList.Warning.Data`](gcp_compute.md#ForwardingRuleAggregatedList.Warning.Data) [`data`](#ForwardingRuleAggregatedList.Warning.data) = 2
* `string` [`message`](#ForwardingRuleAggregatedList.Warning.message) = 3

### `code` {#ForwardingRuleAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#ForwardingRuleAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.ForwardingRuleAggregatedList.Warning.Data`](gcp_compute.md#ForwardingRuleAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#ForwardingRuleAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceTemplateList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceTemplateList.Warning.code) = 1
* `repeated` [`compute.InstanceTemplateList.Warning.Data`](gcp_compute.md#InstanceTemplateList.Warning.Data) [`data`](#InstanceTemplateList.Warning.data) = 2
* `string` [`message`](#InstanceTemplateList.Warning.message) = 3

### `code` {#InstanceTemplateList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceTemplateList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceTemplateList.Warning.Data`](gcp_compute.md#InstanceTemplateList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceTemplateList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RegionList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RegionList.Warning.code) = 1
* `repeated` [`compute.RegionList.Warning.Data`](gcp_compute.md#RegionList.Warning.Data) [`data`](#RegionList.Warning.data) = 2
* `string` [`message`](#RegionList.Warning.message) = 3

### `code` {#RegionList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RegionList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RegionList.Warning.Data`](gcp_compute.md#RegionList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RegionList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#FirewallList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#FirewallList.Warning.code) = 1
* `repeated` [`compute.FirewallList.Warning.Data`](gcp_compute.md#FirewallList.Warning.Data) [`data`](#FirewallList.Warning.data) = 2
* `string` [`message`](#FirewallList.Warning.message) = 3

### `code` {#FirewallList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#FirewallList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.FirewallList.Warning.Data`](gcp_compute.md#FirewallList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#FirewallList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetInstanceAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetInstanceAggregatedList.Warning.code) = 1
* `repeated` [`compute.TargetInstanceAggregatedList.Warning.Data`](gcp_compute.md#TargetInstanceAggregatedList.Warning.Data) [`data`](#TargetInstanceAggregatedList.Warning.data) = 2
* `string` [`message`](#TargetInstanceAggregatedList.Warning.message) = 3

### `code` {#TargetInstanceAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetInstanceAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetInstanceAggregatedList.Warning.Data`](gcp_compute.md#TargetInstanceAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetInstanceAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#DisksScopedList.Warning}

[Output Only] Informational warning which replaces the list of disks when
the list is empty.
[Output Only] Informational warning which replaces the list of disks when
the list is empty.


### Inputs for `Warning`

* `string` [`code`](#DisksScopedList.Warning.code) = 1
* `repeated` [`compute.DisksScopedList.Warning.Data`](gcp_compute.md#DisksScopedList.Warning.Data) [`data`](#DisksScopedList.Warning.data) = 2
* `string` [`message`](#DisksScopedList.Warning.message) = 3

### `code` {#DisksScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#DisksScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.DisksScopedList.Warning.Data`](gcp_compute.md#DisksScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#DisksScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstancesScopedList.Warning}

[Output Only] Informational warning which replaces the list of instances
when the list is empty.
[Output Only] Informational warning which replaces the list of instances
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#InstancesScopedList.Warning.code) = 1
* `repeated` [`compute.InstancesScopedList.Warning.Data`](gcp_compute.md#InstancesScopedList.Warning.Data) [`data`](#InstancesScopedList.Warning.data) = 2
* `string` [`message`](#InstancesScopedList.Warning.message) = 3

### `code` {#InstancesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstancesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstancesScopedList.Warning.Data`](gcp_compute.md#InstancesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstancesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetInstanceList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetInstanceList.Warning.code) = 1
* `repeated` [`compute.TargetInstanceList.Warning.Data`](gcp_compute.md#TargetInstanceList.Warning.Data) [`data`](#TargetInstanceList.Warning.data) = 2
* `string` [`message`](#TargetInstanceList.Warning.message) = 3

### `code` {#TargetInstanceList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetInstanceList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetInstanceList.Warning.Data`](gcp_compute.md#TargetInstanceList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetInstanceList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#DiskTypesScopedList.Warning}

[Output Only] Informational warning which replaces the list of disk types
when the list is empty.
[Output Only] Informational warning which replaces the list of disk types
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#DiskTypesScopedList.Warning.code) = 1
* `repeated` [`compute.DiskTypesScopedList.Warning.Data`](gcp_compute.md#DiskTypesScopedList.Warning.Data) [`data`](#DiskTypesScopedList.Warning.data) = 2
* `string` [`message`](#DiskTypesScopedList.Warning.message) = 3

### `code` {#DiskTypesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#DiskTypesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.DiskTypesScopedList.Warning.Data`](gcp_compute.md#DiskTypesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#DiskTypesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RegionInstanceGroupsListInstances.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RegionInstanceGroupsListInstances.Warning.code) = 1
* `repeated` [`compute.RegionInstanceGroupsListInstances.Warning.Data`](gcp_compute.md#RegionInstanceGroupsListInstances.Warning.Data) [`data`](#RegionInstanceGroupsListInstances.Warning.data) = 2
* `string` [`message`](#RegionInstanceGroupsListInstances.Warning.message) = 3

### `code` {#RegionInstanceGroupsListInstances.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RegionInstanceGroupsListInstances.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RegionInstanceGroupsListInstances.Warning.Data`](gcp_compute.md#RegionInstanceGroupsListInstances.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RegionInstanceGroupsListInstances.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetInstancesScopedList.Warning}

Informational warning which replaces the list of addresses when the list is
empty.
Informational warning which replaces the list of addresses when the list is
empty.


### Inputs for `Warning`

* `string` [`code`](#TargetInstancesScopedList.Warning.code) = 1
* `repeated` [`compute.TargetInstancesScopedList.Warning.Data`](gcp_compute.md#TargetInstancesScopedList.Warning.Data) [`data`](#TargetInstancesScopedList.Warning.data) = 2
* `string` [`message`](#TargetInstancesScopedList.Warning.message) = 3

### `code` {#TargetInstancesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetInstancesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetInstancesScopedList.Warning.Data`](gcp_compute.md#TargetInstancesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetInstancesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#DiskTypeList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#DiskTypeList.Warning.code) = 1
* `repeated` [`compute.DiskTypeList.Warning.Data`](gcp_compute.md#DiskTypeList.Warning.Data) [`data`](#DiskTypeList.Warning.data) = 2
* `string` [`message`](#DiskTypeList.Warning.message) = 3

### `code` {#DiskTypeList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#DiskTypeList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.DiskTypeList.Warning.Data`](gcp_compute.md#DiskTypeList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#DiskTypeList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InterconnectAttachmentAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InterconnectAttachmentAggregatedList.Warning.code) = 1
* `repeated` [`compute.InterconnectAttachmentAggregatedList.Warning.Data`](gcp_compute.md#InterconnectAttachmentAggregatedList.Warning.Data) [`data`](#InterconnectAttachmentAggregatedList.Warning.data) = 2
* `string` [`message`](#InterconnectAttachmentAggregatedList.Warning.message) = 3

### `code` {#InterconnectAttachmentAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InterconnectAttachmentAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InterconnectAttachmentAggregatedList.Warning.Data`](gcp_compute.md#InterconnectAttachmentAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InterconnectAttachmentAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RegionInstanceGroupManagerList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RegionInstanceGroupManagerList.Warning.code) = 1
* `repeated` [`compute.RegionInstanceGroupManagerList.Warning.Data`](gcp_compute.md#RegionInstanceGroupManagerList.Warning.Data) [`data`](#RegionInstanceGroupManagerList.Warning.data) = 2
* `string` [`message`](#RegionInstanceGroupManagerList.Warning.message) = 3

### `code` {#RegionInstanceGroupManagerList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RegionInstanceGroupManagerList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RegionInstanceGroupManagerList.Warning.Data`](gcp_compute.md#RegionInstanceGroupManagerList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RegionInstanceGroupManagerList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupManagersScopedList.Warning}

[Output Only] The warning that replaces the list of managed instance groups
when the list is empty.
[Output Only] The warning that replaces the list of managed instance groups
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupManagersScopedList.Warning.code) = 1
* `repeated` [`compute.InstanceGroupManagersScopedList.Warning.Data`](gcp_compute.md#InstanceGroupManagersScopedList.Warning.Data) [`data`](#InstanceGroupManagersScopedList.Warning.data) = 2
* `string` [`message`](#InstanceGroupManagersScopedList.Warning.message) = 3

### `code` {#InstanceGroupManagersScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupManagersScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupManagersScopedList.Warning.Data`](gcp_compute.md#InstanceGroupManagersScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupManagersScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetPoolAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetPoolAggregatedList.Warning.code) = 1
* `repeated` [`compute.TargetPoolAggregatedList.Warning.Data`](gcp_compute.md#TargetPoolAggregatedList.Warning.Data) [`data`](#TargetPoolAggregatedList.Warning.data) = 2
* `string` [`message`](#TargetPoolAggregatedList.Warning.message) = 3

### `code` {#TargetPoolAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetPoolAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetPoolAggregatedList.Warning.Data`](gcp_compute.md#TargetPoolAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetPoolAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#DiskList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#DiskList.Warning.code) = 1
* `repeated` [`compute.DiskList.Warning.Data`](gcp_compute.md#DiskList.Warning.Data) [`data`](#DiskList.Warning.data) = 2
* `string` [`message`](#DiskList.Warning.message) = 3

### `code` {#DiskList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#DiskList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.DiskList.Warning.Data`](gcp_compute.md#DiskList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#DiskList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InterconnectAttachmentList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InterconnectAttachmentList.Warning.code) = 1
* `repeated` [`compute.InterconnectAttachmentList.Warning.Data`](gcp_compute.md#InterconnectAttachmentList.Warning.Data) [`data`](#InterconnectAttachmentList.Warning.data) = 2
* `string` [`message`](#InterconnectAttachmentList.Warning.message) = 3

### `code` {#InterconnectAttachmentList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InterconnectAttachmentList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InterconnectAttachmentList.Warning.Data`](gcp_compute.md#InterconnectAttachmentList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InterconnectAttachmentList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RegionInstanceGroupList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RegionInstanceGroupList.Warning.code) = 1
* `repeated` [`compute.RegionInstanceGroupList.Warning.Data`](gcp_compute.md#RegionInstanceGroupList.Warning.Data) [`data`](#RegionInstanceGroupList.Warning.data) = 2
* `string` [`message`](#RegionInstanceGroupList.Warning.message) = 3

### `code` {#RegionInstanceGroupList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RegionInstanceGroupList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RegionInstanceGroupList.Warning.Data`](gcp_compute.md#RegionInstanceGroupList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RegionInstanceGroupList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetPoolList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetPoolList.Warning.code) = 1
* `repeated` [`compute.TargetPoolList.Warning.Data`](gcp_compute.md#TargetPoolList.Warning.Data) [`data`](#TargetPoolList.Warning.data) = 2
* `string` [`message`](#TargetPoolList.Warning.message) = 3

### `code` {#TargetPoolList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetPoolList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetPoolList.Warning.Data`](gcp_compute.md#TargetPoolList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetPoolList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#DiskAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#DiskAggregatedList.Warning.code) = 1
* `repeated` [`compute.DiskAggregatedList.Warning.Data`](gcp_compute.md#DiskAggregatedList.Warning.Data) [`data`](#DiskAggregatedList.Warning.data) = 2
* `string` [`message`](#DiskAggregatedList.Warning.message) = 3

### `code` {#DiskAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#DiskAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.DiskAggregatedList.Warning.Data`](gcp_compute.md#DiskAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#DiskAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#LicensesListResponse.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#LicensesListResponse.Warning.code) = 1
* `repeated` [`compute.LicensesListResponse.Warning.Data`](gcp_compute.md#LicensesListResponse.Warning.Data) [`data`](#LicensesListResponse.Warning.data) = 2
* `string` [`message`](#LicensesListResponse.Warning.message) = 3

### `code` {#LicensesListResponse.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#LicensesListResponse.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.LicensesListResponse.Warning.Data`](gcp_compute.md#LicensesListResponse.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#LicensesListResponse.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InterconnectAttachmentsScopedList.Warning}

Informational warning which replaces the list of addresses when the list is
empty.
Informational warning which replaces the list of addresses when the list is
empty.


### Inputs for `Warning`

* `string` [`code`](#InterconnectAttachmentsScopedList.Warning.code) = 1
* `repeated` [`compute.InterconnectAttachmentsScopedList.Warning.Data`](gcp_compute.md#InterconnectAttachmentsScopedList.Warning.Data) [`data`](#InterconnectAttachmentsScopedList.Warning.data) = 2
* `string` [`message`](#InterconnectAttachmentsScopedList.Warning.message) = 3

### `code` {#InterconnectAttachmentsScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InterconnectAttachmentsScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InterconnectAttachmentsScopedList.Warning.Data`](gcp_compute.md#InterconnectAttachmentsScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InterconnectAttachmentsScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RegionDiskTypeList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RegionDiskTypeList.Warning.code) = 1
* `repeated` [`compute.RegionDiskTypeList.Warning.Data`](gcp_compute.md#RegionDiskTypeList.Warning.Data) [`data`](#RegionDiskTypeList.Warning.data) = 2
* `string` [`message`](#RegionDiskTypeList.Warning.message) = 3

### `code` {#RegionDiskTypeList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RegionDiskTypeList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RegionDiskTypeList.Warning.Data`](gcp_compute.md#RegionDiskTypeList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RegionDiskTypeList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InterconnectList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InterconnectList.Warning.code) = 1
* `repeated` [`compute.InterconnectList.Warning.Data`](gcp_compute.md#InterconnectList.Warning.Data) [`data`](#InterconnectList.Warning.data) = 2
* `string` [`message`](#InterconnectList.Warning.message) = 3

### `code` {#InterconnectList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InterconnectList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InterconnectList.Warning.Data`](gcp_compute.md#InterconnectList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InterconnectList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#CommitmentsScopedList.Warning}

[Output Only] Informational warning which replaces the list of commitments
when the list is empty.
[Output Only] Informational warning which replaces the list of commitments
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#CommitmentsScopedList.Warning.code) = 1
* `repeated` [`compute.CommitmentsScopedList.Warning.Data`](gcp_compute.md#CommitmentsScopedList.Warning.Data) [`data`](#CommitmentsScopedList.Warning.data) = 2
* `string` [`message`](#CommitmentsScopedList.Warning.message) = 3

### `code` {#CommitmentsScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#CommitmentsScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.CommitmentsScopedList.Warning.Data`](gcp_compute.md#CommitmentsScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#CommitmentsScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetPoolsScopedList.Warning}

Informational warning which replaces the list of addresses when the list is
empty.
Informational warning which replaces the list of addresses when the list is
empty.


### Inputs for `Warning`

* `string` [`code`](#TargetPoolsScopedList.Warning.code) = 1
* `repeated` [`compute.TargetPoolsScopedList.Warning.Data`](gcp_compute.md#TargetPoolsScopedList.Warning.Data) [`data`](#TargetPoolsScopedList.Warning.data) = 2
* `string` [`message`](#TargetPoolsScopedList.Warning.message) = 3

### `code` {#TargetPoolsScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetPoolsScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetPoolsScopedList.Warning.Data`](gcp_compute.md#TargetPoolsScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetPoolsScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#CommitmentList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#CommitmentList.Warning.code) = 1
* `repeated` [`compute.CommitmentList.Warning.Data`](gcp_compute.md#CommitmentList.Warning.Data) [`data`](#CommitmentList.Warning.data) = 2
* `string` [`message`](#CommitmentList.Warning.message) = 3

### `code` {#CommitmentList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#CommitmentList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.CommitmentList.Warning.Data`](gcp_compute.md#CommitmentList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#CommitmentList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AcceleratorTypesScopedList.Warning}

[Output Only] An informational warning that appears when the accelerator
types list is empty.
[Output Only] An informational warning that appears when the accelerator
types list is empty.


### Inputs for `Warning`

* `string` [`code`](#AcceleratorTypesScopedList.Warning.code) = 1
* `repeated` [`compute.AcceleratorTypesScopedList.Warning.Data`](gcp_compute.md#AcceleratorTypesScopedList.Warning.Data) [`data`](#AcceleratorTypesScopedList.Warning.data) = 2
* `string` [`message`](#AcceleratorTypesScopedList.Warning.message) = 3

### `code` {#AcceleratorTypesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AcceleratorTypesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AcceleratorTypesScopedList.Warning.Data`](gcp_compute.md#AcceleratorTypesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AcceleratorTypesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RegionAutoscalerList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RegionAutoscalerList.Warning.code) = 1
* `repeated` [`compute.RegionAutoscalerList.Warning.Data`](gcp_compute.md#RegionAutoscalerList.Warning.Data) [`data`](#RegionAutoscalerList.Warning.data) = 2
* `string` [`message`](#RegionAutoscalerList.Warning.message) = 3

### `code` {#RegionAutoscalerList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RegionAutoscalerList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RegionAutoscalerList.Warning.Data`](gcp_compute.md#RegionAutoscalerList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RegionAutoscalerList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#MachineTypeAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#MachineTypeAggregatedList.Warning.code) = 1
* `repeated` [`compute.MachineTypeAggregatedList.Warning.Data`](gcp_compute.md#MachineTypeAggregatedList.Warning.Data) [`data`](#MachineTypeAggregatedList.Warning.data) = 2
* `string` [`message`](#MachineTypeAggregatedList.Warning.message) = 3

### `code` {#MachineTypeAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#MachineTypeAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.MachineTypeAggregatedList.Warning.Data`](gcp_compute.md#MachineTypeAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#MachineTypeAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#MachineTypeList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#MachineTypeList.Warning.code) = 1
* `repeated` [`compute.MachineTypeList.Warning.Data`](gcp_compute.md#MachineTypeList.Warning.Data) [`data`](#MachineTypeList.Warning.data) = 2
* `string` [`message`](#MachineTypeList.Warning.message) = 3

### `code` {#MachineTypeList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#MachineTypeList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.MachineTypeList.Warning.Data`](gcp_compute.md#MachineTypeList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#MachineTypeList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#MachineTypesScopedList.Warning}

[Output Only] An informational warning that appears when the machine types
list is empty.
[Output Only] An informational warning that appears when the machine types
list is empty.


### Inputs for `Warning`

* `string` [`code`](#MachineTypesScopedList.Warning.code) = 1
* `repeated` [`compute.MachineTypesScopedList.Warning.Data`](gcp_compute.md#MachineTypesScopedList.Warning.Data) [`data`](#MachineTypesScopedList.Warning.data) = 2
* `string` [`message`](#MachineTypesScopedList.Warning.message) = 3

### `code` {#MachineTypesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#MachineTypesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.MachineTypesScopedList.Warning.Data`](gcp_compute.md#MachineTypesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#MachineTypesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#NetworkList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#NetworkList.Warning.code) = 1
* `repeated` [`compute.NetworkList.Warning.Data`](gcp_compute.md#NetworkList.Warning.Data) [`data`](#NetworkList.Warning.data) = 2
* `string` [`message`](#NetworkList.Warning.message) = 3

### `code` {#NetworkList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#NetworkList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.NetworkList.Warning.Data`](gcp_compute.md#NetworkList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#NetworkList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetSslProxyList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetSslProxyList.Warning.code) = 1
* `repeated` [`compute.TargetSslProxyList.Warning.Data`](gcp_compute.md#TargetSslProxyList.Warning.Data) [`data`](#TargetSslProxyList.Warning.data) = 2
* `string` [`message`](#TargetSslProxyList.Warning.message) = 3

### `code` {#TargetSslProxyList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetSslProxyList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetSslProxyList.Warning.Data`](gcp_compute.md#TargetSslProxyList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetSslProxyList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#CommitmentAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#CommitmentAggregatedList.Warning.code) = 1
* `repeated` [`compute.CommitmentAggregatedList.Warning.Data`](gcp_compute.md#CommitmentAggregatedList.Warning.Data) [`data`](#CommitmentAggregatedList.Warning.data) = 2
* `string` [`message`](#CommitmentAggregatedList.Warning.message) = 3

### `code` {#CommitmentAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#CommitmentAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.CommitmentAggregatedList.Warning.Data`](gcp_compute.md#CommitmentAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#CommitmentAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AcceleratorTypeList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#AcceleratorTypeList.Warning.code) = 1
* `repeated` [`compute.AcceleratorTypeList.Warning.Data`](gcp_compute.md#AcceleratorTypeList.Warning.Data) [`data`](#AcceleratorTypeList.Warning.data) = 2
* `string` [`message`](#AcceleratorTypeList.Warning.message) = 3

### `code` {#AcceleratorTypeList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AcceleratorTypeList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AcceleratorTypeList.Warning.Data`](gcp_compute.md#AcceleratorTypeList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AcceleratorTypeList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#OperationAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#OperationAggregatedList.Warning.code) = 1
* `repeated` [`compute.OperationAggregatedList.Warning.Data`](gcp_compute.md#OperationAggregatedList.Warning.Data) [`data`](#OperationAggregatedList.Warning.data) = 2
* `string` [`message`](#OperationAggregatedList.Warning.message) = 3

### `code` {#OperationAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#OperationAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.OperationAggregatedList.Warning.Data`](gcp_compute.md#OperationAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#OperationAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RouterList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#RouterList.Warning.code) = 1
* `repeated` [`compute.RouterList.Warning.Data`](gcp_compute.md#RouterList.Warning.Data) [`data`](#RouterList.Warning.data) = 2
* `string` [`message`](#RouterList.Warning.message) = 3

### `code` {#RouterList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RouterList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RouterList.Warning.Data`](gcp_compute.md#RouterList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RouterList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#ForwardingRulesScopedList.Warning}

Informational warning which replaces the list of forwarding rules when the
list is empty.
Informational warning which replaces the list of forwarding rules when the
list is empty.


### Inputs for `Warning`

* `string` [`code`](#ForwardingRulesScopedList.Warning.code) = 1
* `repeated` [`compute.ForwardingRulesScopedList.Warning.Data`](gcp_compute.md#ForwardingRulesScopedList.Warning.Data) [`data`](#ForwardingRulesScopedList.Warning.data) = 2
* `string` [`message`](#ForwardingRulesScopedList.Warning.message) = 3

### `code` {#ForwardingRulesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#ForwardingRulesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.ForwardingRulesScopedList.Warning.Data`](gcp_compute.md#ForwardingRulesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#ForwardingRulesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetTcpProxyList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetTcpProxyList.Warning.code) = 1
* `repeated` [`compute.TargetTcpProxyList.Warning.Data`](gcp_compute.md#TargetTcpProxyList.Warning.Data) [`data`](#TargetTcpProxyList.Warning.data) = 2
* `string` [`message`](#TargetTcpProxyList.Warning.message) = 3

### `code` {#TargetTcpProxyList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetTcpProxyList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetTcpProxyList.Warning.Data`](gcp_compute.md#TargetTcpProxyList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetTcpProxyList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#BackendServiceList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#BackendServiceList.Warning.code) = 1
* `repeated` [`compute.BackendServiceList.Warning.Data`](gcp_compute.md#BackendServiceList.Warning.Data) [`data`](#BackendServiceList.Warning.data) = 2
* `string` [`message`](#BackendServiceList.Warning.message) = 3

### `code` {#BackendServiceList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#BackendServiceList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.BackendServiceList.Warning.Data`](gcp_compute.md#BackendServiceList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#BackendServiceList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InterconnectLocationList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InterconnectLocationList.Warning.code) = 1
* `repeated` [`compute.InterconnectLocationList.Warning.Data`](gcp_compute.md#InterconnectLocationList.Warning.Data) [`data`](#InterconnectLocationList.Warning.data) = 2
* `string` [`message`](#InterconnectLocationList.Warning.message) = 3

### `code` {#InterconnectLocationList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InterconnectLocationList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InterconnectLocationList.Warning.Data`](gcp_compute.md#InterconnectLocationList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InterconnectLocationList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#SubnetworksScopedList.Warning}

An informational warning that appears when the list of addresses is empty.
An informational warning that appears when the list of addresses is empty.


### Inputs for `Warning`

* `string` [`code`](#SubnetworksScopedList.Warning.code) = 1
* `repeated` [`compute.SubnetworksScopedList.Warning.Data`](gcp_compute.md#SubnetworksScopedList.Warning.Data) [`data`](#SubnetworksScopedList.Warning.data) = 2
* `string` [`message`](#SubnetworksScopedList.Warning.message) = 3

### `code` {#SubnetworksScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SubnetworksScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SubnetworksScopedList.Warning.Data`](gcp_compute.md#SubnetworksScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SubnetworksScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#BackendServiceAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#BackendServiceAggregatedList.Warning.code) = 1
* `repeated` [`compute.BackendServiceAggregatedList.Warning.Data`](gcp_compute.md#BackendServiceAggregatedList.Warning.Data) [`data`](#BackendServiceAggregatedList.Warning.data) = 2
* `string` [`message`](#BackendServiceAggregatedList.Warning.message) = 3

### `code` {#BackendServiceAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#BackendServiceAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.BackendServiceAggregatedList.Warning.Data`](gcp_compute.md#BackendServiceAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#BackendServiceAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetVpnGatewayAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetVpnGatewayAggregatedList.Warning.code) = 1
* `repeated` [`compute.TargetVpnGatewayAggregatedList.Warning.Data`](gcp_compute.md#TargetVpnGatewayAggregatedList.Warning.Data) [`data`](#TargetVpnGatewayAggregatedList.Warning.data) = 2
* `string` [`message`](#TargetVpnGatewayAggregatedList.Warning.message) = 3

### `code` {#TargetVpnGatewayAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetVpnGatewayAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetVpnGatewayAggregatedList.Warning.Data`](gcp_compute.md#TargetVpnGatewayAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetVpnGatewayAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#BackendBucketList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#BackendBucketList.Warning.code) = 1
* `repeated` [`compute.BackendBucketList.Warning.Data`](gcp_compute.md#BackendBucketList.Warning.Data) [`data`](#BackendBucketList.Warning.data) = 2
* `string` [`message`](#BackendBucketList.Warning.message) = 3

### `code` {#BackendBucketList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#BackendBucketList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.BackendBucketList.Warning.Data`](gcp_compute.md#BackendBucketList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#BackendBucketList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupManagerList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupManagerList.Warning.code) = 1
* `repeated` [`compute.InstanceGroupManagerList.Warning.Data`](gcp_compute.md#InstanceGroupManagerList.Warning.Data) [`data`](#InstanceGroupManagerList.Warning.data) = 2
* `string` [`message`](#InstanceGroupManagerList.Warning.message) = 3

### `code` {#InstanceGroupManagerList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupManagerList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupManagerList.Warning.Data`](gcp_compute.md#InstanceGroupManagerList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupManagerList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetVpnGatewayList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#TargetVpnGatewayList.Warning.code) = 1
* `repeated` [`compute.TargetVpnGatewayList.Warning.Data`](gcp_compute.md#TargetVpnGatewayList.Warning.Data) [`data`](#TargetVpnGatewayList.Warning.data) = 2
* `string` [`message`](#TargetVpnGatewayList.Warning.message) = 3

### `code` {#TargetVpnGatewayList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetVpnGatewayList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetVpnGatewayList.Warning.Data`](gcp_compute.md#TargetVpnGatewayList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetVpnGatewayList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AutoscalersScopedList.Warning}

[Output Only] Informational warning which replaces the list of autoscalers
when the list is empty.
[Output Only] Informational warning which replaces the list of autoscalers
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#AutoscalersScopedList.Warning.code) = 1
* `repeated` [`compute.AutoscalersScopedList.Warning.Data`](gcp_compute.md#AutoscalersScopedList.Warning.Data) [`data`](#AutoscalersScopedList.Warning.data) = 2
* `string` [`message`](#AutoscalersScopedList.Warning.message) = 3

### `code` {#AutoscalersScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AutoscalersScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AutoscalersScopedList.Warning.Data`](gcp_compute.md#AutoscalersScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AutoscalersScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#RoutersScopedList.Warning}

Informational warning which replaces the list of routers when the list is
empty.
Informational warning which replaces the list of routers when the list is
empty.


### Inputs for `Warning`

* `string` [`code`](#RoutersScopedList.Warning.code) = 1
* `repeated` [`compute.RoutersScopedList.Warning.Data`](gcp_compute.md#RoutersScopedList.Warning.Data) [`data`](#RoutersScopedList.Warning.data) = 2
* `string` [`message`](#RoutersScopedList.Warning.message) = 3

### `code` {#RoutersScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#RoutersScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.RoutersScopedList.Warning.Data`](gcp_compute.md#RoutersScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#RoutersScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#TargetVpnGatewaysScopedList.Warning}

[Output Only] Informational warning which replaces the list of addresses
when the list is empty.
[Output Only] Informational warning which replaces the list of addresses
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#TargetVpnGatewaysScopedList.Warning.code) = 1
* `repeated` [`compute.TargetVpnGatewaysScopedList.Warning.Data`](gcp_compute.md#TargetVpnGatewaysScopedList.Warning.Data) [`data`](#TargetVpnGatewaysScopedList.Warning.data) = 2
* `string` [`message`](#TargetVpnGatewaysScopedList.Warning.message) = 3

### `code` {#TargetVpnGatewaysScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#TargetVpnGatewaysScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.TargetVpnGatewaysScopedList.Warning.Data`](gcp_compute.md#TargetVpnGatewaysScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#TargetVpnGatewaysScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AutoscalerList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#AutoscalerList.Warning.code) = 1
* `repeated` [`compute.AutoscalerList.Warning.Data`](gcp_compute.md#AutoscalerList.Warning.Data) [`data`](#AutoscalerList.Warning.data) = 2
* `string` [`message`](#AutoscalerList.Warning.message) = 3

### `code` {#AutoscalerList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AutoscalerList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AutoscalerList.Warning.Data`](gcp_compute.md#AutoscalerList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AutoscalerList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupManagerAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupManagerAggregatedList.Warning.code) = 1
* `repeated` [`compute.InstanceGroupManagerAggregatedList.Warning.Data`](gcp_compute.md#InstanceGroupManagerAggregatedList.Warning.Data) [`data`](#InstanceGroupManagerAggregatedList.Warning.data) = 2
* `string` [`message`](#InstanceGroupManagerAggregatedList.Warning.message) = 3

### `code` {#InstanceGroupManagerAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupManagerAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupManagerAggregatedList.Warning.Data`](gcp_compute.md#InstanceGroupManagerAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupManagerAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#BackendServicesScopedList.Warning}

Informational warning which replaces the list of backend services when the
list is empty.
Informational warning which replaces the list of backend services when the
list is empty.


### Inputs for `Warning`

* `string` [`code`](#BackendServicesScopedList.Warning.code) = 1
* `repeated` [`compute.BackendServicesScopedList.Warning.Data`](gcp_compute.md#BackendServicesScopedList.Warning.Data) [`data`](#BackendServicesScopedList.Warning.data) = 2
* `string` [`message`](#BackendServicesScopedList.Warning.message) = 3

### `code` {#BackendServicesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#BackendServicesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.BackendServicesScopedList.Warning.Data`](gcp_compute.md#BackendServicesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#BackendServicesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#SubnetworkList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#SubnetworkList.Warning.code) = 1
* `repeated` [`compute.SubnetworkList.Warning.Data`](gcp_compute.md#SubnetworkList.Warning.Data) [`data`](#SubnetworkList.Warning.data) = 2
* `string` [`message`](#SubnetworkList.Warning.message) = 3

### `code` {#SubnetworkList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SubnetworkList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SubnetworkList.Warning.Data`](gcp_compute.md#SubnetworkList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SubnetworkList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupList.Warning.code) = 1
* `repeated` [`compute.InstanceGroupList.Warning.Data`](gcp_compute.md#InstanceGroupList.Warning.Data) [`data`](#InstanceGroupList.Warning.data) = 2
* `string` [`message`](#InstanceGroupList.Warning.message) = 3

### `code` {#InstanceGroupList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupList.Warning.Data`](gcp_compute.md#InstanceGroupList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#HealthCheckList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#HealthCheckList.Warning.code) = 1
* `repeated` [`compute.HealthCheckList.Warning.Data`](gcp_compute.md#HealthCheckList.Warning.Data) [`data`](#HealthCheckList.Warning.data) = 2
* `string` [`message`](#HealthCheckList.Warning.message) = 3

### `code` {#HealthCheckList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#HealthCheckList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.HealthCheckList.Warning.Data`](gcp_compute.md#HealthCheckList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#HealthCheckList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#UrlMapList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#UrlMapList.Warning.code) = 1
* `repeated` [`compute.UrlMapList.Warning.Data`](gcp_compute.md#UrlMapList.Warning.Data) [`data`](#UrlMapList.Warning.data) = 2
* `string` [`message`](#UrlMapList.Warning.message) = 3

### `code` {#UrlMapList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#UrlMapList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.UrlMapList.Warning.Data`](gcp_compute.md#UrlMapList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#UrlMapList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AutoscalerAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#AutoscalerAggregatedList.Warning.code) = 1
* `repeated` [`compute.AutoscalerAggregatedList.Warning.Data`](gcp_compute.md#AutoscalerAggregatedList.Warning.Data) [`data`](#AutoscalerAggregatedList.Warning.data) = 2
* `string` [`message`](#AutoscalerAggregatedList.Warning.message) = 3

### `code` {#AutoscalerAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AutoscalerAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AutoscalerAggregatedList.Warning.Data`](gcp_compute.md#AutoscalerAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AutoscalerAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#SubnetworkAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#SubnetworkAggregatedList.Warning.code) = 1
* `repeated` [`compute.SubnetworkAggregatedList.Warning.Data`](gcp_compute.md#SubnetworkAggregatedList.Warning.Data) [`data`](#SubnetworkAggregatedList.Warning.data) = 2
* `string` [`message`](#SubnetworkAggregatedList.Warning.message) = 3

### `code` {#SubnetworkAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SubnetworkAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SubnetworkAggregatedList.Warning.Data`](gcp_compute.md#SubnetworkAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SubnetworkAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#HttpHealthCheckList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#HttpHealthCheckList.Warning.code) = 1
* `repeated` [`compute.HttpHealthCheckList.Warning.Data`](gcp_compute.md#HttpHealthCheckList.Warning.Data) [`data`](#HttpHealthCheckList.Warning.data) = 2
* `string` [`message`](#HttpHealthCheckList.Warning.message) = 3

### `code` {#HttpHealthCheckList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#HttpHealthCheckList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.HttpHealthCheckList.Warning.Data`](gcp_compute.md#HttpHealthCheckList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#HttpHealthCheckList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#OperationList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#OperationList.Warning.code) = 1
* `repeated` [`compute.OperationList.Warning.Data`](gcp_compute.md#OperationList.Warning.Data) [`data`](#OperationList.Warning.data) = 2
* `string` [`message`](#OperationList.Warning.message) = 3

### `code` {#OperationList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#OperationList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.OperationList.Warning.Data`](gcp_compute.md#OperationList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#OperationList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#SnapshotList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#SnapshotList.Warning.code) = 1
* `repeated` [`compute.SnapshotList.Warning.Data`](gcp_compute.md#SnapshotList.Warning.Data) [`data`](#SnapshotList.Warning.data) = 2
* `string` [`message`](#SnapshotList.Warning.message) = 3

### `code` {#SnapshotList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SnapshotList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SnapshotList.Warning.Data`](gcp_compute.md#SnapshotList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SnapshotList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceGroupAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceGroupAggregatedList.Warning.code) = 1
* `repeated` [`compute.InstanceGroupAggregatedList.Warning.Data`](gcp_compute.md#InstanceGroupAggregatedList.Warning.Data) [`data`](#InstanceGroupAggregatedList.Warning.data) = 2
* `string` [`message`](#InstanceGroupAggregatedList.Warning.message) = 3

### `code` {#InstanceGroupAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceGroupAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceGroupAggregatedList.Warning.Data`](gcp_compute.md#InstanceGroupAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceGroupAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#SslCertificateList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#SslCertificateList.Warning.code) = 1
* `repeated` [`compute.SslCertificateList.Warning.Data`](gcp_compute.md#SslCertificateList.Warning.Data) [`data`](#SslCertificateList.Warning.data) = 2
* `string` [`message`](#SslCertificateList.Warning.message) = 3

### `code` {#SslCertificateList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SslCertificateList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SslCertificateList.Warning.Data`](gcp_compute.md#SslCertificateList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SslCertificateList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#HttpsHealthCheckList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#HttpsHealthCheckList.Warning.code) = 1
* `repeated` [`compute.HttpsHealthCheckList.Warning.Data`](gcp_compute.md#HttpsHealthCheckList.Warning.Data) [`data`](#HttpsHealthCheckList.Warning.data) = 2
* `string` [`message`](#HttpsHealthCheckList.Warning.message) = 3

### `code` {#HttpsHealthCheckList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#HttpsHealthCheckList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.HttpsHealthCheckList.Warning.Data`](gcp_compute.md#HttpsHealthCheckList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#HttpsHealthCheckList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AcceleratorTypeAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#AcceleratorTypeAggregatedList.Warning.code) = 1
* `repeated` [`compute.AcceleratorTypeAggregatedList.Warning.Data`](gcp_compute.md#AcceleratorTypeAggregatedList.Warning.Data) [`data`](#AcceleratorTypeAggregatedList.Warning.data) = 2
* `string` [`message`](#AcceleratorTypeAggregatedList.Warning.message) = 3

### `code` {#AcceleratorTypeAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AcceleratorTypeAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AcceleratorTypeAggregatedList.Warning.Data`](gcp_compute.md#AcceleratorTypeAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AcceleratorTypeAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AddressesScopedList.Warning}

[Output Only] Informational warning which replaces the list of addresses
when the list is empty.
[Output Only] Informational warning which replaces the list of addresses
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#AddressesScopedList.Warning.code) = 1
* `repeated` [`compute.AddressesScopedList.Warning.Data`](gcp_compute.md#AddressesScopedList.Warning.Data) [`data`](#AddressesScopedList.Warning.data) = 2
* `string` [`message`](#AddressesScopedList.Warning.message) = 3

### `code` {#AddressesScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AddressesScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AddressesScopedList.Warning.Data`](gcp_compute.md#AddressesScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AddressesScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#VpnTunnelAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#VpnTunnelAggregatedList.Warning.code) = 1
* `repeated` [`compute.VpnTunnelAggregatedList.Warning.Data`](gcp_compute.md#VpnTunnelAggregatedList.Warning.Data) [`data`](#VpnTunnelAggregatedList.Warning.data) = 2
* `string` [`message`](#VpnTunnelAggregatedList.Warning.message) = 3

### `code` {#VpnTunnelAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#VpnTunnelAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.VpnTunnelAggregatedList.Warning.Data`](gcp_compute.md#VpnTunnelAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#VpnTunnelAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AddressList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#AddressList.Warning.code) = 1
* `repeated` [`compute.AddressList.Warning.Data`](gcp_compute.md#AddressList.Warning.Data) [`data`](#AddressList.Warning.data) = 2
* `string` [`message`](#AddressList.Warning.message) = 3

### `code` {#AddressList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AddressList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AddressList.Warning.Data`](gcp_compute.md#AddressList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AddressList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#InstanceAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#InstanceAggregatedList.Warning.code) = 1
* `repeated` [`compute.InstanceAggregatedList.Warning.Data`](gcp_compute.md#InstanceAggregatedList.Warning.Data) [`data`](#InstanceAggregatedList.Warning.data) = 2
* `string` [`message`](#InstanceAggregatedList.Warning.message) = 3

### `code` {#InstanceAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#InstanceAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.InstanceAggregatedList.Warning.Data`](gcp_compute.md#InstanceAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#InstanceAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#VpnTunnelList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#VpnTunnelList.Warning.code) = 1
* `repeated` [`compute.VpnTunnelList.Warning.Data`](gcp_compute.md#VpnTunnelList.Warning.Data) [`data`](#VpnTunnelList.Warning.data) = 2
* `string` [`message`](#VpnTunnelList.Warning.message) = 3

### `code` {#VpnTunnelList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#VpnTunnelList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.VpnTunnelList.Warning.Data`](gcp_compute.md#VpnTunnelList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#VpnTunnelList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#AddressAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#AddressAggregatedList.Warning.code) = 1
* `repeated` [`compute.AddressAggregatedList.Warning.Data`](gcp_compute.md#AddressAggregatedList.Warning.Data) [`data`](#AddressAggregatedList.Warning.data) = 2
* `string` [`message`](#AddressAggregatedList.Warning.message) = 3

### `code` {#AddressAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#AddressAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.AddressAggregatedList.Warning.Data`](gcp_compute.md#AddressAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#AddressAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#SslPoliciesList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#SslPoliciesList.Warning.code) = 1
* `repeated` [`compute.SslPoliciesList.Warning.Data`](gcp_compute.md#SslPoliciesList.Warning.Data) [`data`](#SslPoliciesList.Warning.data) = 2
* `string` [`message`](#SslPoliciesList.Warning.message) = 3

### `code` {#SslPoliciesList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SslPoliciesList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SslPoliciesList.Warning.Data`](gcp_compute.md#SslPoliciesList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SslPoliciesList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#VpnTunnelsScopedList.Warning}

Informational warning which replaces the list of addresses when the list is
empty.
Informational warning which replaces the list of addresses when the list is
empty.


### Inputs for `Warning`

* `string` [`code`](#VpnTunnelsScopedList.Warning.code) = 1
* `repeated` [`compute.VpnTunnelsScopedList.Warning.Data`](gcp_compute.md#VpnTunnelsScopedList.Warning.Data) [`data`](#VpnTunnelsScopedList.Warning.data) = 2
* `string` [`message`](#VpnTunnelsScopedList.Warning.message) = 3

### `code` {#VpnTunnelsScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#VpnTunnelsScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.VpnTunnelsScopedList.Warning.Data`](gcp_compute.md#VpnTunnelsScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#VpnTunnelsScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#DiskTypeAggregatedList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#DiskTypeAggregatedList.Warning.code) = 1
* `repeated` [`compute.DiskTypeAggregatedList.Warning.Data`](gcp_compute.md#DiskTypeAggregatedList.Warning.Data) [`data`](#DiskTypeAggregatedList.Warning.data) = 2
* `string` [`message`](#DiskTypeAggregatedList.Warning.message) = 3

### `code` {#DiskTypeAggregatedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#DiskTypeAggregatedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.DiskTypeAggregatedList.Warning.Data`](gcp_compute.md#DiskTypeAggregatedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#DiskTypeAggregatedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#XpnHostList.Warning}

[Output Only] Informational warning message.
[Output Only] Informational warning message.


### Inputs for `Warning`

* `string` [`code`](#XpnHostList.Warning.code) = 1
* `repeated` [`compute.XpnHostList.Warning.Data`](gcp_compute.md#XpnHostList.Warning.Data) [`data`](#XpnHostList.Warning.data) = 2
* `string` [`message`](#XpnHostList.Warning.message) = 3

### `code` {#XpnHostList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#XpnHostList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.XpnHostList.Warning.Data`](gcp_compute.md#XpnHostList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#XpnHostList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warning` {#OperationsScopedList.Warning}

[Output Only] Informational warning which replaces the list of operations
when the list is empty.
[Output Only] Informational warning which replaces the list of operations
when the list is empty.


### Inputs for `Warning`

* `string` [`code`](#OperationsScopedList.Warning.code) = 1
* `repeated` [`compute.OperationsScopedList.Warning.Data`](gcp_compute.md#OperationsScopedList.Warning.Data) [`data`](#OperationsScopedList.Warning.data) = 2
* `string` [`message`](#OperationsScopedList.Warning.message) = 3

### `code` {#OperationsScopedList.Warning.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#OperationsScopedList.Warning.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.OperationsScopedList.Warning.Data`](gcp_compute.md#OperationsScopedList.Warning.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#OperationsScopedList.Warning.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warnings` {#Operation.Warnings}

[Output Only] If warning messages are generated during processing of the
operation, this field will be populated.


### Inputs for `Warnings`

* `string` [`code`](#Operation.Warnings.code) = 1
* `repeated` [`compute.Operation.Warnings.Data`](gcp_compute.md#Operation.Warnings.Data) [`data`](#Operation.Warnings.data) = 2
* `string` [`message`](#Operation.Warnings.message) = 3

### `code` {#Operation.Warnings.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#Operation.Warnings.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.Operation.Warnings.Data`](gcp_compute.md#Operation.Warnings.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#Operation.Warnings.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warnings` {#Route.Warnings}

[Output Only] If potential misconfigurations are detected for this route,
this field will be populated with warning messages.


### Inputs for `Warnings`

* `string` [`code`](#Route.Warnings.code) = 1
* `repeated` [`compute.Route.Warnings.Data`](gcp_compute.md#Route.Warnings.Data) [`data`](#Route.Warnings.data) = 2
* `string` [`message`](#Route.Warnings.message) = 3

### `code` {#Route.Warnings.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#Route.Warnings.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.Route.Warnings.Data`](gcp_compute.md#Route.Warnings.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#Route.Warnings.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `Warnings` {#SslPolicy.Warnings}

[Output Only] If potential misconfigurations are detected for this SSL
policy, this field will be populated with warning messages.


### Inputs for `Warnings`

* `string` [`code`](#SslPolicy.Warnings.code) = 1
* `repeated` [`compute.SslPolicy.Warnings.Data`](gcp_compute.md#SslPolicy.Warnings.Data) [`data`](#SslPolicy.Warnings.data) = 2
* `string` [`message`](#SslPolicy.Warnings.message) = 3

### `code` {#SslPolicy.Warnings.code}

| Property | Comments |
|----------|----------|
| Field Name | `code` |
| Type | `string` |

[Output Only] A warning code, if applicable. For example, Compute Engine
returns NO_RESULTS_ON_PAGE if there are no results in the response.
Valid values:
    CLEANUP_FAILED
    DEPRECATED_RESOURCE_USED
    DEPRECATED_TYPE_USED
    DISK_SIZE_LARGER_THAN_IMAGE_SIZE
    EXPERIMENTAL_TYPE_USED
    EXTERNAL_API_WARNING
    FIELD_VALUE_OVERRIDEN
    INJECTED_KERNELS_DEPRECATED
    MISSING_TYPE_DEPENDENCY
    NEXT_HOP_ADDRESS_NOT_ASSIGNED
    NEXT_HOP_CANNOT_IP_FORWARD
    NEXT_HOP_INSTANCE_NOT_FOUND
    NEXT_HOP_INSTANCE_NOT_ON_NETWORK
    NEXT_HOP_NOT_RUNNING
    NOT_CRITICAL_ERROR
    NO_RESULTS_ON_PAGE
    REQUIRED_TOS_AGREEMENT
    RESOURCE_IN_USE_BY_OTHER_RESOURCE_WARNING
    RESOURCE_NOT_DELETED
    SCHEMA_VALIDATION_IGNORED
    SINGLE_INSTANCE_PROPERTY_TEMPLATE
    UNDECLARED_PROPERTIES
    UNREACHABLE

### `data` {#SslPolicy.Warnings.data}

| Property | Comments |
|----------|----------|
| Field Name | `data` |
| Type | [`compute.SslPolicy.Warnings.Data`](gcp_compute.md#SslPolicy.Warnings.Data) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `message` {#SslPolicy.Warnings.message}

| Property | Comments |
|----------|----------|
| Field Name | `message` |
| Type | `string` |

[Output Only] A human-readable description of the warning code.

## Message `XpnHostList` {#XpnHostList}



### Inputs for `XpnHostList`

* `string` [`id`](#XpnHostList.id) = 1
* `repeated` [`compute.Project`](gcp_compute.md#Project) [`items`](#XpnHostList.items) = 2
* `string` [`kind`](#XpnHostList.kind) = 3
* `string` [`nextPageToken`](#XpnHostList.nextPageToken) = 4
* `string` [`selfLink`](#XpnHostList.selfLink) = 5
* [`compute.XpnHostList.Warning`](gcp_compute.md#XpnHostList.Warning) [`warning`](#XpnHostList.warning) = 6

### `id` {#XpnHostList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#XpnHostList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Project`](gcp_compute.md#Project) |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] A list of shared VPC host project URLs.

### `kind` {#XpnHostList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of resource. Always compute#xpnHostList for lists of
shared VPC hosts.

### `nextPageToken` {#XpnHostList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#XpnHostList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#XpnHostList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.XpnHostList.Warning`](gcp_compute.md#XpnHostList.Warning) |

## Message `XpnResourceId` {#XpnResourceId}

Service resource (a.k.a service project) ID.


### Inputs for `XpnResourceId`

* `string` [`id`](#XpnResourceId.id) = 1
* `string` [`type`](#XpnResourceId.type) = 2

### `id` {#XpnResourceId.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

The ID of the service resource. In the case of projects, this field matches
the project ID (e.g., my-project), not the project number (e.g., 12345678).

### `type` {#XpnResourceId.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | `string` |

The type of the service resource.
Valid values:
    PROJECT
    XPN_RESOURCE_TYPE_UNSPECIFIED

## Message `Zone` {#Zone}

A Zone resource. (== resource_for beta.zones ==) (== resource_for v1.zones
==)


### Inputs for `Zone`

* `repeated` `string` [`availableCpuPlatforms`](#Zone.availableCpuPlatforms) = 1
* `string` [`creationTimestamp`](#Zone.creationTimestamp) = 2
* [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) [`deprecated`](#Zone.deprecated) = 3
* `string` [`description`](#Zone.description) = 4
* `string` [`id`](#Zone.id) = 5
* `string` [`kind`](#Zone.kind) = 6
* `string` [`name`](#Zone.name) = 7 (**Required**)
* `string` [`region`](#Zone.region) = 8
* `string` [`selfLink`](#Zone.selfLink) = 9
* `string` [`status`](#Zone.status) = 10

### `availableCpuPlatforms` {#Zone.availableCpuPlatforms}

| Property | Comments |
|----------|----------|
| Field Name | `availableCpuPlatforms` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

[Output Only] Available cpu/platform selections for the zone.

### `creationTimestamp` {#Zone.creationTimestamp}

| Property | Comments |
|----------|----------|
| Field Name | `creationTimestamp` |
| Type | `string` |

[Output Only] Creation timestamp in RFC3339 text format.

### `deprecated` {#Zone.deprecated}

| Property | Comments |
|----------|----------|
| Field Name | `deprecated` |
| Type | [`compute.DeprecationStatus`](gcp_compute.md#DeprecationStatus) |

[Output Only] The deprecation status associated with this zone.

### `description` {#Zone.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

[Output Only] Textual description of the resource.

### `id` {#Zone.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] The unique identifier for the resource. This identifier is
defined by the server.

### `kind` {#Zone.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

[Output Only] Type of the resource. Always compute#zone for zones.

### `name` {#Zone.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

[Output Only] Name of the resource.

### `region` {#Zone.region}

| Property | Comments |
|----------|----------|
| Field Name | `region` |
| Type | `string` |

[Output Only] Full URL reference to the region which hosts the zone.

### `selfLink` {#Zone.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for the resource.

### `status` {#Zone.status}

| Property | Comments |
|----------|----------|
| Field Name | `status` |
| Type | `string` |

[Output Only] Status of the zone, either UP or DOWN.
Valid values:
    DOWN
    UP

## Message `ZoneList` {#ZoneList}

Contains a list of zone resources.


### Inputs for `ZoneList`

* `string` [`id`](#ZoneList.id) = 1
* `repeated` [`compute.Zone`](gcp_compute.md#Zone) [`items`](#ZoneList.items) = 2
* `string` [`kind`](#ZoneList.kind) = 3
* `string` [`nextPageToken`](#ZoneList.nextPageToken) = 4
* `string` [`selfLink`](#ZoneList.selfLink) = 5
* [`compute.ZoneList.Warning`](gcp_compute.md#ZoneList.Warning) [`warning`](#ZoneList.warning) = 6

### `id` {#ZoneList.id}

| Property | Comments |
|----------|----------|
| Field Name | `id` |
| Type | `string` |

[Output Only] Unique identifier for the resource; defined by the server.

### `items` {#ZoneList.items}

| Property | Comments |
|----------|----------|
| Field Name | `items` |
| Type | [`compute.Zone`](gcp_compute.md#Zone) |
| Repeated | Any number of instances of this type is allowed in the schema. |

A list of Zone resources.

### `kind` {#ZoneList.kind}

| Property | Comments |
|----------|----------|
| Field Name | `kind` |
| Type | `string` |

Type of resource.

### `nextPageToken` {#ZoneList.nextPageToken}

| Property | Comments |
|----------|----------|
| Field Name | `nextPageToken` |
| Type | `string` |

[Output Only] This token allows you to get the next page of results for list
requests. If the number of results is larger than maxResults, use the
nextPageToken as a value for the query parameter pageToken in the next list
request. Subsequent list requests will have their own nextPageToken to
continue paging through the results.

### `selfLink` {#ZoneList.selfLink}

| Property | Comments |
|----------|----------|
| Field Name | `selfLink` |
| Type | `string` |

[Output Only] Server-defined URL for this resource.

### `warning` {#ZoneList.warning}

| Property | Comments |
|----------|----------|
| Field Name | `warning` |
| Type | [`compute.ZoneList.Warning`](gcp_compute.md#ZoneList.Warning) |

## Message `ZoneSetLabelsRequest` {#ZoneSetLabelsRequest}



### Inputs for `ZoneSetLabelsRequest`

* `string` [`labelFingerprint`](#ZoneSetLabelsRequest.labelFingerprint) = 1
* `repeated` [`compute.ZoneSetLabelsRequest.LabelsEntry`](gcp_compute.md#ZoneSetLabelsRequest.LabelsEntry) [`labels`](#ZoneSetLabelsRequest.labels) = 2

### `labelFingerprint` {#ZoneSetLabelsRequest.labelFingerprint}

| Property | Comments |
|----------|----------|
| Field Name | `labelFingerprint` |
| Type | `string` |

The fingerprint of the previous set of labels for this resource, used to
detect conflicts. The fingerprint is initially generated by Compute Engine
and changes after every request to modify or update labels. You must always
provide an up-to-date fingerprint hash in order to update or change labels.
Make a get() request to the resource to get the latest fingerprint.

### `labels` {#ZoneSetLabelsRequest.labels}

| Property | Comments |
|----------|----------|
| Field Name | `labels` |
| Type | [`compute.ZoneSetLabelsRequest.LabelsEntry`](gcp_compute.md#ZoneSetLabelsRequest.LabelsEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The labels to set for this resource.



# Enumerations


---
Generated from `schema/gcp/compute/compute-api.proto`.
