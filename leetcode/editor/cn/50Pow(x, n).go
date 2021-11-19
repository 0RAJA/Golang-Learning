//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x⁴
//
// Related Topics 递归 数学 👍 787 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/*
x→x^2→x^4→x^9→x^19→x^38→x^77
从后往前,递归式解法
*/
func myPow(x float64, n int) float64 {
	var quickMul func(float64, int) float64
	quickMul = func(x float64, n int) (ret float64) {
		if n == 0 {
			return 1
		}
		k := quickMul(x, n/2)
		ret = k * k
		if n%2 != 0 {
			ret *= x
		}
		return
	}
	if n < 0 {
		return 1.0 / quickMul(x, -n)
	}
	return quickMul(x, n)
}

//leetcode submit region end(Prohibit modification and deletion)
