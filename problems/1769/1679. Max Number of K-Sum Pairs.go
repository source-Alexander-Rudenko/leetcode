/*
You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
Return the maximum number of operations you can perform on the array.

Example 1:
Input: nums = [1,2,3,4], k = 5
Output: 2

Example 2:
Input: nums = [3,1,3,4,3], k = 6
Output: 1
*/

package main

import (
	"fmt"
)

func maxOperations(nums []int, k int) int {
	minInt := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	mp := make(map[int]int)
	res := 0
	for _, n := range nums {
		mp[n]++
	}
	if count, ok := mp[k/2]; ok && k%2 == 0 {
		res += count / 2
		delete(mp, k/2)
	}
	for key, value := range mp {
		if _, exists := mp[k-key]; exists {
			res += minInt(value, mp[k-key])
			delete(mp, key)
			delete(mp, k-key)
		}
	}
	return res
}

//Решение гения без лишних проходов удалений быстрое шо пиздец
/*
func maxOperations(nums []int, k int) int {
    couples := 0
    counts := make(map[int]int, len(nums))
    for _, v := range nums {
        counts[k-v]++
    }
    for _, v := range nums {
        compliment := k - v
        if counts[v] < 1 || counts[compliment] < 1 {
            continue
        }
        if compliment == v && counts[v] < 2 {
            continue
        }

        counts[v]--
        counts[compliment]--
        couples++
    }
    return couples
}
*/
func main() {
	a := []int{1, 2, 3, 4}
	b := []int{3, 1, 3, 4, 3}

	fmt.Println(maxOperations(a, 5))
	fmt.Println(maxOperations(b, 6))

}
