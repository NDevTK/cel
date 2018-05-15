# Contributing

## Quick Overview For The Restless

Same requirements apply as for the Chromium project. If you are not a Chromium
contributor, then please sign the [CLA](https://cla.developers.google.com/).

### Suggested Workflow

1. Create a new branch:

   ``` sh
   git new-branch your-awesome-feature
   ```

   This will create a new branch called `your-awesome-feature` configured to
   track the `master` branch off the CEL repository.

   **Note**: `git new-branch`, and `git cl` are tools included in `depot_tools`.

2. Make your changes. Commit as often as you need. All the commits gets squashed
   in the next step.

3. Check formatting:

   ``` sh
   python build.py format
   ```

   This will ensure that the code you are about to upload is correctly
   formatted. In addition, it runs a set of checks on Markdown source to ensure
   that links are correct. For more information, see `build.py format --help`.

4. Upload your changes:

   ``` sh
   git cl upload
   ```

   If successful, `git-cl` will create a new codereview entry. Send it out for
   review. Move `cel-reviews@chromium.org` from the CC line to the reviewers
   line if you don't know of a specific reviewer.

5. The review may result in changes being requested. Address review comments and
   make additional comments are necessary. The codereview you created in step 3 is
   associated with a Git branch. So you can run `git cl upload` to upload new
   patchsets as necessary.

6. Once the change is approved, someone will submit the change.

7. Good job!

   Optionally, you can clean up your local repository following a successful
   submission by running the following:

   ``` sh
   git rebase-update
   ```

   (This assumes that you are using the Git rebase workflow. See
   [here](https://chromium.googlesource.com/chromium/tools/depot_tools.git) for
   more details).

## Adding a new Go Dependency

This project uses [Dep][] to manage dependencies.

If you add a new import, then it's likely that `Dep` will complain. In that
case, just run the following:

``` sh
dep ensure
```

This will fetch the dependency in to `vendor` and update the version
constraints. You may need to check in `Gopkg.lock` and `Gopkg.toml` files if
they change.

[Dep]: https://golang.github.io/dep/

## Updating Go Dependencies

*** note
**Warning**: This is not safe unless you take special care to understand the new
changes you are pulling in. There might be version incompatibilities that you'll
need to resolve by hand.
***

When updates to dependent packages are available, run:

``` sh
dep ensure -update
```

... to fetch the updates. Don't forget to run the tests and verify that the new
code drops are safe and doesn't regress.

Add the `Gopkg.lock` and `Gopkg.toml` files and commit the change.

See [Go Dep](https://golang.github.io/dep/docs/introduction.html) documentation
for dealing with dependencies using Dep and also dealing with version mismatch
issues across dependencies.

*** note
**Note** : Please update deps as a separate commit. Don't mix depependecy
updates with code changes.
***

