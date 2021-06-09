package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func shipWithinDays(weights []int, D int) int {
	var left, right = 0, 0
	for _, w := range weights {
		left = max(left, w)
		right += w
	}
	for left < right {
		mid := left + (right-left)/2
		data, cost := 1, 0
		for _, w := range weights {
			if cost+w > mid {
				data++
				cost = 0
			}
			cost += w
		}
		if data <= D {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(shipWithinDays(num, 5))
}
