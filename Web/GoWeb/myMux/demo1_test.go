package main

import (
	"fmt"
	"testing"
)


//可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("开始测试:")
	//通过m.Run来执行子测试函数
	m.Run()
}

//测试函数必须以Test开头,否则不执行,但可以设置为一个子测试函数
func TestOne(t *testing.T) {
	fmt.Println("子测试程序One")
}

func TestTwo(t *testing.T) {
	fmt.Println("子测试程序Two")
}
