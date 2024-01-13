package main

import "fmt"

/*
	1704. Determine if String Halves Are Alike
	Easy
	https://leetcode.com/problems/determine-if-string-halves-are-alike/description/?envType=daily-question&envId=2024-01-12
	Runtime 2ms Beats 73.81%, Memory 2.22MB Beats 21.43%
 */
func main() {
	s:= "AbCdEfGh"

	res := halvesAreAlike(s)
	fmt.Println(res)
}

var allowed = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func halvesAreAlike(s string) bool {
	lenS := len(s)
	count := 0
	j:=lenS-1
	for i:=0; i<=lenS/2-1; i++ {
		if _, okI := allowed[rune(s[i])]; okI {
			count++
		}

		if _, okJ := allowed[rune(s[j])]; okJ {
			count--
		}
		j--
	}

	return count == 0
}
