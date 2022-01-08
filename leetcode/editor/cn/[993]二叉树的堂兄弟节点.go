//在二叉树中，根节点位于深度 0 处，每个深度为 K 的节点的子节点位于深度 K+1 处。
//
// 如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
//
// 我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
//
// 只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
//
//
//
// 示例 1：
//
//
//
//输入：root = [1,2,3,4], x = 4, y = 3
//输出：false
//
//
// 示例 2：
//
//
//
//输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
//输出：true
//
//
// 示例 3：
//
//
//
//
//输入：root = [1,2,3,null,4], x = 2, y = 3
//输出：false
//
//
//
// 提示：
//
//
// 二叉树的节点数介于 2 到 100 之间。
// 每个节点的值都是唯一的、范围为 1 到 100 的整数。
//
//
//
// Related Topics 树 广度优先搜索
// 👍 173 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/*
 * Definition for a binary tree myNode.
 */

type Node struct {
	node   *TreeNode
	level  int
	father int
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := make([]Node, 110)
	tail, head := 0, 1
	queue[tail].node = root
	queue[tail].level = 0
	queue[tail].father = -1
	var (
		flag   bool
		father int
		level  int
	)
	check := func(head int) (int, bool) {
		if queue[head].node != nil && (queue[head].node.Val == x || queue[head].node.Val == y) {
			if flag {
				if queue[head].level == level && queue[head].father != father {
					return 2, true
				} else {
					return 2, false
				}
			} else {
				flag = true
				father = queue[head].father
				level = queue[head].level
				return 1, true
			}
		}
		return 0, false
	}
	for tail < head {
		top := queue[tail]
		n, ok := check(tail)
		if n == 2 {
			return ok
		}
		if top.node.Left != nil {
			queue[head].node = top.node.Left
			queue[head].father = tail
			queue[head].level = top.level + 1
			head++

		}
		if top.node.Right != nil {
			queue[head].node = top.node.Right
			queue[head].father = tail
			queue[head].level = top.level + 1
			head++
			check(head)
		}
		tail++
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
