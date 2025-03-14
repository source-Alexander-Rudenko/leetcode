package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	var ans int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans = sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		ans = max(ans, sum)
	}
	return float64(ans) / float64(k)
}

func main() {
	a := []int{0, 1, 1, 3, 3}
	b := []int{5}
	fmt.Println(findMaxAverage(a, 4))
	fmt.Println(findMaxAverage(b, 1))
}
