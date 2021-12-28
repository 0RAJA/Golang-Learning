//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// Related Topics 位运算 数组 回溯 👍 1426 👎 0

package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) (ret [][]int) {
	result := make([]int, 0, len(nums))
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ret = append(ret, append([]int{}, result...))
			return
		}
		result = append(result, nums[index])
		dfs(index + 1)
		result = result[:len(result)-1]
		dfs(index + 1)
	}
	dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
