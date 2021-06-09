package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1. channel通道是引用类型，需要使用 make 进行创建
非缓冲通道: make(chan T)
	一次发送，一次接收，都是阻塞的
缓冲通道: make(chan T , capacity)
	发送:缓冲区的数据满了，才会阻塞
	接收:缓冲区的数据空了，才会阻塞
双向:
	chan T
		chan <- data, 发送数据，写出
		data <- chan, 获取数据，读取
单向:定向
	chan <- T,只支持写
	<- chan T,只读
2.注意:
panic:
	(1)向已经关闭的通道写数据会导致panic。
		最佳实践是由写入者关闭通道，能最大程度地避免向已经关闭的通道写数据而导致的panic.
	(2)重复关闭的通道会导致panic.
阻塞:
	(1)向未初始化的通道写数据或读数据都会导致当前goroutine的永久阻塞。
	(2)向缓冲区已满的通道写入数据会导致goroutine阻塞。
	(3)通道中没有数据，读取该通道会导致goroutine阻塞。
非阻塞:
	(1)读取已经关闭的通道不会引发阻塞，而是立即返回通道元素类型的零值，可以使用comma, ok语法判断通道是否已经关闭。
	(2)向有缓冲且没有满的通道读/写不会引发阻塞。
*/
var wg sync.WaitGroup
var ch1 = make(chan int) //无缓冲通道
var ch2 = make(chan int)

func main() {
	if false { //通道的使用
		wg.Add(2)
		func() {
			var mutex sync.Mutex
			go func() {
				defer wg.Done()
				for i := 0; i < 5; i++ {
					mutex.Lock()
					fmt.Println("写入", i) //写入数据
					ch1 <- i
					time.Sleep(time.Second)
					mutex.Unlock()
				}
			}()
			for {
				data := <-ch1 //读出数据
				fmt.Println("读出", data)
				if data == 4 {
					break
				}
			}
		}()
		wg.Wait()
	}
	if false { //channel的类型
		func() {
			var ch1 = make(chan int)
			fmt.Printf("%T %v\n", ch1, ch1)
		}()
	}
	if false { //只读通道和只写通道
		func() {
			ch := make(chan int)      //创造一个双向通道
			type Send = chan<- int    //只写通道
			type Receive = <-chan int //只读通道
			var (
				sender   Send    = ch //只写
				receiver Receive = ch //只读
			)
			wg.Add(2)
			go sendChan(sender, 10)
			go receiveChan(receiver)
		}()
		wg.Wait()
	}
	if false { //for循环的for range形式可用于从通道接收值，直到它关闭为止。
		// 记得 close 信道
		// 不然主函数中遍历完并不会结束，而是会阻塞.
		wg.Add(2)
		func() {
			ch := make(chan int)
			go addChan(ch)
			go rangeChan(ch)
		}()
		wg.Wait()
	}
	if false { //非阻塞式接收数据--可以判断通道是否关闭
		wg.Add(2)
		func() {
			var ch = make(chan int)
			go addChan(ch)
			go func() {
				defer wg.Done()
				for {
					data, ok := <-ch
					if !ok {
						fmt.Println(data, ok)
						break
					}
					fmt.Println(data, ok)
				}
			}()
		}()
		wg.Wait()
	}
	if false { //缓冲通道的读取--队列式读取
		func() {
			var mutex sync.Mutex
			var ch = make(chan int, 5)
			wg.Add(2)
			go func() { //写入
				defer wg.Done()
				mutex.Lock()
				for i := 0; i < 5; i++ {
					ch <- i
				}
				close(ch)
				fmt.Println("写入完毕")
				mutex.Unlock()
			}()
			go func() { //读取
				defer wg.Done()
				//mutex.Lock()//这里不能再次上锁
				for data := range ch {
					fmt.Println(data, len(ch))
				}
				fmt.Println("读取完毕")
				//mutex.Unlock()
			}()
			wg.Wait()
		}()
	}
	if false { //交替打印数字和字母第一种
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				ch1 <- 0
				<-ch2
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				<-ch1
				fmt.Printf("%c\n", i+'a')
				ch2 <- 0
			}
		}()
		wg.Wait()
	}
	if false { //交替打印数字和字母第二种
		wg.Add(2)
		go printNum(ch1, ch2)
		go printAlpha(ch1, ch2)
		wg.Wait()
	}
	if false { //计算数字各位的平方和和立方和的和
		wg.Add(2)
		func() {
			var num = 123
			go sumThird(num)
			go sumDouble(num)
			n1, n2 := <-ch2, <-ch1
			fmt.Println(n1, n2)
		}()
		wg.Wait()
	}
}

func printNum(ch, ch2 chan int) { //打印数字
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ch2 //ch2输入
		fmt.Println(i)
		ch <- 0 //ch输出
	}
}

func printAlpha(ch, ch2 chan int) { //打印字母
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch2 <- 0 //ch2输出
		<-ch     //ch输入
		fmt.Printf("%c\n", i+'A')
	}
}

func sumDouble(n int) {
	defer wg.Done()
	sum := 0
	for n != 0 {
		x := n % 10
		sum += x * x
		n /= 10
	}
	ch1 <- sum
}

func sumThird(n int) {
	defer wg.Done()
	sum := 0
	for n != 0 {
		x := n % 10
		sum += x * x * x
		n /= 10
	}
	ch2 <- sum
}

func sendChan(sender chan<- int, n int) {
	defer wg.Done()
	sender <- n
}

func receiveChan(receiver <-chan int) {
	defer wg.Done()
	fmt.Println(<-receiver)
}

func addChan(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //写入方记得关闭通道
	fmt.Println("add结束")
}

func rangeChan(ch chan int) {
	defer wg.Done()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("range结束")
}