package main

import (
	"errors"
	"fmt"
	"os"
)

/*
错误指的是可能出现问题的地方出现了问题。比如打开一个文件时失败，这种情况在人们的意料之中。
而异常指的是不应该出现问题的地方出现了问题。比如引用了空指针，这种情况在人们的意料之外。
可见，错误是业务过程的一部分，而异常不是。
error:内置的数据类型，内置的接口
	定义方法: Error() string
使用go语言提供好的包:
	errors包下的函数: New(), 创建一个error对象
	fmt包下的Errorf()函数:
		func Errorf(format string, a ... interface{}) error
*/
func main() {
	if false {
		f, err := os.Open("test.txt")
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		fmt.Println(f.Name())
	}
	//1.创建error
	err1 := errors.New("这是错误1")
	fmt.Printf("%T %v\n", err1, err1)
	err2 := fmt.Errorf("这是错误%d", 2)
	fmt.Printf("%T %v\n", err2, err2)
	fmt.Println("-------------------------------")
	err3 := checkAge(17)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("Go on")
}
func checkAge(age int) error {
	if age < 18 {
		//return errors.New("年龄不合法")
		return fmt.Errorf("年龄%d是不合法的", age)
	}
	return nil
}
