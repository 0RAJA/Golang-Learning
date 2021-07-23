//给你一个二维整数数组 ranges 和两个整数 left 和 right 。每个 ranges[i] = [starti, endi] 表示一个从 star
//ti 到 endi 的 闭区间 。 
//
// 如果闭区间 [left, right] 内每个整数都被 ranges 中 至少一个 区间覆盖，那么请你返回 true ，否则返回 false 。 
//
// 已知区间 ranges[i] = [starti, endi] ，如果整数 x 满足 starti <= x <= endi ，那么我们称整数x 被覆盖了
//。 
//
// 
//
// 示例 1： 
//
// 
//输入：ranges = [[1,2],[3,4],[5,6]], left = 2, right = 5
//输出：true
//解释：2 到 5 的每个整数都被覆盖了：
//- 2 被第一个区间覆盖。
//- 3 和 4 被第二个区间覆盖。
//- 5 被第三个区间覆盖。
// 
//
// 示例 2： 
//
// 
//输入：ranges = [[1,10],[10,20]], left = 21, right = 21
//输出：false
//解释：21 没有被任何一个区间覆盖。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= ranges.length <= 50 
// 1 <= starti <= endi <= 50 
// 1 <= left <= right <= 50 
// 
// Related Topics 数组 哈希表 前缀和 
// 👍 21 👎 0
package main

import (
	"fmt"
	"sort"
)

func main() {
	ranges := [][]int{{13, 43}, {19, 20}, {32, 38}, {26, 33}, {3, 38}, {16, 31}, {26, 48}, {27, 43}, {12, 24}}
	left := 36
	right := 45
	fmt.Println(isCovered(ranges, left, right))
}

//leetcode submit region begin(Prohibit modification and deletion)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	low := ranges[0][1]
	l := ranges[0][0]
	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i+1][0] <= low+1 {
			low = max(low, ranges[i+1][1])
		} else {
			if l <= left && low >= right {
				return true
			}
			low = ranges[i+1][1]
			l = ranges[i+1][0]
		}
	}
	return l <= left && low >= right
}

//leetcode submit region end(Prohibit modification and deletion)
