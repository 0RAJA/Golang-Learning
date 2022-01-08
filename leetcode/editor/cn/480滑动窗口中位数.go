//中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
//
// 例如：
//
//
// [2,3,4]，中位数是 3
// [2,3]，中位数是 (2 + 3) / 2 = 2.5
//
//
// 给你一个数组 nums，有一个长度为 K 的窗口从最左端滑动到最右端。窗口中有 K 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗
//口中元素的中位数，并输出由它们组成的数组。
//
//
//
// 示例：
//
// 给出 nums = [1,3,-1,-3,5,3,6,7]，以及 K = 3。
//
//
//窗口位置                      中位数
//---------------               -----
//[1  3  -1] -3  5  3  6  7       1
// 1 [3  -1  -3] 5  3  6  7      -1
// 1  3 [-1  -3  5] 3  6  7      -1
// 1  3  -1 [-3  5  3] 6  7       3
// 1  3  -1  -3 [5  3  6] 7       5
// 1  3  -1  -3  5 [3  6  7]      6
//
//
// 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
//
//
//
// 提示：
//
//
// 你可以假设 K 始终有效，即：K 始终小于等于输入的非空数组的元素个数。
// 与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。
//
// Related Topics 数组 哈希表 滑动窗口 堆（优先队列） 👍 318 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func medianSlidingWindow(nums []int, k int) (ret []float64) {
	for i, j := 0, k-1; j < len(nums); i, j = i+1, j+1 {
		if (j-i+1)%2 == 0 {
			ret = append(ret, float64(nums[(i+j)/2]))
		} else {
			ret = append(ret, float64((nums[(i+j)/2]+nums[(i+j)/2+1])/2))
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
