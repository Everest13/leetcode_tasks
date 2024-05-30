package main

import "sort"

func main() {

}

/*
	2706. Buy Two Chocolates
	https://leetcode.com/problems/buy-two-chocolates/?envType=daily-question&envId=2024-05-11 E
*/
func buyChoco(prices []int, money int) int {
	sort.Ints(prices)

	twoChoco := prices[0] + prices[1]
	if money > twoChoco {
		return money-twoChoco
	}

	return money
}
