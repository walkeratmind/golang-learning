package main

import (
	"fmt"
	"math"
)

/*
------------------------------------------------------------
- title: 17-exercise-errors.go
- Author: Rakesh Niraula
- date: 2020-04-14 16:39:59
------------------------------------------------------------
*/

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt. negative number: %g", float64(e))
}

func Sqrt(a float64) (float64, error) {
	if a > 0 {
		return math.Sqrt(a), ErrNegativeSqrt(a)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
