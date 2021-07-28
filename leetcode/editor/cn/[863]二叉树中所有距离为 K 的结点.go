//给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。 
//
// 返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
//输出：[7,4,1]
//解释：
//所求结点为与目标结点（值为 5）距离为 2 的结点，
//值分别为 7，4，以及 1
//
//
//
//注意，输入的 "root" 和 "target" 实际上是树上的结点。
//上面的输入仅仅是对这些对象进行了序列化描述。
// 
//
// 
//
// 提示： 
//
// 
// 给定的树是非空的。 
// 树上的每个结点都具有唯一的值 0 <= node.val <= 500 。 
// 目标结点 target 是树上的结点。 
// 0 <= K <= 1000. 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 
// 👍 309 👎 0
package main

import "fmt"

func main() {

}

//ERROR
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type PNode struct {
	val int
	k   int
	pos int //0,1,2
}

func distanceK(root *TreeNode, target *TreeNode, k int) (ret []int) {
	var thisH PNode
	var pSlice []PNode
	if root == target {
		thisH.k = 0
		thisH.pos = 0
	}
	node := PNode{
		val: root.Val,
		k:   0,
		pos: 1,
	}
	pSlice = append(pSlice, node)
	var dfs func(root *TreeNode, k int, pos int)
	dfs = func(root *TreeNode, k int, pos int) {
		if root == nil {
			return
		}
		if root == target {
			thisH.k = k
			thisH.pos = pos
		}
		node := PNode{
			val: root.Val,
			k:   k,
			pos: pos,
		}
		pSlice = append(pSlice, node)
		dfs(root.Left, k+1, pos)
		dfs(root.Right, k+1, pos)
	}
	dfs(root.Left, 1, 0)
	dfs(root.Right, 1, 2)
	fmt.Println(pSlice)
	fmt.Println(thisH)
	for _, v := range pSlice {
		if v.pos == thisH.pos && (v.k-thisH.k == k || thisH.k-v.k == k) {
			ret = append(ret, v.val)
		} else if v.pos != thisH.pos && (v.k+thisH.k == k) {
			ret = append(ret, v.val)
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
