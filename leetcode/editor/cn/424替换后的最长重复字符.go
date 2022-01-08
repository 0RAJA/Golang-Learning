//给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 K 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
//
//
// 注意：字符串长度 和 K 不会超过 10⁴。
//
//
//
// 示例 1：
//
//
//输入：s = "ABAB", K = 2
//输出：4
//解释：用两个'A'替换为两个'B',反之亦然。
//
//
// 示例 2：
//
//
//输入：s = "AABABBA", K = 1
//输出：4
//解释：
//将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
//子串 "BBBB" 有最长重复字母, 答案为 4。
//
// Related Topics 哈希表 字符串 滑动窗口 👍 508 👎 0

package main

import "fmt"

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println(characterReplacement(s, k))
}

//leetcode submit region begin(Prohibit modification and deletion)
func characterReplacement(s string, k int) (ret int) {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	count := func(nums [26]int) (m int) {
		for _, v := range nums {
			m = max(m, v)
		}
		return
	}

	cnt := [26]int{}                    // 计算数组
	for l, r := 0, 0; r < len(s); r++ { // 定左右边界 每一轮让右边界往右走
		cnt[s[r]-'A']++              // 计入新情况
		for (r-l+1)-count(cnt) > k { // 判断是否满足条件,如果杂物太多就缩小左边界(不满足条件)
			cnt[s[l]-'A']--
			l++
		}
		ret = max(ret, r-l+1) //统计结果
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
