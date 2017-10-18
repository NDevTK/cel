# UNDERSTANDING THE LAB

The Chrome Enterprise Lab project aims to build a set of tools for autmating the
deployment of enterprise labs. Have a look through the following links for more
information on why we did this, and how the labs work.

## Contents

* [Background]
  * [Use Cases]
  * [Additional Considerations]
  * [Frameworks/Tools Used]

* [Design Summary][Design]
  * [Objective]
  * [The Product]
  * [Concepts]
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
  * [Deployment Details]
  * [Google Services]
  * [On-Premise Fixtures]
  * [Scalability]
  * [Bootstrapping]
  * [Private Google Compute Images]
  * [Source Locations]
  * [Asset Description Schema]
  * [Integration With Chromium Waterfall]
  * [Workflows]


<!-- INSERT-INDEX -->
<!-- BEGIN-INDEX -->
<!--
Index of tags used throughout the documentation. This list lives in
//docs/index.md and should be included in all documents that depend on these
tags. Whenever the list changes, run the following command:

   ./update-index.sh

This will replace any line containing the string 'INSERT-INDEX' with the
contents of this file. It'll also remove everything  between the BEGIN-INDEX,
END-INDEX block. So each time the script is run it'll replace the index with the
latest version.
-->

[ASSET MANIFEST]: design-summary.md#asset-manifest
[Additional Considerations]: background.md#additional-considerations
[Asset Description Schema]: schema-guidelines.md
[Background]: background.md
[Bootstrapping]: bootstrapping.md
[Concepts]: design-summary.md#concepts
[DEPLOYER]: design-summary.md#deployer
[Deployment Details]: deployment.md
[Deploying Scripted Assets]: deployment.md#deploying-scripted-assets
[Design]: design-summary.md
[Frameworks/Tools Used]: background.md#tools-used
[GREETER]: design-summary.md#greeter
[Google Services]: google-services.md
[HOST ENVIRONMENT]: design-summary.md#host-environment
[HOST TEST RUNNER]: design-summary.md#host-test-runner
[ISOLATE]: design-summary.md#isolate
[Integration With Chromium Waterfall]: chrome-ci-integration.md
[Objective]: design-summary.md#objective
[On-Premise Fixtures]: on-premise-fixtures.md
[Private Google Compute Images]: private-images.md
[SYSTEM TEST RUNNER]: design-summary.md#system-test-runner
[Scalability]: scalability.md
[Source Locations]: source-locations.md
[TEST HOST]: design-summary.md#test-host
[TEST]: design-summary.md#test
[The Product]: design-summary.md#the-product
[Use Cases]: background.md#use-cases
[Workflows]: workflows.md
[cel_bot]: design-summary.md#cel_bot
[cel_py]: design-summary.md#cel_py

<!-- END-INDEX -->
