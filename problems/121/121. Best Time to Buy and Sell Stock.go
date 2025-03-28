/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a
different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
*/
package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	maxPrice, profit := 0, 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else if profit < maxPrice-prices[i] {
			profit = maxPrice - prices[i]
		}
	}
	return profit
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	b := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(a))
	fmt.Println(maxProfit(b))
}
