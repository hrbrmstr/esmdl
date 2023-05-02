#!/bin/bash

BIN="esmdl"
MACOS_TARGET="macosx13.3"

CLANGWRAP="$(PWD)/clangwrap.sh"
MACOS_CLANGWRAP="${CLANGWRAP}"

export CGO_CFLAGS_ALLOW="-fembed-bitcode"
export CGO_LDFLAGS_ALLOW="-bitcode_bundle -headerpad_max_install_names"
export LDFLAGS="-linkmode=external"

CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=arm64 \
SDK="${MACOS_TARGET}" \
CC="${MACOS_CLANGWRAP}" \
CGO_CFLAGS="-fembed-bitcode" \
go build -ldflags=-linkmode=external -tags macos -o "${BIN}-macos-arm" .

CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=amd64 \
SDK="${MACOS_TARGET}" \
CC="${MACOS_CLANGWRAP}" \
go build -ldflags=-linkmode=external -tags macos -o "${BIN}-macos-x86" .

lipo "${BIN}-macos-x86" "${BIN}-macos-arm" -create -output "${BIN}-macos"

codesign -s "${APPLE_SIGN}" --options runtime --entitlements hardened.plist "${BIN}-macos"
