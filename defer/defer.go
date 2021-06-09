package main

import "fmt"

/*
defer的词义: "延迟"，"推迟"
在go语言中，使用defer关键字来延迟一个函数或者方法的执行。
1. defer函数或方法:一个函数或方法的执行被延迟了。
2。defer的用法:
	A:对象.close() ,临时文件的删除...
		文件.open()
		defer close()
		读或写
	B: go语言中关于异常的处理，使用panic()和recover()
		panic函数用于引发恐慌，导致程序中断执行
		recover函数用于恢复程序的执行，recover()语法上要求必须在defer中执行。
3.如果多个defer函数:
	先进后出
4. defer函数传递参数的时候;
	在注册时就将参数传入了,之后实参的改变不会影响形参
5. defer函数注意点:
	遇到return语句时先把之前的defer全部执行完才会执行return
	使用os.Exit()退出进程时不会执行defer
*/
func main() {
	fmt.Println("First")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
