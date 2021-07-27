//给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一
//个。 
//
// 更正式地说，root.val = min(root.left.val, root.right.val) 总成立。 
//
// 给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [2,2,5,null,null,5,7]
//输出：5
//解释：最小的值是 2 ，第二小的值是 5 。
// 
//
// 示例 2： 
//
// 
//输入：root = [2,2,2]
//输出：-1
//解释：最小的值是 2, 但是不存在第二小的值。
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [1, 25] 内 
// 1 <= Node.val <= 231 - 1 
// 对于树中每个节点 root.val == min(root.left.val, root.right.val) 
// 
// Related Topics 树 深度优先搜索 二叉树 
// 👍 159 👎 0
package main

func main() {

}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode1 struct {
 *     Val int
 *     Left *TreeNode1
 *     Right *TreeNode1
 * }
 */
func findSecondMinimumValue(root *TreeNode1) int {
	if root.Left == nil {
		return -1
	}
	a := find(root.Left, root.Val)
	b := find(root.Right, root.Val)
	return IsOK(a, b)
}

func find(root *TreeNode1, target int) int {
	if root.Left == nil || root.Val != target {
		if root.Val != target {
			return root.Val
		}
		return -1
	}
	a := find(root.Left, target)
	b := find(root.Right, target)
	return IsOK(a, b)
}

func IsOK(a, b int) int {
	if a < 0 && b < 0 {
		return -1
	}
	if a > 0 && b < 0 {
		return a
	}
	if a < 0 && b > 0 {
		return b
	}
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
