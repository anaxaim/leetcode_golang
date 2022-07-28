package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}

func main() {
	l := ListNode{}
	l.Next = &ListNode{Val: 1}
	l.Next = &ListNode{Val: 2}
	l.Next = &ListNode{Val: 3}
	l.Next = &ListNode{Val: 4}
	l.Next = &ListNode{Val: 5}

	newL := reverseList(&l)

	fmt.Printf("%T %p %v", newL, &newL, newL)
}
