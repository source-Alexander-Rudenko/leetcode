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

/* топовое решение
func equalPairs(grid [][]int) int {
	rows := make(map[string]int)
	cols := make(map[string]int)
	rowBuf := make([]byte, len(grid)*4)
	colBuf := make([]byte, len(grid)*4)
	for i := range grid {
		for j := range grid[i] {
			binary.BigEndian.PutUint32(rowBuf[j*4:], uint32(grid[i][j]))
			binary.BigEndian.PutUint32(colBuf[j*4:], uint32(grid[j][i]))
		}
		rows[string(rowBuf)]++
		cols[string(colBuf)]++
	}

	count := 0
	for k := range cols {
		count += rows[k] * cols[k]
	}

	return count
}

func equalPairs(grid [][]int) int {
    rows := map[string]int{}
    cols := map[string]int{}
    rowBuf := make([]rune, len(grid))
    colBuf := make([]rune, len(grid))
    res := 0

    for i, row := range grid {
        for j, cell := range row {
            rowBuf[j] = rune(cell)
            colBuf[j] = rune(grid[j][i])
        }

        rows[string(rowBuf)]++
        cols[string(colBuf)]++
    }

    for s := range rows {
        res += rows[s] * cols[s]
    }

    return res
}
*/

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
