package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>C语言中文网 | Go语言入门教程</title>
</head>
<body>
    <div>Go语言简介</div>
    <div>Go语言基本语法
    Go语言变量的声明
    Go语言教程简明版
    </div>
    <div>Go语言容器</div>
    <div>Go语言函数</div>
</body>
</html>
    `
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		return
	}
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(len(result), result)
}
