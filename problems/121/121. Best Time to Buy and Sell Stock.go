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
