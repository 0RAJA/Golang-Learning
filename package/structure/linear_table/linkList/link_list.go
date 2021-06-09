package linkList

import (
	"errors"
)

type Linker interface {
	Lister
	Deleter
	Moder
	Finder
}

type Deleter interface { //需要删除的数据
	DeleteCmp() bool //删除比较方法
}

type Moder interface { //需要修改的数据
	ModCmp() bool //修改比较方法
	Mod()         //修改方法
}

type Finder interface {
	FindCmp() bool //查找方法
}

type Lister interface{} //仅仅储存数据

// Node 链表节点结构体--data为所存储的数据
type Node struct {
	data Linker
	next *Node
}

// Point 链表管理器
type Point struct {
	head *Node
	cnt  int
}

// Creat 创建空链表并返回链表管理器
func Creat() *Point {
	p := new(Point)
	p.head = new(Node)
	p.head.next, p.cnt = nil, 0
	return p
}

//Add 增加一个data数据
func (p *Point) Add(data Linker) {
	inode := new(Node)
	inode.data = data
	inode.next = p.head.next
	p.head.next = inode
	p.cnt++
}

//Adds 增加多个数据
func (p *Point) Adds(data ...Linker) {
	for _, v := range data {
		p.Add(v)
	}
}

func del(deleter Deleter) bool {
	return deleter.DeleteCmp()
}

// Delete 从链表中删除指定元素,如果删除成功error为nil,否则error不为nil
//f 为判断是否是指定元素的函数
func (p *Point) Delete() error {
	if p.Empty() {
		return errors.New("empty")
	}
	for i, j := p.head.next, p.head; i != nil; j, i = i, i.next {
		if del(i.data) {
			j.next = i.next
			return nil
		}
	}
	return errors.New("no find")
}

func find(finder Finder) bool {
	return finder.FindCmp()
}

//Find 从链表中返回指定的元素,如果没找到 error不为nil
//f 为判断是否是指定元素的函数
func (p *Point) Find() (interface{}, error) {
	for i := p.head.next; i != nil; i = i.next {
		if find(i.data) {
			return i.data, nil
		}
	}
	return nil, errors.New("no find")
}

func mod(moder Moder) bool {
	return moder.ModCmp()
}

//Mod 修改指定元素数据,修改成功返回true,否则返回false
//f1 为判断是否是指定元素的函数
//f2 为修改指定元素数据的函数
func (p *Point) Mod() bool {
	for i := p.head.next; i != nil; i = i.next {
		if mod(i.data) {
			i.data.Mod()
			return true
		}
	}
	return false
}

//Empty 链表为空返回true,链表不为空返回false
func (p *Point) Empty() bool {
	return p.cnt == 0
}

//Reverse 逆转链表
func (p *Point) Reverse() {
	i, j := p.head.next, p.head
	p.head.next = nil
	for i != nil {
		j, i = i, i.next
		j.next = p.head.next
		p.head.next = j
	}
}

//ReverseD 递归逆转链表
func (p *Point) ReverseD() *Node {
	if p.head == nil || p.head.next == nil {
		return p.head
	}
	newHead := p.ReverseD()
	t := p.head.next
	t.next = p.head
	p.head.next = nil
	return newHead
}

//Print 将对应元素作为切片返回,返回存储链表所有元素的切片list和元素个数cnt
func (p *Point) Print() ([]interface{}, int) {
	cnt := 0
	list := make([]interface{}, 0)
	for i := p.head.next; i != nil; i = i.next {
		list = append(list, i.data)
		cnt++
	}
	return list, cnt
}
