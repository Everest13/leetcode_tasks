package main

func main() {
	digits := []int{0,1,0,3,12}

	moveZeroes(digits)
}

/*
	283. Move Zeroes
	Easy
	https://leetcode.com/problems/move-zeroes/
 */
func moveZeroes(nums []int)  {
	next := 0
	countZero := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] == 0 {
			countZero++
			continue
		}
		nums[next] = nums[i]
		next++
	}

	for next < len(nums) {
		nums[next] = 0
		next++
	}
}