package main

import "fmt"

/*
go语言的数据类型:
	基本数据类型;
		int, float, bool, string ,complex
	复合数据类型;
		array,     slice,   map,  function, pointer, struct, interface。 。
		[size]type []type   map[keyType][valueType]
	函数的类型:
		func(参数列表的类型)(返回值列表的数据类型)
*/
func main() {
	fmt.Printf("%T\n", main)
	fmt.Printf("%T\n", fun2)
	fmt.Printf("%T\n", fun1)
	//定义一个函数类型变量
	var c func(int, int) (int, int)
	c = fun2
	fmt.Printf("c == %T\n", c)
	fmt.Printf("%T\n", c)
	var c1 = func() {
		fmt.Println("OK")
	}
	c1()
}
func fun1(a, b int) (f func(int, int) (int, int)) {
	a = a + b
	return fun2
}

func fun2(a, b int) (int, int) {
	return a + b, a - b
}
