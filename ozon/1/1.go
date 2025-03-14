//package main
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//
//	var t int
//	fmt.Fscan(in, &t)
//
//	for i := 0; i < t; i++ {
//		var n, m int
//		fmt.Fscan(in, &n, &m)
//
//		switch {
//		case n == 1:
//			fmt.Fprintln(out, 1)
//			fmt.Fprintln(out, "1 1 R")
//		case m == 1:
//			fmt.Fprintln(out, 1)
//			fmt.Fprintln(out, "1 1 D")
//		case n == m || n > m:
//			fmt.Fprintln(out, 2)
//			fmt.Fprintln(out, "1 1 D")
//			fmt.Fprintf(out, "%d %d U\n", n, m)
//		case n < m:
//			fmt.Fprintln(out, 2)
//			fmt.Fprintln(out, "1 1 R")
//			fmt.Fprintf(out, "%d %d L\n", n, m)
//		}
//	}
//}
