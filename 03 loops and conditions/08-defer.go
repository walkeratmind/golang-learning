package main

import "fmt"

/*
------------------------------------------------------------
- title: 08-defer.go
- Author: Rakesh Niraula
- date: 2020-04-13 10:59:05
------------------------------------------------------------
*/

func main() {
	defer fmt.Println("Hello World")

	fmt.Println("Greet")
}
