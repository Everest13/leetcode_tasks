package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}

	res := bstToGst(root)
	fmt.Println(res)
}

/**
	1038. Binary Search Tree to Greater Sum Tree
	Medium
	https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/
 */
func bstToGst(root *TreeNode) *TreeNode {
	nodes := []*TreeNode{}
	recursion := func(node *TreeNode) {}
	recursion = func(node *TreeNode) {
		right := node.Right
		left := node.Left

		if right != nil {
			recursion(right)
		}

		nodes = append(nodes, node)

		if left != nil {
			recursion(left)
		}
	}

	recursion(root)

	sumNodes := 0
	for _, node := range nodes {
		val := node.Val
		node.Val += sumNodes
		sumNodes += val
	}

	return root
}
