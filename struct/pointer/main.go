package main

import (
	"fmt"
)

//使用new() go原因中专门用于创建某种类型指针的函数
//结构体是值类型
func main() {
	p := Person{
		name:    "Raja",
		age:     16,
		sex:     "female",
		address: "xi`an",
	}
	//定义结构体指针
	var pp = &p
	fmt.Println(pp)
	pp.name = "WW"
	fmt.Println(pp)
	//使用new() go原因中专门用于创建某种类型指针的函数
	pp2 := new(Person) //默认为0值
	fmt.Println(*pp2)
	pp2.name = "张三"
	pp2.age = 18
	pp2.sex = "male"
	pp2.address = "xi`an"
	fmt.Println(*pp2)
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
