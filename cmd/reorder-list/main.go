package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	lnList := 5

	head := &ListNode{
		Val: 1,
	}

	curr := head
	for i := 2; i < lnList; i++ {
		next := ListNode{
			Val: i,
		}

		curr.Next = &next
		curr = &next
	}

	reorderList(head)

	curr = head
	for curr.Next != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

/*
	143. Reorder List
	Medium
	https://leetcode.com/problems/reorder-list/description/?envType=daily-question&envId=2024-05-11
	Runtime 3ms Beats 96.7%
	Memory 6.3MB Beats 31.68%
 */
func reorderList(head *ListNode)  {
	arr := []*ListNode{}
	curr := head
	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}

	n := len(arr)
	curr = arr[0]
	i:=1
	j:=n-1
	for i<j {
		curr.Next = arr[j]
		curr.Next.Next = arr[i]
		j--
		i++

		curr = curr.Next.Next
	}

	if j == i {
		curr.Next = arr[i]
		curr = curr.Next
	}

	curr.Next = nil
}
