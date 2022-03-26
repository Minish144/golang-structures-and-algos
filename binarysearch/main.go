package main

import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, val int) bool {
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] == val {
			return true
		}
		if array[median] < val {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return false
}

func main() {
	array := []int{5, 10, 2, 3, 9, 12, 55, 70, 8, 15, 99}
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	fmt.Println(array)

	for _, val := range array {
		fmt.Println(BinarySearch(array, val))
	}
	fmt.Println(BinarySearch(array, 1))
	fmt.Println(BinarySearch(array, 0))
}
