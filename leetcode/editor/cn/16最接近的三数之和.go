//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 双指针 排序 👍 942 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) (ret int) {
	sort.Ints(nums)
	ret = math.MaxInt32
	n := len(nums)
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b, c := a+1, n-1; b < c; {
			for b > a+1 && nums[b] == nums[b-1] && b < c {
				b++
			}
			if b < c {
				sum := nums[a] + nums[b] + nums[c]
				if abs(sum-target) < abs(ret-target) {
					ret = sum
					if ret == target {
						return
					}
				}
				if sum < target {
					b++
				} else {
					c--
				}
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
