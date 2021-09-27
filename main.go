package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDecodings("**"))
}

const mod = 1e9 + 7

func numDecodings(s string) int {
	mapCnt := make(map[string]int)
	for i := 1; i <= 26; i++ {
		mapCnt[strconv.Itoa(i)] = 1
	}
	mapCnt["*"] = 9
	mapCnt["1*"] = 9
	mapCnt["2*"] = 6
	mapCnt["**"] = 15
	fx := make([]int, len(s)+1)
	fx[0] = 1
	for i := 1; i <= len(s); i++ {
		if i == 1 {
			fx[i] = mapCnt[s[i-1:i]] % mod
		} else {
			var num1 int
			if s[i-2] == '*' && s[i-1] != '*' {
				if s[i-1] <= '6' {
					num1 = 2
				} else {
					num1 = 1
				}
			} else {
				num1 = mapCnt[s[i-2:i]]
			}
			num2 := mapCnt[s[i-1:i]]
			fx[i] = (num1*fx[i-2] + num2*fx[i-1]) % mod
		}
	}
	return fx[len(s)]
}
