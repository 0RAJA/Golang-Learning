//给你一个大小为 m x n 的网格和一个球。球的起始坐标为 [startRow, startColumn] 。你可以将球移到在四个方向上相邻的单元格内（可以
//穿过网格边界到达网格之外）。你 最多 可以移动 maxMove 次球。
//
// 给你五个整数 m、n、maxMove、startRow 以及 startColumn ，找出并返回可以将球移出边界的路径数量。因为答案可能非常大，返回对
//109 + 7 取余 后的结果。
//
//
//
// 示例 1：
//
//
//输入：m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
//输出：6
//
//
// 示例 2：
//
//
//输入：m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
//输出：12
//
//
//
//
// 提示：
//
//
// 1 <= m, n <= 50
// 0 <= maxMove <= 50
// 0 <= startRow < m
// 0 <= startColumn < n
//
// Related Topics 动态规划
// 👍 150 👎 0
package main

import "fmt"

func main() {
	fmt.Println(findPaths(1, 3, 3, 0, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
const Mod = int(1e9 + 7)

var Next = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var Done = [51][51][51]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < maxMove+1; k++ {
				Done[i][j][k] = -1
			}
		}
	}
	var dfs func(r, c, k int) int
	dfs = func(r, c, k int) (ret int) {
		if r < 0 || c < 0 || r >= m || c >= n {
			return 1
		}
		if k <= 0 {
			return 0
		}
		if Done[r][c][k] != -1 {
			return Done[r][c][k]
		}
		for i := 0; i < len(Next); i++ {
			ret += dfs(r+Next[i][0], c+Next[i][1], k-1) % Mod
		}
		Done[r][c][k] = ret
		return ret
	}
	return dfs(startRow, startColumn, maxMove) % Mod
}

//leetcode submit region end(Prohibit modification and deletion)
