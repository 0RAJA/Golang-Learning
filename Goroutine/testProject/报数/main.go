package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MaxNum = 10 //随机数最大数
	M      = 10 //游戏轮数
	N      = 10 //玩家数
)

//运动员
func person(id int, Survive []bool, key int, Dead chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	num := RandNum()
	if num == key { //判断出自己被淘汰
		Survive[id] = true
		Dead <- true
	}
}

//裁判--如果有人淘汰减少总人数
func center(sum *int, Dead chan bool, over chan bool) {
	for {
		select {
		case <-Dead:
			*sum--
		case <-over:
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	Dead := make(chan bool) //告知裁判有人淘汰
	over := make(chan bool)
	sum := N                    //总人数
	Survive := make([]bool, 15) //不同编号的人是否存活的标志(默认false存活)
	go center(&sum, Dead, over) //开启裁判
	for i := 0; i < M; i++ {
		key := RandNum() //每轮一个随机数传进去让选手自己判断
		for j := 1; j <= N; j++ {
			if Survive[j] == true {
				continue
			}
			wg.Add(1)
			go person(j, Survive, key, Dead, &wg)
		}
		wg.Wait() //等待上一轮所有人完成
	}
	over <- true //结束裁判
	fmt.Println(N, "个人,", M, "轮游戏后,", "剩余", sum, "个人")
}

func RandNum() int {
	time.Sleep(10 * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MaxNum)
}
