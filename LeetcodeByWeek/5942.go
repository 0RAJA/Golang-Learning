package main

import (
	"fmt"
	"sort"
)

func main() {
	digits := []int{0, 2, 0, 0}
	fmt.Println(findEvenNumbers(digits))
}

func findEvenNumbers(digits []int) (ret []int) {
	iMap := map[int]bool{}
	box := [3]int{}
	var dfs func(step int)
	dfs = func(step int) {
		if step >= 3 {
			x := box[0]*100 + box[1]*10 + box[2]
			if box[0] != 0 && x%2 == 0 {
				ret = append(ret, x)
			}
			return
		}
		for i := 0; i < len(digits); i++ {
			if (i > 0 && digits[i] == digits[i-1] && iMap[i-1] == false) || iMap[i] {
				continue
			}
			iMap[i] = true
			box[step] = digits[i]
			dfs(step + 1)
			iMap[i] = false
		}
	}
	sort.Ints(digits)
	dfs(0)
	return
}
