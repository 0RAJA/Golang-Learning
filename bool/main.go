package main

import "fmt"

//bool值
func main() {
	b1 := true
	var b2 bool                   //默认是false
	fmt.Printf("%T %t\n", b1, b1) // %t是bool类型的输出
	fmt.Printf("%T %t", b2, b2)   //%T是类型 %v是值
}
