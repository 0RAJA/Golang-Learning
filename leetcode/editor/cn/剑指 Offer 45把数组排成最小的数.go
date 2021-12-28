//输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//
//
//
// 示例 1:
//
// 输入: [10,2]
//输出: "102"
//
// 示例 2:
//
// 输入: [3,30,34,5,9]
//输出: "3033459"
//
//
//
// 提示:
//
//
// 0 < nums.length <= 100
//
//
// 说明:
//
//
// 输出结果可能非常大，所以你需要返回一个字符串而不是整数
// 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
//
// Related Topics 贪心 字符串 排序 👍 341 👎 0

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 0}
	fmt.Println(minNumber(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minNumber(nums []int) (ret string) {
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a+b < b+a
	})
	for i := range nums {
		ret += strconv.Itoa(nums[i])
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
