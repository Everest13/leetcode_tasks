package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	arr := []int{5,3,1,2,5,1,2} //1,3
	//arr = []int{1,3,2,2,3,2,2,2,7} //3,3

	head := &ListNode{
		Val: arr[0],
	}
	curr := head
	for i:=1; i<len(arr); i++ {
		next := &ListNode{
			Val: arr[i],
		}
		curr.Next = next
		curr = next
	}

	res := nodesBetweenCriticalPoints(head)
	fmt.Println(res)

}

/*
	2058. Find the Minimum and Maximum Number of Nodes Between Critical Points
	Medium
	https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/?envType=daily-question&envId=2024-07-05
	Runtime 129ms Beats 72.73%, Memory 8.44MB Beats 90.09%
*/
func nodesBetweenCriticalPoints(head *ListNode) []int {
	last := 0
	i:=2
	prevVal := head.Val
	curr := head.Next
	maxDist := 0
	minDist := 0
	for curr.Next != nil {
		next := curr.Next
		val := curr.Val
		if (val > prevVal && val > next.Val) || (val < prevVal && val < next.Val) {
			if last != 0 {
				dist := i - last
				maxDist += dist

				if minDist == 0 {
					minDist = dist
				}

				if dist < minDist {
					minDist = dist
				}
			}

			last = i
		}

		prevVal = val
		curr = next
		i++
	}

	if maxDist < 1 {
		return []int{-1, -1}
	}

	res := []int{minDist, maxDist}
	return res
}

// Runtime 136ms Beats 54.55%, Memory 9.61MB Beats 72.73%
func nodesBetweenCriticalPoints2(head *ListNode) []int {
	minDist := 0
	maxDist := 0
	lastCritId := 0

	prev := head
	curr := head.Next
	i:=2
	for curr.Next != nil {
		if (curr.Val > prev.Val && curr.Val > curr.Next.Val) || (curr.Val < prev.Val && curr.Val < curr.Next.Val) {
			if lastCritId != 0 {
				dist := i - lastCritId
				maxDist += dist

				if minDist == 0 {
					minDist = dist
				}

				if dist < minDist {
					minDist = dist
				}
			}

			lastCritId = i
		}

		prev = curr
		curr = curr.Next
		i++
	}

	if minDist == 0 {
		return []int{-1, -1}
	}

	res := []int{minDist, maxDist}
	return res
}

// Runtime 162ms Beats 18.18%
func nodesBetweenCriticalPoints3(head *ListNode) []int {
	arr := []int{}

	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	minDist := 0
	maxDist := 0
	last := 0
	for i:=1; i<len(arr)-1; i++ { //1,3,2,2,3,2,2,2,7
		if (arr[i] > arr[i-1] && arr[i] > arr[i+1]) || (arr[i] < arr[i-1] && arr[i] < arr[i+1]) {
			if last != 0 {
				dist := i - last
				maxDist += dist

				if minDist == 0 {
					minDist = dist
				}

				if dist < minDist {
					minDist = dist
				}
			}
			last = i
		}
	}


	if minDist == 0 {
		return []int{-1, -1}
	}

	res := []int{minDist, maxDist}
	return res
}
