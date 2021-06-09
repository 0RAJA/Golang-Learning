package main

import (
	"fmt"
)

/*
水仙花数:三位数: [100 , 999]
每个位上的数字的立方和，刚好等于该数字本身。那么就叫水仙花数，4个
比如: 153
1*1*1 + 5*5*5 + 3*3*3 = 1+125+27=153
*/

func numberOfDaffodils(n int) (result bool) {
	other, sum := n, 0
	for other != 0 {
		temp := other % 10
		sum += temp * temp * temp
		other /= 10
	}
	return sum == n
}

func main() {
	const (
		low  = 100
		high = 999
	)
	for i := low; i <= high; i++ {
		if numberOfDaffodils(i) {
			fmt.Println(i)
		}
	}
}
