#!/usr/bin/env bash

rm -r server/kodata
(cd ui && pnpm run build --base=/home/)
cp -r ui/dist server/kodata
