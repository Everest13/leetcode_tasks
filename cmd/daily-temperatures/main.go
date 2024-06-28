package main

import "fmt"

func main() {
	temperatures := []int{73,74,75,71,69,72,76,73}
	//temperatures = []int{89,62,70,58,47,47,46,76,100,70}

	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}

/*
	739. Daily Temperatures
	Medium
	https://leetcode.com/problems/daily-temperatures/?envType=daily-question&envId=2024-05-11
	Runtime 141ms Beats 46.34%, Memory 11.09MB Beats 55.19%
 */
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n, n)
	st := make([]int, 0, n)
	for i:=0; i<n; i++ {
		j := len(st)
		for j>0 && temperatures[st[j-1]] < temperatures[i] {
			res[st[j-1]] = i - st[j-1]
			j--
		}

		st = st[:j]
		st = append(st, i)
	}

	return res
}

// Runtime 345ms Beats 5.03%, Memory 9.14 MB Beats 82.84%
func dailyTemperatures3(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n, n)
	undetermT := make([]int, 0, n)
	for i:=0; i<n-1; i++ {
		nextT := temperatures[i+1]
		if temperatures[i] < nextT {
			res[i] = 1
			cutIdx := len(undetermT)
			for j:=len(undetermT)-1; j>=0; j-- {
				idxT := undetermT[j]
				if nextT > temperatures[idxT] {
					cutIdx = j
					res[idxT] = i+1-idxT
				}
				continue
			}

			undetermT = undetermT[:cutIdx]
			continue
		}

		undetermT = append(undetermT, i)
	}

	return res
}

// Time Limit Exceeded
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n, n)
	mapT := map[int]int{}
	for i:=0; i<n-1; i++ {
		nextT := temperatures[i+1]
		if temperatures[i] < nextT {
			res[i] = 1

			for idx, temp := range mapT {
				if nextT > temp {
					res[idx] = i+1-idx
					delete(mapT, idx)
				}
			}
			continue
		}

		mapT[i] = temperatures[i]
	}

	return res
}