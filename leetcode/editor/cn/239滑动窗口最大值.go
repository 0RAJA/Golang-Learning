//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
//
// 示例 4：
//
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
//
// 示例 5：
//
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 1342 👎 0

package main

import "container/list"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type WinNum int

func (n WinNum) CompareTo(v Comparable) int {
	return int(n - v.(WinNum))
}

func maxSlidingWindow(nums []int, k int) (ret []int) {
	queue := NewStructure(new(Queue))
	for l, r := 0, 0; r < len(nums); r++ {
		queue.Push(WinNum(nums[r]))
		for l < r && r-l+1 > k {
			queue.Pop()
			l++
		}
		if r-l+1 == k {
			ret = append(ret, int(queue.Max().(WinNum)))
		}
	}
	return
}

type Comparable interface {
	CompareTo(v Comparable) int
}

type monotonous interface {
	Pop() interface{}
	Push(v Comparable)
	Top() interface{}
	Min() interface{}
	Max() interface{}
}

type Structure struct {
	monotonous
}

func NewStructure(monotonous monotonous) *Structure {
	return &Structure{monotonous: monotonous}
}

type Assists struct {
	nums      list.List
	assistMax list.List
	assistMin list.List
}

type Queue Assists

func (q *Queue) Push(v Comparable) {
	q.nums.PushBack(v)
	q.PushAssist(v)
}

func (q *Queue) Pop() interface{} {
	if q.nums.Len() <= 0 {
		return nil
	}
	t := q.nums.Front()
	q.PopAssist(t.Value.(Comparable))
	q.nums.Remove(t)
	return t.Value
}
func (q *Queue) Top() interface{} {
	if q.nums.Len() <= 0 {
		return nil
	}
	return q.nums.Front().Value
}
func (q *Queue) PopAssist(v Comparable) {
	if v.CompareTo(q.assistMin.Front().Value.(Comparable)) == 0 {
		q.assistMin.Remove(q.assistMin.Front())
	}
	if v.CompareTo(q.assistMax.Front().Value.(Comparable)) == 0 {
		q.assistMax.Remove(q.assistMax.Front())
	}
}

func (q *Queue) PushAssist(v Comparable) {
	for q.assistMin.Len() > 0 && v.CompareTo(q.assistMin.Back().Value.(Comparable)) < 0 {
		q.assistMin.Remove(q.assistMin.Back())
	}
	q.assistMin.PushBack(v)
	for q.assistMax.Len() > 0 && v.CompareTo(q.assistMax.Back().Value.(Comparable)) > 0 {
		q.assistMax.Remove(q.assistMax.Back())
	}
	q.assistMax.PushBack(v)
}

func (q *Queue) Min() interface{} {
	return q.assistMin.Front().Value
}

func (q *Queue) Max() interface{} {
	return q.assistMax.Front().Value
}

//leetcode submit region end(Prohibit modification and deletion)
