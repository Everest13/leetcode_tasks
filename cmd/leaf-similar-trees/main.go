package main

import "reflect"

/*
	872. Leaf-Similar Trees
	https://leetcode.com/problems/leaf-similar-trees/description/?envType=daily-question&envId=2024-01-09
	Runtime 2 ms Beats 71.85%, Memory 2.6 MB Beats 52.52% E
 */
func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs := map[int][]int{
		1: {},
		2: {},
	}

	recursion := func(node *TreeNode, n int) {}

	recursion = func(node *TreeNode, n int) {
		if node.Left == nil && node.Right == nil {
			leafs[n] = append(leafs[n], node.Val)
			return
		}

		if node.Left != nil {
			recursion(node.Left, n)
		}

		if node.Right != nil {
			recursion(node.Right, n)
		}
	}

	recursion(root1, 1)
	recursion(root2, 2)

	return reflect.DeepEqual(leafs[1], leafs[2])
}
