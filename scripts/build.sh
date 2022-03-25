#!/bin/bash

# go to project root
cd "$PWD/$(dirname $0)/.."

VERSION=$1
GIT_REV=$(git rev-parse --short HEAD)
BUILD_TIME=$(date +'%Y-%m-%dT%TZ%z')

pkg="github.com/keng42/xwallet/pkg/utilities/info"
vd="${VERSION}"

infoFlags="-X ${pkg}.Version=${VERSION} -X ${pkg}.GitRev=${GIT_REV} -X ${pkg}.BuildTime=${BUILD_TIME}"
alias gob='go build -ldflags "${infoFlags}" '
alias gobw='go build -ldflags "-s -w ${infoFlags}" '

if [ "$OSTYPE" = "linux-gnu" ]; then
  shopt -s expand_aliases
  source $HOME/.bashrc
fi

# disable CGO so it can run in docker's alpine image
export CGO_ENABLED=0
export GOARCH=amd64

mkdir -p ./build/ ./build/xwallet-darwin ./build/xwallet-windows

echo "building ${vd}"

# build for linux
echo "building for linux"

export GOOS=linux
gob -o ./build/xwallet-linux/xwallet ./cmd/xwallet

gobResult=$?
if [ "$gobResult" != "0" ]; then
  exit
fi

# build for macos
echo "building for macos"

export GOOS=darwin
gob -o ./build/xwallet-darwin/xwallet ./cmd/xwallet

gobResult=$?
if [ "$gobResult" != "0" ]; then
  exit
fi

# build for windows
echo "building for windows"

export GOOS=windows
gob -o ./build/xwallet-windows/xwallet.exe ./cmd/xwallet

gobResult=$?
if [ "$gobResult" != "0" ]; then
  exit
fi

# package

cp README.md CHANGELOG.md RELEASELOG.md LICENSE ./build/xwallet-linux/
cp README.md CHANGELOG.md RELEASELOG.md LICENSE ./build/xwallet-darwin/
cp README.md CHANGELOG.md RELEASELOG.md LICENSE ./build/xwallet-windows/

chmod +x ./build/xwallet-linux/xwallet
chmod +x ./build/xwallet-darwin/xwallet

cd ./build
tar zcvf xwallet-linux.tar.gz xwallet-linux
tar zcvf xwallet-darwin.tar.gz xwallet-darwin
tar zcvf xwallet-windows.tar.gz xwallet-windows
cd ..

echo "done"
