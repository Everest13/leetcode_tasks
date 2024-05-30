package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	2.Task Add Two Numbers
	https://leetcode.com/problems/add-two-numbers/description/
	Runtime - 3ms(beats 99%), memory: 4.5mb(beats 65%) M
*/
func main() {
	arr1 := []int{2, 4, 3}
	head1 := &ListNode{
		Val:  arr1[0],
		Next: nil,
	}

	arr2 := []int{5, 6, 4}
	head2 := &ListNode{
		Val:  arr2[0],
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

	var node2 *ListNode
	current2 := head2
	for i := 1; i < len(arr2); i++ {
		node2 = &ListNode{
			Val:  arr2[i],
			Next: nil,
		}

		current2.Next = node2
		current2 = node2
	}

	resultNode := addTwoNumbers(head1, head2)
	for resultNode != nil {
		fmt.Println(resultNode)
		resultNode = resultNode.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	checkValSum := func(valSum int) (int, int) {
		if valSum > 9 {
			return valSum % 10, 1
		}
		return valSum, 0
	}

	unit, valSum := checkValSum(l1.Val + l2.Val)
	head := &ListNode{
		Val: unit,
	}
	previous := head
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil || l2 != nil || valSum != 0 {
		if l1 != nil {
			valSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			valSum += l2.Val
			l2 = l2.Next
		}

		unit, valSum = checkValSum(valSum)
		current := &ListNode{
			Val: unit,
		}
		previous.Next = current
		previous = current
	}

	return head
}
