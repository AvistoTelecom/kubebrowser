#!/usr/bin/env bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

rm -r server/kodata
pnpm build:ui --base=/home/
cp -r ui/dist server/kodata
