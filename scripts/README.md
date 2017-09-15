Chrome Auth Lab
===============

These tools are meant to make your life easier if you are using Google Cloud
Storage for transferring files to this machine.

To update the tools:

``` sh
gsutil -m rsync -d gs://chrome-auth-lab-staging/tools C:\Tools\chrome-auth-lab
```

Or:

``` sh
gsutil -m rsync -d gs://chrome-auth-lab-staging/tools $HOME/tools
```

If you edit anything here, push to the project source repository and also run
the following to sync the storage bucket. The VMs don't have access to the
source repository:

``` sh
gsutil -m rsync -d . gs://chrome-auth-lab-staging/tools
```

