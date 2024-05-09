package main

import (
	"fmt"
	"sort"
)

func main() {
	happiness := []int{1,2,3}
	k := 2

	happiness = []int{1,1,1,1}
	k = 2

	res := maximumHappinessSum(happiness, k)

	fmt.Println(res)
}

/*
	3075. Maximize Happiness of Selected Children
	Medium
	https://leetcode.com/problems/maximize-happiness-of-selected-children/description/?envType=daily-question&envId=2024-05-09
	Runtime 156ms Beats 91.08%, Memory 12.38MB Beats 42.68%
*/
func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(a, b int) bool {
		return happiness[b] < happiness[a]
	})

	summ := happiness[0]
	minus := 1
	for i:=1; i<k; i++ {
		if happiness[i] > minus {
			happiness[i] -= minus
			summ += happiness[i]
		}

		minus++
	}

	return int64(summ)
}

