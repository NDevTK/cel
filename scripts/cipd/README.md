The script `upload_to_cipd.py` is used to upload cel releases to CIPD.

## CIPD structure
The structure of celab stuff on cipd is:

```
infra/celab                      The root of everything celab related
     ├── celab
     │     ├── windows-amd64     celab release for Windows
     │     └── linux-amd64       celab release for Linux
```

In the future, there will be other directories to contain 3rd party dependecies, e.g.
`infra/celab/dep/choco` where chocolatey packages are stored, `infra/celab/dep/ps`
where PowerShell DSC packages are stored, etc.

## Script usage

The script relies on `cipd` so make sure that `depot_tools` is already installed.

- To upload a build to CIPD and mark it as the latest, run:
  ```
  python upload_to_cipd.py \
    --input_file=gs://celab/Windows/2019/05/24/8912558978083449328/cel.zip \
    --platform=windows-amd64
  ```