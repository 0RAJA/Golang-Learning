//有效数字（按顺序）可以分成以下几个部分： 
//
// 
// 一个 小数 或者 整数 
// （可选）一个 'e' 或 'E' ，后面跟着一个 整数 
// 
//
// 小数（按顺序）可以分成以下几个部分： 
//
// 
// （可选）一个符号字符（'+' 或 '-'） 
// 下述格式之一：
// 
// 至少一位数字，后面跟着一个点 '.' 
// 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字 
// 一个点 '.' ，后面跟着至少一位数字 
// 
// 
// 
//
// 整数（按顺序）可以分成以下几个部分： 
//
// 
// （可选）一个符号字符（'+' 或 '-'） 
// 至少一位数字 
// 
//
// 部分有效数字列举如下： 
//
// 
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1",
// "53.5e93", "-123.456e789"] 
// 
//
// 部分无效数字列举如下： 
//
// 
// ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"] 
// 
//
// 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "0"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "e"
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：s = "."
//输出：false
// 
//
// 示例 4： 
//
// 
//输入：s = ".1"
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 20 
// s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。 
// 
// Related Topics 数学 字符串 
// 👍 259 👎 0

package main

import "fmt"

func main() {
	fmt.Println(isNumber("95a54e53"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isNumber(s string) bool {
	Sign := false      //有没有正负号
	Point := false     //有没有小数点
	integerCount := 0  //整数部分
	decimalCount := 0  //小数部分
	e := false         //有没有e
	eSign := false     //e后整数的符号
	eIntegerCount := 0 //e后整数部分
	isTrue := func() bool {
		return Point && (integerCount > 0 || decimalCount > 0) || !Point && integerCount > 0
	}
	for i := 0; i < len(s); i++ {
		if isE(s[i]) {
			if isTrue() && e != true {
				e = true
			} else {
				return false
			}
		} else if e == true {
			if isSign(s[i]) && eIntegerCount == 0 {
				if eSign { //避免正负号重复
					return false
				}
				eSign = true
			} else if isNum(s[i]) {
				eIntegerCount++
			} else {
				return false
			}
		} else if isSign(s[i]) {
			if Sign || integerCount != 0 || Point { //多个正负号或者前面有数字或者有小数点
				return false
			}
			Sign = true
		} else if isPoint(s[i]) {
			if Point { //确保小数点只有一个且为小数点前只有一个数字
				return false
			}
			Point = true
		} else if isNum(s[i]) {
			if Point {
				decimalCount++ //小数部分++
			} else {
				integerCount++ //整数部分++
			}
		} else {
			return false
		}
	}
	if e {
		return eIntegerCount > 0
	}
	return isTrue()
}
func isNum(key byte) bool {
	return key >= '0' && key <= '9'
}
func isE(key byte) bool {
	return key == 'E' || key == 'e'
}
func isPoint(key byte) bool {
	return key == '.'
}
func isSign(key byte) bool {
	return key == '+' || key == '-'
}

//leetcode submit region end(Prohibit modification and deletion)
