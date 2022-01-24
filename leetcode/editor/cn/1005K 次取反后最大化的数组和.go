//给你一个整数数组 nums 和一个整数 K ，按以下方法修改该数组：
//
//
// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//
//
// 重复这个过程恰好 K 次。可以多次选择同一个下标 i 。
//
// 以这种方式修改数组后，返回数组 可能的最大和 。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,2,3], K = 1
//输出：5
//解释：选择下标 1 ，nums 变为 [4,-2,3] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,-1,0,2], K = 3
//输出：6
//解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
//
//
// 示例 3：
//
//
//输入：nums = [2,-3,-1,5,-4], K = 2
//输出：13
//解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -100 <= nums[i] <= 100
// 1 <= K <= 10⁴
//
// Related Topics 贪心 数组 排序 👍 127 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestSumAfterKNegations(nums []int, k int) (sum int) {
	min := math.MaxInt32
	sort.Ints(nums)
	for i := range nums {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
		if nums[i] < min {
			min = nums[i]
		}
		sum += nums[i]
	}
	if k == 0 || k%2 == 0 {
		return sum
	} else {
		sum -= min * 2
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
