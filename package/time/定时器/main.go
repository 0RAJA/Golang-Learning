package main

import (
	"fmt"
	"time"
)

//一次性定时器 Timer
func main() {
	if false {
		timer := time.NewTimer(time.Second)
	Loop:
		for {
			select {
			case <-timer.C:
				fmt.Println("超时")
				break Loop
			}
		}
	}
	if true {
		timer := time.NewTimer(time.Second)
		timer.Stop() //终止
		s1 := time.Now()
		timer.Reset(time.Second) //重启设置新时间
		<-timer.C
		fmt.Println(time.Now().Sub(s1))
	}
}

//周期性定时器 Ticker
