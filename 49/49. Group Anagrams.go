/*Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]
*/

package main

import (
	"fmt"
	"sort"
)

// TODO переписать на массив счетчик букав
func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]int)
	for ind, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		str = string(runes)
		mp[str] = append(mp[str], ind)
	}
	var res [][]string
	var anagramsArr []string
	for _, value := range mp {
		for _, val := range value {
			anagramsArr = append(anagramsArr, strs[val])
		}
		res = append(res, anagramsArr)
		anagramsArr = nil
	}
	return res
}

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	b := []string{""}
	c := []string{"a"}
	fmt.Println(groupAnagrams(a))
	fmt.Println(groupAnagrams(b))
	fmt.Println(groupAnagrams(c))
}
