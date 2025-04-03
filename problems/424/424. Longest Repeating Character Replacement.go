package main

import "fmt"

func characterReplacement(s string, k int) int {
	arr := []rune(s)
	freqArr := make([]int, 26)
	start, maxFrequency, longestSubstring := 0, 0, 0
	for end, char := range arr {
		charPosition := char - 'A'
		freqArr[charPosition]++
		if maxFrequency < freqArr[charPosition] {
			maxFrequency = freqArr[charPosition]
		}
		if !(end+1-start-maxFrequency <= k) {
			freqArr[arr[start]-'A']--
			start++
		}
		longestSubstring = end + 1 - start
	}
	return longestSubstring
}

func main() {
	a, b := "ABAB", "AABABBA"
	c := "AABA"
	fmt.Println(characterReplacement(c, 0))
	fmt.Println(characterReplacement(a, 2))
	fmt.Println(characterReplacement(b, 1))
}
