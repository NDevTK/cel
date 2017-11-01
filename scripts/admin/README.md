# Informational Python Scripts

*** note
**Note:** The scripts in this directory are meant to be informational. The logic
there should be recreated in `src/go`.
***

The python scripts in this directory requires the following Python modules to be
installed:

  * google-api-python-client
  * google-cloud
  * pycrypto

Installing these modules for Python 2.7 on Windows also require the installation
of Microsoft Visual C++ Compiler For Python 2.7. Yes there is such a thing. You
can find it [here][vcforpython27]. In case that link doesn't work, search for
`vcforpython27`.

[vcforpython27]: https://www.microsoft.com/en-us/download/confirmation.aspx?id=44266

Install Microsoft Visual C++ Compiler For Python 2.7 into its default
installation path.

Install `pip`, which in turn could be installed via the `get-pip.py` script as
follows:

```ps1
Invoke-WebRequest -Uri https://bootstrap.pypa.io/get-pip.py -OutFile get-pip.py
python get-pip.py --user

$env:APPDATA\Python\Scripts\pip.exe install --user --upgrade google-cloud
$env:APPDATA\Python\Scripts\pip.exe install --user --upgrade google-api-python-client
$env:APPDATA\Python\Scripts\pip.exe install --user --upgrade pycrypto
```

The script rely on Gcloud being able to exercise the Application Default
Credentials ([more info here][appdefaultcredentials]).

[appdefaultcredentials]: https://developers.google.com/identity/protocols/application-default-credentials

First let's make sure that the `cloud` configuration is correct:

```ps1
gcloud config list
```

... That should display the correct account, project, region and zone for the
Chrome Enterprise Lab GCE project.

Now let's populate the Application Default Credentials. The following command
will start an interactive authentication flow which will populate the default
credentials using your `@google.com` credentials.

```ps1
gcloud auth application-default login
```

Now you should be ready to run the scripts.

