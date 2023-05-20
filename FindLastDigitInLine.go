package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	str := "This is string #1.But the last digit is 2"
	fmt.Println(lastDigit(str))
}

func lastDigit(str string) int {
	re := regexp.MustCompile("[0-9]+")
	allDigits := re.FindAllString(str, -1)
	lastElem := len(allDigits) - 1

	toInt, err := strconv.Atoi(allDigits[lastElem])
	if err != nil {
		log.Fatal(err)
	}
	return toInt
}
