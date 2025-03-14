/*
You are given a string s.
Your task is to remove all digits by doing this operation repeatedly:
Delete the first digit and the closest non-digit character to its left.
Return the resulting string after removing all digits.

Example 1:
Input: s = "abc"
Output: "abc"

Example 2:
Input: s = "cb34"
Output: ""
*/
package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	st := []int32{}
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			st = append(st, char)
		} else if unicode.IsDigit(char) {
			st = st[:len(st)-1]
		}
	}
	return string(st)
}

func main() {
	a := "abc"
	b := "cb34"
	fmt.Println(clearDigits(a))
	fmt.Println(clearDigits(b))
}
