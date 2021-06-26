package main

import "fmt"

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	for i, v := range a {
		for j, x := range v {
			fmt.Println(i, j, x)
		}
	}
}
