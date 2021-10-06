package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rb := bufio.NewReader(os.Stdin)
	a, size, err := rb.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(a), size)
}
