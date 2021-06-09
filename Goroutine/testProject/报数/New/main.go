package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NumSize  = 10
	M        = 5  //M轮游戏
	N        = 10 //N个人
	GameOver = 0
	GameNext = 1
)

func main() {
	var ReadyGame, StartGame = sync.WaitGroup{}, sync.WaitGroup{}
	ReadyGame.Add(N)             //N个人准备
	var ch = make(chan chan int) //传送手机的chan
	for i := 0; i < N; i++ {     //N个人M轮报数
		go func() {
			phone := make(chan int)
			for j := 0; j < M; j++ {
				ReadyGame.Done()                           //准备完成
				ch <- phone                                //先传手机过去
				phone <- RandNum()                         //再往手机里传数据
				if result := <-phone; result == GameOver { //判断自己是否被淘汰
					break
				}
				StartGame.Wait() //等待新一轮
			}
		}()
	}
	StartGame.Add(1)         //让所有人等待这一轮结束
	cnt := N                 // 确定当前剩余人数
	for i := 0; i < M; i++ { //裁判
		key := RandNum()
		tempCnt := cnt
		for j := 0; j < tempCnt; j++ {
			phone := <-ch
			if result := <-phone; result == key {
				phone <- GameOver
				cnt--
			} else {
				phone <- GameNext
			}
		}
		if i == M-1 { //最后一轮完了就不用再准备下一轮的东西了
			break
		}
		ReadyGame.Add(cnt) //准备人数
		StartGame.Done()   //开始下一轮
		ReadyGame.Wait()   //等待所有人准备完成
		StartGame.Add(1)   //让所有人完成后等待
	}
	fmt.Println(N, "个人,", M, "轮游戏后,", "剩余", cnt, "个人")
}

// RandNum 生成随机数
func RandNum() int {
	time.Sleep(15 * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(NumSize)
}
