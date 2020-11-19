# INTEGRATION WITH CHROMIUM WATERFALL

Initially we would like to introduce an FYI bot that can be scheduled
periodically to run the enterprise test suite. Once that's been kicked around a
bit, we can move onto a try bot and later consider adding a bot on the main
waterfall.

Based on
[this](https://chromium.googlesource.com/chromium/tools/build/+/master/scripts/slave/recipe_modules/chromium_tests/chromium_recipe.md)
document.


1.  Add a builder to the `master.chromium.fyi` master that is:
    1.  Called `Windows Enterprise`.
    1.  Uses MB configuration `release_bot`.
    1.  Compiles `chrome`.
    1.  Lists `system_instrumentation_tests` as a test. There maybe others, but this is the first planned system test and the one that's described in this document.
    1.  Enables swarming with dimension `pool: enterprise_lab`
1.  Add the [GREETER][] as a swarming bot to the `enterprise_lab` pool.
1.  ???
1.  Profit.

*** note
**Note**: We are splitting the builder and swarming tester so that it would be
possible for users to manually schedule swarming tests inside the lab using
their own binaries.
***

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

