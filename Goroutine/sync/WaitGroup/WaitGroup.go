package main

import (
	"fmt"
	"sync"
)

/*
type WaitGroup struct {
    // 包含隐藏或非导出字段
}
	WaitGroup用于等待一组线程的结束。
	父线程调用Add方法来设定应等待的线程的数量。
	每个被等待的线程在结束时应调用Done方法.
	同时，主线程里可以调用Wait方法阻塞至所有线程结束。
func (wg *WaitGroup) Push(delta int)
	Add方法向内部计数加上delta，delta可以是负数；
	如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。
	注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。
	一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。
func (wg *WaitGroup) Done()
	Done方法减少WaitGroup计数器的值，应在线程的最后执行。
func (wg *WaitGroup) Wait()
	Wait方法阻塞当前协程直到WaitGroup计数器减为0。
死锁:
	WaitGroup的计数器始终不为0而此时存在wait则会导致死锁
*/
var wg sync.WaitGroup

func main() {
	wg.Add(2) //设置两条协程要启动
	go fun1()
	go fun2()
	fmt.Println("main 阻塞...")
	wg.Wait() //进入阻塞状态,除非当wg的计数器的值为0
	fmt.Println("main over...")
}

func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("A")
	}
	wg.Done() //计数减一
}

func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("\tB")
	}
	wg.Done()
}
