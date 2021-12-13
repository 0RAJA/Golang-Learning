//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个
//整数，判断数组中是否含有该整数。
//
//
//
// 示例:
//
// 现有矩阵 matrix 如下：
//
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//
//
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false。
//
//
//
// 限制：
//
// 0 <= n <= 1000
//
// 0 <= m <= 1000
//
//
//
// 注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// Related Topics 数组 二分查找 分治 矩阵 👍 513 👎 0

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(findNumberIn2DArray(matrix, 5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	for r, c := 0, len(matrix[0])-1; r < len(matrix) && c >= 0; {
		if k := matrix[r][c]; k == target {
			return true
		} else if k < target {
			r++
		} else {
			c--
		}
	}
	return false

}

//leetcode submit region end(Prohibit modification and deletion)
