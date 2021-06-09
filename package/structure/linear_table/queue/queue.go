package queue

type Node struct {
	data interface{}
	next *Node
}

type Point struct {
	tail, head *Node
	cnt        int
}

func Creat() *Point {
	p := new(Point)
	p.tail = new(Node)
	p.head = p.tail
	p.head.next = nil
	p.cnt = 0
	return p
}

func (p *Point) Push(data interface{}) {
	node := new(Node)
	node.data = data
	node.next = nil
	p.head.next = node
	p.head = node
	p.cnt++
}

func (p *Point) Pushs(data ...interface{}) {
	for _, v := range data {
		p.Push(v)
	}
}

func (p *Point) IsEmpty() bool {
	return p.cnt == 0
}

func (p *Point) Pop() interface{} {
	if p.IsEmpty() {
		return nil
	}
	data := p.tail.next.data
	p.tail = p.tail.next
	p.cnt--
	return data
}

func (p *Point) Count() int {
	return p.cnt
}

func (p *Point) Clear() {
	p.cnt = 0
	p.head = p.tail
	p.tail.next = nil
}

func (p *Point) Peek() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.tail.data
}
func (p *Point) Print() (int, []interface{}) {
	returnList := make([]interface{}, 0)
	for i := p.tail.next; i != nil; i = i.next {
		returnList = append(returnList, i.data)
	}
	return len(returnList), returnList
}
