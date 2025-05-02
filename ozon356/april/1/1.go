package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')

		score := make(map[string]int)
		var action string

		for i := 0; i < n; i++ {
			line, _ := in.ReadString('\n')
			line = strings.TrimSpace(line)

			colon := strings.Index(line, ":")
			person := line[:colon]
			rest := strings.TrimSpace(line[colon+1:])

			tokens := strings.Split(rest, " ")

			if tokens[0] == "I" && tokens[1] == "am" {
				if tokens[2] == "not" {
					action = strings.TrimSuffix(tokens[3], "!")
					score[person] -= 1
				} else {
					action = strings.TrimSuffix(tokens[2], "!")
					score[person] += 2
				}
			} else {
				target := tokens[0]
				if tokens[2] == "not" {
					action = strings.TrimSuffix(tokens[3], "!")
					score[target] -= 1
				} else {
					action = strings.TrimSuffix(tokens[2], "!")
					score[target] += 1
				}
				if _, ok := score[person]; !ok {
					score[person] = 0
				}
			}

		}

		maxScore := math.MinInt
		for _, v := range score {
			if v > maxScore {
				maxScore = v
			}
		}

		var arr []string
		for name, val := range score {
			if val == maxScore {
				arr = append(arr, name)
			}
		}
		sort.Strings(arr)

		for _, name := range arr {
			fmt.Fprintf(out, "%s is %s.\n", name, action)
		}
	}
}
