package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}
	for _, char := range t {
		freq[char]--
		if freq[char] < 0 {
			return false
		}
	}
	return true
}
func main() {
	a := "anagram"
	b := "nagaram"
	c := "rat"
	d := "cat"
	fmt.Println(isAnagram(a, b))
	fmt.Println(isAnagram(c, d))
}
