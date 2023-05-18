package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 5, 12, 145, 234, 1013, 12342}
	fmt.Println(binarySearch(nums, 1013))
}

func binarySearch(slice []int, n int) bool {
	first := 0
	last := len(slice) - 1
	for last >= first {
		mid := (first + last) / 2
		if slice[mid] == n {
			return true
		} else {
			if n > slice[mid] {
				first = mid + 1
			} else {
				last = mid - 1
			}
		}
	}
	return false
}
