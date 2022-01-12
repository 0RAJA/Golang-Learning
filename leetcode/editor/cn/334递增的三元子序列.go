//给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回
//true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4,5]
//输出：true
//解释：任何 i < j < k 的三元组都满足题意
//
//
// 示例 2：
//
//
//输入：nums = [5,4,3,2,1]
//输出：false
//解释：不存在满足题意的三元组
//
// 示例 3：
//
//
//输入：nums = [2,1,5,0,4,6]
//输出：true
//解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
// Related Topics 贪心 数组 👍 463 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{5, 4, 3, 1, 2, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func increasingTriplet(nums []int) bool {
	leftMin := make([]int, len(nums))
	for i := range nums {
		var t = nums[i]
		if i > 0 && nums[i] > leftMin[i-1] {
			t = leftMin[i-1]
		}
		leftMin[i] = t
	}
	rightMax := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		var t = nums[i]
		if i < len(nums)-1 && nums[i] < rightMax[i+1] {
			t = rightMax[i+1]
		}
		rightMax[i] = t
	}
	for i := range nums {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}
func increasingTriplet2(nums []int) bool {
	min, mid := math.MaxInt32, math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] < mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
