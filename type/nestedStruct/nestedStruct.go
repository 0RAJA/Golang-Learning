package main

import "fmt"

type Person struct {
	name string
}

//创建方法
func (p Person) show() {
	fmt.Println("Person:", p.name)
}
func (p People) show1() { //相当于给Person加方法
	fmt.Println("People:", p.name)
}

type People = Person //起别名

// Student 结构体嵌套
type Student struct {
	People
	Person
}

func main() {
	var s Student
	//s.name = "raja" //混淆
	s.People.name = "张三"
	s.Person.name = "李四"
	//s.show() //混淆
	s.Person.show()
	s.People.show()

	fmt.Printf("%T %T\n", s.Person, s.People) //main.Person main.Person 起别名

}
