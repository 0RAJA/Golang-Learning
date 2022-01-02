//列表 arr 由在范围 [1, n] 中的所有整数组成，并按严格递增排序。请你对 arr 应用下述算法：
//
//
//
//
// 从左到右，删除第一个数字，然后每隔一个数字删除一个，直到到达列表末尾。
// 重复上面的步骤，但这次是从右到左。也就是，删除最右侧的数字，然后剩下的数字每隔一个删除一个。
// 不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
//
//
// 给你整数 n ，返回 arr 最后剩下的数字。
//
//
//
// 示例 1：
//
//
//输入：n = 9
//输出：6
//解释：
//arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
//arr = [2, 4, 6, 8]
//arr = [2, 6]
//arr = [6]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
//
//
//
// Related Topics 数学 👍 194 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func lastRemaining(n int) int {
	l, r := 1, n
	k, cnt, step := 0, n, 1
	for cnt > 0 {
		if k%2 == 0 { //从左往右
			l += step
			if cnt%2 == 1 { //奇数
				r -= step
			}
		} else { //从右往左
			r -= step
			if cnt%2 == 1 {
				l += step
			}
		}
		k++       //到下一轮
		step *= 2 //步长*2
		cnt /= 2  //数量/2
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
