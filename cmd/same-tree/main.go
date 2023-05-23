package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
100. Same Tree. Easy.
https://leetcode.com/problems/same-tree/description/
Runtime 2ms, beats 61,77%; Memory 2.1mb, beats 92.11%.
 */
func main() {
	head1 := &TreeNode{
		Val: 1,
	}

	current := head1
	for i:=2; i < 3; i++ {
		next := &TreeNode{
			Val: i,
		}
		current.Right = next
		current = next
	}

	head2Left := &TreeNode{
		Val: 2,
		Right: nil,
		Left: nil,
	}

	head2 := &TreeNode{
		Val: 1,
		Right: nil,
		Left: head2Left,
	}

	//current = head2
	//for i:=2; i < 4; i++ {
	//	next := &TreeNode{
	//		Val: i,
	//	}
	//	current.Right = next
	//	current = next
	//}

	result := isSameTree(head1, head2)

	fmt.Println("Result is ", result)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return extrVal(p, q)
}

func extrVal(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		if !extrVal(p.Left, q.Left) || !extrVal(p.Right, q.Right) {
			return false
		}
	}

	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}

	return true
}

