//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,3,2,5]
//输出：[3,5]
//解释：[5, 3] 也是有效的答案。
//
//
// 示例 2：
//
//
//输入：nums = [-1,0]
//输出：[-1,0]
//
//
// 示例 3：
//
//
//输入：nums = [0,1]
//输出：[1,0]
//
//
// 提示：
//
//
// 2 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次
//
// Related Topics 位运算 数组 👍 460 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
//先找到两个数的异或值,异或值为相同为0不同为1,找到x最低位的1,说明两个数中一个此位为0,另一个为1,依次将原数据分为两组,再进行异或求解
func singleNumber(nums []int) []int {
	n := 0
	for _, v := range nums {
		n ^= v
	}
	x := n & (-n)
	n1, n2 := 0, 0
	for _, v := range nums {
		if v&x == 0 {
			n1 ^= v
		} else {
			n2 ^= v
		}
	}
	return []int{n1, n2}
}

//leetcode submit region end(Prohibit modification and deletion)
