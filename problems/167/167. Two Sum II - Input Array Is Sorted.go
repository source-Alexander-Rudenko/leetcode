/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order
, find two numbers such that they add up to a specific target number. Let these two numbers
be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array
[index1, index2] of length 2.
The tests are generated such that there is exactly one solution. You may not use the
same element twice.
Your solution must use only constant extra space.

Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]

Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]

Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
*/
package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	// BRUTE FORCE
	//res := make([]int, 2)
	//for left := 0; left < len(numbers); left++ {
	//	for right := left + 1; right < len(numbers) && numbers[left]+numbers[right] <= target; right++ {
	//		if numbers[left]+numbers[right] == target {
	//			res[0], res[1] = left+1, right+1
	//			return res
	//		}
	//	}
	//}
	//return res
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}

func main() {
	a := []int{2, 7, 11, 15}
	b := []int{2, 3, 4}
	c := []int{5, 25, 75}
	fmt.Println(twoSum(a, 9))
	fmt.Println(twoSum(b, 6))
	fmt.Println(twoSum(c, 100))
}
