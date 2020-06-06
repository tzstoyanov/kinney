// Package tools is a "fake" package that is used merely to declare dependencies
// on tools.  This file will not compile, as it imports "main" packages, so the
// "tools" build constraint is used to guarantee that the file will not be
// included in any normal build.
//
// See https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
// for more information.

// +build tools

package tools

import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
