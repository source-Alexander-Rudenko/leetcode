/*
Given an integer array nums, return an array answer such that answer[i] is equal to the
product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without
using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	prefixProduct := make([]int, len(nums))
	prefixProduct[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixProduct[i] = nums[i] * prefixProduct[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] = nums[i] * nums[i+1]
	}
	result := make([]int, len(nums))
	result[0] = nums[1]
	result[len(nums)-1] = prefixProduct[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		result[i] = prefixProduct[i-1] * nums[i+1]
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(a))
	fmt.Println(productExceptSelf(b))
}
