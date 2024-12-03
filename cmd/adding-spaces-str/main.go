package main

import "fmt"

func main() {
	s := "icodeinpython"
	spaces := []int{1,5,7,9} // icodeinpython - i codeinpython - i code inpython

	res := addSpaces(s, spaces)
	fmt.Println(res)
}

/*
	2109. Adding Spaces to a String
	Medium
	https://leetcode.com/problems/adding-spaces-to-a-string/
	Runtime 0ms Beats 100.00%, Memory 19.22MB Beats 63.42%
 */
func addSpaces(s string, spaces []int) string {
	sb := make([]byte, len(s)+len(spaces))
	j := 0
	k := 0

	for i := 0; i < len(s); i++ {
		if j < len(spaces) && i == spaces[j] {
			sb[k] = ' '
			k++
			j++
		}
		sb[k] = s[i]
		k++
	}

	return string(sb)
}