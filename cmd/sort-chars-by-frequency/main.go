package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "Aabb"
	res := frequencySort(str)
	fmt.Println(res)
}

/*
	451. Sort Characters By Frequency
	https://leetcode.com/problems/sort-characters-by-frequency/description/?envType=daily-question&envId=2024-02-07
	Runtime 6ms Beats75.94% Memory 6.42MB Beats 43.32% M
 */
func frequencySort(s string) string {
	charsMap := map[rune]int{}
	for _, char := range s {
		charsMap[char]++
	}

	keys := make([]rune, 0, len(charsMap))
	for char := range charsMap {
		keys = append(keys, char)
	}

	sort.SliceStable(keys, func(i, j int) bool{
		return charsMap[keys[i]] > charsMap[keys[j]]
	})

	var res string
	for _, key := range keys {
		res += strings.Repeat(string(key), charsMap[key])
	}

	return res
}

