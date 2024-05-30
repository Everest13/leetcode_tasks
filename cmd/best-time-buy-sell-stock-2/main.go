package main

import "fmt"

func main() {
	prices := []int{7,1,5,3,4,6}

	res := maxProfit(prices)
	fmt.Println(res)
}

/*
	122. Best Time to Buy and Sell Stock II
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
	Runtime 4ms Beats 67.72%, Memory 3.41MB Beats 9.81% M
 */
func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		// If the current price is greater than the previous day's price,
		// we can make a profit by buying and selling on the same day.
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}