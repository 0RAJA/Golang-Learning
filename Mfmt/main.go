package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "hello world"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	/*
		但是，紧跟在verb之前的[n]符号表示应格式化第n个参数（索引从1开始）。同样的在'*'之前的[n]符号表示采用第n个参数的值作为宽度或精度。在处理完方括号表达式[n]后，除非另有指示，会接着处理参数n+1，n+2……（就是说移动了当前处理位置）
	*/
	fmt.Printf("%[2]d %[1]d", 11, 2)
	fmt.Printf("%[3]*.[2]*[1]d", 12, 2, 6)
	fmt.Printf("%6.2d", 12)
}
