#!/usr/bin/env sh

# Build otterscan distribution and dump production build inside /ots module.
#
# Those files are meant to be commited in source control.
cd ../otterscan
npm ci && npm run build && rm -fr ../ots/dist && cp -R dist ../ots/dist
