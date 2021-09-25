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

func Test_getNext(t *testing.T) {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	fmt.Println(1)
	fmt.Println(GetNext(s))
}
