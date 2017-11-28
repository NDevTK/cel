# Chrome Enterprise Lab

Chrome Enterprise Lab is an inaccurately named set of tools for building
enterprise labs quickly and easily. The labs so built can be used for system
level end-to-end testing of Google Chrome/Chromium.

Have a peek at the [Design](/docs/design-overview.md) document.

Also have a peek at the [Code of Conduct](./CODE_OF_CONDUCT.md).

Most of the code is in Go. See the [README](/src/go/README.md/) file.

## Building

### Prerequisites

* [Go](https://golang.org/): Download Go from the [Go download
    page](https://golang.org/dl/).

* [Depot Tools](https://dev.chromium.org/developers/how-tos/install-depot-tools)

* [Protocol Buffers Compiler](https://developers.google.com/protocol-buffers/)

* [Dep](https://github.com/golang/dep) : This can be installed automatically by
    running `./build.py deps --install` from the root of the source tree.

* [Go support for Protocol Buffers](https://github.com/golang/protobuf) : This
    can be installed automatically by running `./build.py deps --install` from
    the root of the source tree.

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

### Build It

1. Use the build script:

   ``` sh
   python build.py build
   ```

2. Also make sure the tests pass.

   ``` sh
   python build.py test
   ```

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md).

