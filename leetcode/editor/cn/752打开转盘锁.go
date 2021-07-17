//你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
// 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
// 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
//
//
//
// 示例 1:
//
//
//输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//输出：6
//解释：
//可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
//注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
//因为当拨动到 "0102" 时这个锁就会被锁定。
//
//
// 示例 2:
//
//
//输入: deadends = ["8888"], target = "0009"
//输出：1
//解释：
//把最后一位反向旋转一次即可 "0000" -> "0009"。
//
//
// 示例 3:
//
//
//输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], targ
//et = "8888"
//输出：-1
//解释：
//无法旋转到目标数字且不被锁定。
//
//
// 示例 4:
//
//
//输入: deadends = ["0000"], target = "8888"
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= deadends.length <= 500
// deadends[i].length == 4
// target.length == 4
// target 不在 deadends 之中
// target 和 deadends[i] 仅由若干位数字组成
//
// Related Topics 广度优先搜索 数组 哈希表 字符串
// 👍 293 👎 0
package main

import "fmt"

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}

//leetcode submit region begin(Prohibit modification and deletion)

type pair struct {
	status string
	step   int
}

func openLock(deadends []string, target string) int {
	const start = "0000"
	if start == target {
		return 0
	}
	deadMap := make(map[string]bool, 0)
	for _, s := range deadends {
		deadMap[s] = true
	}
	if deadMap[target] == true || deadMap[start] == true {
		return -1
	}
	//获取所有改变的列表
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, v := range s {
			s[i] = v - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = v + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = v //恢复原样
		}
		return
	}
	//广搜
	queue := []pair{{
		status: start,
		step:   0,
	}}
	deadMap[start] = true
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, v := range get(p.status) {
			if deadMap[v] == false {
				if v == target {
					return p.step + 1
				}
				queue = append(queue, pair{
					status: v,
					step:   p.step + 1,
				})
				deadMap[v] = true
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
