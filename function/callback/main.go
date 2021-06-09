package main

import "fmt"

/*
高阶函数:
	根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数。

fun1(),fun2()
	将fun1函数作为了fun2这个函数的参数。

		fun2函数:就叫高阶函数
			接收了一个函数作为参数的函数，高阶函数
		fun1函数:回调函数
			作为另一个函数的参数的函数，叫做回调函数。
*/
func main() {
	a, b := 1, 2
	fmt.Println(operate(a, b, add), operate(a, b, sub), operate(a, b, mul), operate(a, b, div))
}
func add(a, b int) int { //回调函数
	return a + b
}
func sub(a, b int) int { //回调函数
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	if b == 0 {
		fmt.Println("Error")
		return 0
	}
	return a / b
}
func operate(a, b int, f func(int, int) int) int { //高阶函数
	return f(a, b)
}
