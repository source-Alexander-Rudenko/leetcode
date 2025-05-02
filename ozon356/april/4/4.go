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
		var n, k int
		fmt.Fscan(in, &n, &k)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		diff := make([]int, n+2)
		cur := 0
		ok := true

		for i := 0; i < n; i++ {
			cur += diff[i]
			a[i] -= cur

			if a[i] < 0 {
				ok = false
				break
			}
			if a[i] == 0 {
				continue
			}
			if i+k > n {
				ok = false
				break
			}
			cur += a[i]
			diff[i+k] -= a[i]
		}

		if ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
