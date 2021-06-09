package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	copyTest(mu)
}

//这里复制了一个锁，造成了死锁
//在拷贝锁时会同时拷贝它的状态，mu没有解锁直接又进行加锁操作所以发生死锁了。
func copyTest(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("ok")
}
