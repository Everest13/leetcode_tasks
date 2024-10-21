package main

import (
	"fmt"
)

func main() {
	candidates := []int{2,3,6,7}
	target := 7

	candidates = []int{2,3,5}
	target = 8

	candidates = []int{2}
	target = 1

	candidates = []int{7,3,2}
	target = 18

	candidates = []int{2,3,6,7,1,8,4,10}
	target = 18

	res := combinationSum(candidates, target)
	fmt.Println(res)
}

/*
	https://leetcode.com/problems/combination-sum/
	39. Combination Sum
	Runtime 0ms Beats 100.00%, Memory 4.98MB Beats 19.14%
 */
func combinationSum(candidates []int, target int) [][]int {
	candidates = mergeSort(candidates)
	res := [][]int{}
	n := len(candidates)

	recursion := func(i, el, tRest int, subArr []int){}
	recursion = func(i, el, tRest int, subArr []int) {
		elSum := el
		for elSum <= tRest {
			subArr = append(subArr, el)
			rest := tRest - elSum

			for j:=i+1; j<n; j++ {
				nextEl := candidates[j]
				if nextEl <= rest {
					recursion(j, nextEl, rest, subArr)
					continue
				}
				break
			}

			elSum+=el
		}

		subArrSum := 0
		for _, subEl := range subArr {
			subArrSum+=subEl
		}

		if subArrSum == target {
			newArr := make([]int, len(subArr))
			copy(newArr, subArr)

			res = append(res, newArr)
		}
	}

	for i:=0; i<n; i++ {
		el := candidates[i]
		if el > target {
			break
		}
		if el == target {
			res = append(res, []int{target})
			break
		}

		recursion(i, el, target, []int{})
	}


	return res
}

func quickSort(arr []int) []int {
	recursion := func(subArr []int) []int {return []int{}}
	recursion = func(subArr []int) []int{
		n := len(subArr)
		if n == 1 {
			return []int{subArr[0]}
		}

		el := subArr[0]
		left := make([]int, 0, n)
		right := make([]int, 0, n)
		for i:=1; i<n; i++ {
			if subArr[i] > el {
				right = append(right, subArr[i])
				continue
			}
			left = append(left, subArr[i])
		}

		resLeft, resRight := make([]int, 0, len(left)), make([]int, 0, len(right))
		if len(left) > 0 {
			resLeft = recursion(left)
		}
		if len(right) > 0 {
			resRight = recursion(right)
		}

		return append(append(resLeft, el), resRight...)
	}

	return recursion(arr)
}

func mergeSort(arr []int) []int {
	recursion := func(subArr []int) []int {return []int{}}
	recursion = func(subArr []int) []int {
		n := len(subArr)
		if n == 2 {
			one := subArr[0]
			two := subArr[1]
			if one < two {
				return []int{one, two}
			}

			return []int{two, one}
		}


		mid := n/2
		firstHalf := subArr[0:mid]
		if mid > 2 {
			firstHalf = recursion(subArr[0:mid])
		}
		secHalf := recursion(subArr[mid:])

		//sort res
		res := make([]int, 0, n)
		i,j := 0,0
		for len(res) < n {
			if i >= len(firstHalf) || j >= len(secHalf) {
				if i <= len(firstHalf) {
					res = append(res, firstHalf[i:]...)
				}
				if j <= len(secHalf) {
					res = append(res, secHalf[j:]...)
				}
				break
			}

			if firstHalf[i] < secHalf[j] {
				res = append(res, firstHalf[i])
				i++
				continue
			}

			res = append(res, secHalf[j])
			j++
		}

		return res
	}

	return recursion(arr)
}