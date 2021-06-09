package main

import "fmt"

/*
面向对象:OOP
		继承,封装,多态
提升字段:
	在结构体中属于匿名结构体的字段称为提升字段，因为它们可以被访问，就好像它们属于拥有匿名结构字段的结构一样.
	1.模拟继承性:
		type A struct{
			field
		}
		type B struct{
			A //匿名字段直接访问
		}
	2.模拟聚合关系
		type C struct{
			field
		}
		type D struct{
			c C//聚合关系:需要通过C来进行访问
		}
*/
func main() {
	//创建父类
	p1 := Person{name: "张三", age: 30}
	fmt.Println(p1.name, p1.age)
	s1 := Student{Person: Person{name: "李四", age: 20}, school: "xiYou"}
	fmt.Println(s1)
	//
	var s2 Student
	s2.Person.name = "王麻子"
	s2.Person.age = 20
	s2.school = "xiYou"
	fmt.Println(s2)
	//提升字段
	var s3 Student
	s3.school = "xiYou"
	s3.name = "儿子"
	s3.age = 20
	fmt.Println(s3)
}

/*
	s3.Person.name -> S3.name
	Student结构体将Person结构体作为一个匿名字段了,那么Person中的字段.对于Student来讲,就是提升字段
	Student对象直接访问Person中的字段
*/
type Person struct { //父类
	name string
	age  int
}

type Student struct { //子类
	Person
	school string
}
