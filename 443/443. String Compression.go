package main

import (
	"fmt"
	"strconv"
)

//	func compress(chars []byte) int {
//		count := 1
//		firstIndex, seccondIndex := 0, 0
//		var res []byte
//		var countStr string
//		for firstIndex < len(chars)-1 {
//			if chars[firstIndex] == chars[firstIndex+1] {
//				count++
//				seccondIndex = firstIndex + 1
//				for seccondIndex+1 < len(chars) && chars[seccondIndex] == chars[seccondIndex+1] {
//					count++
//					seccondIndex++
//				}
//				res = append(res, chars[firstIndex])
//				countStr = strconv.Itoa(count)
//				if count > 10 {
//					for i := range countStr {
//						res = append(res, countStr[i])
//					}
//				} else {
//					res = append(res, countStr[0])
//				}
//			}
//			if count == 1 {
//				res = append(res, chars[firstIndex])
//				firstIndex++
//			} else {
//				count = 1
//				firstIndex = seccondIndex
//			}
//		}
//		fmt.Println(res)
//		return len(chars)
//	}
//func compress(chars []byte) int {
//	count := 1
//	var res []byte
//	var countStr string
//	for i := 0; i < len(chars); i++ {
//		if i < len(chars)-1 && chars[i] == chars[i+1] {
//			count++
//		} else {
//			res = append(res, chars[i])
//			if count > 1 {
//				countStr = strconv.Itoa(count)
//				if count > 10 {
//					for i := range countStr {
//						res = append(res, countStr[i])
//					}
//				} else {
//					res = append(res, countStr[0])
//				}
//			}
//		}
//	}
//	chars = res
//	fmt.Println(chars)
//	fmt.Println(r)
//	return len(chars)
//}

func compress(chars []byte) int {
	count := 1
	writeIndex := 0
	for i := 0; i < len(chars); i++ {
		if i < len(chars)-1 && chars[i] == chars[i+1] {
			count++
		} else {
			chars[writeIndex] = chars[i]
			writeIndex++
			if count > 1 {
				countStr := strconv.Itoa(count)
				for j := range countStr {
					chars[writeIndex] = countStr[j]
					writeIndex++
				}
			}
			count = 1
		}
	}
	fmt.Println(chars)
	return writeIndex
}

func main() {
	a := []byte{97, 97, 98, 98, 99, 99, 99}
	b := []byte{97, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98}
	fmt.Println(compress(a))
	fmt.Println(compress(b))
}
