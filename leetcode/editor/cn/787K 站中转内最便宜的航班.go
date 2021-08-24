//有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城
//市 fromi 开始，以价格 pricei 抵达 toi。
//
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便
//宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
//
//
// 示例 1：
//
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
//输出: 200
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
//
// 示例 2：
//
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 0
//输出: 500
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 10⁴
// 航班没有重复，且不存在自环
// 0 <= src, dst, k < n
// src != dst
//
// Related Topics 深度优先搜索 广度优先搜索 图 动态规划 最短路 堆（优先队列） 👍 333 👎 0
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 3
	edges := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 1
	fmt.Println(findCheapestPrice(n, edges, src, dst, k))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if n < 1 {
		return -1
	}
	matrix := make([][]int, n)
	save := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		save[i] = make([]int, K+1)
	}

	for i := 0; i < len(flights); i++ {
		s, e, v := flights[i][0], flights[i][1], flights[i][2]
		matrix[s][e] = v
	}

	res := math.MaxInt32
	queue := make([]int, 0, 2*n)
	queue = append(queue, src)
	step := 0

	for len(queue) > 0 && step <= K {
		preLen := len(queue)
		for j := 0; j < preLen; j++ {
			s := queue[0]
			queue = queue[1:]
			for i := 0; i < n; i++ {
				if matrix[s][i] <= 0 {
					continue
				}
				if step == 0 {
					save[i][step] = 0 + matrix[s][i]
				} else {
					s := save[s][step-1] + matrix[s][i]
					if s < save[i][step] || save[i][step] == 0 {
						save[i][step] = s
					} else {
						continue
					}
				}
				if save[i][step] > res {
					continue
				}

				if i == dst {
					res = save[i][step]
				}
				queue = append(queue, i)
			}
		}
		step++
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
