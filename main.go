package main

import "fmt"

func main() {
	var a map[int]int
	a = make(map[int]int)
	fmt.Println(a == nil)
}
