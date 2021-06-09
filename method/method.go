package main

import (
	"fmt"
)

/*
方法: method
	一个方法就是一个包含了接受者的函数.接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
	所有给定类型的方法属于该类型的方法集

语法:
	func (接受者) 方法名 (参数列表) (返回值列表){

	}
	总结: method, 同函数类似,区别需要有接受者.(也就是调用者)

	对比函数:
		A:意义
			方法:某个类别的行为功能，需要指定的接受者调用
			函数:一段独立功能的代码，可以直接调用
		B:语法
			方法:方法名可以相同，只要接受者不同
			函数:命名不能冲突
注意:
	类型方法有如下特点:
	(1)可以为命名类型增加方法(除了接口)，非命名类型不能自定义方法。
		比如不能为[]int类型增加方法，因为[]int 是非命名类型。
		命名接口类型本身就是一个方法的签名集合，所以不能为其增加具体的实现方法。
	(2)为类型增加方法有一个限制,就是方法的定义必须和类型的定义在同一个包中。
		不能再为int bool等预声明类型增加方法，因为它们是命名类型，但它们是Go语言内置的
		预声明类型，作用域是全局的,为这些类型新增的方法是在某个包中.这与第2条规则冲突,所以Go编译器拒绝为int增加方法。
	(3)方法的命名空间的可见性和变量一样，大写开头的方法可以在包外被访问，否则只能在包内可见。
	(4)使用type定义的自定义类型是一个新类型，新类型不能调用原有类型的方法，但是底层类型支持的运算可以被新类型继承。

方法值:
	变量x的静态类型是T, M是类型T的一个方法，x.M被称为方法值(method value)。x.M
	是一个函数类型变量，可以赋值给其他变量，并像普通的函数名一样使用。
*/
func main() {
	w1 := Worker{name: "张三", age: 30, sex: "男"}
	w1.work()
	w2 := &Worker{name: "李四", age: 29, sex: "女"}
	fmt.Printf("%T\n", w2)
	w2.work() //可以通过结构体指针直接调用结构体方法--传递结构体
	w1.rest() //可以通过结构体直接调用结构体指针方法--传递结构体指针
	//指针和本身共用方法--内部编译器会自动转换

	c1 := Cat{color: "red", age: 1}
	c1.PrintInfo()
	w1.printInfo()
	wfun := w1.printInfo //方法值--隐藏接受者
	wfun()
	fmt.Println("-----------------")
	wfun2 := Worker.printInfo //方法表达式--显式地把接受者传递进去
	fmt.Printf("%T\n", wfun2)
	wfun2(w1)

	arr := Arr{1, 2, 3, 5}
	arr.add()
	f1 := (&w1).rest
	f2 := (&w1).rest
	f1()
	f2()

}

//定义结构体

type Worker struct {
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}
type Arr []int

//定义行为方法--函数不可同名,方法在接收方不同的情况下可以同名

func (w Worker) work() { //w = w1
	fmt.Println(w.name, "在工作...")
}

func (w *Worker) rest() {
	fmt.Println(w.name, "在休息")
	fmt.Printf("%p\n", w)
}

func (w Worker) printInfo() {
	fmt.Println(w.name, w.age, w.sex)
}

func (c Cat) PrintInfo() {
	fmt.Println(c.color, c.age)
}

func (arr Arr) add() { //为命名类型创建方法
	sum := 0
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)
}
func (w *Worker) myPrint() {
	fmt.Printf("%p\n", w)
}
