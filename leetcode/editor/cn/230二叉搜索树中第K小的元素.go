//给定一个二叉搜索树的根节点 root ，和一个整数 K ，请你设计一个算法查找其中第 K 个最小元素（从 1 开始计数）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,1,4,null,2], K = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [5,3,6,2,4,null,null,1], K = 3
//输出：3
//
//
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n 。
// 1 <= K <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 K 小的值，你将如何优化算法？
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 486 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) (ret int) {
	var dfs func(root *TreeNode)
	cnt := 0
	isFind := false
	dfs = func(root *TreeNode) {
		if isFind || root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		cnt++
		if cnt == k {
			isFind = true
			ret = root.Val
		} else if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
