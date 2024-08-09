package main

import "fmt"

func main() {
	points := [][]int{{0,0},{0,1},{1,0},{0,2},{2,0}}

	points = [][]int{{1,0},{0,0},{0,1}} //0.5

	res := largestTriangleArea(points)
	fmt.Println(res)
}

/*
	812. Largest Triangle Area
	https://leetcode.com/problems/largest-triangle-area/
	Runtime 4ms Beats 54.55%, Memory 2.38MB Beats 77.27%
 */
func largestTriangleArea(points [][]int) float64 {
	square := func(a1, a2, a3 []int) float64 {
		s := (a1[0]-a3[0])*(a2[1]-a3[1]) - (a1[1] - a3[1])*(a2[0] - a3[0])
		if s < 0 {
			s *= -1
		}

		return float64(s)/2
	}

	if len(points) < 4 {
		return square(points[0], points[1], points[2])
	}

	var maxS float64
	for i:=0; i<len(points); i++ {
		for j:=1; j<len(points); j++ {
			for k:=2; k<len(points); k++ {
				s := square(points[i], points[j], points[k])
				if maxS < s {
					maxS = s
				}
			}
		}
	}

	return maxS
}
