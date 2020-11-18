As part of its setup, Celab install some third party dependencies on the test
client machines using Chocolatey (https://chocolatey.org/). To ensure stable
installs, celab caches these pip packages in CIPD, and download them from CIPD
during setup. This directory contains the CIPD .yaml files for uploading and
updating Chocolatey third-party deps.

## .yaml file usage.

Run 'cipd -create --pkg-def <yaml file path>' to update a package with a new
version. The command will output the new version id. Run 'cipd set-ref
<cipd path> -version <version id> -sef ref <actual version string>' to update
the new version with its actual, human readable version string (e.g. 2.7.15).
Run 'cipd set-ref <cipd path> -version <version id> -sef ref latest' to set the
'latest' ref on the new package version.
