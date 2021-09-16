#!/usr/bin/env lucicfg

# Constants shared by multiple definitions below.
REPO_URL = "https://chromium.googlesource.com/enterprise/cel"
RECIPE_BUNDLE = "infra/recipe_bundles/chromium.googlesource.com/chromium/tools/build"

lucicfg.check_version("1.27.0", "Please update depot_tools")

# Enable LUCI Realms support and launch 100% of Swarming tasks for builds in
# "realms-aware mode".
lucicfg.enable_experiment("crbug.com/1085650")
luci.builder.defaults.experiments.set({"luci.use_realms": 100})

# Enable v2 bucket names in LUCI Scheduler config.
lucicfg.enable_experiment("crbug.com/1182002")

# Tell lucicfg what files it is allowed to touch.
lucicfg.config(
    config_dir = "generated",
    fail_on_warnings = True,
    lint_checks = ["default"],
)

# Definition of what LUCI micro-services to use and global ACLs that apply to
# all buckets.
luci.project(
    name = "celab",
    buildbucket = "cr-buildbucket.appspot.com",
    logdog = "luci-logdog.appspot.com",
    milo = "luci-milo.appspot.com",
    notify = "luci-notify.appspot.com",
    scheduler = "luci-scheduler.appspot.com",
    swarming = "chromium-swarm.appspot.com",
    acls = [
        # This project is publicly readable.
        acl.entry(
            roles = [
                acl.BUILDBUCKET_READER,
                acl.LOGDOG_READER,
                acl.PROJECT_CONFIGS_READER,
                acl.SCHEDULER_READER,
            ],
            groups = "all",
        ),
        # Allow committers to use CQ and to force-trigger and stop CI builds.
        acl.entry(
            roles = acl.CQ_COMMITTER,
            groups = "project-celab-committers",
        ),
        # Ability to schedule a build
        acl.entry(
            roles = acl.SCHEDULER_OWNER,
            groups = "project-celab-admins",
        ),
        # Ability to launch CQ dry runs.
        acl.entry(
            roles = acl.CQ_DRY_RUNNER,
            groups = "project-celab-tryjob-access",
        ),
        # Group with robots that have write access to the Logdog prefix.
        acl.entry(
            roles = acl.LOGDOG_WRITER,
            groups = "luci-logdog-chromium-writers",
        ),
    ],
)

# Required Logdog configuration.
luci.logdog(gs_bucket = "chromium-luci-logdog")

luci.milo(
    logo = "https://storage.googleapis.com/chrome-infra-public/logo/celab.png",
)

# list_view definition for milo
luci.list_view(name = "Daily Test Builders")
luci.list_view(name = "Rotation Console")
luci.list_view(name = "Try Builders")

# The Milo console with all post-submit builders, referenced below.
luci.console_view(
    name = "Main Console",
    repo = REPO_URL,
)
luci.cq(
    status_host = "chromium-cq-status.appspot.com",
    # Optional arguments.
    submit_max_burst = 4,
    submit_burst_delay = 480 * time.second,
)

luci.notify(tree_closing_enabled = True)
luci.notifier_template(
    name = "celab_ci",
    # To test template changes:
    #   1) Branch `luci-notify/email-templates/celab_ci.template` to `celab_debug.template`.
    #   2) Submit `celab_debug.template` changes to master.
    #   3) Update the [celab config](https://luci-config.appspot.com/#/projects/celab).
    #   4) Run: `git cl try -B "luci.celab.try" -b "linux" -p email_notify="[{\"email\": \"<you>@google.com\", \"template\": \"celab_debug\"}]"`
    #   5) If it looks good, move changes to `celab_ci.template` and delete `celab_debug.template`.
    body = io.read_file("./luci-notify/email-templates/celab_ci.template"),
)

# Bucket with post-submit builders.
luci.bucket(
    name = "ci",
    acls = [
        acl.entry(
            roles = acl.BUILDBUCKET_READER,
            groups = "all",
        ),
    ],
)

# Bucket with pre-submit builders.
luci.bucket(
    name = "try",
    acls = [
        # Allow launching tryjobs directly (in addition to doing it through CQ).
        acl.entry(
            roles = acl.BUILDBUCKET_TRIGGERER,
            groups = "project-celab-tryjob-access",
        ),
    ],
)

# The CQ group with all pre-submit builders, referenced below.
luci.cq_group(
    name = "Main",
    watch = cq.refset(repo = REPO_URL, refs = ["refs/heads/.+"]),
    retry_config = cq.retry_config(
        # Optional arguments.
        single_quota = 1,
        global_quota = 2,
        failure_weight = 1,
        transient_failure_weight = 1,
        timeout_weight = 2,
    ),
)

# The gitiles poller: a source of commits that trigger CI builders.
luci.gitiles_poller(
    name = "master-gitiles-trigger",
    bucket = "ci",
    repo = REPO_URL,
)

