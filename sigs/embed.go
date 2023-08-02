package sigs

import (
	_ "embed"

	"github.com/otterscan/go-otterscan/triemap"
)

//go:embed data/trie_with_names
var hashTrieNamed []byte

//go:embed data/with_names
var sigDataNamed []byte

//go:embed data/trie_no_names
var hashTrie []byte

//go:embed data/no_names
var sigData []byte

//go:embed data/zdict
var zdict []byte

var Both = triemap.NewCompoundContainer(Named, Abi)
var Named = triemap.NewContainer(sigDataNamed, hashTrieNamed, zdict)
var Abi = triemap.NewContainer(sigData, hashTrie, zdict)

// gets the method signature given a 8 character hex string
// returns empty string if not found
var Hex = Both.Hex

// looks up the method signature given 4 bytes, cast to a string
// returns empty string if not found
var Lookup = Both.Lookup

// gets the method signature given a 8 character hex string
// returns empty string if not found
var HexNamed = Named.Hex

// looks up the method signature given 4 bytes, cast to a string
// returns empty string if not found
var LookupNamed = Named.Lookup

// gets the method signature given a 8 character hex string
// returns empty string if not found
var HexAbi = Abi.Hex

// looks up the method signature given 4 bytes, cast to a string
// returns empty string if not found
var LookupAbi = Abi.Lookup
