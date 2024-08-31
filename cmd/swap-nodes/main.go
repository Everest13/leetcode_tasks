package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	arr := []int{1,2,3,4}
	head := &ListNode{
		Val: arr[0],
	}

	current := head
	for i:=1; i<len(arr); i++ {
		next := &ListNode{
			Val: arr[i],
		}

		current.Next = next
		current = next
	}

	head = swapPairs(head)
	current = head
	for current != nil {
		fmt.Println(current.Val)
	}
}

/**
	24. Swap Nodes in Pairs
	https://leetcode.com/problems/swap-nodes-in-pairs/description/
	Runtime 0ms Beats 100.00%, Memory 2.30MB Beats 80.68%
*/
func swapPairs(head *ListNode) *ListNode {
	current := head
	var prev *ListNode

	for current != nil && current.Next != nil {
		next := current.Next
		nextNext := next.Next

		current.Next = nextNext
		next.Next = current

		if prev != nil {
			prev.Next = next
		} else {
			head = next
		}

		prev = current
		current = nextNext
	}

	return head
}