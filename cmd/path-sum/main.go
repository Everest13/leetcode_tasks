package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	nodeRight := &TreeNode{
		Val: -3,
	}
	root := &TreeNode{
		Val: -2,
		Right: nodeRight,
		Left: nil,
	}

	targetSum := -5
	res := hasPathSum(root, targetSum)
	fmt.Println(res)
}

/*
	112. Path Sum
	Easy
	https://leetcode.com/problems/path-sum/description/
	Runtime 0ms Beats 100.00%, Memory 6.71MB Beats 36.38%
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	recursion := func(node *TreeNode, rest int) bool {return false}
	recursion = func(node *TreeNode, rest int) bool {
		if node == nil {
			return false
		}

		nodeRest := rest - node.Val
		if nodeRest == 0 && node.Left == nil && node.Right == nil {
			return true
		}

		return recursion(node.Left, nodeRest) || recursion(node.Right, nodeRest)
	}

	return recursion(root, targetSum)
}
