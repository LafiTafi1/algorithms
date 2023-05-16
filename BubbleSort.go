package main

import (
	"fmt"
)

func main() {
	slice := []int{92, 1, 3, 4, 5, 6, 7, 12, 421}
	fmt.Println(bubbleSort(slice))
}

func bubbleSort(slice []int) []int {
	lenght := len(slice) - 1
	for i := 0; i < lenght; i++ {
		noSwaps := true
		for j := 0; j < lenght-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				noSwaps = false
			}
		}
		if noSwaps {
			return slice
		}
	}
	return slice
}
