package main

import "fmt"

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
func main() {
	array := []int{5, 10, 2, 3, 9, 12, 55, 70, 8, 15, 99}
	array2 := []int{14, 17, 8, 33, 99, 0}
	fmt.Println(merge(array, array2))
}
