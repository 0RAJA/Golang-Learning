package main

import (
	"fmt"
	"time"
)

/*
主要就是定时器，标准库中的Timer让用户可以定义自己的超时逻辑，尤其是在应对select处理多个channel的超时、单channel读写的超时等情形时尤为方便。
Timer是一次性的时间触发事件，这点与Ticker不同，Ticker是按一定时间间隔持续触发时间事件。
简介:
type Timer struct {
    C <-chan Time
    // 内含隐藏或非导出字段
}
	Timer类型代表单次时间事件。当Timer到期时，当时的时间会被发送给C，除非Timer是被AfterFunc函数创建的。
func NewTimer(d Duration) *Timer
	NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。
func (t *Timer) Stop() bool
	Stop停止Timer的执行。如果停止了t会返回真；如果t已经被停止或者过期了会返回假。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。
func After(d Duration) <-chan Time
	After会在另一线程经过时间段d后向返回值发送当时的时间的通道。等价于NewTimer(d).C。
*/

func main() {
	if false { //创建计时器
		func() {
			myChan := make(chan int)
			timer := time.NewTimer(3 * time.Second) //3秒后触发
			fmt.Println(time.Now())
			//阻塞3秒
			go func() {
				time.Sleep(4 * time.Second)
				flag := timer.Stop() //停止计时器,如果成功返回true,否则返回false
				if flag {
					fmt.Println("timer停止...")
				} else {
					ch := timer.C
					fmt.Println(<-ch)
				}
				myChan <- 0
			}()
			<-myChan
		}()
	}
	if true {
		func() {
			myChan := make(chan int)
			ch := time.After(3 * time.Second)
			fmt.Println(time.Now())
			go func() {
				fmt.Println(<-ch)
				myChan <- 0
			}()
			<-myChan
		}()
	}
}
