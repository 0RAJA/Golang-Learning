//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。 
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下： 
//
// 
//P   A   H   N
//A P L S I I G
//Y   I   R 
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。 
//
// 请你实现这个将字符串进行指定行数变换的函数： 
//
// 
//string convert(string s, int numRows); 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
// 
//示例 2：
//
// 
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
// 
//
// 示例 3： 
//
// 
//输入：s = "A", numRows = 1
//输出："A"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 由英文字母（小写和大写）、',' 和 '.' 组成 
// 1 <= numRows <= 1000 
// 
// Related Topics 字符串 
// 👍 1216 👎 0
package main

import (
	"fmt"
)

func main() {
	s := "AB"
	numRows := 1
	fmt.Println(convert(s, numRows))
}

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) (ret string) {
	num := make([][]byte, numRows+1)
	for i := 1; i <= numRows; i++ {
		num[i] = make([]byte, len(s)+1)
	}
	x, y, cnt := 1, 1, 1
	flag := true
	for i := 0; i < len(s); i++ {
		num[x][y] = s[i]
		if flag && x < numRows {
			x++
			cnt++
			if cnt == numRows {
				flag = false
			}
		} else if flag == false {
			x--
			y++
			cnt--
			if cnt == 1 {
				flag = true
			}
		} else {
			y++
		}
	}
	bt := make([]byte, 0)
	for i := 1; i <= numRows; i++ {
		for j := 1; j <= len(s); j++ {
			if num[i][j] != 0 {
				bt = append(bt, num[i][j])
			}
		}
	}
	return string(bt)
}

//leetcode submit region end(Prohibit modification and deletion)
