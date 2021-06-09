package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
type Mutex struct {
    // 包含隐藏或非导出字段
}
	Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。
	Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。
func (m *Mutex) Lock()
	Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。
func (m *Mutex) Unlock()
	Unlock方法解锁m，如果m未加锁会导致运行时错误。
	锁和线程无关，可以由不同的线程加锁和解锁。
死锁:
	当两个或两个以上的进程在执行过程中，因争夺资源而处理一种互相等待的状态，
	如果没有外部干涉无法继续下去，这时我们称系统处于死锁或产生了死锁。
注意:
	同一个协程不能连续多次调用 Lock, 否则发生死锁
	锁资源时尽量缩小资源的范围，以免引起其它协程超长时间等待
	mutex 传递给外部的时候需要传指针，不然就是实例的拷贝，会引起锁失败
	善用 defer 确保在函数内释放了锁
	使用 - race 在运行时检测数据竞争问题，go test -race ....，go build -race ....
	善用静态工具检查锁的使用问题
	使用 go-deadlock 检测死锁，和指定锁超时的等待问题 (自己百度工具用法)
	能用 channel 的场景别使用成了 lock
*/

var tickets = 100

var mutex sync.Mutex  //创建互斥锁对象
var wg sync.WaitGroup //创建同步等待组对象

func main() {
	wg.Add(4) //等待4个协程
	go sale("n4")
	go sale("n2")
	go sale("n3")
	go sale("n1")
	wg.Wait() //主程序阻塞
	fmt.Println("main over...")
}

func sale(name string) {
	defer wg.Done() //表示一个协程的结束
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock() //上锁--无法被其他协程抢占
		if tickets <= 0 {
			mutex.Unlock()
			fmt.Println(name, "票售完了")
			break
		}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Nanosecond)
		tickets--
		fmt.Println(name, tickets)
		mutex.Unlock()//解锁
	}
}
