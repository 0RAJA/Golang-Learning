package main

import (
	"fmt"
	"math"
)

/*
方法一:接口类型断言:
	i.(TypeName)
		i必须是接口变量
		TypeName可以是具体类型名或者接口类型名
			若TypeName是具体类型名:判断i所绑定的实例类型是否就是具体类型TypeName
			若TypeName是接口类型名:判断i所绑定的实例对象是否同时实现了TypeName接口
	方式一:
		1. o := i.(TypeName) //不安全， 会panic()
			(1) TypeName是具体类型名,此时如果接口i绑定的实例类型就是具体类型TypeName,
				则变量o的类型就是TypeName,变量o的值就是i接口绑定的实例值的副本(当然实例可能是
				指针值,那就是指针值的副本)。
			(2) TypeName是接口类型名,如果接口i绑定的实例类型满足接口类型TypeName,则变
				量o的类型就是接口类型TypeName，o底层绑定的具体类型实例是i绑定的实例的副本(当然
				实例可能是指针值，那就是指针值的副本)。
			(3)如果上述两种情况都不满足，则程序抛出panic。
		2. o, ok := i.(TypeName) //安全
			o的类型和方法一一致,唯一不同在于方法一如果不匹配就panic(),而方法二可以避免
			(3)如果上述两个都不满足，则ok为false(满足一个就是true), 变量o是TypeName类型的“零值”，此种条
		       件分支下程序逻辑不应该再去引用o，因为此时的o没有意义。
方式二:接口类型查询:
	switch v := i.(type){
	case type1:
		...
	case type2:
		...
	...
	}
接口查询有两层语义，一是查询一个接口变量底层绑定的底层变量的具体类型是什么,二是查询接口变量绑定的底层变量是否还实现了其他接口。
(1) i必须是接口类型。
	具体类型实例的类型是静态的，在类型声明后就不再变化，所以具体类型的变量不存在类型查询，
	类型查询一定是对一个接口变量进行操作。也就是说上文中的i必须是接口变量,如果i是未初始化接口变量，则v的值是nil.
(2) case 字句后面可以跟非接口类型名，也可以跟接口类型名，匹配是按照case子句的顺序进行的。
		●如果case后面是一个接口类型名，且接口变量i绑定的实例类型实现了该接口类型的方法,
			则匹配成功，v的类型是接口类型，v底层绑定的实例是i绑定具体类型实例的副本.
		●如果case后面是一个具体类型名，且接口变量i绑定的实例类型和该具体类型相同，则匹配成功,
			此时v就是该具体类型变量，v的值是i绑定的实例值的副本。
		●如果case后面跟着多个类型，使用逗号分隔，接口变量i绑定的实例类型只要和其中一个类型匹配，则直接使用i赋值给v,相当于V := i.
			这个语法有点奇怪,按理说编译器不应该允许这种操作,语言实现者可能想让type switch语句和普通的switch语句保持一样的语法规则.
		●如果所有的case字句都不满足，则执行default语句，此时执行的仍然是v:=i,最终v的值是i。此时使用v没有任何意义。
		●fallthrough语句不能在Type Switch 语句中使用。
*/
func main() {
	t1 := Triangle{a: 3, b: 4, c: 5}
	c1 := Circle{radius: 5}
	var s shaper = c1
	fmt.Println(s.peri(), s.area())
	getType(t1)
	getType(c1)
	getType(s) //s的值类型是c1
	var p *Triangle = &Triangle{a: 4, b: 5, c: 6}
	fmt.Printf("p:%T %p %p\n", p, &p, p)
	getType2(p)
	getType(p)
	getType2(t1)
	fmt.Println("------------------------")
	t2 := t1
	getType3(t2)
	var num int = 23
	getType3(num)
}

//1.定义接口
type shaper interface {
	peri() float64 // 形状周长
	area() float64 //形状面积
}
type allShaper interface {
	peri() float64 //形状周长
}

// Triangle 定义实现类:三角形
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt((p - t.a) * (p - t.b) * (p - t.c))
	return s
}

// Circle 定义实现类:圆形
type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

//断言
func getType(s shaper) {
	if ins, ok := s.(Triangle); ok { //ins类型是Triangle值是s所绑定的具体实例
		fmt.Println("s是Triangle实例对象:", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(allShaper); ok {
		fmt.Println("s所绑定的实例可以实现allShaper接口", ins.peri()) //ins类型是allShaper值是s所绑定的具体实例
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("s是Circle的实例对象", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("%p %#x\n", ins, ins)
	} else {
		fmt.Println("Error")
	}
}

//查询
func getType2(s shaper) {
	switch ins := s.(type) { //ins是该类型的实例对象
	case Triangle: //ins类型是Triangle值是s所绑定的具体实例
		fmt.Println("Triangle:", ins.peri(), ins.area())
	case Circle:
		fmt.Println("Circle:", ins.peri(), ins.area())
	case *Triangle:
		fmt.Printf("ins:%T %p %p\n", ins, &ins, ins)
		fmt.Printf("s:%T %p %p\n", s, &s, s)
	case allShaper:
		fmt.Println("s所绑定的实例可以实现allShaper接口", ins.peri()) //ins类型是allShaper值是s所绑定的具体实例
	default:
		fmt.Println("Error")
	}
}

func getType3(s interface{}) { //s所绑定的实例对象可以同时实现allShaper和Triangle接口
	switch ins := s.(type) {
	case allShaper: //allShaper接口
		fmt.Printf("%T %v\n", ins, ins.peri())
	case Triangle: //Triangle接口)
		fmt.Printf("%T %v %v %v\n", ins, ins.a, ins.b, ins.c)
	default:
		fmt.Printf("%T %v\n", ins, ins)
	}
}
