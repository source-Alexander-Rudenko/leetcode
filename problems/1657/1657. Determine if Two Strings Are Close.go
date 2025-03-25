/*
Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do
the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
Hint 1
Operation 1 allows you to freely reorder the string.
Hint 2
Operation 2 allows you to freely reassign the letters' frequencies.
*/

package main

import (
	"fmt"
	"sort"
)

/*
	func closeStrings(word1 string, word2 string) bool {
		mp, mp2, countMap := make(map[rune]int), make(map[rune]int), make(map[int]int)
		if len(word1) != len(word2) {
			return false
		}
		for _, ch := range word2 {
			mp2[ch]++
		}
		for _, ch := range word1 {
			if _, ok := mp2[ch]; !ok {
				return false
			}
			mp[ch]++
		}
		for key, value := range mp {
			if _, ok := mp2[key]; !ok {
				return false
			}
			countMap[value]++
		}
		for _, val := range mp2 {
			countMap[val]--
		}
		for _, val := range countMap {
			if val != 0 {
				return false
			}
		}
		return true
	}

это мое творческое, а вот эдиториал с битовым сдвигом, без мап без нихуя
*/
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1Arr, word2Arr := make([]int, 26), make([]int, 26)
	word1Bit, word2Bit := 0, 0
	for _, char := range word1 {
		word1Arr[char-'a']++
		word1Bit = word1Bit | 1<<(char-'a')
	}
	for _, char := range word2 {
		word2Arr[char-'a']++
		word2Bit = word2Bit | 1<<(char-'a')
	}
	if word1Bit != word2Bit {
		return false
	}
	sort.Ints(word1Arr)
	sort.Ints(word2Arr)
	for i := 0; i < 26; i++ {
		if word1Arr[i] != word2Arr[i] {
			return false
		}
	}
	return true
}

func main() {
	a, b := "abc", "bca"
	c, d := "cabbba", "abbccc"
	fmt.Println(closeStrings(a, b))
	fmt.Println(closeStrings(c, d))
}
