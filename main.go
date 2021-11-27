package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`(?m)^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	var str = `1647193241@qq.com`
	fmt.Println(re.FindStringSubmatch(str)[0])
}
