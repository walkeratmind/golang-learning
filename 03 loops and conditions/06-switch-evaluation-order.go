package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Where's Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Saturday is Today")
	case today + 1:
		fmt.Println("Saturday is Tomorrow")
	case today + 2:
		fmt.Println("Day after Tomorrow is Saturday")
	default:
		fmt.Println("Not coming for 2 days: see What's next")
	}
}
