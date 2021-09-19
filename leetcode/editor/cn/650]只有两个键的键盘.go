//最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：
//
//
// Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
// Paste（粘贴）：粘贴 上一次 复制的字符。
//
//
// 给你一个数字 n ，你需要使用最少的操作次数，在记事本上输出 恰好 n 个 'A' 。返回能够打印出 n 个 'A' 的最少操作次数。
//
//
//
// 示例 1：
//
//
//输入：3
//输出：3
//解释：
//最初, 只有一个字符 'A'。
//第 1 步, 使用 Copy All 操作。
//第 2 步, 使用 Paste 操作来获得 'AA'。
//第 3 步, 使用 Paste 操作来获得 'AAA'。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
// Related Topics 数学 动态规划 👍 393 👎 0
package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSteps(3))
}

//leetcode submit region begin(Prohibit modification and deletion)
//dp
/*
设 f[i] 表示打印出 i 个 A 的最少操作次数。
由于我们只能使用「复制全部」和「粘贴」两种操作，那么要想得到 i 个 A，我们必须首先拥有 j 个 A，
使用一次「复制全部」操作，再使用若干次「粘贴」操作得到 i 个 A。
因此，这里的 j 必须是 i 的因数，
「粘贴」操作的使用次数即为i/j−1。
我们可以枚举 j 进行状态转移，这样就可以得到状态转移方程：
f[i]=min(f[j]+j/i)   (i|j)
其中 j∣i 表示 j 可以整除 i，即 j 是 i 的因数。
动态规划的边界条件为 f[1]=0，最终的答案即为 f[n]。

func minSteps(n int) (cnt int) {
	f := make([]int, n+1)
	f[1] = 0
	for i := 2; i <= n; i++ {
		f[i] = math.MaxInt64
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				f[i] = min(f[i], f[j]+i/j)
				f[i] = min(f[i], f[i/j]+j)
			}
		}
	}
	return f[n]
}
*/

//优雅分解质因数
/*
f[i] = min(f[i/j]+j) (i|j)
n>1时 其实就是将n分解为m个数字的乘积 且m个数字的和最小 即把一个数分解为n个质数的和 从小到大的去试探
*/
func minSteps(n int) (cnt int) {
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			cnt += i
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
