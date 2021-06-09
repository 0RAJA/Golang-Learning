package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(3))
}

func judgeSquareSum(c int) bool {
	for a := int(math.Sqrt(float64(c))); a >= 0; a-- {
		b := int(math.Sqrt(float64(c - a*a)))
		if b*b == c-a*a {
			return true
		}
	}
	return false
}
