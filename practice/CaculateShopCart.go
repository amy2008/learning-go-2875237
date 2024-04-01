//package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = true
const showHints = true

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	var sum float64
	for _, fruit := range cart {
		fmt.Println(fruit)
		sum += fruit.price * float64(fruit.quantity)
	}
	return sum
}

func main() {
	// This is how your code will be called.
	// Your answer should be the largest value in the numbers array.
	// You can edit this code to try different testing cases.
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 2.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Println(result)

}
