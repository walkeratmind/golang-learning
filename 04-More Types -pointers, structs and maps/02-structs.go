package main

import "fmt"

/*
------------------------------------------------------------
- title: 02-structs.go
- Author: Rakesh Niraula
- date: 2020-04-13 11:40:10
------------------------------------------------------------
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 3})

	v := Vertex{1, 8}
	v.X = 7
	// access structs fields using dot as Struct.<Field>
	fmt.Println(v.X)
}
