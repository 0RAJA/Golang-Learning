//句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。
//
//
// 例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。
//
//
// 给你一个句子 s 和一个整数 K ，请你将 s 截断 ，使截断后的句子仅含 前 K 个单词。返回 截断 s 后得到的句子。
//
//
//
// 示例 1：
//
// 输入：s = "Hello how are you Contestant", K = 4
//输出："Hello how are you"
//解释：
//s 中的单词为 ["Hello", "how" "are", "you", "Contestant"]
//前 4 个单词为 ["Hello", "how", "are", "you"]
//因此，应当返回 "Hello how are you"
//
//
// 示例 2：
//
// 输入：s = "What is the solution to this problem", K = 4
//输出："What is the solution"
//解释：
//s 中的单词为 ["What", "is" "the", "solution", "to", "this", "problem"]
//前 4 个单词为 ["What", "is", "the", "solution"]
//因此，应当返回 "What is the solution"
//
// 示例 3：
//
// 输入：s = "chopper is not a tanuki", K = 5
//输出："chopper is not a tanuki"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// K 的取值范围是 [1, s 中单词的数目]
// s 仅由大小写英文字母和空格组成
// s 中的单词之间由单个空格隔开
// 不存在前导或尾随空格
//
// Related Topics 数组 字符串 👍 24 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func truncateSentence(s string, k int) string {
	var index int
	isOK := false
	for i := range s {
		if s[i] == ' ' {
			index++
		}
		if index == k {
			index = i
			isOK = true
			break
		}
	}
	if !isOK {
		index = len(s)
	}
	return s[:index]
}

//leetcode submit region end(Prohibit modification and deletion)
