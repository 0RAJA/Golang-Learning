//给定一个字符串 s 和一个整数 K，从字符串开头算起，每 2k 个字符反转前 K 个字符。
//
//
// 如果剩余字符少于 K 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 K 个，则反转前 K 个字符，其余字符保持原样。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abcdefg", K = 2
//输出："bacdfeg"
//
//
// 示例 2：
//
//
//输入：s = "abcd", K = 2
//输出："bacd"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= K <= 10⁴
//
// Related Topics 双指针 字符串 👍 143 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	x := append([]byte("0"), s...)
	//fmt.Println(x)
	for i := 1; i < len(s); i += 2 * k {
		left := i
		right := left + k - 1
		if right > len(s) {
			right = len(s)
		}
		for left < right {
			x[left], x[right] = x[right], x[left]
			left++
			right--
		}
	}
	return string(x[1:])
}

//leetcode submit region end(Prohibit modification and deletion)
