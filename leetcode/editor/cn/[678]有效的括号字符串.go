//给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：
//
//
// 任何左括号 ( 必须有相应的右括号 )。
// 任何右括号 ) 必须有相应的左括号 ( 。
// 左括号 ( 必须在对应的右括号之前 )。
// * 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
// 一个空字符串也被视为有效字符串。
//
//
// 示例 1:
//
//
//输入: "()"
//输出: True
//
//
// 示例 2:
//
//
//输入: "(*)"
//输出: True
//
//
// 示例 3:
//
//
//输入: "(*))"
//输出: True
//
//
// 注意:
//
//
// 字符串大小将在 [1，100] 范围内。
//
// Related Topics 栈 贪心 字符串 动态规划 👍 325 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func checkValidString(s string) bool {
	var left []int
	var star []int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			left = append(left, i)
		case '*':
			star = append(star, i)
		case ')':
			if len(left) > 0 {
				left = left[:len(left)-1]
			} else if len(star) > 0 {
				star = star[:len(star)-1]
			} else {
				return false
			}
		}
	}
	for len(left) > 0 && len(star) > 0 {
		a := left[len(left)-1]
		b := star[len(star)-1]
		left = left[:len(left)-1]
		star = star[:len(star)-1]
		if a > b {
			return false
		}
	}
	if len(left) > 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
