package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning")
	case time.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
