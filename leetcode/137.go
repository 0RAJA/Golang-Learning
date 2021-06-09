package main

import "fmt"

func main() {
	nums := []int{-1, 0, 0, 0}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	result := [33]int{}
	for _, i := range nums {
		cnt := 0
		for i != 0 && cnt < 32 {
			result[cnt] += i & 1
			i >>= 1
			cnt++
		}
	}
	var key int32 = 0
	for index, value := range result {
		if value%3 == 1 {
			key += 1 << index
		}
	}
	return int(key)
}
