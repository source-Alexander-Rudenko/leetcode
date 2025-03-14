/*
A phrase is a palindrome if, after converting all uppercase letters into
lowercase letters and removing all non-alphanumeric characters, it reads the
same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true

Example 2:
Input: s = "race a car"
Output: false

Example 3:
Input: s = " "
Output: true
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	r := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleanedString := r.ReplaceAllString(s, "")
	length := len(cleanedString)
	if length == 0 {
		return true
	}
	cleanedString = strings.ToLower(cleanedString)
	for i, _ := range cleanedString {
		if cleanedString[i] != cleanedString[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	a := "A man, a plan, a canal: Panama"
	b := "race a car"
	c := " "
	fmt.Println(isPalindrome(a))
	fmt.Println(isPalindrome(b))
	fmt.Println(isPalindrome(c))
}
