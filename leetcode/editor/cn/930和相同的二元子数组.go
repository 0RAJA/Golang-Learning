//给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。 
//
// 子数组 是数组的一段连续部分。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,0,1,0,1], goal = 2
//输出：4
//解释：
//如下面黑体所示，有 4 个满足题目要求的子数组：
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,0,0,0,0], goal = 0
//输出：15
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 3 * 104 
// nums[i] 不是 0 就是 1 
// 0 <= goal <= nums.length 
// 
// Related Topics 数组 哈希表 前缀和 滑动窗口 
// 👍 117 👎 0
package main

import "fmt"

func main() {
	nums := []int{0, 0, 0, 0, 0}
	goal := 0
	fmt.Println(numSubarraysWithSum(nums, goal))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
我们用哈希表记录每一种前缀和出现的次数，假设我们当前枚举到元素 nums[j]，我们只需要查询哈希表中元素 sum[j]−goal 的数量即可，
这些元素的数量即对应了以当前 j 值为右边界的满足条件的子数组的数量。
最后这些元素的总数量即为所有和为goal 的子数组数量
*/
func numSubarraysWithSum(nums []int, goal int) (ret int) {
	var prefix = make(map[int]int, 0)
	sum := 0
	for _, num := range nums {
		prefix[sum]++ //只用保存前n-1个前缀和
		sum += num
		ret += prefix[sum-goal]
	}
	return
}

//func numSubarraysWithSum(nums []int, goal int) (sum int) {
//	prefix := make([]int, len(nums)+1)
//	for i := 1; i <= len(nums); i++ {
//		prefix[i] = prefix[i-1] + nums[i-1]
//	}
//	for i := 1; i <= len(nums); i++ {
//		for j := 0; j < i; j++ {
//			ret := prefix[i] - prefix[j]
//			if ret == goal {
//				sum++
//			} else if ret < goal {
//				break
//			}
//		}
//	}
//	return
//}

//leetcode submit region end(Prohibit modification and deletion)
