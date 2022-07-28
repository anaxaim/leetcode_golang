package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// # 1
func middleNode1(head *ListNode) *ListNode {
	t := 0
	tmp := head

	for ; tmp != nil; t++ {
		tmp = tmp.Next
	}

	for i := 0; i < t/2; i++ {
		head = head.Next
	}

	return head
}

// # 2
func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	l := ListNode{}
	l.Next = &ListNode{Val: 1}
	l.Next = &ListNode{Val: 2}
	l.Next = &ListNode{Val: 3}
	l.Next = &ListNode{Val: 4}
	l.Next = &ListNode{Val: 5}

	l1 := middleNode2(&l)

	fmt.Printf("%T %p %v", l1, &l1, l1)
}
