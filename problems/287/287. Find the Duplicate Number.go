package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow2 == slow {
			return slow
		}
	}
}

func main() {
	a := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(a))
}
