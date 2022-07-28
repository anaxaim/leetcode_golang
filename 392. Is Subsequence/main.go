package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, 0

	for j < len(t) {
		if s[i] == t[j] {
			i++
		}

		if i == len(s) {
			return true
		}
		j++
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))    // true
	fmt.Println(isSubsequence("axc", "ahbgdc"))    // false
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa")) // false
}
