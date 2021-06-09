package main

import "fmt"

func main() {
	//声明变量
	//第一种:定义变量然后赋值
	var name string
	name = "Raja"
	fmt.Printf("name == %s\n", name) //格式化

	//第二种:一行中定义并赋值
	var age int = 19
	fmt.Println(age) //自动加上一个换行

	//类型推导
	var isOk = true
	fmt.Println(isOk) //直接打印出来

	//简短声明(只能在函数内部声明): :=
	studentOK := true
	fmt.Println(studentOK)

	//推荐使用驼峰式命名 且声明时同时赋值
	var studentName string = "张三"
	fmt.Println(studentName)

	//多个变量同时定义
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	var n1, n2, n3 = 100, 1.23, "Go"
	fmt.Println(n1, n2, n3)

	var (
		x1 = 100
		x2 = 200
		x3 = 300
	)
	fmt.Println(x1, x2, x3)

	//Go语言中变量声明必须被使用--常量声明不需要
	fmt.Println("学生性别:", studentOK)

	//匿名变量 _ 不占用空间,不分配内存,不存在重复声明--表示忽略值用于占位置
	_, m := 10, 5
	fmt.Println(m)
}
