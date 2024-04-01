package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "10"
	value2 := "5.5  "
	result := calculate(value1, value2)
	fmt.Println("Float sum is ", result)
}

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	var float1, float2 float64

	// Convert the first string to a float64
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	// Convert the second string to a float64
	float2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err2 != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	// Calculate and return the result
	return float1 + float2
}
