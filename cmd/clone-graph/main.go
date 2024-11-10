package main

import "fmt"

type Node struct {
	Val int
	Neighbors []*Node
}

func main() {
	node1 := &Node{
		Val: 1,
	}
	node2 := &Node{
		Val: 2,
	}
	node3 := &Node{
		Val: 3,
	}
	node4 := &Node{
		Val: 4,
	}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	res := cloneGraph(node3)
	fmt.Println(res)
}

/*
	133. Clone Graph
	Medium
	https://leetcode.com/problems/clone-graph/description/
	Runtime 0-3ms Beats 56.48%-100.00%, Memory 4.64MB Beats 17.52%
 */
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	mapGraph := map[int]*Node{}
	recursion := func(node *Node) {}
	recursion = func(node *Node) {
		nodeVal := node.Val
		if _, ok := mapGraph[nodeVal]; ok {
			return
		}

		mapGraph[node.Val] = &Node{Val: nodeVal}
		neighbors := make([]*Node, 0, len(node.Neighbors))
		for _, n := range node.Neighbors {
			nVal := n.Val
			if _, ok := mapGraph[nVal]; !ok {
				recursion(n)
			}
			neighbors = append(neighbors, mapGraph[nVal])
		}
		mapGraph[nodeVal].Neighbors = neighbors
	}

	recursion(node)

	return mapGraph[1]
}
