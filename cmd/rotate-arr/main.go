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