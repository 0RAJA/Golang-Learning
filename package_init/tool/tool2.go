package tool

import "fmt"

func MyTime2() {
	MyTime()
}
func init() {
	fmt.Println("tool2的init函数...1")
}
func init() {
	fmt.Println("tool2的init函数...2")
}

type Student struct {
	Name string
	Age  int
}
