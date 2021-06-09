package main

import "fmt"

/*
由于Go是强类型的语言,如果不满足自动转换的条件,则必须进行强制类型转换。任意两个不相干的类型如果进行强制转换，则必须符合一- 定的规则.
强制类型的语法格式: var a T = (T)(b).使用括号将类型和要转换的变量或表达式的值括起来。
	非常量类型的变量x可以强制转化并传递给类型T,需要满足如下任一条件:
	(1)x可以直接赋值给T类型变量.
	(2)x的类型和T具有相同的底层类型.
	(3)x的类型和T都是未命名的指针类型，并且指针指向的类型具有相同的底层类型。
	(4)x的类型和T都是整型，或者都是浮点型。
	(5)x的类型和T都是复数类型。
	(6)x是整数值或[]byte类型的值，T是string类型。
	(7)x是一个字符串，T是[]byte或[]rune。
*/
func main() {
	//1.相同的底层类型
	type (
		myMap1 map[int]int
		myMap2 map[int]int
	)
	var map1 = myMap1{1: 1, 2: 2, 3: 3}
	var map2 = myMap2{1: 2, 2: 1, 3: 0}
	fmt.Println(map1, map2)
	//map1 = map2 //不能直接赋值(虽然底层类型相同但没有一个是未命名类型)
	map1 = myMap1(map2) //因为底层类型相同可以进行强制类型转换
	fmt.Println(map1, map2)
	//2.字符串和切片的转换
	s := "HelloWorld"
	world := make([]byte, 20)
	world = []byte(s)
	fmt.Printf("%s %s\n", world, s)
	/*
		(1)数值类型和string类型之间的相互转换可能造成值部分丢失;其他的转换仅是类型的转换,不会造成值的改变。
			string和数字之间的转换可使用标准库strconv.
		(2)Go语言没有语言机制支持指针和interger之间的直接转换,可以使用标准库中的unsafe包进行处理。
	*/
	//3.浮点类型和整型
	f1 := 2.3
	num1 := int(f1)
	f2 := float64(num1)
	fmt.Println(num1, f2)
	//4.
	x1 := []rune{'你', '好', '李', '1'}
	str1 := string(x1)
	fmt.Println(str1)
}
