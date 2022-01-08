//有台奇怪的打印机有以下两个特殊要求：
//
//
// 打印机每次只能打印由 同一个字符 组成的序列。
// 每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
//
//
// 给你一个字符串 s ，你的任务是计算这个打印机打印它需要的最少打印次数。
//
//
// 示例 1：
//
//
//输入：s = "aaabbb"
//输出：2
//解释：首先打印 "aaa" 然后打印 "bbb"。
//
//
// 示例 2：
//
//
//输入：s = "aba"
//输出：2
//解释：首先打印 "aaa" 然后在第二个位置打印 "b" 覆盖掉原来的字符 'a'。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 100
// s 由小写英文字母组成
//
// Related Topics 深度优先搜索 动态规划
// 👍 137 👎 0

//太难了...
package main

import "fmt"

func main() {
	fmt.Println(strangePrinter("aba"))
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
	dp[i][j] 中 i表示起始比较点,j表示终止比较点,dp中 当 i == j 时 , dp[i][j] = 1 表示自身到自身一个字符只需打印一次
	当 s[i] == s[j] 表示可以和上一次一起打印,所以 dp[i][j] = dp[i][j-1]
	当 s[i] != s[j] 表示不能和上一次一起打印了,需要找到一个最优解来打印它,即找到一个k点使得 dp[i][j] = dp[i][K]+ dp[K+1][j]
		即 dp[i][j] min(dp[i][K] + dp[K+1][j]),i <= K < j
	而此时 dp[i][j] = dp[i][K] + dp[K+1][j] 最小值的判断需要i从大往下进行初始化,以保证判断的数据均是已经初始化过的了
	dp[i][j] =
	{
		dp[i][j-1] 						(s[i] == s[j])
		min(dp[i][K] + dp[K+1][j])		(s[i] != s[j])
	}(i 从 len(s)-1 到 0,j 从 i+1 到 len(s)-1)
*/
func strangePrinter(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				min := 101
				dp[i][j] = func() int {
					for k := i; k < j; k++ {
						if min > dp[i][k]+dp[k+1][j] {
							min = dp[i][k] + dp[k+1][j]
						}
					}
					return min
				}()
			}
		}
	}
	return dp[0][len(s)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
