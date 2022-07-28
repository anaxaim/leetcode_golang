package main

import "fmt"

func sum(n []int) int {
	var s int
	for _, v := range n {
		s += v
	}
	return s
}

func runningSum(nums []int) []int {
	newNums := make([]int, len(nums))
	for i, _ := range nums {
		newNums[i] = sum(nums[:i+1])
	}
	return newNums
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))     // [1,3,6,10]
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))  // [1,2,3,4,5]
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1})) // [3,4,6,16,17]
}
