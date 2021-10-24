 //给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
//
//
//
//
// 示例 1：
//
//
//输入：[3,2,3]
//输出：[3]
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：[1,1,1,3,3,2,2,2]
//输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
// Related Topics 数组 哈希表 计数 排序 👍 447 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement2(nums []int) (ret []int) {
	ret = make([]int, 0, 2)
	cand1 := nums[0]
	count1 := 0
	cand2 := nums[0]
	count2 := 0
	for _, v := range nums {
		if v == cand1 {
			count1++
		} else if v == cand2 {
			count2++
		} else {
			if count1 == 0 {
				cand1 = v
				count1 = 1
			} else if count2 == 0 {
				cand2 = v
				count2 = 1
			} else {
				count1--
				count2--
			}
		}
	}
	count1 = 0
	count2 = 0
	for _, v := range nums {
		if v == cand1 {
			count1++
		}
		if v == cand2 {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		ret = append(ret, cand1)
	}
	if cand1 != cand2 && count2 > len(nums)/3 {
		ret = append(ret, cand2)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
