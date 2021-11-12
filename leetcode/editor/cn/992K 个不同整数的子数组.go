//给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定不同的子数组为好子数组。
//
// （例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
//
// 返回 A 中好子数组的数目。
//
//
//
// 示例 1：
//
//
//输入：A = [1,2,1,2,3], K = 2
//输出：7
//解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
//
//
// 示例 2：
//
//
//输入：A = [1,2,1,3,4], K = 3
//输出：3
//解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 20000
// 1 <= A[i] <= A.length
// 1 <= K <= A.length
//
// Related Topics 数组 哈希表 计数 滑动窗口 👍 327 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 3}
	k := 2
	fmt.Println(subarraysWithKDistinct(nums, k))
}

//leetcode submit region begin(Prohibit modification and deletion)
func subarraysWithKDistinct(nums []int, k int) (ret int) {
	return count(nums, k) - count(nums, k-1)
}
func count(nums []int, k int) (ret int) {
	cntMap := make(map[int]int)
	newWord := 0
	for l, r := 0, 0; r < len(nums); r++ {
		if cntMap[nums[r]] == 0 {
			newWord++
		}
		cntMap[nums[r]]++
		for newWord > k {
			cntMap[nums[l]]--
			if cntMap[nums[l]] == 0 {
				newWord--
			}
			l++
		}
		ret += r - l + 1
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
