# Schema `asset` {#asset}

Asset manifest.


Messages that are valid in package `asset` are as follows:

*** note
Note that this document uses the term "message" to refer to the same concept as
a "message" in Protocol Buffers. Hence every asset and host resource
description is a *message*. So is their embedded structures.
***

## Message `ActiveDirectoryDomain` {#ActiveDirectoryDomain}

Describes an Active Directory domain.

All ActiveDirectoryDomain definitions live under the "asset.ad_domain"
collection. When authoring, leave out the "asset" component. E.g.:

``` textpb
ad_domain {
  name: "my-domain"
  ...
}
```

Deploying a new forest
----------------------
An `ActiveDirectoryDomain` message with no `parent_name` and no `forest`
results in a new forest. The domain becomes the root domain for the forest.

E.g.: The following messages describes a new forest where the root domain is
`win.example.com`.

``` textpb
ad_domain {
  name: "win.example.com"
  domain_controller {
    windows_machine: "my-ad"
  }
}
```

Deploying a new child domain
----------------------------
An `ActiveDirectoryDomain` message with non-empty `parent_name` constructs a
child domain.

E.g.: The following message describes a new child domain named `sales` under
the parent domain `win.example.com`.

``` textpb
ad_domain {
  name: "sales.win.example.com"
  parent_name: "win.example.com"
  domain_controller {
    windows_machine: "my-ad"
  }
}
```

Deploying a new tree in an existing forest
------------------------------------------
An `ActiveDirectoryDomain` message with a non-empty `forest` constructs a
new tree in an existing forest.

E.g.: The following message describes a new tree rooted at
`research.example.org` in the forest identified by `win.example.com`. Note
that the forest is identified by the `ActiveDirectoryDomain` message
corresponding to the root domain in the forest.

``` textpb
ad_domain {
  name: "research.example.org"
  forest: "win.example.com"
  domain_controller {
    windows_machine: "my-ad"
  }
}
```


### Inputs for `ActiveDirectoryDomain`

