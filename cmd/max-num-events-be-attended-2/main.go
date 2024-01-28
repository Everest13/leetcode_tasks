package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
	1751. Maximum Number of Events That Can Be Attended II
	Hard
	https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/description/?envType=list&envId=r57h7kzc
 */
func main() {
	events := [][]int{{1,1,1}, {2,2,2}, {3,3,3}, {4,4,4}} //k=3, exp=9
	k := 3

	//events = [][]int{{1,2,4},{2,6,7},{2,9,10}, {8,9,10}} //k=4 exp=
	//k=4
	events = [][]int{{13,47,98},{49,74,28},{46,63,74},{56,71,94},{12,51,21},{12,58,50},{27,68,97},{4,12,37},{13,17,17},
		{57,59,95},{45,98,53},{41,52,86},{40,51,94},{14,44,49},{10,66,91},{56,63,92},{30,48,74},{21,80,55},
		{16,37,43},{5,69,67},{81,100,87}} //exp=356 k=12

	k = 12

	//events = [][]int{{1,2,4},{3,4,3},{2,3,1}} //k=2 exp=7
	//k=2

	//events = [][]int{{1,2,4},{3,4,3},{2,3,10}} //k=2 exp=10


	//events = [][]int{{13,70,78},{39,79,2},{21,50,66},{23,87,50},{23,62,5},{29,41,81},{32,52,69},{22,25,61},{86,92,59},
	//	{39,50,35},{43,56,27},{83,90,5},{87,95,74},{8,100,71},{25,96,5},{64,68,43},{33,98,100},{8,99,17},{7,48,91},{52,94,54},
	//	{32,48,45},{47,57,51},{4,74,50},{17,95,18},{53,61,33},{22,99,92}}
	//k = 24

	//events = [][]int{{32,68,21},{13,97,27},{5,41,11},{16,35,8},{37,83,79},{81,90,23},{20,94,84},{17,54,91},{11,93,72},
	//	{47,89,1},{71,93,20},{53,65,20},{28,39,62},{9,65,49},{21,22,33},{51,66,69},{39,88,52}} //exp=91
	//k=1

	res := maxValue(events, k)

	fmt.Println(res)
}

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	lenEvents := len(events)
	upperBoundSearch := func(j, end int) int {
		first := j
		last := lenEvents - 1
		middle := first + (last - first)/2
		//binary upper bound search
		for first - middle > 2  {
			if  end > events[middle][0] {
				first = middle + 1
				continue
			}
			last = middle - 1
			middle = first + (last - first)/2
		}

		for end >= events[first][0] {
			first++
		}

		return first
	}

	res := 0
	lastEventStart := events[lenEvents-1][0]
	chains := [][]int{} //remove

	//i - добавленный уже элемент
	//vals []int - slice of i events
	// usedChains - использованные цепочки уже, уходим от повторов
	rec := func(i int, vals []int, recNum int, usedChains map[int][][]int) {}
	rec = func(i int, vals []int, recNum int, usedChains map[int][][]int) {
		event := events[i] //уже лежит в vals
		j := i+1 //следующий

		//find next event
		for j < lenEvents {
			lastJ := j
			//если нет след эл-та - выходим
			if event[1] > lastEventStart  {
				break
			}

			// try to take next
			if event[1] < events[j][0] {
				event = events[j]
				vals = append(vals, j)
				j++
				continue
			}

			// upper_bound search
			j = upperBoundSearch(j, event[1])
			if j == -1 { //TODO нет элемента?
				break
			}

			// запускаем рекурсии только на разветвления
			// на элементы которые пересекаются с event[j]
			// чтобы проверить все возможные ветки
			// надо убрать повторы, которые берутся при проваливании в рекурсии
			// (гепы наслаиваются, и некоторые значения повторяются)
			lastValEl := len(vals)-1
			if lastValEl > 0 {

				for m:=lastJ; m<j; m++ {
					localVals := make([]int, lastValEl)
					copy(localVals, vals[:lastValEl])

					test := events[m]
					_=test

					localVals = append(localVals, m)
					recNum++
					rec(m, localVals, recNum, usedChains)
				}

			}

			event = events[j]
			vals = append(vals, j)
			j++
		}



		chains = append(chains, vals)
		//считаем общее val chain
		if len(vals) > k {
			sort.Slice(vals, func(i, j int) bool {
				return events[i][2] > events[j][2]
			})
		}

		localRes := 0
		for j:=0; j<k && j<len(vals); j++ {
			localRes+=events[j][2]
		}
		if localRes > res {
			res = localRes
		}
	}

	vals := []int{0}
	//идем по всем значениям
	for i:=0; i<lenEvents; i++ {
		rec(0, vals, i, map[int][][]int{})
	}

	return res
}

func maxValue2(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	res := 0
	lenEvents := len(events)
	chains := [][]int{}
	bigChains := [][][]int{}

	rec := func(i int, vals []int, usedChains map[int][][]int) {}
	rec = func(i int, vals []int, usedChains map[int][][]int) {

		j:=i+1
		for j < lenEvents {
			lastValEl := len(vals)-1

			if events[vals[lastValEl]][1] < events[j][0] {
				vals = append(vals, j)
				j++
				continue
			}

			if lastValEl < 1 {
				j++
				continue
			}

			localVals := make([]int, lastValEl)
			copy(localVals, vals[:lastValEl])

			//todo skip copies
			isEqual := false
			jChains, ok := usedChains[j] //map[int][][]int
			if !ok {
				usedChains[j] = [][]int{localVals}
			} else {
				for _, jChain := range jChains {
					if slices.Equal(jChain, localVals) {
						isEqual = true
						continue
					}
				}
			}
			if isEqual {
				j++
				continue
			}
			usedChains[j] = append(usedChains[j], localVals)

			localVals = append(localVals, j)
			rec(j, localVals, usedChains)
			j++
		}

		if len(vals) > k {
			sort.Slice(vals, func(i, j int) bool {
				return events[i][2] < events[j][2]
			})
		}

		chains = append(chains, vals)

		localRes := 0
		bigChain := [][]int{}
		for j:=0; j<k && j<len(vals); j++ {
			localRes+=events[vals[j]][2]
			bigChain = append(bigChain, events[vals[j]])
		}

		bigChains = append(bigChains, bigChain)

		if localRes > res {
			res = localRes
		}
	}

	//идем по всем значениям
	for i:=0; i<lenEvents; i++ {
		rec(i, []int{i}, map[int][][]int{})
	}

	return res
}

/*
	Worked
	Time Limit Exceeded: 97.857494ms-2.272893521s
*/
func maxValue3(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	res := 0
	lenEvents := len(events)
	chains := [][][]int{}
	rec := func(i int, vals [][]int) {}

	rec = func(i int, vals [][]int) {
		//todo почему беруться повторы
		for i < lenEvents {
			lastValEl := len(vals)-1
			// check intersection
			if lastValEl >= 0 && vals[lastValEl][1] >= events[i][0] {
				localVals := make([][]int, lastValEl)
				copy(localVals, vals[:lastValEl])

				rec(i, localVals)
				i++
				continue
			}

			vals = append(vals, events[i])
			i++
		}



		chains = append(chains, vals)
		if len(vals) > k {
			sort.Slice(vals, func(i, j int) bool {
				return vals[i][2] > vals[j][2]
			})
		}

		localRes := 0
		for j:=0; j<k && j<len(vals); j++ {
			localRes+=vals[j][2]
		}

		if localRes > res {
			res = localRes
		}
	}

	rec(0, [][]int{})

	return res
}
