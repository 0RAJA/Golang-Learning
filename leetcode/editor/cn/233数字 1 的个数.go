//给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
//
//
//
// 示例 1：
//
//
//输入：n = 13
//输出：6
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 2 * 109
//
// Related Topics 递归 数学 动态规划
// 👍 239 👎 0
package main

import "fmt"

func main() {
	for x := 0; x <= 10000; x += 10 {
		fmt.Println(x, countDigitOne(x))
	}

	//fmt.Println(countDigitOne(3000))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countDigitOne(n int) int {
	index, num := 0, 0
	for ; index <= n; index++ {
		num += isOne(index)
	}
	return num
}
func isOne(n int) (ret int) {
	for n > 0 {
		if n%10 == 1 {
			ret++
		}
		n /= 10
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
