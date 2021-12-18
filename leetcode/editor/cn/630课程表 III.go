//这里有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，其中 courses[i] = [durationi,
//lastDayi] 表示第 i 门课将会 持续 上 durationi 天课，并且必须在不晚于 lastDayi 的时候完成。
//
// 你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。
//
// 返回你最多可以修读的课程数目。
//
//
//
// 示例 1：
//
//
//输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
//输出：3
//解释：
//这里一共有 4 门课程，但是你最多可以修 3 门：
//首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
//第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
//第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
//第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。
//
// 示例 2：
//
//
//输入：courses = [[1,2]]
//输出：1
//
//
// 示例 3：
//
//
//输入：courses = [[3,2],[4,3]]
//输出：0
//
//
//
//
// 提示:
//
//
// 1 <= courses.length <= 10⁴
// 1 <= durationi, lastDayi <= 10⁴
//
// Related Topics 贪心 数组 堆（优先队列） 👍 241 👎 0

package main

import (
	"container/heap"
	"sort"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func scheduleCourse(courses [][]int) int {
	/*
		通过截止日期进行排序,同时将每个被选择课目的信息保存在一个优先队列中,每次在选择一个新的科目时看能否被选择,
		如果时间不够就看优先队列中耗费时间最多的和此课目的耗费时间比较,选择较小的那个.最后返回优先队列中的元素个数
	*/
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	hp := &Heap{}
	total := 0
	for _, course := range courses {
		if t := course[0]; t+total <= course[1] {
			heap.Push(hp, t)
			total += t
		} else if hp.Len() > 0 && hp.IntSlice[0] > t {
			total += t - hp.IntSlice[0]
			hp.IntSlice[0] = t
			heap.Fix(hp, 0)
		}
	}
	return hp.IntSlice.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
