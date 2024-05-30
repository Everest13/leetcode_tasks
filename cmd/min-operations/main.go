package main

func main() {

}

/*
	1758. Minimum Changes To Make Alternating Binary String
	https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/?envType=daily-question&envId=2024-05-11
	Runtime 2ms Beats 68.42% Memory 2.72MB Beats 13.16% E
*/
func minOperations(s string) int {
	countOps := func(prev rune) (count int) {
		for _, char := range s {
			if prev == char {
				if char == '0' {
					prev = '1'
				} else {
					prev = '0'
				}
				count++
				continue
			}
			prev = char
		}

		return
	}

	count0 := countOps('0')
	count1 := countOps('1')

	if count0 > count1 {
		return count1
	}

	return count0
}
