package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
分支语句: if, switch, select
select语句类型于switch语句
	但是select语句会随机执行一个可运行的case
	如果没有case可以运行，要看是否有default,如果有就执行default,否则就进入阻塞，直到有case可以运行
	每一个case后的内容都会被求值--**需要解释**
	注意:
		每个 case 都必须是一个通信
		所有 channel 表达式都会被求值
		所有被发送的表达式都会被求值
		如果任意某个通信可以进行，它就执行，其他被忽略。
		如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		否则：
			如果有 default 子句，则执行该语句。
			如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
扇入(Fan in)扇出(Fan out)
	所谓的扇入是指将多路通道聚合到一条通道中处理，
	Go语言最简单的扇入就是使用select聚合多条通道服务;
	所谓的扇出是指将一条通道发散到多条通道中处理.在Go语直里面具体实现就是使用ga关键字启动多个goroutine并发处理。
通知退出机制
	读取已经关闭的通道不会引起阻塞，也不会导致panic,而是立即返回该通道存储类型的零值。
	关闭select监听的某个通道能使select立即感知这种通知，然后进行相应的处理，这就是所谓的退出通知机制。
	多个监听器监听一个chan时,chan的退出通知会传递到所有的监听器
*/
func main() {
	if false { //select的使用--发送信号
		func() {
			ch := make(chan int)
			go func() {
				for {
					select { //随机选择一个可以执行的
					case ch <- 0:
					case ch <- 1:
					default:
						ch <- 2
					}
				}
			}()
			for i := 0; i < 10; i++ {
				fmt.Println(<-ch)
			}
		}()
	}
	if false { //select的使用--接收信号
		func() {
			ch := make(chan int)
			go func() {
				//time.Sleep(2 * time.Second)
				ch <- 100
			}()
			time.Sleep(1 * time.Second)
			select {
			case num, ok := <-ch: //尝试进行通信,可以判断通道是否关闭
				if ok {
					fmt.Println(num)
				} else {
					fmt.Println("ERROR")
				}
			default:
				fmt.Println("default")
			}
		}()
	}
	if false { //通知退出机制的理解--随机数的生成
		func() {
			done := make(chan int) //通知chan
			ch := randNum(done)    //ch是生产chan
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			close(done) //通知生成函数关闭chan
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			fmt.Println("over")
		}()
	}
	if false { //随机数生成器
		func() {
			ch := generateIntA()
			for i := 0; i < 10000; i++ {
				<-ch
			}
		}()
	}
	if false { //随机数生成器增强型--多源
		func() {
			ch := generateInt()
			for i := 0; i < 10000; i++ {
				<-ch
			}
		}()
	}
	if true { //）一个融合了并发、缓冲、退出通知等多重特性的生成器。
		func() {
			fmt.Println("Go:", runtime.NumGoroutine())
			done := make(chan int)
			fmt.Println("建立done")
			ch := randNumInt(done)
			time.Sleep(1 * time.Second)
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			time.Sleep(1 * time.Second)
			fmt.Println("Go:", runtime.NumGoroutine())
			close(done)
			time.Sleep(1 * time.Second)
			fmt.Println(<-ch)
			fmt.Println(<-ch)
		}()
	}
	if false { //管道
		func() { //累加函数
			fun := func(in chan int) chan int {
				out := make(chan int) //结果chan
				go func() {             //累加goroutine
					for i := range in { //只要in中加入一个数加一然后送到结果数组out
						out <- i + 1
					}
					close(out) //计算完毕就关闭结果chan,使得最后的打印range结束
				}()
				return out
			}
			in := make(chan int) //输入chan
			go func() {
				for i := 0; i < 10; i++ {
					in <- i
				}
				close(in) //此处的close使得累加goroutine结束range
			}()
			out := fun(fun(fun(in)))
			for i := range out { //最后在进行调用,每次都是一个数据传送
				fmt.Println(i)
			}
		}()
	}
}

func randNum(done chan int) chan int {
	ch := make(chan int, 10) //生产chan
	go func() {
	Lable:
		for {
			select { //多路监听
			case ch <- rand.Int():
			case <-done: //如果此通道关闭会通知此处执行下面的操作
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}
func generateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func generateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func generateInt() chan int { //使用Fan in增加生成的随机源
	ch := make(chan int, 10)
	go func() {
		for {
			select {
			case ch <- <-generateIntA():
			case ch <- <-generateIntB():
			}
		}
	}()
	return ch
}

func randNumA(done chan int) chan int {
	fmt.Println("进入NumA")
	var mutex sync.Mutex
	ch := make(chan int,10)
	go func() {
	Lable:
		for {
			rand.Seed(time.Now().UnixNano())
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		mutex.Lock()
		close(ch)
		fmt.Println("randNumA关闭")
		mutex.Unlock()
	}()
	return ch
}

func randNumB(done chan int) chan int {
	fmt.Println("进入NumB")
	var mutex sync.Mutex
	ch := make(chan int,10)
	go func() {
	Lable:
		for {
			rand.Seed(time.Now().UnixNano())
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		mutex.Lock()
		close(ch)
		fmt.Println("randNumB关闭")
		mutex.Unlock()
	}()
	return ch
}

func randNumInt(done chan int) chan int { //扇入
	fmt.Println("进入randNumInt")
	var mutex sync.Mutex
	ch := make(chan int)
	send := make(chan int) //控制其他通道关闭
	go func() {
	Lable:
		for { //一直for循环是为了一直监听着done,不然执行一次就结束了,缺点每次都会开一个新通道
			select {
			case ch <- <-randNumA(send): //都会先求值
			case ch <- <-randNumB(send): //都会先求值
			case <-done:
				close(send)
				break Lable
			}
		}
		mutex.Lock()
		close(ch)
		fmt.Println("randNumINt关闭")
		mutex.Unlock()
	}()
	return ch
}
