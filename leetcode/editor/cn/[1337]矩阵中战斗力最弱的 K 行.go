//给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
//
// 请你返回矩阵中战斗力最弱的 K 行的索引，按从最弱到最强排序。
//
// 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
//
// 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
//
//
//
// 示例 1：
//
//
//输入：mat =
//[[1,1,0,0,0],
// [1,1,1,1,0],
// [1,0,0,0,0],
// [1,1,0,0,0],
// [1,1,1,1,1]],
//K = 3
//输出：[2,0,3]
//解释：
//每行中的军人数目：
//行 0 -> 2
//行 1 -> 4
//行 2 -> 1
//行 3 -> 2
//行 4 -> 5
//从最弱到最强对这些行排序后得到 [2,0,3,1,4]
//
//
// 示例 2：
//
//
//输入：mat =
//[[1,0,0,0],
// [1,1,1,1],
// [1,0,0,0],
// [1,0,0,0]],
//K = 2
//输出：[0,2]
//解释：
//每行中的军人数目：
//行 0 -> 1
//行 1 -> 4
//行 2 -> 1
//行 3 -> 1
//从最弱到最强对这些行排序后得到 [0,2,3,1]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 2 <= n, m <= 100
// 1 <= K <= m
// matrix[i][j] 不是 0 就是 1
//
// Related Topics 数组 二分查找 矩阵 排序 堆（优先队列）
// 👍 59 👎 0
package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type ANum struct {
	row int
	val int
}

func kWeakestRows(mat [][]int, k int) (ret []int) {
	num := make([]ANum, 0)
	for i := 0; i < len(mat); i++ {
		low, high := 0, len(mat[i])-1
		for low+1 < high {
			mid := low + (high-low)/2
			if mat[i][mid] == 0 {
				high = mid
			} else {
				low = mid
			}
		}
		n := ANum{
			row: i,
		}
		if mat[i][high] == 1 {
			n.val = high
		} else if mat[i][low] == 1 {
			n.val = low
		} else {
			n.val = low - 1
		}
		num = append(num, n)
	}
	sort.Slice(num, func(i, j int) bool {
		if num[i].val != num[j].val {
			return num[i].val < num[j].val
		} else {
			return num[i].row < num[j].row
		}
	})
	for i := 0; i < k; i++ {
		ret = append(ret, num[i].row)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
