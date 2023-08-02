# Introduction

This is a data go module meant to repackage a production build of Otterscan as a go module to be embedded inside Erigon.

That allows Erigon to reference this module without having to add Otterscan git submodules into their repo.

> This repository incorporates code originally developed by external contributor Eddie Lee at https://gfx.cafe/open/4bytes.

# Update instructions

> Every time one of the submodules below is updated, this go module should be tagged and released as well, and a PR to the [Erigon GitHub repository](https://github.com/ledgerwatch/erigon) must then be issued to upgrade the embedded Otterscan version.

## Otterscan

For every Otterscan release, update the `otterscan` submodule to the proper git tag and regenerate the production build using the `scripts/build-otterscan.sh` script.

## Otterscan assets

Every time the external static assets distribution for Otterscan changes, update the `otterscan-assets` submodule to the proper git tag and regenerate the embeddable assets by:

1 - Regenerate the embedded chains metadata using the `scripts/build-chains.sh` script.
2 - Rebuild the `zstd` dictionary for 4bytes/topic0 signatures using the `scripts/build-dict.sh` script.
3 - Regenerate the compressed embeddable 4bytes/topic0 data using the `go run ./cmd/gen` command.
