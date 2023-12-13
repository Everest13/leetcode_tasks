package main

import (
	"fmt"
	"sort"
)

func main() {
	events := [][]int{{1,1,1}, {2,2,2}, {3,3,3}, {4,4,4}}
	events = [][]int{{1,2,4},{3,4,3},{2,3,10}}
	events = [][]int{{1,2,4},{3,4,3},{2,3,1}}
	//events = [][]int{{1,1,1},{2,2,2},{3,3,3},{4,4,4}}

	//events = [][]int{{1,3,4},{2,4,1},{1,1,4},{3,5,1},{2,5,5}}
	events = [][]int{{21,77,43},{2,74,47},{6,59,22},{47,47,38},{13,74,57},{27,55,27},{8,15,8}}
	events = [][]int{{31,57,53},{5,63,29},{54,57,32},{55,83,28},{25,64,5},{5,33,33},{32,68,27},{30,99,54}}


	k := 4


	res := maxValue(events, k)
	fmt.Println(res)
}

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][2] > events[j][2]
	})

	testArr := [][]int{
		events[0],
	}
	res := events[0][2]
	k--
	timeTable := map[int]int{
		events[0][0]: events[0][1],
	}

	for i := 1; i < len(events); i++ {
		if k == 0 {
			break
		}

		var isBusy bool
		for start, end := range timeTable {
			test := events[i]
			_=test
			s1 := events[i][0] - start
			s2 := end - events[i][1]
			_,_ = s1,s2

			if (events[i][0] <= start && events[i][1] >= start) ||
				(events[i][0] <= end && events[i][1] >= end) ||
				(events[i][0] >= start && events[i][1] <= end) {

				isBusy = true
				break
			}
		}


		if isBusy {
			continue
		}

		testArr = append(testArr, events[i])


		res += events[i][2]
		k--
	}

	return res
}