#!/bin/bash


cd "$( dirname "${BASH_SOURCE[0]}" )"

# REGISTRY=registry.tce.com/infra

GOOS=linux
GOARCH="arm64"
BUILD_DIR=./bin


if [ -d $BUILD_DIR ]; then
  rm -rf $BUILD_DIR
fi
mkdir -p $BUILD_DIR


# build auditctl tools
go build -o $BUILD_DIR/auditctl


# Show all built files
ls -alh ${BUILD_DIR}
cp -raf $BUILD_DIR/auditctl  ./auditctl
