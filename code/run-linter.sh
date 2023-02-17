#!/usr/bin/env bash

set -ex

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

cd ../dumb_linter
go build -buildmode=plugin .
mv dumb_linter.so ../code
cd ../code
golangci-lint run ./...