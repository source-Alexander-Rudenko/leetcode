/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23
*/
package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, n := range piles {
		if n > r {
			r = n
		}
	}
	res := r

	for l <= r {
		k := (l + r) / 2
		totalTime := 0
		for _, n := range piles {
			totalTime += int(math.Ceil(float64(n) / float64(k)))
		}
		if totalTime <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}

func main() {
	a := []int{3, 6, 7, 11}
	b := []int{30, 11, 23, 4, 20}
	c := []int{30, 11, 23, 4, 20}
	d := []int{312884470}
	e := []int{2, 2}
	fmt.Println(minEatingSpeed(e, 4))
	fmt.Println(minEatingSpeed(d, 968709470))
	fmt.Println(minEatingSpeed(a, 8), minEatingSpeed(b, 5), minEatingSpeed(c, 6))
}
