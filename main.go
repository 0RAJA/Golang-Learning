package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	//buildTree(preorder, inorder)
	//fmt.Println(isNumber("0.8"))
	//str := "1^1^1^"
	//fmt.Println(len(strings.Split(str, "^")))
	fmt.Println(fib(3))
}

func isDecimals(s string) bool {
	if s == "" {
		return false
	}
	if num := strings.Count(s, "+") + strings.Count(s, "-"); num != 0 {
		if num > 1 || s[0] != '+' && s[0] != '-' {
			return false
		}
		s = s[1:]
	}
	if len(s) > 0 && strings.Count(s, ".") == 1 {
		var ok error
		index := strings.Index(s, ".")
		if index == 0 {
			_, ok = strconv.Atoi(s[1:])
		} else if index == len(s)-1 {
			_, ok = strconv.Atoi(s[:len(s)-1])
		} else {
			_, ok1 := strconv.Atoi(s[:index])
			_, ok2 := strconv.Atoi(s[index+1:])
			if ok1 == nil && ok2 == nil {
				ok = nil
			} else {
				ok = errors.New(" ")
			}
		}
		if ok == nil {
			return true
		}
	}
	return false
}
func isInteger(s string) bool {
	if s == "" {
		return false
	}
	if num := strings.Count(s, "+") + strings.Count(s, "-"); num != 0 {
		if num > 1 || s[0] != '+' && s[0] != '-' {
			return false
		}
		s = s[1:]
	}
	if len(s) > 0 && strings.Count(s, ".") == 0 {
		_, ok := strconv.Atoi(s)
		if ok == nil {
			return true
		}
	}
	return false
}
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if num := strings.Count(s, "e") + strings.Count(s, "E"); num == 0 || num == 1 {
		index := strings.Index(s, "e")
		if index == -1 {
			index = strings.Index(s, "E")
		}
		if index == -1 {
			return isInteger(s) || isDecimals(s)
		} else if index > 0 && index < len(s)-1 {
			return (isInteger(s[:index]) || isDecimals(s[:index])) && isInteger(s[index+1:])
		}
	}
	return false
}

func fib(n int) int {
	nums := make([]int, 101)
	nums[1], nums[2] = 1, 1
	for i := 3; i <= 100; i++ {
		nums[i] = (nums[i-1]+nums[i-2])%1e9 + 7
	}
	return nums[n]
}
