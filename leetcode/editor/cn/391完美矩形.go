//给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是
// (xi, yi) ，右上顶点是 (ai, bi) 。
//
// 如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
//
//
// 示例 1：
//
//
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
//输出：true
//解释：5 个矩形一起可以精确地覆盖一个矩形区域。
//
//
// 示例 2：
//
//
//输入：rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
//输出：false
//解释：两个矩形之间有间隔，无法覆盖成一个矩形。
//
// 示例 3：
//
//
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]
//输出：false
//解释：图形顶端留有空缺，无法覆盖成一个矩形。
//
// 示例 4：
//
//
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
//输出：false
//解释：因为中间有相交区域，虽然形成了矩形，但不是精确覆盖。
//
//
//
// 提示：
//
//
// 1 <= rectangles.length <= 2 * 10⁴
// rectangles[i].length == 4
// -10⁵ <= xi, yi, ai, bi <= 10⁵
//
// Related Topics 数组 扫描线 👍 116 👎 0

package main

import "math"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func isRectangleCover(rectangles [][]int) bool {
	left, right, high, low := math.MaxInt32, math.MinInt32, math.MinInt32, math.MaxInt32
	sum := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	mapCnt := make(map[struct{ x, y int }]bool)
	for i := range rectangles {
		a, b, c, d := rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]
		sum += (c - a) * (d - b)
		left = min(left, b)
		low = min(low, a)
		high = max(high, c)
		right = max(right, d)
		points := [4]struct{ x, y int }{{x: a, y: d}, {x: a, y: b}, {x: c, y: d}, {x: c, y: b}}
		for p := range points {
			if _, ok := mapCnt[points[p]]; ok {
				delete(mapCnt, points[p])
			} else {
				mapCnt[points[p]] = true
			}
		}
	}
	return len(mapCnt) == 4 && sum == (right-left)*(high-low) && mapCnt[struct{ x, y int }{x: low, y: left}] && mapCnt[struct{ x, y int }{x: low, y: right}] && mapCnt[struct{ x, y int }{x: high, y: left}] && mapCnt[struct{ x, y int }{x: high, y: right}]
}

//leetcode submit region end(Prohibit modification and deletion)
