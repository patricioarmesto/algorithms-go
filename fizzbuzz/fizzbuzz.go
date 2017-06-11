package fizzBuzz

import "strconv"

/*
FizzBuzz Example
*/
func FizzBuzz(val int) string {

	var result string

	if val%15 == 0 {
		result = "FizzBuzz"
	} else if val%3 == 0 {
		result = "Fizz"
	} else if val%5 == 0 {
		result = "Buzz"
	} else {
		result = strconv.Itoa(val)
	}
	return result
}
