package main

import (
	"fmt"
	"sort"
)

/*
	func carFleet(target int, position []int, speed []int) int {
		n := len(position)
		pairs := make([][2]int, n)
		stack := []float64{}
		for i := 0; i < n; i++ {
			pairs[i] = [2]int{position[i], speed[i]}
		}
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i][0] > pairs[j][0]
		})
		for _, p := range pairs {
			time := float64(target-p[0]) / float64(p[1])
			stack = append(stack, time)
			if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
				stack = stack[:len(stack)-1]
			}
		}
		return len(stack)
	}
*/
type Car struct {
	p int
	s int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	road := make([]Car, n)
	for i := 0; i < n; i++ {
		road[i] = Car{position[i], speed[i]}
	}
	sort.Slice(road, func(i, j int) bool {
		return road[i].p > road[j].p
	})
	stack := []float64{}
	for _, car := range road {
		time := float64(target-car.p) / float64(car.s)
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}

func main() {
	a, b := []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet(12, a, b))
}
