package main

import "fmt"

/*
------------------------------------------------------------
- title: 03-struct-pointers.go
- Author: Rakesh Niraula
- date: 2020-04-14 15:13:59
------------------------------------------------------------
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 9}
	p := &v // pointer for vertex v

	fmt.Println("V: ", v)

	p.X = 1e7

	fmt.Println("V: ", v)
}
