package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Point struct {
	x, y int
}

var directions = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // по вертикали и горизонтали
	{-1, -1}, {1, 1}, {1, -1}, {-1, 1}, // по диагоналям
}

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n, &m)
		in.ReadString('\n')

		grid := make([][]int, n)
		for i := range grid {
			grid[i] = make([]int, m)
		}

		for i := 0; i < n; i++ {
			str, _ := in.ReadString('\n')
			for j := 0; j < m; j++ {
				grid[i][j], _ = strconv.Atoi(string(str[j]))
			}
		}

		var x, y, p int
		fmt.Fscan(in, &x, &y, &p)
		x--
		y--

		quake := make([][]int, n)
		for i := 0; i < n; i++ {
			quake[i] = make([]int, m)
		}

		queue := []Point{{x, y}}
		quake[x][y] = p

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			currX, currY := curr.x, curr.y
			currPower := quake[currX][currY]

			if currPower <= 0 {
				continue
			}

			for _, dir := range directions {
				nx, ny := currX+dir.x, currY+dir.y
				if nx >= 0 && ny >= 0 && nx < n && ny < m && quake[nx][ny] == 0 {
					quake[nx][ny] = currPower - 1
					queue = append(queue, Point{nx, ny})
				}
			}
		}

		houseID := make([][]int, n)
		for i := 0; i < n; i++ {
			houseID[i] = make([]int, m)
		}

		id := 1
		idToDestroyed := make(map[int]bool)

		var dfs func(x, y, id int)
		dfs = func(x, y, id int) {
			houseID[x][y] = id
			for _, dir := range directions[:4] {
				nx, ny := x+dir.x, y+dir.y
				if nx >= 0 && ny >= 0 && nx < n && ny < m && houseID[nx][ny] == 0 && grid[nx][ny] > 0 {
					dfs(nx, ny, id)
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] > 0 && houseID[i][j] == 0 {
					dfs(i, j, id)
					shouldDestroy := false
					for x := 0; x < n; x++ {
						for y := 0; y < m; y++ {
							if houseID[x][y] == id {
								if quake[x][y] > grid[x][y] {
									shouldDestroy = true
									break
								}
							}
						}
					}
					if shouldDestroy {
						idToDestroyed[id] = true
					}
					id++
				}
			}
		}
		fmt.Fprintln(out, len(idToDestroyed))
	}
}
