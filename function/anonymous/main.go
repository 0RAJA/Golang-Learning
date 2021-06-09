package main

import "fmt"

/*
匿名函数:没有名字的函数
	可进行调用或者赋值
	GO语言支持函数式编程
	1.将匿名函数座位函数参数.回调函数
	2.将匿名函数座位另一个函数的返回值,可以形成闭包
*/
func main() {
	//1.创建
	func(a, b int) { //直接创建匿名函数并且调用
		fmt.Println(a + b)
	}(1, 2)
	anonymous1 := func() { //创建匿名函数且赋值
		fmt.Println("匿名函数赋值")
	}
	anonymous1()

	a := func() func() int {
		cnt := 0
		return func() int {
			cnt++
			return cnt
		}
	}
	b := a()
	fmt.Println(b(), b(), b())
	c := fun1()
	fmt.Println(c(), c(), c())
}
func fun1() func() int {
	return func() int {
		var cnt = 0
		cnt++
		return cnt
	}
}
