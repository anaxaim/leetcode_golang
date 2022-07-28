package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	q := []*TreeNode{root}
	var node *TreeNode
	for level := 0; len(q) > 0; level++ {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			node, q = q[0], q[1:]

			if len(result) <= level {
				result = append(result, []int{node.Val})
			} else {
				result[level] = append(result[level], node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

	}
	return result
}

func main() {
	/*
	   [3,9,20,null,null,15,7]
	   		3
	           /\
	          9  20
	             /\
	            15 7
	*/
	root := CreateNode(3)
	root.Left = CreateNode(9)
	root.Right = CreateNode(20)
	root.Right.Left = CreateNode(15)
	root.Right.Right = CreateNode(7)
	fmt.Println(levelOrder(root))
}
