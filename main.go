package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 4, 5, 6}
	fmt.Println(sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 7
	}))
}
