//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 3957 👎 0

package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) (ret [][]int) {
	target := 0
	n := len(nums)
	sort.Ints(nums)
	for a := 0; a < n; a++ { //a从左往右
		if a > 0 && nums[a] == nums[a-1] { //a不重复
			continue
		}
		for b, c := a+1, n-1; b < c; { //c 从右往左
			if sum := nums[a] + nums[b] + nums[c]; sum == target { //通过和目标值比较决定两个指针移动
				ret = append(ret, []int{nums[a], nums[b], nums[c]})
				for b++; nums[b] == nums[b-1] && b < c; b++ { //因为不重复所以可以两个指针都动
				}
				for c--; nums[c] == nums[c+1] && b < c; c-- {
				}
			} else if sum < target {
				b++
			} else {
				c--
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
