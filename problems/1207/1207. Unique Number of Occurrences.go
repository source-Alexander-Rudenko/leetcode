/*
Given an array of integers arr, return true if the number of occurrences of each value in the array
is unique or false otherwise.
*/
package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	countMap := map[int]int{}
	occurrences := map[int]struct{}{}
	for _, num := range arr {
		countMap[num]++
	}
	for _, value := range countMap {
		if _, exists := occurrences[value]; exists {
			return false
		}
		occurrences[value] = struct{}{}
	}
	return true
}

func main() {
	a, b, c := []int{1, 2, 2, 1, 1, 3}, []int{1, 2}, []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(b))
	fmt.Println(uniqueOccurrences(a), uniqueOccurrences(b), uniqueOccurrences(c))
}
