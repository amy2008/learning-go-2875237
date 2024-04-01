// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	/* carts := make([]cartItem, 0, 20)

	// Your code goes here.
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	for decoder.More() {
		err := decoder.Decode(&cart)
		if err != nil {
			panic(err)
		}
		carts = append(carts, cart...)
		return carts
	} */
	err := json.Unmarshal([]byte(jsonString), &cart) // From the instructor David.
	if err != nil {
		panic(err)
	}
	return cart
}

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
	{"name":"orange","price":1.50,"quantity":8},
	{"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)
	fmt.Println(result)
}
