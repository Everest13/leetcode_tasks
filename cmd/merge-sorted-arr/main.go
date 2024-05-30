package main

func main() {

}


/*
	88. Merge Sorted Array
	https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
	Runtime 2ms Beats 73.26%, Memory 2.64MB, Beats 14.96% E
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		nums1 = append(nums1[0:m])
		return
	}

	res := make([]int, 0, m+n)
	j:=0
	for i:=0; i<m; i++ {
		test := []int{}
		for j<n && nums2[j] <= nums1[i] {
			test = append(test, nums2[j])
			j++
		}

		res = append(res, test...)
		res = append(res, nums1[i])
	}

	if j < n {
		res = append(res, nums2[j:]...)
	}

	nums1 = append(nums1[0:0], res...)
}