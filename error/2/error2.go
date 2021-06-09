package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println(ins.Op, "\n", ins.Path, "\n", ins.Err)
		}
		return
	}
	fmt.Println(f.Name())
}
