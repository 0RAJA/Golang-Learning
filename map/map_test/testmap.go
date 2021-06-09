package main

import "fmt"

func main() {
	map1 := make(map[id]Student)
	stu := Student{name: "raja", age: 20}
	map1[1] = stu
	fmt.Println(map1[1])
	//map1[1].age = 23 //无法直接修改其内部成员修改
	stu.age = 23
	map1[1] = stu
	fmt.Println(map1[1])
}

type id int

type Student struct {
	name string
	age  int
}
