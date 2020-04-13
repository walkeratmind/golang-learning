package main

import "fmt"

/*
------------------------------------------------------------
- title: 09-defer-multi.go
- Author: Rakesh Niraula
- date: 2020-04-13 11:00:52
------------------------------------------------------------
*/

func main() {
	fmt.Println("Counting: ")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done in reverse count with defer")
}
