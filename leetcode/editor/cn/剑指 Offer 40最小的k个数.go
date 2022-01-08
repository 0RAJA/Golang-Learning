//输入整数数组 arr ，找出其中最小的 K 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
// 示例 1：
//
// 输入：arr = [3,2,1], K = 2
//输出：[1,2] 或者 [2,1]
//
//
// 示例 2：
//
// 输入：arr = [0,1,2,1], K = 1
//输出：[0]
//
//
//
// 限制：
//
//
// 0 <= K <= arr.length <= 10000
// 0 <= arr[i] <= 10000
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 336 👎 0

package main

import (
	"container/heap"
	"sort"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type Test struct {
	sort.IntSlice
}

func (t *Test) Push(x interface{}) {
	t.IntSlice = append(t.IntSlice, x.(int))
}

func (t *Test) Pop() interface{} {
	p := t.IntSlice[len(t.IntSlice)-1]
	t.IntSlice = t.IntSlice[:len(t.IntSlice)-1]
	return p
}

func (t *Test) Less(i, j int) bool { return t.IntSlice[i] < t.IntSlice[j] }

func getLeastNumbers(arr []int, k int) (ret []int) {
	ret = make([]int, k)
	hp := &Test{}
	for i := range arr {
		heap.Push(hp, arr[i])
	}
	for i := 0; i < k; i++ {
		ret[i] = heap.Pop(hp).(int)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
