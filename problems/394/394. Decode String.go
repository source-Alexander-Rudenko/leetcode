package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
это сам родил костыль на костыле
func decodeString(s string) string {
	stack := []rune{}
	var digit int
	var char rune
	for _, ch := range s {
		if ch != ']' {
			stack = append(stack, ch)
		} else {
			var tmp []rune
			char, stack = stack[len(stack)-1], stack[:len(stack)-1]
			tmp = append(tmp, char)
			char, stack = stack[len(stack)-1], stack[:len(stack)-1]
			for char != '[' {
				tmp = append(tmp, char)
				char, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}
			digit, stack = int(stack[len(stack)-1]-'0'), stack[:len(stack)-1]
			if len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				digit += int(stack[len(stack)-1]-'0') * 10
				stack = stack[:len(stack)-1]
				if len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
					digit += int(stack[len(stack)-1]-'0') * 100
					stack = stack[:len(stack)-1]
				}
			}
			if len(tmp) > 1 {
				for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
					tmp[i], tmp[j] = tmp[j], tmp[i]
				}
			}
			for i := 0; i < digit; i++ {
				stack = append(stack, tmp...)
			}
		}
	}
	return string(stack)
}*/

//neetcode

func decodeString(s string) string {
	var stack []rune
	for _, ch := range s {
		if ch != ']' {
			stack = append(stack, ch)
		} else {
			subString := []rune{}
			for stack[len(stack)-1] != '[' {
				subString = append([]rune{stack[len(stack)-1]}, subString...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			digits := []rune{}
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				digits = append([]rune{stack[len(stack)-1]}, digits...)
				stack = stack[:len(stack)-1]
			}
			n, _ := strconv.Atoi(string(digits))
			for i := 0; i < n; i++ {
				stack = append(stack, subString...)
			}

		}
	}
	return string(stack)
}

func main() {
	a := "3[a]2[bc]"
	b := "3[a2[c]]"
	c := "2[abc]3[cd]ef"
	d := "10[l]"

	fmt.Println(decodeString(d))

	fmt.Println(decodeString(a))
	fmt.Println(decodeString(b))
	fmt.Println(decodeString(c))
}
