package stack

// Node 栈结构体
type Node struct {
	data interface{}
	next *Node
}

// Point 栈管理器结构体
type Point struct {
	top *Node
	cnt int
}

// Creat 创建栈
func Creat() *Point {
	p := new(Point)
	p.top, p.cnt = nil, 0
	return p
}

// Push 入栈
func (p *Point) Push(data interface{}) {
	p.cnt++
	node := new(Node)
	node.data = data
	node.next = p.top
	p.top = node
}

func (p *Point) Pushs(data ...interface{}) {
	for _, v := range data {
		p.Push(v)
	}
}

// IsEmpty 判空
func (p *Point) IsEmpty() bool {
	return p.cnt == 0
}

// Pop 弹出栈首元素
func (p *Point) Pop() interface{} {
	if p.IsEmpty() {
		return nil
	}
	data := p.top.data
	p.top = p.top.next
	p.cnt--
	return data
}

// Clear 清空
func (p *Point) Clear() {
	p.top = nil
	p.cnt = 0
}

// Print 返回栈中元素和它的数量
func (p *Point) Print() (int, []interface{}) {
	returnList := make([]interface{}, 0)
	for i := p.top; i != nil; i = i.next {
		returnList = append(returnList, i.data)
	}
	return len(returnList), returnList
}

// Count 返回栈中元素的数量
func (p *Point) Count() int {
	return p.cnt
}

// Peek 获取栈顶元素
func (p *Point) Peek() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.top.data
}
