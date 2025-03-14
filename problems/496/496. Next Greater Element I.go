/*
The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater
element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.
*/

package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for ind, num := range nums1 {
		mp[num] = ind
	}
	stack, res := []int{}, make([]int, len(nums1))
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if _, exists := mp[nums2[i]]; exists {
			if len(stack) == 0 {
				res[mp[nums2[i]]] = -1
			} else {
				res[mp[nums2[i]]] = stack[len(stack)-1]
			}
		}
		stack = append(stack, nums2[i])
	}
	return res
}
func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
