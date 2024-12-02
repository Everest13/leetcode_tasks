package main

import "fmt"

func main() {
	sentence := "i love eating burger"
	searchWord := "burg"

	res := isPrefixOfWord(sentence, searchWord)
	fmt.Println(res)
}

/*
	1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
	Easy
	https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/
 */
func isPrefixOfWord(sentence string, searchWord string) int {
	lenSearch := len(searchWord)
	lenSent := len(sentence)
	countSentWord := 0
	idx := -1

	i,j:=0,0
	for i<len(sentence) {
		countSentWord++

		for j<lenSearch && i<lenSent && sentence[i] == searchWord[j] {
			i++
			j++
		}

		if j == lenSearch && idx == -1 {
			idx = countSentWord
		}

		j=0
		for i<lenSent {
			if sentence[i] == 32 {
				i++
				break
			}
			i++
		}
	}

	return idx
}
