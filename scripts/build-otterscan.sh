#!/usr/bin/env sh

# Build otterscan distribution and dump production build inside /ots module.
#
# Those files are meant to be commited in source control.
cd ../otterscan
npm ci && npm run build && rm -fr ../ots/dist && cp -R dist ../ots/dist

# Remove all .js|css files that have a corresponding precompressed .js|css.gz
cd ../ots/dist
find . -name "*.js.gz" | rev | cut -c 4- | rev | cat | xargs rm
find . -name "*.css.gz" | rev | cut -c 4- | rev | cat | xargs rm

# Remove all .br precompressed brotli files
rm assets/*.br
