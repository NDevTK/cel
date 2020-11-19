# BOOTSTRAPPING

[TOC]

Bootstrapping is the process by which we go from an newly created empty Google
Cloud project to a functioning enterprise lab.

There are two types of bootstrapping:


1.   **Full**: The full enterprise lab with its physical counterparts and
     manually provisioned fixtures that is integrated into the waterfall and
     provides continuous integration testing. Following the sale of Boston
     Dynamics, we are no longer in a position to automate the deployment of
     physical lab components.  Physical labs and some virtual assets are
     expected to be not automated at least in the short term.

1.   **Standalone**: An independent lab sans some of the physical components and
     manually provisioned fixtures.

Given that one of the key use cases is the ability to *partially* clone the
environment for individual or team use, the bootstrapping process needs to be
streamlined and easily reproducible.


## Bootstrapping A New Instance Of The Lab

The process for establishing an instance of an enterprise lab hosted at a new
Google Cloud project follows:


1.  Start by building the Chrome Enterprise Lab infrastructure tools. See
    [Source Locations][] for `cel` repository location. The build process
    generates all the scripts and binaries that are needed for the remainder of
    this process.  The files are copied out to a `staging` directory. This
    directory will be needed by the `cel_admin bootstrap` command later.

1.  Create a Google Cloud project
    ([docs](https://cloud.google.com/resource-manager/docs/creating-managing-projects)).

1.  Create the assets that need to be manually created. These include physical
    networks, VPN gateways, physical machines, virtual machines that can't be
    automatically deployed due to licensing, reserving pools of externally
    visible IP addresses, etc.. The list of such assets obviously depend
    entirely on the list of tests to be run.

1.  Build a [HOST ENVIRONMENT][] configuration file.

1.  Verify the [HOST ENVIRONMENT][] optionally against a real [ASSET MANIFEST][]
    using the `cel_admin validate` command. This command will explain any
    missing pieces or report any errors it finds in the configuration. If an
    [ASSET MANIFEST][] is available, the `cel_admin validate` command should

1.  Invoke `cel_admin start` to construct the enterprise lab infrastructure.
    This command:

    1.  Verifies or enables the necessary Google Cloud APIs for the project.
    1.  Verifies or creates the Google Cloud Storage bucket needed for
    bootstrapping.
    1.  Copies the build artifacts from `staging` to the Google Cloud Storage
    bucket.
    1.  Copies the HOST ENVIRONMENT to a Google Cloud Storage bucket. The
    storage location is available to all VMs in the lab in project metadata
    `cel-host-environment`.
    1.  Purge all existing GCP resources that are of the types that are deployed
    by the lab. This should wipe out and reset the state for a Google Cloud
    project that had been previously used for a lab deployment. Purging also
    ensures that any resources that were orphaned during previous runs are
    cleaned up.
    1.  Starts deploying the intrinsic enterprise lab assets. This process is
	equivalent to what the DEPLOYER does, but instead the `cel_admin`
	command uses a list of builtin assets. The list of assets deployed here
	include:
	1.  Administrative network which hosts the GREETER and DEPLOYER
	1.  Service accounts for used for automated deployment
	1.  GREETER
	1.  DEPLOYER

1.  Start -- or restart -- the GREETER and DEPLOYER services so that they pick
    up the new configuration.

1.  At this point, the lab's GREETER can start accepting test runs and asset
    deployment requests. For Standalone use cases, this is sufficient.


## Updating An Existing Lab

Updates must currently be deployed manually and follow a similar workflow as
what's described above.


1.  Drain the lab using `cel_admin drain`. This notifies the lab's GREETER that
    it should no longer accept new requests to run tests or deploy assets. Any
    requests that are in progress will run to completion.

1.  Make any necessary changes to manually deployed assets.

1.  Update the HOST ENVIRONMENT to match.

1.  Invoke `cel_admin start` with the new HOST ENVIRONMENT. This command
    performs the same set of operations as described above for a new deployment.



## References

*   [tools/build](https://chromium.googlesource.com/chromium/tools/build/): Repository containing build configurations for Chromium waterfalls.
*   [Create a new builder](https://chromium.googlesource.com/chromium/tools/build/+/master/scripts/slave/recipe_modules/chromium_tests/chromium_recipe.md): Document on how to add a new builder to the waterfall.


<!-- INCLUDE index.md (56 lines) -->
<!--
Index of tags used throughout the documentation. This list lives in
/docs/index.md and is included in all documents that depend on these tags.

In order to update the tags:

   1. Update `/docs/index.md`
   2. Run the following command from the root of the source tree:

         ./build.py format

Keep the tags below sorted.
-->

[ASSET MANIFEST]: design-summary.md#asset-manifest
[Additional Considerations]: background.md#additional-considerations
[Asset Description Schema]: schema-guidelines.md
[Asset Example]: /examples/schema/ad/one-domain.asset.textpb
[Asset Schema]: /schema/asset/
[Background]: background.md
[Bootstrapping]: bootstrapping.md
[Coding Patterns for Resolvers]: deployment.md#coding-patterns-for-resolvers
[Completed Asset Manifest]: deployment.md#completed-asset-manifest
[Concepts]: design-summary.md#concepts
[DEPLOYER]: design-summary.md#deployer
[Deploying Scripted Assets]: deployment.md#deploying-scripted-assets
[Deployment Details]: deployment.md
[Deployment Overview]: deployment.md#overview
[Design]: design-summary.md
[Frameworks/Tools Used]: background.md#tools-used
[GREETER]: design-summary.md#greeter
[Google Services]: google-services.md
[HOST ENVIRONMENT]: design-summary.md#host-environment
[HOST TEST RUNNER]: design-summary.md#host-test-runner
[Host Environment Schema]: /schema/host/
[Host Example]: /examples/schema/ad/one-domain.host.textpb
[ISOLATE]: design-summary.md#isolate
[Inline References]: deployment.md#inline-references
[Integration With Chromium Waterfall]: chrome-ci-integration.md
[Key Management]: key-management.md
[Objective]: design-summary.md#objective
[On-Premise Fixtures]: on-premise-fixtures.md
[Private Google Compute Images]: private-images.md
[SYSTEM TEST RUNNER]: design-summary.md#system-test-runner
[Scalability]: scalability.md
[Schema References]: schema-guidelines.md#references
[Schema Validation]: schema-guidelines.md#validation
[Source Locations]: source-locations.md
[TEST HOST]: design-summary.md#test-host
[TEST]: design-summary.md#test
[The Product]: design-summary.md#the-product
[Use Cases]: background.md#use-cases
[Workflows]: workflows.md
[cel_bot]: design-summary.md#cel_bot
[cel_py]: design-summary.md#cel_py

