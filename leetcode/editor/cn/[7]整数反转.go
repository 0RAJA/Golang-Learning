//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。 
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231, 231 − 1] ，就返回 0。 
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
// 
//
// 示例 1： 
//
// 
//输入：x = 123
//输出：321
// 
//
// 示例 2： 
//
// 
//输入：x = -123
//输出：-321
// 
//
// 示例 3： 
//
// 
//输入：x = 120
//输出：21
// 
//
// 示例 4： 
//
// 
//输入：x = 0
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// -231 <= x <= 231 - 1 
// 
// Related Topics 数学 
// 👍 2751 👎 0
package main

import "math"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var flag bool = false
	if x < 0 {
		flag = true
		x = -x
	}
	var num int64 = 0
	for x != 0 {
		num = num*10 + int64(x%10)
		x /= 10
	}
	if flag {
		num = -num
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		num = 0
	}
	return int(num)
}

//leetcode submit region end(Prohibit modification and deletion)
