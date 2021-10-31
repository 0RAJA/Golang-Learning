package main

import "fmt"

func main() {
	var A map[struct{}]struct{}
	fmt.Println(A, A == nil)
}
