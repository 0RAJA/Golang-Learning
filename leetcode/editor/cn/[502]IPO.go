//假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限
//，它只能在 IPO 之前完成最多 K 个不同的项目。帮助 力扣 设计完成最多 K 个不同项目后得到最大总资本的方式。
//
// 给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
//
// 最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
//
// 总而言之，从给定项目中选择 最多 K 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
//
// 答案保证在 32 位有符号整数范围内。
//
//
//
// 示例 1：
//
//
//输入：K = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
//输出：4
//解释：
//由于你的初始资本为 0，你仅可以从 0 号项目开始。
//在完成后，你将获得 1 的利润，你的总资本将变为 1。
//此时你可以选择开始 1 号或 2 号项目。
//由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
//因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
//
//
// 示例 2：
//
//
//输入：K = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= K <= 10⁵
// 0 <= w <= 10⁹
// n == profits.length
// n == capital.length
// 1 <= n <= 10⁵
// 0 <= profits[i] <= 10⁴
// 0 <= capital[i] <= 10⁹
//
// Related Topics 贪心 数组 排序 堆（优先队列） 👍 202 👎 0
package main

import (
	"container/heap"
	"sort"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	type pair struct{ c, p int }
	arr := make([]pair, n)
	for i, p := range profits {
		arr[i] = pair{capital[i], p}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

	h := &Hp{}
	for cur := 0; k > 0; k-- {
		for cur < n && arr[cur].c <= w {
			heap.Push(h, arr[cur].p)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

type Hp struct{ sort.IntSlice }

func (h Hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *Hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

//type Value struct {
//	pro, cap int
//	isOK     bool
//}

//func findMaximizedCapital(K int, w int, profits []int, capital []int) int {
//	values := make([]Value, 0, len(profits))
//	for i := 0; i < len(profits); i++ {
//		x := Value{
//			pro: profits[i],
//			cap: capital[i],
//		}
//		values = append(values, x)
//	}
//	sort.Slice(values, func(i, j int) bool {
//		return values[i].pro > values[j].pro
//	})
//Loop:
//	length := len(values)
//	for i := 0; i < length && K > 0; i++ {
//		if values[i].isOK == false && w >= values[i].cap {
//			w += values[i].pro
//			values[i].isOK = true
//			K--
//			goto Loop
//		}
//	}
//	return w
//}

//leetcode submit region end(Prohibit modification and deletion)
