package main

import "fmt"

func sum(sm []int) int {
	s := 0
	for _, v := range sm {
		s += v
	}
	return s
}

func pivotIndex(nums []int) int {
	for i := range nums {
		left_sum := sum(nums[:i])
		right_sum := sum(nums[i+1:])
		if left_sum == right_sum {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))          // -1
	fmt.Println(pivotIndex([]int{2, 1, -1}))         // 0
}
