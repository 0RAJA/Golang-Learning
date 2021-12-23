package main

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

const (
	maxLevel    int     = 18 //设置最大为18层，可以视情况调节
	probability float64 = 1 / math.E
)

type (
	//这里定义跳表用到的数据结构，可以通过下方的图看得更直观一些

	// Node
	//跳表的节点，这个节点会是一个element切片
	//切片的大小代表着一个节点有多少个索引
	//头结点有18级
	Node struct {
		next []*Element
	}

	// Element
	//存储数据的元素
	//每个元素有key和value更重要的是有下一个节点的指针
	Element struct {
		Node
		key   []byte
		value interface{}
	}

	// SkipList
	//Node指向了跳表节点的开头
	//其余则是跳表需要用到的一些信息，碰到了再进行说明
	SkipList struct {
		Node               //next切片
		maxLevel       int //最大层数
		Len            int //跳表长度
		randSource     rand.Source
		probability    float64
		probTable      []float64 //每层建索引的概率
		prevNodesCache []*Node   //存储经过的最后一个节点
	}
)

// NewSkipList
// 创建跳表
func NewSkipList() *SkipList {
	return &SkipList{
		Node:           Node{next: make([]*Element, maxLevel)},
		prevNodesCache: make([]*Node, maxLevel),
		maxLevel:       maxLevel,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:    probability,
		probTable:      probabilityTable(probability, maxLevel),
	}
}

// Key 获得element的key
func (e *Element) Key() []byte {
	return e.key
}

// Value 获得element的value
func (e *Element) Value() interface{} {
	return e.value
}

// SetValue 设置element的value
func (e *Element) SetValue(val interface{}) {
	e.value = val
}

// Next
//跳表的第一级（索引为0）是一个完整的单向链表
//可以从这个单链表中索引所有的元素
//一个元素是element，他有指向下一个节点的指针next(element中内嵌的Node)
//下一节点的指针其实就是一个指向18级element的切片
//第一级（索引为0）就是这个element的下一个element
func (e *Element) Next() *Element {
	return e.next[0]
}

// Front
// 获取跳表的第一层的第一个element
// 有了这个element通过next就可以访问所有的元素
// Front first element.
// Get the head element of skl, and get all data by traversing backward.
//	e := list.Front()
//	for p := e; p != nil; p = p.Next() {
//		//do something with Element p
//	}
func (t *SkipList) Front() *Element {
	return t.next[0]
}

// Get find value by the key, returns nil if not found.
func (t *SkipList) Get(key []byte) *Element {
	//一开始看prev和next两个变量的类型不同会迷惑
	//其实按照变量名来理解就好
	//prev就是当前指针（相对于next是前一个）
	//next是下一个指针
	var prev = &t.Node
	var next *Element

	//从最高层，即第18层开始
	//这一层循环表示从高级（level大的一端）向低级（level小的一端）移动
	for i := t.maxLevel - 1; i >= 0; i-- {
		//一开始next 指向Node节点的第18级
		next = prev.next[i]

		//如果目标的key比下个的key大，则向右移动
		//这一层循环表示向右移动（key大的方向）
		for next != nil && bytes.Compare(key, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
	}

	//在两个循环结束以后，指针已经到达了第一层，不能再继续往低层移动了
	//并且目标key不比右边的key大了（即key <= next.key）
	//如果（next.key<=key）即key == next.key那就找到了目标
	if next != nil && bytes.Compare(next.key, key) <= 0 {
		return next
	}
	//否则返回nil表示目标不存在
	return nil
}

// Exist 调用get函数，只要返回不为nil即表明目标存在跳表中
func (t *SkipList) Exist(key []byte) bool {
	return t.Get(key) != nil
}

//这个函数与前面的get类似，不过不同的是会记录寻找过程中经过的元素
//因为对链表的某个地方插入或者删除一个元素
//需要获得前一个元素的指针
func (t *SkipList) backNodes(key []byte) []*Node {
	var prev = &t.Node
	var next *Element

	//记录寻找节点中所经过的节点
	//在跳表插入过程中，每一层都有可能需要插入索引
	//所以记录寻找过程中每一层的最终节点，方便后续插入删除
	prevs := t.prevNodesCache

	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		//寻找过程与get保持一致
		for next != nil && bytes.Compare(key, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
		//记录寻找过程
		prevs[i] = prev
	}

	return prevs
}

// Remove 通过key来删除元素
func (t *SkipList) Remove(key []byte) *Element {
	//寻找key，并记录寻找过程每层经过的最后一个节点
	prev := t.backNodes(key)
	//如果寻找过程中得到的节点中
	//第一层的节点的下一个元素不为空
	//并且key与目标key相等
	//那么就是找到了要删除的元素
	if element := prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		//k是层数，v是这个要被删除元素在第k层指向的下一个元素
		for k, v := range element.next {
			//将缓存在第k层找到的元素的第k层索引指向
			//被删除元素的第k层指向的下一个元素
			prev[k].next[k] = v
		}
		//跳表长度减一
		t.Len--
		//返回被删除元素
		return element
	}
	//如果下一个key与目标key不相等
	//即不存在目标元素，返回空
	return nil
}

// Put 向跳表增加一个元素，如果已经存在，则用新的value替代
func (t *SkipList) Put(key []byte, value interface{}) *Element {
	var element *Element
	//寻找key，并记录寻找过程每层经过的最后一个节点
	prev := t.backNodes(key)
	//如果第1层的下一个节点与要插入的key相等
	//那么修改value即可
	if element = prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		element.value = value
		return element
	}
	//否则，新建一个元素用作插入
	element = &Element{
		Node: Node{
			//创造一个随机层数的element切片
			//注意并不是每一个元素都有18层级别
			//切片的大小对应有多少个索引
			next: make([]*Element, t.randomLevel()),
		},
		key:   key,
		value: value,
	}
	//创造的这个切片有多少层，就建立多少个索引
	//即让新建元素指向缓存指向的下一个元素
	//让缓存的节点指向新建的元素即可
	for i := range element.next {
		element.next[i] = prev[i].next[i]
		prev[i].next[i] = element
	}

	t.Len++
	return element
}

//一个简单的数学概率问题，假设向上不建立索引的概率为0.4
//（代码实际的概率为1/e）
//首先第一层（索引为0）是一定会插入的
//从第二层次（索引为1）开始，向上不建立索引的概率为0.4^1
//第n层（索引为n-1）则是0.4^（n-1）
//这个函数就是建立了这样一个表。之后只需要查表插入第几层只需要一次随机数
//然后拿着这个随机数对表比较即可，随机数比表的数小，就插入
//比表的数大则不插入
func probabilityTable(probability float64, maxLevel int) (table []float64) {
	for i := 1; i <= maxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}

//随机决定生成多少层索引
// generate random index level.
func (t *SkipList) randomLevel() (level int) {
	//生成一个随机数
	r := float64(t.randSource.Int63()) / (1 << 63)
	//第一层一定生成
	level = 1
	//如果没到最高层并且随机数比表对应层的概率小，则新建索引
	for level < t.maxLevel && r < t.probTable[level] {
		level++
	}
	return
}

func main() {
	t := NewSkipList()
	t.Put([]byte("111"), 111)
}
