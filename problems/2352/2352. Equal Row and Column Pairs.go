/*
Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row
ri and column cj are equal.
A row and column pair is considered equal if they contain the same elements in the same order
(i.e., an equal array).

Example 1:

Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	mp := make(map[string]int)
	count := 0
	for _, row := range grid {
		var strArr []string
		for _, v := range row {
			strArr = append(strArr, strconv.Itoa(v))
		}
		str := strings.Join(strArr, ",")
		mp[str]++
	}

	for i := 0; i < len(grid); i++ {
		var strArr []string
		for j := 0; j < len(grid); j++ {
			strArr = append(strArr, strconv.Itoa(grid[j][i]))
		}
		str := strings.Join(strArr, ",")
		if val, ok := mp[str]; ok {
			count += val
		}
	}
	return count
}

func main() {
	grid1 := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}

	grid := [][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}
	fmt.Println(equalPairs(grid1))
	fmt.Println(equalPairs(grid))
}
