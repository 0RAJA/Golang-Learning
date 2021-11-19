//给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为
//每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
//
//
//
// 示例 1:
//
//
//输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "xtfn"。
//
// 示例 2:
//
//
//输入: ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4
//解释: 这两个单词为 "ab", "cd"。
//
// 示例 3:
//
//
//输入: ["a","aa","aaa","aaaa"]
//输出: 0
//解释: 不存在这样的两个单词。
//
//
//
//
// 提示：
//
//
// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000
// words[i] 仅包含小写字母
//
// Related Topics 位运算 数组 字符串 👍 263 👎 0

package main

import (
	"sort"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(words []string) (ret int) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	cntMap := make([]int, len(words))
	for i := range words {
		nums := [26]bool{}
		for j := range words[i] {
			nums[int(words[i][j]-'a')] = true
		}
		for j := range nums {
			if nums[j] {
				cntMap[i] += 1
			}
			cntMap[i] = cntMap[i] << 1
		}
	}
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if cntMap[i]&cntMap[j] == 0 {
				if a := len(words[i]) * len(words[j]); a > ret {
					ret = a
				} else {
					break
				}
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
