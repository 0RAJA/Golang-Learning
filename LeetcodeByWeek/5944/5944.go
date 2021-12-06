package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) (ret string) {
	var now, end *TreeNode
	var front func(root *TreeNode, path string)
	startPath, endPath := "", ""
	front = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		if root.Val == startValue {
			startPath = path
			now = root
		}
		if root.Val == destValue {
			endPath = path
			end = root
		}
		if root.Left != nil {
			front(root.Left, path+"L")
		}
		if root.Right != nil {
			front(root.Right, path+"R")
		}
	}
	front(root, "")
	var index int
	isOK := false
	for i := range startPath {
		if i >= len(endPath) {
			break
		}
		if startPath[i] != endPath[i] {
			index = i
			isOK = true
			break
		}
	}
	if !isOK && len(startPath) > len(endPath) {
		index = len(endPath)
	}
	for i := index; i < len(startPath); i++ {
		ret += "U"
	}
	if index < len(endPath) {
		ret += endPath[index:]
	}
	return
}
