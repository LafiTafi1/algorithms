package main

import "fmt"

func main() {
	nums := []int{2, 1, 6, 5, 3, 4} // 1 2 3 5 6
	fmt.Println(insertionSort(nums))
}

func insertionSort(slice []int) []int {

	for i := 1; i < len(slice); i++ {
		for j := i; j != 0 && slice[j-1] > slice[j]; j-- {
			slice[j-1], slice[j] = slice[j], slice[j-1]
		}
	}
	return slice
}
