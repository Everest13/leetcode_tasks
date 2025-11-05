package main

func main() {

}

/*
	66. Plus One
	Easy
	https://leetcode.com/problems/plus-one/description/
	Runtime 0ms Beats 100.00%, Memory 4.05MB Beats 13.11%
 */
func plusOne(digits []int) []int {
	for i:=len(digits)-1; i>=0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}

	return append([]int{1}, digits...)
}

func plusOne2(digits []int) []int {
	n := len(digits)-1

	recursion := func(i int) {}
	recursion = func(i int) {
		if digits[i] != 9 {
			digits[i]++
			return
		}

		if i == 0 {
			digits[i] = 0
			digits = append([]int{1}, digits...)
			return
		}

		digits[i] = 0
		i--
		recursion(i)
	}

	recursion(n)

	return digits
}
