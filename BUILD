load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/iangudger/zinc_demo
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_binary(
    name = "zinc_demo",
    embed = [":zinc_demo_lib"],
    #gotags = ["ne"],
    importpath = "github.com/iangudger/zinc_demo",
    visibility = ["//visibility:public"],
)

go_library(
    name = "zinc_demo_lib",
    srcs = ["main.go"],
    importpath = "github.com/iangudger/zinc_demo",
    tags = ["manual"],
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_zinclabs_zinc//pkg/handlers/auth",
        "@com_github_zinclabs_zinc//pkg/handlers/document",
        "@com_github_zinclabs_zinc//pkg/handlers/index",
        "@com_github_zinclabs_zinc//pkg/handlers/search",
        "@com_github_zinclabs_zinc//pkg/meta/elastic",
    ],
)
