//给你一个整数数组 nums 和一个整数 K ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
//
//
// 子数组大小 至少为 2 ，且
// 子数组元素总和为 K 的倍数。
//
//
// 如果存在，返回 true ；否则，返回 false 。
//
// 如果存在一个整数 n ，令整数 x 符合 x = n * K ，则称 x 是 K 的一个倍数。
//
//
//
// 示例 1：
//
//
//输入：nums = [23,2,4,6,7], K = 6
//输出：true
//解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。
//
// 示例 2：
//
//
//输入：nums = [23,2,6,4,7], K = 6
//输出：true
//解释：[23, 2, 6, 4, 7] 是大小为 5 的子数组，并且和为 42 。
//42 是 6 的倍数，因为 42 = 7 * 6 且 7 是一个整数。
//
//
// 示例 3：
//
//
//输入：nums = [23,2,6,4,7], K = 13
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// 0 <= nums[i] <= 109
// 0 <= sum(nums[i]) <= 231 - 1
// 1 <= K <= 231 - 1
//
// Related Topics 数学 动态规划
// 👍 290 👎 0

package main

func main() {

}

/*
sum[j]−sum[i−1]=n∗K
经过简单的变形可得：
sum[j]/K - sum[i-1]/K = n(整数)
即:sum[j]%K == sum[i]%K
*/
//leetcode submit region begin(Prohibit modification and deletion)
func checkSubarraySum(nums []int, k int) bool {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	myMap := make(map[int]bool)
	for i := 2; i <= len(nums); i++ {
		myMap[preSum[i-2]%k] = true
		if myMap[preSum[i]%k] == true {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
