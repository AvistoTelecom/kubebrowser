#!/usr/bin/env bash

set -e

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

# Get git tag or commit sha
GIT_TAG=$(git describe --tags --exact-match 2>/dev/null || true)
if [ -n "$GIT_TAG" ]; then
  export VERSION="$GIT_TAG"
else
  export VERSION=$(git rev-parse --short HEAD)
fi

rm -rf server/kodata
VITE_VERSION=$VERSION pnpm build:ui --base=/home/
cp -r ui/dist server/kodata
