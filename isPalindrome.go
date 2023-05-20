package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "rAdaR"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {
	newStr := ""
	for _, v := range s {
		newStr = string(v) + newStr
	}

	if strings.ToLower(s) == strings.ToLower(newStr) {
		return true
	}

	return false
}
