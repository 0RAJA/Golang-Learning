package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	cnt := 0
	for cnt < 10 {
		fmt.Println(cnt)
		cnt++
	}
	list := make([]int, 10, 10)
	fmt.Println(len(list), cap(list))

	alist := [20]int{1, 2, 3}
	for i, value := range alist {
		fmt.Println(i, value)
	}
	alice := alist[1:3]
	fmt.Println(len(alice), cap(alice))
}
