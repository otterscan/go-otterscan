#!/usr/bin/env sh

# Copy chains metadata submodule to /ots module.
#
# Those files are meant to be commited in source control.
cd ../otterscan-assets
rm -fr ../ots/chains && cp -R chains/_data/chains ../ots/chains
