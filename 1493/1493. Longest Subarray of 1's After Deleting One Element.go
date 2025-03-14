package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	var left, count, ans int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			count++
		}
		if count > 1 {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans - 1
}

func main() {
	a := []int{1, 1, 0, 1}
	b := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	c := []int{1, 1, 1}

	fmt.Println(longestSubarray(a))
	fmt.Println(longestSubarray(b))
	fmt.Println(longestSubarray(c))
}
