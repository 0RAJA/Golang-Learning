package main

import "strings"

func addSpaces(s string, spaces []int) string {
	sum := 0
	ret := make([]string, 0, len(spaces))
	for i := range spaces {
		ret = append(ret, s[sum:spaces[i]])
		sum = spaces[i]
	}
	ret = append(ret, s[sum:])
	return strings.Join(ret, " ")
}
