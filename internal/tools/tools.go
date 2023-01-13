//go:build tools
// +build tools

package tools

// Installed packages can be listed here and installed with makefile go install
// The go build directive above prevents this file from being built normally without tags, avoiding the error:
//   import "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway" is a program, not an importable package

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
