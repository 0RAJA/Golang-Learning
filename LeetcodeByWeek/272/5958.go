package main

import "fmt"

func main() {
	price := []int{3, 2, 1, 4}
	fmt.Println(getDescentPeriods(price))
}
func getDescentPeriods(prices []int) (ret int64) {
	for i, j := 0, 0; j < len(prices); j++ {
		if j != i && prices[j]-prices[j-1] != -1 {
			ret += int64((1+(j-i))*(j-i)) / 2
			i = j
		}
		if j == len(prices)-1 {
			ret += int64((1+(j-i+1))*(j-i+1)) / 2
		}
	}
	return
}
