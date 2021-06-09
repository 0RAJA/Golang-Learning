package main

import "fmt"

func main() {
	wall := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}
	fmt.Println(leastBricks(wall))
}

func leastBricks(wall [][]int) int {
	map1 := make(map[int]int)
	ln := len(wall)
	for i := 0; i < ln; i++ {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			map1[sum]++
		}
	}
	max := 0
	for _, value := range map1 {
		if value > max {
			max = value
		}
	}
	return ln - max
}
