//给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。
//
// 如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。
//
//
//
//
//
//
// 示例 1：
//
// 输入：1
//输出：true
//
//
// 示例 2：
//
// 输入：10
//输出：false
//
//
// 示例 3：
//
// 输入：16
//输出：true
//
//
// 示例 4：
//
// 输入：24
//输出：false
//
//
// 示例 5：
//
// 输入：46
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= N <= 10^9
//
// Related Topics 数学 计数 枚举 排序 👍 86 👎 0

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println(reorderedPowerOf2(16))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reorderedPowerOf2(n int) bool {
	const MaxNum = 1000000000
	var num = 1
	res := make([]map[byte]int, 100)
	for i := 0; num <= MaxNum; i++ {
		res[i] = make(map[byte]int, 100)
		s := strconv.Itoa(num)
		for j := range s {
			res[i][s[j]]++
		}
		num *= 2
	}
	s := strconv.Itoa(n)
	resMap := make(map[byte]int, 100)
	for i := range s {
		resMap[s[i]]++
	}
	for _, v := range res {
		if reflect.DeepEqual(v, resMap) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
