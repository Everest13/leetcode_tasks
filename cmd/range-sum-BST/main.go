package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	938. Range Sum of BST
	Easy
	https://leetcode.com/problems/range-sum-of-bst/description/?envType=daily-question&envId=2024-01-08
	Runtime 63 ms Beats 94.54%, Memory 7.4 MB Beats 95.56%
 */
func main()  {

}

func rangeSumBST(root *TreeNode, low int, high int) int {
	summ := 0

	recursion := func(node *TreeNode) {}
	recursion = func(node *TreeNode) {

		if node.Val >= low && node.Val <= high {
			summ+=node.Val
		}

		if node.Left != nil {
			recursion(node.Left)
		}

		if node.Right != nil {
			recursion(node.Right)
		}
	}

	recursion(root)

	return summ
}
