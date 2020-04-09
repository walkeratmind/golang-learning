package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

// Naked Return : with named return
func split(sum int) (x, y int) {
	x = sum * 4 / 5
	y = sum - 3
	return // this is an eg. of naked return
}

func main() {
	fmt.Println("Function Demo:")
	c := add(3, 4)
	fmt.Println(c)

	fmt.Println(swap("Hari", "Ram"))

	fmt.Println(split(5))
}
