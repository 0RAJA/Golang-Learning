//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
//
//
//
// 示例 1:
//
// 输入: [7,5,6,4]
//输出: 5
//
//
//
// 限制：
//
// 0 <= 数组长度 <= 50000
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 586 👎 0

package main

import "fmt"

func main() {
	fmt.Println(reversePairs([]int{9, 1, 8, 2, 7, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) (ret int) {
	merge := func(nums []int, begin, mid, end int) {
		tmp := make([]int, 0, end-begin+1)
		l, r := begin, mid
		for l < mid && r < end {
			if nums[r] < nums[l] {
				tmp = append(tmp, nums[r])
				ret += mid - l
				r++
			} else {
				tmp = append(tmp, nums[l])
				l++
			}
		}
		tmp = append(tmp, nums[l:mid]...)
		tmp = append(tmp, nums[r:end]...)
		for i := range tmp {
			nums[begin+i] = tmp[i]
		}
	}
	var mergeSort func([]int, int, int)
	mergeSort = func(nums []int, l, r int) {
		if r-l > 1 {
			mid := l + (r-l)/2
			mergeSort(nums, l, mid)
			mergeSort(nums, mid, r)
			merge(nums, l, mid, r)
		}
	}
	mergeSort(nums, 0, len(nums))
	fmt.Println(nums)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
