package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	mS := make(map[rune]rune, len(s))
	mT := make(map[rune]rune, len(t))

	for i := 0; i < len(s); i++ {
		c1 := rune(s[i])
		c2 := rune(t[i])

		valS, okS := mS[c1]
		valT, okT := mT[c2]

		if !okS && !okT {
			mS[c1] = c2
			mT[c2] = c1
		} else if !(valS == c2 && valT == c1) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))   // false
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
}
