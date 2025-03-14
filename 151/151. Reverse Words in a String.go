/*
Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated
by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words.
Do not include any extra spaces.

Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/
package main

import (
	"fmt"
	"slices"
	"strings"
)

//	func reverseWords(s string) string {
//		wrd := ""
//		var res []string
//		for i := range s {
//			if string(s[i]) != " " {
//				wrd += string(s[i])
//				i++
//			} else if string(s[i]) == " " && i > 0 && string(s[i-1]) != " " {
//				res = append(res, wrd)
//				wrd = ""
//				i++
//			} else {
//				i++
//			}
//		}
//		if wrd != "" {
//			res = append(res, wrd)
//		}
//		slices.Reverse(res)
//		return strings.Join(res, " ")
//	}
func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
func main() {
	a := "the sky is blue"
	b := "  hello world  "
	c := "a good   example"
	fmt.Println(reverseWords(a))
	fmt.Println(reverseWords(b))
	fmt.Println(reverseWords(c))
}
