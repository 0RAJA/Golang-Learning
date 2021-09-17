package TrieNo

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	trie := Constructor()
	trie.Insert("abcde", "1")
	trie.Insert("abcde", "2")
	fmt.Println(trie.Search("abcde"))
}
