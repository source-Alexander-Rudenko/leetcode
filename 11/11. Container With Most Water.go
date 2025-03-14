package main

import "fmt"

func maxArea(height []int) int {
	area, answer, l, r := 0, 0, 0, len(height)-1
	for l < r {
		area = min(height[l], height[r]) * (r - l)
		answer = max(area, answer)
		if height[l] >= height[r] {
			r--
		} else {
			l++
		}
	}
	return answer
}

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	b := []int{1, 1}
	fmt.Println(maxArea(a))
	fmt.Println(maxArea(b))
}
