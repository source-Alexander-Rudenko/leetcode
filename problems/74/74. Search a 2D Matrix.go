package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	top, bot := 0, rows-1
	for top <= bot {
		row := (top + bot) / 2
		if target > matrix[row][cols-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}
	if !(top <= bot) {
		return false
	}
	row := (top + bot) / 2
	l, r := 0, cols-1
	for l <= r {
		mid := (l + r) / 2
		if target > matrix[row][mid] {
			l = mid + 1
		} else if target < matrix[row][mid] {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	a := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Print(searchMatrix(a, 3))
}
