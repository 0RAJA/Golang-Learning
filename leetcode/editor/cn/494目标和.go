//给你一个整数数组 nums 和一个整数 target 。 
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ： 
//
// 
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。 
// 
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], target = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 20 
// 0 <= nums[i] <= 1000 
// 0 <= sum(nums[i]) <= 1000 
// -1000 <= target <= 1000 
// 
// Related Topics 深度优先搜索 动态规划 
// 👍 804 👎 0
package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWaysTwo(nums, target))
}

//leetcode submit region begin(Prohibit modification and deletion)

//DFS
func findTargetSumWays(nums []int, target int) int {
	var dfs func(int, int) int
	dfs = func(sum, cnt int) (all int) {
		if cnt == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}
		return dfs(sum+nums[cnt], cnt+1) + dfs(sum-nums[cnt], cnt+1)
	}
	return dfs(0, 0)
}

//DFS记忆化搜索
func findTargetSumWaysTwo(nums []int, target int) int {
	value := make(map[string]int)
	var dfs func(int, int) int
	dfs = func(sum, cnt int) (all int) {
		key := strconv.Itoa(sum) + "_" + strconv.Itoa(cnt)
		if v, ok := value[key]; ok == true {
			return v
		}
		if cnt == len(nums) {
			if sum == target {
				value[key] = sum + 1
			} else {
				value[key] = sum
			}
			return value[key]
		}
		value[key] = dfs(sum+nums[cnt], cnt+1) + dfs(sum-nums[cnt], cnt+1)
		return value[key]
	}
	return dfs(0, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
