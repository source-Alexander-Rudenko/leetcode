package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	i := 0
	for i < len(nums) {
		start := nums[i]
		for i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			i++
		}
		if start != nums[i] {
			ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			ans = append(ans, strconv.Itoa(nums[i]))
		}
		i++
	}
	return ans
}

func main() {
	a := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	b := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	fmt.Println(a)
	fmt.Println(b)
}

//TODO в obsidian записать