* `string` [`name`](#ActiveDirectoryDomain.name) = 1
* `string` [`parent_name`](#ActiveDirectoryDomain.parent_name) = 2 (**Required**)
* `string` [`forest`](#ActiveDirectoryDomain.forest) = 3 (**Required**)
* [`asset.ActiveDirectoryDomain.FunctionalLevel`](asset.md#ActiveDirectoryDomain.FunctionalLevel) [`domain_mode`](#ActiveDirectoryDomain.domain_mode) = 4
* [`asset.ActiveDirectoryDomain.FunctionalLevel`](asset.md#ActiveDirectoryDomain.FunctionalLevel) [`forest_mode`](#ActiveDirectoryDomain.forest_mode) = 5
* `string` [`netbios_name`](#ActiveDirectoryDomain.netbios_name) = 6 (**Required**)
* `repeated` [`asset.ActiveDirectoryDomainController`](asset.md#ActiveDirectoryDomainController) [`domain_controller`](#ActiveDirectoryDomain.domain_controller) = 100 (**Required**)

### Outputs for `ActiveDirectoryDomain`

* [`common.Secret`](common.md#Secret) [`safe_mode_admin_password`](#ActiveDirectoryDomain.safe_mode_admin_password) = 200

### `name` {#ActiveDirectoryDomain.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |

FQDN of the domain in lower case.

### `parent_name` {#ActiveDirectoryDomain.parent_name}

| Property | Comments |
|----------|----------|
| Field Name | `parent_name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_domain` |

Parent domain name. Only specify this if this domain is going to be a
child domain.

### `forest` {#ActiveDirectoryDomain.forest}

| Property | Comments |
|----------|----------|
| Field Name | `forest` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_domain` |

Forest name. Only specify this if this domain should be a new tree in an
existing forest. The value of this field should refer to the
ActiveDirectoryDomain entry corresponding to the root domain of the
forest.

If this field is empty, then this domain will become the root domain of a
new forest. Don't set this field to the same value as the `name` field.

### `domain_mode` {#ActiveDirectoryDomain.domain_mode}

| Property | Comments |
|----------|----------|
| Field Name | `domain_mode` |
| Type | [`asset.ActiveDirectoryDomain.FunctionalLevel`](asset.md#ActiveDirectoryDomain.FunctionalLevel) |

### `forest_mode` {#ActiveDirectoryDomain.forest_mode}

| Property | Comments |
|----------|----------|
| Field Name | `forest_mode` |
| Type | [`asset.ActiveDirectoryDomain.FunctionalLevel`](asset.md#ActiveDirectoryDomain.FunctionalLevel) |

Active Directory forest functinoal level. A.k.a. Forest Mode.

This value only applies to ActiveDirectoryDomain objects that describe the
root domain of a new forest. I.e. Whenever the `forest` field is empty.

### `netbios_name` {#ActiveDirectoryDomain.netbios_name}

| Property | Comments |
|----------|----------|
| Field Name | `netbios_name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

NetBIOS name.

### `domain_controller` {#ActiveDirectoryDomain.domain_controller}

| Property | Comments |
|----------|----------|
| Field Name | `domain_controller` |
| Type | [`asset.ActiveDirectoryDomainController`](asset.md#ActiveDirectoryDomainController) |
| Repeated | Any number of instances of this type is allowed in the schema. |
| Required | This field is required. It is an error to omit this field. |

Domain controllers for this domain. At least one is required.

### `safe_mode_admin_password` {#ActiveDirectoryDomain.safe_mode_admin_password}

| Property | Comments |
|----------|----------|
| Field Name | `safe_mode_admin_password` |
| Type | [`common.Secret`](common.md#Secret) |

Safe Mode / Directory Services Restore Mode Administrator password.

## Message `ActiveDirectoryDomainController` {#ActiveDirectoryDomainController}

Describes a single Active Directory Domain Controller. This message should
be embedded inside an ActiveDirectoryDomain message which implicitly binds
the Domain Controller to the containing domain.


### Inputs for `ActiveDirectoryDomainController`

* `string` [`windows_machine`](#ActiveDirectoryDomainController.windows_machine) = 1 (**Required**)
* `bool` [`install_dns`](#ActiveDirectoryDomainController.install_dns) = 2
* `bool` [`no_dns_on_network`](#ActiveDirectoryDomainController.no_dns_on_network) = 3
* `bool` [`no_global_catalog`](#ActiveDirectoryDomainController.no_global_catalog) = 4
* `bool` [`create_dns_delegation`](#ActiveDirectoryDomainController.create_dns_delegation) = 5

### `windows_machine` {#ActiveDirectoryDomainController.windows_machine}

| Property | Comments |
|----------|----------|
| Field Name | `windows_machine` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_machine` |

Machine hosting the ADDS. Must match the `name` of a WindowsMachine entry.

### `install_dns` {#ActiveDirectoryDomainController.install_dns}

| Property | Comments |
|----------|----------|
| Field Name | `install_dns` |
| Type | `bool` |

### `no_dns_on_network` {#ActiveDirectoryDomainController.no_dns_on_network}

| Property | Comments |
|----------|----------|
| Field Name | `no_dns_on_network` |
| Type | `bool` |

Assume DNS service is not available on the network. Only applicable when
installing DNS services. If this field is not set, or set to false, then
the installation can assume that the TCP/IP client settings of the host OS
specifies the DNS server to use.

### `no_global_catalog` {#ActiveDirectoryDomainController.no_global_catalog}

| Property | Comments |
|----------|----------|
| Field Name | `no_global_catalog` |
| Type | `bool` |

This domain controller should not be a global catalog server. Default is
to run with global catalog for Win2012 or later.

### `create_dns_delegation` {#ActiveDirectoryDomainController.create_dns_delegation}

| Property | Comments |
|----------|----------|
| Field Name | `create_dns_delegation` |
| Type | `bool` |

If true, attempts to create a DNS delegation for the new DNS server. Only
applicable when installing a DNS server. E.g.: If the authoritative DNS
server for foo.example.com is using ActiveDirectory, and we are installing
the subordinate domain bar, then setting this value to true causes
foo.example.com to delegate the bar domain to the new DNS server.

## Message `ActiveDirectoryGroupPolicy` {#ActiveDirectoryGroupPolicy}

Describes an Active Directory GPO. The GPO itself may not contain anything
particularly important on creation, and is entirely based on the starter
GPO. The GPO is also not linked anywhere by default. Use the
ActiveDirectoryGroupPolicyLink message to create links, and
ActiveDirectoryRegistryPolicy to add registry based policies.

E.g.:

    # Create a group policy.
    ad_group_policy {
      name: 'foo'
      ad_domain: 'my-domain'
    }

    # Add some registry values to it.
    ad_registry_policy {
      name: 'reg-pol-0001'
      ad_group_policy: 'foo'

      key {
        path: 'HKCU\\Software\\My Company\\Foo\\Bar'

        value {
          name: 'version'
          value: '1.0'
        }

        value {
          name: 'FooCount'
          type: DWORD
          value: '10'
        }
      }
    }

    # And link it to one or more OUs.
    ad_group_policy_link {
      name: 'foo-link'
      ad_group_policy: 'foo'
      container { ad_organizational_unit: 'my-ou' }
      container { ad_organizational_unit: 'your-ou' }
      enforced: true
    }


### Inputs for `ActiveDirectoryGroupPolicy`

* `string` [`name`](#ActiveDirectoryGroupPolicy.name) = 1 (**Required**)
* `string` [`full_name`](#ActiveDirectoryGroupPolicy.full_name) = 2
* `string` [`ad_domain`](#ActiveDirectoryGroupPolicy.ad_domain) = 3 (**Required**)
* `string` [`comment`](#ActiveDirectoryGroupPolicy.comment) = 4
* `string` [`based_on`](#ActiveDirectoryGroupPolicy.based_on) = 5 (**Required**)
* `repeated` [`asset.ActiveDirectoryRegistryPolicy`](asset.md#ActiveDirectoryRegistryPolicy) [`registry`](#ActiveDirectoryGroupPolicy.registry) = 6
* `repeated` [`asset.ActiveDirectoryRegistryPrefPolicy`](asset.md#ActiveDirectoryRegistryPrefPolicy) [`registry_pref`](#ActiveDirectoryGroupPolicy.registry_pref) = 7

### `name` {#ActiveDirectoryGroupPolicy.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

The name of the GPO.

### `full_name` {#ActiveDirectoryGroupPolicy.full_name}

| Property | Comments |
|----------|----------|
| Field Name | `full_name` |
| Type | `string` |

The full name. Only use this if the full name cannot be specified in
|name|.

### `ad_domain` {#ActiveDirectoryGroupPolicy.ad_domain}

| Property | Comments |
|----------|----------|
| Field Name | `ad_domain` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_domain` |

The domain in which this GPO is created.

### `comment` {#ActiveDirectoryGroupPolicy.comment}

| Property | Comments |
|----------|----------|
| Field Name | `comment` |
| Type | `string` |

A comment for the GPO. This can contain up to 2047 characters.

### `based_on` {#ActiveDirectoryGroupPolicy.based_on}

| Property | Comments |
|----------|----------|
| Field Name | `based_on` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_group_policy` |

If a template GPO is specified, that will be used as the starter GPO.
Creates a starter GPO is no value is specified here.

### `registry` {#ActiveDirectoryGroupPolicy.registry}

| Property | Comments |
|----------|----------|
| Field Name | `registry` |
| Type | [`asset.ActiveDirectoryRegistryPolicy`](asset.md#ActiveDirectoryRegistryPolicy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `registry_pref` {#ActiveDirectoryGroupPolicy.registry_pref}

| Property | Comments |
|----------|----------|
| Field Name | `registry_pref` |
| Type | [`asset.ActiveDirectoryRegistryPrefPolicy`](asset.md#ActiveDirectoryRegistryPrefPolicy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `ActiveDirectoryGroupPolicyLink` {#ActiveDirectoryGroupPolicyLink}

Describes one or more GPO links.


### Inputs for `ActiveDirectoryGroupPolicyLink`

* `string` [`name`](#ActiveDirectoryGroupPolicyLink.name) = 1 (**Required**)
* `string` [`ad_group_policy`](#ActiveDirectoryGroupPolicyLink.ad_group_policy) = 2 (**Required**)
* `repeated` [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#ActiveDirectoryGroupPolicyLink.container) = 3
* `bool` [`enforced`](#ActiveDirectoryGroupPolicyLink.enforced) = 4
* `bool` [`enabled`](#ActiveDirectoryGroupPolicyLink.enabled) = 5

### `name` {#ActiveDirectoryGroupPolicyLink.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

A convenient identifier for this set of GPO links.

### `ad_group_policy` {#ActiveDirectoryGroupPolicyLink.ad_group_policy}

| Property | Comments |
|----------|----------|
| Field Name | `ad_group_policy` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_group_policy` |

The policy that will be linked.

### `container` {#ActiveDirectoryGroupPolicyLink.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of containers for which the GPO referred to in |ad_group_policy|
will be linked. Only the |ad_domain| and |ad_organizational_unit| fields
can be used. There should be at least one of these.

### `enforced` {#ActiveDirectoryGroupPolicyLink.enforced}

| Property | Comments |
|----------|----------|
| Field Name | `enforced` |
| Type | `bool` |

### `enabled` {#ActiveDirectoryGroupPolicyLink.enabled}

| Property | Comments |
|----------|----------|
| Field Name | `enabled` |
| Type | `bool` |

## Message `ActiveDirectoryOrganizationalUnit` {#ActiveDirectoryOrganizationalUnit}

Describes an Active Directory Organizational-Unit.

By default the `name` field corresponds to the `ou` attribute of the
`Organizational-Unit` object. If the `name` field cannot contain the value
(e.g. because the `ou` value needs to contain characters that are not
allowed in the `name` field), then the `full_name` field can be used
instead.

Note that this message omits location properties like postal code, city,
state, and country. Use the 'property' field to specify these if required.

E.g.:
    ad_organizational_unit {
       name: 'foo'
       ...
       property { key: 'l', value: 'Cambridge' }
       property { key: 'st', value: 'MA' }
       ..
    }


### Inputs for `ActiveDirectoryOrganizationalUnit`

* `string` [`name`](#ActiveDirectoryOrganizationalUnit.name) = 1 (**Required**)
* `string` [`full_name`](#ActiveDirectoryOrganizationalUnit.full_name) = 2
* [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#ActiveDirectoryOrganizationalUnit.container) = 3 (**Required**)
* `string` [`server`](#ActiveDirectoryOrganizationalUnit.server) = 4 (**Required**)
* `string` [`based_on`](#ActiveDirectoryOrganizationalUnit.based_on) = 5 (**Required**)
* `string` [`display_name`](#ActiveDirectoryOrganizationalUnit.display_name) = 6
* `string` [`description`](#ActiveDirectoryOrganizationalUnit.description) = 7
* [`asset.UserOrGroupReference`](asset.md#UserOrGroupReference) [`managed_by`](#ActiveDirectoryOrganizationalUnit.managed_by) = 8
* `repeated` [`asset.ActiveDirectoryOrganizationalUnit.AttributeEntry`](asset.md#ActiveDirectoryOrganizationalUnit.AttributeEntry) [`attribute`](#ActiveDirectoryOrganizationalUnit.attribute) = 9

### `name` {#ActiveDirectoryOrganizationalUnit.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the OU. Also populates the 'ou' attribute
(`Organizational-Unit-Name`) of the AD object unless overridden by
|full_name|.

### `full_name` {#ActiveDirectoryOrganizationalUnit.full_name}

| Property | Comments |
|----------|----------|
| Field Name | `full_name` |
| Type | `string` |

The 'name' property of the AD object. Only specify this if |name| can't be
used to store the display name.

### `container` {#ActiveDirectoryOrganizationalUnit.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |
| Required | This field is required. It is an error to omit this field. |

The container in which this OU is created. Only the 'ad_domain' and
'ad_organizational_unit' values are valid.

### `server` {#ActiveDirectoryOrganizationalUnit.server}

| Property | Comments |
|----------|----------|
| Field Name | `server` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_machine` |

The AD DS server which should service the request for creating the OU if
necessary. If left unspecified, one of the candidate AD DS instances
associated with 'container' will be used.

### `based_on` {#ActiveDirectoryOrganizationalUnit.based_on}

| Property | Comments |
|----------|----------|
| Field Name | `based_on` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_organizational_unit` |

Use the referred AD OU as the template. Any properties specified in this
message will override corresponding properties from the template OU. Any
properties specified in the template, but not in this message will be
copied over.

### `display_name` {#ActiveDirectoryOrganizationalUnit.display_name}

| Property | Comments |
|----------|----------|
| Field Name | `display_name` |
| Type | `string` |

The 'displayName' attribute (`Display-Name`) of the AD object.

### `description` {#ActiveDirectoryOrganizationalUnit.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

The 'description' attribute (`Description`) of the AD object.

### `managed_by` {#ActiveDirectoryOrganizationalUnit.managed_by}

| Property | Comments |
|----------|----------|
| Field Name | `managed_by` |
| Type | [`asset.UserOrGroupReference`](asset.md#UserOrGroupReference) |

The principal managing this OU. Note that the 'container' property for a
user or a group is implicit and should be omitted.

### `attribute` {#ActiveDirectoryOrganizationalUnit.attribute}

| Property | Comments |
|----------|----------|
| Field Name | `attribute` |
| Type | [`asset.ActiveDirectoryOrganizationalUnit.AttributeEntry`](asset.md#ActiveDirectoryOrganizationalUnit.AttributeEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Additional attributes. The key is the ldapDisplayName of the attribute.
The value can be a single string. Repeat using the same key to specify
more than one value for a single attribute.

E.g.:
    attribute {
      key: 'telephoneNumber'
      value: '+1-617-555-1234'
    }

    attribute {
      key: 'postalCode'
      value: '02141'
    }

## Message `ActiveDirectoryRegistryPolicy` {#ActiveDirectoryRegistryPolicy}

Describes a set of registry keys that should be applied for a GPO.


### Inputs for `ActiveDirectoryRegistryPolicy`

* `bool` [`additive`](#ActiveDirectoryRegistryPolicy.additive) = 3
* `repeated` [`asset.RegistryKey`](asset.md#RegistryKey) [`key`](#ActiveDirectoryRegistryPolicy.key) = 4

### `additive` {#ActiveDirectoryRegistryPolicy.additive}

| Property | Comments |
|----------|----------|
| Field Name | `additive` |
| Type | `bool` |

If true, the registry values defined in this message will be added to the
respective registry keys. The default behavior -- also the behavior when
this field is set to false -- is to delete all the values under each key
prior to adding the new values.

### `key` {#ActiveDirectoryRegistryPolicy.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | [`asset.RegistryKey`](asset.md#RegistryKey) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Set of registry keys and values that will be applied (in order).

## Message `ActiveDirectoryRegistryPrefPolicy` {#ActiveDirectoryRegistryPrefPolicy}

Describes a registry change that should happen when applying a GPO.


### Inputs for `ActiveDirectoryRegistryPrefPolicy`

* [`asset.ActiveDirectoryRegistryPrefPolicy.Action`](asset.md#ActiveDirectoryRegistryPrefPolicy.Action) [`action`](#ActiveDirectoryRegistryPrefPolicy.action) = 3
* `repeated` [`asset.RegistryKey`](asset.md#RegistryKey) [`key`](#ActiveDirectoryRegistryPrefPolicy.key) = 4

### `action` {#ActiveDirectoryRegistryPrefPolicy.action}

| Property | Comments |
|----------|----------|
| Field Name | `action` |
| Type | [`asset.ActiveDirectoryRegistryPrefPolicy.Action`](asset.md#ActiveDirectoryRegistryPrefPolicy.Action) |

### `key` {#ActiveDirectoryRegistryPrefPolicy.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | [`asset.RegistryKey`](asset.md#RegistryKey) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Set of registry keys and values that will be applied (in order).

## Message `Address` {#Address}

Address is an IPv4 or IPv6 address.


### Inputs for `Address`

* `string` [`ip`](#Address.ip) = 1 (**Required**)

### `ip` {#Address.ip}

| Property | Comments |
|----------|----------|
| Field Name | `ip` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

## Message `AddressRange` {#AddressRange}

AddressRange is an IPv4 or IPv6 CIDR range.


### Inputs for `AddressRange`

* `string` [`cidr`](#AddressRange.cidr) = 1 (**Required**)

### `cidr` {#AddressRange.cidr}

| Property | Comments |
|----------|----------|
| Field Name | `cidr` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

## Message `AssetManifest` {#AssetManifest}

AssetManifest is the main container for all top level assets that go into an
asset manifest and can be identified by name. Please refer to each
individual message for naming and other requirements.

The intended use for this message is to act as the primary message into
which text format protobufs can be deserialized when reading configuration
files. I.e., a file could be formatted as follows:

``` textpb
  network  { name: "foo" }

  dns_zone  {
    origin: "foo.example."
    record  {
      name: "bar"
      ttl: 3600
      record_type: "A"
      answer: "10.10.2.20"
    }
  }
```

This file can be parsed into an AssetManifest message. Furthermore, multiple
such messages can be concatenated without losing information.  The latter
property is important since it allows a single asset manifest to be
distributed across a number of files.

The manner in which the names appear in a text format protobuf should
underscore the style descision to use singular forms when naming repeated
fields.


### Inputs for `AssetManifest`

* `repeated` [`asset.Network`](asset.md#Network) [`network`](#AssetManifest.network) = 1
* `repeated` [`asset.DNSZone`](asset.md#DNSZone) [`dns_zone`](#AssetManifest.dns_zone) = 2
* `repeated` [`asset.ActiveDirectoryDomain`](asset.md#ActiveDirectoryDomain) [`ad_domain`](#AssetManifest.ad_domain) = 100
* `repeated` [`asset.ActiveDirectoryOrganizationalUnit`](asset.md#ActiveDirectoryOrganizationalUnit) [`ad_organizational_unit`](#AssetManifest.ad_organizational_unit) = 101
* `repeated` [`asset.ActiveDirectoryGroupPolicy`](asset.md#ActiveDirectoryGroupPolicy) [`ad_group_policy`](#AssetManifest.ad_group_policy) = 102
* `repeated` [`asset.ActiveDirectoryGroupPolicyLink`](asset.md#ActiveDirectoryGroupPolicyLink) [`ad_group_policy_link`](#AssetManifest.ad_group_policy_link) = 103
* `repeated` [`asset.WindowsGroup`](asset.md#WindowsGroup) [`windows_group`](#AssetManifest.windows_group) = 104
* `repeated` [`asset.WindowsMachine`](asset.md#WindowsMachine) [`windows_machine`](#AssetManifest.windows_machine) = 105
* `repeated` [`asset.WindowsUser`](asset.md#WindowsUser) [`windows_user`](#AssetManifest.windows_user) = 106
* `repeated` [`asset.Certificate`](asset.md#Certificate) [`certificate`](#AssetManifest.certificate) = 200
* `repeated` [`asset.CertificatePool`](asset.md#CertificatePool) [`certificate_pool`](#AssetManifest.certificate_pool) = 201
* `repeated` [`asset.IISApplication`](asset.md#IISApplication) [`iis_application`](#AssetManifest.iis_application) = 300
* `repeated` [`asset.IISServer`](asset.md#IISServer) [`iis_server`](#AssetManifest.iis_server) = 301
* `repeated` [`asset.IISSite`](asset.md#IISSite) [`iis_site`](#AssetManifest.iis_site) = 302
* `repeated` [`asset.RemoteDesktopHost`](asset.md#RemoteDesktopHost) [`remote_desktop_host`](#AssetManifest.remote_desktop_host) = 400

### `network` {#AssetManifest.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | [`asset.Network`](asset.md#Network) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Core networking. Use field numbers 1-99

### `dns_zone` {#AssetManifest.dns_zone}

| Property | Comments |
|----------|----------|
| Field Name | `dns_zone` |
| Type | [`asset.DNSZone`](asset.md#DNSZone) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `ad_domain` {#AssetManifest.ad_domain}

| Property | Comments |
|----------|----------|
| Field Name | `ad_domain` |
| Type | [`asset.ActiveDirectoryDomain`](asset.md#ActiveDirectoryDomain) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Windows and ActiveDirectory. Use field numbers 100-199

### `ad_organizational_unit` {#AssetManifest.ad_organizational_unit}

| Property | Comments |
|----------|----------|
| Field Name | `ad_organizational_unit` |
| Type | [`asset.ActiveDirectoryOrganizationalUnit`](asset.md#ActiveDirectoryOrganizationalUnit) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `ad_group_policy` {#AssetManifest.ad_group_policy}

| Property | Comments |
|----------|----------|
| Field Name | `ad_group_policy` |
| Type | [`asset.ActiveDirectoryGroupPolicy`](asset.md#ActiveDirectoryGroupPolicy) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `ad_group_policy_link` {#AssetManifest.ad_group_policy_link}

| Property | Comments |
|----------|----------|
| Field Name | `ad_group_policy_link` |
| Type | [`asset.ActiveDirectoryGroupPolicyLink`](asset.md#ActiveDirectoryGroupPolicyLink) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `windows_group` {#AssetManifest.windows_group}

| Property | Comments |
|----------|----------|
| Field Name | `windows_group` |
| Type | [`asset.WindowsGroup`](asset.md#WindowsGroup) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `windows_machine` {#AssetManifest.windows_machine}

| Property | Comments |
|----------|----------|
| Field Name | `windows_machine` |
| Type | [`asset.WindowsMachine`](asset.md#WindowsMachine) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `windows_user` {#AssetManifest.windows_user}

| Property | Comments |
|----------|----------|
| Field Name | `windows_user` |
| Type | [`asset.WindowsUser`](asset.md#WindowsUser) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `certificate` {#AssetManifest.certificate}

| Property | Comments |
|----------|----------|
| Field Name | `certificate` |
| Type | [`asset.Certificate`](asset.md#Certificate) |
| Repeated | Any number of instances of this type is allowed in the schema. |

PKI. Use field numbers 200-299

### `certificate_pool` {#AssetManifest.certificate_pool}

| Property | Comments |
|----------|----------|
| Field Name | `certificate_pool` |
| Type | [`asset.CertificatePool`](asset.md#CertificatePool) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `iis_application` {#AssetManifest.iis_application}

| Property | Comments |
|----------|----------|
| Field Name | `iis_application` |
| Type | [`asset.IISApplication`](asset.md#IISApplication) |
| Repeated | Any number of instances of this type is allowed in the schema. |

IIS. Use field numbers 300-399

### `iis_server` {#AssetManifest.iis_server}

| Property | Comments |
|----------|----------|
| Field Name | `iis_server` |
| Type | [`asset.IISServer`](asset.md#IISServer) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `iis_site` {#AssetManifest.iis_site}

| Property | Comments |
|----------|----------|
| Field Name | `iis_site` |
| Type | [`asset.IISSite`](asset.md#IISSite) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `remote_desktop_host` {#AssetManifest.remote_desktop_host}

| Property | Comments |
|----------|----------|
| Field Name | `remote_desktop_host` |
| Type | [`asset.RemoteDesktopHost`](asset.md#RemoteDesktopHost) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Remote Desktop

## Message `AttributeEntry` {#ActiveDirectoryOrganizationalUnit.AttributeEntry}



### Inputs for `AttributeEntry`

* `string` [`key`](#ActiveDirectoryOrganizationalUnit.AttributeEntry.key) = 1
* `string` [`value`](#ActiveDirectoryOrganizationalUnit.AttributeEntry.value) = 2

### `key` {#ActiveDirectoryOrganizationalUnit.AttributeEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#ActiveDirectoryOrganizationalUnit.AttributeEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `Certificate` {#Certificate}

An x509 certificate.


### Inputs for `Certificate`

* `string` [`name`](#Certificate.name) = 1 (**Required**)
* [`common.FileReference`](common.md#FileReference) [`certficate`](#Certificate.certficate) = 2 (**Required**)
* [`common.FileReference`](common.md#FileReference) [`private_key`](#Certificate.private_key) = 3

### `name` {#Certificate.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

A convenient identifier for this certificate.

### `certficate` {#Certificate.certficate}

| Property | Comments |
|----------|----------|
| Field Name | `certficate` |
| Type | [`common.FileReference`](common.md#FileReference) |
| Required | This field is required. It is an error to omit this field. |

PEM encoded file containing a *single* certificate, and no private key.

### `private_key` {#Certificate.private_key}

| Property | Comments |
|----------|----------|
| Field Name | `private_key` |
| Type | [`common.FileReference`](common.md#FileReference) |

PEM encoded private key. Optional.

## Message `CertificatePool` {#CertificatePool}

A pool of certifiates. Usually used to establish a set of trust roots.


### Inputs for `CertificatePool`

* `string` [`name`](#CertificatePool.name) = 1 (**Required**)
* `repeated` `string` [`include_named`](#CertificatePool.include_named) = 2 (**Required**)
* `repeated` [`common.FileReference`](common.md#FileReference) [`include_file`](#CertificatePool.include_file) = 3

### `name` {#CertificatePool.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

A convenient identifier for this certificate pool

### `include_named` {#CertificatePool.include_named}

| Property | Comments |
|----------|----------|
| Field Name | `include_named` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.certificate` |

Named certificates. Should match one of the |Certificate| entries
specified in this manifest.

### `include_file` {#CertificatePool.include_file}

| Property | Comments |
|----------|----------|
| Field Name | `include_file` |
| Type | [`common.FileReference`](common.md#FileReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

PEM encoded file containing one or more certificates. Note that private
keys cannot be specified this way.

## Message `DNSRecord` {#DNSRecord}

DNSRecord describes a single DNS record in a Zone.


### Inputs for `DNSRecord`

* `string` [`name`](#DNSRecord.name) = 1 (**Required**)
* `int32` [`ttl`](#DNSRecord.ttl) = 2 (**Required**)
* `string` [`record_class`](#DNSRecord.record_class) = 3
* `string` [`record_type`](#DNSRecord.record_type) = 4 (**Required**)
* `int32` [`priority`](#DNSRecord.priority) = 5
* `string` [`answer`](#DNSRecord.answer) = 6 (**Required**)

### `name` {#DNSRecord.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name on record.

### `ttl` {#DNSRecord.ttl}

| Property | Comments |
|----------|----------|
| Field Name | `ttl` |
| Type | `int32` |
| Required | This field is required. It is an error to omit this field. |

TTL in seconds.

### `record_class` {#DNSRecord.record_class}

| Property | Comments |
|----------|----------|
| Field Name | `record_class` |
| Type | `string` |

Must be omitted or is always IN.

### `record_type` {#DNSRecord.record_type}

| Property | Comments |
|----------|----------|
| Field Name | `record_type` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Type of record. E.g. A, AAAA, NS, MX, ...

### `priority` {#DNSRecord.priority}

| Property | Comments |
|----------|----------|
| Field Name | `priority` |
| Type | `int32` |

Priority value.

### `answer` {#DNSRecord.answer}

| Property | Comments |
|----------|----------|
| Field Name | `answer` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Answer section.

## Message `DNSZone` {#DNSZone}

DNSZone describes a DNS zone. Conceptually it encompasses the same
information included in a DNS zone file as described in
https://en.wikipedia.org/wiki/Zone_file

TODO(asanka): Document the behavior when name == ActiveDirectoryDomain.name
for some ActiveDirectoryDomain entry. If the AD Domain also deploys a DNS
server, then the |record| entries should be added to the Active Directory
DNS server. Otherwise we need to deploy a different DNS server.


### Inputs for `DNSZone`

* `string` [`name`](#DNSZone.name) = 1
* `repeated` [`asset.DNSRecord`](asset.md#DNSRecord) [`record`](#DNSZone.record) = 2

### `name` {#DNSZone.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |

Origin described by zone. E.g. foo.example.com.

### `record` {#DNSZone.record}

| Property | Comments |
|----------|----------|
| Field Name | `record` |
| Type | [`asset.DNSRecord`](asset.md#DNSRecord) |
| Repeated | Any number of instances of this type is allowed in the schema. |

DNS records that should be included in the zone.

## Message `FixedAddress` {#FixedAddress}

FixedAddress described an address that is determined either by the host
environment or by the asset manifest.


### Inputs for `FixedAddress`

* [`asset.Address`](asset.md#Address) [`address`](#FixedAddress.address) = 1
* `string` [`address_pool`](#FixedAddress.address_pool) = 2 (**Required**)

### Outputs for `FixedAddress`

* [`asset.Address`](asset.md#Address) [`resolved_address`](#FixedAddress.resolved_address) = 3

### `address` {#FixedAddress.address}

| Property | Comments |
|----------|----------|
| Field Name | `address` |
| Type | [`asset.Address`](asset.md#Address) |

Single literal address.

### `address_pool` {#FixedAddress.address_pool}

| Property | Comments |
|----------|----------|
| Field Name | `address_pool` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `host.address_pool` |

One of the available addresses from the named |host.AddressPool| will be
used.

### `resolved_address` {#FixedAddress.resolved_address}

| Property | Comments |
|----------|----------|
| Field Name | `resolved_address` |
| Type | [`asset.Address`](asset.md#Address) |

Upon resolution, this field will contain the actual IP address that's
assigned to this network interface.

## Message `GroupReference` {#GroupReference}

A reference to a group. The combination of |name| and |container| must match
one of the WindowsGroup entries.


### Inputs for `GroupReference`

* `string` [`windows_group`](#GroupReference.windows_group) = 1 (**Required**)
* [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#GroupReference.container) = 2

### `windows_group` {#GroupReference.windows_group}

| Property | Comments |
|----------|----------|
| Field Name | `windows_group` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_group` |

The name of the group.

### `container` {#GroupReference.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |

Location. Since GroupReference messages are typically specified as a field
of an object that already has a container, omiting this field results in
the GroupReference inheriting the parent object's container. Take for
example, the following WindowsUser definition:

``` textpb
    windows_user {
      name: 'joe'
      container: { domain: 'foo.example' }
      member_of: { windows_group: 'bar' }
    }
```

This results in the user being a member of the group 'bar' in the
'foo.example' AD domain because that's the enclosing container.  Note
however, that inheriting in this manner isn't always correct since it is
possible for users to be members of groups from other containers.

## Message `IISApplication` {#IISApplication}



### Inputs for `IISApplication`

* `string` [`name`](#IISApplication.name) = 1 (**Required**)
* `string` [`iis_site`](#IISApplication.iis_site) = 2 (**Required**)
* [`common.FileReference`](common.md#FileReference) [`contents`](#IISApplication.contents) = 3
* [`common.FileReference`](common.md#FileReference) [`web_config_file`](#IISApplication.web_config_file) = 4
* `string` [`web_config_string`](#IISApplication.web_config_string) = 5

### `name` {#IISApplication.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of application.

### `iis_site` {#IISApplication.iis_site}

| Property | Comments |
|----------|----------|
| Field Name | `iis_site` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.iis_site` |

Name of the iis.Site that's hosting this application.

### `contents` {#IISApplication.contents}

| Property | Comments |
|----------|----------|
| Field Name | `contents` |
| Type | [`common.FileReference`](common.md#FileReference) |

Relative path to directory containing the files that will be hosted on
this application. The entire subtree will be copied over to the target
host.

While it is legal, avoid nesting sites and applications within a single
directory tree.

### `web_config_file` {#IISApplication.web_config_file}

| Property | Comments |
|----------|----------|
| Field Name | `web_config_file` |
| Type | [`common.FileReference`](common.md#FileReference) |

Relative path to an optional web.config file that will override any
existing web.config in |contents|.

### `web_config_string` {#IISApplication.web_config_string}

| Property | Comments |
|----------|----------|
| Field Name | `web_config_string` |
| Type | `string` |

Inline web.config contents. Any contents here will be merged with the
|web_config_file| if there is one, a web config file found in the top
level of |contents|.

## Message `IISBindings` {#IISBindings}

Bindings for a website onto a webserver.


### Inputs for `IISBindings`

* `string` [`hostname`](#IISBindings.hostname) = 1
* [`asset.Protocol`](asset.md#Protocol) [`protocol`](#IISBindings.protocol) = 2
* `uint32` [`port`](#IISBindings.port) = 3
* `string` [`certificate`](#IISBindings.certificate) = 4 (**Required**)
* `bool` [`use_sni`](#IISBindings.use_sni) = 5

### `hostname` {#IISBindings.hostname}

| Property | Comments |
|----------|----------|
| Field Name | `hostname` |
| Type | `string` |

Hostname. Can be empty to bind to all hostnames.

### `protocol` {#IISBindings.protocol}

| Property | Comments |
|----------|----------|
| Field Name | `protocol` |
| Type | [`asset.Protocol`](asset.md#Protocol) |

Protocol. Must be either HTTP or HTTPS. If using HTTPS, the SSL
information is also required.

### `port` {#IISBindings.port}

| Property | Comments |
|----------|----------|
| Field Name | `port` |
| Type | `uint32` |

Port to use. Leave empty to use the default port based on protocol.

### `certificate` {#IISBindings.certificate}

| Property | Comments |
|----------|----------|
| Field Name | `certificate` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.certificate` |

Name of a Certificate. This certificate should include a private key. See
the definition in cert.proto.

### `use_sni` {#IISBindings.use_sni}

| Property | Comments |
|----------|----------|
| Field Name | `use_sni` |
| Type | `bool` |

Whether or not to use SNI. Only applicable when |protocol| is HTTPS. If
this value is false, then only one HTTPS site can be bound to a single
server.

## Message `IISServer` {#IISServer}

An IIS server.

The following Windows features will be automatically installed for all IIS
servers.
  *  Web-Common-HTTP
  *  Web-Security


### Inputs for `IISServer`

* `string` [`name`](#IISServer.name) = 1 (**Required**)
* `string` [`windows_machine`](#IISServer.windows_machine) = 2 (**Required**)

### `name` {#IISServer.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of server instance.

### `windows_machine` {#IISServer.windows_machine}

| Property | Comments |
|----------|----------|
| Field Name | `windows_machine` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_machine` |

Name of host machine. Should refer to a |WindowsMachine| entry.
Container information for the IIS instance will be inherited from the
WindowsMachine entry.

## Message `IISSite` {#IISSite}

A single web site.


### Inputs for `IISSite`

* `string` [`name`](#IISSite.name) = 1 (**Required**)
* `string` [`iis_server`](#IISSite.iis_server) = 2 (**Required**)
* [`asset.IISBindings`](asset.md#IISBindings) [`bindings`](#IISSite.bindings) = 3
* [`common.FileReference`](common.md#FileReference) [`contents`](#IISSite.contents) = 4

### `name` {#IISSite.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of site.

### `iis_server` {#IISSite.iis_server}

| Property | Comments |
|----------|----------|
| Field Name | `iis_server` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.iis_server` |

Name of server. Should refer to a |Server| entry. A single server can host
multiple sites as long as those sites use distinct binding information.

### `bindings` {#IISSite.bindings}

| Property | Comments |
|----------|----------|
| Field Name | `bindings` |
| Type | [`asset.IISBindings`](asset.md#IISBindings) |

Bindings. These are bindings. Specify bindings here. If left unset, will
use default bindings (HTTP, port 80, all hostnames).

### `contents` {#IISSite.contents}

| Property | Comments |
|----------|----------|
| Field Name | `contents` |
| Type | [`common.FileReference`](common.md#FileReference) |

Relative path to directory containing the files that will be hosted on
this site. The entire subtree will be copied over to the target host.

While it is legal, avoid nesting sites and applications within a single
directory tree.

## Message `Machine` {#Machine}

Machine describes a generic machine. All substrates of machines should use
the same fieldnames and field numbers for common fields. See WindowsMachine
for an example.


### Inputs for `Machine`

* `string` [`name`](#Machine.name) = 1 (**Required**)
* `string` [`machine_type`](#Machine.machine_type) = 2 (**Required**)
* `repeated` [`asset.NetworkInterface`](asset.md#NetworkInterface) [`network_interface`](#Machine.network_interface) = 3 (**Required**)

### `name` {#Machine.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the machine. Also becomes the hostname of the machine.

### `machine_type` {#Machine.machine_type}

| Property | Comments |
|----------|----------|
| Field Name | `machine_type` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `host.machine_type` |

The name of a host.MachineType entry that describes the host machine.

### `network_interface` {#Machine.network_interface}

| Property | Comments |
|----------|----------|
| Field Name | `network_interface` |
| Type | [`asset.NetworkInterface`](asset.md#NetworkInterface) |
| Repeated | Any number of instances of this type is allowed in the schema. |
| Required | This field is required. It is an error to omit this field. |

Network interfaces. There can be more than one for multihomed machines.
There MUST be at least one of these.

## Message `MultiString` {#RegistryValue.MultiString}

MultiString represents a single REG_MULTI_SZ registry value. Each |value|
represents a single \0 delimited string. All values are concatenated in
order to construct the REG_MULTI_SZ value. No additional \0 values are
required and the final registry value will be correctly \0\0 terminated.


### Inputs for `MultiString`

* `repeated` `string` [`value`](#RegistryValue.MultiString.value) = 1

### `value` {#RegistryValue.MultiString.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

A single string value in UTF-8. It will be re-encoded to UTF-16 before
being written to the registry. Do not include any \0 characters here.
|value| also cannot be the empty string.

## Message `Network` {#Network}



### Inputs for `Network`

* `string` [`name`](#Network.name) = 1 (**Required**)
* [`asset.AddressRange`](asset.md#AddressRange) [`address_range`](#Network.address_range) = 2

### `name` {#Network.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the network.

### `address_range` {#Network.address_range}

| Property | Comments |
|----------|----------|
| Field Name | `address_range` |
| Type | [`asset.AddressRange`](asset.md#AddressRange) |

The address range assigned to the network. If left unspecified, an address
range will be determined when deploying the assets. This is the preferred
option unless an explicit address range is required.

Two networks in the same asset manifest can't have overlapping address
ranges even if they aren't peers.

## Message `NetworkInterface` {#NetworkInterface}

NetworkInterface describes a single network interface on a machine.


### Inputs for `NetworkInterface`

* `string` [`network`](#NetworkInterface.network) = 1 (**Required**)
* [`asset.FixedAddress`](asset.md#FixedAddress) [`fixed_address`](#NetworkInterface.fixed_address) = 2

### `network` {#NetworkInterface.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.network` |

Name of Network entry describing the network that this interface is
attached to.

### `fixed_address` {#NetworkInterface.fixed_address}

| Property | Comments |
|----------|----------|
| Field Name | `fixed_address` |
| Type | [`asset.FixedAddress`](asset.md#FixedAddress) |

Fixed address, if this interface is to have one. Leave undefined if the
interface should obatain an address automatically.

## Message `NetworkPeer` {#NetworkPeer}

NetworkPeer describes a peering group. All networks that are a member of a
peering group can route traffic across each other.


### Inputs for `NetworkPeer`

* `repeated` `string` [`network`](#NetworkPeer.network) = 1 (**Required**)

### `network` {#NetworkPeer.network}

| Property | Comments |
|----------|----------|
| Field Name | `network` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.network` |

List of networks that should form a full mesh. Individual networks are
isolated by default and are only able to talk to each other if:
  * They are a part of a peering group, or
  * They are connected via a VPN gateway, or
  * They are connected via a virtual router.

A single nework can participate in multiple disjoint peering groups,
however peering is not transitive. I.e. If {A,B} is a peering group, and
{B,C} is a peering group, traffic from A still can't route to C.

## Message `RegistryKey` {#RegistryKey}

A registry key.


### Inputs for `RegistryKey`

* `string` [`path`](#RegistryKey.path) = 1 (**Required**)
* `repeated` [`asset.RegistryValue`](asset.md#RegistryValue) [`value`](#RegistryKey.value) = 2

### `path` {#RegistryKey.path}

| Property | Comments |
|----------|----------|
| Field Name | `path` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Path to registry key. Either backslashes or forwardslashes can be used as
separators. The first component of the path selects the hive. The
following values and aliases are accepted fas the first component of the
path:

    HKEY_CLASSES_ROOT, HKCR
    HKEY_CURRENT_USER, HKCU
    HKEY_LOCAL_MACHINE, HKLM
    HKEY_USERS, HKU
    HKEY_CURRENT_CONFIG, HKCC

E.g.:
    key {
      path: 'HKEY_LOCAL_MACHINE\\System\\CurrentControlSet\\Foo\\bar'
      value: {
        name: 'x'
        type: STRING
        value: 'y'
      }
    }

### `value` {#RegistryKey.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`asset.RegistryValue`](asset.md#RegistryValue) |
| Repeated | Any number of instances of this type is allowed in the schema. |

The list of values under the key.

## Message `RegistryValue` {#RegistryValue}

A single registry value. The registry key is implied based on the containing
message.


### Inputs for `RegistryValue`

* `string` [`name`](#RegistryValue.name) = 1
* `string` [`string_value`](#RegistryValue.string_value) = 10
* `string` [`expand_string_value`](#RegistryValue.expand_string_value) = 11
* `bytes` [`binary_value`](#RegistryValue.binary_value) = 12
* `int32` [`dword_value`](#RegistryValue.dword_value) = 13
* `int64` [`qword_value`](#RegistryValue.qword_value) = 14
* [`asset.RegistryValue.MultiString`](asset.md#RegistryValue.MultiString) [`multi_string_value`](#RegistryValue.multi_string_value) = 15

### `name` {#RegistryValue.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |

Value name. Leave empty to set the default value.

### `string_value` {#RegistryValue.string_value}

| Property | Comments |
|----------|----------|
| Field Name | `string_value` |
| Type | `string` |

REG_SZ value in UTF-8. Will be re-encoded into UTF-16 when writing to the
registry.

### `expand_string_value` {#RegistryValue.expand_string_value}

| Property | Comments |
|----------|----------|
| Field Name | `expand_string_value` |
| Type | `string` |

REG_EXPAND_SZ value in UTF-8.

### `binary_value` {#RegistryValue.binary_value}

| Property | Comments |
|----------|----------|
| Field Name | `binary_value` |
| Type | `bytes` |

REG_BINARY

### `dword_value` {#RegistryValue.dword_value}

| Property | Comments |
|----------|----------|
| Field Name | `dword_value` |
| Type | `int32` |

REG_DWORD

### `qword_value` {#RegistryValue.qword_value}

| Property | Comments |
|----------|----------|
| Field Name | `qword_value` |
| Type | `int64` |

REG_QWORD

### `multi_string_value` {#RegistryValue.multi_string_value}

| Property | Comments |
|----------|----------|
| Field Name | `multi_string_value` |
| Type | [`asset.RegistryValue.MultiString`](asset.md#RegistryValue.MultiString) |

REG_MULTI_SZ

## Message `RemoteDesktopHost` {#RemoteDesktopHost}



### Inputs for `RemoteDesktopHost`

* `string` [`windows_machine`](#RemoteDesktopHost.windows_machine) = 1 (**Required**)
* `string` [`collection_name`](#RemoteDesktopHost.collection_name) = 2
* `string` [`collection_description`](#RemoteDesktopHost.collection_description) = 3

### `windows_machine` {#RemoteDesktopHost.windows_machine}

| Property | Comments |
|----------|----------|
| Field Name | `windows_machine` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_machine` |

Machine hosting the remote desktop . Must match the `name` of a
WindowsMachine entry.

### `collection_name` {#RemoteDesktopHost.collection_name}

| Property | Comments |
|----------|----------|
| Field Name | `collection_name` |
| Type | `string` |

The name of the collection to create. This is required for
Windows Server >= 2012 but is ignored for older versions.
RDS Collections were introduced in Windows Server 2012.
More info:
https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/rds-create-collection

### `collection_description` {#RemoteDesktopHost.collection_description}

| Property | Comments |
|----------|----------|
| Field Name | `collection_description` |
| Type | `string` |

The description of the collection to create. This is optional for
Windows Server >= 2012 and is ignored for older versions.

## Message `UserOrGroupReference` {#UserOrGroupReference}

A reference to a user or a group. Provided as a convenience for cases where
either a single user or a group can be specified.


### Inputs for `UserOrGroupReference`

* [`asset.UserReference`](asset.md#UserReference) [`user`](#UserOrGroupReference.user) = 1
* [`asset.GroupReference`](asset.md#GroupReference) [`group`](#UserOrGroupReference.group) = 2

### `user` {#UserOrGroupReference.user}

| Property | Comments |
|----------|----------|
| Field Name | `user` |
| Type | [`asset.UserReference`](asset.md#UserReference) |

### `group` {#UserOrGroupReference.group}

| Property | Comments |
|----------|----------|
| Field Name | `group` |
| Type | [`asset.GroupReference`](asset.md#GroupReference) |

## Message `UserReference` {#UserReference}

A reference to a user. The combination of |name| and |container| must match
one of the WindowsUser entries.


### Inputs for `UserReference`

* `string` [`windows_user`](#UserReference.windows_user) = 1 (**Required**)
* [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#UserReference.container) = 2

### `windows_user` {#UserReference.windows_user}

| Property | Comments |
|----------|----------|
| Field Name | `windows_user` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_user` |

The name of the user.

### `container` {#UserReference.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |

Location. See GroupReference for how the container parameters are
determined by default.

## Message `WindowsContainer` {#WindowsContainer}

Describes a container that a Windows asset can reside in.

Resources like machines, users, and groups can be specified per domain, per
machine, or per organizational unit. When specifying one of these asset
types, use the WindowsContainer member to specify where to create the asset.

Note that this is not related to the Active Directory class named
[Container](https://msdn.microsoft.com/en-us/library/ms680997(v=vs.85).aspx).
The latter is an abstract class for Active Directory objects that can hold
other objects. In theory, any Active Directory object could contain other
objects, but this schema intends to keep things simple any only include a
few possible container objects that typically arise when describing an
enterprise deployment. Feel free to add to this list as more container
object types are needed.


### Inputs for `WindowsContainer`

* `string` [`ad_domain`](#WindowsContainer.ad_domain) = 1 (**Required**)
* `string` [`windows_machine`](#WindowsContainer.windows_machine) = 2 (**Required**)
* `string` [`ad_organizational_unit`](#WindowsContainer.ad_organizational_unit) = 3 (**Required**)

### `ad_domain` {#WindowsContainer.ad_domain}

| Property | Comments |
|----------|----------|
| Field Name | `ad_domain` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_domain` |

Domain name. Users and groups that have an `ad_domain` container are
automatically assigne as tenants of the domain controller for the
associated domain.

### `windows_machine` {#WindowsContainer.windows_machine}

| Property | Comments |
|----------|----------|
| Field Name | `windows_machine` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.windows_machine` |

Machine name.

### `ad_organizational_unit` {#WindowsContainer.ad_organizational_unit}

| Property | Comments |
|----------|----------|
| Field Name | `ad_organizational_unit` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `asset.ad_organizational_unit` |

Organizational unit. Users and groups that have an `ad_domain` container
are automatically assigned as tenants of the domain controller for the
associated domain.

## Message `WindowsGroup` {#WindowsGroup}

Descibes an Active Directory or Windows local group.


### Inputs for `WindowsGroup`

* `string` [`name`](#WindowsGroup.name) = 1 (**Required**)
* `string` [`full_name`](#WindowsGroup.full_name) = 2
* `string` [`description`](#WindowsGroup.description) = 3
* [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#WindowsGroup.container) = 4 (**Required**)

### `name` {#WindowsGroup.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the group. Exclude the domain name. The name alone is not
sufficient if this group corresponds to a Well Known group. Use the
|well_known_sid| field for that.

### `full_name` {#WindowsGroup.full_name}

| Property | Comments |
|----------|----------|
| Field Name | `full_name` |
| Type | `string` |

The actual Unicode Windows group name. Only specify this if the desired
name is different from the |name| field due to not being an RFC 1035
label.

### `description` {#WindowsGroup.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

Human readable description of the group.

### `container` {#WindowsGroup.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |
| Required | This field is required. It is an error to omit this field. |

Container for the group. A container must be specified for a WindowsGroup.

## Message `WindowsMachine` {#WindowsMachine}

A Windows machine.

WindowsMachine messages are located under the "asset.windows_machine"
collection.  When authoring, leave out the "asset." prefix. E.g.:

``` textpb
  windows_machine {
    name: "foo"
    container { ad_domain: "win-domain" }
  }
```


### Inputs for `WindowsMachine`

* `string` [`name`](#WindowsMachine.name) = 1 (**Required**)
* `string` [`machine_type`](#WindowsMachine.machine_type) = 2 (**Required**)
* `repeated` [`asset.NetworkInterface`](asset.md#NetworkInterface) [`network_interface`](#WindowsMachine.network_interface) = 3 (**Required**)
* [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#WindowsMachine.container) = 10
* `string` [`locale`](#WindowsMachine.locale) = 11
* `string` [`timezone`](#WindowsMachine.timezone) = 12
* `repeated` `string` [`windows_feature`](#WindowsMachine.windows_feature) = 13
* `repeated` [`asset.RegistryKey`](asset.md#RegistryKey) [`registry_key`](#WindowsMachine.registry_key) = 14
* [`common.FileReference`](common.md#FileReference) [`configuration_file`](#WindowsMachine.configuration_file) = 15

### Outputs for `WindowsMachine`

* `repeated` `string` [`all_features`](#WindowsMachine.all_features) = 102

## Runtime fields for `WindowsMachine`

* `repeated` `string` [`fqdn`](#WindowsMachine.fqdn) = 100
* `repeated` [`asset.Address`](asset.md#Address) [`address`](#WindowsMachine.address) = 101

### `name` {#WindowsMachine.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the machine. This name will become the hostname for the machine,
both absolute and domain relative (if applicable). Hence must be globally
unique.

For Windows machines, it's advisable to have *short* hostnames, ideally
shorter than 11 characters. This allows the name to do double duty as a
NetBios name as well as a DNS hostname. If the name needs to be longer,
then you can specify a shorter NetBIOS name using the `NetbiosName`
property.

### `machine_type` {#WindowsMachine.machine_type}

| Property | Comments |
|----------|----------|
| Field Name | `machine_type` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `host.machine_type` |

The name of a host.MachineType entry that describes the host machine.

### `network_interface` {#WindowsMachine.network_interface}

| Property | Comments |
|----------|----------|
| Field Name | `network_interface` |
| Type | [`asset.NetworkInterface`](asset.md#NetworkInterface) |
| Repeated | Any number of instances of this type is allowed in the schema. |
| Required | This field is required. It is an error to omit this field. |

Network interfaces. There can be more than one for multihomed machines.
There must be at least one of these.

### `container` {#WindowsMachine.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |

Container. Only |domain| and |ou| values are acceptable. Currently
|machine| is not a valid option. If no container is specified, the machine
will be brought up as a standalone workstation or server depending on the
installed operating system.

This field should be empty for a machine that's referenced in a
ActiveDirectoryDomainController entry.

Specifying this field results in the machine being joined to the specified
domain and, if necessary, placed in the specified container.

### `locale` {#WindowsMachine.locale}

| Property | Comments |
|----------|----------|
| Field Name | `locale` |
| Type | `string` |

System locale. If left unspecified, the default is left unchanged. Use the
following PowerShell command to determine the list of available locales on
a Windows machine:

``` ps1
[System.Globalization.CultureInfo]::GetCultures([System.Globalization.CultureTypes]::AllCultures).name
```

PS DSC Reference:
* https://github.com/PowerShell/SystemLocaleDsc

### `timezone` {#WindowsMachine.timezone}

| Property | Comments |
|----------|----------|
| Field Name | `timezone` |
| Type | `string` |

Set the system timezone. If left unspecified, the default is left
unchanged. Use the following PowerShell command to determine the lsit of
available timezone identifiers on a Windows machine:

``` ps1
[System.TimeZoneInfo]::GetSystemTimeZones().Id
```

PS DSC Reference:
* https://github.com/PowerShell/xTimeZone

### `windows_feature` {#WindowsMachine.windows_feature}

| Property | Comments |
|----------|----------|
| Field Name | `windows_feature` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

List of additional Windows features or roles to install. The values here
should be valid for the selected host machine type. You can use the
'Get-WindowsFeature' PowerShell commandlet to retrieve a list of available
Windows features. E.g.:

``` textpb
windows_feature: "Web-Server"
```

Note: This method cannot be used to specify all sub-features. All
features that needs to be installed should be listed individually and
explicitly.

Note: Addition of roles can cause features to be installed implicitly.
E.g. specifying a machine as the host for an IIS site will automatically
install the necessary web server roles. The |windows_feature| field should
be used for features that otherwise won't be installed as part of any such
role assignment.

Note: Any 'package' values specified in the
host.machine_type.<machine-type> will be prepended to this list.

### `registry_key` {#WindowsMachine.registry_key}

| Property | Comments |
|----------|----------|
| Field Name | `registry_key` |
| Type | [`asset.RegistryKey`](asset.md#RegistryKey) |
| Repeated | Any number of instances of this type is allowed in the schema. |

Registry keys to set on the machine. These will be applied prior to the
user login.

### `configuration_file` {#WindowsMachine.configuration_file}

| Property | Comments |
|----------|----------|
| Field Name | `configuration_file` |
| Type | [`common.FileReference`](common.md#FileReference) |

A configuration file. Specify this if you've run Windows Server Manager
and produced a configration file already.

### `fqdn` {#WindowsMachine.fqdn}

| Property | Comments |
|----------|----------|
| Field Name | `fqdn` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

List of FQDNs that are known to refer to this machine. There can be
multiple if there are aliases configured for this machine.

### `address` {#WindowsMachine.address}

| Property | Comments |
|----------|----------|
| Field Name | `address` |
| Type | [`asset.Address`](asset.md#Address) |
| Repeated | Any number of instances of this type is allowed in the schema. |

IPv4 or IPv6 addresses that refer to this instance. These include both
internal and, if applicable, external addresses.

### `all_features` {#WindowsMachine.all_features}

| Property | Comments |
|----------|----------|
| Field Name | `all_features` |
| Type | `string` |
| Repeated | Any number of instances of this type is allowed in the schema. |

List of all Windows features that must be installed and configured on this
Windows machine. This is a union of `windows_feature` and features that
are required due to additional services hosted on this machine.

## Message `WindowsUser` {#WindowsUser}

Describes a Active Directory or a Windows local user.

WindowsUser messages are located under the "asset.windows_user" collection.
When authoring, leave out the "asset." prefix. E.g.:

``` textpb
  windows_user {
    name: "joe"
    container { ad_domain: "win-domain" }
  }
```


### Inputs for `WindowsUser`

* `string` [`name`](#WindowsUser.name) = 1 (**Required**)
* `string` [`full_name`](#WindowsUser.full_name) = 2
* `string` [`description`](#WindowsUser.description) = 3
* [`asset.WindowsContainer`](asset.md#WindowsContainer) [`container`](#WindowsUser.container) = 4 (**Required**)
* `string` [`hardcoded_password`](#WindowsUser.hardcoded_password) = 5
* `repeated` [`asset.GroupReference`](asset.md#GroupReference) [`member_of`](#WindowsUser.member_of) = 6

### Outputs for `WindowsUser`

* [`common.Secret`](common.md#Secret) [`password`](#WindowsUser.password) = 100

### `name` {#WindowsUser.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

Name of the user. Exclude the domain name.

E.g.: joe

### `full_name` {#WindowsUser.full_name}

| Property | Comments |
|----------|----------|
| Field Name | `full_name` |
| Type | `string` |

The actual Unicode Windows user name. Only specify this if the desired
name is different from the |name| field due to not being an RFC 1035
label.

### `description` {#WindowsUser.description}

| Property | Comments |
|----------|----------|
| Field Name | `description` |
| Type | `string` |

Description.

### `container` {#WindowsUser.container}

| Property | Comments |
|----------|----------|
| Field Name | `container` |
| Type | [`asset.WindowsContainer`](asset.md#WindowsContainer) |
| Required | This field is required. It is an error to omit this field. |

Container for the user. A container must be specified for a WindowsUser.

### `hardcoded_password` {#WindowsUser.hardcoded_password}

| Property | Comments |
|----------|----------|
| Field Name | `hardcoded_password` |
| Type | `string` |

Hardcoded password in UTF-8. Don't use this field.

### `member_of` {#WindowsUser.member_of}

| Property | Comments |
|----------|----------|
| Field Name | `member_of` |
| Type | [`asset.GroupReference`](asset.md#GroupReference) |
| Repeated | Any number of instances of this type is allowed in the schema. |

List of groups that the user belongs to.

### `password` {#WindowsUser.password}

| Property | Comments |
|----------|----------|
| Field Name | `password` |
| Type | [`common.Secret`](common.md#Secret) |



# Enumerations

## Enumeration `FunctionalLevel` {#ActiveDirectoryDomain.FunctionalLevel}

Active Directory functional level. A.k.a. Domain Mode. See
https://docs.microsoft.com/en-us/windows-server/identity/ad-ds/active-directory-functional-levels
for more details on the specific features that are available at each
functional level.


Values:
* ["`DEFAULT`"](#ActiveDirectoryDomain.FunctionalLevel.DEFAULT)
* "`Win2003`"
* "`Win2008`"
* "`Win2008R2`"
* "`Win2012`"
* "`Win2012R2`"
* "`Win2016`"

### "`DEFAULT`" {#ActiveDirectoryDomain.FunctionalLevel.DEFAULT}

Use the default. The default functional level depends on the host OS and
on the other AD DS servers in the domain or forest.


## Enumeration `Action` {#ActiveDirectoryRegistryPrefPolicy.Action}

The action to take when processing the GPO.


Values:
* "`CREATE`"
* "`REPLACE`"
* "`UPDATE`"
* "`DELETE`"

## Enumeration `Protocol` {#Protocol}

List of web protocols. Not necessarily exhaustive. We just need a convenient
enum so that we don't need to defined it everywhere. Values should be self
explanatory.


Values:
* "`UNKNOWN`"
* "`HTTP`"
* "`HTTPS`"


---
Generated from `schema/asset/network.proto`, `schema/asset/machine.proto`, `schema/asset/active_directory.proto`, `schema/asset/cert.proto`, `schema/asset/dns.proto`, `schema/asset/iis.proto`, `schema/asset/remote_desktop.proto`, `schema/asset/asset_manifest.proto`.
