package main

import "fmt"

func main() {
	const (
		n = 10
		m = 10
	)
	var (
		left  = 1
		right = m
		up    = 1
		down  = n
		cnt   = 1
		num   [n + 1][m + 1]int
	)
	for up <= down || left <= right {
		for j := left; j <= right; j++ {
			num[up][j] = cnt
			cnt++
		}
		for j := up + 1; j <= down; j++ {
			num[j][right] = cnt
			cnt++
		}
		for j := right - 1; j >= left; j-- {
			num[down][j] = cnt
			cnt++
		}
		for j := down - 1; j >= up+1; j-- {
			num[j][left] = cnt
			cnt++
		}
		up++
		down--
		left++
		right--
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%4d ", num[i][j])
		}
		fmt.Println()
	}
}
