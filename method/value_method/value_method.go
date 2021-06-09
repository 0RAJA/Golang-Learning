package main

import (
	"fmt"
)

/*
前面介绍方法集时我们知道，具体类型实例变量直接调用其方法时，编译器会所调用方法进行自动转换，
即使接收者是指针的方法，仍然可以使用值类型变量进行调用。下面讨论在以下两种情况下编译器是否会进行方法的自动转换。

(1)通过 类型字面量 显式地进行值调用和表达式调用，可以看到在这种情况下编译器不会做自动转换，会进行严格的方法集检查.
(2)通过 类型变量 进行值调用和表达式调用，在这种情况下，使用值调用(method value)
	方式调用时编译器会进行自动转换，使用表达式调用(methodExpression)方式调用时编译器不会进行转换,会进行严格的方法集检查。例如:
*/
func main() {
	//普通变量调用--自动类型转换
	var a = Data{name: "李四"}
	a.testValue(1)
	a.testPointer(1)
	(&a).testValue(2)
	(&a).testValue(2)
	//1.通过类型字面量进行调用--不进行自动转换
	(*Data)(&struct{ name string }{"张三"}).testPointer(3) //显示调用
	(*Data)(&struct{ name string }{"张三"}).testValue(3)

	(Data)(struct{ name string }{"张三"}).testValue(3)
	Data.testValue(Data{name: "张三"}, 3) //method expression
	//失败案例
	//data.testPointer(data{name: "张三"}, 3)//invalid method expression data.testPointer
	//(data)(struct{ name string }{"张三"}).testPointer(3)//cannot call pointer method on data
	//2.通过类型变量进行值调用和表达式调用--值调用会自动转换,表达式调用不会自动转换
	Data.testValue(a, 4)
	//data.testValue(&a, 4) //invalid method expression data.testPointer
	//data.testPointer(&a, 4)
	//data.testPointer(a, 4)
	(*Data).testPointer(&a, 4)
	(*Data).testValue(&a, 4)
	//值调用会自动转换--下面这些本身就会自动转换所以赋值也一样
	(&a).testValue(4)
	a.testPointer(4)
	(&a).testPointer(4)
	(&a).testValue(4)
	//fmt.Println(unsafe.Sizeof(S{}))
}

/*
	Data 的方法集是testValue
	*Data 的方法集是testValue和testPointer
*/
type Data struct {
	name string
}

type S struct {
	n1 int
	c1 byte
	n2 int
}

func (d Data) testValue(n int) {
	fmt.Println(d.name, n)
}

func (d *Data) testPointer(n int) {
	fmt.Println(d.name, n)
}
