//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
//
//
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
//
//
//
// 示例 1:
//
// 输入: 1
//输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
//
//
// 示例 2:
//
// 输入: 2
//输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0
//.05556,0.02778]
//
//
//
// 限制：
//
// 1 <= n <= 11
// Related Topics 数学 动态规划 概率与统计 👍 338 👎 0

package main

import "fmt"

func main() {
	fmt.Println(dicesProbability(2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func dicesProbability(n int) (ret []float64) {
	ret = make([]float64, 6)
	for i := range ret {
		ret[i] = 1.0 / 6
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, i*5+1)   //新建一层
		for j := 0; j < len(ret); j++ { //计算这一层给下一层做出的贡献 ret[j]*(1/6)
			for k := 0; k < 6; k++ {
				tmp[j+k] += ret[j] / 6
			}
		}
		ret = tmp
	}
	return
}

/*
1/6	1/6	1/6	1/6	1/6	1/6     i=1 len=i*5+1=6
							  i=2 len=i*5+1=11
*/
//leetcode submit region end(Prohibit modification and deletion)
