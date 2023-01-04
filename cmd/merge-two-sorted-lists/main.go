package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Task #8 Merge Two Sorted Lists (Easy)
https://leetcode.com/problems/merge-two-sorted-lists/
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.5 MB, less than 88.59% of Go online submissions for Merge Two Sorted Lists.
*/
func main() {
	//first
	arr := []int{6, 7, 11, 14}
	//second
	arr1 := []int{4, 8, 13, 14}

	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	var node *ListNode
	current := head
	for i := 1; i < len(arr); i++ {
		node = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		current.Next = node
		current = node
	}

	head1 := &ListNode{
		Val:  arr1[0],
		Next: nil,
	}

	var node1 *ListNode
	current1 := head1
	for i := 1; i < len(arr1); i++ {
		node1 = &ListNode{
			Val:  arr1[i],
			Next: nil,
		}

		current1.Next = node1
		current1 = node1
	}

	result := mergeTwoLists(head, head1)
	next := result
	for next != nil {
		fmt.Println(next)
		next = next.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//var recursion func(current1 *ListNode, current2 *ListNode) *ListNode
	//recursion = func(current1 *ListNode, current2 *ListNode) *ListNode {
	//	if current1 == nil {
	//		return current2
	//	}
	//
	//	if current2 == nil {
	//		return current1
	//	}
	//
	//	if current1.Val < current2.Val {
	//		current1.Next = recursion(current2, current1.Next)
	//		return current1
	//	}
	//
	//	current2.Next = recursion(current1, current2.Next)
	//	return current2
	//}
	//
	//return recursion(list1, list2)

	var head, current *ListNode //head - is head!, current - last result node

	if list1.Val < list2.Val {
		head = list1
		current = list1
		list1 = list1.Next
	} else {
		head = list2
		current = list2
		list2 = list2.Next
	}

	for list1 != nil || list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			current = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			current = list2
			list2 = list2.Next
		}

		if list1 == nil {
			current.Next = list2
			break
		}

		if list2 == nil {
			current.Next = list1
			break
		}
	}

	return head
}
