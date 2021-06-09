package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
输入和输出:
	fmt包:输入，输出
输出:
	Print() //打印
	Printf() //格式化打印
	PrintLn() //打印之后换行
格式化打印占位符:
	%v,原样输出
		%#v,带类型符输出
	%T,打印类型
	%t,bool类型
	%s,字符串
	%f,浮点
	%d, 10进制的整数
	%b, 2进制的整数
	%o, 8进制
	%x,%X 16进制
		%x: 0-9, a-f
		%X: 0-9, A-F
	%C,打印字符
	%p,打印地址
*/
func main() {
	//var x int
	//var y float64
	//fmt.Println("请输入一个整数和浮点数")
	//fmt.Scanln(&x, &y) //读取一行中间以空格分隔
	//fmt.Scanf("%d%f",&x,&y)//格式化读取,不能用\n来分隔
	//fmt.Printf("%d,%f", x, y)
	//bufio包 文件的输入输出
	fmt.Println("输入一个字符串")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Printf("读到的数据:%s", s1)

	for _, i := range s1 {
		if i == '\n' {
			fmt.Println("读到\\n")
		}
	}
}
