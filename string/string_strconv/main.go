package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv 字符串与基本类型之间的转化
	*/
	//1.bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1) //string转bool
	if err != nil {
		fmt.Println(nil)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)
	s1 = strconv.FormatBool(b1) //bool转string
	fmt.Printf("%T,%s\n", s1, s1)
	//2.整数
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)
	s2 = strconv.FormatInt(i2, 10)
	fmt.Printf("%T,%s\n", s2, s2)
	//3.itoa(),atoi()
	i3, err := strconv.Atoi("-10") //转整数
	s3 := strconv.Itoa(100)        //转字符串
	fmt.Printf("%d %T\n%s %T\n\n", i3, i3, s3, s3)

}
