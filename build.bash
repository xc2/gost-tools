#!/usr/bin/env bash
set -e

BUILDDIR=build
ENTRY="$PWD"

build() {
  local GOOS="$1"
  local GOARCH="$2"
  local UPX="$3"
  T="gost-tools_${GOOS}_${GOARCH}"
  env CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$T" -ldflags '-s -w' "$ENTRY"
  if [[ -n "$UPX" ]]; then
    upx -1 --lzma "$T"
  fi
}

mkdir -p "$BUILDDIR"
pushd "$BUILDDIR"
rm -f *

build linux amd64 1

ls -allh

sha256sum gost-tools_* | tee SHA256SUM
