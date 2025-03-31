package main

import "fmt"

func generateParenthesis(n int) []string {
	stack := []rune{}
	result := []string{}
	var backtracking func(int, int)
	backtracking = func(open int, close int) {
		if open == n && close == n {
			result = append(result, string(stack))
			return
		}
		if open < n {
			stack = append(stack, '(')
			backtracking(open+1, close)
			stack = stack[:len(stack)-1]
		}
		if close < open {
			stack = append(stack, ')')
			backtracking(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtracking(0, 0)
	return result
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
