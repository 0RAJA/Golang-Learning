//输入一个字符串，打印出该字符串中字符的所有排列。 
//
// 
//
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。 
//
// 
//
// 示例: 
//
// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
// 
//
// 
//
// 限制： 
//
// 1 <= s 的长度 <= 8 
// Related Topics 回溯算法 
// 👍 317 👎 0

package main

import "fmt"

func main() {
	fmt.Println(permutation("abc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func permutation(s string) []string {
	var returnStr = make([]string, 0)
	keyMap := make(map[int]bool)
	haveMap := make(map[string]bool)
	var dfs func(int, string)
	dfs = func(pos int, str string) {
		if pos == len(s) {
			ok, _ := haveMap[str]
			if ok == false {
				returnStr = append(returnStr, str)
				haveMap[str] = true
			}
			return
		}
		for i := 0; i < len(s); i++ {
			if keyMap[i] == false {
				keyMap[i] = true
				dfs(pos+1, str+string(s[i]))
				keyMap[i] = false
			}
		}
	}
	dfs(0, "")
	return returnStr
}

//leetcode submit region end(Prohibit modification and deletion)
