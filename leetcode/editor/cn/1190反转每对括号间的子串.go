//给出一个字符串 s（仅含有小写英文字母和括号）。 
//
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。 
//
// 注意，您的结果中 不应 包含任何括号。 
//
// 
//
// 示例 1： 
//
// 输入：s = "(abcd)"
//输出："dcba"
// 
//
// 示例 2： 
//
// 输入：s = "(u(love)i)"
//输出："iloveu"
// 
//
// 示例 3： 
//
// 输入：s = "(ed(et(oc))el)"
//输出："leetcode"
// 
//
// 示例 4： 
//
// 输入：s = "a(bcdefghijkl(mno)p)q"
//输出："apmnolkjihgfedcbq"
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 2000 
// s 中只有小写英文字母和括号 
// 我们确保所有括号都是成对出现的 
// 
// Related Topics 栈 
// 👍 137 👎 0
package main

import "fmt"

func main() {
	s := "a(bcdefghijkl(mno)p)q"
	fmt.Println(reverseParentheses(s))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseParentheses(s string) string {
	char := make([]byte, 0, 2001) //放不是括号的字符
	brackets := make([]int, 0)    //放左括号的下标
	for _, c := range s {
		if c == '(' { //遇到括号存起来下标
			brackets = append(brackets, len(char))
		} else if c == ')' { //遇到右括号把最近的左括号到这里的字符串翻转
			left, right := brackets[len(brackets)-1], len(char)
			for i := left; i < (left+right)/2; i++ {
				char[i], char[left+right-i-1] = char[left+right-i-1], char[i]
			}
			brackets = brackets[:len(brackets)-1]
		} else { //非括号字符直接存起来
			char = append(char, byte(c))
		}
	}
	return string(char)
}

//leetcode submit region end(Prohibit modification and deletion)
