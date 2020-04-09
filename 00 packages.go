package main

import (
	"fmt"
	"math/rand"
)

/*
------------------------------------------------------------
	Note: The environment in which these programs are executed is deterministic,
	so each time you run the example program rand.Intn will return the same number.

	(To see a different number, seed the number generator;
	see rand.Seed. Time is constant in the playground,
	so you will need to use something else as the seed.)
------------------------------------------------------------
*/

func main() {
	fmt.Println(rand.Intn(100))
}
