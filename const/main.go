package main

import "fmt"

// PI 常量定义之后不能修改
const PI = 3.1415926

//批量声明常量
const (
	statusOK = 200
	noFound  = 404
)

//批量声明常量值,如果某一行声明后没有赋值,默认和上一行一样
const (
	n1 = 100
	n2 //100
	n3 //100
)

//iota 常量计数器,只能在常量的表达式中使用
//iota 在const关键字出现时被重置为0, !const中每新增一行将使iota计数一次! --可以定义枚举

const (
	num1 = iota //num1 == 0 iota == 0
	_           //		   iota == 1
	num2 = iota //num2 == 2 iota == 2
	num3 = iota //num3 == 3 iota == 3
)

//多个变量声明在一行 -- iota 只有新增一行才行
const (
	d1, d2 = iota, iota // d1 == 0 d2 == 0 iota == 0
	d3, d4 = iota, iota // d3 == 1 d4 == 1 iota == 1
)

//定义数量级
const (
	_        = iota
	KB       = 1 << (10 * iota) //1024
	MB       = 1 << (10 * iota)
	GB       = 1 << (10 * iota)
	TB int64 = 1 << (10 * iota)
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(num1, num2, num3)
	fmt.Printf("d1 == %d d2 ==%d\nd3 == %d d4 == %d\n", d1, d2, d3, d4)
	fmt.Println(KB, MB, GB, TB)
}
