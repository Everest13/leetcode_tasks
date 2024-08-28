package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	arr := []int{18,6,10,3}

	head := &ListNode{
		Val: arr[0],
	}

	current := head
	for i:=1; i<len(arr);i++ {
		next := &ListNode{
			Val: arr[i],
		}

		current.Next = next
		current = next
	}

	resHead := insertGreatestCommonDivisors(head)
	current = resHead
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

/*
	2807. Insert Greatest Common Divisors in Linked List
	Medium
	https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/description/
	Runtime 9ms 84.89%, Memory 6.20MB 87.77%
*/
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	current := head
	for current.Next != nil {
		next := current.Next
		current.Next = &ListNode{
			Val: gcd(current.Val, next.Val),
			Next: next,
		}

		current = next
	}

	return head
}

func gcd(n, m int) (rest int) {
	if n > m {
		rest = n%m
		if rest == 0 {
			return m
		}
		return gcd(rest, m)
	}

	rest = m%n
	if rest == 0 {
		return n
	}
	return gcd(n, rest)
}
