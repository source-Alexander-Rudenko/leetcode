package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	countArr1 := [26]int{}
	countArr2 := [26]int{}
	size := len(s1)
	if size > len(s2) {
		return false
	}
	for i := 0; i < size; i++ {
		countArr1[s1[i]-'a']++
		countArr2[s2[i]-'a']++
	}

	for i := size; i < len(s2); i++ {
		if countArr2 == countArr1 {
			return true
		}
		countArr2[s2[i]-'a']++
		countArr2[s2[i-size]-'a']--
	}
	if countArr2 == countArr1 {
		return true
	} else {
		return false
	}
}

func main() {
	a, b := "ab", "eidbaooo"
	c, d := "ab", "eidboaoo"
	e, f := "adc", "dcda"
	fmt.Println(checkInclusion(a, b))
	fmt.Println(checkInclusion(c, d))
	fmt.Println(checkInclusion(e, f))
}
