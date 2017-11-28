# Workflows

[TOC]

## Using the Chromium waterfall


### Trigger existing tests

This is basically what a non-Windows-using developer would use to verify that their code doesn't break some enterprise use case.



1.  Hack hack hack.
1.  `git cl upload`
1.  `git cl try -B luci.chrome.enterprise -b win`
1.  Wait for results to show up on Gerrit.


### Trigger existing tests using locally built binary

Expected to be used by someone who has built their binary on Windows and would like to quickly run a test case or two in the lab.



1.  Hack hack hack.
1.  Build binaries locally.
1.  `mb.py isolate out/win-foo system_tests`
1.  `isolate.py archive -i out/win-foo/system_tests.isolate -s out/win-foo/system_tests.isolated`
1.  `swarming.py …` or some other script to trigger builds.


## Using a private lab

TODO(asanka): Fill these out


### Creating a private lab


### Adding "permanent" fixtures


### Interactive development and debugging

Once they run into problems and printf is no longer their friend, it's perhaps time to login to a user VM and debug their binary. Shuttling of debug symbols is TBD.

This may also be used by a developer who hasn't written a system level test yet and would just like to manually verify that the fix works.


1.  Hack hack hack.
1.  Build binaries locally.
1.  `mb.py isolate out/win-foo enterprise_system_tests`
1.  `isolate.py archive -i out/win-foo/enterprise_system_tests.isolate -s out/win-foo/enterprise_system_tests.isolated`
1.  On VM: `isolateserver.py download -s <hash>`
1.  On VM: Run and debug binary manually.


<!-- INCLUDE index.md (51 lines) -->
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
[Background]: background.md
[Bootstrapping]: bootstrapping.md
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
[Host Example]: /examples/schema/ad/one-domain.host.textpb
[ISOLATE]: design-summary.md#isolate
[Integration With Chromium Waterfall]: chrome-ci-integration.md
[Objective]: design-summary.md#objective
[On-Premise Fixtures]: on-premise-fixtures.md
[Private Google Compute Images]: private-images.md
[SYSTEM TEST RUNNER]: design-summary.md#system-test-runner
[Scalability]: scalability.md
[Schema References]: schema-guidelines.md#references
[Schema Validation]: schema-guidelines.md#validation
[Inline References]: schema-guidelines.md#inline-references
[Source Locations]: source-locations.md
[TEST HOST]: design-summary.md#test-host
[TEST]: design-summary.md#test
[The Product]: design-summary.md#the-product
[Use Cases]: background.md#use-cases
[Workflows]: workflows.md
[cel_bot]: design-summary.md#cel_bot
[cel_py]: design-summary.md#cel_py

