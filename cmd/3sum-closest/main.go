package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1,2,1,-4}
	target := 1 //2

	nums = []int{1,1,1,0}
	target = -100 //2

	nums = []int{1,3,4,7,8,9}
	target = 15 //15

	nums = []int{4,0,5,-5,3,3,0,-4,-5}
	target = -2 //-2

	res := threeSumClosest(nums, target)
	fmt.Println(res)
}

/*
	16. 3Sum Closest
	Medium
	https://leetcode.com/problems/3sum-closest/
	Runtime 4ms Beats 98.87%, Memory 4.54MB Beats 5.10%
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)

	closest := nums[0] + nums[1] + nums[2]
	diff := getDiff(target, closest)
	for i:=0; i<n-2; i++ {
		searched := target - nums[i]

		j:=i+1
		k:=n-1
		twoSum := nums[j] + nums[k]
		diffTwoSum := getDiff(searched, twoSum)
		bestSum := twoSum
		bestDiff := diffTwoSum

		for {
			switch true {
			case twoSum > searched:
				k--
			case twoSum < searched:
				j++
			case twoSum == searched:
				return nums[i] + twoSum
			}

			if j<k {
				twoSum = nums[j] + nums[k]
				diffTwoSum = getDiff(searched, twoSum)

				if diffTwoSum < bestDiff {
					bestDiff = diffTwoSum
					bestSum = twoSum
				}
				continue
			}
			break
		}

		if bestDiff < diff {
			closest = bestSum + nums[i]
			diff = bestDiff
		}
	}

	return closest
}

func getDiff(t, s int) int {
	diff := t-s
	if diff < 0 {
		diff *= -1
	}

	return diff
}

/*
	Runtime 580ms Beats 6.19%, Memory 4.54MB Beats 5.10%
 */
func threeSumClosest2(nums []int, target int) int {
	n := len(nums)

	closest := nums[0] + nums[1] + nums[2]
	diff := getDiff(target, closest)
	for i:=0; i<n-2; i++ {
		for j:=i+1; j<n-1; j++ {
			for k:=j+1; k<n; k++ {
				numSumm := nums[i] + nums[j] + nums[k]
				numDiff := getDiff(target, numSumm)
				if numDiff < diff {
					closest = numSumm
					diff = numDiff
					continue
				}
				break
			}
		}
	}

	return closest
}

