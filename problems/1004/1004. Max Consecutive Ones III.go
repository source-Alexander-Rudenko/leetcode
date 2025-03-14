package main

import (
	"fmt"
)

//func longestOnes(nums []int, k int) int {
//	window := 0
//	start := 0
//	sum := 0
//	for end := 0; end < len(nums); end++ {
//		sum += nums[end]
//		if end-start+1-sum <= k {
//			window = end - start + 1
//		} else {
//			sum -= nums[start]
//			start++
//		}
//	}
//	return window
//}

func longestOnes(nums []int, k int) int {
	var left, cur, ans int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			cur++
		}
		if cur > k {
			if nums[left] == 0 {
				cur--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func main() {
	a := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	b := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}

	fmt.Println(longestOnes(a, 2))
	fmt.Println(longestOnes(b, 3))
}
