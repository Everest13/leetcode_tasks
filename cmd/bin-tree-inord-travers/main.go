package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
94. Binary Tree Inorder Traversal. Easy.
https://leetcode.com/problems/binary-tree-inorder-traversal/description/
Runtime 2ms, beats 64,32%; Memory 2mb, beats 99.8%.
 */
func main() {
	head := &TreeNode{
		Val: 1,
	}

	current := head
	//to right
	for i:=0; i < 3; i++ {
		next := &TreeNode{
			Val: i,
		}
		current.Right = next
		current = next
	}

	result := inorderTraversal(head)

	fmt.Println("Result is ", result)
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	extrVal := func(root *TreeNode) {}
	extrVal = func(root *TreeNode) {
		for root != nil {
			if root.Right != nil {
				extrVal(root.Right)
			}

			result = append(result, root.Val)

			if root.Left != nil {
				extrVal(root.Left)
			}
			break
		}
	}

	extrVal(root)

	return result
}
