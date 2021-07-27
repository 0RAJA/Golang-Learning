package main

import (
	"fmt"
	"time"
)

//广播退出
func main() {
	cancelCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		i := i
		go func(cancelCh chan struct{}) {
			for {
				if IsClose(cancelCh) {
					break
				}
				time.Sleep(2 * time.Second)
			}
			fmt.Println(i, "cancel!")
		}(cancelCh)
	}
	time.Sleep(1 * time.Second)
	CancelChan(cancelCh)
	time.Sleep(2 * time.Second)
}

func CancelChan(cancelCh chan struct{}) {
	close(cancelCh)
}

func IsClose(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}
