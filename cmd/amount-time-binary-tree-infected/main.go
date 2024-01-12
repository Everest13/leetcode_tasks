package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	headRightR := &TreeNode{
		Val: 6,
	}
	headRightL := &TreeNode{
		Val: 10,
	}

	headLeftRL := &TreeNode{
		Val: 9,
	}
	headLeftRR := &TreeNode{
		Val: 2,
	}

	headLeftR := &TreeNode{
		Val: 4,
		Right: headLeftRR,
		Left: headLeftRL,
	}

	headLeft := &TreeNode{
		Val: 5,
		Right: headLeftR,
	}
	headRight := &TreeNode{
		Val: 3,
		Right: headRightR,
		Left: headRightL,
	}

	head := &TreeNode{
		Val: 1,
		Left: headLeft,
		Right: headRight,
	}

	start := 3

	//headRight1 := &TreeNode{
	//	Val: 5,
	//	Right: nil,
	//	Left: nil,
	//}
	//
	//head1 := &TreeNode{
	//	Val: 2,
	//	Left: nil,
	//	Right: headRight1,
	//}
	//
	//start1 := 5
	//res := amountOfTime(head1, start1)


	res := amountOfTime(head, start)
	fmt.Println(res)
}

func amountOfTime(root *TreeNode, start int) int {
	type Node struct {
		Val int
		Ribs []*Node
		Mark bool
	}

	var startNode *Node
	isStarted := false
	rec :=  func(originalNode *TreeNode, parent *Node) {}
	rec = func(originalNode *TreeNode, parent *Node) {
		node := &Node{
			Val: originalNode.Val,
			Ribs: []*Node{},
		}

		if originalNode.Val == start && !isStarted {
			startNode = node
			isStarted = true
		}

		if parent != nil {
			parent.Ribs = append(parent.Ribs, node)

			node.Ribs = []*Node{0: parent}
		}

		if originalNode.Right != nil {
			rec(originalNode.Right, node)
		}
		if originalNode.Left != nil {
			rec(originalNode.Left, node)
		}

	}

	rec(root,nil)
	root = nil
	time := 0
	infectionTimeRec := func(nodes []*Node) {}
	infectionTimeRec = func(nodes []*Node) {
		nextInfectedNodes := make([]*Node, 0, len(nodes))
		for _, ribNode := range nodes {
			if ribNode.Mark {
				continue
			}
			ribNode.Mark = true
			nextInfectedNodes = append(nextInfectedNodes, ribNode.Ribs...)
		}

		if len(nextInfectedNodes) > 0 {
			time++
			infectionTimeRec(nextInfectedNodes)
		}
	}

	startNode.Mark = true
	infectionTimeRec(startNode.Ribs)

	return time
}
