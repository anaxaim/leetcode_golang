package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := ListNode{}
	curr := &list3

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}

	return list3.Next
}

func main() {
	l1 := ListNode{}
	l1.Next = &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}

	l2 := ListNode{}
	l2.Next = &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next = &ListNode{Val: 4}

	l3 := mergeTwoLists(&l1, &l2)

	fmt.Printf("%T %p %v", l3, &l3, l3)
}
