//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
//
// Related Topics 数组 二分查找 矩阵 👍 531 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	findRow := func() int {
		l, r := 0, len(matrix)-1
		for l < r {
			mid := l + (r-l)/2
			if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
				return mid
			} else if matrix[mid][0] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
	row := findRow()
	l, r := 0, len(matrix[row])-1
	for l < r {
		mid := l + (r-l)/2
		if t := matrix[row][mid]; t == target {
			return true
		} else if t < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return matrix[row][l] == target
}

//leetcode submit region end(Prohibit modification and deletion)
