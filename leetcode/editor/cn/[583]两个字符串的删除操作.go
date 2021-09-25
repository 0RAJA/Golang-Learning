//给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
//
//
//
// 示例：
//
// 输入: "sea", "eat"
//输出: 2
//解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
//
//
//
//
// 提示：
//
//
// 给定单词的长度不超过500。
// 给定单词中的字符只含有小写字母。
//
// Related Topics 字符串 动态规划 👍 272 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	n := myCmp(word1, word2)
	return len(word1) + len(word2) - 2*n
}

// myCmp 返回s1和s2的最长子序列长度
func myCmp(s1, s2 string) int {
	dp := make([][]int, len(s1)+1) //行是len(s1)+1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1) //列是len(s2)+1
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = func() int {
					if dp[i-1][j] > dp[i][j-1] {
						return dp[i-1][j]
					}
					return dp[i][j-1]
				}()
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

//leetcode submit region end(Prohibit modification and deletion)
