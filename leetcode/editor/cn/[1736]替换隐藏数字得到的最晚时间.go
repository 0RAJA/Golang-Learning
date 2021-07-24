//给你一个字符串 time ，格式为 hh:mm（小时：分钟），其中某几位数字被隐藏（用 ? 表示）。 
//
// 有效的时间为 00:00 到 23:59 之间的所有时间，包括 00:00 和 23:59 。 
//
// 替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。 
//
// 
//
// 示例 1： 
//
// 
//输入：time = "2?:?0"
//输出："23:50"
//解释：以数字 '2' 开头的最晚一小时是 23 ，以 '0' 结尾的最晚一分钟是 50 。
// 
//
// 示例 2： 
//
// 
//输入：time = "0?:3?"
//输出："09:39"
// 
//
// 示例 3： 
//
// 
//输入：time = "1?:22"
//输出："19:22"
// 
//
// 
//
// 提示： 
//
// 
// time 的格式为 hh:mm 
// 题目数据保证你可以由输入的字符串生成有效的时间 
// 
// Related Topics 字符串 
// 👍 16 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumTime(time string) string {
	times := []byte(time)
	for i := 0; i < 5; i++ {
		if i == 0 {
			if times[0] == '?' {
				if times[1] >= '4' && times[1] != '?' {
					times[0] = '1'
				} else {
					times[0] = '2'
				}
			}
		} else if i == 1 {
			if times[1] == '?' {
				if times[0] == '2' {
					times[1] = '3'
				} else {
					times[1] = '9'
				}
			}
		} else if i == 3 {
			if times[3] == '?' {
				times[3] = '5'
			}
		} else if i == 4 {
			if times[4] == '?' {
				times[4] = '9'
			}
		}
	}
	return string(times)
}

//leetcode submit region end(Prohibit modification and deletion)
