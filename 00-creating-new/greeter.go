package main

import (
	"fmt"
	"time"
)

/*
------------------------------------------------------------
- title: greeter.go
- Author: Rakesh Niraula
- date: 2020-04-13 10:47:31
------------------------------------------------------------
*/

func main() {
	fmt.Println("--------------------------")
	fmt.Println("--- Welcome to Greeter ---")
	fmt.Println("--------------------------")

	time := time.Now()

	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning, It's ", time.Clock())
	case time.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

}
