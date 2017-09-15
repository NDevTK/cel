# Chrome Enterprise Lab

This directory and its subdirectories contain the code for administering an
instance of the Chrome Enterprise Lab.

## Notes on Google Cloud APIs.

Most of the interactions we do with Google Cloud is via the REST APIs, and the
Go API bindings at [google-api-go-client][]. In addition, where a stable API is
available, we use the [google-cloud-go][] APIs. The latter is more idiomatic Go
and is the recommended options where available. We should consider migrating as
more APIs are made available in [google-cloud-go][].

[google-api-go-client]: https://github.com/google/google-api-go-client
[google-cloud-go]: https://github.com/GoogleCloudPlatform/google-cloud-go

## Building

Use the [build.py][../../build.py] script at the root of the source tree to
build the Go code.

``` sh
./build.py build
```

## Testing

To run the tests:

``` sh
./build.py test
```

## Environment

The build environment needs to be customized to allow for the directory
structure used by the lab code. If you would like the same environment during
editing and debugging the Go code, you can use:

``` sh
./build.py run bash
```

... which will start `bash` running under the build time Go environment.
Alternatively, you can do something like ...

``` sh
./build.py run vim
```

... to start your editor with the correct environment.
