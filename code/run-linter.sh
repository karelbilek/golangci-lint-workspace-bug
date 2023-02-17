#!/usr/bin/env bash

set -ex

cd ../dumb_linter
go build -buildmode=plugin .
mv dumb_linter.so ../code
cd ../code
golangci-lint run ./...