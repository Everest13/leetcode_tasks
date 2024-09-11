package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	arr := []int{1,2,3,4,5}
	n := 2

	arr = []int{1}
	n = 1

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

	resHead := removeNthFromEnd(head, n)
	current = resHead
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

/**
	19. Remove Nth Node From End of List
	Medium
	https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
	Runtime 0ms Beats 100.00%, Memory 2.50MB Beats 13.42%
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := []*ListNode{}
	current := head
	for current != nil {
		arr = append(arr, current)
		current = current.Next
	}

	arrLen := len(arr)
	idx := arrLen - n
	if idx == 0 {
		return head.Next
	}

	prev := arr[idx-1]
	prev.Next = nil
	if idx+1 < arrLen {
		prev.Next = arr[idx+1]
	}

	return head
}