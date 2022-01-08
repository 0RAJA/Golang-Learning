//有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， w
//i 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, K = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：times = [[1,2,1]], n = 2, K = 1
//输出：1
//
//
// 示例 3：
//
//
//输入：times = [[1,2,1]], n = 2, K = 2
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= K <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
//
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列）
// 👍 295 👎 0
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
const MAX = 101

type LNode struct {
	index  int
	weight int
	next   *LNode
}
type Vex struct {
	first *LNode
}
type AdjList struct {
	list            []Vex
	vexNum, edgeNum int
}

func networkDelayTime(times [][]int, n int, k int) int {
	adjList := AdjList{
		list:    make([]Vex, n+1),
		vexNum:  n,
		edgeNum: len(times),
	}
	book := make(map[int]bool, n)
	for i := 0; i < len(times); i++ {
		x, y, w := times[i][0], times[i][1], times[i][2]
		node := LNode{
			index:  y,
			weight: w,
		}
		node.next = adjList.list[x].first
		adjList.list[x].first = &node
	}
	dis := make([]int, n+1)
	for i := 1; i < len(dis); i++ {
		dis[i] = MAX
	}
	dis[k] = 0
	book[k] = true
	for p := adjList.list[k].first; p != nil; p = p.next {
		dis[p.index] = p.weight
	}
	for {
		min := MAX
		t := -1
		for i := 1; i <= adjList.vexNum; i++ {
			if book[i] == false && dis[i] < min {
				min = dis[i]
				t = i
			}
		}
		if t == -1 {
			break
		}
		book[t] = true
		for p := adjList.list[t].first; p != nil; p = p.next {
			if dis[p.index] > dis[t]+p.weight {
				dis[p.index] = dis[t] + p.weight
			}
		}
	}
	max := 0
	for i := 1; i < len(dis); i++ {
		if dis[i] == MAX {
			return -1
		}
		if max < dis[i] {
			max = dis[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
