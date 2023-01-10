package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	tr1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	tr2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTree(tr1, tr2)) // true

	tr3 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	tr4 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(tr3, tr4)) // false

	tr5 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	tr6 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(tr5, tr6)) // false
}
