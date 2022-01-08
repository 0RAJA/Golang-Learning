//给你一个整数数组 arr 。
//
// 现需要从数组中取三个下标 i、j 和 K ，其中 (0 <= i < j <= K < arr.length) 。
//
// a 和 b 定义如下：
//
//
// a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
// b = arr[j] ^ arr[j + 1] ^ ... ^ arr[K]
//
//
// 注意：^ 表示 按位异或 操作。
//
// 请返回能够令 a == b 成立的三元组 (i, j , K) 的数目。
//
//
//
// 示例 1：
//
// 输入：arr = [2,3,1,6,7]
//输出：4
//解释：满足题意的三元组分别是 (0,1,2), (0,2,2), (2,3,4) 以及 (2,4,4)
//
//
// 示例 2：
//
// 输入：arr = [1,1,1,1,1]
//输出：10
//
//
// 示例 3：
//
// 输入：arr = [2,3]
//输出：0
//
//
// 示例 4：
//
// 输入：arr = [1,3,5,7,9]
//输出：3
//
//
// 示例 5：
//
// 输入：arr = [7,11,12,9,5,2,7,17,22]
//输出：8
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 300
// 1 <= arr[i] <= 10^8
//
// Related Topics 位运算 数组 数学
// 👍 113 👎 0

package main

import "fmt"

func main() {
	arr := []int{7, 11, 12, 9, 5, 2, 7, 17, 22}
	fmt.Println(countTriplets(arr))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countTriplets(arr []int) int {
	num := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ { //先存放异或完的数组
		num[i] = num[i-1] ^ arr[i-1]
	}
	cnt := 0
	for j := 2; j <= len(arr); j++ {
		map1 := make(map[int]int)
		for i := j - 1; i >= 1; i-- { //将i决定的值a存到map里
			map1[num[j-1]^num[i-1]]++
		}
		for k := j; k <= len(arr); k++ { //检索k决定的b在map中能否找到
			cnt += map1[num[k]^num[j-1]]
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
