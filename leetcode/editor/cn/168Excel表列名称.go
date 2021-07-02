//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。 
//
// 例如： 
//
// 
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28 
//...
// 
//
// 
//
// 示例 1： 
//
// 
//输入：columnNumber = 1
//输出："A"
// 
//
// 示例 2： 
//
// 
//输入：columnNumber = 28
//输出："AB"
// 
//
// 示例 3： 
//
// 
//输入：columnNumber = 701
//输出："ZY"
// 
//
// 示例 4： 
//
// 
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= columnNumber <= 231 - 1 
// 
// Related Topics 数学 字符串 
// 👍 420 👎 0
package main

import (
	"fmt"
)

func main() {
	columnNumber := 701
	fmt.Println(convertToTitle(columnNumber))
}

//leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	var ret []byte
	for columnNumber > 0 {
		columnNumber-- //往左偏移一位
		ret = append(ret, byte(columnNumber%26+'A'))
		columnNumber /= 26
	}
	for i, n := 0, len(ret); i < n/2; i++ {
		ret[i], ret[n-1-i] = ret[n-1-i], ret[i]
	}
	return string(ret)
}

//leetcode submit region end(Prohibit modification and deletion)
