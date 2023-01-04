package main

import "fmt"

func minimumRounds(tasks []int) int {
	if len(tasks) == 1 {
		return -1
	}

	result := 0
	count := make(map[int]int)

	for _, t := range tasks {
		count[t]++
	}

	for _, v := range count {
		for v > 0 {
			switch {
			case v%3 == 0:
				result = result + v/3
				v = 0
			case v > 4:
				result++
				v -= 3
			case v%2 == 0:
				result = result + v/2
				v = 0
			default:
				return -1
			}
		}
	}

	return result
}

func main() {
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4})) // 4
	fmt.Println(minimumRounds([]int{2, 3, 3}))                      // -1
}
