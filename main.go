package main

import (
	"fmt"
	"time"
)

type Test struct {
	Name string `json:"name,omitempty"`
	Age  int64  `json:"age"`
}

func main() {
	num := 0
	go t(&num)
	t(&num)
	time.Sleep(time.Second)
	fmt.Println(num)
}

func t(num *int) {
	for i := 0; i < 1000; i++ {
		*num++
	}
}
