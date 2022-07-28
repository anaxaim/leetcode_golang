package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	maxCount := 0
	haveOne := false
	for _, v := range m {
		if v != 1 {
			if v%2 == 0 {
				maxCount += v
			} else {
				maxCount += v - 1
				haveOne = true
			}
		} else {
			haveOne = true
		}
	}

	if haveOne {
		maxCount++
	}

	return maxCount
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))  // 7
	fmt.Println(longestPalindrome("a"))         // 1
	fmt.Println(longestPalindrome("ccc"))       // 3
	fmt.Println(longestPalindrome("ababababa")) // 9
}
