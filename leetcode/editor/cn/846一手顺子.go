//Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，使每一组的牌数都是 groupSize ，并且由 groupSize 张连续的牌组成。
//
// 给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。如果她可能重新排列这些牌，返回 true ；否则，
//返回 false 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
//输出：true
//解释：Alice 手中的牌可以被重新排列为 [1,2,3]，[2,3,4]，[6,7,8]。
//
// 示例 2：
//
//
//输入：hand = [1,2,3,4,5], groupSize = 4
//输出：false
//解释：Alice 手中的牌无法被重新排列成几个大小为 4 的组。
//
//
//
// 提示：
//
//
// 1 <= hand.length <= 10⁴
// 0 <= hand[i] <= 10⁹
// 1 <= groupSize <= hand.length
//
//
//
//
// 注意：此题目与 1296 重复：https://leetcode-cn.com/problems/divide-array-in-sets-of-k-
//consecutive-numbers/
// Related Topics 贪心 数组 哈希表 排序 👍 147 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 1, 2, 2, 3, 3}, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
type PQueue struct {
	sort.IntSlice
}

func (p *PQueue) Push(x interface{}) {
	p.IntSlice = append(p.IntSlice, x.(int))
}

func (p *PQueue) Pop() (ret interface{}) {
	ret = (p.IntSlice)[len(p.IntSlice)-1]
	p.IntSlice = (p.IntSlice)[:len(p.IntSlice)-1]
	return
}

func isNStraightHand(hand []int, groupSize int) bool {
	pq := &PQueue{}
	hash := make(map[int]int, len(hand))
	for _, v := range hand {
		heap.Push(pq, v)
		hash[v]++
	}
	for pq.Len() > 0 {
		q := heap.Pop(pq).(int)
		if hash[q] == 0 {
			continue
		}
		for i := q; i < q+groupSize; i++ {
			if hash[i] <= 0 {
				return false
			}
			hash[i]--
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
