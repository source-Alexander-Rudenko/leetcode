/*
Given an unsorted array of integers nums, return the length of the longest
consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/
package main

import "fmt"

func longestConsecutive(nums []int) int {
	longest, count := 0, 0

	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	for key := range mp {
		if _, exist := mp[key-1]; !exist {
			count++
			for i := key + 1; key < len(nums); i++ {
				if val := mp[i]; val {
					count++
				} else {
					break
				}
			}
			longest = max(longest, count)
			count = 0
		}
	}
	return longest
}

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	b := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	c := []int{0}
	d := []int{}
	fmt.Println(longestConsecutive(a))
	fmt.Println(longestConsecutive(b))
	fmt.Println(longestConsecutive(c))
	fmt.Println(longestConsecutive(d))
}
