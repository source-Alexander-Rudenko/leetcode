/*
You are given two strings s1 and s2 of equal length. A string swap
is an operation where you choose two indices in a string
(not necessarily different) and swap the characters at these indices.
Return true if it is possible to make both strings equal by
performing at most one string swap on exactly one of the strings.
Otherwise, return false.

Example 1:
Input: s1 = "bank", s2 = "kanb"
Output: true

Example 2:
Input: s1 = "attack", s2 = "defend"
Output: false

Example 3:
Input: s1 = "kelb", s2 = "kelb"
Output: true
*/

package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	var swapable []int32
	for i, char := range s1 {
		if int32(s2[i]) != char {
			swapable = append(swapable, char, int32(s2[i]))
		}
		if len(swapable) > 4 {
			return false
		}
	}
	if len(swapable) == 2 {
		return false
	}
	if len(swapable) == 0 || (swapable[0] == swapable[3] && swapable[1] == swapable[2]) {
		return true
	} else {
		return false
	}
}

func main() {
	s1 := "aa"
	s2 := "ac"
	s3 := "attack"
	s4 := "defend"
	s5 := "kelb"
	s6 := "kelb"
	fmt.Println(areAlmostEqual(s1, s2))
	fmt.Println(areAlmostEqual(s3, s4))
	fmt.Println(areAlmostEqual(s5, s6))
}
