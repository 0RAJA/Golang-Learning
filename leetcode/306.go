package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := "1023"
	fmt.Println(isAdditiveNumber(nums))
}

func isAdditiveNumber(num string) (ret bool) {
	if len(num) < 3 {
		return false
	}
	res := make([]int, 0, len(num))
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(num) && len(res) >= 3 {
			ret = true
			return true
		}
		for i := 1; idx+i <= len(num); i++ {
			if i > 1 && num[idx] == '0' {
				return false
			}
			next, _ := strconv.Atoi(num[idx : idx+i])
			if len(res) <= 1 || next == res[len(res)-2]+res[len(res)-1] {
				res = append(res, next)
				if dfs(idx + i) {
					return true
				}
				res = res[:len(res)-1]
			}
		}
		return false
	}
	return dfs(0)
}
