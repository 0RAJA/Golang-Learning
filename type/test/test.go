package main

import "fmt"

func main() {
	type A string         //string 		     预声明类型
	type B A              //string 			 B的类型向上查找
	type M map[string]int //map[string]int   类型字面量
	type N M              //map[string]int   N的类型向上查找
	type P *N             //*N				 类型字面量
	type S string         //string 		     预声明类型
	type T map[S]int      //map[S]int 	     类型字面量
	type U T              //map[S]int 		 向上查找
	var (
		t T
		u T
	)
	t = u
	fmt.Printf("%T\n", t)
	var a int
	var b int
	fmt.Printf("%T %T\n", a, b)
	//var a T1
	//var b T2
	//var num1 int = 10
	//var fl1 float64 = 2.3
	//num1 = int(fl1)
	//fl1 = float64(30)
	//struct {
	//	name string
	//	age int
	//}
	//var a interface{eat()}
	s := Student{name: "张三", age: 19}
	s.eat()
	Student.eat(s)
}

type name interface {
	eat()
}

type Student struct {
	name string
	age  int
}

func (s Student) eat() {
	fmt.Println(s.name, "正在吃饭")
}
