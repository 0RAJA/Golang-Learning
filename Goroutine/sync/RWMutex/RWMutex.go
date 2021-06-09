package main

import (
	"fmt"
	"sync"
	"time"
)

/*
type RWMutex struct {
    // 包含隐藏或非导出字段
}
	RWMutex是读写互斥锁。
	该锁可以被同时多个读取者持有或唯一个写入者持有。
	RWMutex可以创建为其他结构体的字段；零值为解锁状态。
	RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
func (rw *RWMutex) Lock()
	Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入。
func (rw *RWMutex) Unlock()
	Unlock方法解除rw的写入锁状态，如果m未加写入锁会导致运行时错误。
func (rw *RWMutex) RLock()
	RLock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取。
func (rw *RWMutex) RUnlock()
	RUnlock方法解除rw的读取锁状态，如果m未加读取锁会导致运行时错误。
func (rw *RWMutex) RLocker() Locker
	RLocker方法返回一个互斥锁，通过调用rw.RLock和rw.RunLock实现了Locker接口。
总结:
	读锁不能阻塞读锁
	读锁需要阻塞写锁，直到所有读锁都释放
	写锁需要阻塞读锁，直到所有写锁都释放
	写锁需要阻塞写锁
*/
var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex) //读写锁
	wg = new(sync.WaitGroup)    //同步等待组
	if true {                   //验证读锁可以同时进行读操作,但不能写
		wg.Add(3)
		go readDate(1)
		go writeDate(2)
		go readDate(3)
	}
	if false { //写锁不可同时进行写操作,也不能同时进行读操作
		wg.Add(3)
		go writeDate(1)
		go readDate(2)
		go writeDate(3)
	}
	wg.Wait()
	fmt.Println("main over...")
}

func readDate(n int) { //读锁后 可多次读锁,但不能写锁
	defer wg.Done()
	rwMutex.RLock() //读操作上锁
	fmt.Println(n, "开始读...")
	fmt.Println(n, "正在读取...")
	time.Sleep(100 * time.Nanosecond)
	fmt.Println(n, "结束读...")
	rwMutex.RUnlock() //读操作解锁
}

func writeDate(n int) { //写锁后 不能读锁也不能写锁
	defer wg.Done()
	rwMutex.Lock()
	fmt.Println(n, "开始写...")
	fmt.Println(n, "正在写...")
	time.Sleep(100 * time.Nanosecond)
	fmt.Println(n, "结束写...")
	rwMutex.Unlock()
}
