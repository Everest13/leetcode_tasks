package main

import (
	"fmt"
)

/*
	844. Backspace String Compare
	https://leetcode.com/problems/backspace-string-compare/description/?envType=daily-question&envId=2023-10-19
	Runtime 2 ms Beats 15.41%, Memory 2.2 MB Beats 17.79% E
 */
func main() {
	s := "ab#c"
	t := "ad#c"

	s = "ab##"
	t = "c#d#"

	s = "a#c"
	t = "b"

	s = "xywrrmp"
	t = "xywrrmu#p"

	s = "y#fo##f" //f
	t = "y#f#o##f" //f

	s = "gtc#uz#" //gtu
	t = "gtcm##uz#" //gtu

	res := backspaceCompare(s, t)
	fmt.Println("Result is ", res)

}

func backspaceCompare(s string, t string) bool {
	lenT := len(t)
	lenS := len(s)

	var s2, t2 string
	i := 0

	for  {
		if i < lenS {
			if s[i] != '#' {
				s2 += string(s[i])
			} else if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}

		}

		if i < lenT {
			if t[i] != '#' {
				t2 += string(t[i])
			} else if len(t2) > 0 {
				t2 = t2[:len(t2)-1]
			}
		}

		i++

		if i > lenS && i > lenT {
			break
		}
	}

	return s2 == t2
}



