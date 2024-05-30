package main

import "fmt"

func main() {
	rat := []int{1,0,2} //5
	rat = []int{1,2,2} //4
	//rat = []int{1,3,2,2,1} //7
	//rat = []int{1,2,87,87,87,2,1} //13
	rat = []int{29,51,87,87,72,12} //12
	//rat = []int{1,2,3,1,0} //9
	rat = []int{1,6,10,8,7,3,2} //18

	res := candy(rat)
	fmt.Println(res)
}

/*
	135. Candy
	https://leetcode.com/problems/candy/description/?envType=daily-question&envId=2024-02-07
	Runtime 171ms Beats 9.23%, Memory 6.53MB Beats 32.92% H
 */
func candy(ratings []int) int {
	candies := len(ratings)

	var candyInc int
	var beforeZeroIdx int
	var decreaseCount int
	var maxCandies int

	i:=0
	j:=i+1
	for j<len(ratings) {
		if ratings[i] < ratings[j] {
			if decreaseCount > 0 {
				candyInc = 0
				decreaseCount = 0
				maxCandies = 0
			}

			candyInc++
			candies += candyInc
		}

		if ratings[i] == ratings[j] {
			candyInc = 0
		}

		if ratings[i] > ratings[j] {
			decreaseCount++

			if candyInc != 0 {
				beforeZeroIdx = i
				maxCandies = candyInc
				candyInc = 0
				i++
				j++
				continue
			}

			k := i
			l:=j
			for k>=0 && ratings[k] > ratings[l] {
				if k == beforeZeroIdx && maxCandies >= decreaseCount {
					break
				}

				candies++
				k--
				l--
			}
		}

		i++
		j++
	}

	return candies
}
