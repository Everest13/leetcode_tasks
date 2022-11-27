package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Task #6 Remove Nth Node From End of List (Medium)
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.3 MB, less than 35.84% of Go online submissions for Remove Nth Node From End of List.
*/
func main() {
	//example
	//arr := []int{1, 2}
	//arr := []int{1, 2, 3, 4, 5, 6}
	arr := []int{1}
	n := 1
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

	result := removeNthFromEnd(head, n)
	next := result
	for next != nil {
		fmt.Println(next)
		next = next.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	counter := 0
	for current != nil {
		current = current.Next
		counter++
	}

	current = head
	var previous *ListNode
	count := 1
	deleted := counter - n + 1
	for current != nil {
		if count == deleted {
			if previous != nil {
				previous.Next = current.Next
				return head
			}

			head = head.Next
			return head
		}

		count++
		previous = current
		current = current.Next
	}

	return head
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	current1 := head
	current2 := head
	var prevCurrent2 *ListNode
	i := 0
	for current1 != nil {
		if i < n {
			i++
		} else {
			prevCurrent2 = current2
			current2 = current2.Next
		}

		current1 = current1.Next
	}

	if prevCurrent2 != nil {
		prevCurrent2.Next = current2.Next
		return head
	} else {
		return current2.Next
	}

	return head
}
