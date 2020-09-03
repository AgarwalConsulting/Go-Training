package fizzbuzz

import "strconv"

// FizzBuzz fizz buzzes the numbers
func FizzBuzz(number int) string {
	isDivisibleBy3 := number%3 == 0
	isDivisibleBy5 := number%5 == 0

	if isDivisibleBy3 && isDivisibleBy5 {
		return "FizzBuzz"
	} else if isDivisibleBy3 {
		return "Fizz"
	} else if isDivisibleBy5 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}
