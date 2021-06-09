package main

import (
	"fmt"
	"time"
)

/*
panic:
1.内建函数
2.假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表,按照defer的逆序执行
3.返回函数F的调用者G,在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表,
	按照defer的逆序执行，这里的defer 有点类似try-catch-finally中的finally
4、直到goroutine整个退出,并报告错误
recover:
1、内建函数
2、用来控制一个goroutine的panicking行为，捕获panic, 从而影响应用的行为
3、一般的调用建议
	a).在defer函数中，通过recover来终 止一个goroutine的panicking过程, 从而恢复正常代码的执行
	b).可以获取通过panic传递的error

简单来讲: go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
*/
func main() {
	start := time.Now()
	funA()
	defer fmt.Println("main 3...")
	funB()
	defer fmt.Println("main 4...")
	fmt.Println("main over")
	fmt.Println("用时:", time.Since(start))
}
func funA() {
	fmt.Println("funA")
}
func funB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "恢复程序")
		}
	}()
	fmt.Println("funB")
	defer fmt.Println("funB 1...")
	num := [5]int{}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", num[i])
		//os.Exit(0)
	}
	fmt.Println()
	defer fmt.Println("funB 2...")
	fmt.Println("funB over")
}
