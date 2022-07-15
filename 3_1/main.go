package main

import (
	"fmt"
)

// Написать функцию itoa (integer to ascii),
// которая принимает на вход целое число и возвращает строку с этим же числом

func main() {
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(1))
	fmt.Println(Itoa(123))
	fmt.Println(Itoa(-1))
	fmt.Println(Itoa(-123))
}

func abs(number int) (int, bool) {
	if number < 0 {
		return -number, true
	}

	return number, false
}

func Itoa(number int) string {
	result := ""
	remainder, hasSign := abs(number)

	for {
		lastNumber := remainder % 10
		lastNumberAsString := string('0' + rune(lastNumber))
		result = lastNumberAsString + result

		remainder = remainder / 10
		if remainder == 0 {
			break
		}
	}

	if hasSign {
		result = "-" + result
	}

	return result
}
