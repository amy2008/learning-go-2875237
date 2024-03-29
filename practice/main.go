package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 3, 4, 5
	intSum := i1 + i2 + i3
	fmt.Println("Integre Sum: ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum: ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float Sum: ", floatSum)

	circleRadius := 15.5
	circleference := circleRadius * 2 * math.Pi
	fmt.Printf("Circleference: %.4f\n", circleference) //keep the 4 demical numbers

}
