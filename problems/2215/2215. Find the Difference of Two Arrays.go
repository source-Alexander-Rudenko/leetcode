/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

можно еще поменять bool  на пустую структуру, она же вообще ничего не занимает
*/

package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	mp1, mp2 := make(map[int]bool), make(map[int]bool)
	res := make([][]int, 2)

	for _, n := range nums1 {
		mp1[n] = true
	}
	for _, n := range nums2 {
		mp2[n] = true
	}
	for key := range mp1 {
		if !mp2[key] {
			res[0] = append(res[0], key)
		}
	}
	for key := range mp2 {
		if !mp1[key] {
			res[1] = append(res[1], key)
		}
	}
	return res
}

func main() {
	num1, num2 := []int{1, 2, 3}, []int{2, 4, 6}
	num3, num4 := []int{1, 2, 3, 3}, []int{1, 1, 2, 2}
	fmt.Println(findDifference(num1, num2))
	fmt.Println(findDifference(num3, num4))
}
