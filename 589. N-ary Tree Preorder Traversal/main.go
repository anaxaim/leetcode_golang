package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func getTreeNode(key int) *Node {
	var n *Node = &Node{}
	n.Val = key
	n.Children = make([]*Node, 0)
	return n
}

func (n *Node) addChild(key int) {
	var t *Node = getTreeNode(key)
	n.Children = append(n.Children, t)
}

type NAryTree struct {
	root *Node
}

func getNAryTree() *NAryTree {
	return &NAryTree{nil}
}

func preorder(root *Node) []int {
	sl := make([]int, 0)

	if root == nil {
		return sl
	}

	var pop *Node
	stack := []*Node{root}
	for len(stack) > 0 {
		pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
		sl = append(sl, pop.Val)

		for i := len(pop.Children) - 1; i >= 0; i-- {
			stack = append(stack, pop.Children[i])
		}
	}
	return sl
}

func main() {
	var tree *NAryTree = getNAryTree()
	tree.root = getTreeNode(1)
	tree.root.addChild(3)
	tree.root.addChild(2)
	tree.root.addChild(4)
	tree.root.Children[0].addChild(5)
	tree.root.Children[0].addChild(6)

	fmt.Println(preorder(tree.root))
}
