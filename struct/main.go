package main

import (
	"fmt"
)

/*
结构体:是由一系列具有相同类型或不同类型的数据构成的数据集合
	结构体成员是由系列的成员变量构成.这些成员变量也被称为“字段”
	值类型:深拷贝
*/
func main() {
	//方法1
	var p1 Person //初始化为0值
	fmt.Println(p1)
	p1.name = "李四"
	//方法2--推荐
	p2 := Person{
		name:    "张三",
		age:     15,
		sex:     "male",
		address: "西安",
	}
	fmt.Println(p2)
	//方法3
	p3 := struct {
		name    string
		age     int
		sex     string
		address string
	}{
		name:    "张三",
		age:     15,
		sex:     "male",
		address: "西安",
	}
	fmt.Println(p3)
	//方法4
	p4 := new(Person) //返回指针
	fmt.Println(*p4)
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
