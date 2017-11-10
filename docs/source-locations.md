## Source Code Locations


Proposed locations for source code is as follows:


*   `//chrome/test/enterprise/system_tests`: Is the "home" of system tests in
    the Chromium code base. It contains:
    *   `doc/`: Documentation for how to run the test suite, how to add tests,
        how to add new test suites other than `system_tests`, a Markdown version
        of this document distilled for use by future generations.
    *   `assets/`: Test assets used during the deployment phase.
	*   `assets.textpb`: Text format protobuf describing all the test
	    infrastructure needed by the Chromium System Tests suite. This file
	    may be split into multiple files for easier maintenance.
	*   Everything else: Additional files can be referenced by
	    `assets.textpb` for the purpose of describing assets. These include
	    files used to populate `wwwroot` on Microsoft IIS servers,
	    configuration files for proxies and other services, etc.

    *   `scripts/`: Test runner scripts and their dependencies go here. These
        don't define tests, but rather just the manner in which the tests run.
	*   `system_test_runner.py`: [SYSTEM TEST RUNNER] script. Will complain
	    loudly if it is run from anywhere else other than a [GREETER].
	*   `host_	test_runner.py`: [HOST TEST RUNNER] script. Will complain
	    loudly if it is run from anywhere else other than a [TEST HOST].
    *   `selenium_tests/`: Selenium based tests. Any file that ends in
        `*_test.py` will get picked up by the `host_test_runner.py` as a
        candidate test and executed. The entire subdirectory is made available
        as-is on every [TEST HOST] in addition to any binaries under test. As
        such, tests can rely on additional files in this directory during test
        execution.

    *   **Other**: It is possible we'd need to run C++ GoogleTest test binaries
        as well. These would need to be compiled and isolated separately from
        the Selenium based tests since an isolate ideally only contains a single
        test invocation. If any such tests are written, we could place their
        source here.

*   `chromium.googlesource.com/enterprise/cel` ↔ `github.com/chromium/cel` :
    Code for bootstrapping, administration, [DEPLOYER] service, startup code for
    [TEST HOST], and [cel_py].

It should be possible for anyone to use the code in `cel` and build their own
lab on a Google Cloud project. There's nothing secret about how the lab works
and the `cel` project does not include any hosting configuration. The
`site.textpb` host configuration file needs to be injected during the
[Bootstrapping] phase.

Some parts of the infrastructure *maybe* proprietary and are discussed in the
CAVEATS section under [Private Google Compute Images], and [On-Premise Fixtures]
sections.

*** note
**Note** Files that are needed on the [TEST HOST] should not go in the
`//chrome/test/enterprise/system_tests/assets` directory. The files in `assets/`
are used only during the DEPLOY phase.
***


<!-- BEGIN-INDEX -->
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
