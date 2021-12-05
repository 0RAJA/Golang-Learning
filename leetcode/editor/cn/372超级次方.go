//你的任务是计算 aᵇ 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
//
//
//
// 示例 1：
//
//
//输入：a = 2, b = [3]
//输出：8
//
//
// 示例 2：
//
//
//输入：a = 2, b = [1,0]
//输出：1024
//
//
// 示例 3：
//
//
//输入：a = 1, b = [4,3,3,8,5,2]
//输出：1
//
//
// 示例 4：
//
//
//输入：a = 2147483647, b = [2,0,0]
//输出：1198
//
//
//
//
// 提示：
//
//
// 1 <= a <= 2³¹ - 1
// 1 <= b.length <= 2000
// 0 <= b[i] <= 9
// b 不含前导 0
//
// Related Topics 数学 分治 👍 133 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/*
高精度幂乘:
2*1234 = (2*123)^10 * 2^4
         ((2^12)^10 * 2^3)^10 * 2^4
         (((2^1)^10 * 2^2)^10 * 2^3)^10) * 2^4
a^[]nums = a^nums[length-1] * ((a^nums[:length-1])^10)
快速幂:
2^10 = 2^5 * 2^5
2^11 = 2^5 * 2^5 *2
a^b = a^(b/2) * a^(b/2) *a (如果b%2!=0)
if b == 0{
	return 1
}
*/
func superPow(a int, b []int) int {
	k := 1337
	var myPow func(int, int) int
	myPow = func(num int, cnt int) (ret int) {
		if cnt == 0 {
			return 1
		}
		num %= k
		sum := myPow(num, cnt/2) % k
		ret = sum * sum
		if cnt%2 != 0 {
			ret *= num
			ret %= k
		}
		return
	}
	var myPow2 func(int, []int) int
	myPow2 = func(a int, b []int) int {
		if len(b) == 0 {
			return 1
		}
		a %= k
		back := b[len(b)-1]
		t1 := myPow(a, back)
		t2 := myPow(myPow2(a, b[:len(b)-1]), 10)
		return (t1 * t2) % k
	}
	return myPow2(a, b)
}

//leetcode submit region end(Prohibit modification and deletion)
