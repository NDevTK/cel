# Chrome Enterprise Lab

Chrome Enterprise Lab is an inaccurately named set of tools for building
enterprise labs quickly and easily. The labs so built can be used for system
level end-to-end testing of Google Chrome/Chromium.

Have a peek at the [Design](http://goto.google.com/chrome-enterprise-lab) document.

Also have a peek at the [Code of Conduct](./CODE_OF_CONDUCT.md).

Most of the code is in Go. See the [README](/src/go/README.md/) file.

## Building

Make sure you have [Go](https://golang.org/) and
[depot_tools](https://dev.chromium.org/developers/how-tos/install-depot-tools)
installed. If you are setup for developing Chromium, you only need to worry
about Go.

1. Clone this repository.

2. Make sure it builds.

   ``` sh
   python build.py build
   ```

3. Also make sure the tests pass.

   ``` sh
   python build.py test
   ```

## Contributing

Same requirements apply as for the Chromium project. If you are not a Chromium
contributor, then please sign the [CLA](https://cla.developers.google.com/).

1. Create a new branch:

   ``` sh
   git new-branch your-awesome-feature
   ```
   
   This will create a new branch called `your-awesome-feature` configured to
   track the `master` branch off the CEL repository.

2. Make your changes. Commit as often as you need. All the commits gets squashed
   in the next step.

3. Upload your changes:

   ``` sh
   git cl upload
   ```

   If successful, `git-cl` will create a new codereview entry. Send it out for
   review. Move `cel-reviews@chromium.org` from the CC line to the reviewers
   line if you don't know of a specific reviewer.

4. The review may result in changes being requested. Address review comments and
   make additional comments are necessary. The codereview you created in step 3 is
   associated with a Git branch. So you can run `git cl upload` to upload new
   patchsets as necessary.

5. Once the change is approved, someone will submit the change.

6. Good job!

   Optionally, you can clean up your local repository following a successful
   submission by running the following:

   ``` sh
   git rebase-update
   ```

   (This assumes that you are using the Git rebase workflow. See
   [here](https://chromium.googlesource.com/chromium/tools/depot_tools.git) for
   more details).
