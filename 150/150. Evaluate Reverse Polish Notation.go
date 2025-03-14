/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.
*/

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []string{}
	a, b := "", ""
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			a, stack = stack[len(stack)-1], stack[:len(stack)-1]
			b, stack = stack[len(stack)-1], stack[:len(stack)-1]
			aDidgit, _ := strconv.Atoi(a)
			bDidgit, _ := strconv.Atoi(b)
			switch token {
			case "+":
				stack = append(stack, strconv.Itoa(aDidgit+bDidgit))
			case "-":
				stack = append(stack, strconv.Itoa(aDidgit-bDidgit))
			case "*":
				stack = append(stack, strconv.Itoa(aDidgit*bDidgit))
			case "/":
				stack = append(stack, strconv.Itoa(bDidgit/aDidgit))
			}
		} else {
			stack = append(stack, token)
		}
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}

/*
EDITORAL
func evalRPN(tokens []string) int {
    operations := map[string]func(int, int) int{
        "+": func(a, b int) int { return a + b },
        "-": func(a, b int) int { return a - b },
        "*": func(a, b int) int { return a * b },
        "/": func(a, b int) int { return a / b },
    }
    stack := []int{}
    for _, token := range tokens {
        if operation, exists := operations[token]; exists {
            num1 := stack[len(stack)-2]
            num2 := stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            result := operation(num1, num2)
            stack = append(stack, result)
        } else {
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }
    return stack[0]
}*/

func main() {
	a := []string{"2", "1", "+", "3", "*"}
	b := []string{"4", "13", "5", "/", "+"}

	fmt.Println(evalRPN(a))
	fmt.Println(evalRPN(b))
}
