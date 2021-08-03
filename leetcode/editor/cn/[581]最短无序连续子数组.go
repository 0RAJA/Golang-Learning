//给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。 
//
// 请你找出符合题意的 最短 子数组，并输出它的长度。 
//
// 
//
// 
// 
// 示例 1： 
//
// 
//输入：nums = [2,6,4,8,10,9,15]
//输出：5
//解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,4]
//输出：0
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 104 
// -105 <= nums[i] <= 105 
// 
//
// 
//
// 进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？ 
// 
// 
// Related Topics 栈 贪心 数组 双指针 排序 单调栈 
// 👍 591 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findUnsortedSubarray(nums []int) int {
	prefixMax := make([]int, len(nums))
	suffixMin := make([]int, len(nums))
	prefixMax[0], suffixMin[len(nums)-1] = nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
		j := len(nums) - i - 1
		if nums[j] < suffixMin[j+1] {
			suffixMin[j] = nums[j]
		} else {
			suffixMin[j] = suffixMin[j+1]
		}
	}
	lBool, rBool := true, true
	left, right := 0, 0
	for i := 0; i < len(nums) && (lBool || rBool); i++ {
		if lBool && prefixMax[i] > suffixMin[i] {
			left = i
			lBool = false
		}
		j := len(nums) - i - 1
		if rBool && suffixMin[j] < prefixMax[j] {
			right = j
			rBool = false
		}
	}
	if right == left {
		return 0
	}
	return right - left + 1
}

//leetcode submit region end(Prohibit modification and deletion)
