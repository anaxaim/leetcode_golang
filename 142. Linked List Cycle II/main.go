package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {
	l := ListNode{}
	l.Next = &ListNode{Val: 3}
	l.Next = &ListNode{Val: 2}
	l.Next = &ListNode{Val: 0}
	l.Next = &ListNode{Val: -4}

	l1 := detectCycle(&l)

	fmt.Printf("%T %p %v", l1, &l1, l1)
}
