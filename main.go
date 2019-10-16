package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(input int) string {

	result := ""

	if input%3 == 0 {
		result += "Fizz"
	}

	if input%5 == 0 {
		result += "Buzz"
	}

	if result != "" {
		return result
	}
	return strconv.Itoa(input)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(strconv.Itoa(i)+" : "+FizzBuzz(i))
	}

}
