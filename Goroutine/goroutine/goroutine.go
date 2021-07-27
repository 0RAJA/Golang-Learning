package main

import (
	"fmt"
	"time"
)

/*
1.进程:
	进程是一个程序在一个数据集中的一次动态执行过程，可以简单理解为“正在执行的程序”，它是CPU资源分配和调度的独立单位。
	进程一般由程序、数据集、进程控制块三部分组成。
	我们编写的程序用来描述进程要完成哪些功能以及如何完成；
	数据集则是程序在执行过程中所需要使用的资源；
	进程控制块用来记录进程的外部特征，描述进程的执行变化过程，系统可以利用它来控制和管理进程，它是系统感知进程存在的唯一标志。
	进程的局限是创建、撤销和切换的开销比较大。
2.线程:
	线程也叫轻量级进程，它是一个基本的CPU执行单元，也是程序执行过程中的最小单元，由线程ID、程序计数器、寄存器集合和堆栈共同组成。
	一个进程可以包含多个线程。
	线程的优点是减小了程序并发执行时的开销，提高了操作系统的并发性能
	缺点是线程没有自己的系统资源，只拥有在运行时必不可少的资源，但同一进程的各线程可以共享进程所拥有的系统资源，
3.协程:
	协程是一种用户态的轻量级线程，又称微线程，英文名Coroutine，协程的调度完全由用户控制。
	人们通常将协程和子程序（函数）比较着理解。
		子程序调用总是一个入口，一次返回，一旦退出即完成了子程序的执行。
	协程的特点在于是一个线程执行，与多线程相比，其优势体现在：协程的执行效率极高。
	因为子程序切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销，和多线程比，线程数量越多，协程的性能优势就越明显。
4.goroutine
	1.:go中使用goroutine来实现并发。
	2.主goroutine:
		封装main函数的goroutine称为主goroutine。
	3.1
		当新的Goroutine开始时，Goroutine调用立即返回。
		与函数不同，go不等待Goroutine执行结束。
		当Goroutine调用，并且Goroutine的任何返回值被忽略之后，go立即执行到下一行代码。
	3.2
		main的Goroutine应该为其他的Goroutines执行。
		如果main的Goroutine终止了，程序将被终止，而其他Goroutine将不会运行。
*/

func main() {
	if false {
		go fun1()
		go fun2()
		for i := 0; i < 100; i++ {
			fmt.Println("\t\tB", "main")
		}
		time.Sleep(2 * time.Second)
	}
	if true { //分开进行运行
		ch1 := make(chan int)
		ch2 := make(chan int)
		go CountSheep(ch1)
		go CountMoo(ch2)
		for i := 0; i < 10; i++ {
			select {
			case msg := <-ch1:
				fmt.Println("\t", msg)
			case msg := <-ch2:
				fmt.Println(msg)
			}
		}
	}
}
func fun1() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "fun1")
	}
}
func fun2() {
	for i := 0; i < 100; i++ {
		fmt.Println("\tA", "fun2")
	}
}
func CountSheep(ch chan int) {
	for {
		ch <- 1
		time.Sleep(1 * time.Nanosecond)
	}
}

func CountMoo(ch chan int) {
	for {
		ch <- 2
		time.Sleep(2 * time.Nanosecond)
	}
}
