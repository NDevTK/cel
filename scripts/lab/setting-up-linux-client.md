Setting up the Linux client
===========================

The linux client VM is by default named `linux-client` and is based on the most
recent Ubuntu image availabe on GCE.

In order to prepare the VM, run the following commands:

1. Install `crcmod`:

       ```sh
       # Lookup the latest instructions:
       gsutil help crcmod

       sudo apt-get update

       # And follow them:
       sudo apt-get install gcc python-dev python-setuptools
       sudo easy_install -U pip
       sudo pip uninstall crcmod
       sudo pip install -U crcmod
       ```
2. Install `unzip`:

       ```sh
       sudo apt-get install unzip
       ```

3. Get the `chrome-auth-lab-tools`:

       ```sh
       gsutil -m rsync -d gs://chrome-auth-lab-staging/tools $HOME/tools
       ```

4. Follow the instructions in `~/tools/README.md` to sync and extract the latest
   set of binaries from the staging area.


