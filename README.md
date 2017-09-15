# Introduction

Chrome Enterprise Lab is a virtual lab hosted on Google Compute Engine that can
be used for automated and manual tests of Chrome binaries.

Have a peek at the [Design](http://goto.google.com/chrome-enterprise-lab) document.

Note that this repository uses submodules. After checking out the repository
run:

``` sh
git submodule update --init --recursive
```

Most of the code is in Go. See the [README](/src/go/README.md/) file.
