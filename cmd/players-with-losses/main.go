package main

import (
	"fmt"
	"sort"
)

/*
	2225. Find Players With Zero or One Losses
	Medium
	https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/?envType=daily-question&envId=2024-01-15
	Runtime 279ms Beats83.46% Memory 25.70MB Beats18.04%
 */
func main() {
	matches := [][]int{{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}}
	matches = [][]int{{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}}

	res := findWinners(matches)
	fmt.Println(res)
}

/*
	Runtime 279ms Beats 81.38%
	Memory 25.70MB Beats 60.46%
 */
func findWinners(matches [][]int) [][]int {
	losses := map[int]int{}
	for i:=0; i<len(matches); i++ {
		losses[matches[i][0]]+=0
		losses[matches[i][1]]++
	}

	res := [][]int{0: {}, 1: {}}
	for loser, count := range losses {
		if count == 0 {
			res[0] = append(res[0], loser)
			continue
		}

		if count == 1 {
			res[1] = append(res[1], loser)
			continue
		}
	}

	sort.Ints(res[0])
	sort.Ints(res[1])
	return res
}