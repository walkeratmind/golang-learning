package main

import "fmt"

/*
------------------------------------------------------------
- title: 01-pointers.go
- Author: Rakesh Niraula
- copyright (C) 2020. All rights reserved
- date: 2020-04-13 11:35:47
------------------------------------------------------------
*/

func main() {
	i, j := 7, 13

	p := &i         // point to i
	fmt.Println(*p) // read i through pointer

	*p = 21 // sets value of i through pointer
	fmt.Println(i)

	p = &j // point to j
	fmt.Println(*p)
	fmt.Println("Is value of *p == j: ", *p == j)
}
