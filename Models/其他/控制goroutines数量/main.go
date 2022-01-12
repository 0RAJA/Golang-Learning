package main

import (
	"fmt"
	"runtime"
	"sync"
)

func work1(out chan struct{}, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(id, runtime.NumGoroutine())
	<-out
}

// Test1 利用chan和wg进行控制
func Test1() {
	works := 10000
	goroutinesNum := 5
	in := make(chan struct{}, goroutinesNum)
	wg := sync.WaitGroup{}
	for i := 1; i <= works; i++ {
		in <- struct{}{}
		wg.Add(1)
		go work1(in, i, &wg)
	}
	wg.Wait()
}

func work2(task chan int, wg *sync.WaitGroup) {
	for id := range task {
		fmt.Println(id, runtime.NumGoroutine())
		wg.Done()
	}
}

// SendTask 协程池,工人和任务模型
func SendTask(id int, task chan int) {
	task <- id
}

func Test2() {
	wg := sync.WaitGroup{}
	workers := 10
	taskCnt := 1000
	task := make(chan int)
	for i := 0; i < workers; i++ {
		go work2(task, &wg)
	}
	for i := 1; i <= taskCnt; i++ {
		wg.Add(1)
		SendTask(i, task)
	}
	wg.Wait()
}

func main() {
	//Test1()
	//Test2()
}
