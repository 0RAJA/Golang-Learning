//给定一个 N 叉树，找到其最大深度。
//
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
// N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：3
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：5
//
//
//
//
// 提示：
//
//
// 树的深度不会超过 1000 。
// 树的节点数目位于 [0, 10⁴] 之间。
//
// Related Topics 树 深度优先搜索 广度优先搜索 👍 230 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Children []*Node
* }
 */

func maxDepth() int {
	type Node struct {
		Val      int
		Children []*Node
	}
	var root = new(Node)
	var maxDepth func(root *Node) (ret int)
	maxDepth = func(root *Node) (ret int) {
		if root == nil {
			return 0
		}
		for _, p := range root.Children {
			if k := maxDepth(p); k > ret {
				ret = k
			}
		}
		return ret + 1
	}
	return maxDepth(root)
}

//leetcode submit region end(Prohibit modification and deletion)
