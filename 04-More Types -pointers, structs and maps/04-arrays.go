package main

import "fmt"

/*
------------------------------------------------------------
- title: 04-arrays.go
- Author: Rakesh Niraula
- date: 2020-04-14 15:17:23
------------------------------------------------------------
*/

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Universe"

	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes: ", primes)

	// slicing an array
	primesBelowTen := primes[0:4]
	fmt.Println("PrimesBelowTen: ", primesBelowTen)
}
