package main

import (
	"fmt"
)

func main() {
	t := "ababcbacadefegdehijhklij"
	//t = "caedbdedda"
	//t = "jyb"
	
	res := partitionLabels(t)
	fmt.Println(res)
}

/*
	763. Partition Labels
	Medium
	https://leetcode.com/problems/partition-labels/description/
 */
func partitionLabels(s string) []int {
	lenS := len(s)
	last := make(map[byte]int, lenS)
	for i:=0; i<lenS; i++ {
		last[s[i]] = i
	}

	res := []int{}
	firsId := 0
	lastId := 0
	for i:=0; i<lenS; i++ {
		if last[s[i]] > lastId {
			lastId = last[s[i]]
		}

		if i == lastId {
			res = append(res, lastId-firsId+1)
			firsId = lastId+1
		}
	}

	return res
}

