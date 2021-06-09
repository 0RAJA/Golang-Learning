package main

import (
	"fmt"
)

/*
complex有两种类型:complex64(实部和虚部分别是complex32)和complex128(实部和虚部分别是complex64)
*/
func main() {
	//1.var声明再初始化
	var com1 complex64
	com1 = 10 + 5i
	fmt.Println(com1)
	//2.var 声明并初始化
	var com2 complex64 = 10 + 5i
	var com3 = 10 + 5i //不写类型默认complex128(和电脑有关)
	fmt.Printf("%T %T\n", com2, com3)
	//3.简短声明加类型推断--函数内可简短声明
	com4 := 10 + 5i
	fmt.Println(com4)
	//4.特殊函数
	var com5 = complex(10, 5) //构造
	fmt.Println(com5)
	a := real(com5) //获取实部
	b := imag(com5) //获取虚部
	fmt.Println(a, b)
}
