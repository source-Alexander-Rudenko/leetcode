package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pop(s []string) string {
	n := len(s) - 1
	char := s[n]
	s = s[:n]
	return char
}

func decodeString(s string) string {
	var stack []string
	var subString, tmp string
	var didgit int
	arr := strings.Split(s, "")
	for _, char := range arr {
		if char != "]" {
			stack = append(stack, char)
		} else {
			subString = pop(stack)
			tmp = pop(stack)
			for tmp != "[" {
				subString = tmp + subString
				tmp = pop(stack)
			}
			_ = pop(stack)
			tmp = pop(stack)
			didgit, _ = strconv.Atoi(tmp)
			subString = strings.Repeat(subString, didgit)
			stack = append(stack, subString)
		}
	}
	return strings.Join(stack, "")
}

func main() {
	a := "3[a]2[bc]"
	b := "3[a2[c]]"
	c := "2[abc]3[cd]ef"

	fmt.Println(decodeString(a))
	fmt.Println(decodeString(b))
	fmt.Println(decodeString(c))
}
