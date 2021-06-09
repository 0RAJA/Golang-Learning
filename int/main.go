package main

import "fmt"

//整型

func main() {
	//十进制
	var i1 int = 101       //int 看电脑--int32
	fmt.Printf("%d\n", i1) // 101
	fmt.Printf("%o\n", i1) //八进制
	fmt.Printf("%x\n", i1) //十六进制
	fmt.Printf("%b\n", i1) //二进制
	//八进制
	i2 := 010
	fmt.Printf("%d\n", i2) // 8
	//十六进制
	i3 := 0x123
	fmt.Printf("%d\n", i3) // 16

	//查看变量类型
	i1 = 2<<30 - 1
	fmt.Println(i1)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%T\n", int64(i1))
}
