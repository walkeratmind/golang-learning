package main

import "fmt"

/*
------------------------------------------------------------
- title: 05-slice-literals.go
- Author: Rakesh Niraula
- date: 2020-04-14 15:22:35
------------------------------------------------------------
*/

func main() {
	odd := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("odd: ", odd)

	r := []bool{true, false, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, false},
		{6, true},
	}

	fmt.Println(s)
}
