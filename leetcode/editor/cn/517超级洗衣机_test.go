//假设有 n 台超级洗衣机放在同一排上。开始的时候，每台洗衣机内可能有一定量的衣服，也可能是空的。
//
// 在每一步操作中，你可以选择任意 m (1 <= m <= n) 台洗衣机，与此同时将每台洗衣机的一件衣服送到相邻的一台洗衣机。
//
// 给定一个整数数组 machines 代表从左至右每台洗衣机中的衣物数量，请给出能让所有洗衣机中剩下的衣物的数量相等的 最少的操作步数 。如果不能使每台洗衣
//机中衣物的数量相等，则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：machines = [1,0,5]
//输出：3
//解释：
//第一步:    1     0 <-- 5    =>    1     1     4
//第二步:    1 <-- 1 <-- 4    =>    2     1     3
//第三步:    2     1 <-- 3    =>    2     2     2
//
//
// 示例 2：
//
//
//输入：machines = [0,3,0]
//输出：2
//解释：
//第一步:    0 <-- 3     0    =>    1     2     0
//第二步:    1     2 --> 0    =>    1     1     1
//
//
// 示例 3：
//
//
//输入：machines = [0,2,0]
//输出：-1
//解释：
//不可能让所有三个洗衣机同时剩下相同数量的衣物。
//
//
//
//
// 提示：
//
//
// n == machines.length
// 1 <= n <= 10⁴
// 0 <= machines[i] <= 10⁵
//
// Related Topics 贪心 数组 👍 141 👎 0

package main

import (
	"math"
	"testing"
)

func TestSuperWashingMachines(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
/*
将前 i 台洗衣机看成一组，记作 A，其余洗衣机看成另一组，记作 B。
设 sum[i]=∑j=0~i machines'[j]
若 sum[i] 为正则说明需要从 A 向 B 移入 sum[i] 件衣服，为负则说明需要从 B 向 A 移入 −sum[i] 件衣服。

我们分两种情况来考虑操作步数：

A 与 B 两组之间的衣服，最多需要 max i=0~n−1 ∣sum[i]∣ 次衣服移动；
组内的某一台洗衣机内的衣服数量过多，需要向左右两侧移出衣服，这最多需要 max i=0~n−1 machines'[i]次衣服移动。
取两者的最大值即为答案。
*/
func findMinMoves(machines []int) (max int) {
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%len(machines) != 0 {
		return -1
	}
	avl := sum / len(machines)
	sum = 0
	for i := range machines {
		machines[i] -= avl
		sum += machines[i]
		max = mMax(max, mMax(int(math.Abs(float64(sum))), machines[i]))
	}
	return
}
func mMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
