package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(input int) string {

	if Contains(input, "3") && Contains(input, "5") {
		return "FizzBuzz"
	}

	if Contains(input, "3") {
		return "Fizz"
	}

	if Contains(input, "5") {
		return "Buzz"
	}

	if input%15 == 0 {
		return "FizzBuzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}

	if input%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(input)
}

func Contains(input int, s string) bool {
	return strings.Contains(strconv.Itoa(input), s)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(strconv.Itoa(i) + " : " + FizzBuzz(i))
	}
}
