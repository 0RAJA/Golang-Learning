package myString

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	s1 := "1232233"
	s2 := "33"
	fmt.Println(Kmp(s1, s2))
}

func Test_getNext(t *testing.T) {
	fmt.Println(Kmp("dassadas", "das"))
}
