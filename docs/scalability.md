## Scalability

Enterprise system tests are resource intensive. So much so that a single test
suite takes over an entire GCP project while running. Assuming we manage to
build our entire test suite such that the properties listed under the [TEST][]
section are met, then it would be possible for multiple test suites to run at
the same time in a single GCP project.

In practice, this isn't perfect since there's too much room for two tests to
interact with each other. For example, a change made to an IIS Web Application
can affect a test suite from a prior revision that's running concurrently.
Fortunately, such changes should be infrequent and the lab should be able to
parallelize multiple test session as long as their [ASSET MANIFEST][]s are the same.

When a change is detected to the [ASSET MANIFEST][], the [GREETER][] can request all
other test suites to be drained so that the ambient asset definitions can be
rotated to the new version. Once this is done, all new test runs based on the
new asset definition can once again be parallelized.

This process is a bit complicated and is something we can implement as we move
further along in the implementation. It's not something that's being considered
for the initial deployment.

A more reliable (less flaky or complicated) means of scaling would be to allow
multiple deployments of the lab, each into its own GCP project. A [GREETER][] from
each instance could be registered with Luci Isolate effectively providing load
balancing for enterprise tests.

The number of such labs that can be integrated with the waterfall is limited by
the availability of physical resources as described in [On-Premise Fixtures][].
Each instance of the lab must have a set of identically configured physical
devices for tests to be reliable.

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

