package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isClose(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isClose(ctx) {
					break
				}
				time.Sleep(2 * time.Second)
			}
			fmt.Println(i, "cancel!")
		}(i, ctx)
	}
	cancel() // 关闭
	time.Sleep(2 * time.Second)
}
