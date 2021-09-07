package Trie

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	tire := Constructor()
	tire.Insert("abc", "a")
	tire.Insert("abcd", "hdd")
	tire.Insert("babcc", "hsadsd")
	ret := tire.Search("abc")
	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i])
	}
}
