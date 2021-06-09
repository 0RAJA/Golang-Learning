package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
runtime包提供和go运行时环境的互操作，如控制go程的函数。
它也包括用于reflect包的低层次类型信息；参见reflect报的文档获取运行时类型系统的可编程接口。
*/
func init() { //在程序之前使用
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//获取GORoot目录
	fmt.Println("GOROOT->", runtime.GOROOT())
	//获取操作系统
	fmt.Println("os/platform->", runtime.GOOS)
	//获取当前电脑CPU数量
	fmt.Println("CPU数量->", runtime.NumCPU())
	//设置Go执行的最大CPU的数量[1,256]在使用Go之前设置[1.8]版本之前才使用
	if false {
		runtime.GOMAXPROCS(2)
		fmt.Println("CPU数量->", runtime.NumCPU())
	}
	//Gosched()//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
	if false {
		//fmt.Println(runtime.NumGoroutine())
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println("fun()")
			}
		}()
		//fmt.Println(runtime.NumGoroutine())
		for i := 0; i < 10; i++ {
			runtime.Gosched() //每次到这个就让出时间片-据说也会抢占资源,可能会不一样
			fmt.Println("\tmain")
		}
	}
	if false {
		fmt.Println("Goroutine开始")
		go fun1()
		runtime.Gosched()
		fmt.Println("Goroutine结束")
		time.Sleep(2 * time.Second)
	}
}

func fun1() {
	defer fmt.Println("fun1()1")
	go func() {
		fmt.Println("\tnew Goroutine")
	}()
	//return //只是终止这个函数
	runtime.Goexit() //终止当前的Goroutine
	defer fmt.Println("fun1()2")
}
