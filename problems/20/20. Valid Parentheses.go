/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
package main

import "fmt"

/*func isValid(s string) bool {
	st := []int32{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			st = append(st, char)
		} else if len(st) > 0 && char == ')' {
			if st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		} else if len(st) > 0 {
			if st[len(st)-1] != char-2 {
				return false
			}
			st = st[:len(st)-1]
		} else {
			return false
		}
	}
	if len(st) == 0 {
		return true
	} else {
		return false
	}
}*/

func isValid(s string) bool {
	mapping := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make([]rune, 0)
	for _, char := range s {
		if val, ok := mapping[char]; ok {
			topElement := '#'
			if len(stack) != 0 {
				topElement, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}
			if topElement != val {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func main() {
	a := "]"
	b := "()[]{}"
	c := "(]"
	d := "([])"
	fmt.Println(isValid(a))
	fmt.Println(isValid(b))
	fmt.Println(isValid(c))
	fmt.Println(isValid(d))
}
