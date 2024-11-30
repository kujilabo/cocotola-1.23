#!/bin/bash

# echo cocotola-web
# pushd cocotola-web
# npm run lint
# popd
# code=$?
# if [ $code -ne 0 ]; then
#   exit 1
# fi

# echo cocotola-core
# pushd cocotola-core
# golangci-lint run --config ../.github/.golangci.yml && \
# golangci-lint run --disable-all --config ../.github/.golangci.yml \
# -E bodyclose \
# -E errorlint \
# -E exhaustive \
# -E forbidigo \
# -E forcetypeassert \
# -E gocognit \
# -E gocyclo \
# -E gofmt \
# -E goimports \
# -E gomnd \
# -E gosec \
# -E noctx \
# -E testpackage \
# -E thelper \
# -E unconvert \
# -E whitespace && \
# pkgforbid -config=../.github/pkgforbid.yml ./... && \
# popd
# code=$?
# if [ $code -ne 0 ]; then
#   exit 1
# fi

# echo cocotola-auth
# pushd cocotola-auth
# golangci-lint run --config ../.github/.golangci.yml && \
# golangci-lint run --disable-all --config ../.github/.golangci.yml \
# -E bodyclose \
# -E errorlint \
# -E exhaustive \
# -E forbidigo \
# -E forcetypeassert \
# -E gocognit \
# -E gocyclo \
# -E gofmt \
# -E goimports \
# -E gosec \
# -E noctx \
# -E testpackage \
# -E thelper \
# -E unconvert \
# -E whitespace && \
# pkgforbid -config=../.github/pkgforbid.yml ./... && \
# popd
# code=$?
# if [ $code -ne 0 ]; then
#   exit 1
# fi

task lint
if [ $code -ne 0 ]; then
  exit 1
fi

# echo cocotola-lib
# pushd lib
# golangci-lint run --config ../.github/.golangci.yml && \
# golangci-lint run --disable-all --config ../.github/.golangci.yml \
# -E bodyclose \
# -E errorlint \
# -E exhaustive \
# -E forbidigo \
# -E forcetypeassert \
# -E gocognit \
# -E gocyclo \
# -E gofmt \
# -E goimports \
# -E gosec \
# -E noctx \
# -E testpackage \
# -E thelper \
# -E unconvert \
# -E whitespace && \
# pkgforbid -config=../.github/pkgforbid.yml ./... && \
# popd
# code=$?
# if [ $code -ne 0 ]; then
#   exit 1
# fi
