/*
Given an integer array nums, return true if any value appears at least twice
in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, val := range nums {
		if seen[val] {
			return true
		}
		seen[val] = true
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 1}
	b := []int{1, 2, 3, 4}
	c := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(a))
	fmt.Println(containsDuplicate(b))
	fmt.Println(containsDuplicate(c))
}
