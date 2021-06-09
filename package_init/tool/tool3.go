package tool

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("tool3的init函数...1")
}
func init() {
	fmt.Println("tool3的init函数...2")
}
func MyTime() {
	fmt.Println(time.Now())
}
