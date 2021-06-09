package main

import (
	"fmt"
)

/*
1.结构体指针仍使用 . 访问成员
2.Go不支持指针运算
3.函数中允许返回局部变量的地址:栈逃逸
	函数指针:一个指针，指向了一个函数的指针。
		因为go语言中，function, 默认看作- -个指针，没有*。
		slice , map, function
	指针函数:一个函数，该函数的返回值是一个指针。
*/

func main() {
	//局部变量的地址
	p := fun1()
	*p++
	fmt.Println(*p)
	//数组指针
	arr := [5]int{}
	p1 := &arr
	fmt.Printf("%T\n", p1)
	//指针数组
	arr2 := [5]*int{}
	p2 := &arr2
	fmt.Printf("%T\n", p2)
	//函数指针
	var p3 func() *int
	p3 = fun1
	fmt.Printf("%T\n", p3)
	//指针函数
	p4 := fun2()
	fmt.Println(*p4)
}

func fun1() *int {
	i := 1
	return &i
}

func fun2() *[]int {
	arr := make([]int, 0, 100)
	return &arr
}
