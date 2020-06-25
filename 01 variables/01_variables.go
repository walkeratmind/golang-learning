package main

/*
------------------------------------------------------------
- title: 02_variables.go
- Author: Rakesh Niraula
- copyright (C) 2020. All rights reserved
- date: 2020-04-09 16:27:25
------------------------------------------------------------
*/

import (
	"fmt"
	"math"
	"math/cmplx"
)

// var statement declares list of variables,
// var statement can be at package level or function level.
var c, python, java bool

// var with initializers
var a, b int = 1, 2

func main() {
	var i int

	var (
		isGo   bool       = true
		maxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const pi = math.Pi
	var rust, godot = true, "!no"
	fmt.Println(i, c, python, java)
	fmt.Println(a, b)
	fmt.Println(rust, godot)

	fmt.Printf("Type: %T Value: %v\n", isGo, isGo)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("Pi: %v\n", pi)
}
