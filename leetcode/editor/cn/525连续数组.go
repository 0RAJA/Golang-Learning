//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量0和1的最长连续子数组。 
//
// 示例 2: 
//
// 
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 105 
// nums[i] 不是 0 就是 1 
// 
// Related Topics 哈希表 
// 👍 318 👎 0
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 0, 1}
	fmt.Println(findMaxLength(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums []int) int {
	max := 0
	preSum := make([]int, len(nums)+1)
	myMap := make(map[int]int) //key:sum value:index
	for i := 0; i <= len(nums); i++ {
		if i != 0 {
			if nums[i-1] == 1 {
				preSum[i] = 1
			} else {
				preSum[i] = -1
			}
			preSum[i] += preSum[i-1]
		}
		if myMap[preSum[i]] == 0 {
			myMap[preSum[i]] = i + 1
		} else {
			if (i+1)-myMap[preSum[i]] > max {
				max = (i + 1) - myMap[preSum[i]]
			}
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
