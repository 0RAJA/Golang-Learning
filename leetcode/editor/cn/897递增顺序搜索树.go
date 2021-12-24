//给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
//
//
//
// 示例 1：
//
//
//输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
//输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//
//
// 示例 2：
//
//
//输入：root = [5,1,7]
//输出：[1,null,5,null,7]
//
//
//
//
// 提示：
//
//
// 树中节点数的取值范围是 [1, 100]
// 0 <= Node.val <= 1000
//
// Related Topics 栈 树 深度优先搜索 二叉搜索树 二叉树 👍 254 👎 0

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
func increasingBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) []*TreeNode
	dfs = func(root *TreeNode) (ret []*TreeNode) {
		if root == nil {
			return nil
		}
		s1 := append(dfs(root.Left), root)
		s2 := dfs(root.Right)
		return append(s1, s2...)
	}
	ret := dfs(root)
	for i := 1; i < len(ret); i++ {
		ret[i-1].Right = ret[i]
		ret[i-1].Left = nil
	}
	return ret[0]
}

//leetcode submit region end(Prohibit modification and deletion)
