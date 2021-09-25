package main

type node struct {
	val  int
	next *node
}

type tree struct {
	val         int
	left, right *tree
	isOK        bool
}

func main() {

}

func reserve(head *node) {
	p := head
	q := p.next
	for q != nil {
		t := q.next
		q.next = p
		q = t
	}
	head.next.next = nil
	head.next = p
}

func insert(n **tree, val int) {
	if n == nil {
		*n = &tree{
			val:   val,
			left:  nil,
			right: nil,
		}
	} else {
		if val < (*n).val {
			insert(&(*n).left, val)
		} else {
			insert(&(*n).right, val)
		}
	}
}

func treeNum(n *tree, cnt int) int {
	if n == nil {
		return cnt
	}
	return max(treeNum(n.left, cnt+1), treeNum(n.right, cnt+1))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func treeNum2(n *tree) int {
	stack := []*tree{n}
	max := 1
	cnt := 1
Loop:
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		if p.left != nil && p.left.isOK == false {
			p.left.isOK = true
			stack = append(stack, p.left)
			cnt++
			continue Loop
		} else if p.right != nil && p.right.isOK == false {
			p.right.isOK = true
			stack = append(stack, p.right)
			cnt++
			continue Loop
		} else {
			if max < cnt {
				max = cnt
			}
			stack = stack[:len(stack)-1]
			cnt--
		}
	}
	return cnt
}
