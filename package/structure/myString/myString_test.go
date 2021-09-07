package myString

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	s1 := "123"
	s2 := "3"
	fmt.Println(Kmp(s1, s2))
}
