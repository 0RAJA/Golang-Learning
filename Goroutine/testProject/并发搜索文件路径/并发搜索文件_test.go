package main

import (
	"fmt"
	"testing"
)

func Test_kmp(t *testing.T) {
	fmt.Println(kmp("123", "23"))
	next := make([]int, len("abaabcac")+1)
	getNext("abaabcac", next)
	fmt.Println(next)
}

func Test_inputMessage(t *testing.T) {
	inputMessage()
}
