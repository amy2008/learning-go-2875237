// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	var result float64
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operator!")
	}
	return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic("Invalid number.")
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	fmt.Println(calculate("8", "3", "+"))
	fmt.Println(calculate("8", "3", "-"))
	fmt.Println(calculate("8", "3", "*"))
	fmt.Println(calculate("8", "2", "/"))
	//fmt.Println(calculate("a", "b", "+"))
	fmt.Println(calculate("8", "2", "&"))
}
