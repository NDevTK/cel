# Schema `common` {#common}



Messages that are valid in package `common` are as follows:

*** note
Note that this document uses the term "message" to refer to the same concept as
a "message" in Protocol Buffers. Hence every asset and host resource
description is a *message*. So is their embedded structures.
***

## Message `FileReference` {#FileReference}

A file reference. Use this when the configuration should refer to a file or
a directory that's specified as part of the configuration. All such paths
are resolved *relative to the source file* in which the reference appears.


### Inputs for `FileReference`

* `string` [`source`](#FileReference.source) = 1
* `string` [`target_path`](#FileReference.target_path) = 2

### Outputs for `FileReference`

* `string` [`full_path`](#FileReference.full_path) = 3
* `string` [`object_reference`](#FileReference.object_reference) = 4
* `string` [`integrity`](#FileReference.integrity) = 5
* [`common.FileReference.Type`](common.md#FileReference.Type) [`resolved_type`](#FileReference.resolved_type) = 6

### `source` {#FileReference.source}

| Property | Comments |
|----------|----------|
| Field Name | `source` |
| Type | `string` |

Path relative to the location of the .textpb file in which the reference
appears. Absolute paths are not allowed.

  * Path separators are always forward slashes.
  * No parent directory traversal. I.e. ".." is an invalid path component.
  * '.' is a valid component, but you might as well leave it out.
  * The referenced path must exist.

This, of course, means that textpb files that refer to other files should
live higher up in the directory tree or live beside the relevant files.

Note: This field is only valid during authoring. It is cleared when
constructing the Completed Asset Manifest.

### `target_path` {#FileReference.target_path}

| Property | Comments |
|----------|----------|
| Field Name | `target_path` |
| Type | `string` |

The path where the target of this FileReference should be copied to on the
target machine. This field is optional and is only meaningful when a file
or directory should be placed in a specific location on a machine. The
machine is implied based on the schema containing the FileReference.

The containing schema also determines how the contents of the
FileReference is treated. In cases where the local path doesn't matter
then an indeterminate temporary path may be used.

Must be a full path or empty.

### `full_path` {#FileReference.full_path}

| Property | Comments |
|----------|----------|
| Field Name | `full_path` |
| Type | `string` |

The full path to file or directory contents being referred.

During authoring the deployer sets this field to the full path to the
source.

Within the lab, this field is used to store the full path to the location
where the referred files or directories are stored on the local file
system. This could be the same as target_path if the latter was not empty
and valid.

### `object_reference` {#FileReference.object_reference}

| Property | Comments |
|----------|----------|
| Field Name | `object_reference` |
| Type | `string` |

Opaque reference to an accessible location where the contents of the file
or folder could be found. For GCP, this would be a Google Cloud Storage
URL.

This reference is understood by the ObjectStorage service.

### `integrity` {#FileReference.integrity}

| Property | Comments |
|----------|----------|
| Field Name | `integrity` |
| Type | `string` |

Subresource integrity string.

This string is required for all deployed objects. On the client side, the
object is rejected if the integrity check fails. Currently only SHA-384
digests are supported.

See https://w3c.github.io/webappsec-subresource-integrity/

### `resolved_type` {#FileReference.resolved_type}

| Property | Comments |
|----------|----------|
| Field Name | `resolved_type` |
| Type | [`common.FileReference.Type`](common.md#FileReference.Type) |

## Message `MapFieldEntry` {#TestMessageWithTypes.MapFieldEntry}



### Inputs for `MapFieldEntry`

* `string` [`key`](#TestMessageWithTypes.MapFieldEntry.key) = 1
* [`common.TestGoodProto`](common.md#TestGoodProto) [`value`](#TestMessageWithTypes.MapFieldEntry.value) = 2

### `key` {#TestMessageWithTypes.MapFieldEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#TestMessageWithTypes.MapFieldEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |

## Message `MapStringEntry` {#TestMessageWithTypes.MapStringEntry}



### Inputs for `MapStringEntry`

* `string` [`key`](#TestMessageWithTypes.MapStringEntry.key) = 1
* `string` [`value`](#TestMessageWithTypes.MapStringEntry.value) = 2

### `key` {#TestMessageWithTypes.MapStringEntry.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |

### `value` {#TestMessageWithTypes.MapStringEntry.value}

| Property | Comments |
|----------|----------|
| Field Name | `value` |
| Type | `string` |

## Message `Secret` {#Secret}

A secret. Could be a password or a private key. The enclosing object is
responsible for storing the secret as a string of octets in the `final`
field.

The Secret's message handlers will then store the contents in ObjectStorage
and store the resulting reference in the `object_reference` field. The
contents will be encrypted using a key that'll be available to the instance
VMs in the lab.


### Inputs for `Secret`

* `string` [`hardcoded`](#Secret.hardcoded) = 1

### Outputs for `Secret`

* `bytes` [`final`](#Secret.final) = 2
* `string` [`object_reference`](#Secret.object_reference) = 3

### `hardcoded` {#Secret.hardcoded}

| Property | Comments |
|----------|----------|
| Field Name | `hardcoded` |
| Type | `string` |

The hardcoded secret as a string. This field is typically only used when
storing a hardcoded password. The `final` field will contain a UTF-8
encoding of the string.

Use of this field is discouraged.

If this field is not set or empty, the enclosing object is responsible for
generating a suitable secret and storing it in the `final` field during
the `ResolveGeneratedContent` phase.

### `final` {#Secret.final}

| Property | Comments |
|----------|----------|
| Field Name | `final` |
| Type | `bytes` |

The actual secret. In the case of a password, this field will contain a
UTF-8 encoded string. Otherwise it will contain whatever sequence of bytes
that was stored in here by the enclosing object.

### `object_reference` {#Secret.object_reference}

| Property | Comments |
|----------|----------|
| Field Name | `object_reference` |
| Type | `string` |

Once the secret is encoded and stored in object storage, the resulting
object reference is stored here.

## Message `TestBadMessageWithOptions` {#TestBadMessageWithOptions}



### Inputs for `TestBadMessageWithOptions`

* `int32` [`name`](#TestBadMessageWithOptions.name) = 1 (**Required**)

### `name` {#TestBadMessageWithOptions.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `int32` |
| Required | This field is required. It is an error to omit this field. |

## Message `TestBadOneOf` {#TestBadOneOf}



### Inputs for `TestBadOneOf`

* `string` [`name`](#TestBadOneOf.name) = 1 (**Required**)
* [`common.TestBadProto`](common.md#TestBadProto) [`field`](#TestBadOneOf.field) = 2

### `name` {#TestBadOneOf.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field` {#TestBadOneOf.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestBadProto`](common.md#TestBadProto) |

## Message `TestBadProto` {#TestBadProto}



### Inputs for `TestBadProto`

* `string` [`name`](#TestBadProto.name) = 1 (**Required**)

### `name` {#TestBadProto.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

## Message `TestBadReturnType` {#TestBadReturnType}



### Inputs for `TestBadReturnType`

* `string` [`name`](#TestBadReturnType.name) = 2 (**Required**)

### `name` {#TestBadReturnType.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

## Message `TestBadValidateArgs` {#TestBadValidateArgs}



### Inputs for `TestBadValidateArgs`

* `string` [`name`](#TestBadValidateArgs.name) = 1 (**Required**)

### `name` {#TestBadValidateArgs.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

## Message `TestContainer` {#TestContainer}



### Inputs for `TestContainer`

* `repeated` [`common.TestGoodProto`](common.md#TestGoodProto) [`a`](#TestContainer.a) = 1
* `repeated` [`common.TestMessageWithOptions`](common.md#TestMessageWithOptions) [`b`](#TestContainer.b) = 2

### `a` {#TestContainer.a}

| Property | Comments |
|----------|----------|
| Field Name | `a` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `b` {#TestContainer.b}

| Property | Comments |
|----------|----------|
| Field Name | `b` |
| Type | [`common.TestMessageWithOptions`](common.md#TestMessageWithOptions) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `TestFileRefProto` {#TestFileRefProto}



### Inputs for `TestFileRefProto`

* [`common.FileReference`](common.md#FileReference) [`ref`](#TestFileRefProto.ref) = 1

### `ref` {#TestFileRefProto.ref}

| Property | Comments |
|----------|----------|
| Field Name | `ref` |
| Type | [`common.FileReference`](common.md#FileReference) |

## Message `TestGoodOneOf` {#TestGoodOneOf}



### Inputs for `TestGoodOneOf`

* `string` [`name`](#TestGoodOneOf.name) = 1 (**Required**)
* [`common.TestGoodProto`](common.md#TestGoodProto) [`field`](#TestGoodOneOf.field) = 2

### `name` {#TestGoodOneOf.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field` {#TestGoodOneOf.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |

## Message `TestGoodProto` {#TestGoodProto}



### Inputs for `TestGoodProto`

* `string` [`name`](#TestGoodProto.name) = 1 (**Required**)

### `name` {#TestGoodProto.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

## Message `TestHasBadField` {#TestHasBadField}



### Inputs for `TestHasBadField`

* `string` [`name`](#TestHasBadField.name) = 1 (**Required**)
* [`common.TestBadProto`](common.md#TestBadProto) [`field`](#TestHasBadField.field) = 2

### `name` {#TestHasBadField.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field` {#TestHasBadField.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestBadProto`](common.md#TestBadProto) |

## Message `TestHasBadSlice` {#TestHasBadSlice}



### Inputs for `TestHasBadSlice`

* `string` [`name`](#TestHasBadSlice.name) = 1 (**Required**)
* `repeated` [`common.TestBadProto`](common.md#TestBadProto) [`field`](#TestHasBadSlice.field) = 2

### `name` {#TestHasBadSlice.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field` {#TestHasBadSlice.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestBadProto`](common.md#TestBadProto) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `TestHasGoodField` {#TestHasGoodField}



### Inputs for `TestHasGoodField`

* `string` [`name`](#TestHasGoodField.name) = 1 (**Required**)
* [`common.TestGoodProto`](common.md#TestGoodProto) [`field`](#TestHasGoodField.field) = 2

### `name` {#TestHasGoodField.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field` {#TestHasGoodField.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |

## Message `TestHasGoodSlice` {#TestHasGoodSlice}



### Inputs for `TestHasGoodSlice`

* `string` [`name`](#TestHasGoodSlice.name) = 1 (**Required**)
* `repeated` [`common.TestGoodProto`](common.md#TestGoodProto) [`field`](#TestHasGoodSlice.field) = 2

### `name` {#TestHasGoodSlice.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field` {#TestHasGoodSlice.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `TestMessageWithOptions` {#TestMessageWithOptions}



### Inputs for `TestMessageWithOptions`

* `string` [`name`](#TestMessageWithOptions.name) = 1 (**Required**)
* `string` [`key`](#TestMessageWithOptions.key) = 2 (**Required**)
* `string` [`label`](#TestMessageWithOptions.label) = 3 (**Required**)
* `string` [`optional_key`](#TestMessageWithOptions.optional_key) = 4 (**Required**)
* `string` [`fqdn`](#TestMessageWithOptions.fqdn) = 5
* `string` [`reqd`](#TestMessageWithOptions.reqd) = 6 (**Required**)
* `string` [`optional_string`](#TestMessageWithOptions.optional_string) = 7

### Outputs for `TestMessageWithOptions`

* `string` [`output`](#TestMessageWithOptions.output) = 8
* `string` [`output_alt`](#TestMessageWithOptions.output_alt) = 9
* `int32` [`output_int`](#TestMessageWithOptions.output_int) = 10
* [`common.TestGoodProto`](common.md#TestGoodProto) [`output_proto`](#TestMessageWithOptions.output_proto) = 11

### `name` {#TestMessageWithOptions.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `key` {#TestMessageWithOptions.key}

| Property | Comments |
|----------|----------|
| Field Name | `key` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `a.b.with_types.repeated_field` |

### `label` {#TestMessageWithOptions.label}

| Property | Comments |
|----------|----------|
| Field Name | `label` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `optional_key` {#TestMessageWithOptions.optional_key}

| Property | Comments |
|----------|----------|
| Field Name | `optional_key` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |
| Reference | The value of this field is a named reference to a `a.b.with_types.repeated_field` |

### `fqdn` {#TestMessageWithOptions.fqdn}

| Property | Comments |
|----------|----------|
| Field Name | `fqdn` |
| Type | `string` |

### `reqd` {#TestMessageWithOptions.reqd}

| Property | Comments |
|----------|----------|
| Field Name | `reqd` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `optional_string` {#TestMessageWithOptions.optional_string}

| Property | Comments |
|----------|----------|
| Field Name | `optional_string` |
| Type | `string` |

### `output` {#TestMessageWithOptions.output}

| Property | Comments |
|----------|----------|
| Field Name | `output` |
| Type | `string` |

### `output_alt` {#TestMessageWithOptions.output_alt}

| Property | Comments |
|----------|----------|
| Field Name | `output_alt` |
| Type | `string` |

### `output_int` {#TestMessageWithOptions.output_int}

| Property | Comments |
|----------|----------|
| Field Name | `output_int` |
| Type | `int32` |

### `output_proto` {#TestMessageWithOptions.output_proto}

| Property | Comments |
|----------|----------|
| Field Name | `output_proto` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |

## Message `TestMessageWithTypes` {#TestMessageWithTypes}



### Inputs for `TestMessageWithTypes`

* `string` [`name`](#TestMessageWithTypes.name) = 1 (**Required**)
* `bool` [`bool_value`](#TestMessageWithTypes.bool_value) = 2
* `int32` [`int_value`](#TestMessageWithTypes.int_value) = 3
* [`common.TestGoodProto`](common.md#TestGoodProto) [`field`](#TestMessageWithTypes.field) = 4
* `repeated` [`common.TestGoodProto`](common.md#TestGoodProto) [`repeated_field`](#TestMessageWithTypes.repeated_field) = 5
* [`common.TestGoodProto`](common.md#TestGoodProto) [`optional_field`](#TestMessageWithTypes.optional_field) = 6
* `repeated` [`common.TestMessageWithTypes.MapFieldEntry`](common.md#TestMessageWithTypes.MapFieldEntry) [`map_field`](#TestMessageWithTypes.map_field) = 7
* `repeated` [`common.TestMessageWithTypes.MapStringEntry`](common.md#TestMessageWithTypes.MapStringEntry) [`map_string`](#TestMessageWithTypes.map_string) = 8

### `name` {#TestMessageWithTypes.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `bool_value` {#TestMessageWithTypes.bool_value}

| Property | Comments |
|----------|----------|
| Field Name | `bool_value` |
| Type | `bool` |

### `int_value` {#TestMessageWithTypes.int_value}

| Property | Comments |
|----------|----------|
| Field Name | `int_value` |
| Type | `int32` |

### `field` {#TestMessageWithTypes.field}

| Property | Comments |
|----------|----------|
| Field Name | `field` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |

### `repeated_field` {#TestMessageWithTypes.repeated_field}

| Property | Comments |
|----------|----------|
| Field Name | `repeated_field` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `optional_field` {#TestMessageWithTypes.optional_field}

| Property | Comments |
|----------|----------|
| Field Name | `optional_field` |
| Type | [`common.TestGoodProto`](common.md#TestGoodProto) |

### `map_field` {#TestMessageWithTypes.map_field}

| Property | Comments |
|----------|----------|
| Field Name | `map_field` |
| Type | [`common.TestMessageWithTypes.MapFieldEntry`](common.md#TestMessageWithTypes.MapFieldEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

### `map_string` {#TestMessageWithTypes.map_string}

| Property | Comments |
|----------|----------|
| Field Name | `map_string` |
| Type | [`common.TestMessageWithTypes.MapStringEntry`](common.md#TestMessageWithTypes.MapStringEntry) |
| Repeated | Any number of instances of this type is allowed in the schema. |

## Message `TestMultipleOneOf` {#TestMultipleOneOf}



### Inputs for `TestMultipleOneOf`

* `string` [`name`](#TestMultipleOneOf.name) = 1 (**Required**)
* `string` [`field2`](#TestMultipleOneOf.field2) = 2
* `string` [`field3`](#TestMultipleOneOf.field3) = 3
* `string` [`field4`](#TestMultipleOneOf.field4) = 4
* `string` [`field5`](#TestMultipleOneOf.field5) = 5
* [`common.TestBadProto`](common.md#TestBadProto) [`field6`](#TestMultipleOneOf.field6) = 6
* [`common.TestBadProto`](common.md#TestBadProto) [`field7`](#TestMultipleOneOf.field7) = 7

### `name` {#TestMultipleOneOf.name}

| Property | Comments |
|----------|----------|
| Field Name | `name` |
| Type | `string` |
| Required | This field is required. It is an error to omit this field. |

### `field2` {#TestMultipleOneOf.field2}

| Property | Comments |
|----------|----------|
| Field Name | `field2` |
| Type | `string` |

### `field3` {#TestMultipleOneOf.field3}

| Property | Comments |
|----------|----------|
| Field Name | `field3` |
| Type | `string` |

### `field4` {#TestMultipleOneOf.field4}

| Property | Comments |
|----------|----------|
| Field Name | `field4` |
| Type | `string` |

### `field5` {#TestMultipleOneOf.field5}

| Property | Comments |
|----------|----------|
| Field Name | `field5` |
| Type | `string` |

### `field6` {#TestMultipleOneOf.field6}

| Property | Comments |
|----------|----------|
| Field Name | `field6` |
| Type | [`common.TestBadProto`](common.md#TestBadProto) |

### `field7` {#TestMultipleOneOf.field7}

| Property | Comments |
|----------|----------|
| Field Name | `field7` |
| Type | [`common.TestBadProto`](common.md#TestBadProto) |

## Message `Validation` {#Validation}



### Inputs for `Validation`

* [`common.Validation.FieldType`](common.md#Validation.FieldType) [`type`](#Validation.type) = 1
* `string` [`ref`](#Validation.ref) = 2
* `bool` [`optional`](#Validation.optional) = 3

### `type` {#Validation.type}

| Property | Comments |
|----------|----------|
| Field Name | `type` |
| Type | [`common.Validation.FieldType`](common.md#Validation.FieldType) |

### `ref` {#Validation.ref}

| Property | Comments |
|----------|----------|
| Field Name | `ref` |
| Type | `string` |

The foreign key. If non-empty, designates that this string field is a
reference to a collection designated by `ref`.

E.g.: Declare `network` to be a field that refers to an `asset.network` by
name.

``` proto
  message Foo {
    string network = 1 [(v).ref="asset.network"]
  }
```

The annotated field must be a string.

### `optional` {#Validation.optional}

| Property | Comments |
|----------|----------|
| Field Name | `optional` |
| Type | `bool` |



# Enumerations

## Enumeration `Type` {#FileReference.Type}

Type of object.


Values:
* "`UNDEFINED`"
* ["`FILE`"](#FileReference.Type.FILE)
* ["`ZIP_ARCHIVE`"](#FileReference.Type.ZIP_ARCHIVE)

### "`FILE`" {#FileReference.Type.FILE}

A regular file. The object referenced by object_reference contains the
contents of the file.


### "`ZIP_ARCHIVE`" {#FileReference.Type.ZIP_ARCHIVE}

A directory tree. The object referenced by object_reference is a Zipped
archive which must be expanded at the location specified by target_path.
All paths in the Zip files are considered to be relative to target_path.

In the event that target_path is not specified, the archive contents are
expanded to a location determined based on the integrity string.


## Enumeration `FieldType` {#Validation.FieldType}



Values:
* ["`UNKNOWN`"](#Validation.FieldType.UNKNOWN)
* ["`REQUIRED`"](#Validation.FieldType.REQUIRED)
* ["`LABEL`"](#Validation.FieldType.LABEL)
* ["`FQDN`"](#Validation.FieldType.FQDN)
* ["`ORGLABEL`"](#Validation.FieldType.ORGLABEL)
* ["`OUTPUT`"](#Validation.FieldType.OUTPUT)
* ["`RUNTIME`"](#Validation.FieldType.RUNTIME)
* ["`TOPLEVEL`"](#Validation.FieldType.TOPLEVEL)

### "`UNKNOWN`" {#Validation.FieldType.UNKNOWN}

No validation is to be performed.


### "`REQUIRED`" {#Validation.FieldType.REQUIRED}

When applied to a `string` field, implies that the field value cannot be
empty. This is implied for any field that has a non-empty external
reference (i.e. `ref != ""`), or the validation type is
[LABEL](#Validation.FieldType.LABEL), [FQDN](#Validation.FieldType.FQDN)
or [ORGLABEL](#Validation.FieldType.ORGLABEL).

When applied to a `repeated` field, implies that there must be at least
one instance of the field.

When applied to `oneof` implies that at least one of the alternatives
must be specified.

When applied to a `map` implies that there should be at least one
mapping.

E.g.: Declare the `h` field of Foo message to be required:

``` proto
  message Foo {
    string h = 1 [(v).type=REQUIRED]
  }
```


### "`LABEL`" {#Validation.FieldType.LABEL}

The field value cannot be empty and must match the `<label>` production
in [RFC 1035][]. This validation type is applied by default for for any
field named 'name'.

[RFC 1035]: https://www.ietf.org/rfc/rfc1035.txt


### "`FQDN`" {#Validation.FieldType.FQDN}

The field value cannot be empty, and must match the `<subdomains>`
production in [RFC 1035][]. Can only be applied to `string` fields.


### "`ORGLABEL`" {#Validation.FieldType.ORGLABEL}

A label with an optional org component. These look like:
`example.com:foo`. Cannot be empty.


### "`OUTPUT`" {#Validation.FieldType.OUTPUT}

This is an output field and is not expected to be populated in a asset
manifest input. The field will be populated during the deployment
process and made available to downstream consumers of the manifest.


### "`RUNTIME`" {#Validation.FieldType.RUNTIME}

This is a runtime field. It's value is only available *after* the
corresponding asset has been deployed and running. Runtime values can be
looked up via RuntimeConfiguration service from within the lab.

Unlike [OUTPUT](#Validation.FieldType.OUTPUT) fields,
[RUNTIME](#Validation.FieldType.RUNTIME) fields can only be defined on
top-level assets. In other words, they can only appear if the asset in
question is a direct child of the HostEnvironment or AssetManifest
messages.


### "`TOPLEVEL`" {#Validation.FieldType.TOPLEVEL}

This is a top-level field. This type is applicable only to fields that
constititue top level collections. Any element in a top level collection
that doesn't have a rooted reference will be removed during the pruning
phase.

See description of pruning in
https://chromium.googlesource.com/enterprise/cel/+/HEAD/docs/deployment.md#Pruning.



---
Generated from `schema/common/validation.proto`, `schema/common/file_reference.proto`, `schema/common/secret.proto`, `schema/common/testmsgs.proto`.
