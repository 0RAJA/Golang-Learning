package main

import (
	"fmt"
	"strings"
)

//测试键盘输入的反斜杠能否被识别->可以
func main() {
	var name string
	fmt.Scan(&name)
	fmt.Println(strings.ContainsAny(name, "\\/,"))
	//fmt.Println(name)
}
