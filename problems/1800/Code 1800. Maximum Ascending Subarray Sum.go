/*
Given an array of positive integers nums, return the maximum possible sum
of an ascending subarray in nums.
A subarray is defined as a contiguous sequence of numbers in an array.
A subarray [numsl, numsl+1, ..., numsr-1, numsr] is ascending if for
all i where l <= i < r, numsi  < numsi+1. Note that a subarray of size 1 is ascending.

Example 1:
Input: nums = [10,20,30,5,10,50]
Output: 65

Example 2:
Input: nums = [10,20,30,40,50]
Output: 150

Example 3:
Input: nums = [12,17,15,13,10,11,12]
Output: 33
*/
package main

import "fmt"

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		res = max(res, sum)
	}
	return res
}

func main() {
	a := []int{10, 20, 30, 5, 10, 50}
	b := []int{10, 20, 30, 40, 50}
	c := []int{100, 10, 1}
	fmt.Println(maxAscendingSum(a))
	fmt.Println(maxAscendingSum(b))
	fmt.Println(maxAscendingSum(c))
}
