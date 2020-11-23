# Go Style Guide

In general follow [Effective Go](https://golang.org/doc/effective_go.html).

## Notes on Google Cloud APIs

Most of the interactions we do with Google Cloud is via the REST APIs, and the
Go API bindings at [google-api-go-client][]. In addition, where a stable API is
available, we use the [google-cloud-go][] APIs. The latter is more idiomatic Go
and is the recommended options where available. We should consider migrating as
more APIs are made available in [google-cloud-go][].

[google-api-go-client]: https://github.com/google/google-api-go-client
[google-cloud-go]: https://github.com/GoogleCloudPlatform/google-cloud-go

## Dependencies

Currently we use `dep` to manage dependencies. The build scripts should invoke
`dep` if it already exists. If not, dependencies may be installed in your
`GOPATH`.

See <https://github.com/golang/dep> for details on dependency management using
`dep`.
