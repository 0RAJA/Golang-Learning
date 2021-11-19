//给定一个正整数 n ，你可以做如下操作：
//
//
// 如果 n 是偶数，则用 n / 2替换 n 。
// 如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
//
//
// n 变为 1 所需的最小替换次数是多少？
//
//
//
// 示例 1：
//
//
//输入：n = 8
//输出：3
//解释：8 -> 4 -> 2 -> 1
//
//
// 示例 2：
//
//
//输入：n = 7
//输出：4
//解释：7 -> 8 -> 4 -> 2 -> 1
//或 7 -> 6 -> 3 -> 2 -> 1
//
//
// 示例 3：
//
//
//输入：n = 4
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 贪心 位运算 记忆化搜索 动态规划 👍 164 👎 0

package main

import "fmt"

func main() {
	fmt.Println(integerReplacement(2147483647))
}

//leetcode submit region begin(Prohibit modification and deletion)
func integerReplacement(n int) (ret int) {
	memory := make(map[int]int)
	isPower2 := func(a int) bool {
		return a&(a-1) == 0
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var find func(int) int
	find = func(m int) (ret int) {
		if m == 1 {
			return 0
		}
		if v, ok := memory[m]; ok {
			return v
		}
		defer func() {
			memory[m] = ret
		}()
		if m%2 == 0 {
			ret = find(m/2) + 1
		} else if isPower2(m - 1) {
			ret = find(m/2) + 2
		} else {
			ret = min(find(m/2), find(m/2+1)) + 2
		}
		return
	}
	return find(n)
}

//leetcode submit region end(Prohibit modification and deletion)
