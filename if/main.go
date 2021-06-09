package main

import "fmt"

func main() {
	if a, b := 1, 2; a < 10 {
		fmt.Println(a, b)
	}
	fmt.Println("ok")
}
