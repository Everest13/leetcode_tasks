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
	Hard
	Runtime 171ms Beats 9.23%, Memory 6.53MB Beats 32.92%
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
		//если идем на повышение
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

		//если идем на понижение
		if ratings[i] > ratings[j] {
			decreaseCount++

			if candyInc != 0 { //первый 0
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

// ChatGPT
func candyChatGPT(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)

	// Give each child 1 candy to start
	for i := range candies {
		candies[i] = 1
	}

	// Scan from left to right to ensure children with higher ratings get more candies than their left neighbors
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Scan from right to left to ensure children with higher ratings get more candies than their right neighbors
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	// Calculate total candies
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}


func candy2(ratings []int) int {
	candies := len(ratings)
	candies2 := candies

	idxCandiesMap := map[int]int{0: 0}
	i:=0
	j:=i+1
	for j<len(ratings) {
		//если идем на повышение
		if ratings[i] < ratings[j] {
			idxCandiesMap[j] = idxCandiesMap[i] + 1

			candies2 += idxCandiesMap[i] + 1 //todo
		}

		if ratings[i] == ratings[j] {
			idxCandiesMap[j] = 0
		}

		//если идем на понижение
		if ratings[i] > ratings[j] {
			idxCandiesMap[j] = 0

			//надо проверить предыдущее значение на 0
			if idxCandiesMap[i] == 0 {
				idxCandiesMap[i]++
				candies2++ //todo
				// надо прибавить по 1 ко всем предыдущим значениям пока не дойжем до 0 или начала и пока идет на увеличиение обратно
				k := i
				l:=k-1
				for l>=0 && ratings[l] > ratings[k] && idxCandiesMap[k] == idxCandiesMap[l] {
					idxCandiesMap[l]++
					candies2++ //todo
					k--
					l--
				}

			}
		}

		i++
		j++
	}

	return candies2
}

func candy3(ratings []int) int {
	candies := len(ratings)

	i:=0
	j:=i+1

	var maxCandies int
	var beforeZero int
	var beforeZeroIdx int
	var zeroInOrder int

	for j<len(ratings) {
		//если идем на повышение
		if ratings[i] < ratings[j] {
			candies++
			maxCandies++
			zeroInOrder = 0
		}

		//если идем на понижение
		if ratings[i] > ratings[j] {
			zeroInOrder++

			if maxCandies != 0 {
				beforeZero = maxCandies //самое большое значение до первого сброса
				beforeZeroIdx = i
				maxCandies = 0
				continue
			}

			// если maxCandies=0 значит сборс в 0 уже был и надо накидать конфет
			k := i
			for k >= beforeZeroIdx {
				if k == beforeZeroIdx && beforeZero > zeroInOrder {
					break
				}

				candies++
				k--
			}
		}

		i++
		j++
	}

	return candies
}
