package main

import "fmt"

/*
	接口: interface--常以er结尾
		在Go中，接口是一组方法签名。--不需要func引导且只有声明没有实现
	当某个类型为这个接口中的所有方法提供了方法的实现，它被称为实现接口。
	Go语言中，接口和类型的实现关系，是非侵入式
	//其他语言中，要显示的定义
		class Mouse impLaments USB{}

	1.当需要接口类型的对象时,可以使用任意实现类对象代替
	2.接口对象不能访问实现类中的属性

	多态:一个事物的多种形态
		go语言通过接口模拟多态

		就一个接口的实现
			1.看成实现本身的类型，能够访问实现类中的属性和方法
			2.看成是对应的接口类型，那就只能够访问接口中的方法

	接口的用法:
		1.一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数。
		2.定义一个类型为接口类型，实际上可以赋值为任意实现类的对象

	接口的类型:
		1.动态类型:接口所绑定的具体实例的类型
		2.静态类型:接口的方法集
*/
func main() {
	var inter USBer //定义接口类型--只能访问对应的方法
	//1.创建Mouse类型
	m1 := Mouse{name: "罗技"}
	f1 := FlashDish{name: "U盘"}
	testInterface(m1)
	testInterface(f1)
	inter = m1 //接口初始化--没有初始化的接口是没有意义的
	/*
		接口初始化:
		1.实例对象赋值接口
		2.接口赋值接口
			例:接口A = 具体类型实例T(T的方法集是A方法集的超集)
			   接口B = 接口A(A的方法集是B方法集的超集)
			此时B绑定的具体实例是A绑定的具体实例T的拷贝
	*/
	inter.start(1)
	//fmt.Println(inter.name)//接口只能访问其对应的方法,不能访问属性
	fmt.Println(m1.name) //而实现类属性的变量可以访问其对应方法和属性
	//多态
	f1.deleteData() //f1此时作为U盘可以访问其属性
	var inter2 USBer
	inter2 = f1 //f1现在作为USB可以访问其方法
	inter2.start(1)

	var arr [3]USBer
	arr[0] = m1
	arr[1] = f1
	fmt.Println(arr)
}

//1.定义接口

type USBer interface {
	//方法只有声明,没有实现,写方法参数和返回值
	start(int) int
	end()
}

//2.实现类

type Mouse struct {
	name string
}

type FlashDish struct {
	name string
}

func (m Mouse) start(int) int {
	fmt.Println(m.name, "鼠标开始")
	return 0
}
func (m Mouse) end() {
	fmt.Println(m.name, "鼠标结束")
}
func (f FlashDish) start(int) int {
	fmt.Println(f.name, "U盘开始")
	return 0
}
func (f FlashDish) end() {
	fmt.Println(f.name, "U盘结束")
}

//3.测试方法
func testInterface(usb USBer) { //相当于给接口赋值,然后通过接口来调用对应方法
	usb.start(1)
	usb.end()
}

func (f FlashDish) deleteData() {
	fmt.Println(f.name, "已删除")
}
