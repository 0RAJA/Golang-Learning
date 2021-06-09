package main

import "fmt"

//类似C语言里的宏展开
func main() {
	var cat Cat //cat是Cat类型
	cat.test1()
	cat.test2()
	cat.test3()
	var a1 A               //A1是A接口
	fmt.Printf("%T\n", a1) //nil类型
	a1 = cat               //把cat类型看成是A接口类型
	a1.test1()
	var b1 B = cat //把cat的类型看成B接口的类型
	b1.test2()
	var c1 C = cat //把cat类型看成是C接口类型
	c1.test1()
	c1.test3()
	c1.test3()
	var a2 A = c1 //把C接口类型看成A接口类型则c1只能访问A接口的方法
	a2.test1()
}

type A interface { //A的实现类只需要实现test1
	test1()
}

type B interface { //B的实现类只需要实现test2
	test2()
}

type C interface { //C的实现类需要实现test1,test2,test3
	A
	B
	test3()
}

type Cat struct { //ABC接口的实现

}

func (c Cat) test1() {
	fmt.Println("test1")
}
func (c Cat) test2() {
	fmt.Println("test2")
}
func (c Cat) test3() {
	fmt.Println("test3")
}
