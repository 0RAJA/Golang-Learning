package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumNum(10))
	fmt.Println(sumNum2(10, 20))
	fmt.Println(getAdd(1, 2, 3, 4))
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(getAdd(list...)) //变参传元素加 ...
	fmt.Println(test(list))
}

func sumNum(n int) (sum int) {
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func sumNum2(n1, n2 int) (sum int) {
	for i := n1; i <= n2; i++ {
		sum += i
	}
	return sum
}

//边长参数相当于切片
func getAdd(arr ...int) (sum int) {
	for _, i := range arr {
		sum += i
	}
	return sum
}

//传进去一个切片
func test(arr []int) (sum int) {
	for _, i := range arr {
		sum += i
	}
	return sum
}
