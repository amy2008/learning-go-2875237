package main

import (
	"fmt"
)

func main() {
	sum := addValues(4, 9)
	fmt.Println("The sum is ", sum)

	multiSum, count := multiValues(4, 9, 2, 5)
	fmt.Println("The multiSum is ", multiSum)
	fmt.Println("The count is ", count)
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func multiValues(values ...int) (int, int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum, len(values)
}
