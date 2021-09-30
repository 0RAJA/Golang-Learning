//给你 二维 平面上两个 由直线构成的 矩形，请你计算并返回两个矩形覆盖的总面积。
//
// 每个矩形由其 左下 顶点和 右上 顶点坐标表示：
//
//
//
// 第一个矩形由其左下顶点 (ax1, ay1) 和右上顶点 (ax2, ay2) 定义。
// 第二个矩形由其左下顶点 (bx1, by1) 和右上顶点 (bx2, by2) 定义。
//
//
//
//
//
// 示例 1：
//
//
//输入：ax1 = -3, ay1 = 0, ax2 = 3, ay2 = 4, bx1 = 0, by1 = -1, bx2 = 9, by2 = 2
//输出：45
//
//
// 示例 2：
//
//
//输入：ax1 = -2, ay1 = -2, ax2 = 2, ay2 = 2, bx1 = -2, by1 = -2, bx2 = 2, by2 = 2
//输出：16
//
//
//
//
// 提示：
//
//
// -10⁴ <= ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 <= 10⁴
//
// Related Topics 几何 数学 👍 144 👎 0

package main

import "testing"

func TestRectangleArea(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	//将给定的两个数按从小到大排序
	sort := func(a, b int) (int, int) {
		if a > b {
			return a, b
		}
		return b, a
	}
	//a,b,c,d分别表示重叠区间的下,上,左,右
	a, _ := sort(ax1, bx1)
	_, b := sort(ax2, bx2)
	c, _ := sort(ay1, by1)
	_, d := sort(ay2, by2)
	sum := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	//排除不重叠的情况
	if (b-a) > 0 && (d-c) > 0 {
		sum -= (b - a) * (d - c)
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
