package main

import (
	"fmt"
	"unicode"
)

//处理非ASCII码类型的字符,定义了新类型rune--除了ASCII码里面的
//uint8即byte类型,代表ASCII码中的一个字符
//rune类型--int32,代表一个UTF-8字符
func main() {
	s := "你好李焕英122" //len(s) == 18
	//len()求byte字节的数目
	fmt.Println(len(s))

	//判断中文
	cnt := 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			cnt++
		}
	}
	fmt.Println(cnt)
	//字符串修改
	list1 := []rune(s) //把字符串强制转换为一个rune切片
	for i := 0; i < len(list1); i++ {
		fmt.Printf("%c\n", list1[i])
	}
	list1[0] = '我'
	fmt.Printf("%s\n", string(list1)) //把rune切片强制转换为字符串
	s2 := "Hello"
	fmt.Printf("%T %d\n", s2[0], s2[0])

	//类型转换 --整型和浮点型,字符串和切片可以

}
