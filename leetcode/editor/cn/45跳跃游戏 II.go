//给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 假设你总是可以到达数组的最后一个位置。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 示例 2:
//
//
//输入: nums = [2,3,0,1,4]
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
//
// Related Topics 贪心 数组 动态规划 👍 1266 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(jump(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) (cnt int) {
	//好好思考下
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	end := 0
	maxPoint := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPoint = max(maxPoint, i+nums[i]) //找到当前路过点能到达的最大位置
		if i == end {                       //到达当前的目的地后,将之前走过的点能到达的最大位置更新为新的目的地(如果没到达新的目的地就到最后一个格子了说明能够通过上一次直接到),且更新下step+1(最开始目的地为0)
			end = maxPoint
			step++
		}
	}
	return step
}

//贪心
func jump2(nums []int) (ret int) {
	n := len(nums) - 1
	idx := 0
	search := func(idx int) (res int) {
		max := 0
		k := nums[idx]
		for i := 1; i <= k; i++ {
			point := idx + i
			if point == n {
				return point
			}
			if m := point + nums[point]; m > max {
				max = m
				res = point
			}
		}
		return
	}
	for idx != n {
		idx = search(idx)
		ret++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
