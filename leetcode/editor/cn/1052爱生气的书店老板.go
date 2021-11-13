//今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分
//钟结束后离开。
//
// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一
//分钟的顾客就会不满意，不生气则他们是满意的。
//
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
//
// 请你返回这一天营业下来，最多有多少客户能够感到满意。
//
//
// 示例：
//
//
//输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
//输出：16
//解释：
//书店老板在最后 3 分钟保持冷静。
//感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.
//
//
//
//
// 提示：
//
//
// 1 <= X <= customers.length == grumpy.length <= 20000
// 0 <= customers[i] <= 1000
// 0 <= grumpy[i] <= 1
//
// Related Topics 数组 滑动窗口 👍 196 👎 0

package main

import "fmt"

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	fmt.Println(maxSatisfied(customers, grumpy, minutes))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxSatisfied(customers []int, grumpy []int, minutes int) (ret int) {
	if minutes >= len(customers) { // 如果能全部不生气就直接返回所有值
		for _, v := range customers {
			ret += v
		}
		return
	}
	sum := 0
	for i := range customers { //先把所有不生气的人数加起来
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}
	for l, r := 0, 0; r < len(customers); r++ { // 维持一个 不生气的区间 往右走,遇到生气的就加入sum,抛弃生气的要从sum中减去
		if grumpy[r] == 1 {
			sum += customers[r]
		}
		if r-l+1 != minutes {
			continue
		}
		if sum > ret {
			ret = sum
		}
		if grumpy[l] == 1 {
			sum -= customers[l]
		}
		l++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
