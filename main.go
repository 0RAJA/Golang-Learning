package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fun := func(nums ...int) {
		fmt.Println(nums)
		nums[0] = 4
	}
	fun(nums[0:2]...)
	fmt.Println(nums)
}
