package triemap

import (
	_ "embed"
	"sync"

	"github.com/klauspost/compress/zstd"
	"github.com/openacid/slim/encode"
	"github.com/openacid/slim/trie"
)

type Container interface {
	// gets the method signature given a hex string
	// returns empty string if not found
	Hex(string) string
	// Gets the method signature given raw bytes, cast to a string
	// returns empty string if not found
	Lookup(string) string
}

type Ptr uint64

func NewPtr(pos, sz uint32) Ptr {
	return Ptr(uint64(pos)<<32 | uint64(sz))
}

func (p Ptr) Pos() uint32 {
	return uint32(p >> 32)
}

func (p Ptr) Sz() uint32 {
	return uint32(0>>32 | p)
}

var PtrCodec, _ = encode.NewTypeEncoder(new(Ptr))

type container struct {
	trie []byte
	data []byte
	dict []byte

	p sync.Pool
	//idx *boomphf.HMap[string, uint64]
	//idx *trie.SlimTrie
	idx Getter
}

type Getter interface {
	Get(string) (uint64, bool)
}

func NewContainer(data []byte, trb []byte, dict []byte) Container {
	return &container{
		trie: trb,
		data: data,
		dict: dict,
		p: sync.Pool{
			New: func() any {
				a, _ := zstd.NewReader(nil, zstd.WithDecoderDicts(dict))
				return a
			},
		},
		//idx: func() *boomphf.HMap[string, uint64] {
		//	tr := boomphf.NewHMap[string, uint64](nil, nil, boomphf.String)
		//	err := tr.Unmarshal(trb)
		//	if err != nil {
		//		panic(err)
		//	}
		//	return tr
		//}(),
		idx: func() *SlimTrieBackend {
			tr, err := trie.NewSlimTrie(PtrCodec, nil, nil)
			if err != nil {
				panic(err)
			}
			err = tr.Unmarshal(trb)
			if err != nil {
				panic(err)
			}
			return &SlimTrieBackend{tr}
		}(),
	}
}

func (c *container) Decode(in []byte) string {
	dd := c.p.Get().(*zstd.Decoder)
	defer c.p.Put(dd)
	ans, _ := dd.DecodeAll(in, make([]byte, 0, len(in)*2))
	return string(ans)
}

func (c *container) LookupPtr(raw string) *Ptr {
	val, ok := c.idx.Get(raw)
	if !ok {
		return nil
	}
	p := Ptr(val)
	return &p
}

func (c *container) Hex(hex string) string {
	return c.Lookup(RawFromHex(hex))
}

func (c *container) Lookup(raw string) string {
	ptr := c.LookupPtr(raw)
	if ptr == nil {
		return ""
	}
	if int(ptr.Pos()+ptr.Sz()) > len(c.data) {
		return ""
	}
	return c.Decode(c.data[ptr.Pos() : ptr.Pos()+ptr.Sz()-1])
}

type compoundContainer struct {
	c1 Container
	c2 Container
}

func NewCompoundContainer(first, second Container) Container {
	return &compoundContainer{
		c1: first,
		c2: second,
	}
}

func (c *compoundContainer) Hex(hex string) string {
	return c.Lookup(RawFromHex(hex))
}

func (c *compoundContainer) Lookup(raw string) string {
	one := c.c1.Lookup(raw)
	if one == "" {
		return c.c2.Lookup(raw)
	}
	return one
}
