package main

import "fmt"

/*
数据类型:
	一: 按照数据类型来分:
			基本数据类型:
				int, float, string, bool
			复合数据类型:
				array, slice, map, struct, interface.
	二:按照数据的存储特点来分:
			值类型的数据:操作的是数值本身。
				int, float64, bool, string, array
			引用类型的数据:操作的是数据的地址
				slice, map, chan
参数传递:
	A:值传递:
		值类型的都是值传递
	B:引用传递:
		传递的是数据的地址,导致多变量指向同一个地址
不定参数:
	声明: name ...type
	注:不定参数是函数的最后一个参数,不定参数在函数体内相当于切片,切片可作为实参传递给不定参数不过要加上...表示切片元素
*/
func main() {
	s1 := make([]int, 0, 10)
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Println(s1)
	fun1(s1...)
	fmt.Println(s1)
	fun2(s1)
	fmt.Println(s1)
	function := func() []int {
		return s1
	}()
	fmt.Println(function)
}

func fun1(s1 ...int) {
	s1[0] = 10
}

func fun2(s1 []int) {
	s1[0] = 20
}
