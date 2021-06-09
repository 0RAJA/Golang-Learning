package main

import "fmt"

/*
匿名结构体和匿名字段;

匿名结构体:没有名字的结构体，
	在创建匿名结构体时，同时创建对象
	 变量名 := struct{
		定义字段Field
	}{
		字段赋值
	}

匿名字段:一个结构体的字段没有字段名--被匿名嵌入的字段必须是命名类型或命名类型的指针
在定义struct 的过程中，如果字段只给出字段类型，没有给出字段名，则称这样的字段为
“匿名字段”。被匿名嵌入的字段必须是命名类型或命名类型的指针，类型字面量不能作为匿名
字段使用。匿名字段的字段名默认就是类型名，如果匿名字段是指针类型，则默认的字段名就
是指针指向的类型名。但一个结构体里面不能同时存在某一类型及其指针类型的匿名字段，原
因是二者的字段名相等。如果嵌入的字段来自其他包，则需要加上包名，并且必须是其他包可以导出的类型

匿名函数:
*/
func main() {
	p1 := struct { //匿名结构体
		name string
		age  int
	}{
		name: "ww",
		age:  15,
	}
	p2 := Worker{
		name: "Raja",
		age:  15,
	}
	p3 := Student{
		string: "raja",
		int:    14,
	}
	fmt.Println(p1, p2, p3.string)
}

type Worker struct {
	name string
	age  int
}

type Student struct { //匿名字段默认使用类型作为名字,匿名字段类型不能重复
	string
	int
}
