package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

/*
------------------------------------------------------------
- title: 07-exercise-slices.go
- Author: Rakesh Niraula
- date: 2020-04-14 15:42:30
------------------------------------------------------------
*/

func Pic(dx, dy int) [][]uint8 {
	fmt.Printf("%d x %d\n\n", dx, dy)

	pixels := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		pixels[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pixels[y][x] = uint8(x * y)
		}
	}
	return pixels
}
func main() {
	pic.Show(Pic)
}
