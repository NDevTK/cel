# Generating Documentation

Documentation for schema is generated based on the `.proto` files using the
following method:

* During the file generation process, the build script invokes the ProtoBuf
  compiler (`protoc`) with the following options:

  ```
  protoc --descriptor_set_out=<path> --include_source_info ...
  ```

  These options ensure that a file at `path` is created containing a binary
  encoded [`FileDescriptorProtoSet`][descriptor.proto] buffer.

* At a later stage, the build script invokes the code in this directory (called
  `gen_doc_proto`). `gen_doc_proto` in turn reads in the
  `FileDescriptorProtoSet` and writes out the documentation to a set of `.md`
  files rooted at the output directory.

The `FileDescriptorProtoSet` contains a parsed representation of the `.proto`
files that comprise the Chrome Enterprise Lab schema. This data is used to
compile and output the documentation.

[descriptor.proto]: https://github.com/golang/protobuf/blob/master/protoc-gen-go/descriptor/descriptor.proto

## TODO List

* Generate documentation on validation constraints and type information.

* Autolinkify cross references between symbols in comments.

* Include a summary and a TOC for easier navigation.


