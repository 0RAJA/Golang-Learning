package main

import (
	"fmt"
	"math/rand"
	"time"
)

//step1:设置种子数

func main() {
	//时间戳:指定时间,距离1970年1月1日0点0时0分0秒之间的差值,秒
	fmt.Println(time.Now().Unix())     //秒
	fmt.Println(time.Now().UnixNano()) //纳秒
	//设置种子数
	rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	//求随机数--[12,67]--[0,55]+12
	fmt.Println(rand.Intn(56) + 12)
}
