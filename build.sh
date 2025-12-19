#!/usr/bin/env bash
set -euo pipefail

DARWIN_AMD64="./dist/lista-darwin-amd64"
DARWIN_ARM64="./dist/lista-darwin-arm64"
LINUX_AMD64="./dist/lista-linux-amd64"
LINUX_ARM64="./dist/lista-linux-arm64"

buildsArray=(
  "$DARWIN_AMD64"
  "$DARWIN_ARM64"
  "$LINUX_AMD64"
  "$LINUX_ARM64"
)

BUILD_DIR="./dist"

if [ ! -d "$BUILD_DIR" ]; then
  mkdir "$BUILD_DIR"
fi

for build in "${buildsArray[@]}"; do
    if test -f "$build" 
    then
        echo "$build exists."
    else
        # GOOS=darwin GOARCH=amd64 go build -o "./$BUILD_DIR/$DARWIN_AMD64"
        # GOOS=darwin GOARCH=arm64 go build -o "./$BUILD_DIR/$DARWIN_ARM64"
        # GOOS=linux GOARCH=amd64 go build -o "./$BUILD_DIR/$LINUX_AMD64"
        # GOOS=linux GOARCH=arm64 go build -o "./$BUILD_DIR/$LINUX_ARM64"
        echo "$build does not exist"
    fi
done

