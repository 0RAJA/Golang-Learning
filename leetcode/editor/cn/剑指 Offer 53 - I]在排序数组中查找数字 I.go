//统计一个数字在排序数组中出现的次数。 
//
// 
//
// 示例 1: 
//
// 输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2 
//
// 示例 2: 
//
// 输入: nums = [5,7,7,8,8,10], target = 6
//输出: 0 
//
// 
//
// 限制： 
//
// 0 <= 数组长度 <= 50000 
//
// 
//
// 注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-
//position-of-element-in-sorted-array/ 
// Related Topics 数组 二分查找 
// 👍 149 👎 0
package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	target := 5
	fmt.Println(search(nums, target))
}

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) (ret int) {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	for i := low; i<len(nums)&&nums[i] == target; i++ {
		ret++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
