# UNDERSTANDING THE LAB

The Chrome Enterprise Lab project aims to build a set of tools for autmating the
deployment of enterprise labs. Have a look through the following links for more
information on why we did this, and how the labs work.

## Contents

* [Background][]
  * [Use Cases][]
  * [Additional Considerations][]
  * [Frameworks/Tools Used][]

* [Design Summary][Design]
  * [Objective][]
  * [The Product][]
  * [Concepts][]
    * [The Isolate][ISOLATE]
    * [The Asset Manifest][ASSET MANIFEST]
    * [The Host Environment][HOST ENVIRONMENT]
    * [The System Test Runner][SYSTEM TEST RUNNER]
    * [The Deployer][DEPLOYER]
    * [The Test Host][TEST HOST]
    * [The `cel_bot` tool][cel_bot]
    * [The Host Test Runner][HOST TEST RUNNER]
    * [The Test][TEST]
    * [The `cel_py` library][cel_py]

* Design Details
  * [Deployment Details][]
  * [Google Services][]
  * [On-Premise Fixtures][]
  * [Scalability][]
  * [Bootstrapping][]
  * [Private Google Compute Images][]
  * [Source Locations][]
  * [Asset Description Schema][]
  * [Integration With Chromium Waterfall][]
  * [Workflows][]


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

