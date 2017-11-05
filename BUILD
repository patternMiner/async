load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "async",
    srcs = [
        "dispatcher.go",
        "event.go",
        "event_handler.go"
    ],
    deps = [],
    importpath = "github.com/patternMiner/async",
    visibility = ["//visibility:public"],
)

go_test(
    name = "event_test",
    srcs = ["event_test.go"],
    importpath = "github.com/patternMiner/async",
    library = ":async",
)
