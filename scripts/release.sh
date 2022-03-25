#!/bin/bash

# go to project root
cd "$PWD/$(dirname $0)/.."

VERSION=$1

if [ -z $VERSION ]; then
  echo "empty version"
  exit
fi

git tag $1 && \
  git-chglog -o CHANGELOG.md && \
  git-chglog --template=".chglog/RELEASE.tpl.md" -o RELEASELOG.md ${VERSION} && \
  git tag -d $1 && \
  git add CHANGELOG.md RELEASELOG.md && \
  git commit -m "release ${VERSION}" && \
  git tag $1

