subinclude("///pleasings2//go:compat")

go_library(
    name = "posthook",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["PUBLIC"],
    deps = [
        "//internal/ark",
        "//internal/ark/api",
        "//internal/ark/sync",
        "//src/auth",
    ],
)
