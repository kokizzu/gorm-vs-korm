#!/usr/bin/env bash
set -euo pipefail

go mod edit -droprequire=github.com/sourcegraph/sourcegraph/lib
GOFLAGS=-mod=mod go get github.com/sourcegraph/conc@latest
GOFLAGS=-mod=mod go get github.com/kokizzu/gotro@latest
GOFLAGS=-mod=mod go get golang.org/x/crypto@v0.53.0
GOFLAGS=-mod=mod go get github.com/opencontainers/runc@v1.3.6
go mod edit -exclude=github.com/docker/docker@v20.10.22+incompatible
go mod edit -exclude=github.com/docker/docker@v27.1.1+incompatible
GOFLAGS=-mod=mod go mod tidy
go mod edit -require=golang.org/x/crypto@v0.53.0
GOFLAGS=-mod=mod go mod download golang.org/x/crypto
