package main

import (
	"fmt"
	"sync"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var lists []*ListNode
	arrs := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 23, 45, 67, 90}, {3, 6, 8, 10, 11, 15, 16, 17}, {8, 11, 12, 15, 16, 18, 19, 25, 37}}

	var node *ListNode
	for _, arr := range arrs {
		head := &ListNode{
			Val:  arr[0],
			Next: nil,
		}
		current := head
		for i := 1; i < len(arr); i++ {
			node = &ListNode{
				Val:  arr[i],
				Next: nil,
			}

			current.Next = node
			current = node
		}

		lists = append(lists, node)
		node = nil
	}

	result := mergeKLists(lists)
	next := result
	for next != nil {
		fmt.Println(next)
		next = next.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	chList := make(chan *ListNode)
	var wg sync.WaitGroup

	for i := 0; i < len(lists); i++ {
		go checkAndMerge(&wg, chList, lists[i], lists[i+1])
		i += 2
	}

	wg.Wait()
	result := <-chList

	close(chList)
	
	return result
}

func checkAndMerge(wg *sync.WaitGroup, chList chan *ListNode, list1 *ListNode, list2 *ListNode) {
	result := mergeTwoLists(list1, list2)
	for {
		list, ok := <-chList
		if !ok {
			chList <- result
			break
		}

		mergeTwoLists(result, list)
	}

	wg.Done()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, current *ListNode
	var least *ListNode
	for list1 != nil || list2 != nil {
		if list1.Val < list2.Val {
			least = list1
		} else {
			least = list2
		}

		if head == nil {
			head = least
		} else {
			current.Next = least
		}

		current = least
		least = least.Next

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
