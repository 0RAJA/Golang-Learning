//给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
//
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：s = "bbbab"
//输出：4
//解释：一个可能的最长回文子序列为 "bbbb" 。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出：2
//解释：一个可能的最长回文子序列为 "bb" 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由小写英文字母组成
//
// Related Topics 字符串 动态规划
// 👍 501 👎 0
package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("aabaa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 1; i <= len(s); i++ {
		dp[i][i] = 1
		for j := i - 1; j >= 1; j-- {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j+1] + 2 //找到相同的则+2个
			} else {
				dp[i][j] = max1(dp[i-1][j], dp[i][j+1]) //选择不走i或者j的大的那个
			}
		}
	}
	return dp[len(s)][1]
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
