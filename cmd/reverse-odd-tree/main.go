package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildPerfectTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}

	for i := 0; i < len(vals)/2; i++ {
		if 2*i+1 < len(vals) {
			nodes[i].Left = nodes[2*i+1]
		}
		if 2*i+2 < len(vals) {
			nodes[i].Right = nodes[2*i+2]
		}
	}

	return nodes[0] // root
}

func printLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		n := len(queue)
		fmt.Printf("Level %d: ", level)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
		level++
	}
}

func main() {
	vals := []int{2,3,5,8,13,21,34}

	root := buildPerfectTree(vals)
	res := reverseOddLevels(root)
	printLevelOrder(res)
}

/*
	2415. Reverse Odd Levels of Binary Tree
	Medium
	https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/description/
 */
// BFS
func reverseOddLevels(root *TreeNode) *TreeNode {
	nodeArr := []*TreeNode{root}
	newArr := []*TreeNode{}
	level := 0
	for nodeArr[0].Right != nil {
		//собираем массив текущего уровня
		for _, node := range nodeArr {
			newArr = append(newArr, node.Left)
			newArr = append(newArr, node.Right)
		}

		nodeArr = newArr
		newArr = []*TreeNode{}
		level++
		//если уровень нечетный - меняем значения
		if level%2 != 0 {
			i:=0
			j:=len(nodeArr)-1
			for i<j {
				nodeIVal := nodeArr[i].Val
				nodeArr[i].Val = nodeArr[j].Val
				nodeArr[j].Val = nodeIVal
				i++
				j--
			}
		}
	}

	return root
}

// DFS
func reverseOddLevelsDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dfs(root.Left, root.Right, 1)
	return root
}

func dfs(left, right *TreeNode, level int) {
	if left == nil || right == nil {
		return
	}
	if level%2 == 1 {
		left.Val, right.Val = right.Val, left.Val
	}
	dfs(left.Left, right.Right, level+1)
	dfs(left.Right, right.Left, level+1)
}
