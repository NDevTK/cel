# Contributing

## Quick Overview For The Restless

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

## Updating Dependencies

If you add a new import, then it's likely that `Dep` will complain. In that
case, just run the following:

``` sh
dep ensure
```

This will fetch the dependency in to `vendor` and pin the version.

When updates to dependent packages are available, run:

``` sh
dep update
```

... to fetch the updates. Don't forget to run the tests and verify that the new
code drops are safe and doesn't regress.

Add the `Gopkg.lock` and `Gopkg.toml` files and commit the change.

*** note
**Note** : Please update deps as a separate commit. Don't mix depependecy
updates with code changes.
***

