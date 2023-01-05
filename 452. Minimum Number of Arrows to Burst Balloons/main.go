package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	count := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if end < points[i][0] {
			count++
			end = points[i][1]
		}
	}

	return count
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})) // 2
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))    // 4
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))    // 2
}
