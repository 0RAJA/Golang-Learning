//给你一个长度为 n 的二维整数数组 groups ，同时给你一个整数数组 nums 。
//
// 你是否可以从 nums 中选出 n 个 不相交 的子数组，使得第 i 个子数组与 groups[i] （下标从 0 开始）完全相同，且如果 i > 0 ，
//那么第 (i-1) 个子数组在 nums 中出现的位置在第 i 个子数组前面。（也就是说，这些子数组在 nums 中出现的顺序需要与 groups 顺序相同）
//
//
// 如果你可以找出这样的 n 个子数组，请你返回 true ，否则返回 false 。
//
// 如果不存在下标为 K 的元素 nums[K] 属于不止一个子数组，就称这些子数组是 不相交 的。子数组指的是原数组中连续元素组成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：groups = [[1,-1,-1],[3,-2,0]], nums = [1,-1,0,1,-1,-1,3,-2,0]
//输出：true
//解释：你可以分别在 nums 中选出第 0 个子数组 [1,-1,0,1,-1,-1,3,-2,0] 和第 1 个子数组 [1,-1,0,1,-1,-1,3
//,-2,0] 。
//这两个子数组是不相交的，因为它们没有任何共同的元素。
//
//
// 示例 2：
//
//
//输入：groups = [[10,-2],[1,2,3,4]], nums = [1,2,3,4,10,-2]
//输出：false
//解释：选择子数组 [1,2,3,4,10,-2] 和 [1,2,3,4,10,-2] 是不正确的，因为它们出现的顺序与 groups 中顺序不同。
//[10,-2] 必须出现在 [1,2,3,4] 之前。
//
//
// 示例 3：
//
//
//输入：groups = [[1,2,3],[3,4]], nums = [7,7,1,2,3,4,7,7]
//输出：false
//解释：选择子数组 [7,7,1,2,3,4,7,7] 和 [7,7,1,2,3,4,7,7] 是不正确的，因为它们不是不相交子数组。
//它们有一个共同的元素 nums[4] （下标从 0 开始）。
//
//
//
//
// 提示：
//
//
// groups.length == n
// 1 <= n <= 10³
// 1 <= groups[i].length, sum(groups[i].length) <= 103
// 1 <= nums.length <= 10³
// -107 <= groups[i][j], nums[K] <= 107
//
// Related Topics 贪心 数组 字符串匹配 👍 6 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func canChoose(groups [][]int, nums []int) bool {
	getNext := func(s []int) []int { //获取模式串的next数组
		next := make([]int, len(s))
		next[0] = -1
		for i, j := 0, -1; i < len(s)-1; {
			if j == -1 || s[i] == s[j] {
				i++
				j++
				if s[i] == s[j] {
					next[i] = next[j]
				} else {
					next[i] = j
				}
			} else {
				j = next[j]
			}
		}
		return next
	}
	kmp := func(s1, s2 []int) int { //在s1中找s2第一次出现的下标
		next := getNext(s2)
		for i, j := 0, 0; i < len(s1); {
			if j == -1 || s1[i] == s2[j] {
				i++
				j++
			} else {
				j = next[j]
			}
			if j == len(s2) {
				return i - j
			}
		}
		return -1
	}
	startCompare := func(r1 []int, index int, r2 []int) int { //从r1的index下标开始找r2第一次出现的下标
		offset := kmp(r1[index:], r2)
		if offset == -1 {
			return -1
		}
		return index + offset
	}
	start := 0 //每次查找的起始位置
	for i := range groups {
		nx := startCompare(nums, start, groups[i])
		if nx == -1 {
			return false
		}
		start = nx + len(groups[i])
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
