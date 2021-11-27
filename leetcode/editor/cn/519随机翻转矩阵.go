//给你一个 m x n 的二元矩阵 matrix ，且所有值被初始化为 0 。请你设计一个算法，随机选取一个满足 matrix[i][j] == 0 的下标
//(i, j) ，并将它的值变为 1 。所有满足 matrix[i][j] == 0 的下标 (i, j) 被选取的概率应当均等。
//
// 尽量最少调用内置的随机函数，并且优化时间和空间复杂度。
//
// 实现 Solution 类：
//
//
// Solution(int m, int n) 使用二元矩阵的大小 m 和 n 初始化该对象
// int[] flip() 返回一个满足 matrix[i][j] == 0 的随机下标 [i, j] ，并将其对应格子中的值变为 1
// void reset() 将矩阵中所有的值重置为 0
//
//
//
//
// 示例：
//
//
//输入
//["Solution", "flip", "flip", "flip", "reset", "flip"]
//[[3, 1], [], [], [], [], []]
//输出
//[null, [1, 0], [2, 0], [0, 0], null, [2, 0]]
//
//解释
//Solution solution = new Solution(3, 1);
//solution.flip();  // 返回 [1, 0]，此时返回 [0,0]、[1,0] 和 [2,0] 的概率应当相同
//solution.flip();  // 返回 [2, 0]，因为 [1,0] 已经返回过了，此时返回 [2,0] 和 [0,0] 的概率应当相同
//solution.flip();  // 返回 [0, 0]，根据前面已经返回过的下标，此时只能返回 [0,0]
//solution.reset(); // 所有值都重置为 0 ，并可以再次选择下标返回
//solution.flip();  // 返回 [2, 0]，此时返回 [0,0]、[1,0] 和 [2,0] 的概率应当相同
//
//
//
// 提示：
//
//
// 1 <= m, n <= 10⁴
// 每次调用flip 时，矩阵中至少存在一个值为 0 的格子。
// 最多调用 1000 次 flip 和 reset 方法。
//
// Related Topics 水塘抽样 哈希表 数学 随机化 👍 67 👎 0

package main

import (
	"math/rand"
	time2 "time"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type SolutionA struct {
	sum, m, n int
	hash      map[int]int
}

func ConstructorA(m int, n int) SolutionA {
	return SolutionA{
		sum:  m * n,
		m:    m,
		n:    n,
		hash: make(map[int]int),
	}
}

func (this *SolutionA) Flip() []int {
	rand.Seed(time2.Now().UnixNano())
	idx := rand.Intn(this.sum)
	this.sum--
	if k, ok := this.hash[idx]; ok {
		idx = k
	}
	x := idx % this.m
	y := idx / this.m
	this.hash[idx] = this.sum
	return []int{x, y}
}

func (this *SolutionA) Reset() {
	this.hash = make(map[int]int)
	this.sum = this.m * this.n
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
//leetcode submit region end(Prohibit modification and deletion)
