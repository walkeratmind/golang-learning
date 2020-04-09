package main

import "fmt"

func fibonacci(limit int) {
	a := 0
	b := 1
	next := 0

	for i := 0; i <= limit; i++ {
		fmt.Printf("%v\t", a)
		next = a + b
		a = b
		b = next
	}
	fmt.Printf("\n")
}

func main() {

	// sum := 0 // similar to, sum int = 0, or var sum int = 0
	var sum int = 0
	// var x int = 3

	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	/*
		------------------------------------------------------------
			In Go, for == while
		------------------------------------------------------------
	*/
	sum = 1
	for sum < 10 {
		sum += sum
	}

	/*
		------------------------------------------------------------
			Infinite Loop
		------------------------------------------------------------
		for {

		}
	*/

	fmt.Println("sum:", sum)
	fibonacci(20)
}
