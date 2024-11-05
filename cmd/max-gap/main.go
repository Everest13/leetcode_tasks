package main

import (
	"fmt"
)

func main() {
	nums := []int{9,3,7,1,6}
	nums = []int{1,3,100}
	nums = []int{1,1,1,1}
	nums = []int{93,94,95,96,97,99,100}

	res := maximumGap(nums)
	fmt.Println(res)
}

/*
	164. Maximum Gap
	Medium
	https://leetcode.com/problems/maximum-gap/description/
	Runtime 9ms Beats 97.73%, Memory 11.70MB Beats 34.32%
 */
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	n := len(nums)
	bucketSize := getMax(1, (maxVal-minVal)/(n-1))
	bucketCount := (maxVal-minVal)/bucketSize + 1 //сколько корзин

	minBucket := make([]int, bucketCount)
	maxBucket := make([]int, bucketCount)
	//разрыв между значениями внутри одной корзины всегда будет меньше, чем разрыв между значениями в соседних корзинах
	for i := range minBucket {
		minBucket[i] = -1
		maxBucket[i] = -1
	}

	for _, num := range nums {
		bucketIndex := (num - minVal) / bucketSize
		if minBucket[bucketIndex] == -1 || num < minBucket[bucketIndex] {
			minBucket[bucketIndex] = num
		}
		if maxBucket[bucketIndex] == -1 || num > maxBucket[bucketIndex] {
			maxBucket[bucketIndex] = num
		}
	}

	maxGap := 0
	previousMax := minVal
	for i := 0; i < bucketCount; i++ {
		if minBucket[i] == -1 {
			continue
		}
		maxGap = getMax(maxGap, minBucket[i]-previousMax)
		previousMax = maxBucket[i]
	}

	return maxGap
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
