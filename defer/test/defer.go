package main

import "fmt"

func fx(n int) (ret int) {
	defer func() {
		ret += n
	}()
	var f = func() {
		ret += 2
	}
	defer f()
	return n + 1
}
func main() {
	fmt.Println(fx(3))
}

//说明return 先把n+1赋给ret,然后再逐层调用defer
