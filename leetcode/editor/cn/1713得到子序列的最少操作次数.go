//给你一个数组 target ，包含若干 互不相同 的整数，以及另一个整数数组 arr ，arr 可能 包含重复元素。 
//
// 每一次操作中，你可以在 arr 的任意位置插入任一整数。比方说，如果 arr = [1,4,1,2] ，那么你可以在中间添加 3 得到 [1,4,3,1,
//2] 。你可以在数组最开始或最后面添加整数。 
//
// 请你返回 最少 操作次数，使得 target 成为 arr 的一个子序列。 
//
// 一个数组的 子序列 指的是删除原数组的某些元素（可能一个元素都不删除），同时不改变其余元素的相对顺序得到的数组。比方说，[2,7,4] 是 [4,2,3,
//7,2,1,4] 的子序列（加粗元素），但 [2,4,2] 不是子序列。 
//
// 
//
// 示例 1： 
//
// 输入：target = [5,1,3], arr = [9,4,2,3,4]
//输出：2
//解释：你可以添加 5 和 1 ，使得 arr 变为 [5,9,4,1,2,3,4] ，target 为 arr 的子序列。
// 
//
// 示例 2： 
//
// 输入：target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// 1 <= target.length, arr.length <= 105 
// 1 <= target[i], arr[i] <= 109 
// target 不包含任何重复元素。 
// 
// Related Topics 贪心 数组 哈希表 二分查找 
// 👍 43 👎 0
package main

import (
	"fmt"
	"sort"
)

func main() {
	target := []int{16, 7, 20, 11, 15, 13, 10, 14, 6, 8}
	arr := []int{11, 14, 15, 7, 5, 5, 6, 10, 11, 6}
	fmt.Println(minOperations(target, arr))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(target []int, arr []int) int {
	mapIndex := make(map[int]int)
	//将target的值变为下标
	for i, v := range target {
		mapIndex[v] = i
	}
	slice := make([]int, 0)
	//遍历arr.找
	for i := 0; i < len(arr); i++ {
		if k, f := mapIndex[arr[i]]; f {
			if p := sort.SearchInts(slice, k); p < len(slice) {
				slice[p] = k
			} else {
				slice = append(slice, k)
			}
		}
	}
	return len(target) - len(slice)
}

//func minOperations(target []int, arr []int) int {
//mapIndex := make(map[int][]int)
//for i, v := range arr {
//	mapIndex[v] = append(mapIndex[v], i)
//}
//cnt := 0
//now := 0
//for i := 0; i < len(target); i++ {
//	getNew := func() int {
//		for j := 0; j < len(mapIndex[target[i]]); j++ {
//			if mapIndex[target[i]][j] >= now {
//				return mapIndex[target[i]][j]
//			}
//		}
//		return -1
//	}
//	next := getNew()
//	if next == -1 {
//		cnt++
//	} else {
//		now = next + 1
//	}
//}
//return cnt
//}

//leetcode submit region end(Prohibit modification and deletion)
