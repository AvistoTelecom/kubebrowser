#!/usr/bin/env bash

rm -r server/kodata
(cd ui && pnpm run build)
cp -r ui/dist server/kodata
