package main

func main() {

}

/*
	121. Best Time to Buy and Sell Stock
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
	Runtime 92ms Beats 73.08% Memory 8.08MB Beats 56.07% E
*/
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	minP := prices[0]
	diff := 0
	for i:=1; i<len(prices); i++ {
		if prices[i] < minP {
			minP = prices[i]
			continue
		}

		if prices[i] > minP {
			locDiff := prices[i] - minP
			if locDiff > diff {
				diff = locDiff
			}
		}
	}

	return diff
}


