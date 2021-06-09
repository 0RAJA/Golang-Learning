package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	N         = 10 //10个人
	M         = 10 // 10轮游戏
	MaxNumber = 10
	gameOver  = 0
	gameNext  = 1
)

func main() {
	ch := make(chan chan int)
	readyGame := sync.WaitGroup{}
	readyGame.Add(N)
	start := sync.WaitGroup{}
	// N个人
	for i := 0; i < N; i++ {
		go func() {
			//每个人执行M轮游戏
			phone := make(chan int)
			for j := 0; j < M; j++ {
				// 告诉裁判我可以开始游戏了
				readyGame.Done()
				ch <- phone
				phone <- rand.Intn(MaxNumber)
				if val := <-phone; val == gameOver {
					break
				}
				//这一轮游戏结束等待
				// 如果没有这一步的话 他会直接进入到写管道 这个时候他可能提前进入下一轮游戏 这样是不对的
				start.Wait()
			}
		}()
	}
	//裁判
	start.Add(1)
	cnt := N
	for i := 0; i < M; i++ {
		num := rand.Intn(MaxNumber)
		tempCnt := cnt
		for j := 0; j < tempCnt; j++ {
			phone := <-ch
			val := <-phone
			//fmt.Println(val)
			if num == val {
				phone <- gameOver
				cnt--
			} else {
				phone <- gameNext
			}
		}
		if i == M-1 {
			break
		}
		readyGame.Add(cnt)
		start.Done()
		readyGame.Wait()
		start.Add(1)
	}
	fmt.Println(M, " 轮游戏还有", cnt, "个人")
}
