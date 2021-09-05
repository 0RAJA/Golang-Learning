//已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
//
// 不要使用系统的 Math.random() 方法。
//
//
//
//
//
//
// 示例 1:
//
//
//输入: 1
//输出: [7]
//
//
// 示例 2:
//
//
//输入: 2
//输出: [8,4]
//
//
// 示例 3:
//
//
//输入: 3
//输出: [8,1,10]
//
//
//
//
// 提示:
//
//
// rand7 已定义。
// 传入参数: n 表示 rand10 的调用次数。
//
//
//
//
// 进阶:
//
//
// rand7()调用次数的 期望值 是多少 ?
// 你能否尽量少调用 rand7() ?
//
// Related Topics 数学 拒绝采样 概率与统计 随机化 👍 233 👎 0
package main

func main() {

}

func rand7() int {
	return 0
}

//leetcode submit region begin(Prohibit modification and deletion)

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		index := (row-1)*7 + col
		if index > 40 {
			continue
		}
		return 1 + (index-1)%10
	}
}

//leetcode submit region end(Prohibit modification and deletion)
