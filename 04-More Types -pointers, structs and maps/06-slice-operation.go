package main

import "fmt"

/*
------------------------------------------------------------
- title: 06-slice-operation.go
- Author: Rakesh Niraula
- date: 2020-04-14 15:27:23
------------------------------------------------------------
*/

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {

	// slice boundary
	even := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	fmt.Println("Even: ", even)

	fmt.Println("even[:]", even[:])
	fmt.Println("even[0:]: ", even[0:])
	fmt.Println("event[:10]", even[:9])

	// making slices.go
	a := make([]int, 5)
	printSlice("a: ", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// appending in slice
	d = append(d, 1, 1, 1)
	printSlice("d", d)

	// range
	var pow = []int{1, 2, 4, 8, 16, 32}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
