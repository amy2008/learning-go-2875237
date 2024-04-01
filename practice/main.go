package main

import (
	"fmt"
	"sort"
)

func main() {
	//var colors = [3]string{"Red", "Green", "Blue"} //This is an array
	var colors = []string{"Red", "Green", "Blue"} //This is a slice, [] without a number
	fmt.Println(colors)
	colors = append(colors, "purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)]) //copy the part of a slice
	fmt.Println(colors)

	colors = append(colors, colors[1:len(colors)]...) //Myself practice: Connect 2 slices, ... is important
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1]) //Delete the last colors items
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 12123
	numbers[1] = 32
	numbers[2] = 53
	numbers[3] = 12
	numbers[4] = 283
	fmt.Println(numbers)

	numbers = append(numbers, 123)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
