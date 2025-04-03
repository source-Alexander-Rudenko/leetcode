/*
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	mp := make(map[rune]struct{})
	arr := []rune(s)
	l, result := 0, 1
	for r, ch := range arr {
		if _, ok := mp[ch]; ok {
			for arr[l] != ch {
				delete(mp, arr[r])
				l++
			}
			delete(mp, arr[r])
			l++
		}
		mp[ch] = struct{}{}
		if (r-l)+1 > result {
			result = r - l + 1
		}
	}
	return result
}

func main() {
	a := "abcabcbb"
	b := "bbbbb"
	c := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(a))
	fmt.Println(lengthOfLongestSubstring(b))
	fmt.Println(lengthOfLongestSubstring(c))
}
