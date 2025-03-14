package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n)

		// Чтение пропускных способностей серверов
		throughputs := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&throughputs[j])
		}

		// Чтение размеров изображений
		fmt.Scan(&m)
		weights := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&weights[j])
		}

		// Создадим список серверов с их индексами
		servers := make([][2]int, n)
		for j := 0; j < n; j++ {
			servers[j] = [2]int{throughputs[j], j} // [пропускная способность, индекс сервера]
		}

		// Сортируем серверы по пропускной способности
		sort.Slice(servers, func(i, j int) bool {
			return servers[i][0] < servers[j][0]
		})

		// Создаем список изображений с их индексами
		images := make([][2]int, m)
		for j := 0; j < m; j++ {
			images[j] = [2]int{weights[j], j} // [размер изображения, индекс изображения]
		}

		// Сортируем изображения по их размеру
		sort.Slice(images, func(i, j int) bool {
			return images[i][0] < images[j][0]
		})

		// Массив для хранения серверов, на которых будут храниться изображения
		serverAssignments := make([]int, m)

		// Массив для хранения времени доставки для каждого изображения на каждом сервере
		times := make([][]int, m)
		for j := 0; j < m; j++ {
			times[j] = make([]int, n)
			for k := 0; k < n; k++ {
				times[j][k] = int(math.Ceil(float64(images[j][0]) / float64(servers[k][0])))
			}
		}

		// Распределение изображений по серверам
		// Каждое изображение получит сервер с минимальным временем доставки
		for j := 0; j < m; j++ {
			minTime := math.MaxInt
			bestServer := -1
			for k := 0; k < n; k++ {
				if times[j][k] < minTime {
					minTime = times[j][k]
					bestServer = servers[k][1] + 1 // Индекс сервера + 1 (для вывода)
				}
			}
			serverAssignments[images[j][1]] = bestServer
		}

		// Для минимизации разницы между самым быстрым и самым медленным временем доставки
		// нам нужно посчитать разницу времени между самым быстрым и самым медленным изображением
		minTime, maxTime := math.MaxInt, -1
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				if times[j][k] < minTime {
					minTime = times[j][k]
				}
				if times[j][k] > maxTime {
					maxTime = times[j][k]
				}
			}
		}

		// Разница времени
		diff := maxTime - minTime

		// Выводим минимальную разницу и распределение изображений
		fmt.Println(diff)
		for _, server := range serverAssignments {
			fmt.Printf("%d ", server)
		}
		fmt.Println()
	}
}
