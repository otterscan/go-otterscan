package triemap

import "github.com/openacid/slim/trie"

type SlimTrieBackend struct {
	*trie.SlimTrie
}

func (s *SlimTrieBackend) Get(x string) (uint64, bool) {
	val, ok := s.SlimTrie.GetI64(x)
	if !ok {
		return 0, false
	}
	return uint64(val), ok
}
