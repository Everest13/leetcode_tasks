package main

import "fmt"

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo","bar"}

	s = "barfoofoobarthefoobarman" //6,9,12
	words = []string{"bar","foo","the"}

	s = "wordgoodgoodgoodbestword"
	words = []string{"word","good","best","good"} //8

	//s = "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	//words = []string{"fooo","barr","wing","ding","wing"} //13

	res := findSubstring(s, words)
	fmt.Println(res)
}
/*
	30. Substring with Concatenation of All Words
	Hard
	https://leetcode.com/problems/substring-with-concatenation-of-all-words/
	Runtime 1394ms Beats 28.85%, Memory 6.52MB Beats 63.19%
 */
func findSubstring(s string, words []string) []int {
	lenWord := len(words[0])
	lenWords := len(words)
	n := len(s)
	if n < lenWord*lenWords {
		return nil
	}

	wordsMap := make(map[string]int, len(words))
	for _, word := range words {
		wordsMap[word]++
	}

	checkConcatenated := func(usedWords map[string]int, j int) bool {
		usedWordsAmount := 1
		for usedWordsAmount < lenWords {
			bunchS := s[j:j+lenWord]
			count, ok := wordsMap[bunchS]
			if !ok {
				return false
			}

			countUsed, ok := usedWords[bunchS]
			if ok && countUsed == count {
				return false
			}

			usedWords[bunchS]++
			usedWordsAmount++
			j = j+lenWord
		}

		return true
	}

	res := []int{}
	lenS := len(s)
	for i:=0; i<lenWord; i++ {

		for j:=i; j<=lenS-lenWord*lenWords; j+=lenWord {
			bunchS := s[j:j+lenWord]
			_, ok := wordsMap[bunchS]
			if !ok {
				continue
			}

			checkRes := checkConcatenated(map[string]int{bunchS: 1}, j+lenWord)
			if checkRes {
				res = append(res, j)
			}
		}
	}

	return res
}
