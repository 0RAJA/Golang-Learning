//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口 👍 6399 👎 0

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) (ret int) {
	cnt := 0
	mapCnt := make(map[byte]int)
	var c byte
	for l, r := 0, 0; r < len(s); r++ {
		if mapCnt[s[r]] > 0 {
			cnt++
			c = s[r]
		}
		mapCnt[s[r]]++
		for cnt != 0 {
			if s[l] == c {
				cnt--
			}
			mapCnt[s[l]]--
			l++
		}
		if ret < r-l+1 {
			ret = r - l + 1
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
