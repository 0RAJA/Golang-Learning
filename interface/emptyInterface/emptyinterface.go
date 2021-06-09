package main

import "fmt"

/*
	空接口(interface{})-->任意类型
		不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。
*/
func main() {
	var cat = Cat{colour: "red"}
	var person = Person{name: "Raja", age: 19}
	test1(cat, person) //任意类型
	var temp A
	fmt.Printf("%T\n", temp) //nil类型
	//应用:
	//map
	map1 := make(map[interface{}]interface{})
	map1["12"] = 2
	map1[21] = "2"
	fmt.Println(map1)
	//slice
	slice := []interface{}{1, "2", true, 'p', cat, person}
	fmt.Println(slice)
	//等等
}

//接口A是空接口,代表任意类型
func test1(arr ...A) { //空接口切片
	fmt.Println(arr)
}

type A interface{}

type Cat struct {
	colour string
}
type Person struct {
	name string
	age  int
}
