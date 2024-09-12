package main

import (
	"fmt"
	"sync"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	lists := [][]int{{1,4,5},{2,3,4},{3,6}}

	listsListNode := make([]*ListNode, 0, len(lists))
	for _, list := range lists {
		head := &ListNode{
			Val: list[0],
		}
		current := head
		for i:=1; i<len(list); i++ {
			next := &ListNode{
				Val: list[i],
			}
			current.Next = next
			current = next
		}

		listsListNode = append(listsListNode, head)
	}

	res := mergeKLists(listsListNode)
	fmt.Println(res)
}

/*
	23. Merge k Sorted Lists
	Hard
	https://leetcode.com/problems/merge-k-sorted-lists/
	Runtime 7ms Beats 76.82%, Memory 5.18MB Beats 45.01%
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	mergeTwoLists := func(listHead1, listHead2 *ListNode) {
		head := &ListNode{}
		if listHead1.Val < listHead2.Val {
			head = listHead1
			listHead1 = listHead1.Next
		} else {
			head = listHead2
			listHead2 = listHead2.Next
		}

		current := head
		for listHead1 != nil && listHead2 != nil {
			if listHead1.Val < listHead2.Val {
				current.Next = listHead1
				current = current.Next
				listHead1 = listHead1.Next
				continue
			}

			current.Next = listHead2
			current = current.Next
			listHead2 = listHead2.Next
		}

		if listHead1 == nil {
			current.Next = listHead2
		} else {
			current.Next = listHead1
		}

		lists = append(lists, head)
	}

	for len(lists) > 1 {
		heads := lists[0:2]
		lists = lists[2:]
		if heads[0] == nil && heads[1] == nil {
			continue
		}
		if heads[0] == nil {
			lists = append(lists, heads[1])
			continue
		}
		if heads[1] == nil {
			lists = append(lists, heads[0])
			continue
		}

		mergeTwoLists(heads[0], heads[1])
	}

	if len(lists) == 0 {
		return nil
	}

	return lists[0]
}

/*
	Time Limit Exceeded
 */
type Locker struct {
	Lists           []*ListNode
	mutex             sync.RWMutex
}

func (l *Locker) addLists(node *ListNode) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	l.Lists = append(l.Lists, node)
}

func (l *Locker) takeList() []*ListNode {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	list := l.Lists[0:2]
	l.Lists = l.Lists[2:]
	return list
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	locker := Locker{
		mutex: sync.RWMutex{},
		Lists: lists,
	}
	var wg sync.WaitGroup
	mergeTwoLists := func(listHead1, listHead2 *ListNode) {
		defer wg.Done()

		head := &ListNode{}
		if listHead1.Val < listHead2.Val {
			head = listHead1
			listHead1 = listHead1.Next
		} else {
			head = listHead2
			listHead2 = listHead2.Next
		}
		current := head
		for listHead1 != nil && listHead2 != nil {
			if listHead1.Val < listHead2.Val {
				current.Next = listHead1
				current = current.Next
				listHead1 = listHead1.Next
			} else {
				current.Next = listHead2
				current = current.Next
				listHead2 = listHead2.Next
			}
		}

		if listHead1 == nil {
			current.Next = listHead2
		} else {
			current.Next = listHead1
		}

		locker.addLists(head)
	}

	for {
		if len(locker.Lists) <= 1 {
			wg.Wait()
		}
		if len(locker.Lists) <= 1 {
			break
		}

		heads := locker.takeList()
		if heads[0] == nil && heads[1] == nil {
			continue
		}
		if  heads[0] == nil {
			locker.addLists(heads[1])
			continue
		}
		if heads[1] == nil {
			locker.addLists(heads[0])
			continue
		}

		go mergeTwoLists(heads[0], heads[1])
		wg.Add(1)
	}

	if len(locker.Lists) == 0 {
		return nil
	}

	return locker.Lists[0]
}
