package main

import (
	"fmt"
)

func main() {
	stones := []int{0, 1, 2, 3, 4, 8, 9, 11}
	fmt.Println(canCross(stones))
}

func canCross(stones []int) bool {
	ln := len(stones)
	// dp定义 dp[i][k]bool  i是目标位置，k是跳跃距离 bool是否能抵达
	dp := make([][]bool, ln)
	for i := range dp {
		dp[i] = make([]bool, ln)
	}
	dp[0][0] = true
	for i := 1; i < ln; i++ {
		// 两个石头距离大于当前所在石子位置的索引值+1 则不可能跳的到
		// 当前位置i-1，目标位置i 如果 stones[i]-stones[i-1]>i-1+1 必然不能跳过去
		if stones[i]-stones[i-1] > i {
			return false
		}
	}

	for i := 1; i < ln; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			// 离i相对近的点j，都无法跳到i，那么j点之前的点也必然跳不过去 每个j点最大跳跃距离就是j+1
			if k > j+1 {
				break
			}
			// 到i点的k距离 是从 前面某个可跳点j 跳了k+1/k/k-1步过来的
			dp[i][k] = dp[j][k-1] || dp[j][k+1] || dp[j][k]
			// 抵达最终点
			if i == ln-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}
