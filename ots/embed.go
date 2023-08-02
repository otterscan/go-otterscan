package ots

import "embed"

//go:embed dist/*
var Dist embed.FS

//go:embed chains/*
var Chains embed.FS
