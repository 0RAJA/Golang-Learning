//在有向图中，以某个节点为起始节点，从该点出发，每一步沿着图中的一条有向边行走。如果到达的节点是终点（即它没有连出的有向边），则停止。
//
// 对于一个起始节点，如果从该节点出发，无论每一步选择沿哪条有向边行走，最后必然在有限步内到达终点，则将该起始节点称作是 安全 的。
//
// 返回一个由图中所有安全的起始节点组成的数组作为答案。答案数组中的元素应当按 升序 排列。
//
// 该有向图有 n 个节点，按 0 到 n - 1 编号，其中 n 是 graph 的节点数。图以下述形式给出：graph[i] 是编号 j 节点的一个列表，
//满足 (i, j) 是图的一条有向边。
//
//
//
//
//
// 示例 1：
//
//
//输入：graph = [[1,2],[2,3],[5],[0],[5],[],[]]
//输出：[2,4,5,6]
//解释：示意图如上。
//
//
// 示例 2：
//
//
//输入：graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]
//输出：[4]
//
//
//
//
// 提示：
//
//
// n == graph.length
// 1 <= n <= 104
// 0 <= graph[i].length <= n
// graph[i] 按严格递增顺序排列。
// 图中可能包含自环。
// 图中边的数目在范围 [1, 4 * 104] 内。
//
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序
// 👍 159 👎 0
package main

import (
	"fmt"
)

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println(eventualSafeNodes(graph))
}

//leetcode submit region begin(Prohibit modification and deletion)
func eventualSafeNodes(graph [][]int) (slice []int) {
	cntMap := make(map[int]int, len(graph)) //记录出度
	linkMap := make(map[int][]int)          //记录连接者
	ret := make([]int, 0, len(graph))       //队列
	for i := 0; i < len(graph); i++ {
		cntMap[i] = len(graph[i])
		for j := 0; j < len(graph[i]); j++ {
			linkMap[graph[i][j]] = append(linkMap[graph[i][j]], i)
		}
		if cntMap[i] == 0 {
			ret = append(ret, i)
		}
	}
	for index := 0; ; index++ {
		if index >= len(ret) {
			break
		}
		for i := 0; i < len(linkMap[ret[index]]); i++ {
			cntMap[linkMap[ret[index]][i]]--
			if cntMap[linkMap[ret[index]][i]] == 0 {
				ret = append(ret, linkMap[ret[index]][i])
			}
		}
	}
	slice = make([]int, 0, len(graph))
	for i := 0; i < len(graph); i++ {
		if cntMap[i] == 0 {
			slice = append(slice, i)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
