package _944

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	root *TreeNode
	path string
}

func getDirections(root *TreeNode, startValue int, destValue int) (ret string) {
	father := map[*TreeNode]*TreeNode{}
	var now, end *TreeNode
	var front func(root *TreeNode)
	front = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == startValue {
			now = root
		}
		if root.Val == destValue {
			end = root
		}
		if root.Left != nil {
			father[root.Left] = root
			front(root.Left)
		}
		if root.Right != nil {
			father[root.Right] = root
			front(root.Right)
		}
	}
	front(root)
	iMap := make(map[*TreeNode]bool)
	stack := []*Node{}
	stack = append(stack, &Node{
		root: now,
		path: "",
	})
	iMap[now] = true
	for len(stack) > 0 {
		p := stack[0]
		stack = stack[1:]
		now := p.root
		if x := now.Left; x != nil && !iMap[x] {
			t := &Node{
				root: x,
				path: p.path + "L",
			}
			stack = append(stack, t)
			iMap[x] = true
			if x == end {
				return t.path
			}
		}
		if x := now.Right; x != nil && !iMap[x] {
			t := &Node{
				root: x,
				path: p.path + "R",
			}
			stack = append(stack, t)
			iMap[x] = true
			if x == end {
				return t.path
			}
		}
		if x := father[now]; x != nil && !iMap[x] {
			t := &Node{
				root: x,
				path: p.path + "U",
			}
			stack = append(stack, t)
			iMap[x] = true
			if x == end {
				return t.path
			}
		}
	}
	return
}
