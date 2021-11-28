//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
//
//
//
// 示例 1:
//
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
//
// 示例 2:
//
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
//
// Related Topics 哈希表 字符串 滑动窗口 👍 663 👎 0

package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) (ret []int) {
	hash := map[byte]int{}
	for i := range p {
		hash[p[i]]++
	}
	cnt := len(hash)
	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := hash[s[r]]; ok {
			hash[s[r]]--
			if hash[s[r]] == 0 {
				cnt--
			}
		}
		for l < r && r-l+1 > len(p) {
			if _, ok := hash[s[l]]; ok {
				if hash[s[l]] == 0 {
					cnt++
				}
				hash[s[l]]++
			}
			l++
		}
		if r-l+1 == len(p) && cnt == 0 {
			ret = append(ret, l)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
