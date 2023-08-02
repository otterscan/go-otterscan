package main

import (
	"bytes"
	"encoding/hex"
	"io/fs"
	"os"
	"path"
	"sort"

	"github.com/klauspost/compress/zstd"
	"github.com/openacid/slim/encode"
	"github.com/openacid/slim/trie"
	"github.com/spf13/afero"
)

// generate dict with /scripts/build-dict.sh
type pair struct {
	k string
	v []byte
}

type mapping struct {
	toSig []pair
}

func (m *mapping) add(sig []byte, hash string) {
	m.toSig = append(m.toSig, pair{k: hash, v: sig})
}

func main() {
	doGenerate("topics", "no_names", "./otterscan-assets/topic0/signatures")
	doGenerate("sigs", "no_names", "./otterscan-assets/4bytes/signatures")

	doGenerate("topics", "with_names", "./otterscan-assets/topic0/with_parameter_names")
	doGenerate("sigs", "with_names", "./otterscan-assets/4bytes/with_parameter_names")
}

func doGenerate(tp string, kind string, path string) {
	m := &mapping{
		toSig: make([]pair, 0, 500000),
	}
	sigs := afero.NewBasePathFs(afero.NewOsFs(), path)
	err := afero.Walk(sigs, ".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			bts, err := afero.ReadFile(sigs, path)
			if err != nil {
				return err
			}
			hb, err := hex.DecodeString(info.Name())
			if err != nil {
				return err
			}
			m.add(bts, string(hb))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	err = m.create(tp, kind)
	if err != nil {
		panic(err)
	}
}

func NewPtr(pos, sz uint32) uint64 {
	return uint64(pos)<<32 | uint64(sz)
}

func (m *mapping) create(tp string, kind string) error {
	sort.SliceStable(m.toSig, func(i, j int) bool {
		return m.toSig[i].k < m.toSig[j].k
	})
	k := make([]string, 0)
	v := make([][]byte, 0)
	cx := make([]uint64, 0)
	cur := 0
	dict, _ := os.ReadFile(path.Join("./", tp, "/data/zdict"))
	buf := new(bytes.Buffer)
	cmp, _ := zstd.NewWriter(nil, zstd.WithEncoderDict(dict))
	for _, vv := range m.toSig {
		k = append(k, vv.k)
		v = append(v, vv.v)
		enc := cmp.EncodeAll(vv.v, nil)
		cx = append(cx, NewPtr(
			uint32(cur),
			uint32(len(enc)),
		))
		buf.Write(enc)
		cur = cur + len(enc)
	}
	//tr := boomphf.NewHMap[string, uint64](k, cx, boomphf.String)
	tr, err := trie.NewSlimTrie(encode.U64{}, k, cx, trie.Opt{
		Complete: trie.Bool(true),
	})
	if err != nil {
		return err
	}
	bts, err := tr.Marshal()
	if err != nil {
		return err
	}
	sb := new(bytes.Buffer)
	for _, vv := range m.toSig {
		sb.Write(vv.v)
	}
	err = os.WriteFile(path.Join("./", tp, "data", kind), buf.Bytes(), 0o644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join("./", tp, "data", "trie_"+kind), bts, 0o644)
	if err != nil {
		return err
	}
	return nil
}
