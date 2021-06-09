package main

import (
	"fmt"
)

/*
switch语句:
switch中case是顺序进行的
语法结构:数值不一定是数字,只要和变量类型相同就行
	switch [赋值;] 变量名 {
	case数值1:
		分支1
	case数值2:
		分支2
	case数值3:
		分支3
	default:
		最后一个分支
	}

	switch [赋值;] {
	case bool类型:
		分支1
	case bool类型:
		分支2
	}
fallthrough 穿透case
只能放在case里的最下面
*/
func main() {
	//1
	if false {
		var num int = 3
		switch num {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		default:
			fmt.Println("None")
		}
	}
	//2-计算器
	if false {
		var (
			num1, num2 int
			opera      byte
			sum        int
		)
		fmt.Println("请输入整数num1和num2")
		fmt.Scanf("%d%d", &num1, &num2)
		fmt.Println("请输入一个操作符 +,-,*,/")
		fmt.Scanf("%c", &opera)
		switch opera {
		case '+':
			sum = num1 + num2
		case '-':
			sum = num1 - num2
		case '*':
			sum = num1 * num2
		case '/':
			sum = num1 / num2
		default:
			fmt.Println("No")
		}
		fmt.Println(sum)
	}
	if true {
		p := 2
		switch a := true; {
		case p == 2:
			fmt.Println(a)
		case p == 3:
			fmt.Println(!a)
		}
	}
}
