package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func main() {

	fmt.Println(sqrt(16))
	fmt.Println(sqrt(-16))

	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