def ci_builder(name, *, short_name, os, category, cpu = "x86-64", schedule = None, test_coverage = None):
    """Defines a post-submit builder.
    Args:
      name: name of the builder to define.
      short_name: short name for console view in milo.
      os: the target OS.
      category: the category to put it under in the console.
      cpu: the target CPU.
      schedule: cron definition, if it is None then gitiles_poller will be a trigger point.
      test_coverage: possible values are full, quick, none.
                "full"  : execute all test cases.
                "quick" : execute few quick tests that cover key parts of our test framework.
                none    : no test cases will be executed.
    """
    triggered_by = ["master-gitiles-trigger"]

    #if schedule is none then gitiles_poller will be a trigger point.
    if schedule != None:
        triggered_by = None

    properties = {
        "pool_name": "celab-ci",
        "pool_size": 5,
        "target_cpu": "x64",
    }

    # Property to execute all test cases.
    if test_coverage != None:
        properties.update({"tests": "*"})

    # Property for quick tests that cover key parts of our test framework.
    # These are quick tests that exercise the key parts of our test infra.
    if test_coverage == "quick":
        properties.update({"include": "core"})

    luci.builder(
        name = name,
        bucket = "ci",
        executable = luci.recipe(
            name = "celab",
            recipe = "celab",
            cipd_package = RECIPE_BUNDLE,
            cipd_version = "refs/heads/main",
        ),
        properties = properties,
        dimensions = {
            "pool": "luci.flex.ci",
            "cpu": cpu,
            "os": os,
        },
        service_account = "celab-ci-builder@chops-service-accounts.iam.gserviceaccount.com",
        # Run this builder on commits to REPO_URL.
        triggered_by = triggered_by,
        schedule = schedule,
    )

    name_with_bucket = "ci/" + name  # disambiguate by prefixing the bucket name

    # Non-scheduler streams displayed in main-console.
    if schedule == None:
        # Add it to the console as well.
        luci.console_view_entry(
            builder = name_with_bucket,
            console_view = "Main Console",
            category = name,
        )
    else:
        # Add it to the separate console for every scheduler.
        luci.list_view_entry(
            builder = name_with_bucket,
            list_view = "Daily Test Builders",
        )

    # Add it to the rotation console as well.
    luci.list_view_entry(
        builder = name_with_bucket,
        list_view = "Rotation Console",
    )

    # Defines notifiers on luci-notify.appspot.com.
    luci.notifier(
        name = "celab-check",
        on_status_change = ["SUCCESS", "FAILURE"],
        on_new_status = ["SUCCESS", "FAILURE"],
        notify_emails = ["chrome-enterprise-lab+luci@google.com"],
        template = "celab_ci",
        notified_by = [name_with_bucket],
    )

# definition of CI builders.
ci_builder("Daily", short_name = "daily", os = "Ubuntu-16.04", category = "Linux|64", schedule = "0 10 * * * ", test_coverage = "full")
ci_builder("Linux", short_name = "linux", os = "Ubuntu-16.04", category = "Linux|16.04", test_coverage = "full")
ci_builder("Windows", short_name = "win", os = "Windows-10", category = "Win|64", test_coverage = "quick")

def try_builder(name, *, short_name, category, os, cpu = "x86-64", test_coverage = None):
    """Defines a pre-submit builder.
    Args:
      name: name of the builder to define.
      short_name: short name for console view in milo.
      category: the category to put it under in the console.
      os: the target OS.
      cpu: the target CPU.
            test_coverage: possible values are full, quick, none.
                "full"  : execute all test cases.
                "quick" : execute few quick tests that cover key parts of our test framework.
                none    : no test cases will be executed.
    """
    properties = {
        "pool_name": "celab-try",
        "pool_size": 5,
        "target_cpu": "x64",
    }

    # Property to execute all test cases.
    if test_coverage != None:
        properties.update({"tests": "*"})

    # Property for quick tests that cover key parts of our test framework.
    # These are quick tests that exercise the key parts of our test infra.
    if test_coverage == "quick":
        properties.update({"include": "core"})

    luci.builder(
        name = name,
        bucket = "try",
        executable = luci.recipe(
            name = "try_builder",
            recipe = "celab",
            cipd_package = RECIPE_BUNDLE,
        ),
        properties = properties,
        dimensions = {
            "pool": "luci.flex.try",
            "os": os,
            "cpu": cpu,
        },
        service_account = "celab-try-builder@chops-service-accounts.iam.gserviceaccount.com",
    )

    name_with_bucket = "try/" + name  # disambiguate by prefixing the bucket name

    # Add to the CQ.
    luci.cq_tryjob_verifier(
        builder = name_with_bucket,
        cq_group = "Main",
    )

    # Add it to the console as well.
    luci.list_view_entry(
        builder = name_with_bucket,  # disambiguate by prefixing the bucket name,
        list_view = "Try Builders",
    )

# definition of try builders.
try_builder("linux-build", short_name = "build", category = "linux", os = "Ubuntu-16.04")
try_builder("linux-quick-tests", short_name = "quick", category = "linux|tests", os = "Ubuntu-16.04", test_coverage = "quick")
try_builder("linux-full-tests", short_name = "full", category = "linux|tests", os = "Ubuntu-16.04", test_coverage = "full")
try_builder("windows-build", short_name = "build", category = "windows", os = "Windows-10")
try_builder("windows-quick-tests", short_name = "quick", category = "windows|tests", os = "Windows-10", test_coverage = "quick")
try_builder("windows-full-tests", short_name = "full", category = "windows|tests", os = "Windows-10", test_coverage = "full")
