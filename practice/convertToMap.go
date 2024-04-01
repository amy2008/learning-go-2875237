package main

import "fmt"

// Write your answer here, and then test your code.

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	// Your code goes here

	i := 0
	num := len(items)
	for k := range items {
		result[items[k]] = float64(100 / num)
		i++
	}
	return result
}

func main() {
	slice := []string{"apple", "banana", "orange", "date", "pear", "cherry", "mongo"}
	result := convertToMap(slice)
	fmt.Println(result)

}
