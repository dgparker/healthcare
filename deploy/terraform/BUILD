package(default_visibility = ["//visibility:public"])

licenses(["notice"])  # Apache 2.0

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "apply.go",
        "config.go",
    ],
    importpath = "github.com/GoogleCloudPlatform/healthcare/deploy/terraform",
    deps = [
        "//deploy/runner:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "apply_test.go",
    ],
    embed = [":go_default_library"],
    # Override default run dir to make it easier to find test files.
    rundir = ".",
    deps = [
        "//deploy/runner:go_default_library",
        "@com_github_google_cmp//cmp:go_default_library",
    ],
)
