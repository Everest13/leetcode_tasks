package main

import (
	"fmt"
	"sort"
)

/*
1561. Maximum Number of Coins You Can Get
https://leetcode.com/problems/maximum-number-of-coins-you-can-get/description/?envType=daily-question&envId=2023-11-24
Runtime 107 ms Beats 88%, Memory 8.8 MB Beats 20%
 */
func main() {
	piles := []int{9,8,7,6,5,1,2,3,4}

	res := maxCoins(piles)

	fmt.Println("Res is ", res)

}

func maxCoins(piles []int) int {
	sort.Ints(piles)

	var res int
	j:=0
	i:=len(piles)-1
	for i>0 {
		res += piles[i-1]
		j++
		i-=2
		if j > i-2 {
			break
		}
	}

	return res
}