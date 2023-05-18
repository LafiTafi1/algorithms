package main

import "fmt"

func main() {
	nums := []int{91, 94, 1, 2, 3, 4, 120}
	fmt.Println(sort(nums))
}

func sort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := slice[:mid]
	right := slice[mid:]

	left = sort(left)
	right = sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
