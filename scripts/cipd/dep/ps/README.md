Celab install Powershell modules on the test client machines as part of its
setup. Celab caches these Powershell modules in CIPD. This directory contains
the CIPD .yaml files for uploading and updating Powershell modules.

## .yaml file usage.

Run 'cipd -create --pkg-def <yaml file path>' to update a package with a new
version. The command will output the new version id. Run 'cipd set-ref
<cipd path> -version <version id> -sef ref <actual version string>' to update
the new version with its actual, human readable version string (e.g. 2.19.0.).
