package main

import "fmt"

func main() {
	nums := []int{2,3,1,1,4}

	res := canJump(nums)
	fmt.Println(res)
}

/*
	55. Jump Game
	https://leetcode.com/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150
	Runtime 38ms Beats 96.00%, Memory 6.96MB Beats 87.41% M
 */
func canJump(nums []int) bool {
	num := nums[0]
	for i:=1; i<len(nums); i++ { // 2,3,1,1,4
		if num >= 1 {
			num = num - 1
			next := nums[i]
			if next > num {
				num = next
			}
			continue
		}

		return false
	}

	return true
}
