package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "arc Car"
	s2 := "car Arc"
	fmt.Println(isAnagram(s1, s2))
}

func isAnagram(s1, s2 string) bool {
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	slice1 := []byte(s1)
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })

	slice2 := []byte(s2)
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })

	if bytes.Equal(slice1, slice2) {
		return true
	}
	return false
}
