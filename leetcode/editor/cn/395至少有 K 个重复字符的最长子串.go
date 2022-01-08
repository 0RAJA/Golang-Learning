//给你一个字符串 s 和一个整数 K ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 K 。返回这一子串的长度。
//
//
//
// 示例 1：
//
//
//输入：s = "aaabb", K = 3
//输出：3
//解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
//
//
// 示例 2：
//
//
//输入：s = "ababbc", K = 2
//输出：5
//解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文字母组成
// 1 <= K <= 10⁵
//
// Related Topics 哈希表 字符串 分治 滑动窗口 👍 571 👎 0

package main

import "strings"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func mMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//函数定义:s字符串中每个字符出现次数均不小于k的最大字符串长度
func longestSubstring(s string, k int) (ret int) {
	if len(s) < k { // 当此字符串长度小于k直接返回0
		return
	}
	cntMap := make([]int, 26)
	for i := range s {
		cntMap[s[i]-'a']++
	}
	var ch byte // 找到整个字符串中出现次数小于k的字符,依次字符串为分割点,将所有字符串分隔递归进行计算.
	for i, v := range cntMap {
		if v > 0 && v < k {
			ch = byte(i) + 'a'
			break
		}
	}
	if ch == 0 { // 当此字符串中所有字符均满足条件则直接返回总长度
		return len(s)
	}
	for _, v := range strings.Split(s, string(ch)) { //到此步说明中间存在分割点可以继续向下分隔
		ret = mMax(ret, longestSubstring(v, k))
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
