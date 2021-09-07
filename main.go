package main

import (
	"fmt"
	"os"
)

func main() {
	path := "D:/Games"
	f, err := os.Stat(path)
	fmt.Println("file:", f, "error:", err)
}
