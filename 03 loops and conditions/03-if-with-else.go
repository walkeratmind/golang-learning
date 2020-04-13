package main

import "fmt"

func main() {

	var city = "kathmandu"

	if city == "pokhara" {
		fmt.Println("City: ", city)
	} else {
		fmt.Println("Outside of kathmandu")
	}
}
