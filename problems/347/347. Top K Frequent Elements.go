/*
Given an integer array nums and an integer k, return the
k most frequent elements. You may return the answer in any order.
Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]
*/
package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	for _, i := range nums {
		count[i]++
	}
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}
	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func main() {
	a := []int{1, 1, 1, 2, 2, 3}
	b := []int{1}
	fmt.Println(topKFrequent(a, 2))
	fmt.Println(topKFrequent(b, 1))
}
