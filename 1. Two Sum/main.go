package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, vi := range nums {
		for j, vj := range nums[i+1:] {
			if target-vi == vj {
				return []int{i, j + i + 1}
			}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, val := range nums {
		if mp[target-val] != 0 {
			return []int{mp[target-val] - 1, i}
		}
		mp[val] = i + 1
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]

	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum2([]int{3, 3}, 6))         // [0,1]
}
