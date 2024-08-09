package main

import "fmt"

func main() {
	letters := []byte{'c','f','j'}
	var target byte = 'a'

	res := nextGreatestLetter(letters, target)
	fmt.Println(res)
}

/*
	https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
	https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
	Runtime 2ms Beats 60.23%
	Memory 2.82MB Beats 65.50%
 */
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[0] < target {
		return letters[0]
	}

	for i:=0; i<len(letters); i++ {
		if target < letters[i] {
			return letters[i]
		}
	}

	return letters[0]
}
