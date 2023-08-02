#!/usr/bin/env sh

# zstd train 4bytes/topic
cd ../otterscan-assets
zstd -T0 --train -o ../sigs/data/zdict --maxdict=32000 -r "./4bytes/signatures"
zstd -T0 --train -o ../topics/data/zdict --maxdict=32000 -r "./topic0/signatures"
