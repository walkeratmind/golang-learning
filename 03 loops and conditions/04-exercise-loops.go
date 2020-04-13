package main

import (
	"fmt"
	"math"
)

// Given a number x, we want to find the number z for which z² is most nearly x.
// Input: Number x
// Output: Number z for which z square is most nearly to x

// z -= (z*z - x) / (2*z)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
func main() {
	fmt.Println("Sq. root of 81: ", math.Sqrt(81))

	fmt.Println(Sqrt(81))
}
