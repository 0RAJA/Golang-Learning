package main

import (
	"fmt"
	"math/rand"
	"time"
)

var tickets = 100

func main() {
	go sale("n1")
	go sale("n2")
	go sale("n3")
	go sale("n4")
	time.Sleep(3 * time.Second)
}
func sale(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if tickets <= 0 {
			fmt.Println("票售完了")
			break
		}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Nanosecond)
		tickets--
		fmt.Println(name, tickets)
	}
}
