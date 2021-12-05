package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	m := &ListNode{Next: head}
	q := head
	for q != nil && q.Next != nil {
		q = q.Next.Next
		m = m.Next
	}
	if m.Next == head {
		return nil
	}
	m.Next = m.Next.Next
	return head
}
