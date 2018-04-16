# Private Google Compute Images

Google Compute Engine doesn't support all the versions of the operating systems
that we would like to support in the enterprise lab. For example, the [GCE
public images](https://cloud.google.com/compute/docs/images#os-compute-support)
only include server editions of Microsoft Windows. While some amount of Windows
client functionality can be tested on the server editions, we would prefer to
test using the desktop editions of Windows, including versions of Windows that
are not congruent to the server editions available in the GCE public images.

The private images we need to support are:

*   Windows 7
*   Windows 10


## Substitutions

In the short run, it is possible to approximate desktop editions of Windows
using their server equivalents as described below. However, if any notable
incompatibilities are found, then we'd need to expedite the deployment of
private desktop images.

Currently we are not aware of any differences between the client OSs and their
corresponding substitutions that would directly affect the behavior of Chrome.

*** note
**Substitutions aren't always a solution.**

Even when a server edition can be configured to be functionally equivalent to a
client edition of Windows, some applications can and are known to discriminate
between client and server editions. An example is AV products, which demand
different licensing schemes and different product lines be used on client and
server editions. Testing with AV enabled would be one case where supporting
client images would be a blocking requirement.
***


### Windows 8.1 ≈ Windows Server 2012 R2

Windows 8.1 is functionally equivalent to Windows Server 2012 R2 with the
Desktop Experience feature installed. The [Windows Server 2012 R2 Desktop
Experience
Overview](https://technet.microsoft.com/en-us/library/dn609826(v=ws.11).aspx) in
Microsoft Technet states that the installed features are from Windows 8.1 and
includes the Windows Store. The server edition contains additional server
oriented features that are not expected to interfere with desktop functionality.

Windows Server 2012 R2 is currently a [public
image](https://cloud.google.com/compute/docs/images#os-compute-support)
supported on GCE.


### Windows 7 ≈ Windows Server 2008 R2

Much like Windows 8.1 and Windows Server 2012 R2, Windows 7 is functionally
equivalent to Windows Server 2008 R2. The [Windows Server 2008 R2 Desktop
Experience](https://technet.microsoft.com/en-us/library/cc772567(v=ws.11).aspx)
description in TechNet indicates that the addition of the Desktop Experience
feature will add Windows Defender, AVI support and Windows Media Player in
addition to a number of desktop mode applications. While the additional desktop
applications shouldn't be a compatibility factor, Windows Defender might be.
Hence we should include Desktop Experience when substituting Windows Server 2008
R2 for Windows 7.

Windows Server 2008 R2 is currently a [public
image](https://cloud.google.com/compute/docs/images#os-compute-support)
supported on GCE.


### Windows 10 ≈ Windows Server 2016

GCE currently includes two Windows Server 2016 flavors: Server 2016 Core and
Server 2016. Of these, the Server 2016 (not Core) image is the spiritual 2016
equivalent of Windows Server 2012 R2 with Desktop Experience. Server 2016
includes the Windows 10 shell among other desktop features by default according
to Microsoft's
[documents](https://docs.microsoft.com/en-us/windows-server/get-started/getting-started-with-server-with-desktop-experience).

Windows Server 2016 is currently a [public
image](https://cloud.google.com/compute/docs/images#os-compute-support)
supported on GCE.

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

