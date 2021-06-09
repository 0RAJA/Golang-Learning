//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。 
//
// 
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。 
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
//输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
//其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 
//n 的值 3 。
// 
//
// 示例 2： 
//
// 
//输入：strs = ["10", "0", "1"], m = 1, n = 1
//输出：2
//解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 600 
// 1 <= strs[i].length <= 100 
// strs[i] 仅由 '0' 和 '1' 组成 
// 1 <= m, n <= 100 
// 
// Related Topics 动态规划 
// 👍 440 👎 0
package main

import (
	"fmt"
	"sort"
)

//不会
func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	fmt.Println(findMaxForm(strs, 5, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)

func findMaxForm(strs []string, m int, n int) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	x1, x2 := 0, 0
	cnt := 0
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) > m-x1+n-x2 {
			break
		}
		t1, t2 := 0, 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				t1++
			} else {
				t2++
			}
		}
		if t1+x1 <= m && t2+x2 <= n {
			x1 += t1
			x2 += t2
			cnt++
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
