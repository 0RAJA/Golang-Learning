//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划 👍 4293 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) (ret string) {
	isOK := func(l, r int) (int, int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return l + 1, r - 1
	}
	var left, right int
	for i := range s {
		l1, r1 := isOK(i, i)
		l2, r2 := isOK(i, i+1)
		if r1-l1 > right-left {
			right = r1
			left = l1
		}
		if r2-l2 > right-left {
			right = r2
			left = l2
		}
	}
	return s[left : right+1]
}

//leetcode submit region end(Prohibit modification and deletion)
