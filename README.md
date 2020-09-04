# Chrome Enterprise Lab

Chrome Enterprise Lab is an inaccurately named set of tools for building
enterprise labs quickly and easily. The labs so built can be used for system
level end-to-end testing of Google Chrome/Chromium.

Have a peek at the [Design](/docs/design-overview.md) document.

Also have a peek at the [Code of Conduct](./CODE_OF_CONDUCT.md).

Most of the code is in Go. See [Guide to code](/docs/guide-to-code.md) to get a
head start on the code in this repository.

## Building

### Prerequisites

* [Go][]: Download Go from the [Go download page][]. **This project requires Go
    1.9.4 or later**. The build will fail with prior versions.

* [Depot Tools][]: Used for managing the checkout and the contributor workflow.

* [Protocol Buffers Compiler][]: Protocol buffers are used extensively for
    wranging all the data that needs to be shuttled around. Run `protoc --version` to
    check the protoc version. The protoc version should be at least 3.5.1. If
    protoc is not installed, or if the version is less than 3.5.1, download it from
    https://github.com/protocolbuffers/protobuf/releases.

    Remember that you need to copy the contents of the `include`
    directory of the zip file into `/usr/local/include/` as well, i.e.
    ```
    >> cd tmp
    >> wget https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip
    >> unzip protoc-3.13.0-linux-x86_64.zip -d protoc
    >> sudo cp protoc/bin/protoc /usr/local/bin/protoc
    >> sudo chmod 751 /usr/local/bin/protoc
    >> cp protoc/
    >> sudo mv include/* /usr/local/include/
    >> rm -rf protoc3 protoc-3.13.0-linux-x86_64.zip
    ```

* [Dep][] : Used for Go depedency management. This can be installed
    automatically by running `./build.py deps --install` from the root of the
    source tree.

* [Go support for Protocol Buffers][] : This can be installed automatically by
  running `./build.py deps --install` from the root of the source tree.

* [absl-py][]: This Python package is used by tests. Install it by running
  `pip install absl-py`.

[Go]: https://golang.org/
[Go download page]: https://golang.org/dl/
[Depot Tools]: https://dev.chromium.org/developers/how-tos/install-depot-tools
[Protocol Buffers Compiler]: https://developers.google.com/protocol-buffers/
[Dep]: https://github.com/golang/dep
[Go support for Protocol Buffers]: https://github.com/golang/protobuf
[absl-py]: https://pypi.org/project/absl-py/

### Get The Source

There are two ways to get the source. One is to use managed deps, and the other
is to use plain `go get`. The latter workflow doesn't quite work yet due to this
repository not being integrated with `go.chromium.org`. So this page only
mentions the managed dependency workflow.

1. Clone this repository:

   Assumes that `$GOPATH` is a single path and not a `:` delimited list.

   ``` sh
   mkdir -p ${GOPATH}/src/chromium.googlesource.com/enterprise
   cd ${GOPATH}/src/chromium.googlesource.com/enterprise
   git clone https://chromium.googlesource.com/enterprise/cel
   cd cel
   ```

2. Get the dependencies:

   ``` sh
   python build.py deps --install
   ```
   If you see an error message complaining about protoc missing, see the protoc
   prerequisites above.

### Build It

1. Use the build script:

   ``` sh
   python build.py build
   ```

   If you see an error message like `google/protobuf/descriptor.proto: File not found`,
   see the protoc prerequisites (specifically the includes step).

2. Also make sure that all unit tests for the Go source files are passing.

   ``` sh
   python build.py test
   ```

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md).

## Release

The framework is uploaded into [CIPD](https://github.com/luci/luci-go/tree/master/cipd) as a package.

Here are the general steps to release the package.
1. After the code review, merge in the change
2. Wait for windows CI build finished (https://ci.chromium.org/p/celab/g/main/console)
3. Run upload_to_cipd with the link to binaries in the CI build (instructions here: https://chromium.googlesource.com/enterprise/cel/+/refs/heads/master/scripts/cipd)
4. (Optional) Check the output of that command and update the infra version in the .vpython file for your app
