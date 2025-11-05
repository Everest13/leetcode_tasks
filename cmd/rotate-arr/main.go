package main

func main() {

}

/*
	189. Rotate Array
	https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
	Runtime 23ms Beats 78.54%, Memory 8.35MB Beats 14.90% E
*/
func rotate(nums []int, k int)  {
	l:=len(nums)
	res := make([]int, l)

	if k > l {
		k = k-(l*( k/l))
	}

	res = append(nums[l-k:], nums[0:l-k]...)
	nums = append(nums[0:0], res...)
}

// classic approach
func rotate2(nums []int, k int) {
	lenNums := len(nums)
	if k > lenNums {
		k = k%lenNums
	}

	reverse(nums, 0, lenNums-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, lenNums-1)
}

func reverse(nums []int, from, to int) {
	for from < to {
		nums[from], nums[to] = nums[to], nums[from]
		from++
		to--
	}
}