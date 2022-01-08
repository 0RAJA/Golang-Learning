//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: K[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 K 次。注意 K 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 K ，例如不会出现像 3a 或 2[4] 的输入。
//
//
//
// 示例 1：
//
// 输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
//
// 示例 2：
//
// 输入：s = "3[a2[c]]"
//输出："accaccacc"
//
//
// 示例 3：
//
// 输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//
//
// 示例 4：
//
// 输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
//
// Related Topics 栈 递归 字符串 👍 911 👎 0

package main

import "fmt"

func main() {
	fmt.Println(decodeString("100[a2[c]]"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) (ret string) {
	stack := NewStackM()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			str := ""
			for {
				p := stack.Pop()
				if p == '[' {
					break
				}
				str = string(p) + str
			}
			var n int
			for cnt := 1; ; cnt *= 10 {
				if stack.IsEmpty() || (stack.Top() < '0' || stack.Top() > '9') {
					break
				}
				n += int(stack.Pop()-'0') * cnt
			}
			for j := 0; j < n; j++ {
				stack.PushStr(str)
			}
		default:
			stack.Push(s[i])
		}
	}
	for !stack.IsEmpty() {
		ret = string(stack.Pop()) + ret
	}
	return
}

type StackM struct {
	stack []byte
}

func NewStackM() *StackM {
	return &StackM{stack: make([]byte, 0, 1000)}
}

func (s *StackM) Push(c byte) {
	s.stack = append(s.stack, c)
}
func (s *StackM) PushStr(str string) {
	for i := 0; i < len(str); i++ {
		s.Push(str[i])
	}
}
func (s *StackM) Pop() byte {
	p := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return p
}
func (s *StackM) Top() byte {
	return s.stack[len(s.stack)-1]
}

func (s *StackM) IsEmpty() bool {
	return len(s.stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
