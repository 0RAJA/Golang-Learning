//给定一个未排序的整数数组，找到最长递增子序列的个数。
//
// 示例 1:
//
//
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//
//
// 示例 2:
//
//
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
//
// 注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。
// Related Topics 树状数组 线段树 数组 动态规划 👍 447 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/*
 dp[i] 表示以 nums[i] 结尾的最长上升子序列的长度，cnt[i] 表示以 nums[i] 结尾的最长上升子序列的个数
*/
func findNumberOfLIS(nums []int) (ret int) {
	maxLen := 0
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	for i, a := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, b := range nums[:i] {
			if b < a {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j] //入乡随俗
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j] //收至麾下
				}
			}
		}
		if maxLen == dp[i] {
			ret += cnt[i]
		} else if dp[i] > maxLen {
			ret = cnt[i]
			maxLen = dp[i]
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
