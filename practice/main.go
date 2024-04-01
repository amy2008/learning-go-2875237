package main

import (
	"fmt"
)

func main() {
	anInt := 42
	p := &anInt
	fmt.Println("Value of pointer: ", *p)
	fmt.Println("Address of pointer: ", p)

	value1 := 42.12
	pointer1 := &value1
	fmt.Println("Value 1: ", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 2: ", *pointer1)
}
