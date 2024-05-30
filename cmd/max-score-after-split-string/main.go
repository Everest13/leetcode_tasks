package main

func main() {

}


/*
	1422. Maximum Score After Splitting a String
	https://leetcode.com/problems/maximum-score-after-splitting-a-string/?envType=daily-question&envId=2024-05-11
	Runtime 1ms Beats80.46% Memory 2.17MB Beats97.70% E
*/
func maxScore(s string) int {
	pref := 0
	for i:=0; i<len(s); i++ {
		if s[i] == '1' {
			pref++
		}
	}

	maxS := 0
	suff := 0
	for i:=0; i<len(s)-1; i++ {
		if s[i] == '0' {
			suff++
		} else {
			pref--
		}
		if pref+suff > maxS {
			maxS = pref+suff
		}
	}

	return maxS
}




