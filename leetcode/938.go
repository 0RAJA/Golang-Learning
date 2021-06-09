package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ranges(root *TreeNode, low, high int) int {
	sum := 0
	if root == nil {
		return sum
	} else if root.Val >= low && root.Val <= high {
		sum = ranges(root.Left, low, high) + ranges(root.Right, low, high) + root.Val
	} else if root.Val < low {
		sum = ranges(root.Right, low, high)
	} else {
		sum = ranges(root.Left, low, high)
	}
	return sum
}
func rangeSumBST(root *TreeNode, low int, high int) int {
	return ranges(root, low, high)
}
