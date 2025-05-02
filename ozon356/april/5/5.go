package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n, &m)

		histograms := make([][]int, n)
		for i := 0; i < n; i++ {
			histograms[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &histograms[i][j])
			}
		}

		validCount := make(map[int]int)

		for _, h := range histograms {
			sum := h[0] + h[m-1]
			valid := true
			for i := 1; i < m/2; i++ {
				if h[i]+h[m-1-i] != sum {
					valid = false
					break
				}
			}
			if m%2 == 1 {
				if h[m/2]*2 != sum {
					valid = false
				}
			}
			if valid {
				validCount[sum]++
			}
		}

		totalPairs := 0
		for _, cnt := range validCount {
			totalPairs += cnt * (cnt - 1) / 2
		}

		fmt.Fprintln(out, totalPairs)
	}
}
