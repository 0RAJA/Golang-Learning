package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup
	//var mutex sync.Mutex
	//十个协程数量
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			//1万叠加
			for j := 0; j < 10000; j++ {
				//mutex.Lock()
				count++//可能会同时给原来的值+1
				//mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
