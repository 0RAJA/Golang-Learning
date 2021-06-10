package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/*
1.定义命令行参数:
flag.Type()
	flag.Type(flag名, 默认值, 帮助信息)*Type
flag.TypeVar()
	flag.TypeVar(Type指针变量, flag名, 默认值, 帮助信息)
2.解析命令行参数
flag.Parse()
-flag xxx （使用空格，一个-符号）
--flag xxx （使用空格，两个-符号）
-flag=xxx （使用等号，一个-符号）
--flag=xxx （使用等号，两个-符号）
其中，布尔类型的参数必须使用等号的方式指定
Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。
3.其他函数
flag.Args() //返回命令行参数后的其他参数，以[]string类型
flag.NArg() //返回命令行参数后的其他参数个数
flag.NFlag() //返回使用的命令行参数个数
4.演示
 ./flag -name raja -age 19 -married false -d 0h
raja 19 true 0s
[false -d 0h]
3
3

$ ./flag -help
Usage of D:\WorkSpace\Go\goTest\src\package\flag\flag.exe:
  -age int
        年龄 (default 18)
  -d duration
        延迟的时间间隔
  -married
        婚否
  -name string
        姓名 (default "张三")
*/
func main() {
	if false { //简单的想要获取命令行参数
		if len(os.Args) > 0 {
			for index, arg := range os.Args {
				fmt.Println(index, arg)
			}
		}
	}
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
