package main

import (
	"fmt"
	"strings"
)

func maxVowels(s string, k int) int {
	var sum int
	var ans int
	vovelLetters := "aeiou"
	for i := 0; i < k; i++ {
		if strings.Contains(vovelLetters, string(s[i])) {
			sum++
		}
	}
	ans = sum
	for j := k; j < len(s); j++ {
		if strings.Contains(vovelLetters, string(s[j])) {
			sum++
		}
		if strings.Contains(vovelLetters, string(s[j-k])) {
			sum--
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
func main() {
	a := "abciiidef"
	b := "aeiou"
	c := "leetcode"

	fmt.Println(maxVowels(a, 3))
	fmt.Println(maxVowels(b, 2))
	fmt.Println(maxVowels(c, 3))
}
