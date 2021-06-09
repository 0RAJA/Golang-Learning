package main

import (
	"fmt"
)

/*
OOP中的继承性:
	如果两个类(class)存在继承关系，其中一个是子类，另一个作为父类，那么:
		1.子类可以直接访问父类的属性和方法:❤
		2.子类可以新增自己的属性和方法❤
		3.子类可以重写父类的方法(rewrite,就是将父类已有的方法，重新实现)❤

	模拟继承性:
		type A struct{	//父类
			field
		}
		type B struct{	//子类
			A 			//匿名字段直接访问
			field		//新增属性
		}
		A的方法B可以使用或者重写.
		A也可以新增自己的方法
*/
func main() {
	p1 := Person{name: "人类", age: 50}
	s1 := Student{Person: Person{name: "学生", age: 18}, school: "xiYou"}
	c1 := Child{Student: Student{Person: Person{name: "儿童", age: 3}, school: "xiXou"}, isCrying: false}
	p1.eat()
	s1.eat()
	c1.eat()
}

//Person 父类
type Person struct {
	name string
	age  int
}

//Student 是 Person 的子类
type Student struct {
	Person
	school string
}

//Child 是 Student 的子类
type Child struct {
	Student
	isCrying bool
}

//方法
func (p Person) eat() { //子类继承父类的方法
	fmt.Println(p.name, "正在吃饭")
}

func (s Student) eat() { //子类重写父类的方法
	fmt.Println(s.name, "正在", s.school, "吃饭")
}

func (c Child) eat() {
	fmt.Println(c.isCrying, "\b爱哭的", c.name, "正在", c.school, "吃饭")
}
