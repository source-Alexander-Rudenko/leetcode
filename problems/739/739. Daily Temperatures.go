/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i]
is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for
which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack [][2]int
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1][0] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1][1] - i
		}
		stack = append(stack, [2]int{temperatures[i], i})
	}
	return result
}

func main() {
	a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	b := []int{30, 40, 50, 60}
	c := []int{30, 60, 90}
	fmt.Println(dailyTemperatures(a))
	fmt.Println(dailyTemperatures(b))
	fmt.Println(dailyTemperatures(c))
}
