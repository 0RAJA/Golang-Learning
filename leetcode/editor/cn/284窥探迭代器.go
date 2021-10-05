//请你设计一个迭代器，除了支持 hasNext 和 next 操作外，还支持 peek 操作。
//
// 实现 PeekingIterator 类：
//
//
//
//
// PeekingIterator(int[] nums) 使用指定整数数组 nums 初始化迭代器。
// int next() 返回数组中的下一个元素，并将指针移动到下个元素处。
// bool hasNext() 如果数组中存在下一个元素，返回 true ；否则，返回 false 。
// int peek() 返回数组中的下一个元素，但 不 移动指针。
//
//
//
//
// 示例：
//
//
//输入：
//["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
//[[[1, 2, 3]], [], [], [], [], []]
//输出：
//[null, 1, 2, 2, 3, false]
//
//解释：
//PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
//peekingIterator.next();    // 返回 1 ，指针移动到下一个元素 [1,2,3]
//peekingIterator.peek();    // 返回 2 ，指针未发生移动 [1,2,3]
//peekingIterator.next();    // 返回 2 ，指针移动到下一个元素 [1,2,3]
//peekingIterator.next();    // 返回 3 ，指针移动到下一个元素 [1,2,3]
//peekingIterator.hasNext(); // 返回 False
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 1000
// 对 next 和 peek 的调用均有效
// next、hasNext 和 peek 最多调用 1000 次
//
//
//
//
//
//
// 进阶：你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？
// Related Topics 设计 数组 迭代器 👍 101 👎 0

package main

func main() {

}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 1
}

//leetcode submit region begin(Prohibit modification and deletion)
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
type node struct {
	val  interface{}
	next *node
}

type PeekingIterator struct {
	cur *node
}

func mConstructor(iter *Iterator) (h *PeekingIterator) {
	h = new(PeekingIterator)
	h.cur = new(node)
	tail := h.cur
	for iter.hasNext() {
		p := &node{val: iter.next()}
		tail.next = p
		tail = p
	}
	return h
}

func (this *PeekingIterator) hasNext() bool {
	return this.cur.next != nil
}

func (this *PeekingIterator) next() (ret int) {
	ret = this.cur.next.val.(int)
	this.cur = this.cur.next
	return
}

func (this *PeekingIterator) peek() int {
	return this.cur.next.val.(int)
}

//leetcode submit region end(Prohibit modification and deletion)
