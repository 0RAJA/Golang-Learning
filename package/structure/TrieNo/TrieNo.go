package TrieNo

type Trie struct {
	suffix  map[byte]*Trie
	results []string
}

func Constructor() *Trie {
	return &Trie{suffix: map[byte]*Trie{}}
}

func (t *Trie) Insert(k, v string) {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] == nil {
			root.suffix[k[i]] = &Trie{suffix: make(map[byte]*Trie)}
		}
		root = root.suffix[k[i]]
	}
	root.results = append(root.results, v)
}

func (t *Trie) Search(k string) []string {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] != nil {
			root = root.suffix[k[i]]
		} else {
			return []string{}
		}
	}
	return root.results
}
